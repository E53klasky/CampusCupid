package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION_HEADER_KEY  = "authorization"
	AUTHORIZATION_TYPE_KEY = "bearer"
	AUTHORIZATION_PAYLOAD = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(tokenMaker token.PasetoMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION_HEADER_KEY)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:           http.StatusUnauthorized,
				Message:        err.Error(),
				HttpStatusCode: http.StatusUnauthorized,
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:           http.StatusUnauthorized,
				Message:        err.Error(),
				HttpStatusCode: http.StatusUnauthorized,
			})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AUTHORIZATION_TYPE_KEY {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:           http.StatusUnauthorized,
				Message:        err.Error(),
				HttpStatusCode: http.StatusUnauthorized,
			})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:           http.StatusUnauthorized,
				Message:        err.Error(),
				HttpStatusCode: http.StatusUnauthorized,
			})
			return
		}

		ctx.Set(AUTHORIZATION_PAYLOAD, payload)
		ctx.Next()
	}
}
