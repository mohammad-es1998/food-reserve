package repository

import "food-reserve/db/model"

type IUserRepository interface {
	GetByUsername(username string) (*model.User, error)
	Create(user *model.User) error
}
