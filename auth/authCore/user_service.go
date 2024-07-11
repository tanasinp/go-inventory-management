package authcore

import (
	"github.com/tanasinp/go-inventory-management/database"
)

type UserService interface {
	CreateUser(user *database.User) error
	LoginUser(user *database.User) (string, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user *database.User) error {
	if err := s.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) LoginUser(user *database.User) (string, error) {
	t, err := s.repo.LoginUser(user)
	if err != nil {
		return "", err
	}
	return t, nil
}
