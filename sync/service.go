package sync

//import "github.com/mrtomyum/nopadol/sales"

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	GetNewQoutaion() (interface{},error)
}



func (s *service)GetNewQoutaion()(interface{},error){
	resp,err := s.repo.GetNewQoutaion()
	if err != nil {
		return nil,err
	}
	return resp,nil
}