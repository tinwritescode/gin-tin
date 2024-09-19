package service

import (
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
)

type UserService interface {
	GetUsers() ([]model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) GetUsers() ([]model.User, error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}
