package v2

import (
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	distributor "github.com/achintya-7/dating-app/pkg/worker/distributor"
)

type RouteHandler struct {
	store       *db.Store
	tokenMake   *token.PasetoMaker
	distributor distributor.TaskDistributor
}

func NewRouteHandler(store *db.Store, tokenMaken *token.PasetoMaker, distributor distributor.TaskDistributor) *RouteHandler {
	return &RouteHandler{
		store:     store,
		tokenMake: tokenMaken,
		distributor: distributor,
	}
}
