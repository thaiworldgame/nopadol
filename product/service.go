package product

func New(repo Repository) (Service){
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	SearchByBarcode(req *SearchByBarcodeTemplate) (interface{}, error)
}


func (s *service)SearchByBarcode(req *SearchByBarcodeTemplate)(interface{}, error){
	resp, err := s.repo.SearchByBarcode(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}