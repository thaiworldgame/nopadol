package mysqldb

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
)

// NewDomain1Repository creates domain1 repository implements by domain4
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
