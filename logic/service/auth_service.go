package service

import (
	"errors"
	"food-reserve/db/repository"
	service "food-reserve/logic/service/interface"
	"food-reserve/logic/utils"
)

type authService struct {
	uow repository.IUnitOfWork
}

func NewAuthService(uow repository.IUnitOfWork) service.IAuthService {
	return &authService{uow: uow}
}

func (s *authService) CheckPermission(token string, requiredPermission string) (*utils.Claims, error) {

	claims, err := utils.ValidateToken(token)
	if err != nil {
		return nil, errors.New(utils.InvalidToken)
	}

	user, err := s.uow.AuthRepository().GetUser(claims.Username)
	if err != nil {
		return nil, errors.New(utils.UserNotFound)
	}

	if !user.Role.HasPermission(requiredPermission) {
		return nil, errors.New(utils.Forbidden)
	}
	return claims, nil

}
