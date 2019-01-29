package mysqldb

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/product"
)

const (
	// stsocktype constant
	itemStockTypeKeepStock int = 0
	itemStockTypeNoneStock int = 1

	// sale type Constant
	cashSaleType int = 0
)

type itemUnit struct {
	Id           int64          `db:"id"`
	UnitCode     string         `db:"unit_code"`
	UnitName     string         `db:"unit_name"`
	ActiveStatus int            `db:"active_status"`
	CreateBy     string         `db:"create_by"`
	CreateTime   mysql.NullTime `db:"create_time"`
}

type itemModel struct {
	Id         int64          `db:"id"`
	Code       string         `db:"code"`
	Name       string         `db:"name"`
	ShortName  string         `db:"short_name"`
	UnitCode   string         `db:"unit_code"`
	StockType  int            `db:"stock_type"`
	PicPath1   string         `db:"picture_path1"`
	PicPath2   string         `db:"picture_path2"`
	ItemStatus int            `db: "item_status"`
	CompanyID  int            `db:"company_id"`
	CreateBy   string         `db:"create_by"`
	CreateTime mysql.NullTime `db:"create_time"`
	EditBy     string         `db:"edit_by"`
	EditTime   mysql.NullTime `db:"edit_time"`
}

func (it *itemModel) map2itemModel(db *sqlx.DB, req *product.ProductNewRequest) (err error) {
	u := itemUnitModel{id: req.UnitID}
	fmt.Println("map2itemModel  unitid -->", req.UnitID)
	err = u.getByID(db)
	it.Code = req.ItemCode
	it.Name = req.ItemName
	it.UnitCode = u.unitCode
	it.PicPath1 = req.Picture
	it.StockType = req.StockType
	fmt.Println("map2itemModel return ", it.UnitCode)
	return
}

func (it *itemModel) save(db *sqlx.DB) (newID int64, err error) {

	return -1, nil
}

func (it *itemModel) getItemIDbyCode(db *sqlx.DB, code string) (int64, error) {
	lccommand := `select id from Item where code=?`
	db.QueryRow(lccommand, code).Scan(&it.Id)
	fmt.Printf("unit id from code %s is %v \n", code, it.Id)
	return it.Id, nil
}

func (it *itemModel) getItemCodeById(db *sqlx.DB) error {
	db.QueryRow(`select code from Item where id=?`, it.Id).Scan(&it.Code)
	fmt.Printf("UnitCode from UnitID %s is %v \n", it.Id, it.Code)
	return nil
}
