package application

// CreateQuizRequest DTO for creating a new quiz
type CreateQuizRequest struct {
	Question string `json:"question"`
	Choice1  string `json:"choice1"`
	Choice2  string `json:"choice2"`
	Choice3  string `json:"choice3"`
	Choice4  string `json:"choice4"`
}

// QuizResponse DTO for quiz responses
type QuizResponse struct {
	ID           string `json:"id"`
	Question     string `json:"question"`
	Choice1      string `json:"choice1"`
	Choice2      string `json:"choice2"`
	Choice3      string `json:"choice3"`
	Choice4      string `json:"choice4"`
	DisplayOrder int    `json:"display_order"`
}
