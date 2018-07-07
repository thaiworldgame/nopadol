package product

type Repository interface {
	SearchByBarcode(req *SearchByBarcodeTemplate)(interface{}, error)
}
