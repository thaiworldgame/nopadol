package mariadb

func NewSaleRepository() sale.Repository {
	return &saleRepository{}
}

type saleRepository struct {}

