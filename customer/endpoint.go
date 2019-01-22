package customer

import (
	"context"
	"fmt"
)

//type Endpoint interface {
//	SearchCustomerById(context.Context, *SearchCustomerByIdRequest) (*SearchCustomerResponse, error)
//}

type (
	SearchCustomerByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchCustomerResponse struct {
		Id         int64  `json:"id"`
		Code       string `json:"code"`
		Name       string `json:"name"`
		Address    string `json:"address"`
		Telephone  string `json:"telephone"`
		BillCredit int64  `json:"bill_credit"`
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

func SearchByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
