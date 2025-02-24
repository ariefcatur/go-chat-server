package handlers

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/config"
	"go-chat-server/helpers"
	"go-chat-server/models"
	"go-chat-server/utils"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid input"})
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid input", gin.H{
			"details": err.Error()})
		return
	}

	// Cek apakah username sudah digunakan
	if utils.UserExists(user.Username) {
		//c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Username sudah terdaftar"})
		helpers.ErrorResponse(c, http.StatusConflict, "Registration is failed", gin.H{
			"details": "Username already registered",
		})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengenkripsi password"})
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", gin.H{
			"details": err.Error(),
		})
		return
	}
	user.Password = hashedPassword

	// Simpan ke database
	err = utils.SaveUser(user)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menyimpan pengguna"})
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Adding new user is failed", gin.H{
			"details": err.Error(),
		})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"success": true, "message": "Registrasi berhasil"})
	helpers.APIResponse(c, http.StatusCreated, "Registration is success", user)
}

func Login(c *gin.Context) {
	var userInput models.User
	var user models.User

	// Ambil input JSON
	if err := c.ShouldBindJSON(&userInput); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid input", gin.H{
			"details": err.Error(),
		})
		return
	}

	// Cari user di database
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"message": "User tidak ditemukan"})
		helpers.ErrorResponse(c, http.StatusNotFound, "User not found", gin.H{
			"details": err.Error(),
		})
		return
	}

	// Cek password
	if !utils.CheckPassword(userInput.Password, user.Password) {
		//c.JSON(http.StatusUnauthorized, gin.H{"message": "Password salah"})
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Wrong password", gin.H{
			"details": "Invalid input",
		})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal generate token"})
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Generating token is failed", gin.H{
			"details": err.Error(),
		})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
	helpers.APIResponse(c, http.StatusOK, "success", gin.H{"success": true, "token": token})
}
