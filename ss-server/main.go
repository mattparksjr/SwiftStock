package main

import (
	"log"
	"swiftstock_api/app/controller"
	"swiftstock_api/app/database"
	"swiftstock_api/app/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()
	loadDatabase()
	serve()
}

func serve() {

	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	router.Run() // listen and serve on 0.0.0.0:8080
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("Error loading .env file. " + err.Error())
	}
}