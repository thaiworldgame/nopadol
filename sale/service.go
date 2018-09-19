package sale

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

// Service is Sale service interface
type Service interface {
	// Create creates new Entity1
	Create(req *NewQTTemplate) (interface{}, error)

}


func (s *service) Create(req *NewQTTemplate)(interface{},error){
	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}