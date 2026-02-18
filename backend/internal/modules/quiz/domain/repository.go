package domain

import "context"

// QuizRepository defines the interface for quiz data access
type QuizRepository interface {
	// GetAll returns all quizzes ordered by display_order
	GetAll(ctx context.Context) ([]Quiz, error)

	// GetByID returns a quiz by its ID
	GetByID(ctx context.Context, id string) (*Quiz, error)

	// Create inserts a new quiz
	Create(ctx context.Context, quiz *Quiz) error

	// Delete removes a quiz by its ID
	Delete(ctx context.Context, id string) error

	// GetMaxDisplayOrder returns the current maximum display_order
	GetMaxDisplayOrder(ctx context.Context) (int, error)

	// DecrementDisplayOrdersAbove decrements display_order for all quizzes with order > given value
	DecrementDisplayOrdersAbove(ctx context.Context, order int) error
}
