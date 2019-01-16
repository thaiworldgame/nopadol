package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
)

type drivethruRepository struct{ db *sqlx.DB }

type BranchModel struct {
	Id   int64  `json:"id" db:"id"`
	Code string `json:"code" db:"code"`
	Name string `json:"name" db:"branch_name"`
}

func NewDrivethruRepository(db *sqlx.DB) drivethru.Repository {
	return &drivethruRepository{db}
}

func (d *drivethruRepository) SearchListCompany() (interface{}, error) {
	rs, err := d.db.Query("select id,code,branch_name from npdl.branch ")
	if err != nil {
		fmt.Println("error query database ")
		return nil, err
	}

	Bms := []BranchModel{}
	bm := BranchModel{}
	for rs.Next() {
		rs.Scan(&bm.Id, &bm.Code, &bm.Name)
		Bms = append(Bms, bm)
	}

	fmt.Println("mysqldb recive databranch -> ", Bms)
	return Bms, nil
}

func (d *drivethruRepository) SearchListMachine() (interface{}, error) {
	rs, err := d.db.Query("select id,company_id,branch_id,machine_no,machine_code,def_wh_id,def_shelf_id" +
		",is_open from npdl.pos_machine")
	if err != nil {
		fmt.Println("error query database ")
		return nil, err
	}
	type machineModel struct {
		id          int64  `json:"id"`
		CompanyID   int64  `json:"company_id" `
		BranchID    int64  `json:"branch_id"`
		MachineNo   string `json:"machine_no"`
		MachineCode string `json:"machine_code"`
		DefWhID     int64  `json:"def_wh_id"`
		DefShelfID  int64  `json:"def_shelf_id"`
		IsOpen      int    `json:"is_open"`
	}
	Mcs := []machineModel{}
	mc := machineModel{}
	for rs.Next() {
		rs.Scan(&mc.id, &mc.CompanyID, &mc.BranchID, &mc.MachineNo, &mc.MachineCode, &mc.DefWhID, &mc.DefShelfID, &mc.IsOpen)
		Mcs = append(Mcs, mc)
	}

	fmt.Println("mysqldb recive databranch -> ", Mcs)
	return Mcs, nil
}

func (d *drivethruRepository) SearchCarBrand(keyword string) (interface{}, error) {
	lccommand := "select id," +
		"car_brand," +
		"active_status " +
		" from npdl.car_brand where car_brand like '%" + keyword + "%'"
	fmt.Println(lccommand)
	rs, err := d.db.Query(lccommand)
	if err != nil {
		fmt.Println("error query database ")
		return nil, err
	}
	type brandModel struct {
		Id           int64  `json:"id"`
		CarBrand     string `json:"name"`
		ActiveStatus int    `json:"active_status"`
	}
	Mcs := []brandModel{}
	mc := brandModel{}
	for rs.Next() {
		rs.Scan(&mc.Id, &mc.CarBrand, &mc.ActiveStatus)
		Mcs = append(Mcs, mc)
	}

	fmt.Println("mysqldb recived brand data -> ", Mcs)
	return Mcs, nil
}

func (d *drivethruRepository) SearchCustomer(keyword string) (interface{}, error) {
	// todo: search by keyword here  by code,name,telno
	lccommand := "select id," +
		"code," +
		"name, " +
		"address, " +
		"telephone" +
		" from npdl.Customer where code like '%" + keyword + "%' or name like '%" + keyword + "%'"
	fmt.Println(lccommand)
	rs, err := d.db.Query(lccommand)
	if err != nil {
		fmt.Println("error query database ")
		return nil, err
	}
	type customerModel struct {
		Id        int64  `json:"id"`
		Code      string `json:"code"`
		Name      string `json:"name"`
		Address   string `json:"address"`
		Telephone string `json:"telephone"`
	}
	Mcs := []customerModel{}
	mc := customerModel{}
	for rs.Next() {
		rs.Scan(&mc.Id, &mc.Code, &mc.Name, &mc.Address, &mc.Telephone)
		Mcs = append(Mcs, mc)
	}

	fmt.Println("mysqldb recived brand data -> ", Mcs)
	return Mcs, nil
}
