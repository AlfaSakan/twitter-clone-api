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

type SessionHandler struct {
	sessionService services.ISessionService
	userService    services.IUserService
}

func NewSessionHandler(sessionService services.ISessionService, userService services.IUserService) *SessionHandler {
	return &SessionHandler{sessionService, userService}
}

func (s *SessionHandler) PostSessionHandler(ctx *gin.Context) {
	request := schemas.SessionRequest{}
	response := &helpers.Response{}
	user := entities.User{}
	session := entities.Session{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	request.UserAgent = ctx.Request.UserAgent()

	err = s.sessionService.Login(&request, &user, &session)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	accessToken, refreshToken, err := s.sessionService.GenerateAccessRefresh(&user, &session)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = &entities.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionId:    session.Id,
	}

	ctx.JSON(response.Status, response)
}

func (s *SessionHandler) DeleteSessionHandler(ctx *gin.Context) {
	response := &helpers.Response{}

	sessionId := ctx.Param("sessionId")
	request := entities.Session{
		Id: sessionId,
	}

	err := s.sessionService.Logout(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Message = "Success Logout"
	response.Status = http.StatusOK

	ctx.JSON(response.Status, response)
}
