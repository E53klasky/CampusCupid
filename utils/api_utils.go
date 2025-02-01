package utils

import (
	"log"
	"net/http"

	"github.com/achintya-7/dating-app/internal/dto"

	"github.com/gin-gonic/gin"
)

// HandlerFunction is the type of handler function callback.
type HandlerFunction[T any] func(*gin.Context) (*T, *dto.ErrorResponse)

// HandlerWrapper is a wrapper around all the handler methods.
// Only this method will handle the response creation (i.e., context.JSON).
// All the handler methods in this wrapper should return a pair of values:
// 1. Result computed (which can be of any type)
// 2. A pointer to ErrorResponse
func HandlerWrapper[T any](callback HandlerFunction[T]) func(context *gin.Context) {
	return func(context *gin.Context) {
		defer handlePanic(context)

		result, err := callback(context)
		if err != nil {
			sendErrorResponse(context, err)
			return
		}

		sendSuccessResponse(context, result)
	}
}

func handlePanic(context *gin.Context) {
	if r := recover(); r != nil {
		log.Println("Recovered from panic: ", r)
		sendErrorResponse(context, &dto.ErrorResponse{
			Code:           http.StatusInternalServerError,
			Message:        "Internal Server Error",
			HttpStatusCode: http.StatusInternalServerError,
		})
	}
}

func sendErrorResponse(context *gin.Context, err *dto.ErrorResponse) {
	context.AbortWithStatusJSON(err.HttpStatusCode, dto.ApiResponse[any]{
		Status: false,
		Result: nil,
		Error: &dto.ApiError{
			Code:           err.Code,
			Message:        err.Message,
			HttpStatusCode: err.HttpStatusCode,
		},
	})
}

func sendSuccessResponse[T any](context *gin.Context, result *T) {
	if result == nil {
		sendNotFoundResponse(context)
		return
	}
	context.JSON(http.StatusOK, dto.ApiResponse[T]{
		Status: true,
		Result: result,
		Error:  nil,
	})
}

func sendNotFoundResponse(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusNotFound, dto.ApiResponse[any]{
		Status: false,
		Result: nil,
		Error: &dto.ApiError{
			Code:           http.StatusNotFound,
			Message:        "Oops! No data found",
			HttpStatusCode: http.StatusNotFound,
		},
	})
}
