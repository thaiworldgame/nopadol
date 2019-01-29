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
