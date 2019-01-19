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
	SearchItem(string) (interface{}, error)
	ShiftOpen(*ShiftOpenRequest) (interface{},error)
}

func (s *service) SearchListCompany() (interface{}, error) {
	return  s.repo.SearchListCompany()
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("service recive data -> ", resp)
	//
	//return resp, nil
}

func (s *service) SearchListMachine() (interface{}, error) {
	resp, err := s.repo.SearchListMachine()
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service Search Machine data -> ", resp)

	return resp, nil
}

func (s *service) SearchCarBrand(keyword string) (interface{}, error) {
	resp, err := s.repo.SearchCarBrand(keyword)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service Search Car Brand data -> ", resp)

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

func (s *service) SearchItem(keyword string) (interface{}, error){
	resp, err := s.repo.SearchItem(keyword)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service Search Item data -> ", resp)
	return resp, nil
}

func (s *service)ShiftOpen(req *ShiftOpenRequest)(interface{},error){
	return s.repo.ShiftOpen(req)
}