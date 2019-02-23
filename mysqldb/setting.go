package mysqldb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/sales"
)

type Setting struct {
	Id             int64  `json:"company_id"`
	BranchId       int64  `json:"branch_id"`
	TaxType        int64  `json:"tax_type"`
	TaxRate        int64  `json:"tax_rate"`
	LogoPath       string `json:"logo_path"`
	DepartId       int64  `json:"depart_id"`
	DefSaleWhId    int64  `json:"def_sale_wh_id"`
	DefSaleShelfId int64  `json:"def_sale_shelf_id"`
	DefBuyWhId     int64  `json:"def_buy_wh_id"`
	DefBuyShelfId  int64  `json:"def_buy_shelf_id"`
	StockStatus    int64  `json:"stock_status"`
	SaleTaxType    int64  `json:"sale_tax_type"`
	BuyTaxType     int64  `json:"buy_tax_type"`
	SaleBillType   int64  `json:"sale_bill_type"`
	BuyBillType    int64  `json:"buy_bill_type"`
	UseAddress     string `json:"use_address"`
	PosDefCustId   int64  `json:"pos_def_cust_id"`
	PosDefStock    int64  `json:"pos_def_stock"`
	DefCustId      int64  `json:"def_cust_id"`
	CreateBy       string `json:"create_by"`
	CreateTime     string `json:"create_time"`
	EditBy         string `json:"edit_by"`
	EditTime       string `json:"edit_time"`
}
type SettingRepository struct{ db *sqlx.DB }

func NewSettingRepository(db *sqlx.DB) sales.Repository {
	return &SettingRepository{db}
}

func (repo *SettingRepository) SettingSys(req *sales.SearchHisCustomerTemplate) (resp interface{}, err error) {
	var sql string
	d := []NewSearchHisCustomerModel{}

	sql = `select distinct a.Id, ifnull(a.DocDate,'') as DocDate, ifnull(a.DocNo,'') as DocNo, 
		a.ArId, a.ArName, a.SaleName , a.TotalAmount
		from SaleOrder a 
		where a.ArCode like concat(?) 
		order by a.Id desc limit 20`
	err = repo.db.Select(&d, sql, req.ArCode)

	fmt.Println("sql = ", sql, req.ArCode)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	dp := []sales.NewSearchHisCustomerTemplate{}
	for _, dep := range d {
		dpline := map_hiscustomer_template(dep)
		dp = append(dp, dpline)
	}

	return dp, nil
}
