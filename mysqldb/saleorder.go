package mysqldb

import (
	"github.com/mrtomyum/nopadol/sale"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type SaleOrderModel struct {
	DocNo   string       `db:"DocNo"`
	DocDate string       `db:"DocDate"`
	ArCode  string       `db:"ArCode"`
	ArName  string       `db:"ArName"`
	Subs    []*SubsModel `db:"subs"`
}

type SubsModel struct {
	ItemCode string  `db:"ItemCode"`
	ItemName string  `db:"ItemName"`
	Qty      float64 `db:"Qty"`
	UnitCode string  `db:"UnitCode"`
}

type saleRepository struct{ db *sqlx.DB }

// NewDomain1Repository creates domain1 repository implements by domain4
func NewSaleRepository(db *sqlx.DB) sale.Repository {
	return &saleRepository{db}
}


func (sr *saleRepository) Search(ctx context.Context, kw *sale.EntitySearch) (resp sale.SaleOrderTemplate, err error) {
	fmt.Println("From API Keyword =", kw.Keyword)

	SO1 := SaleOrderModel{}
	//sql := `select DocNo,DocDate,ArCode,ArName from SaleOrder where DocNo = ?`
	sql := `select ifnull(docNo,'') as DocNo,ifnull(docDate,'') as DocDate,ifnull(customerCode,'') as ArCode,'moo' as ArName from Queue where docNo = ?`
	//err = sr.db.Get(&so, sql, kw.Keyword)
	err = sr.db.Get(&SO1, sql, kw.Keyword)
	if err != nil {
		return resp, err
	}

	//fmt.Println("From API ArName =",so.ArName)
	fmt.Println("From API DocNo =", SO1.DocNo)

	submodel := []SubsModel{}
	//sqlsub := `select ItemCode,ItemName,Qty,UnitCode from SaleOrderSub where DocNo = ?`
	sqlsub := `select itemCode as ItemCode,itemName as ItemName,qty as Qty,unitCode as UnitCode from QItem where docNo = ?`
	err = sr.db.Select(&submodel, sqlsub, SO1.DocNo)
	if err != nil {
		return resp, err
	}

	//fmt.Println("sub =", so1.Subs)

	Resp := buildsaleorder(SO1)

	for _, v := range submodel {
		//fmt.Println(v)
		soline := buildsaleordersub(v)
		Resp.Subs = append(Resp.Subs, soline)

		//fmt.Println("So line = ",soline)
	}

	//fmt.Println("DocNo =", Resp.DocNo)
	//fmt.Println("DocDate =", Resp.DocDate)
	//fmt.Println("ArCode =", Resp.ArCode)
	//fmt.Println("so =",Resp)

	return Resp, nil
}

func (sr *saleRepository) NewSaleOrder(ctx context.Context, so *sale.SaleOrderTemplate) (Id int64, err error) {
	fmt.Println("data = ", so)

	//sql := `insert into SaleOrder(DocNo,DocDate,ArCode,ArName) values(?,?,?,?)`
	if (so.DocNo != "") {
		sql := `insert into Queue(docNo,customerCode,customerName) values(?,?,?)`
		_, err = sr.db.Exec(sql, so.DocNo, so.ArCode, so.ArName)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}

		for _, sub := range so.Subs {
			//sqlsub := `insert into SaleOrderSub(DocNo,DocDate,ItemCode,ItemName,Qty,UnitCode) values(?,?,?,?,?,?)`
			sqlsub := `insert into QItem(docNo,docDate,itemCode,itemName,qty,unitCode) values(?,?,?,?,?,?)`
			_, err = sr.db.Exec(sqlsub, so.DocNo, so.DocDate, sub.ItemCode, sub.ItemName, sub.Qty, sub.UnitCode)
			if err != nil {
				fmt.Println(err.Error())
				return 0, err
			}
		}

	}

	return 0, nil
}

func (sr *saleRepository) Register(ctx context.Context, entity *sale.Entity1) (string, error) {
	//sr.db.
	fmt.Println("Entity1 = ", entity.Field1)
	return "moo", nil
}


func buildsaleorder(x SaleOrderModel) sale.SaleOrderTemplate {
	var subs []sale.SubsTemplate
	return sale.SaleOrderTemplate{
		DocNo:   x.DocNo,
		DocDate: x.DocDate,
		ArCode:  x.ArCode,
		ArName:  x.ArName,
		Subs:    subs,
	}
}

func buildsaleordersub(x SubsModel) sale.SubsTemplate {
	return sale.SubsTemplate{
		ItemCode: x.ItemCode,
		ItemName: x.ItemName,
		Qty:      x.Qty,
		UnitCode: x.UnitCode,
	}
}
