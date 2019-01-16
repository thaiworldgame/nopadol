package delivery

import "fmt"

func NewService(repo Repository) (Service) {
	s := service{repo}
	return &s
}

type service struct {
	repo Repository
}

type Service interface {
	ReportDaily(req string) (interface{}, error)
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
