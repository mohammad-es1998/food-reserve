package service

import (
	"errors"
	"food-reserve/db/model"
	repository "food-reserve/db/repository/interface"
	service "food-reserve/logic/service/interface"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) service.IUserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Login(username, password string) (*model.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *userService) Register(username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     model.Role{Name: role},
	}

	return s.userRepo.Create(user)
}
