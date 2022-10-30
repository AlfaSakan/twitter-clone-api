package services

import (
	"net/http"
	"time"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/AlfaSakan/twitter-clone-api/src/schemas"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	FindUser(user *entities.User) (statusCode int, errorMessage error)
	CreateUser(user *entities.User) (statusCode int, errorMessage error)
	UpdateUser(user *schemas.UpdateUserSchema, id string) error
	DeleteUser(user *entities.User) error
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) FindUser(user *entities.User) (int, error) {
	status, err := s.userRepository.FindUser(user)
	user.Password = ""

	return status, err
}

func (s *UserService) CreateUser(user *entities.User) (int, error) {
	user.Id = helpers.GenerateId()

	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user.Password = string(hashedPassword)

	status, err := s.userRepository.CreateUser(user)
	user.Password = ""

	return status, err
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
