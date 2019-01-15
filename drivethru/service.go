package drivethru

func New(repo Repository) (Service){
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	SearchListCompany() (interface{}, error)
}

func (s *service) SearchListCompany() (interface{}, error){
	resp, err := s.repo.SearchListCompany()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
