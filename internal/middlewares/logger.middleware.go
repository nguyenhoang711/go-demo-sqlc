package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/global"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		global.Logger.Info(fmt.Sprintf("Started %s %s ",ctx.Request.Method, ctx.Request.URL.Path))
		ctx.Next()
		global.Logger.Info(fmt.Sprintf("Completed in %v", time.Since(start)))
	}
}