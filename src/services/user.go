package services

import (
	"time"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/AlfaSakan/twitter-clone-api/src/schemas"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	FindUser(user *entities.User) error
	CreateUser(user *entities.User) error
	UpdateUser(user *schemas.UpdateUserSchema, id string) error
	DeleteUser(user *entities.User) error
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) FindUser(user *entities.User) error {
	err := s.userRepository.FindUser(user)
	user.Password = ""

	return err
}

func (s *UserService) CreateUser(user *entities.User) error {
	user.Id = helpers.GenerateId()

	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = s.userRepository.CreateUser(user)
	user.Password = ""

	return err
}

func (s *UserService) UpdateUser(schema *schemas.UpdateUserSchema, id string) error {
	user := entities.User{
		Username:  schema.Username,
		Name:      schema.Name,
		Email:     schema.Email,
		UpdatedAt: time.Now().UnixMilli(),
	}

	return s.userRepository.UpdateUser(&user, id)
}

func (s *UserService) DeleteUser(user *entities.User) error {
	return s.userRepository.DeleteUser(user)
}
