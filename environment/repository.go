package environment

type Repository interface {
	SearchDepartmentById(req *SearchByIdTemplate) (interface{}, error)
	SearchDepartmentByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchProjectById(req *SearchByIdTemplate) (interface{}, error)
	SearchProjectByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchAllocateById(req *SearchByIdTemplate) (interface{}, error)
	SearchAllocateByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchCustContactByIdRepo(req *SearchByIdTemplate) ([]FindCustContactModel, error)
	SearchCustContactByKeywordRepo(req *SearchByKeywordTemplate) ([]FindCustContactModel, error)
}
