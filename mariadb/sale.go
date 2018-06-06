package mariadb

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
)

// NewSaleRepository creates sale repository implements by mariadb
func NewSaleRepository() sale.Repository {
	return &saleRepository{}
}

type saleRepository struct {}

func (saleRepository) Register(ctx context.Context, entity *sale.Entity1) (string, error) {
	return "", nil
}

// SetField3 sets field3 for Entity1
func (saleRepository) SetField3(ctx context.Context, entityID string, field3 int) error {
	return nil
}

