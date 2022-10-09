//go:build wireinject
// +build wireinject

package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedUserService(db *gorm.DB) *UserService {
	wire.Build(NewUserService, repositories.NewUserRepository, wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)))

	return nil
}

func InitializedTweetService(db *gorm.DB) *TweetService {

	wire.Build(
		repositories.NewTweetRepository,
		repositories.NewTweetLikeRepository,
		repositories.NewUserRepository,
		NewTweetService,
		wire.Bind(new(repositories.ITweetRepository), new(*repositories.TweetRepository)),
		wire.Bind(new(repositories.ITweetLikeRepository), new(*repositories.TweetLikeRepository)),
		wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)),
	)

	return nil
}

func InitializedSessionService(db *gorm.DB) *SessionService {
	wire.Build(
		repositories.NewSessionRepository,
		repositories.NewUserRepository,
		NewSessionService,
		wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)),
		wire.Bind(new(repositories.ISessionRepository), new(*repositories.SessionRepository)),
	)

	return nil
}
