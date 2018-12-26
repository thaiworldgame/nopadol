package dataimport

import "github.com/mrtomyum/nopadol/product"

type Repository interface {
	ProductUpdate(product.ProductTemplate)(string,error)
}
