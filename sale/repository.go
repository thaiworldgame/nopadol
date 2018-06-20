package sale

import (
	"context"
)


// Repository is the domain1 storage
type Repository interface {
	Search(ctx context.Context, kw *EntitySearch) (so SaleOrderTemplate, err error)
	
	NewSaleOrder(ctx context.Context, so *SaleOrderTemplate) (Id int64, err error)

	Register(ctx context.Context, entity *Entity1) (entityID string, err error)



	// SetField3 sets field3 for Entity1
	//SetField3(ctx context.Context, entityID string, field3 int) error
}