package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cananga-odorata/golang-template/internal/modules/quiz/application"
	"github.com/cananga-odorata/golang-template/internal/modules/quiz/domain"
	sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"
	"github.com/go-chi/chi/v5"
)

// mockQuizService is a mock implementation of application.QuizService
type mockQuizService struct {
	quizzes   []application.QuizResponse
	createErr error
	deleteErr error
	created   *application.QuizResponse
}

func (m *mockQuizService) GetAll(_ context.Context) ([]application.QuizResponse, error) {
	return m.quizzes, nil
}

func (m *mockQuizService) Create(_ context.Context, _ application.CreateQuizRequest) (*application.QuizResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	return m.created, nil
}

func (m *mockQuizService) Delete(_ context.Context, _ string) error {
	return m.deleteErr
}

// ============ Test Cases ============

func TestListHandler_Empty(t *testing.T) {
	svc := &mockQuizService{quizzes: []application.QuizResponse{}}
	handler := NewQuizHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/quizzes", nil)
	rec := httptest.NewRecorder()

	handler.List(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	json.NewDecoder(rec.Body).Decode(&resp)
	if resp["success"] != true {
		t.Errorf("expected success true, got %v", resp["success"])
	}
}

func TestListHandler_WithData(t *testing.T) {
	svc := &mockQuizService{
		quizzes: []application.QuizResponse{
			{ID: "1", Question: "Q1", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 1},
			{ID: "2", Question: "Q2", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 2},
		},
	}
	handler := NewQuizHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/quizzes", nil)
	rec := httptest.NewRecorder()

	handler.List(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

func TestCreateHandler_Success(t *testing.T) {
	svc := &mockQuizService{
		created: &application.QuizResponse{
			ID:           "test-id",
			Question:     "Test Q",
			Choice1:      "A",
			Choice2:      "B",
			Choice3:      "C",
			Choice4:      "D",
			DisplayOrder: 1,
		},
	}
	handler := NewQuizHandler(svc)

	body, _ := json.Marshal(application.CreateQuizRequest{
		Question: "Test Q",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	})

	req := httptest.NewRequest(http.MethodPost, "/quizzes", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rec.Code)
	}
}

func TestCreateHandler_InvalidJSON(t *testing.T) {
	svc := &mockQuizService{}
	handler := NewQuizHandler(svc)

	req := httptest.NewRequest(http.MethodPost, "/quizzes", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestCreateHandler_ValidationError(t *testing.T) {
	svc := &mockQuizService{
		createErr: domain.ErrInvalidQuiz,
	}
	handler := NewQuizHandler(svc)

	body, _ := json.Marshal(application.CreateQuizRequest{
		Question: "",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	})

	req := httptest.NewRequest(http.MethodPost, "/quizzes", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestDeleteHandler_Success(t *testing.T) {
	svc := &mockQuizService{deleteErr: nil}
	handler := NewQuizHandler(svc)

	// Use chi router to inject URL params
	r := chi.NewRouter()
	r.Delete("/quizzes/{id}", handler.Delete)

	req := httptest.NewRequest(http.MethodDelete, "/quizzes/test-id", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", rec.Code)
	}
}

func TestDeleteHandler_NotFound(t *testing.T) {
	svc := &mockQuizService{
		deleteErr: sharedDomain.NewNotFoundError("Quiz not found"),
	}
	handler := NewQuizHandler(svc)

	r := chi.NewRouter()
	r.Delete("/quizzes/{id}", handler.Delete)

	req := httptest.NewRequest(http.MethodDelete, "/quizzes/nonexistent", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}
