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

func (d *drivethruRepository)SearchItem(keyword string)(interface{},error){
	fmt.Println("mysql recived param keyword -> ",keyword)
	type itemModel struct {
		ItemBarcode string `json:"item_barcode"`
		ItemCode string `json:"item_code"`
		ItemName string `json:"item_name"`
		ItemCategory string `json:"item_category"`
		ItemPrice float64 `json:"item_price"`
		ItemPrice2 float64 `json:"item_price_2"`
		ItemUnitCode string `json:"item_unit_code"`
		ItemRemark string `json:"item_remark"`
		ItemShortName string `json:"item_short_name"`
		ItemFilePath string `json:"item_file_path"`
	}
	lccommand := "select 	a.bar_code ,"+
		"a.item_code,"+
		"a.unit_code ,"+
		"c.sale_price_1 as price ,"+
		"c.sale_price_2 as price2,"+
		"b.item_name ,"+
		"b.short_name ,"+
		"'-' as item_category,"+
		"'-' as item_remark,"+
		"b.pic_path1 as item_file_path "+
	"from npdl.Barcode a left outer join npdl.Item b on a.item_code=b.code "+
	"left join npdl.Price c on b.code = c.item_code and a.unit_code = c.unit_code "+
	" where a.bar_code='"+keyword+"' limit 1 "
	fmt.Println(lccommand)
	rs := d.db.QueryRow(lccommand)
	it :=itemModel{}
	rs.Scan(&it.ItemBarcode,&it.ItemCode,&it.ItemUnitCode,&it.ItemPrice,&it.ItemPrice2,
	&it.ItemName,&it.ItemShortName,&it.ItemCategory,&it.ItemRemark,&it.ItemFilePath)

	fmt.Println("before mysql return -> ",it)
	return it,nil
}


func (d *drivethruRepository)UserLogIn(req *drivethru.UserLogInRequest)(interface{},error){
	fmt.Println("mysql recived param req -> ",req.UserCode)
	type userLogInModel struct {
		ItemBarcode string `json:"item_barcode"`
		ItemCode string `json:"item_code"`
		ItemName string `json:"item_name"`
		ItemCategory string `json:"item_category"`
		ItemPrice float64 `json:"item_price"`
		ItemPrice2 float64 `json:"item_price_2"`
		ItemUnitCode string `json:"item_unit_code"`
		ItemRemark string `json:"item_remark"`
		ItemShortName string `json:"item_short_name"`
		ItemFilePath string `json:"item_file_path"`
	}
	lccommand := "select 	a.bar_code ,"+
		"a.item_code,"+
		"a.unit_code ,"+
		"c.sale_price_1 as price ,"+
		"c.sale_price_2 as price2,"+
		"b.item_name ,"+
		"b.short_name ,"+
		"'-' as item_category,"+
		"'-' as item_remark,"+
		"b.pic_path1 as item_file_path "+
		"from npdl.Barcode a left outer join npdl.Item b on a.item_code=b.code "+
		"left join npdl.Price c on b.code = c.item_code and a.unit_code = c.unit_code "+
		" where a.bar_code=? limit 1 "
	fmt.Println(lccommand)
	//rs := d.db.QueryRow(lccommand,req.UserCode)
	user :=userLogInModel{}
	//rs.Scan(&it.ItemBarcode,&it.ItemCode,&it.ItemUnitCode,&it.ItemPrice,&it.ItemPrice2,
	//	&it.ItemName,&it.ItemShortName,&it.ItemCategory,&it.ItemRemark,&it.ItemFilePath)

	fmt.Println("before mysql return -> ",user)
	return user,nil
}