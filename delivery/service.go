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
	ReportDaily(string) (interface{}, error)
}

// ListUpdateByVending ส่งคืนรายการ Software, Firmware, Data Update กลับไปยัง Vending ที่ร้องขอมา
func (s *service) ReportDaily(date string) (interface{}, error) {
	fmt.Println("begin delivery service ReportDaily")
	//s.repo.ListUpdateByVending()
	resp, err := s.repo.ReportDaily()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
