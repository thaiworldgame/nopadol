package mariadb

import (
	"github.com/mrtomyum/nopadol/delivery"
	"context"
	"github.com/mrtomyum/nopadol/sale"
	"database/sql"
)

// NewSaleRepository creates sale repository implements by mariadb
func NewDeliveryRepository() delivery.Repository {
	return &deliveryRepo{}
}

type deliveryRepo struct {
	*sql.DB
}

func (deliveryRepo) ReportDeliveryBydate(ctx context.Context, entity *sale.Entity1) (string, error) {
	return "", nil
}



