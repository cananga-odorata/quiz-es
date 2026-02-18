package application

import (
	"context"
	"testing"

	"github.com/cananga-odorata/golang-template/internal/modules/quiz/domain"
)

// mockQuizRepository is a mock implementation of domain.QuizRepository
type mockQuizRepository struct {
	quizzes         []domain.Quiz
	getMaxOrderResp int
	getMaxOrderErr  error
	createErr       error
	deleteErr       error
	decrementErr    error
	getByIDResp     *domain.Quiz
	getByIDErr      error
}

func newMockRepo() *mockQuizRepository {
	return &mockQuizRepository{
		quizzes: []domain.Quiz{},
	}
}

func (m *mockQuizRepository) GetAll(_ context.Context) ([]domain.Quiz, error) {
	return m.quizzes, nil
}

func (m *mockQuizRepository) GetByID(_ context.Context, id string) (*domain.Quiz, error) {
	if m.getByIDErr != nil {
		return nil, m.getByIDErr
	}
	if m.getByIDResp != nil {
		return m.getByIDResp, nil
	}
	for _, q := range m.quizzes {
		if q.ID == id {
			return &q, nil
		}
	}
	return nil, domain.ErrQuizNotFound
}

func (m *mockQuizRepository) Create(_ context.Context, quiz *domain.Quiz) error {
	if m.createErr != nil {
		return m.createErr
	}
	m.quizzes = append(m.quizzes, *quiz)
	return nil
}

func (m *mockQuizRepository) Delete(_ context.Context, id string) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	newQuizzes := []domain.Quiz{}
	for _, q := range m.quizzes {
		if q.ID != id {
			newQuizzes = append(newQuizzes, q)
		}
	}
	m.quizzes = newQuizzes
	return nil
}

func (m *mockQuizRepository) GetMaxDisplayOrder(_ context.Context) (int, error) {
	if m.getMaxOrderErr != nil {
		return 0, m.getMaxOrderErr
	}
	return m.getMaxOrderResp, nil
}

func (m *mockQuizRepository) DecrementDisplayOrdersAbove(_ context.Context, order int) error {
	if m.decrementErr != nil {
		return m.decrementErr
	}
	for i := range m.quizzes {
		if m.quizzes[i].DisplayOrder > order {
			m.quizzes[i].DisplayOrder--
		}
	}
	return nil
}

// ============ Test Cases ============

func TestCreateQuiz_Success(t *testing.T) {
	repo := newMockRepo()
	repo.getMaxOrderResp = 0
	service := NewQuizService(repo)

	req := CreateQuizRequest{
		Question: "ข้อใดต่างจากข้ออื่น",
		Choice1:  "3",
		Choice2:  "5",
		Choice3:  "9",
		Choice4:  "11",
	}

	resp, err := service.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if resp.DisplayOrder != 1 {
		t.Errorf("expected display_order 1, got %d", resp.DisplayOrder)
	}
	if resp.Question != "ข้อใดต่างจากข้ออื่น" {
		t.Errorf("expected question 'ข้อใดต่างจากข้ออื่น', got '%s'", resp.Question)
	}
	if resp.Choice1 != "3" || resp.Choice2 != "5" || resp.Choice3 != "9" || resp.Choice4 != "11" {
		t.Error("choices do not match expected values")
	}
	if len(repo.quizzes) != 1 {
		t.Errorf("expected 1 quiz in repo, got %d", len(repo.quizzes))
	}
}

func TestCreateQuiz_AutoIncrementDisplayOrder(t *testing.T) {
	repo := newMockRepo()
	repo.getMaxOrderResp = 3
	service := NewQuizService(repo)

	req := CreateQuizRequest{
		Question: "X + 2 = 4 จงหาค่า X",
		Choice1:  "1",
		Choice2:  "2",
		Choice3:  "3",
		Choice4:  "4",
	}

	resp, err := service.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if resp.DisplayOrder != 4 {
		t.Errorf("expected display_order 4, got %d", resp.DisplayOrder)
	}
}

func TestCreateQuiz_ValidationError_EmptyQuestion(t *testing.T) {
	repo := newMockRepo()
	service := NewQuizService(repo)

	req := CreateQuizRequest{
		Question: "",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	}

	_, err := service.Create(context.Background(), req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestCreateQuiz_ValidationError_EmptyChoice(t *testing.T) {
	repo := newMockRepo()
	service := NewQuizService(repo)

	req := CreateQuizRequest{
		Question: "What is 1+1?",
		Choice1:  "1",
		Choice2:  "",
		Choice3:  "3",
		Choice4:  "4",
	}

	_, err := service.Create(context.Background(), req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestCreateQuiz_ValidationError_WhitespaceOnly(t *testing.T) {
	repo := newMockRepo()
	service := NewQuizService(repo)

	req := CreateQuizRequest{
		Question: "   ",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	}

	_, err := service.Create(context.Background(), req)
	if err == nil {
		t.Fatal("expected validation error for whitespace-only question, got nil")
	}
}

func TestGetAll_Empty(t *testing.T) {
	repo := newMockRepo()
	service := NewQuizService(repo)

	resp, err := service.GetAll(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(resp) != 0 {
		t.Errorf("expected 0 quizzes, got %d", len(resp))
	}
}

func TestGetAll_WithData(t *testing.T) {
	repo := newMockRepo()
	repo.quizzes = []domain.Quiz{
		{ID: "1", Question: "Q1", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 1},
		{ID: "2", Question: "Q2", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 2},
	}
	service := NewQuizService(repo)

	resp, err := service.GetAll(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(resp) != 2 {
		t.Errorf("expected 2 quizzes, got %d", len(resp))
	}
}

func TestDeleteQuiz_Success_WithRenumber(t *testing.T) {
	repo := newMockRepo()
	repo.quizzes = []domain.Quiz{
		{ID: "a", Question: "Q1", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 1},
		{ID: "b", Question: "Q2", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 2},
		{ID: "c", Question: "Q3", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 3},
	}
	service := NewQuizService(repo)

	// Delete quiz #2 (display_order=2)
	err := service.Delete(context.Background(), "b")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Should have 2 quizzes remaining
	if len(repo.quizzes) != 2 {
		t.Fatalf("expected 2 quizzes, got %d", len(repo.quizzes))
	}

	// Quiz "c" should now have display_order=2 (was 3, decremented because > 2)
	for _, q := range repo.quizzes {
		if q.ID == "c" && q.DisplayOrder != 2 {
			t.Errorf("expected quiz 'c' to have display_order 2 after renumber, got %d", q.DisplayOrder)
		}
		if q.ID == "a" && q.DisplayOrder != 1 {
			t.Errorf("expected quiz 'a' to still have display_order 1, got %d", q.DisplayOrder)
		}
	}
}

func TestDeleteQuiz_NotFound(t *testing.T) {
	repo := newMockRepo()
	repo.quizzes = []domain.Quiz{
		{ID: "a", Question: "Q1", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 1},
	}
	service := NewQuizService(repo)

	err := service.Delete(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error for non-existent quiz, got nil")
	}
}

func TestDeleteQuiz_LastItem(t *testing.T) {
	repo := newMockRepo()
	repo.quizzes = []domain.Quiz{
		{ID: "a", Question: "Q1", Choice1: "A", Choice2: "B", Choice3: "C", Choice4: "D", DisplayOrder: 1},
	}
	service := NewQuizService(repo)

	err := service.Delete(context.Background(), "a")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(repo.quizzes) != 0 {
		t.Errorf("expected 0 quizzes, got %d", len(repo.quizzes))
	}
}
