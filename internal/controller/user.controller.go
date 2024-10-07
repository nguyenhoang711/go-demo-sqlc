package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong.hhh.oing" + uc.userService.GetUserById(uid),
		"uid": uid,
		"users": []string{"cr7", "m10", "anomystick"},
	})
}