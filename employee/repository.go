package employee

import "context"

type Repository interface {
	SearchEmployeeById(ctx context.Context, id *SearchByIdTemplate) (EmployeeTemplate, error)
}