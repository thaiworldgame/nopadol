package employee

import "context"

type Service interface {
	SearchEmployeeById(ctx context.Context, req *SearchByIdTemplate) (emp EmployeeTemplate, err error)
}
