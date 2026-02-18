package infrastructure

import (
	"context"
	"database/sql"

	"github.com/cananga-odorata/golang-template/internal/infra/database"
	"github.com/cananga-odorata/golang-template/internal/modules/quiz/domain"
	"github.com/jmoiron/sqlx"
)

type postgresQuizRepository struct {
	db *sqlx.DB
}

// NewPostgresQuizRepository creates a new PostgreSQL quiz repository
func NewPostgresQuizRepository(db *sqlx.DB) domain.QuizRepository {
	return &postgresQuizRepository{db: db}
}

func (r *postgresQuizRepository) getQueryable(ctx context.Context) database.Queryable {
	return database.GetQueryable(ctx, r.db)
}

// GetAll returns all quizzes ordered by display_order
func (r *postgresQuizRepository) GetAll(ctx context.Context) ([]domain.Quiz, error) {
	var quizzes []domain.Quiz
	query := `SELECT id, question, choice1, choice2, choice3, choice4, display_order, created_at, updated_at
	           FROM quizzes ORDER BY display_order ASC`
	q := r.getQueryable(ctx)
	err := q.SelectContext(ctx, &quizzes, query)
	if err != nil {
		return nil, err
	}
	if quizzes == nil {
		quizzes = []domain.Quiz{}
	}
	return quizzes, nil
}

// GetByID returns a quiz by its ID
func (r *postgresQuizRepository) GetByID(ctx context.Context, id string) (*domain.Quiz, error) {
	var quiz domain.Quiz
	query := `SELECT id, question, choice1, choice2, choice3, choice4, display_order, created_at, updated_at
	           FROM quizzes WHERE id = $1`
	q := r.getQueryable(ctx)
	err := q.GetContext(ctx, &quiz, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrQuizNotFound
	}
	return &quiz, err
}

// Create inserts a new quiz
func (r *postgresQuizRepository) Create(ctx context.Context, quiz *domain.Quiz) error {
	query := `INSERT INTO quizzes (id, question, choice1, choice2, choice3, choice4, display_order, created_at, updated_at)
	           VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())`
	q := r.getQueryable(ctx)
	_, err := q.ExecContext(ctx, query, quiz.ID, quiz.Question, quiz.Choice1, quiz.Choice2, quiz.Choice3, quiz.Choice4, quiz.DisplayOrder)
	return err
}

// Delete removes a quiz by its ID
func (r *postgresQuizRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM quizzes WHERE id = $1`
	q := r.getQueryable(ctx)
	result, err := q.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrQuizNotFound
	}
	return nil
}

// GetMaxDisplayOrder returns the current maximum display_order (0 if no quizzes)
func (r *postgresQuizRepository) GetMaxDisplayOrder(ctx context.Context) (int, error) {
	var maxOrder sql.NullInt64
	query := `SELECT MAX(display_order) FROM quizzes`
	q := r.getQueryable(ctx)
	err := q.GetContext(ctx, &maxOrder, query)
	if err != nil {
		return 0, err
	}
	if !maxOrder.Valid {
		return 0, nil
	}
	return int(maxOrder.Int64), nil
}

// DecrementDisplayOrdersAbove decrements display_order for all quizzes with order > given value
func (r *postgresQuizRepository) DecrementDisplayOrdersAbove(ctx context.Context, order int) error {
	query := `UPDATE quizzes SET display_order = display_order - 1, updated_at = NOW() WHERE display_order > $1`
	q := r.getQueryable(ctx)
	_, err := q.ExecContext(ctx, query, order)
	return err
}
