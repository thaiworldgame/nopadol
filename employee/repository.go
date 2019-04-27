package employee

type Repository interface {
	//SearchEmployeeById(ctx context.Context, id *SearchByIdTemplate) (EmployeeTemplate, error)
	//FindEmployeeById(req *SearchByIdTemplate) (interface{}, error)
	//FindEmployeeByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchById(req *SearchByIdTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}
