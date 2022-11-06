package handlers

import (
	"net/http"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/schemas"
	"github.com/AlfaSakan/twitter-clone-api/src/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RetweetHandler struct {
	retweetService services.IRetweetService
	tweetService   services.ITweetService
}

func NewRetweetHandler(
	retweetService services.IRetweetService,
	tweetService services.ITweetService,
) *RetweetHandler {
	return &RetweetHandler{
		retweetService,
		tweetService,
	}
}

func (h *RetweetHandler) PostRetweetHandler(ctx *gin.Context) {
	var request schemas.PostRetweetSchema
	response := new(helpers.Response)

	user := new(entities.User)

	userToken, ok := ctx.Get("User")
	if ok {
		user = userToken.(*entities.User)
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	retweet, err := h.tweetService.CreateTweet(schemas.TweetRequest{
		Content: request.Content,
		UserId:  user.Id,
		TypeId:  entities.TypeRetweet,
	})
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
	}

	err = h.retweetService.CreateRetweet(request.TweetId, retweet.Id)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	err = h.tweetService.AddRetweetCounts(request.TweetId)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Created Retweet"
	response.Data = retweet
	response.SendJson(ctx)
}

func (h *RetweetHandler) GetRetweetsByIdHandler(ctx *gin.Context) {
	tweetId := ctx.Param("tweetId")
	retweets := []entities.Retweet{}
	tweets := []entities.Tweet{}
	response := new(helpers.Response)
	userId := ""

	userToken, ok := ctx.Get("User")
	if ok {
		userId = userToken.(*entities.User).Id
	}

	if err := h.retweetService.FindRetweets(tweetId, &retweets); err != nil {
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
		response.SendJson(ctx)
	}

	for _, retweet := range retweets {
		tweet := entities.Tweet{
			Id: retweet.RetweetId,
		}

		_, err := h.tweetService.FindTweet(&tweet, userId)
		if err == nil {
			tweets = append(tweets, tweet)
		}
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = tweets
	response.SendJson(ctx)
}
