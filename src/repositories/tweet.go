package repositories

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"gorm.io/gorm"
)

type ITweetRepository interface {
	FindListTweets(tweetRequest *entities.Tweet, tweets *[]entities.Tweet) error
	FindTweet(tweet *entities.Tweet) error
	CreateTweet(tweet *entities.Tweet) error
	DeleteTweet(tweet *entities.Tweet) error
	FindTweetJoin(tweet *entities.Tweet) error
	UpdateTweet(tweet *entities.Tweet) error
	UpdateTweetLikes(request *entities.Tweet) error
	GetAllTweets(tweets *[]entities.Tweet) error
}

type TweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) *TweetRepository {
	return &TweetRepository{db}
}

func (r *TweetRepository) GetAllTweets(tweets *[]entities.Tweet) error {
	return r.db.Select("*").Order("created_at ASC").Find(tweets).Error
}

func (r *TweetRepository) FindListTweets(tweetRequest *entities.Tweet, tweets *[]entities.Tweet) error {
	return r.db.Order("created_at ASC").Find(tweets, tweetRequest).Error
}

func (r *TweetRepository) FindTweet(tweet *entities.Tweet) error {
	return r.db.Where(&entities.Tweet{Id: tweet.Id}).First(tweet).Error
}

func (r *TweetRepository) FindTweetJoin(tweet *entities.Tweet) error {
	return r.db.Select("*").Joins("JOIN users ON tweets.user_id = users.id").First(tweet).Error
}

func (r *TweetRepository) CreateTweet(tweet *entities.Tweet) error {
	return r.db.Create(tweet).Error
}

func (r *TweetRepository) DeleteTweet(tweet *entities.Tweet) error {
	return r.db.Where(tweet).Delete(tweet).Error
}

func (r *TweetRepository) UpdateTweet(tweet *entities.Tweet) error {
	query := &entities.Tweet{
		Id: tweet.Id,
	}
	return r.db.Where(query).Updates(tweet).Error
}

func (r *TweetRepository) UpdateTweetLikes(request *entities.Tweet) error {
	query := &entities.Tweet{
		Id: request.Id,
	}

	tweet := map[string]interface{}{
		"likes": request.Likes,
	}

	return r.db.Model(query).Where(query).Updates(tweet).Error
}
