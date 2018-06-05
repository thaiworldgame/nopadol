package sale


import "context"

type Service interface {
	// Create creates new Entity1
	Create(ctx context.Context, entity *Entity1) (entityID string, err error)

	// Update updates Entity1
	Update(ctx context.Context, entity *Entity1) error

	// Search searchs Entity1
	Search(ctx context.Context, keyword *EntitySearch) (docno EntitySaleOrder, err error)
}