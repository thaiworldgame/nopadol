package mysqldb

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	barCodeStatusActive   int = 1
	barCodeStatusInActive int = 0
)

type itemBarcodeModel struct {
	Id      int64  `db:"id"`
	ItemID  int64  `db:"item_id"`
	Barcode string `db:"barcode"`
	UnitID  int64  `db:"unit_id"`
}

type barcodeModel struct {
	Id           int64          `db:"id"`
	ItemID       int64          `db:"item_id"`
	ItemCode     string         `db:"item_code"`
	BarCode      string         `db:"bar_code"`
	UnitID       int64          `db:"unit_id"`
	UnitCode     string         `db:"unit_code"`
	ActiveStatus int            `db:"active_status"`
	CreateBy     string         `db:"create_by"`
	CreateTime   mysql.NullTime `db:"create_time"`
	EditBy       string         `db:"edit_by"`
	EditTime     mysql.NullTime `db:"edit_time"`
	CompanyID    int            `db:"company_id"`
}

func (b *barcodeModel) checkExistByBarcode(db *sqlx.DB) (bool, int64) {
	var id int64 = -1
	db.QueryRow(`select id from Barcode where bar_code=?`, b.BarCode).Scan(&id)
	if id == -1 {
		return false, id
		// ไม่มี barcode อยู่
	}
	return true, id
}
func (b *barcodeModel) verifyRequestData(db *sqlx.DB) error {
	// check item_id
	if b.ItemID == 0 {
		return fmt.Errorf("item_id is error %v ", b.ItemID)
	}

	// check itemcode and get
	if b.ItemCode == "" {
		it := itemModel{}
		it.Id = b.ItemID
		it.getItemCodeById(db)
		if it.Code ==""{
			return fmt.Errorf(" data missing somting.. itemcode is blank")
		}else{
			b.ItemCode=it.Code
		}

	}

	// check unitcode and get
	if b.UnitCode == "" {
		u := itemUnitModel{}
		u.id = b.UnitID
		u.getByID(db)

		if u.unitCode ==""{
			return fmt.Errorf("Unitcode is blank")
		} else {
			b.UnitCode  = u.unitCode
		}
	}

	if b.BarCode == "" {
		return fmt.Errorf("Barcode  is blank")
	}

	return nil
}


func (b *barcodeModel) save(db *sqlx.DB) (id int64, err error) {

	err = b.verifyRequestData(db)
	if err != nil {
		return -1, err
	}

	ok, curID := b.checkExistByBarcode(db)
	if ok {
		// update
		fmt.Println("update case")
		lcCommand := `update Barcode set
			bar_code =? ,
			item_id=?,
			unit_id = ?,
			active_status=?,
			edit_by = ?,
			edit_time=?,
			company_id=?
			where id = ?`
		_, err := db.Exec(lcCommand,
			b.BarCode, b.ItemID, b.UnitID, b.ActiveStatus, b.EditBy, b.EditTime, b.CompanyID, curID)
		if err != nil {
			log.Printf("sql exec err ", err.Error())
			return -1, err

		}
		id = curID
	} else {
		// new barcode
		//
		fmt.Println("insert case ")
		lcCommand := `insert into Barcode (
		bar_code,
		unit_id,
		unit_code,
		active_status,
		item_id,
		item_code,
		create_by,
		create_time)
	values (?,?,?,?,?,?,?,?)`
		rs, err := db.Exec(lcCommand,
			b.BarCode,
			b.UnitID,
			b.UnitCode,
			b.ActiveStatus,
			b.ItemID,
			b.ItemCode,
			b.CreateBy,
			b.CreateTime.Time)
		if err != nil {
			log.Printf("sql exec err ", err.Error())
			return -1, err
		}

		id, err = rs.LastInsertId()
		if err != nil {
			log.Printf("sql exec err ", err.Error())
			return -1, err
		}

	}
	return id, nil
}
