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

type TweetHandler struct {
	tweetService services.ITweetService
}

func NewTweetHandler(tweetService services.ITweetService) *TweetHandler {
	return &TweetHandler{tweetService}
}

func (h *TweetHandler) GetAllTweetsHandler(ctx *gin.Context) {
	tweets := &[]entities.Tweet{}
	response := new(helpers.Response)

	err := h.tweetService.GetAllTweets(tweets)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = tweets
	ctx.JSON(http.StatusOK, response)
}

func (h *TweetHandler) FindAllTweetsHandler(ctx *gin.Context) {
	tweet := schemas.TweetRequestByUserId{}
	tweets := &[]entities.Tweet{}
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&tweet)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
		}
	}

	err = h.tweetService.FindListTweets(&tweet, tweets)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = tweets
	ctx.JSON(http.StatusOK, response)
}

func (h *TweetHandler) GetTweetByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	tweet := &entities.Tweet{Id: id}
	response := new(helpers.Response)

	err := h.tweetService.FindTweet(tweet)
	if err != nil {
		helpers.ResponseNotFound(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = tweet
	ctx.JSON(http.StatusOK, response)
}

func (h *TweetHandler) PostTweetHandler(ctx *gin.Context) {
	var request schemas.TweetRequest
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	tweet, err := h.tweetService.CreateTweet(request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Created"
	response.Data = tweet
	ctx.JSON(response.Status, response)
}

func (h *TweetHandler) DeleteTweetHandler(ctx *gin.Context) {
	var request schemas.TweetRequestById
	response := &helpers.Response{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	err = h.tweetService.DeleteTweet(request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Data = "Tweet deleted!"
	response.Message = "Deleted"
	response.Status = http.StatusOK
	ctx.JSON(response.Status, response)
}

func (h *TweetHandler) LikeTweetHandler(ctx *gin.Context) {
	var request entities.TweetLike
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	err = h.tweetService.LikeTweetService(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = "Success!"
	ctx.JSON(response.Status, response)
}

func (h *TweetHandler) GetLikeTweetHandler(ctx *gin.Context) {
	var request entities.TweetLike
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	err = h.tweetService.FindLikeTweetService(&request)
	if err != nil {
		helpers.ResponseNotFound(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = request
	ctx.JSON(http.StatusOK, response)
}
