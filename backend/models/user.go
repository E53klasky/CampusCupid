package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string `gorm:"unique;not null"`
	Email             string `gorm:"unique;not null"`
	Password          string `gorm:"not null"`
	Bio               string `gorm:"type:text"`
	Interests         string `gorm:"type:text"`
	ProfilePictureURL string `gorm:"type:text"`
	AgeRange          string `gorm:"type:varchar(50)"`
	Distance          int    `gorm:"type:int"`
	GenderPreference  string `gorm:"type:varchar(50)"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Bio       string `json:"bio"`
	Interests string `json:"interests"`
}

type UpdatePreferencesRequest struct {
	AgeRange         string `json:"age_range"`
	Distance         int    `json:"distance"`
	GenderPreference string `json:"gender_preference"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
