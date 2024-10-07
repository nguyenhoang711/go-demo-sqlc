package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	Code int `json:"code"` // status code
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, responseData{
		Code: code,
		Message: ErrMessageDict()(code),
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusConflict, responseData{
		Code: code,
		Message: ErrMessageDict()(code),
	})
}