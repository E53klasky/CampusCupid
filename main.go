package main

/*
Update main.go to Include Swagger Docs
In your main.go file, you need to set up Swagger documentation for each API endpoint. We will use comments to define metadata for each route.

Updated main.go with Swagger Annotations:
*/


package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files" 
	"github.com/swaggo/gin-swagger" 
	_ "github.com/E53klasky/CampusCupid/docs" // Import the generated docs
)

// User struct for data representation
// @Description User represents a user in our system
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory storage for users
var users = []User{
	{ID: "1", Name: "Alice", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", Email: "bob@example.com"},
}

// @title Go API with Swagger Docs
// @version 1.0
// @description This is a simple API that demonstrates integration with Swagger UI
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// API Routes
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	// Run server
	r.Run(":8080") 
}

// @Summary Get all users
// @Description Fetch all users from the system
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// @Summary Get a single user by ID
// @Description Fetch a user by their unique ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} gin.H{"message": "User not found"}
// @Router /users/{id} [get]
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// @Summary Create a new user
// @Description Add a new user to the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User data"
// @Success 201 {object} User
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Router /users [post]
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

// @Summary Update an existing user
// @Description Modify the details of an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body User true "Updated user data"
// @Success 200 {object} User
// @Failure 404 {object} gin.H{"message": "User not found"}
// @Router /users/{id} [put]
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// @Summary Delete a user
// @Description Remove a user by their ID
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} gin.H{"message": "User deleted"}
// @Failure 404 {object} gin.H{"message": "User not found"}
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
