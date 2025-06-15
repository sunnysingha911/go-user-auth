package services

import (
	"github.com/sunnysingha911/user-service/database"
	"github.com/sunnysingha911/user-service/models"
	"github.com/sunnysingha911/user-service/utils"
)

type UserListResponse struct {
	Users []*models.User
	Total int64
	Page  int
	Limit int
}

type UserResponse struct {
	User *models.User
}

func GetUserList(page, limit int) (*UserListResponse, error) {
	var users []*models.User
	var total int64

	db := database.DB

	// Count total users
	if err := db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// Fetch paginated users
	if err := db.Scopes(utils.Paginate(page, limit)).Find(&users).Error; err != nil {
		return nil, err
	}

	return &UserListResponse{
		Users: users,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func GetUserById(userId string) (*UserResponse, error) {
	var user models.User

	db := database.DB

	if err := db.First(&user, "id = ?", userId).Error; err != nil {
		return nil, err
	}

	return &UserResponse{
		User: &user,
	}, nil
}
