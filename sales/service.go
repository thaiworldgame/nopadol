package sales

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	CreateQuo(req *NewQuoTemplate) (interface{}, error)
	SearchQuoById() (interface{}, error)
	CreateSale(req *NewSaleTemplate) (interface{}, error)
	SearchSaleById() (interface{}, error)
}

func (s *service)CreateQuo(req *NewQuoTemplate) (interface{}, error){
	resp, err := s.repo.CreateQuo(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service)SearchQueById()(interface{}, error){
	resp, err := s.repo.SearchQuoById()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) CreateSale(req *NewSaleTemplate) (interface{}, error){
	resp, err := s.repo.CreateSale(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service)SearchSaleById()(interface{}, error){
	resp, err := s.repo.SearchSaleById()
	if err != nil {
		return nil, err
	}
	return resp, nil
}