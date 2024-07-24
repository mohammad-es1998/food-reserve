package service

import "food-reserve/logic/utils"

type IAuthService interface {
	CheckPermission(token string, requiredPermission string) (*utils.Claims, error)
}
