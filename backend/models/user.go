package models

import (
	"time"

	"golang.org/x/crypto/bcrypt" // Importing the bcrypt package for hashing passwords
)

type User struct {
	ID                uint      `gorm:"primaryKey"`       // Explicitly define ID field and set it as the primary key
	Username          string    `gorm:"unique;not null"`  // Unique and non-nullable username
	Email             string    `gorm:"unique;not null"`  // Unique and non-nullable email
	Password          string    `gorm:"not null"`         // Non-nullable password field
	Bio               string    `gorm:"type:text"`        // Bio can be any length (stored as text)
	Interests         string    `gorm:"type:text"`        // Interests can be a large text field
	ProfilePictureURL string    `gorm:"type:text"`        // URL of the profile picture (stored as text)
	AgeRange          string    `gorm:"type:varchar(50)"` // Age range of the user (string of max length 50)
	Distance          int       `gorm:"type:int"`         // Distance preference (in integer)
	GenderPreference  string    `gorm:"type:varchar(50)"` // Gender preference for potential matches
	CreatedAt         time.Time `gorm:"autoCreateTime"`   // Automatically set when a record is created
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`   // Automatically update when a record is modified
}

// LoginRequest is the structure for login requests containing username and password
type LoginRequest struct {
	Username string `json:"username"` // Username for login
	Password string `json:"password"` // Password for login
}

// UpdateProfileRequest is used to update user profile details (bio and interests)
type UpdateProfileRequest struct {
	Bio               string `json:"bio"`
	Interests         string `json:"interests"`
	ProfilePictureURL string `json:"profile_picture"`
	AgeRange          string `json:"age_range"`
	Distance          int    `json:"distance"`
	GenderPreference  string `json:"gender_preference"`
}

// UpdatePreferencesRequest is used to update user preferences like age range, distance, and gender preference
type UpdatePreferencesRequest struct {
	AgeRange         string `json:"age_range"`         // Age range preference
	Distance         int    `json:"distance"`          // Distance preference
	GenderPreference string `json:"gender_preference"` // Gender preference for matches
}

// HashPassword hashes the provided password using bcrypt and stores it in the User struct
func (u *User) HashPassword(password string) error {
	// Generate a bcrypt hash of the password with a cost of 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err // Return an error if password hashing fails
	}
	u.Password = string(bytes) // Store the hashed password in the User struct
	return nil                 // Return nil if hashing is successful
}

// CheckPassword compares the provided password with the stored hashed password
func (u *User) CheckPassword(password string) error {
	// Compare the hashed password in the User struct with the provided password
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
