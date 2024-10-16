package initialize

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/hoangnguyen/demo-sqlc/internal/controller"
	// "github.com/hoangnguyen/demo-sqlc/internal/middlewares"
	// "github.com/hoangnguyen/demo-sqlc/internal/repo"
	// "github.com/hoangnguyen/demo-sqlc/internal/service"
	"github.com/hoangnguyen/demo-sqlc/internal/controller/account"
	"github.com/hoangnguyen/demo-sqlc/internal/wire"
)

func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> AA")
		ctx.Next()
		fmt.Println("After --> AA")
	}
} 

func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> BB")
		ctx.Next()
		fmt.Println("After --> BB")
	}
} 

func CC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> CC")
		ctx.Next()
		fmt.Println("After --> CC")
	}
} 
func InitRouter() *gin.Engine {
	r := gin.Default() //tao 1 instance cua gin (middleware, version, etc ...)
	//use the middleware
	r.Use(AA(), BB(), CC())
	// r.Use(middlewares.AuthenMiddleware(), BB(), CC())
	//non-dependency
	// userRepoNonDependency := repo.NewUSerRepo()
	// userServiceNonDependency := service.NewUserService(userRepoNonDependency)
	// userHandlerNonDependency := controller.NewUserController(userServiceNonDependency)
	userController, _ := wire.InitUserController()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)          // /v1/2024/ping
		v1.GET("/user/1", Pong) //params
	}

	v2 := r.Group("/v2/2024")
	{
		// v2.GET("/ping", controller.NewUserController().GetUserById) // /v2/2024/ping
		v2.POST("/user/register", userController.Register) // /v2/2024/ping
		v2.POST("/user/login", account.Login.Login)
	}

	return r

}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "anonystick")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong.hhh.oing" + name,
		"uid": uid,
		"users": []string{"cr7", "m10", "anomystick"},
	})
}