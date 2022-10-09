package handlers

import (
	"net/http"
	"time"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/schemas"
	"github.com/AlfaSakan/twitter-clone-api/src/services"
	"github.com/AlfaSakan/twitter-clone-api/src/utils"
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
	request := &schemas.SessionRequest{}

	response := &helpers.Response{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	user := entities.User{
		Username: request.Username,
		Password: request.Password,
	}

	err = s.userService.FindUser(&user)
	if err != nil {
		helpers.ResponseNotFound(ctx, response, err)
		return
	}

	userAgent := ctx.Request.UserAgent()
	session := entities.Session{
		UserAgent: userAgent,
		UserId:    user.Id,
		Valid:     true,
	}

	err = s.sessionService.Login(&session)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	accessClaims := &utils.CustomClaim{
		User: &user,
	}

	expireAccessToken := time.Now().Add(time.Hour * 12).UnixMilli()
	accessToken, err := utils.GenerateToken(accessClaims, expireAccessToken)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	refreshClaim := &utils.CustomClaim{
		User:      &entities.User{},
		SessionId: session.Id,
	}

	expireRefreshToken := time.Now().Add(time.Hour * 24 * 30 * 12).UnixMilli()
	refreshToken, err := utils.GenerateToken(refreshClaim, expireRefreshToken)
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
