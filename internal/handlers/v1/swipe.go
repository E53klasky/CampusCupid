package v1

import (
	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/achintya-7/dating-app/internal/middleware"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/gin-gonic/gin"
)

func (rh *RouteHandler) SwipeUser(ctx *gin.Context) (*gin.H, *dto.ErrorResponse) {
	var req dto.SwipeRequest
	var resp gin.H

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, &dto.ErrorResponse{
			Code:           400,
			Message:        "Invalid request",
			HttpStatusCode: 400,
		}
	}

	authPayload := ctx.MustGet(middleware.AUTHORIZATION_PAYLOAD).(*token.Payload)

	if req.Preference == "NO" {
		resp = gin.H{
			"matched": false,
		}

		return &resp, nil
	}

	txArgs := db.AddSwipeTx{
		User1Id: authPayload.UserId,
		User2Id: req.SwipedUserId,
	}

	match, err := rh.store.AddMatchTx(ctx, txArgs)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Code:           500,
			Message:        "Internal server error",
			HttpStatusCode: 500,
		}
	}

	if match.MatchID != "" {
		resp = gin.H{
			"matched": true,
			"matchId": match.MatchID,
		}
	} else {
		resp = gin.H{
			"matched": false,
		}
	}

	return &resp, nil
}
