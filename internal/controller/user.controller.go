package controller

import (
	// "fmt"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/internal/service"
	"github.com/hoangnguyen/demo-sqlc/internal/vo"
	"github.com/hoangnguyen/demo-sqlc/pkg/response"
)

// "github.com/hoangnguyen/demo-sqlc/pkg/response"

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	uid := c.Query("uid")
// 	response.SuccessResponse(c, 20001, fmt.Sprintf("pong.hhh.ping.%s", uc.userService.GetUserById(uid)))
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"message": "pong.hhh.oing" + uc.userService.GetUserById(uid),
// 	// 	"uid": uid,
// 	// 	"users": []string{"cr7", "m10", "anomystick"},
// 	// })
// }

// INTERFACE_VERSION
type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController{
	return &UserController{
		userService: userService,
	}
}
func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
  		response.ErrorResponse(c, response.ErrCodeParamInvalid)
	}
	fmt.Printf("Email params: %s", params.Email)
	res := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, res, nil)
}