package drivethru

import "fmt"

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)
	SearchCustomer(string) (interface{}, error)
}

func (s *service) SearchListCompany() (interface{}, error) {
	resp, err := s.repo.SearchListCompany()
	if err != nil {
		return nil, err
	}
	fmt.Println("service recive data -> ", resp)

	return resp, nil
}

func (s *service) SearchListMachine() (interface{}, error) {
	resp, err := s.repo.SearchListMachine()
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service SearchListMachine data -> ", resp)

	return resp, nil
}

func (s *service) SearchCarBrand(keyword string) (interface{}, error) {
	resp, err := s.repo.SearchCarBrand(keyword)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service SearchListMachine data -> ", resp)

	return resp, nil
}

func (s *service) SearchCustomer(keyword string) (interface{}, error) {
	resp, err := s.repo.SearchCustomer(keyword)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service SearchListCustommer data -> ", resp)
	return resp, nil
}

