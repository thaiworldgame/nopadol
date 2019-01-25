package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/product"
	"fmt"
)

type itemModel struct {
	Id         int64  `db:"id"`
	Code       string `db:"code"`
	Name       string `db:"name"`
	ShortName  string `db:"short_name"`
	UnitCode   string `db:"unit_code"`
	StockType  int    `db:"stock_type"`
	PicPath1   string `db:"picture_path1"`
	PicPath2   string `db:"picture_path2"`
	ItemStatus int    `db: "item_status"`
}

type priceModel struct {
	Id         int64   `db:"id"`
	ItemId     int64   `db:"item_id"`
	UnitID     int64   `db:"unit_id"`
	UnitCode   string  `db:"unit_code"`
	SalePrice1 float64 `db:"sale_price_1"`
	SalePrice2 float64 `db:"sale_price_2"`
	// todo : add start , end date
	// todo : add volumn qty
}


func (it *itemModel)map2itemModel(db *sqlx.DB,req *product.ProductNewRequest)(err error){
	u := itemUnitModel{id : req.UnitID}
	fmt.Println("map2itemModel  unitid -->",req.UnitID)
	err = u.getByID(db)
		it.Code = req.ItemCode
		it.Name =  req.ItemName
		it.UnitCode = u.unitCode
		it.PicPath1 =  req.Picture
		it.StockType =  req.StockType
	fmt.Println("map2itemModel return ",it.UnitCode)
	return
}

type barcodeModel struct {
	Id           int64  `db:"id"`
	ItemID       int64  `db:"item_id"`
	ItemCode     string `db:"item_code"`
	BarCode      string `db:"bar_code"`
	UnitID       int64  `db:"unit_id"`
	UnitCode     string `db:"unit_code"`
	ActiveStatus int    `db:"active_status"`
}

type packingRate struct {
	Id              int64  `db:"id"`
	ItemID          int64  `db:"item_id"`
	ItemCode        string `db:"item_code"`
	UnitID          int64  `db:"unit_id"`
	UnitCode        string `db:"unit_code"`
	RatePerBaseUnit int    `db:"rate1"`
}

func (it *itemModel) save(db *sqlx.DB) (newID int64, err error) {
	return -1, nil
}

func (pr *priceModel) save(db *sqlx.DB) (newID int64, err error) {
	return -1, nil
}

func (b *barcodeModel) save(db *sqlx.DB) (newID int64, err error) {
	return -1, nil
}

func (r *packingRate) save(db *sqlx.DB) (newID int64, err error) {

	lccommand := `insert into ItemRate (item_id,item_code,unit_code,rate1)
			values (?,?,?,?)`
	rs,err := db.Exec(lccommand,r.ItemID,r.ItemCode,r.UnitCode,r.RatePerBaseUnit)
	newID , err = rs.LastInsertId()
	if err != nil {
		return -1,err
	}
	return newID, nil
}

