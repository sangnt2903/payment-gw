package http

import (
	"payment-gw/internal/adapters/primary/http/middleware"
	"payment-gw/internal/core/ports/input"
	"payment-gw/pkg/conf"
	"time"

	_ "payment-gw/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router         *gin.Engine
	paymentHandler *PaymentHandler
	rdb            *redis.Client
}

func NewServer(paymentUseCase input.PaymentUseCase) *Server {
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:          12 * time.Hour,
	}))

	rdb := redis.NewClient(&redis.Options{
		Addr: conf.GetString("redis", "addr"),
	})

	server := &Server{
		router:         router,
		paymentHandler: NewPaymentHandler(paymentUseCase),
		rdb:           rdb,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Apply rate limit middleware to all routes
	s.router.Use(middleware.RateLimit(s.rdb, 100, time.Second))

	// Swagger docs
	s.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	api := s.router.Group("/api/v1")
	{
		payments := api.Group("/payments")
		{
			payments.POST("", s.paymentHandler.CreatePayment)
			payments.GET("/:id", s.paymentHandler.GetPayment)
		}
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
