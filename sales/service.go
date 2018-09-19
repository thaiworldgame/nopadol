package sales

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	Create(req *PosConfigTemplate1) (interface{}, error)
	SearchById() (interface{}, error)
}

func (s *service)Create(req *PosConfigTemplate1) (interface{}, error){
	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (s *service)SearchById()(interface{}, error){
	resp, err := s.repo.SearchById()
	if err != nil {
		return nil, err
	}
	return resp, nil
}