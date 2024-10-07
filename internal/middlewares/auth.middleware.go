package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrCodeNotAuthen)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}