package customer

import "context"

type Service interface {
	SearchCustomerById(ctx context.Context, req *SearchById) (cust CustomerTemplate, err error)
}
