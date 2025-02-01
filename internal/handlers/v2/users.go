package v2

import (
	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/internal/middleware"
	"github.com/achintya-7/dating-app/logger"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

func (rh *RouteHandler) CreateUserV2(ctx *gin.Context) (*dto.CreateUserResponse, *dto.ErrorResponse) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "Error while binding request: ", err)
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid request",
			HttpStatusCode: 400,
		}
	}

	password, err := utils.HashPassword(req.Password)
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
		Password:  password,
	}

	txArgs := db.CreateUserTx{
		UserReq: args,
	}

	user, err := rh.store.CreateUserTx(ctx, txArgs)
	if err != nil {
		logger.Error(ctx, "Error while creating user: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	user.Password = req.Password

	return user, nil
}

func (rh *RouteHandler) CreateRandomUserV2(ctx *gin.Context) (*dto.CreateUserResponse, *dto.ErrorResponse) {
	var req dto.CreateUserRequest

	err := faker.FakeData(&req)
	if err != nil {
		logger.Error(ctx, "Error while faking data: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	password, err := utils.HashPassword(req.Password)
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
		Password:  password,
	}

	txArgs := db.CreateUserTx{
		UserReq: args,
	}

	user, err := rh.store.CreateUserTx(ctx, txArgs)
	if err != nil {
		logger.Error(ctx, "Error while creating user: ", err)
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	user.Password = req.Password

	return user, nil
}

func (rh *RouteHandler) DiscoverV2(ctx *gin.Context) (*[]dto.DiscoverV2Response, *dto.ErrorResponse) {
	var req dto.DiscoverV2Request
	var resp []dto.DiscoverV2Response

	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err.Error() != "EOF" {
			logger.Error(ctx, "Error while binding request: ", err)
			return nil, &dto.ErrorResponse{
				Code:           400,
				Message:        "Invalid request",
				HttpStatusCode: 400,
			}
		}
	}

	if req.Age.LessThan != nil && *req.Age.LessThan <= 18 {
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Age should be greater than 18",
			HttpStatusCode: 400,
		}
	}

	if req.Age.GreaterThan != nil && *req.Age.GreaterThan <= 18 {
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Age should be greater than 18",
			HttpStatusCode: 400,
		}
	}

	if req.Age.LessThan != nil && req.Age.GreaterThan != nil && *req.Age.LessThan <= *req.Age.GreaterThan {
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Less than age should be greater than greater than age",
			HttpStatusCode: 400,
		}
	}

	authPayload := ctx.MustGet(middleware.AUTHORIZATION_PAYLOAD).(*token.Payload)

	discoverV2Args := db.DiscoverUsersV2Params{
		GreaterThanAge: utils.GetNullInt(req.Age.GreaterThan),
		LowerThanAge:   utils.GetNullInt(req.Age.LessThan),
		Gender:         utils.GetNullString(req.Gender),
		SwiperID:       authPayload.UserId,
		UserID:         authPayload.UserId,
	}

	users, err := rh.store.DiscoverUsersV2(ctx, discoverV2Args)
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

	currentUser, err := rh.store.GetUserByEmail(ctx, authPayload.Email)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	for _, user := range users {
		distance := utils.CalculateDistance(currentUser.Latitude, currentUser.Longitude, user.Latitude, user.Longitude, "K")

		respItem := dto.DiscoverV2Response{
			UserID:   user.UserID,
			Name:     user.Name,
			Gender:   user.Gender,
			Age:      user.Age,
			Distance: distance,
		}

		resp = append(resp, respItem)
	}

	return &resp, nil
}
