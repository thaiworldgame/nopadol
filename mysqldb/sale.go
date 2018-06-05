package mysqldb

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
	"github.com/jmoiron/sqlx"
	"fmt"
)

// NewDomain1Repository creates domain1 repository implements by domain4
func NewSaleRepository(db *sqlx.DB) sale.Repository {
	return &saleRepository{db}
}

type saleRepository struct {db *sqlx.DB}

func (sr *saleRepository) Register(ctx context.Context, entity *sale.Entity1) (string, error) {
	//sr.db.
	fmt.Println("Entity1 = ",entity.Field1)
	return "moo", nil
}

func (sr *saleRepository) Search(ctx context.Context, kw *sale.EntitySearch) (docno sale.EntitySaleOrder, err error) {
	sql := `select DocNo,ArCode,ArName from SaleOrder where DocNo = ? limit 1`
	err = sr.db.Get(&docno, sql, kw.Keyword)
	if err != nil {
		return docno,err
	}

	fmt.Println("DocNo =",docno)
	return docno,nil
}

// SetField3 sets field3 for Entity1
func (saleRepository) SetField3(ctx context.Context, entityID string, field3 int) error {
	return nil
}
