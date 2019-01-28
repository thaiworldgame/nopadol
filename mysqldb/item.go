package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/product"
	"fmt"
	"log"
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

func (it *itemModel) getItemIDbyCode(db *sqlx.DB,code string)(int64,error){
	lccommand := `select id from Item where code=?`
	db.QueryRow(lccommand,code).Scan(&it.Id)
	fmt.Printf("unit id from code %s is %v \n",code,it.Id)
	return it.Id,nil

}

func (pr *priceModel) save(db *sqlx.DB) (newID int64, err error) {
	return -1, nil
}

func (b *barcodeModel) save(db *sqlx.DB) (newID int64, err error) {
	return -1, nil
}

func (r *packingRate)getUnitIDByUnitCode(db *sqlx.DB,unitcode string)(int64,error){
	lccommand := `select id from ItemRate where unit_code=?`
	db.QueryRow(lccommand,unitcode).Scan(&r.UnitID)
	fmt.Printf("unit id from code %s is %v \n",unitcode,r.UnitID)
	return r.UnitID,nil
}
func (r *packingRate) save(db *sqlx.DB) (newID int64, err error) {
	// check data before ins
	if r.ItemID == 0 {
		it := itemModel{}
		r.ItemID,err = it.getItemIDbyCode(db,r.ItemCode)
		if err != nil {
			return -1 , fmt.Errorf("no item_id! ")
		}
	}

	if r.UnitID == 0 {
		r.UnitID,err = r.getUnitIDByUnitCode(db,r.UnitCode)
		if err != nil {
			return -1 , fmt.Errorf("no unit_id! ")
		}
	}

	lccommand := `insert into ItemRate (
		item_id,
		item_code,
		unit_code,
		rate1
		) VALUES(?,?,?,?)
	 ON DUPLICATE KEY UPDATE
	 rate1=?,unit_id=?`

	rs,err := db.Exec(lccommand,r.ItemID,r.ItemCode,r.UnitCode,r.RatePerBaseUnit,r.RatePerBaseUnit,r.UnitID)
	if err != nil {
		log.Printf("sql command %v",err.Error())
		return -1,err
	}
	newID , err = rs.LastInsertId()
	if err != nil {
		log.Printf("error insert new rate to item ",err.Error())
		return -1,err
	}
	return newID, nil
}

