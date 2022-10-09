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

func InitializedTweetHandler(us services.ITweetService) *TweetHandler {
	wire.Build(NewTweetHandler)

	return nil
}

func InitializedSessionHandler(us services.IUserService, ss services.ISessionService) *SessionHandler {
	wire.Build(NewSessionHandler)

	return nil
}
