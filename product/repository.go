package product

type Repository interface {
	SearchByBarcode(req *SearchByBarcodeTemplate) (interface{}, error)
	SearchByItemCode(req *SearchByItemCodeTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}
