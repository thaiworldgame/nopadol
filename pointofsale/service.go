package pointofsale

func New(repo Repository) (Service){
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	Create(req *BasketTemplate) (interface{}, error)
	ManageBasket(req *BasketTemplate)(interface{}, error)
}

func (s *service) Create(req *BasketTemplate) (interface{}, error){
	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) ManageBasket(req *BasketTemplate) (interface{}, error){
	resp, err := s.repo.ManageBasket(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
