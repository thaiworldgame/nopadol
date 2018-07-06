package customer

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	//SearchCustomerById(ctx context.Context, req *SearchByIdTemplate) (cust CustomerTemplate, err error)
	SearchById(req *SearchByIdTemplate) (interface{}, error)
}
func (s *service) SearchById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
