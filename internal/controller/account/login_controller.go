package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/internal/service"
	"github.com/hoangnguyen/demo-sqlc/pkg/response"
)
var Login = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	//implement logic for login
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}