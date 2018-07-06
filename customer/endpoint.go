package customer

import (
	"context"
	"fmt"
)

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

func SearchById(s Service) interface{} {
	return func(ctx context.Context, req *SearchCustomerByIdRequest) (interface{}, error) {
		resp, err := s.SearchById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
