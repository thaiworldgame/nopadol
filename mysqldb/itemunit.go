package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type itemUnitModel struct {
	id int64 `db:"id"`
	unitCode string `db:"unit_code"`
	unitName string `db:"unit_name"`

}

func (u *itemUnitModel)getByCode(db *sqlx.DB)(err error){
	//todo : get itemUnitProfile by code
	fmt.Printf("get UnitID By Code %v \n",u.unitCode)
	lccommand := `select id from item_unit where unit_code=?`
	err = db.QueryRow(lccommand,u.unitCode).Scan(&u.id)
	if err != nil {
		fmt.Printf("error get unit_id by code ... %v",err.Error())
		return err
	}
	return nil
}

func (u *itemUnitModel)getByID(db *sqlx.DB)(err error){
	//todo : get itemUnitProfile by ID
	lccommand := `select unit_code from item_unit where id=?`
	err = db.QueryRow(lccommand,u.id).Scan(&u.unitCode)
	if err != nil {
		fmt.Printf("error get unit_code from unit_id ... %v ",err.Error())
		return err
	}
	return nil
}

