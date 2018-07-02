package customer

import "context"

type Service interface {
	SearchCustomerById(ctx context.Context, req *SearchByIdTemplate) (cust CustomerTemplate, err error)
}
