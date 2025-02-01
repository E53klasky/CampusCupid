package v1

import (
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
)

type RouteHandler struct {
	store     *db.Store
	tokenMake *token.PasetoMaker
}

func NewRouteHandler(store *db.Store, tokenMaken *token.PasetoMaker) *RouteHandler {
	return &RouteHandler{
		store:     store,
		tokenMake: tokenMaken,
	}
}
