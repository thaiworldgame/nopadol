package product

import "context"

type Repository interface {
	SearchProductByBarcode(ctx context.Context, req *SearchByBarcodeTemplate)(ProductTemplate, error)
}
