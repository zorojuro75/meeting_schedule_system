package server

import (
	"fmt"
	"log"
	"meeting_scheduler/config"
	"meeting_scheduler/internal/handlers"
	"meeting_scheduler/internal/middleware"
	"meeting_scheduler/internal/models"
	"meeting_scheduler/internal/repository"
	"meeting_scheduler/internal/scheduler"
	"meeting_scheduler/internal/services"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Server is the application server with dependencies.
type Server struct {
	engine *gin.Engine
	db     *gorm.DB
	cfg    *config.Config
}

// NewServer creates and configures the server.
func NewServer(cfg *config.Config) (*Server, error) {
	dsn := cfg.DatabaseURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// AutoMigrate could be used for initial development.
	// AutoMigrate models (safe for development)
	if err := db.AutoMigrate(&models.User{}, &models.Meeting{}, &models.MeetingParticipant{}, &models.AuditLog{}); err != nil {
		return nil, fmt.Errorf("automigrate: %w", err)
	}

	r := gin.Default()

	// CORS for NextJS frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Simple health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Wire repositories, services and handlers
	repos := repository.NewRepos(db)
	authSvc := services.NewAuthService(repos, cfg.JwtSecret)
	meetingSvc := services.NewMeetingService(repos)

	authHandler := handlers.NewAuthHandler(authSvc)
	meetingHandler := handlers.NewMeetingHandler(meetingSvc)

	// Public routes
	r.POST("/auth/login", authHandler.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTAuth(cfg.JwtSecret))
	api.POST("/meetings", meetingHandler.Create)

	srv := &Server{engine: r, db: db, cfg: cfg}

	// Start scheduler in background
	go func() {
		if err := scheduler.Start(db, cfg); err != nil {
			log.Printf("scheduler error: %v", err)
		}
	}()

	return srv, nil
}

// Run starts the HTTP server on the provided address (e.g. ":8080").
func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
