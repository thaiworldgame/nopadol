package mysqldb

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

// NewDomain1Repository creates domain1 repository implements by domain4
func NewSaleRepository(db *sqlx.DB) sale.Repository {
	return &saleRepository{db}
}

type saleRepository struct {db *sqlx.DB}

func (sr *saleRepository) NewSO(ctx context.Context, so *sale.SaleOrder)(Id int64, err error){
	return 0,nil
}

func (sr *saleRepository) Register(ctx context.Context, entity *sale.Entity1) (string, error) {
	//sr.db.
	fmt.Println("Entity1 = ",entity.Field1)
	return "moo", nil
}

func (sr *saleRepository) Search(ctx context.Context, kw *sale.EntitySearch) (so sale.SaleOrder, err error) {
	fmt.Println("From API Keyword =",kw.Keyword)
	sql := `select DocNo,ArCode,ArName from SaleOrder where DocNo = ?`
	err = sr.db.Get(&so, sql, kw.Keyword)
	if err != nil {
		return so,err
	}

	//fmt.Println("From API ArName =",so.ArName)
	fmt.Println("From API DocNo =",so.DocNo)

	sqlsub := `select ItemCode,ItemName,Qty,UnitCode from SaleOrderSub where DocNo = ?`
	err = sr.db.Select(&so.Subs,sqlsub,so.DocNo)
	if err != nil {
		return so, err
	}



	fmt.Println("DocNo =",so.DocNo)
	return so,nil
}

// SetField3 sets field3 for Entity1
func (saleRepository) SetField3(ctx context.Context, entityID string, field3 int) error {
	return nil
}
