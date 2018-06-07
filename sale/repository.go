package sale

import (
	"context"
)


// Repository is the domain1 storage
type Repository interface {
	NewSO(ctx context.Context, so *SaleOrder) (Id int64, err error)
	// Registers inserts given Entity1 into storage
	Register(ctx context.Context, entity *Entity1) (entityID string, err error)

	// Registers inserts given Entity1 into storage
	Search(ctx context.Context, kw *EntitySearch) (so SaleOrder, err error)

	// SetField3 sets field3 for Entity1
	SetField3(ctx context.Context, entityID string, field3 int) error
}