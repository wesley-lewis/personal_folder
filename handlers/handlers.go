package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/wesley-lewis/personal_folder/models"
)

var usersData []model.UserData

func RegisterHandler(c *gin.Context) {
	var user model.UserData
	// fmt.Println(c.PostForm("name"), c.PostForm("email"), c.PostForm("mobile_no"))
	if err := c.Bind(&user); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "all details must be provided"})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersData = append(usersData, user)

	c.JSON(http.StatusOK, gin.H{"message": "user has been added"})
}

func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": usersData})
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	c.SaveUploadedFile(file, "./files/"+file.Filename)
	c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully"})
}
