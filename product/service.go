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
	SearchByItemStockLocation(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	StoreItem(req *ProductNewRequest)(interface{},error)
	StoreBarcode(req *BarcodeNewRequest) (interface{},error)
	StorePrice(req *PriceTemplate)(interface{},error)
	StorePackingRate(req *PackingRate)(interface{},error)
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

func (s *service) SearchByItemStockLocation(req *SearchByItemCodeTemplate) (interface{}, error) {
	resp, err := s.repo.SearchByItemStockLocation(req)
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

func (s *service) StoreItem(req *ProductNewRequest)(interface{},error){
	return s.repo.StoreItem(req)
}

func (s *service) StoreBarcode(req *BarcodeNewRequest)(interface{},error){
	return s.repo.StoreBarcode(req)

	//return nil,nil
}

func (s *service) StorePrice(req *PriceTemplate) (interface{},error){
	return s.repo.StorePrice(req)
	//return nil,nil
}

func (s *service) StorePackingRate(req *PackingRate)(interface{},error){
	return s.repo.StorePackingRate(req)
}
