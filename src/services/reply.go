package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
)

type IReplyService interface {
	CreateReply(tweetId string, replyId string) error
}

type ReplyService struct {
	replyRepository repositories.IReplyRepository
}

func NewReplyService(replyRepository repositories.IReplyRepository) *ReplyService {
	return &ReplyService{replyRepository}
}

func (rs *ReplyService) CreateReply(tweetId string, replyId string) error {
	tweetReply := &entities.TweetReply{
		TweetId: tweetId,
		ReplyId: replyId,
	}

	return rs.replyRepository.CreateReply(tweetReply)
}
