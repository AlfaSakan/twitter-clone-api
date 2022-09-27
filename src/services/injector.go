//go:build wireinject
// +build wireinject

package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedUserService(ur repositories.IUserRepository) *UserService {
	wire.Build(NewUserService)

	return nil
}
