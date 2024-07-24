package repository

import (
	"food-reserve/db/model"
)

type IAuthRepository interface {
	GetUser(username string) (*model.User, error)
}
