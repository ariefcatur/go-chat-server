package utils

import (
	"go-chat-server/config"
	"go-chat-server/models"
)

// UserExists mengecek apakah username sudah terdaftar di database
func UserExists(username string) bool {
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)
	return result.RowsAffected > 0
}

// SaveUser menyimpan user baru ke database
func SaveUser(user models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}
