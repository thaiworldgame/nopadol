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
	LogIn(req *LoginRequest) (interface{}, error)
	SearchListCompany() (interface{}, error)
	SearchListZone(string) (interface{}, error)
	SearchListMachine() (interface{}, error)
	SearchCarBrand(string) (interface{}, error)
	SearchCustomer(string) (interface{}, error)
	SearchItem(string) (interface{}, error)

	PickupNew(req *NewPickupRequest) (interface{}, error)
	CancelQueue(req *PickupCancelRequest) (interface{}, error)
	ManagePickup(req *ManagePickupRequest) (interface{}, error)
	ManageCheckout(req *ManageCheckoutRequest) (interface{}, error)
	ListQueue(req *ListQueueRequest) (interface{}, error)
	PickupEdit(req *PickupEditRequest) (interface{}, error)
	QueueEdit(req *QueueEditRequest) (interface{}, error)
	QueueStatus(req *QueueStatusRequest) (interface{}, error)
	QueueProduct(req *QueueProductRequest) (interface{}, error)
	BillingDone(req *BillingDoneRequest) (interface{}, error)

	ShiftOpen(*ShiftOpenRequest) (interface{}, error)
	ShiftClose(*ShiftCloseRequest) (interface{}, error)
}

func (s *service) SearchListCompany() (interface{}, error) {
	resp, err := s.repo.SearchListCompany()
	if err != nil {
		return nil, err
	}
	fmt.Println("service recive data -> ", resp)

	return resp, nil
}

func (s *service) SearchListZone(access_token string) (interface{}, error) {
	resp, err := s.repo.SearchListZone(access_token)
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

func (s *service) LogIn(req *LoginRequest) (interface{}, error) {
	resp, err := s.repo.LogIn(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service UserLogIn data -> ", resp)
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

func (s *service) CancelQueue(req *PickupCancelRequest) (interface{}, error) {
	resp, err := s.repo.CancelQueue(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service Cancel Queue -> ", resp)
	return resp, nil
}

func (s *service) ManageCheckout(req *ManageCheckoutRequest) (interface{}, error) {
	resp, err := s.repo.ManageCheckout(req)
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

func (s *service) PickupEdit(req *PickupEditRequest) (interface{}, error) {
	resp, err := s.repo.PickupEdit(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service List Queue data -> ", resp)
	return resp, nil
}

func (s *service) QueueEdit(req *QueueEditRequest) (interface{}, error) {
	resp, err := s.repo.QueueEdit(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service List Queue data -> ", resp)
	return resp, nil
}

func (s *service) QueueProduct(req *QueueProductRequest) (interface{}, error) {
	resp, err := s.repo.QueueProduct(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service List Queue data -> ", resp)
	return resp, nil
}

func (s *service) QueueStatus(req *QueueStatusRequest) (interface{}, error) {
	resp, err := s.repo.QueueStatus(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service List Queue data -> ", resp)
	return resp, nil
}

func (s *service) BillingDone(req *BillingDoneRequest) (interface{}, error) {
	resp, err := s.repo.BillingDone(req)
	if err != nil {
		fmt.Println("error service level ", err.Error())
		return nil, err
	}
	fmt.Println("service List Billing data -> ", resp)
	return resp, nil
}
func (s *service) ShiftOpen(req *ShiftOpenRequest) (interface{}, error) {
	return s.repo.ShiftOpen(req)
}

func (s *service) ShiftClose(req *ShiftCloseRequest) (interface{}, error) {
	return s.repo.ShiftClose(req)
}
