package mysqldb

import (
	"log"
	"github.com/jmoiron/sqlx"
	"fmt"
)
type packingRate struct {
	Id              int64  `db:"id"`
	ItemID          int64  `db:"item_id"`
	ItemCode        string `db:"item_code"`
	UnitID          int64  `db:"unit_id"`
	UnitCode        string `db:"unit_code"`
	RatePerBaseUnit int    `db:"rate1"`
}

func(r *packingRate)verifyRequestData(db *sqlx.DB)(pass bool,err error){
	if r.ItemID == 0 {
		it := itemModel{}
		r.ItemID, err = it.getItemIDbyCode(db, r.ItemCode)
		if err != nil {
			return false, fmt.Errorf("no item_id! ")
		}
	}

	// find unitid if 0 and have code
	if r.UnitID == 0 && r.UnitCode !="" {
		u := itemUnitModel{}
		err = u.getByCode(db)
		if err != nil {
			return false,fmt.Errorf(" no unit_id ",r.UnitID)
		}
		r.UnitID = u.id

	}

	// find unitcode if blank and have unitid
	if r.UnitCode =="" && r.UnitID!=0{
		u := itemUnitModel{}
		err = u.getByID(db)
		if err != nil {
			return false,fmt.Errorf("unit_code not found :",r.UnitCode)
		}
		r.UnitCode = u.unitCode

	}
	return true,nil

}
func (r *packingRate) save(db *sqlx.DB) (newID int64, err error) {
	// check data before ins

	ok,err := r.verifyRequestData(db)
	if !ok || err !=nil {
		return -1,err
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
