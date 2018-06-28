package customer

import "context"

type Endpoint interface {
	SearchCustomerById(context.Context, *SearchCustomerByIdRequest) (*SearchCustomerResponse, error)
}

type (
	SearchCustomerByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchCustomerResponse struct {
		CustomerId   int64  `json:"customer_id"`
		CustomerCode string `json:"customer_code"`
		CustomerName string `json:"customer_name"`
	}
)
