package repository

import "food-reserve/db/model"

type IRoleRepository interface {
	GetByName(roleName string) (*model.Role, error)
}
