package domain

import sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"

var (
	ErrQuizNotFound = sharedDomain.NewNotFoundError("Quiz not found")
	ErrInvalidQuiz  = sharedDomain.NewValidationError("Question and all 4 choices are required")
)
