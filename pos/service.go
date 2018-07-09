package pos

import "fmt"

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	Create(req *NewPosTemplate) (interface{}, error)
	SearchById(req *SearchPosByIdRequestTemplate) (interface{}, error)
}

func (s *service)Create(req *NewPosTemplate) (interface{}, error){
	fmt.Println("Service")
	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service)SearchById(req *SearchPosByIdRequestTemplate)(interface{}, error){
	resp, err := s.repo.SearchById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
