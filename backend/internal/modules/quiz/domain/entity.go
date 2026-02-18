package domain

import "time"

// Quiz represents a quiz question entity
type Quiz struct {
	ID           string    `json:"id" db:"id"`
	Question     string    `json:"question" db:"question"`
	Choice1      string    `json:"choice1" db:"choice1"`
	Choice2      string    `json:"choice2" db:"choice2"`
	Choice3      string    `json:"choice3" db:"choice3"`
	Choice4      string    `json:"choice4" db:"choice4"`
	DisplayOrder int       `json:"display_order" db:"display_order"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
