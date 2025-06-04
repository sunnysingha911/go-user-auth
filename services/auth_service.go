package services

import (
	"errors"

	"github.com/sunnysingha911/user-service/database"
	"github.com/sunnysingha911/user-service/models"
	"github.com/sunnysingha911/user-service/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type LoginResponse struct {
	User  *models.User
	Token string
}

// Login authenticates a user and returns user info + JWT token
func Login(email, password string) (*LoginResponse, error) {
	// Find user
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:  &user,
		Token: token,
	}, nil
}
