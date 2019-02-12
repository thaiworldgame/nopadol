package customer

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	//SearchCustomerById(ctx context.Context, req *SearchByIdTemplate) (cust CustomerTemplate, err error)
	SearchById(req *SearchByIdTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	StoreCustomer(req *CustomerTemplate) (interface{}, error)
}

func (s *service) SearchById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchByKeyword(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) StoreCustomer(req *CustomerTemplate) (interface{}, error){
	return nil,nil
}