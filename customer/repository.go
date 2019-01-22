package customer

type Repository interface {
	//SearchCustomerById(ctx context.Context, id *SearchByIdTemplate) (CustomerTemplate, error)
	SearchById(req *SearchByIdTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}
