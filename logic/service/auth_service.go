package service

import (
	"errors"
	repository "food-reserve/db/repository/interface"
	service "food-reserve/logic/service/interface"
	"food-reserve/logic/utils"
)

type authService struct {
	authRepo repository.IAuthRepository
}

func NewAuthService(authRepo repository.IAuthRepository) service.IAuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) CheckPermission(token string, requiredPermission string) (*utils.Claims, error) {

	claims, err := utils.ValidateToken(token)
	if err != nil {
		return nil, errors.New(utils.InvalidToken)
	}

	user, err := s.authRepo.GetUser(claims.Username)
	if err != nil {
		return nil, errors.New(utils.UserNotFound)
	}

	if !user.Role.HasPermission(requiredPermission) {
		return nil, errors.New(utils.Forbidden)
	}
	return claims, nil

}
