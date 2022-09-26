package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseBadRequest(ctx *gin.Context, response *Response, err error) {
	response.Message = err.Error()
	response.Status = http.StatusBadRequest
	response.Data = ""
	ctx.JSON(response.Status, response)
}

func ResponseNotFound(ctx *gin.Context, response *Response, err error) {
	response.Message = err.Error()
	response.Status = http.StatusNotFound
	response.Data = ""
	ctx.JSON(response.Status, response)
}
