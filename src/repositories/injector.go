//go:build wireinject
// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedUserRepository(db *gorm.DB) *UserRepository {
	wire.Build(NewUserRepository)
	return nil
}
