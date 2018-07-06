package product

import "context"

type Service interface {
	SearchProductByBarcode(ctx context.Context, req *SearchByBarcodeTemplate) (resp ProductTemplate, err error)
}
