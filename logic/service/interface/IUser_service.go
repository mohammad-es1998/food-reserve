package service

import "food-reserve/db/model"

type IUserService interface {
	Login(username, password string) (*model.User, error)
	Register(username, password, role string) error
}
