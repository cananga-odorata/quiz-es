package application

import (
	"context"
	"strings"

	"github.com/cananga-odorata/golang-template/internal/modules/quiz/domain"
	sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"
)

// QuizService defines the quiz business logic interface
type QuizService interface {
	GetAll(ctx context.Context) ([]QuizResponse, error)
	Create(ctx context.Context, req CreateQuizRequest) (*QuizResponse, error)
	Delete(ctx context.Context, id string) error
}

type quizService struct {
	repo domain.QuizRepository
}

// NewQuizService creates a new QuizService
func NewQuizService(repo domain.QuizRepository) QuizService {
	return &quizService{repo: repo}
}

// GetAll returns all quizzes ordered by display_order
func (s *quizService) GetAll(ctx context.Context) ([]QuizResponse, error) {
	quizzes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, sharedDomain.NewInternalError("Failed to fetch quizzes", err)
	}

	responses := make([]QuizResponse, len(quizzes))
	for i, q := range quizzes {
		responses[i] = toQuizResponse(q)
	}
	return responses, nil
}

// Create creates a new quiz with auto-assigned display_order
func (s *quizService) Create(ctx context.Context, req CreateQuizRequest) (*QuizResponse, error) {
	// Validate required fields
	if strings.TrimSpace(req.Question) == "" ||
		strings.TrimSpace(req.Choice1) == "" ||
		strings.TrimSpace(req.Choice2) == "" ||
		strings.TrimSpace(req.Choice3) == "" ||
		strings.TrimSpace(req.Choice4) == "" {
		return nil, domain.ErrInvalidQuiz
	}

	// Get the next display_order
	maxOrder, err := s.repo.GetMaxDisplayOrder(ctx)
	if err != nil {
		return nil, sharedDomain.NewInternalError("Failed to get max display order", err)
	}

	quiz := &domain.Quiz{
		ID:           sharedDomain.NewID(),
		Question:     strings.TrimSpace(req.Question),
		Choice1:      strings.TrimSpace(req.Choice1),
		Choice2:      strings.TrimSpace(req.Choice2),
		Choice3:      strings.TrimSpace(req.Choice3),
		Choice4:      strings.TrimSpace(req.Choice4),
		DisplayOrder: maxOrder + 1,
	}

	if err := s.repo.Create(ctx, quiz); err != nil {
		return nil, sharedDomain.NewInternalError("Failed to create quiz", err)
	}

	resp := toQuizResponse(*quiz)
	return &resp, nil
}

// Delete removes a quiz and renumbers remaining quizzes
func (s *quizService) Delete(ctx context.Context, id string) error {
	// Get the quiz to find its display_order
	quiz, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.ErrQuizNotFound
	}

	// Delete the quiz
	if err := s.repo.Delete(ctx, id); err != nil {
		return sharedDomain.NewInternalError("Failed to delete quiz", err)
	}

	// Renumber: decrement display_order for all quizzes above the deleted one
	if err := s.repo.DecrementDisplayOrdersAbove(ctx, quiz.DisplayOrder); err != nil {
		return sharedDomain.NewInternalError("Failed to renumber quizzes", err)
	}

	return nil
}

func toQuizResponse(q domain.Quiz) QuizResponse {
	return QuizResponse{
		ID:           q.ID,
		Question:     q.Question,
		Choice1:      q.Choice1,
		Choice2:      q.Choice2,
		Choice3:      q.Choice3,
		Choice4:      q.Choice4,
		DisplayOrder: q.DisplayOrder,
	}
}
