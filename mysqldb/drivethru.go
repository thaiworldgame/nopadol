package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
	"fmt"
)

type drivethruRepository struct{ db *sqlx.DB }

type BranchModel struct {
	Id int64 `db:"id"`
	Code string `db:"code"`
	Name string `db:"branch_name"`
}

func NewDrivethruRepository(db *sqlx.DB) drivethru.Repository {
	return &drivethruRepository{db}
}


func (d *drivethruRepository) SearchListCompany() (interface{}, error) {
	rs,err := d.db.Query("select id,code,branch_name from npdl.branch ")
	if err != nil {
		fmt.Println("error query database ")
		return nil,err
	}

	Bms := []BranchModel{}
	bm := BranchModel{}
	for rs.Next() {
		rs.Scan(&bm.Id,&bm.Code,&bm.Name)
		Bms = append(Bms,bm)
	}

	fmt.Println("mysqldb recive databranch -> ",Bms)
	return Bms,nil
}