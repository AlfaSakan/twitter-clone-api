package routes

import (
	"fmt"

	"github.com/AlfaSakan/twitter-clone-api/src/handlers"
	"github.com/gin-gonic/gin"
)

const USER_ROUTE = "/user"

func User(r *gin.RouterGroup, h *handlers.UserHandler) {
	r.GET(fmt.Sprintf("%s/:id", USER_ROUTE), h.GetUserHandler)

	r.POST(USER_ROUTE, h.PostUserHandler)

	r.PATCH(USER_ROUTE, h.PatchUserHandler)

	r.DELETE(USER_ROUTE, h.DeleteUserHandler)
}
