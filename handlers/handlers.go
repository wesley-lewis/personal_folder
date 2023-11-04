package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/wesley-lewis/personal_folder/models"
)

var usersData []model.UserData

func RegisterHandler(c *gin.Context) {
	var user model.UserData

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all details must be provided"})
	}

	usersData = append(usersData, user)

	c.JSON(http.StatusOK, gin.H{"message": "user has been added"})
}

func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": usersData})
}
