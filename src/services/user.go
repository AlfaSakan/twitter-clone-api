package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
)

type IUserService interface {
	FindUser(user *entities.User) error
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User, id string) error
	DeleteUser(user *entities.User) error
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) FindUser(user *entities.User) error {
	return s.userRepository.FindUser(user)
}

func (s *UserService) CreateUser(user *entities.User) error {
	user.Id = helpers.GenerateId()

	return s.userRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(user *entities.User, id string) error {
	return s.userRepository.UpdateUser(user, id)
}

func (s *UserService) DeleteUser(user *entities.User) error {
	return s.userRepository.DeleteUser(user)
}
