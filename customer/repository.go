package customer

import (
	"context"
)

type Repository interface {
	SearchCustomerById(ctx context.Context, req *SearchByIdTemplate)(CustomerTemplate, error)
}
