package service

import (
	"errors"
	"food-reserve/db/model"
	"food-reserve/db/repository"
	service "food-reserve/logic/service/interface"
	"food-reserve/logic/utils"
	bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	uow repository.IUnitOfWork
}

func NewUserService(uow repository.IUnitOfWork) service.IUserService {
	return &userService{uow: uow}
}

func (s *userService) Login(username, password string) (*model.User, error) {
	user, err := s.uow.UserRepository().GetByUsername(username)
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
	accessRole, err := s.uow.RoleRepository().GetByName(role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(utils.RoleNotFound)
		}
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     model.Role{Model: gorm.Model{ID: accessRole.ID}},
	}

	return s.uow.UserRepository().Create(user)
}
