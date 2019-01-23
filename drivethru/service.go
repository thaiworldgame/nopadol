package drivethru

import "fmt"

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	UserLogIn(req *UserLogInRequest) (interface{}, error)
	SearchListCompany() (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)
	SearchCustomer(string) (interface{}, error)
	SearchItem(string) (interface{}, error)
	PickupNew(req *NewPickupRequest) (interface{}, error)
	ManagePickup(req *ManagePickupRequest) (interface{}, error)
	ListQueue(req *ListQueueRequest) (interface{}, error)
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
	fmt.Println("service SearchCarBrand data -> ", resp)

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

func (s *service) SearchItem(keyword string) (interface{}, error) {
	resp, err := s.repo.SearchItem(keyword)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service SearchItem data -> ", resp)
	return resp, nil
}

func (s *service) UserLogIn(req *UserLogInRequest) (interface{}, error) {
	resp, err := s.repo.UserLogIn(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service UserLogIn data -> ", resp)
	return resp, nil
}

func (s *service) PickupNew(req *NewPickupRequest) (interface{}, error) {
	resp, err := s.repo.PickupNew(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service Pickup New data -> ", resp)
	return resp, nil
}

func (s *service) ManagePickup(req *ManagePickupRequest) (interface{}, error) {
	resp, err := s.repo.ManagePickup(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service Pickup Manage Data -> ", resp)
	return resp, nil
}

func (s *service) ListQueue(req *ListQueueRequest) (interface{}, error) {
	resp, err := s.repo.ListQueue(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service List Queue data -> ", resp)
	return resp, nil
}
