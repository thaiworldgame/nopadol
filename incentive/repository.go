package incentive

import (
	"context"
)

// Repository is the domain1 storage
type Repository interface {
	SearchSaleCode(ctx context.Context, kw *EntitySearch) (sic SaleCode, err error)
}