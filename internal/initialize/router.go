package initialize

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default() //tao 1 instance cua gin (middleware, version, etc ...)
	//use the middleware
	// r.Use(AA(), BB(), CC)
	// v1 := r.Group("/v1/2024")
	// {
	// 	v1.GET("/ping", controller.NewPongController().Pong)          // /v1/2024/ping
	// 	v1.GET("/user/1", controller.NewUserController().GetUserById) //params
	// 	v1.GET("/hello", controller.NewPongController().DemoTimeoutContext)
	// 	// v1.DELETE("/ping", controller.NewPongController().Pong)
	// 	// v1.PATCH("/ping", controller.NewPongController().Pong)
	// 	// v1.HEAD("/ping", controller.NewPongController().Pong)
	// 	// v1.OPTIONS("/ping", controller.NewPongController().Pong)
	// }

	// v2 := r.Group("/v2/2024")
	// {
	// 	v2.GET("/ping", controller.NewPongController().Pong) // /v2/2024/ping
	// 	v2.PUT("/ping", controller.NewPongController().Pong)
	// 	v2.DELETE("/ping", controller.NewPongController().Pong)
	// 	v2.PATCH("/ping", controller.NewPongController().Pong)
	// 	v2.HEAD("/ping", controller.NewPongController().Pong)
	// }

	return r

}