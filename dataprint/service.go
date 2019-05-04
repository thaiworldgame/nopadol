package dataprint


func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	DataPrint() (interface{}, error)
}

func (s *service) DataPrint() (interface{}, error) {
	resp, err := s.repo.DataPrint()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
