package sync

//import "github.com/mrtomyum/nopadol/sales"

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	GetNewQuotaion() (interface{},error)
	ConfirmTransfer(Log) (interface{},error)
}

func (s *service)GetNewQuotaion()(interface{},error){
	resp,err := s.repo.GetNewQuotaion()
	if err != nil {
		return nil,err
	}
	return resp,nil
}

func (s *service)ConfirmTransfer(req Log)(interface{}, error){
	resp, err := s.repo.ConfirmTransfer(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}