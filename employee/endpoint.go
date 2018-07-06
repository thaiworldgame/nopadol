package employee

import "context"

type Endpoint interface {
	SearchEmployeeById(context.Context, *SearchEmployeeByIdRequest) (*SearchEmployeeResponse, error)
}

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
