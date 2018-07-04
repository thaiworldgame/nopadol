package mariadb

import (
	"github.com/mrtomyum/nopadol/delivery"
	"context"
	"github.com/mrtomyum/nopadol/sale"
)

// NewSaleRepository creates sale repository implements by mariadb
func NewDeliveryRepository() delivery.Repository {
	return &deliveryRepo{}
}

type deliveryRepo struct{}

func (deliveryRepo) ReportDeliveryBydate(ctx context.Context, entity *sale.Entity1) (string, error) {
	return "", nil
}



