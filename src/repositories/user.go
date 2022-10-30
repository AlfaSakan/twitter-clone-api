package repositories

import (
	"net/http"
	"strings"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindUser(user *entities.User) (statusCode int, errorMessage error)
	CreateUser(user *entities.User) (statusCode int, errorMessage error)
	UpdateUser(user *entities.User, id string) error
	DeleteUser(user *entities.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindUser(user *entities.User) (int, error) {
	err := r.db.Where(user).First(user).Error

	if err.Error() == "record not found" {
		return http.StatusNotFound, helpers.UserNotFound.With(user.Id)
	}

	return http.StatusBadRequest, err
}

func (r *UserRepository) CreateUser(user *entities.User) (int, error) {
	err := r.db.Create(user).Error

	if strings.Contains(err.Error(), "username") {
		return http.StatusBadRequest, helpers.DuplicateUser.From(user.Username, err)
	}

	if strings.Contains(err.Error(), "email") {
		return http.StatusBadRequest, helpers.DuplicateUser.From(user.Email, err)
	}

	return http.StatusCreated, err
}

func (r *UserRepository) UpdateUser(user *entities.User, id string) error {
	return r.db.Where(&entities.User{Id: id}).Updates(user).Error
}

func (r *UserRepository) DeleteUser(user *entities.User) error {
	return r.db.Where(user).Delete(user).Error
}
