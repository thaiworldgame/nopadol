package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type packingRate struct {
	Id              int64  `db:"id"`
	ItemID          int64  `db:"item_id"`
	ItemCode        string `db:"item_code"`
	UnitID          int64  `db:"unit_id"`
	UnitCode        string `db:"unit_code"`
	RatePerBaseUnit int    `db:"rate1"`
	CompanyID       int    `db:"company_id"`
}

<<<<<<< Updated upstream
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
		u := itemUnitModel{}
		u.getByCode(db)
		r.UnitID = u.id
=======
func (r *packingRate) verifyRequestData(db *sqlx.DB) (err error) {
	it := itemModel{}
	u := itemUnitModel{}

	switch {
	case r.ItemID == 0 && r.ItemCode != "":
		r.ItemID, err = it.getItemIDbyCode(db, r.ItemCode)
		if err != nil {
			return fmt.Errorf("no item_id! ")
		}
	case r.ItemID != 0 && r.ItemCode == "":
		r.ItemCode, err = it.getItemCodeById(db)
		if err != nil {
			return fmt.Errorf("no item_id! ")
		}
	case r.UnitID == 0 && r.UnitCode != "": // have Unitcode
		err = u.getByCode(db)
		if err != nil {
			return fmt.Errorf("no unit_id ", r.UnitID)
		}
		r.UnitID = u.id
	case r.UnitCode == "" && r.UnitID != 0: // have unitID
		err = u.getByID(db)
		if err != nil {
			return fmt.Errorf("unit_code not found :", r.UnitCode)
		}
		r.UnitCode = u.unitCode
	case r.RatePerBaseUnit == 0:
		return fmt.Errorf("rate zero error ")
	case r.CompanyID == 0 :
		return fmt.Errorf("company id zero error")
	}
	return nil
}
func (r *packingRate) save(db *sqlx.DB) (newID int64, err error) {
	// check data before ins

	err = r.verifyRequestData(db)
	if err != nil {
		return -1, err
>>>>>>> Stashed changes
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
