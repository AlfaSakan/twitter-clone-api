package repositories

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"gorm.io/gorm"
)

type IRetweetRepository interface {
	CreateRetweet(retweet *entities.Retweet) error
	FindRetweetId(retweet *entities.Retweet) error
	FindRetweetsId(tweetId string, retweets *[]entities.Retweet) error
}

type RetweetRepository struct {
	db *gorm.DB
}

func NewRetweetRepository(db *gorm.DB) *RetweetRepository {
	return &RetweetRepository{db}
}

func (rr *RetweetRepository) CreateRetweet(retweet *entities.Retweet) error {
	return rr.db.Create(retweet).Error
}

func (rr *RetweetRepository) FindRetweetId(retweet *entities.Retweet) error {
	return rr.db.Debug().Where(retweet).First(retweet).Error
}

func (rr *RetweetRepository) FindRetweetsId(
	tweetId string,
	retweets *[]entities.Retweet,
) error {
	return rr.db.Where(&entities.Retweet{TweetId: tweetId}).Find(retweets).Error
}
