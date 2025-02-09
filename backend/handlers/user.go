package handlers

import (
	"datingapp/database"
	"datingapp/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Register handles user registration
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {
	var user models.User

	// Bind JSON request body to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// If error occurs in binding, return bad request with error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password before storing it in the database
	if err := user.HashPassword(user.Password); err != nil {
		// Return an internal server error if password hashing fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create the user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		// Return an internal server error if user creation fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Return success message upon successful user registration
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login authenticates a user and returns a JWT token
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var input models.LoginRequest

	// Bind JSON request body to input struct (login credentials)
	if err := c.ShouldBindJSON(&input); err != nil {
		// If error occurs in binding, return bad request with error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve user from the database based on the provided username
	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		// If user not found, return unauthorized error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the provided password matches the stored hashed password
	if err := user.CheckPassword(input.Password); err != nil {
		// If password check fails, return unauthorized error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a JWT token with user ID and expiration time (24 hours)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the JWT secret key from environment variables
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		// Return an internal server error if token generation fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the generated token in the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
