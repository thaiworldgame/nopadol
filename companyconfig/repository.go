package companyconfig

type Repository interface {
	Create(req *RequestConfigTemplate) (interface{}, error)
	SearchById(req *SearchByIdRequestTemplate) (interface{}, error)
}
