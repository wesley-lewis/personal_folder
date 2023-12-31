package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/wesley-lewis/personal_folder/handlers"
)

func main() {
	server := gin.Default()
	fmt.Println("Personal folder")

	server.StaticFile("/", "./public/index.html")

	server.GET("/get-all-users", handler.GetUsersHandler)
	server.POST("/register", handler.RegisterHandler)
	server.POST("upload-file", handler.UploadFile)

	server.Run(":8080")
}
