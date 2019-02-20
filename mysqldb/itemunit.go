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
	fmt.Println("getByCode ",u.unitCode)
	//lccommand := `select id,unit_code,unit_name from item_unit where unit_code=?`
	lccommand := "select id,unit_code,unit_name from item_unit where unit_code='"+u.unitCode+"'"
	fmt.Println(lccommand)


	rs := db.QueryRow(lccommand)
	err = rs.Scan(&u.id,&u.unitCode,&u.unitName)
	if err != nil {
		fmt.Println("error rs.scan in getByCode -->",err.Error())
		return err
	}
	return nil
}

func (u *itemUnitModel)getByID(db *sqlx.DB)(err error){
	//todo : get itemUnitProfile by ID
	fmt.Println("start itemUnitModel.getByID ")
	lccommand := `select id,unit_code,unit_name from item_unit where id=?`
	err = db.QueryRow(lccommand,u.id).Scan(&u.id,&u.unitCode,&u.unitName)
	if err != nil {
		return err
	}
	return nil
}

