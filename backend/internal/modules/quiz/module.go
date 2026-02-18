package quiz

import (
	"github.com/cananga-odorata/golang-template/internal/modules/quiz/application"
	"github.com/cananga-odorata/golang-template/internal/modules/quiz/infrastructure"
	httpinterface "github.com/cananga-odorata/golang-template/internal/modules/quiz/interfaces/http"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// Module represents the quiz module with all its dependencies
type Module struct {
	Service application.QuizService
}

// NewModule initializes the quiz module with all dependencies
func NewModule(db *sqlx.DB) *Module {
	repo := infrastructure.NewPostgresQuizRepository(db)
	service := application.NewQuizService(repo)

	return &Module{
		Service: service,
	}
}

// RegisterRoutes registers the module's HTTP routes
func (m *Module) RegisterRoutes(r chi.Router) {
	httpinterface.RegisterRoutes(r, m.Service)
}
