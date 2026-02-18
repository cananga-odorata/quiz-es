package http

import (
	"encoding/json"
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/modules/quiz/application"
	"github.com/cananga-odorata/golang-template/internal/shared/dto"
	"github.com/go-chi/chi/v5"
)

// QuizHandler handles HTTP requests for quiz operations
type QuizHandler struct {
	service application.QuizService
}

// NewQuizHandler creates a new QuizHandler
func NewQuizHandler(service application.QuizService) *QuizHandler {
	return &QuizHandler{service: service}
}

// List handles GET /quizzes
func (h *QuizHandler) List(w http.ResponseWriter, r *http.Request) {
	quizzes, err := h.service.GetAll(r.Context())
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.OK(w, quizzes)
}

// Create handles POST /quizzes
func (h *QuizHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req application.CreateQuizRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.Error(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request payload")
		return
	}

	quiz, err := h.service.Create(r.Context(), req)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.Created(w, quiz)
}

// Delete handles DELETE /quizzes/{id}
func (h *QuizHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		dto.Error(w, http.StatusBadRequest, "VALIDATION_ERROR", "Quiz ID is required")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.NoContent(w)
}
