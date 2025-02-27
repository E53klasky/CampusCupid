package main

// go run main.go
// swag init
// go mod tidy

import (
	"datingapp/database"
	"datingapp/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "datingapp/docs"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title CampusCupid API
// @version 1.0
// @description This is the backend API for the CampusCupid dating app.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@campuscupid.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize database connection
	database.Connect()

	// Create a new Gin router instance
	r := gin.Default()

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public authentication routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// new routes once logged in
	// profile info
	// profile info post and get
	r.GET("/profile/:user_id", handlers.GetUserProfile)

	// Start the server on port 8080
	r.Run(":8080")
}
