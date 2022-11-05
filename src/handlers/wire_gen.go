// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handlers

import (
	"github.com/AlfaSakan/twitter-clone-api/src/services"
)

// Injectors from injector.go:

func InitializedUserHandler(us services.IUserService) *UserHandler {
	userHandler := NewUserHandler(us)
	return userHandler
}

func InitializedTweetHandler(us services.ITweetService) *TweetHandler {
	tweetHandler := NewTweetHandler(us)
	return tweetHandler
}

func InitializedSessionHandler(us services.IUserService, ss services.ISessionService) *SessionHandler {
	sessionHandler := NewSessionHandler(ss, us)
	return sessionHandler
}

func InitializedReplyHandler(ts services.ITweetService, rs services.IReplyService) *ReplyHandler {
	replyHandler := NewReplyHandler(rs, ts)
	return replyHandler
}

func InitializedRetweetHandler(ts services.ITweetService, rs services.IRetweetService) *RetweetHandler {
	retweetHandler := NewRetweetHandler(rs, ts)
	return retweetHandler
}
