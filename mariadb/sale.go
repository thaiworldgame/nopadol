package mariadb

import (
	"github.com/mrtomyum/nopadol/sale"
)
func NewSaleRepository() sale.Repository {
	return &saleRepository{}
}

type saleRepository struct {}

