package domain

import (
	"time"

	sharedDomain "github.com/cananga-odorata/golang-template/internal/shared/domain"
)

type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int64     `json:"price" db:"price"`
	Commission  int       `json:"commission" db:"commission"`
	Status      Status    `json:"status" db:"status"`
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDraft    Status = "draft"
)

func NewProduct(name, description string, price int64, commission int, tenantID string) *Product {
	now := time.Now()
	return &Product{
		ID:          sharedDomain.NewID(),
		Name:        name,
		Description: description,
		Price:       price,
		Commission:  commission,
		Status:      StatusDraft,
		TenantID:    tenantID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
