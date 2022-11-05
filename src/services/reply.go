package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
)

type IReplyService interface {
	CreateReply(tweetId string, replyId string) error
	FindReplies(tweetId string, replies *[]entities.TweetReply) error
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

func (rs *ReplyService) FindReplies(tweetId string, replies *[]entities.TweetReply) error {
	return rs.replyRepository.FindRepliesId(tweetId, replies)
}
