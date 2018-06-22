package incentive

import "context"

type Service interface {
	// Search searchs Entity1
	SearchSaleCode(ctx context.Context, keyword *EntitySearch) (sin SaleCode, err error)
}