package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
)

type IRetweetService interface {
	CreateRetweet(tweetId string, retweetId string) error
	FindRetweets(tweetId string, retweets *[]entities.Retweet) error
}

type RetweetService struct {
	retweetRepository repositories.IRetweetRepository
}

func NewRetweetService(retweetRepository repositories.IRetweetRepository) *RetweetService {
	return &RetweetService{retweetRepository}
}

func (rs *RetweetService) CreateRetweet(tweetId string, retweetId string) error {
	retweet := &entities.Retweet{
		TweetId:   tweetId,
		RetweetId: retweetId,
	}

	return rs.retweetRepository.CreateRetweet(retweet)
}

func (rs *RetweetService) FindRetweets(tweetId string, retweets *[]entities.Retweet) error {
	return rs.retweetRepository.FindRetweetsId(tweetId, retweets)
}
