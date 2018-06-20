package customer

import (
	"context"
)

type Repository interface {
	SearchCustomerById(ctx context.Context, req *SearchById)(CustomerTemplate, error)
}
