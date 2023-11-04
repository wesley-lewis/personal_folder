package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/wesley-lewis/personal_folder/handlers"
)

func main() {
	server := gin.Default()
	fmt.Println("Personal folder")

	// server.Static("/", "./public")

	server.GET("/get-all-users", handler.GetUsersHandler)
	server.POST("/register", handler.RegisterHandler)

	server.Run(":8080")
}
