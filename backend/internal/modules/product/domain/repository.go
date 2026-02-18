package domain

import "context"

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter ProductFilter) ([]*Product, int64, error)
}

type ProductFilter struct {
	TenantID string
	Status   *Status
	Search   string
	Limit    int
	Offset   int
}
