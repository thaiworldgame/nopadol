package companyconfig

type Service interface {
	Create(req *RequestConfigTemplate) (interface{}, error)
	SearchById(req *SearchByIdRequestTemplate) (interface{}, error)
}

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

func (s *service)Create(req *RequestConfigTemplate)(interface{}, error){
	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (s *service)SearchById(req *SearchByIdRequestTemplate)(interface{}, error){
	resp, err := s.repo.SearchById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}