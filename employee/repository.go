package employee

import "context"

type Repository interface {
	SearchEmployeeById(ctx context.Context, id *SearchById) (EmployeeTemplate, error)
}