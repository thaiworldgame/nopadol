package product

//import "github.com/mrtomyum/nopadol/auth"

type Repository interface {
	SearchByBarcode(req *SearchByBarcodeTemplate) (interface{}, error)
	SearchByItemCode(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByItemStockLocation(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	StoreItem(req *ProductNewRequest) (interface{},error)
	StoreBarcode(req *BarcodeNewRequest) (interface{},error)
	StorePrice(req *PriceTemplate) (interface{},error)
	StorePackingRate(req *PackingRate)(interface{},error)
}

