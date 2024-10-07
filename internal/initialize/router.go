package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/internal/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default() //tao 1 instance cua gin (middleware, version, etc ...)
	//use the middleware
	// r.Use(AA(), BB(), CC)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)          // /v1/2024/ping
		v1.GET("/user/1", Pong) //params
	}

	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", controller.NewUserController().GetUserById) // /v2/2024/ping
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