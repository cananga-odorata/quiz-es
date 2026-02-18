package http

import (
	"github.com/cananga-odorata/golang-template/internal/modules/quiz/application"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all quiz module routes
func RegisterRoutes(r chi.Router, service application.QuizService) {
	handler := NewQuizHandler(service)

	r.Route("/quizzes", func(r chi.Router) {
		r.Get("/", handler.List)
		r.Post("/", handler.Create)
		r.Delete("/{id}", handler.Delete)
	})
}
