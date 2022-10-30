package repositories

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"gorm.io/gorm"
)

type IReplyRepository interface {
	CreateReply(reply *entities.TweetReply) error
	FindReplyId(reply *entities.TweetReply) error
	FindRepliesId(tweetId string, reply *[]entities.TweetReply) error
}

type ReplyRepository struct {
	db *gorm.DB
}

func NewReplyRepository(db *gorm.DB) *ReplyRepository {
	return &ReplyRepository{db}
}

func (rr *ReplyRepository) CreateReply(reply *entities.TweetReply) error {
	return rr.db.Create(reply).Error
}

func (rr *ReplyRepository) FindReplyId(reply *entities.TweetReply) error {
	return rr.db.First(reply).Error
}

func (rr *ReplyRepository) FindRepliesId(tweetId string, reply *[]entities.TweetReply) error {
	return rr.db.Where(&entities.TweetReply{TweetId: tweetId}).Find(reply).Error
}
