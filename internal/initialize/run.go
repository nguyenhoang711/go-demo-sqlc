package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/global"
	"go.uber.org/zap"
)

func Run() *gin.Engine{
	// env := os.Getenv("ENV")
	// if env == "" {
    //     log.Fatal("ENV variable is not set")
    // }
	LoadConfig()
	// LoadConfig(env)
	InitLogger()
	InitMysqlc()
	InitRedis()
	InitServiceInterface()
	global.Logger.Info("Config Log Ok!!!", zap.String("config", "success"))

	r := InitRouter()

	return r
	// r.Run(":8002")
}