package controllers

import (
	v1 "github.com/achintya-7/dating-app/internal/handlers/v1"
	"github.com/achintya-7/dating-app/internal/middleware"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	"github.com/achintya-7/dating-app/utils"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handlers   *v1.RouteHandler
	tokenMaker *token.PasetoMaker
}

// NewV1Router creates a new router instance
func NewV1Router(store *db.Store, tokenMaker *token.PasetoMaker) *Router {
	return &Router{
		handlers:   v1.NewRouteHandler(store, tokenMaker),
		tokenMaker: tokenMaker,
	}
}

// SetupRoutes sets up the routes for the V1 router
func (r *Router) SetupRoutes(router *gin.RouterGroup) {
	v1Route := router.Group("/v1")

	v1Route.POST("/login", utils.HandlerWrapper(r.handlers.Login))

	// Setup user routes
	usersRoute := v1Route.Group("/users")
	usersRoute.POST("/create", utils.HandlerWrapper(r.handlers.CreateUser))

	// Apply auth middleware
	authRoute := v1Route.Group("/")
	authRoute.Use(middleware.AuthMiddleware(*r.tokenMaker))
	
	// User Discovery route V1
	userAuthRoute := authRoute.Group("/users")
	userAuthRoute.GET("/discover", utils.HandlerWrapper(r.handlers.DiscoverV1))

	// Setup swipe routes
	swipeAuthRoute := authRoute.Group("/swipe")
	swipeAuthRoute.POST("/", utils.HandlerWrapper(r.handlers.SwipeUser))
}
