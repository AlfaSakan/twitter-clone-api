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

type ReplyHandler struct {
	replyService services.IReplyService
	tweetService services.ITweetService
}

func NewReplyHandler(
	replyService services.IReplyService,
	tweetService services.ITweetService,
) *ReplyHandler {
	return &ReplyHandler{replyService, tweetService}
}

func (h *ReplyHandler) PostReplyHandler(ctx *gin.Context) {
	var request schemas.PostReplySchema
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	reply, err := h.tweetService.CreateTweet(schemas.TweetRequest{
		Content: request.Content,
		UserId:  request.UserId,
		TypeId:  entities.TypeReply,
	})
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	err = h.replyService.CreateReply(request.TweetId, reply.Id)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	err = h.tweetService.AddReplyCounts(request.TweetId)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Created Reply"
	response.Data = reply
	ctx.JSON(response.Status, response)
}
