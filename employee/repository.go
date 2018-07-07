package employee

type Repository interface {
	//SearchEmployeeById(ctx context.Context, id *SearchByIdTemplate) (EmployeeTemplate, error)
	SearchById(req *SearchByIdTemplate)(interface{}, error)
}