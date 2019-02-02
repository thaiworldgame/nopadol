package product

import "github.com/mrtomyum/nopadol/auth"

type Repository interface {
	SearchByBarcode(req *SearchByBarcodeTemplate) (interface{}, error)
	SearchByItemCode(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByItemStockLocation(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	StoreItem(req *ProductNewRequest) (interface{},error)
	StoreBarcode(req []BarcodeNewRequest,tk *auth.Token) (interface{},error)
}
