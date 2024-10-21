package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangnguyen/demo-sqlc/global"
	"github.com/hoangnguyen/demo-sqlc/pkg/response"
	"go.uber.org/zap"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // Process the request

        // Check if there are any errors after request processing
        if len(c.Errors) > 0 {
            // Extract the last error recorded
            err := c.Errors.Last()
			global.Logger.Error(err.Error(), zap.Error(err))
            // Assuming the error is wrapped in a structured response code
            if customError, ok := err.Err.(response.CustomError); ok {
                // Use your existing ErrorResponse function
                response.ErrorResponse(c, customError.Code)
            } else {
                // Default to internal server error
                response.ErrorResponse(c, response.ErrCodeConflict)
            }
        }
    }
}