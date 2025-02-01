package middleware

import (
	"github.com/achintya-7/dating-app/constants"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Gin Middleware to set the correlation id in the gin context and response headers
func SetCorrelationIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// if the client has sent the correlation id, use that
		correlationId := ctx.GetHeader(constants.CORRELATION_HEADER)
		if correlationId == "" { // if not, create new
			correlationId = uuid.New().String()
		}

		ctx.Set(constants.CORRELATION_ID, correlationId)
		ctx.Header(constants.CORRELATION_HEADER, correlationId)

		ctx.Next()
	}
}
