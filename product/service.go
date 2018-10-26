package product

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	SearchByBarcode(req *SearchByBarcodeTemplate) (interface{}, error)
	SearchByItemCode(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}

func (s *service) SearchByBarcode(req *SearchByBarcodeTemplate) (interface{}, error) {
	resp, err := s.repo.SearchByBarcode(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchByItemCode(req *SearchByItemCodeTemplate) (interface{}, error) {
	resp, err := s.repo.SearchByItemCode(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchByKeyword(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
