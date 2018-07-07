package employee

import (
	"context"
	"fmt"
)

//type Endpoint interface {
//	SearchEmployeeById(context.Context, *SearchEmployeeByIdRequest) (*SearchEmployeeResponse, error)
//}

type (
	SearchEmployeeByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchEmployeeResponse struct {
		EmployeeId   int64  `json:"employee_id"`
		EmployeeCode string `json:"employee_code"`
		EmployeeName string `json:"employee_name"`
	}
)

func SearchById(s Service) interface{} {
	return func(ctx context.Context, req *SearchEmployeeByIdRequest) (interface{}, error) {
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
