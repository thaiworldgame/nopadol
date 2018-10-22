package environment

type Repository interface {
	SearchDepartmentById (req *SearchByIdTemplate) (interface{}, error)
	SearchDepartmentByKeyword (req *SearchByKeywordTemplate) (interface{}, error)
}
