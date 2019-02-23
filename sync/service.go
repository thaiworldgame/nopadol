package sync

import "github.com/mrtomyum/nopadol/sales"

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	NewQuotation() (sales.NewQuoTemplate,error)
}


func (s *service)NewQuotation()(sales.NewQuoTemplate,error){
	resp,err := s.repo.GetNewQoutation()
	if err != nil {
		return nil,err
	}
	return resp,nil
}