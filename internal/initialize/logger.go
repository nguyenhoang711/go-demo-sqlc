package initialize

import (
	"github.com/hoangnguyen/demo-sqlc/global"
	"github.com/hoangnguyen/demo-sqlc/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}