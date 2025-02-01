package app

import (
	"database/sql"
	"net/http"

	"github.com/achintya-7/dating-app/config"
	"github.com/achintya-7/dating-app/internal/controllers"
	"github.com/achintya-7/dating-app/internal/middleware"
	"github.com/achintya-7/dating-app/logger"
	"github.com/achintya-7/dating-app/pkg/mail"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/token"
	distributor "github.com/achintya-7/dating-app/pkg/worker/distributor"
	processor "github.com/achintya-7/dating-app/pkg/worker/processor"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type Server struct {
	router      *gin.Engine
	store       *db.Store
	tokenMaker  *token.PasetoMaker
	distributor distributor.TaskDistributor
	processor   *processor.RedisTaskProcessor
}

// NewServer creates a new server instance
func NewServer() *Server {
	server := &Server{}

	server.setupDatbases()
	server.setupClient()
	server.setupWorker()
	server.setupRouter()

	return server
}

// setupDatabases sets up the database connection and store
func (s *Server) setupDatbases() {
	conn, err := sql.Open("mysql", config.Values.MySqlUrl)
	if err != nil {
		logger.Fatal(nil, "cannot connect to database")
	}

	store := db.NewStore(conn)

	s.store = store
}

// setupClient sets up the token maker
func (s *Server) setupClient() {
	tokenMaker, err := token.NewPasetoMaker(config.Values.TokenSymmetricKey)
	if err != nil {
		logger.Fatal(nil, "cannot create token maker", err)
	}

	s.tokenMaker = tokenMaker
}

// setupWorker sets up the task distributor and processor
func (s *Server) setupWorker() {
	redisOpt := asynq.RedisClientOpt{
		Addr: config.Values.RedisUrl,
	}

	taskDistributor := distributor.NewRedisTaskDistributor(redisOpt)

	// Create a mail sender
	mailer := mail.NewGmailSender(config.Values.EmailName, config.Values.EmailAddress, config.Values.EmailPassowrd)

	// Create a processor instance
	redisTaskProcessor := processor.NewRedisTaskProcessor(redisOpt, s.store, mailer)

	s.distributor = taskDistributor
	s.processor = redisTaskProcessor
}

// setupRouter sets up the gin router
func (s *Server) setupRouter() {
	router := gin.Default()

	// Create a base router group
	baseRouter := router.Group("/dating-app")

	// Register health route
	s.registerHealthRoute(baseRouter)

	// Add a rate limiter middleware
	limiter := tollbooth.NewLimiter(float64(3), nil) // 3 requests per second
	limiter.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})
	baseRouter.Use(tollbooth_gin.LimitHandler(limiter))

	// Setup middleware
	baseRouter.Use(middleware.SetCorrelationIdMiddleware())

	// Setup v1 routes
	v1Router := controllers.NewV1Router(s.store, s.tokenMaker)
	v1Router.SetupRoutes(baseRouter)

	// Setup v2 routes
	v2Router := controllers.NewV2Router(s.store, s.tokenMaker, s.distributor)
	v2Router.SetupRoutes(baseRouter)

	// Setup CORS middleware
	router.Use(middleware.CorsMiddleware())

	s.router = router
}

// Start starts the server
func (s *Server) Start() error {
	go func() {
		logger.Info(nil, "starting task processor")
		err := s.processor.Start()
		if err != nil {
			logger.Fatal(nil, "cannot start task processor")
		}
	}()

	port := config.Values.HttpPort
	port = ":" + port

	logger.Info(nil, "starting server at port "+port)
	return s.router.Run(port)
}

// registerHealthRoute registers the health route
func (s *Server) registerHealthRoute(router *gin.RouterGroup) {
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})
}
