package v1

import (
	"log"

	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/internal/middleware"
	"github.com/achintya-7/dating-app/logger"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (rh *RouteHandler) CreateUser(ctx *gin.Context) (*dto.CreateUserResponse, *dto.ErrorResponse) {
	var req dto.CreateUserRequest
	var resp dto.CreateUserResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "Error while binding request: ", err)
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid request",
			HttpStatusCode: 400,
		}
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Error(ctx, "Error while hashing password: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	user_id := uuid.New().String()

	args := db.CreateUserParams{
		UserID:    user_id,
		Email:     req.Email,
		Name:      req.Name,
		Gender:    req.Gender,
		Age:       int32(req.Age),
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Password:  hashedPassword,
	}

	_, err = rh.store.CreateUser(ctx, args)
	if err != nil {
		log.Println("Error while creating user: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	resp.ID = user_id
	resp.Email = req.Email
	resp.Name = req.Name
	resp.Age = req.Age
	resp.Password = req.Password
	resp.Gender = req.Gender

	return &resp, nil
}

func (rh *RouteHandler) DiscoverV1(ctx *gin.Context) (*[]db.DiscoverUsersV1Row, *dto.ErrorResponse) {
	authPayload := ctx.MustGet(middleware.AUTHORIZATION_PAYLOAD).(*token.Payload)

	discoverUserArgs := db.DiscoverUsersV1Params{
		UserID:   authPayload.UserId,
		SwiperID: authPayload.UserId,
	}

	users, err := rh.store.DiscoverUsersV1(ctx, discoverUserArgs)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	if len(users) == 0 {
		return nil, &dto.ErrorResponse{
			Code:           404,
			Message:        "No users found",
			HttpStatusCode: 404,
		}
	}

	return &users, nil
}
