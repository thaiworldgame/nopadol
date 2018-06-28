package employee

import "context"

type Service interface {
	SearchEmployeeById(ctx context.Context, req *SearchById) (emp EmployeeTemplate, err error)
}
