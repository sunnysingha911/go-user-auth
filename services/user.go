package services

import (
	"github.com/sunnysingha911/user-service/database"
	"github.com/sunnysingha911/user-service/models"
	"github.com/sunnysingha911/user-service/utils"
)

type UserResponse struct {
	Users []*models.User
	Total int64
	Page  int
	Limit int
}

func GetUserList(page, limit int) (*UserResponse, error) {
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

	return &UserResponse{
		Users: users,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}
