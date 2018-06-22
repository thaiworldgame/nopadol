package customer

import "context"

type Service interface {
	SearchCustomerById(ctx context.Context, cust *SearchById) (id CustomerTemplate, err error)
}
