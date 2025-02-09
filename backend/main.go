package main

import (
	"datingapp/database"
	"datingapp/handlers"
	"datingapp/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "datingapp/docs"
)

func init() {
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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	database.Connect()

	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handlers.GetProfile)
		auth.PUT("/profile", handlers.UpdateProfile)
		auth.DELETE("/profile", handlers.DeleteProfile)
		auth.POST("/profile/picture", handlers.UploadProfilePicture)
		auth.PUT("/preferences", handlers.UpdatePreferences)
		auth.GET("/preferences", handlers.GetPreferences)
	}

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware()) // Add admin-specific middleware if needed
	{
		admin.GET("/users", handlers.GetAllUsers)
	}

	r.Run(":8080")
}
