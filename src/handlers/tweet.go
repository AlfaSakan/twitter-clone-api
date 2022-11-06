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
	user := new(entities.User)

	userToken, ok := ctx.Get("User")
	if ok {
		user = userToken.(*entities.User)
	}

	err := h.tweetService.GetAllTweets(tweets, user.Id)
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

	userId := ""

	if userToken, ok := ctx.Get("User"); ok {
		userId = userToken.(*entities.User).Id
	}

	err = h.tweetService.FindListTweets(&tweet, tweets, userId)
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
	userId := ""

	id := ctx.Param("id")

	userToken, ok := ctx.Get("User")
	if ok {
		userId = userToken.(*entities.User).Id
	}

	tweet := &entities.Tweet{Id: id}
	response := new(helpers.Response)

	status, err := h.tweetService.FindTweet(tweet, userId)
	if err != nil {
		response.Status = status
		response.Message = err.Error()
		response.SendJson(ctx)
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

	user := new(entities.User)

	userToken, ok := ctx.Get("User")
	if ok {
		user = userToken.(*entities.User)
	}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	request.UserId = user.Id

	tweet, err := h.tweetService.CreateTweet(request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Created"
	response.Data = tweet
	response.SendJson(ctx)
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
	userToken, _ := ctx.Get("User")
	userId := userToken.(*entities.User).Id

	var request entities.TweetLike
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	request.UserId = userId

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
	var request schemas.TweetRequestByUserId
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	tweets, err := h.tweetService.FindLikeTweetsService(&request)
	if err != nil {
		helpers.ResponseNotFound(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = tweets
	ctx.JSON(http.StatusOK, response)
}
