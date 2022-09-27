//go:build wireinject
// +build wireinject

package handlers

import (
	"github.com/AlfaSakan/twitter-clone-api/src/services"
	"github.com/google/wire"
)

func InitializedUserHandler(us services.IUserService) *UserHandler {
	wire.Build(NewUserHandler)

	return nil
}
