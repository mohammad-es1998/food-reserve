package repository

import (
	"errors"
	"food-reserve/db/model"
	"food-reserve/logic/utils"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUser(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New(utils.UserNotFound)
	}

	return &user, nil
}
