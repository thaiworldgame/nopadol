package mysqldb

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
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

type itemUnit struct {
	Id           int64          `db:"id"`
	UnitCode     string         `db:"unit_code"`
	UnitName     string         `db:"unit_name"`
	ActiveStatus int            `db:"active_status"`
	CreateBy     string         `db:"create_by"`
	CreateTime   mysql.NullTime `db:"create_time"`
}

type itemModel struct {
	Id           int64          `db:"id"`
	Code         string         `db:"code"`
	Name         string         `db:"name"`
	ShortName    string         `db:"short_name"`
	UnitID       int64          `db:"unit_id"`
	UnitCode     string         `db:"unit_code"`
	BuyUnitCode  string         `db:"buy_unit_code" `
	StockType    int            `db:"stock_type"`
	PicPath1     string         `db:"picture_path1"`
	PicPath2     string         `db:"picture_path2"`
	AverageCost  float64        `db:"average_cost"`
	ActiveStatus int            `db:"active_status"`
	ItemStatus   int            `db: "item_status"`
	CompanyID    int            `db:"company_id"`
	CreateBy     string         `db:"create_by"`
	CreateTime   mysql.NullTime `db:"create_time"`
	EditBy       string         `db:"edit_by"`
	EditTime     mysql.NullTime `db:"edit_time"`
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
	it.CompanyID = req.CompanyID
	fmt.Println("map2itemModel return ", it.UnitCode)
	return
}
func (it *itemModel) verifyRequestData(db *sqlx.DB) (bool ,error){
	if it.Code =="" {
		return false,fmt.Errorf("Item Code not found ")
	}

	if it.UnitCode ==""{
		return false , fmt.Errorf("Unit Code not Found..")
	}

	if it.Name ==""{
		return false, fmt.Errorf("Item Name not found...")
	}


	return true,nil
}

func(it *itemModel)checkExistsByCode(db *sqlx.DB,code string)(int64,bool){
	var id int64=-1
	db.QueryRow(`select id from Item where code=?`,code).Scan(&id)
	if id <= 0  {
		return -1,false
	}
	return id , true
}

func (it *itemModel) save(db *sqlx.DB) (newID int64, err error) {
	//check new data item
	fmt.Println("start item.save() req ",it)
	_,err = it.verifyRequestData(db)

	if err != nil {
		return -1, fmt.Errorf("verify state not pass error %v ", err.Error())
	}

	id,ok := it.checkExistsByCode(db,it.Code)
	if ok  {

		// update
		fmt.Println("update case to item.id -> ",id)
		db.Exec(`update Item set item_name=?,short_name=?,pic_path1 = ? , pic_path2=?
			where id = ?`,
			it.Name,it.ShortName,it.PicPath1,it.PicPath2,id)
		newID = id
	}else {

		fmt.Println("insert case ")
		// case new
		// todo : insert item flage incomplete
		lcCommand := `insert into Item (
			code,
			item_name,
			short_name,
			unit_code,
			buy_unit,
			stock_type,
			pic_path1,
			pic_path2,
			active_status,
			create_by,
			create_time,
			edit_by,
			edit_time,
			company_id) values (
			?,?,?,?,?,
			?,?,?,?,?,
			?,?,?,?
			)
	`
		rs, err := db.Exec(lcCommand,
			it.Code,
			it.Name,
			it.ShortName,
			it.UnitCode,
			it.BuyUnitCode,
			it.StockType,
			it.PicPath1,
			it.PicPath2,
			it.ActiveStatus,
			it.CreateBy,
			it.CreateTime,
			it.EditBy,
			it.EditTime,
			it.CompanyID,
		)
		if err != nil {
			log.Printf("error sql exec %v", err.Error())
			return -1, err
		}
		newID, err = rs.LastInsertId()
		if err != nil {
			log.Printf("error sql exec %v", err.Error())
			return -1, err
		}

	}


	// todo : insert barcode (default barcode = itemcode)
	// todo : insert price (option)
	// todo : insert ItemRate (default baseUnit rate=1)
	// todo : update complete save New

	return newID, nil
}

func (it *itemModel) getItemIDbyCode(db *sqlx.DB, code string) (int64, error) {
	lccommand := `select id from Item where code=?`
	db.QueryRow(lccommand, code).Scan(&it.Id)
	fmt.Printf("unit id from code %s is %v \n", code, it.Id)
	return it.Id, nil
}

func (it *itemModel) getItemCodeById(db *sqlx.DB) (string ,error) {
	db.QueryRow(`select code from Item where id=?`, it.Id).Scan(&it.Code)
	fmt.Printf("UnitCode from UnitID %s is %v \n", it.Id, it.Code)
	return it.Code,nil
}
