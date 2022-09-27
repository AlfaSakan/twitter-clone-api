package repositories

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindUser(user *entities.User) error
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User, id string) error
	DeleteUser(user *entities.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindUser(user *entities.User) error {
	return r.db.Debug().Find(user, user).Error
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	user.Id = helpers.GenerateId()
	return r.db.Debug().Create(user).Error
}

func (r *UserRepository) UpdateUser(user *entities.User, id string) error {
	return r.db.Debug().Where(&entities.User{Id: id}).Updates(user).Error
}

func (r *UserRepository) DeleteUser(user *entities.User) error {
	return r.db.Debug().Where(user).Delete(user).Error
}
