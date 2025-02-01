package v1

import (
	"net/http"

	"github.com/achintya-7/dating-app/constants"
	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/logger"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

func (rh *RouteHandler) Login(ctx *gin.Context) (*dto.LoginResponse, *dto.ErrorResponse) {
	var req dto.LoginRequest
	var resp dto.LoginResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, &dto.ErrorResponse{
			Code:           http.StatusBadRequest,
			Message:        "Invalid request",
			HttpStatusCode: http.StatusBadRequest,
		}
	}

	user, err := rh.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Error(ctx, "Error getting user by email", err)
		return nil, &dto.ErrorResponse{
			Code:           http.StatusNotFound,
			Message:        "User not found",
			HttpStatusCode: http.StatusNotFound,
		}
	}

	check := utils.CheckPassword(req.Password, user.Password)
	if check != nil {
		return nil, &dto.ErrorResponse{
			Code:           http.StatusUnauthorized,
			Message:        "Invalid password",
			HttpStatusCode: http.StatusUnauthorized,
		}
	}

	token, _, err := rh.tokenMake.CreateToken(user.Name, user.UserID, user.Email, constants.TOKEN_EXPIRY)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           http.StatusInternalServerError,
			Message:        "Internal server error",
			HttpStatusCode: http.StatusInternalServerError,
		}
	}

	resp.Token = token

	return &resp, nil
}
