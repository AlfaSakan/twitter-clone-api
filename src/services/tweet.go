package services

import (
	"net/http"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/AlfaSakan/twitter-clone-api/src/schemas"
)

type ITweetService interface {
	FindListTweets(request *schemas.TweetRequestByUserId, tweets *[]entities.Tweet, userId string) error
	FindTweet(tweet *entities.Tweet, userId string) (statusCode int, errorMessage error)
	CreateTweet(tweetRequest schemas.TweetRequest) (*entities.Tweet, error)
	DeleteTweet(tweetRequest schemas.TweetRequestById) error
	LikeTweetService(request *entities.TweetLike) error
	FindLikeTweetService(request *entities.TweetLike) error
	GetAllTweets(tweets *[]entities.Tweet, userId string) error
	AddReplyCounts(tweetId string) error
}

type TweetService struct {
	tweetRepository repositories.ITweetRepository
	tweetLikeRepo   repositories.ITweetLikeRepository
	userRepo        repositories.IUserRepository
}

func NewTweetService(
	tweetRepository repositories.ITweetRepository,
	tweetLikeRepo repositories.ITweetLikeRepository,
	userRepo repositories.IUserRepository,
) *TweetService {
	return &TweetService{tweetRepository, tweetLikeRepo, userRepo}
}

func (s *TweetService) GetAllTweets(tweets *[]entities.Tweet, userId string) error {
	userMap := map[string]entities.User{}

	err := s.tweetRepository.GetAllTweets(tweets)
	if err != nil {
		return err
	}

	for i, tw := range *tweets {
		if userId != "" {
			tweetLike := entities.TweetLike{
				TweetId: tw.Id,
				UserId:  userId,
			}

			e := s.tweetLikeRepo.FindLike(&tweetLike)
			if e != nil {
				(*tweets)[i].IsLike = false
			}

			(*tweets)[i].IsLike = tweetLike.IsLike
		}

		user := entities.User{
			Id: tw.UserId,
		}
		s.userRepo.FindUser(&user)

		if _, ok := userMap[user.Id]; !ok {
			userMap[user.Id] = user
		}

		(*tweets)[i].User = userMap[tw.UserId]
	}

	return nil
}

func (s *TweetService) FindListTweets(request *schemas.TweetRequestByUserId, tweets *[]entities.Tweet, userId string) error {
	tweet := &entities.Tweet{
		UserId:  request.UserId,
		Content: request.Content,
	}

	userInserted := map[string]entities.User{}

	err := s.tweetRepository.FindListTweets(tweet, tweets)
	if err != nil {
		return err
	}

	for i, tw := range *tweets {
		tweetLike := entities.TweetLike{
			TweetId: tw.Id,
			UserId:  userId,
		}
		e := s.tweetLikeRepo.FindLike(&tweetLike)
		if e != nil {
			(*tweets)[i].IsLike = false
		}

		(*tweets)[i].IsLike = tweetLike.IsLike

		if u, ok := userInserted[tw.UserId]; ok {
			(*tweets)[i].User = u
			continue
		}

		user := entities.User{
			Id: tw.UserId,
		}

		s.userRepo.FindUser(&user)
		(*tweets)[i].User = user
		userInserted[user.Id] = user
	}

	return nil
}

func (s *TweetService) FindTweet(tweet *entities.Tweet, userId string) (int, error) {
	err := s.tweetRepository.FindTweet(tweet)
	if err != nil {
		return http.StatusNotFound, err
	}

	tweet.User.Id = tweet.UserId
	status, err := s.userRepo.FindUser(&tweet.User)
	if err != nil {
		return status, err
	}

	tweetLike := entities.TweetLike{
		TweetId: tweet.Id,
		UserId:  userId,
	}

	err = s.tweetLikeRepo.FindLike(&tweetLike)
	if err != nil {
		tweet.IsLike = false
		return http.StatusOK, nil
	}
	tweet.IsLike = tweetLike.IsLike
	return http.StatusOK, nil
}

func (s *TweetService) CreateTweet(tweetRequest schemas.TweetRequest) (*entities.Tweet, error) {
	tweet := &entities.Tweet{
		Content: tweetRequest.Content,
		UserId:  tweetRequest.UserId,
		TypeId:  tweetRequest.TypeId,
	}

	tweet.Id = helpers.GenerateId()

	return tweet, s.tweetRepository.CreateTweet(tweet)
}

func (s *TweetService) DeleteTweet(tweetRequest schemas.TweetRequestById) error {
	tweet := &entities.Tweet{
		Id: tweetRequest.Id,
	}
	return s.tweetRepository.DeleteTweet(tweet)
}

func (s *TweetService) LikeTweetService(request *entities.TweetLike) error {
	tweet := &entities.Tweet{
		Id:     request.TweetId,
		UserId: request.UserId,
	}
	e := s.tweetRepository.FindTweet(tweet)
	if e != nil {
		return e
	}

	err := s.tweetLikeRepo.FindLike(request)
	if err != nil {
		e = s.tweetLikeRepo.CreateLike(request)
		if e != nil {
			return e
		}
		tweet.Likes++
		e = s.tweetRepository.UpdateTweetLikes(tweet)
		if e != nil {
			return e
		}

		return nil
	}

	request.IsLike = !request.IsLike
	err = s.tweetLikeRepo.UpdateIsLike(request)
	if err != nil {
		return err
	}

	if request.IsLike {
		tweet.Likes++
		e = s.tweetRepository.UpdateTweetLikes(tweet)
		if e != nil {
			return e
		}

		return nil
	}

	tweet.Likes--
	e = s.tweetRepository.UpdateTweetLikes(tweet)
	if e != nil {
		return e
	}

	return nil
}

func (s *TweetService) FindLikeTweetService(request *entities.TweetLike) error {
	return s.tweetLikeRepo.FindLike(request)
}

func (s *TweetService) AddReplyCounts(tweetId string) error {
	tweet := entities.Tweet{
		Id: tweetId,
	}

	err := s.tweetRepository.FindTweet(&tweet)
	if err != nil {
		return err
	}

	tweet.ReplyCounts++

	return s.tweetRepository.IncrementReplyCounts(&tweet)
}
