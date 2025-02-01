package v2

import (
	"strings"

	"github.com/achintya-7/dating-app/constants"
	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/internal/middleware"
	"github.com/achintya-7/dating-app/logger"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/pkg/worker"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (rh *RouteHandler) SwipeUserV2(ctx *gin.Context) (*gin.H, *dto.ErrorResponse) {
	var req dto.SwipeRequest
	var resp gin.H
	var errResp *dto.ErrorResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid request",
			HttpStatusCode: 400,
		}
	}

	if req.Preference != string(db.SwipesSwipeTypeYES) && req.Preference != string(db.SwipesSwipeTypeNO) {
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid preference",
			HttpStatusCode: 400,
		}
	}

	authPayload := ctx.MustGet(middleware.AUTHORIZATION_PAYLOAD).(*token.Payload)
	correlationId := ctx.GetString(constants.CORRELATION_ID)

	// If the user swipes NO, then no need to check for a match
	if req.Preference == string(db.SwipesSwipeTypeNO) {
		resp, errResp = rh.handleNoSwipe(ctx, req, authPayload.UserId, correlationId)
		if errResp != nil {
			return nil, errResp
		}
	} else {
		resp, errResp = rh.handleYesSwipee(ctx, req, authPayload.UserId, correlationId)
		if errResp != nil {
			return nil, errResp
		}
	}

	return &resp, nil
}

func (rh *RouteHandler) handleNoSwipe(ctx *gin.Context, req dto.SwipeRequest, userId, correlationId string) (gin.H, *dto.ErrorResponse) {
	resp := gin.H{
		"matched": false,
	}

	// Create a swipe with YES for user1
	createSwipeArgs := db.CreateSwipeParams{
		SwipeID:   uuid.New().String(),
		SwiperID:  userId,
		SwipeeID:  req.SwipedUserId,
		SwipeType: db.SwipesSwipeTypeNO,
	}

	_, err := rh.store.CreateSwipe(ctx, createSwipeArgs)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, &dto.ErrorResponse{
				Code:           409,
				Message:        "Duplicate swipe",
				HttpStatusCode: 409,
			}
		}

		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	// send the swiped user's details for rank calculation
	rankPayload := worker.CalculateUserAttractivenessTaskPayload{
		Userid:        req.SwipedUserId,
		Response:      req.Preference,
		CorrelationId: correlationId,
	}

	go rh.distributor.CalculateUserAttractivenessTask(ctx, &rankPayload)

	return resp, nil
}

func (rh *RouteHandler) handleYesSwipee(ctx *gin.Context, req dto.SwipeRequest, userId, correlationId string) (gin.H, *dto.ErrorResponse) {
	resp := gin.H{
		"matched": false,
	}

	// If the user swipes YES, then check for a match
	txArgs := db.AddSwipeTx{
		User1Id: userId,
		User2Id: req.SwipedUserId,
	}

	match, err := rh.store.AddMatchTx(ctx, txArgs)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, &dto.ErrorResponse{
				Code:           409,
				Message:        "Duplicate swipe",
				HttpStatusCode: 409,
			}
		}

		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	if match.MatchID != "" {
		logger.Info(ctx, "Match found")

		resp = gin.H{
			"matched": true,
			"matchId": match.MatchID,
		}

		// send an email to both users
		emailPayload := worker.SendMatchEmailTaskPayload{
			UserId:        userId,
			MatchedUserId: req.SwipedUserId,
			CorrelationId: correlationId,
		}

		go rh.distributor.SendMatchedEmailTask(ctx, &emailPayload)

	} else {
		logger.Info(ctx, "No match found")

		resp = gin.H{
			"matched": false,
		}
	}

	logger.Info(ctx, "Sending user attractiveness calculation task")

	// send the swiped user's details for rank calculation
	rankPayload := worker.CalculateUserAttractivenessTaskPayload{
		Userid:        req.SwipedUserId,
		Response:      req.Preference,
		CorrelationId: correlationId,
	}

	go rh.distributor.CalculateUserAttractivenessTask(ctx, &rankPayload)

	logger.Info(ctx, "User attractiveness calculation task sent")

	return resp, nil
}
