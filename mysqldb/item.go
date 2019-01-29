package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/product"
	"log"
)

const (
	// stsocktype constant
	itemStockTypeKeepStock int = 0
	itemStockTypeNoneStock int = 1

	// sale type Constant
	cashSaleType int = 0
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
	ItemCode   string  `db:"item_code"`
	UnitID     int64   `db:"unit_id"`
	UnitCode   string  `db:"unit_code"`
	SalePrice1 float64 `db:"sale_price_1"`
	SalePrice2 float64 `db:"sale_price_2"`
	SaleType   int     `db:"salt_type"`
	CompanyID  int     `db:"company_id"`
	// todo : add start , end date
	// todo : add volumn qty
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

func (it *itemModel) getItemIDbyCode(db *sqlx.DB, code string) (int64, error) {
	lccommand := `select id from Item where code=?`
	db.QueryRow(lccommand, code).Scan(&it.Id)
	fmt.Printf("unit id from code %s is %v \n", code, it.Id)
	return it.Id, nil

}

func (pr *priceModel) checkExitsByItemcodeUnitcodeSaletype(db *sqlx.DB) (id int64, result bool) {
	id  =-1
	rs := db.QueryRow(`select id
		from Price
		where item_id=? and unit_code=? and sale_type = ? limit 1 `,
	pr.ItemId, pr.UnitCode, pr.SaleType)
	rs.Scan(&id)

	//fmt.Printf("check price_id = %v  from item_id %v , unit_code %v, sale_type %v \n",id,pr.ItemId, pr.UnitCode, pr.SaleType)


	if id == -1 {
		return -1,false
	}
	//fmt.Printf(" price is exists with itemcode : %v,unitcode : %v ,saletype : %v",pr.ItemCode,getUnitCodeByUnitID(db, pr.UnitID),pr.SaleType)
	return  id,true
}
func (pr *priceModel) save(db *sqlx.DB) (newID int64, err error) {
	if pr.ItemId == 0 || pr.UnitID == 0 || pr.CompanyID == 0 {
		log.Printf(" error verify data is not ready: data -> %v", pr)
		return -1, fmt.Errorf("no item_id !!")
	}
	// todo : check exists item_code+unit_code+sale_type
	curID,ok := pr.checkExitsByItemcodeUnitcodeSaletype(db)
	fmt.Printf("check exists result is : ",ok)

	if  ok {
		// update and replace data with id
		fmt.Printf("case update  \n")
		_,err := db.Exec(`update Price set sale_price_1 = ? , sale_price_2 = ? where id = ?`,pr.SalePrice1,pr.SalePrice2,curID)
		if err != nil {
			log.Printf("update state sql command %v", err.Error())
			return -1, err
		}

	} else {
	//	insert new record
		fmt.Printf("case insert \n")

		lcCommand := `insert into Price (
		item_code,
		sale_type,
		unit_code,
		sale_price_1,
		sale_price_2,
		unit_id,
		company_id
		) VALUES(?,?,?,?,?,?,?)
	 ON DUPLICATE KEY UPDATE
	 sale_price_1=?,sale_price_2=?`

		rs, err := db.Exec(lcCommand,
			pr.ItemCode,
			pr.SaleType,
			pr.UnitCode,
			pr.SalePrice1,
			pr.SalePrice2,
			pr.UnitID,
			pr.CompanyID,
			pr.SalePrice1,
			pr.SalePrice2,
		)

		if err != nil {
			log.Printf("sql command %v", err.Error())
			return -1, err
		}
		newID, err = rs.LastInsertId()
		if err != nil {
			log.Printf("error insert new price ", err.Error())
			return -1, err
		}
	}

	return newID, nil
}

func (b *barcodeModel) save(db *sqlx.DB) (newID int64, err error) {
	return -1, nil
}

func getUnitIDByUnitCode(db *sqlx.DB, unitCode string) int64 {
	var unitID int64
	lccommand := `select id from item_unit where unit_code=?`
	db.QueryRow(lccommand, unitCode).Scan(&unitID)
	fmt.Printf("unit id from code %s is %v \n", unitCode, &unitID)
	return unitID
}
func getUnitCodeByUnitID(db *sqlx.DB, unitID int64) (uCode string) {
	db.QueryRow(`select unit_code from item_unit where id=?`, unitID).Scan(&uCode)
	fmt.Printf("UnitCode from UnitID %s is %v \n", unitID, uCode)
	return
}

func (it *itemModel)getItemCodeById(db *sqlx.DB)error{
	db.QueryRow(`select code from Item where id=?`, it.Id).Scan(&it.Code)
	fmt.Printf("UnitCode from UnitID %s is %v \n", it.Id,it.Code)
	return nil
}

func (r *packingRate) save(db *sqlx.DB) (newID int64, err error) {
	// check data before ins
	if r.ItemID == 0 {
		it := itemModel{}
		r.ItemID, err = it.getItemIDbyCode(db, r.ItemCode)
		if err != nil {
			return -1, fmt.Errorf("no item_id! ")
		}
	}

	if r.UnitID == 0 {
		r.UnitID = getUnitIDByUnitCode(db, r.UnitCode)
	}

	lccommand := `insert into ItemRate (
		item_id,
		item_code,
		unit_code,
		rate1
		) VALUES(?,?,?,?)
	 ON DUPLICATE KEY UPDATE
	 rate1=?,unit_id=?`

	rs, err := db.Exec(lccommand, r.ItemID, r.ItemCode, r.UnitCode, r.RatePerBaseUnit, r.RatePerBaseUnit, r.UnitID)
	if err != nil {
		log.Printf("sql command %v", err.Error())
		return -1, err
	}
	newID, err = rs.LastInsertId()
	if err != nil {
		log.Printf("error insert new rate to item ", err.Error())
		return -1, err
	}
	return newID, nil
}
