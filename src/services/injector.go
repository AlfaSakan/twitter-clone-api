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
		repositories.NewRetweetRepository,
		NewTweetService,
		wire.Bind(new(repositories.ITweetRepository), new(*repositories.TweetRepository)),
		wire.Bind(new(repositories.ITweetLikeRepository), new(*repositories.TweetLikeRepository)),
		wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)),
		wire.Bind(new(repositories.IRetweetRepository), new(*repositories.RetweetRepository)),
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

func InitializedReplyService(db *gorm.DB) *ReplyService {
	wire.Build(
		NewReplyService,
		repositories.NewReplyRepository,
		wire.Bind(new(repositories.IReplyRepository), new(*repositories.ReplyRepository)),
	)

	return nil
}

func InitializedRetweetService(db *gorm.DB) *RetweetService {
	wire.Build(
		NewRetweetService,
		repositories.NewRetweetRepository,
		wire.Bind(new(repositories.IRetweetRepository), new(*repositories.RetweetRepository)),
	)

	return nil
}
