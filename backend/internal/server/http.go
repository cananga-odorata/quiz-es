package server

import (
	"log/slog"
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/config"
	"github.com/cananga-odorata/golang-template/internal/modules/quiz"
	"github.com/cananga-odorata/golang-template/internal/shared/middleware"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

// Server holds the HTTP server dependencies
type Server struct {
	Router *chi.Mux
	Config *config.Config
	DB     *sqlx.DB
}

// New creates a new server with all modules wired
func New(cfg *config.Config, db *sqlx.DB) *Server {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.RequestID)
	// Rate limiting
	r.Use(middleware.RateLimitMiddleware(cfg.RateLimit, cfg.RateLimitBurst))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-ID"},
		ExposedHeaders:   []string{"Link", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check
	r.Get("/health", healthHandler)

	// Initialize modules
	quizModule := quiz.NewModule(db)

	// API v1 routes
	r.Route("/api/v1", func(api chi.Router) {
		// Quiz routes (public - no auth required for this assignment)
		quizModule.RegisterRoutes(api)
	})

	slog.Info("Server initialized",
		"modules", []string{"quiz"},
		"environment", cfg.Environment,
	)

	return &Server{
		Router: r,
		Config: cfg,
		DB:     db,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
