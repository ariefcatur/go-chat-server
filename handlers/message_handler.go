package handlers

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/config"
	"go-chat-server/helpers"
	"go-chat-server/models"
	"net/http"
)

func GetMessages(c *gin.Context) {
	var messages []models.Message
	config.DB.Find(&messages)
	//c.JSON(http.StatusOK, messages)
	helpers.APIResponse(c, http.StatusOK, "Success get messages", messages)
}

func SendMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		helpers.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	config.DB.Create(&msg)
	//c.JSON(http.StatusOK, msg)
	helpers.APIResponse(c, http.StatusOK, "Success send message", msg)
}
