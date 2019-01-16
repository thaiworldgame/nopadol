package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
	"fmt"
)

type drivethruRepository struct{ db *sqlx.DB }

type BranchModel struct {
	Id int64 `json:"id" db:"id"`
	Code string `json:"code" db:"code"`
	Name string `json:"name" db:"branch_name"`
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

func (d *drivethruRepository)SearchListMachine()(interface{},error){
	rs,err := d.db.Query("select id,company_id,branch_id,machine_no,machine_code,def_wh_id,def_shelf_id" +
		",is_open from npdl.pos_machine")
	if err != nil {
		fmt.Println("error query database ")
		return nil,err
	}
	type machineModel struct {
		id          int64 `json:"id"`
		CompanyID   int64 `json:"company_id" `
		BranchID    int64 `json:"branch_id"`
		MachineNo   string `json:"machine_no"`
		MachineCode string `json:"machine_code"`
		DefWhID     int64 `json:"def_wh_id"`
		DefShelfID  int64 `json:"def_shelf_id"`
		IsOpen      int `json:"is_open"`
	}
	Mcs := []machineModel{}
	mc := machineModel{}
	for rs.Next() {
		rs.Scan(&mc.id,&mc.CompanyID,&mc.BranchID,&mc.MachineNo,&mc.MachineCode,&mc.DefWhID,&mc.DefShelfID,&mc.IsOpen)
		Mcs = append(Mcs,mc)
	}

	fmt.Println("mysqldb recive databranch -> ",Mcs)
	return Mcs,nil
}


func (d *drivethruRepository)SearchCarBrand(keyword string)(interface{},error){
	lccommand := "select id,"+
			 "company_id,"+
			 "branch_id,"+
			 "machine_no,"+
			 "machine_code,"+
			 "def_wh_id,"+
			 "def_shelf_id,"+
			 "is_open"+
		      " from npdl.pos_machine where machine_code like '%"+keyword+"%'"
	fmt.Println(lccommand)
	rs,err := d.db.Query(lccommand)
	if err != nil {
		fmt.Println("error query database ")
		return nil,err
	}
	type brandModel struct {
		id          int64 `json:"id"`
		CompanyID   int64 `json:"company_id" `
		BranchID    int64 `json:"branch_id"`
		MachineNo   string `json:"machine_no"`
		MachineCode string `json:"machine_code"`
		DefWhID     int64 `json:"def_wh_id"`
		DefShelfID  int64 `json:"def_shelf_id"`
		IsOpen      int `json:"is_open"`
	}
	Mcs := []brandModel{}
	mc := brandModel{}
	for rs.Next() {
		rs.Scan(&mc.id,&mc.CompanyID,&mc.BranchID,&mc.MachineNo,&mc.MachineCode,&mc.DefWhID,&mc.DefShelfID,&mc.IsOpen)
		Mcs = append(Mcs,mc)
	}

	fmt.Println("mysqldb recived brand data -> ",Mcs)
	return Mcs,nil
}

