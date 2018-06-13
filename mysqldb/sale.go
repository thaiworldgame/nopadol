package mysqldb

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type saleRepository struct {db *sqlx.DB}

// NewDomain1Repository creates domain1 repository implements by domain4
func NewSaleRepository(db *sqlx.DB) sale.Repository {
	return &saleRepository{db}
}

func (sr *saleRepository) NewSO(ctx context.Context, so *sale.SaleOrder)(Id int64, err error){
	fmt.Println("data = ",so.DocNo, so.DocDate, so.ArCode, so.ArName)
	//sql := `insert into SaleOrder(DocNo,DocDate,ArCode,ArName) values(?,?,?,?)`
	sql := `insert into Queue(docNo,customerCode,customerName) values(?,?,?)`
	_, err = sr.db.Exec(sql, so.DocNo, so.ArCode, so.ArName)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	for _, sub := range so.Subs{
		//sqlsub := `insert into SaleOrderSub(DocNo,DocDate,ItemCode,ItemName,Qty,UnitCode) values(?,?,?,?,?,?)`
		sqlsub := `insert into QItem(docNo,docDate,itemCode,itemName,qty,unitCode) values(?,?,?,?,?,?)`
		_, err = sr.db.Exec(sqlsub,so.DocNo,so.DocDate,sub.ItemCode,sub.ItemName,sub.Qty,sub.UnitCode)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}
	return 0,nil
}

func (sr *saleRepository) Register(ctx context.Context, entity *sale.Entity1) (string, error) {
	//sr.db.
	fmt.Println("Entity1 = ",entity.Field1)
	return "moo", nil
}

func (sr *saleRepository) Search(ctx context.Context, kw *sale.EntitySearch) (so sale.SaleOrder, err error) {
	fmt.Println("From API Keyword =",kw.Keyword)
	//sql := `select DocNo,DocDate,ArCode,ArName from SaleOrder where DocNo = ?`
	sql := `select docNo as DocNo,docDate as DocDate,customerCode as ArCode,'moo' as ArName from Queue where docNo = 'S01-QUE590104-0005'`
	//err = sr.db.Get(&so, sql, kw.Keyword)
	err = sr.db.Get(&so, sql)
	if err != nil {
		return so,err
	}

	//fmt.Println("From API ArName =",so.ArName)
	fmt.Println("From API DocNo =",so.DocNo)

	//sqlsub := `select ItemCode,ItemName,Qty,UnitCode from SaleOrderSub where DocNo = ?`
	sqlsub := `select itemCode as ItemCode,itemName as ItemName,qty as Qty,unitCode as UnitCode from QItem where docNo = ?`
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
