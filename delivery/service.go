package delivery

import "fmt"

func NewService(repo Repository) Service {
	s := service{repo}
	return &s
}

type service struct {
	repo Repository
}

type Service interface {
	ReportDaily(req string) (interface{}, error)
	GetTeam() (interface{}, error)
	GetSales() (interface{}, error)
}

// ListUpdateByVending ส่งคืนรายการ Software, Firmware, Data Update กลับไปยัง Vending ที่ร้องขอมา
func (s *service) ReportDaily(req string) (interface{}, error) {
	fmt.Println("begin delivery service ReportDaily")
	fmt.Println("service param is ->", req)
	//s.repo.ListUpdateByVending()
	resp, err := s.repo.ReportDaily(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GetTeam() (interface{}, error) {
	fmt.Println("begin delivery service ReportDaily")
	fmt.Println("service param is ->")
	//s.repo.ListUpdateByVending()
	resp, err := s.repo.GetTeam()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) GetSales() (interface{}, error) {
	fmt.Println("begin delivery service ReportDaily")
	fmt.Println("service param is ->")
	//s.repo.ListUpdateByVending()
	resp, err := s.repo.GetSale()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
