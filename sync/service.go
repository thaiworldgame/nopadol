package sync

import "fmt"

//import "github.com/mrtomyum/nopadol/sales"

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	GetNewQuotaion() (interface{},error)
	GetNewSaleOrder() (interface{}, error)
	ConfirmTransfer(*Logs) (interface{},error)
}

func (s *service)GetNewQuotaion()(interface{},error){
	resp,err := s.repo.GetNewQuotaion()
	if err != nil {
		return nil,err
	}
	return resp,nil
}

func (s *service)GetNewSaleOrder()(interface{},error){
	resp,err := s.repo.GetNewSaleOrder()
	if err != nil {
		return nil,err
	}
	return resp,nil
}

func (s *service)ConfirmTransfer(req *Logs)(interface{}, error){
	resp, err := s.repo.ConfirmTransfer(req)
	if err != nil {
		fmt.Println("Error Service = ", err.Error())
		return nil, err
	}
	return resp, nil
}