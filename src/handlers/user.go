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

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetUserMeHandler(ctx *gin.Context) {
	response := new(helpers.Response)

	userToken, _ := ctx.Get("User")
	user := userToken.(*entities.User)

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = user
	response.SendJson(ctx)
}

func (h *UserHandler) GetUserHandler(ctx *gin.Context) {
	response := new(helpers.Response)

	id := ctx.Param("id")

	user := &entities.User{Id: id}
	status, err := h.userService.FindUser(user)
	if err != nil {
		response.Status = status
		response.Message = err.Error()
		response.SendJson(ctx)
		return
	}

	response.Message = "OK"
	response.Status = http.StatusOK
	response.Data = user
	response.SendJson(ctx)
}

func (h *UserHandler) PostUserHandler(ctx *gin.Context) {
	var request entities.User
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	status, err := h.userService.CreateUser(&request)
	if err != nil {
		response.Status = status
		response.Message = err.Error()
		response.SendJson(ctx)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Created"
	response.Data = request
	response.SendJson(ctx)
}

func (h *UserHandler) PatchUserHandler(ctx *gin.Context) {
	var request schemas.UpdateUserSchema
	response := new(helpers.Response)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	userToken, _ := ctx.Get("User")
	userId := userToken.(*entities.User).Id

	err = h.userService.UpdateUser(&request, userId)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Updated"
	response.Data = "Success Updated"
	ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	var request entities.User
	response := &helpers.Response{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			helpers.ResponseBadRequest(ctx, response, e)
			return
		}
	}

	err = h.userService.DeleteUser(&request)
	if err != nil {
		helpers.ResponseBadRequest(ctx, response, err)
		return
	}

	response.Data = "User deleted!"
	response.Message = "Deleted"
	response.Status = http.StatusOK
	ctx.JSON(response.Status, response)
}
