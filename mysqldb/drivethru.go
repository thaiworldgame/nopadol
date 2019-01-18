package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
	"github.com/satori/go.uuid"
	"time"
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

func (d *drivethruRepository) SearchItem(keyword string) (interface{}, error) {
	fmt.Println("mysql recived param keyword -> ", keyword)
	type itemModel struct {
		ItemBarcode   string  `json:"item_barcode"`
		ItemCode      string  `json:"item_code"`
		ItemName      string  `json:"item_name"`
		ItemCategory  string  `json:"item_category"`
		ItemPrice     float64 `json:"item_price"`
		ItemPrice2    float64 `json:"item_price_2"`
		ItemUnitCode  string  `json:"item_unit_code"`
		ItemRemark    string  `json:"item_remark"`
		ItemShortName string  `json:"item_short_name"`
		ItemFilePath  string  `json:"item_file_path"`
	}
	lccommand := "select 	a.bar_code ," +
		"a.item_code," +
		"a.unit_code ," +
		"c.sale_price_1 as price ," +
		"c.sale_price_2 as price2," +
		"b.item_name ," +
		"b.short_name ," +
		"'-' as item_category," +
		"'-' as item_remark," +
		"b.pic_path1 as item_file_path " +
		"from npdl.Barcode a left outer join npdl.Item b on a.item_code=b.code " +
		"left join npdl.Price c on b.code = c.item_code and a.unit_code = c.unit_code " +
		" where a.bar_code='" + keyword + "' limit 1 "
	fmt.Println(lccommand)
	rs := d.db.QueryRow(lccommand)
	it := itemModel{}
	rs.Scan(&it.ItemBarcode, &it.ItemCode, &it.ItemUnitCode, &it.ItemPrice, &it.ItemPrice2,
		&it.ItemName, &it.ItemShortName, &it.ItemCategory, &it.ItemRemark, &it.ItemFilePath)

	fmt.Println("before mysql return -> ", it)
	return it, nil
}

func (d *drivethruRepository) UserLogIn(req *drivethru.UserLogInRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	//DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	now.String()

	fmt.Println("mysql recived param req -> ", req.UserCode, req.BranchId)
	type userLogInModel struct {
		Id               string `json:"id"`
		Code             string `json:"code"`
		Name             string `json:"name"`
		Password         string `json:"password"`
		ImageFileName    string `json:"image_filename"`
		Role             int    `json:"role"`
		ActiveStatus     int    `json:"activeStatus"`
		IsConfirm        int    `json:"isConfirm"`
		CreatorCode      string `json:"creatorCode"`
		CreateDateTime   string `json:"createdateTime"`
		LastEditorCode   string `json:"lasteditorCode"`
		LastEditDateTime string `json:"lasteditdateTime"`
		BranchCode       string `json:"branchCode"`
		Remark           string `json:"remark"`
		LoginZone        string `json:"loginZone"`
		CompanyId        int    `json:"company_id"`
		BranchId         int    `json:"branch_id"`
	}

	branch_code := getBranch(d.db, req.BranchId)
	lccommand := "select id,code,name,ifnull(pwd,'') as password,ifnull(image_path,'') as image_filename,role,active_status as activesTatus,is_confirm as isConfirm,ifnull(create_by,'') as creatorCode,ifnull(create_time,'') as createdateTime,ifnull(edit_by,'') as lasteditorCode,ifnull(edit_time,'') as lasteditdateTime,ifnull(branch_code,'') as branchCode,'' as remark,ifnull(zone_id,'') as loginZone,ifnull(company_id,1) as company_id,ifnull(branch_id,1) as branch_id from user where code = ? and branch_code = ? and pwd = ?"
	fmt.Println(lccommand, req.UserCode, branch_code)
	rs := d.db.QueryRow(lccommand, req.UserCode, branch_code, &req.Password)
	user := userLogInModel{}
	err := rs.Scan(&user.Id, &user.Code, &user.Name, &user.Password, &user.ImageFileName, &user.Role, &user.ActiveStatus, &user.IsConfirm, &user.CreatorCode, &user.CreateDateTime, &user.LastEditorCode, &user.LastEditDateTime, &user.BranchCode, &user.Remark, &user.LoginZone, &user.CompanyId, &user.BranchId)
	if err != nil || user.Code == "" {
		fmt.Println("error = ", err.Error())
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "login",
				"processDesc": "login fail",
				"isSuccess":   false,
			},
		}, nil
	}
	fmt.Println("before mysql return -> ", user)

	var check_exist int
	var uuid string

	lccommand_check := `select count(*) as vCount from user_access where user_id = ? and user_code = ? `
	err = d.db.Get(&check_exist, lccommand_check, user.Id, req.UserCode)
	if err != nil {
		fmt.Println(err.Error())
	}

	if check_exist == 0 {
		uuid = getUUID()
	}else{
		lccommand_check := `select access_token from user_access where user_id = ? and user_code = ? `
		err = d.db.Get(&uuid, lccommand_check, user.Id, req.UserCode)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(check_exist)
	var expire_date string

	due_date := now.AddDate(0, 0, int(req.CreditDay)).Format("2006-01-02")

	DueDate = due_date

	fmt.Println("duedate =", req.DueDate)

	lccommand = "insert user_access(user_id,user_code,access_token,company_id,branch_id,branch_code,zone_id,create_time,expire_time) values(?,?,?,?,?,?,?,?,ADDDATE(CONVERT_TZ(CURRENT_TIMESTAMP(),'+00:00','+07:00'),INTERVAL 1 DAY)"
	u, err := d.db.Exec(lccommand, user.Id, user.Code, uuid, user.CompanyId, user.BranchId, branch_code, user.LoginZone, now.String())
	if err != nil {
		fmt.Println("error = ", err.Error())
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "login",
				"processDesc": err.Error(),
				"isSuccess":   false,
			},
		}, nil
	}

	fmt.Println(u.LastInsertId())

	return map[string]interface{}{
		"response": map[string]interface{}{
			"process":     "login",
			"processDesc": "successful",
			"isSuccess":   true,
		},
		"accessToken":    uuid,
		"accessDatetime": now.String(),
		"pathPHPUpload":  "http://qserver.nopadol.com/drivethru/upload.php",
		"pathFile":       "http://qserver.nopadol.com/drivethru/tmp/",
		"user":           user,
	}, nil

	//return user, nil
}

func getBranch(db *sqlx.DB, branch_id int) string {
	var branch_code string

	lccommand := `select code from branch where id = ?`
	err := db.Get(&branch_code, lccommand, branch_id)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Branch Code =", branch_code)
	return branch_code
}

func getUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err.Error()
	}
	fmt.Printf("UUIDv4: %s\n", uuid)

	return uuid.String()
}
