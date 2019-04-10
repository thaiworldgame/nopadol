package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	machineOpen      int = 1
	shiftStatusOpen  int = 0
	shiftStatusClose int = 1
)

type ShiftModel struct {
	Id                 int64   `json:"id" db:"id"`
	DocDate            string  `json:"doc_date" db:"doc_date"`
	CompanyID          int     `json:"company_id" db:"company_id"`
	BranchID           int     `json:"branch_id" db:"branch_id"`
	MachineID          int     `json:"machine_id" db:"machine_id"`
	MachineNo          string  `json:"machine_no" db:"machine_no"`
	CashierID          int     `json:"cashier_id" db:"cashier_id"`
	CashierName        string  `json:"cashier_name" db:"cashier_name"`
	ChangeAmount       float64 `json:"change_amount" db:"change_amount"`
	WhID               int     `json:"wh_id" db:"wh_id"`
	Status             int     `json:"status" db:"status"`
	ShiftUUid          string  `json:"shift_uuid" db:"shift_uid"`
	Remark             string  `json:"remark" db:"remark"`
	OpenBy             string  `json:"open_by" db:"open_by"`
	OpenTime           string  `json:"open_time" db:"open_time"`
	EditBy             string  `json:"edit_by" db:"edit_by"`
	EditTime           string  `json:"edit_time" db:"edit_time"`
	CloseTime          string  `json:"close_time" db:"close_time"`
	CloseBy            string  `json:"close_by" db:"close_by"`
	SumOfCashAmount    float64 `json:"sum_of_cash_amount" db:"sum_of_cash_amount"`
	SumOfCreditAmount  float64 `json:"sum_of_credit_amount" db:"sum_of_credit_amount"`
	SumOfBankAmount    float64 `json:"sum_of_bank_amount" db:"sum_of_bank_amount"`
	SumOfCouponAmount  float64 `json:"sum_of_coupon_amount" db:"sum_of_coupon_amount"`
	SumOfDepositAmount float64 `json:"sum_of_deposit_amount" db:"sum_of_deposit_amount"`
}

type Machine struct {
	MachineId   int    `json:"machine_id" db:"machine_id"`
	MachineNo   string `json:"machine_no" db:"machine_no"`
	MachineCode string `json:"machine_code" db:"machine_code"`
	DefWhId     int    `json:"def_wh_id" db:"def_wh_id"`
	DefShelfId  int    `json:"def_shelf_id" db:"def_shelf_id"`
	WHCode      string `json:"wh_code" db:"wh_code"`
	ShelfCode   string `json:"shelf_code" db:"shelf_code"`
	CashierID   int    `json:"cashier_id" db:"cashier_id"`
	ShiftUUID   string `json:"shift_uuid" db:"shift_uuid"`
}

func (m *Machine) SearchMachineId(db *sqlx.DB, company_id int, branch_id int, machine_id int) {
	fmt.Println("Search Machine UUID = ", machine_id, company_id, branch_id)
	lccommand := " select a.id as machine_id,ifnull(machine_no,'') as machine_no,ifnull(CashierId) as cashier_id, ifnull(a.def_wh_id,0) as def_wh_id,ifnull(def_shelf_id,0) as def_shelf_id,ifnull(b.wh_code,'') as wh_code, ifnull(c.shelf_code,'') as shelf_code from  pos_machine as a inner join warehouse b on a.def_wh_id = b.id and a.company_id = b.company_id and a.branch_id = b.branch_id inner join warehouse_shelf c on a.def_shelf_id = c.id and a.company_id = c.company_id and a.branch_id = c.branch_id where a.id   = ? and a.company_id = ? and a.branch_id = ?"
	rs := db.QueryRow(lccommand, machine_id, company_id, branch_id)
	err := rs.Scan(&m.MachineId, &m.MachineNo, &m.MachineCode, &m.DefWhId, &m.DefShelfId, &m.WHCode, &m.ShelfCode)
	//err := db.Get(&m, lccommand, machine_id, company_id, branch_id)
	if err != nil {
		fmt.Println("machine error = ", err.Error())
	}

	fmt.Println("WHCode =", m.DefWhId, m.WHCode, m.MachineNo)

	return
}

func (m *Machine) SearchMachineNo(db *sqlx.DB, company_id int, branch_id int, access_token string) {
	fmt.Println("Search Machine UUID = ", access_token, company_id, branch_id)
	lccommand := "	select ifnull(c.id,0) as machine_id, ifnull(machine_no, '') as machine_no, ifnull(machine_code, '') as machine_code, def_wh_id, def_shelf_id, ifnull(d.wh_code, '') as wh_code, ifnull(e.shelf_code, '') as shelf_code, ifnull(b.shift_uid,'') as shift_uuid,ifnull(b.cashier_id,0) as cashier_id from user_access a  inner join shift b on a.user_id = b.cashier_id inner join pos_machine c on  b.machine_id = c.id inner join warehouse d on c.def_wh_id = d.id inner join warehouse_shelf e on c.def_shelf_id = e.id where   a.access_token = ? and year(b.open_time) = year(CURDATE()) and month(b.open_time) = month(CURDATE()) and day(b.open_time) = day(CURDATE()) and c.company_id = ? and c.branch_id = ?"
	rs := db.QueryRow(lccommand, access_token, company_id, branch_id)
	rs.Scan(&m.MachineId, &m.MachineNo, &m.MachineCode, &m.DefWhId, &m.DefShelfId, &m.WHCode, &m.ShelfCode, &m.ShiftUUID, &m.CashierID)
	//err := db.Get(&m, lccommand, access_token, company_id, branch_id)
	//if err != nil {
	//	fmt.Println("machine error = ",err.Error())
	//}

	fmt.Println("WHCode =", m.WHCode, m.MachineNo)

	return
}

func (s *ShiftModel) SearchUserShift(db *sqlx.DB, company_id int, branch_id int, access_token string) {
	fmt.Println("Search Machine UUID = ", access_token, company_id, branch_id)
	lccommand := "	select ifnull(c.id,0) as machine_id, ifnull(machine_no, '') as machine_no, def_wh_id as wh_id, ifnull(b.shift_uid,'') as shift_uuid,ifnull(b.cashier_id,0) as cashier_id,'' as cashier_name, change_amount, status, ifnull(remark,'') as remark, ifnull(open_by,'') as open_by, ifnull(open_time,'') as open_time,ifnull(close_by,'') as close_by,ifnull(close_time,'') as close_time,ifnull(sum_cash_amount,0) as sum_cash_amount, ifnull(sum_credit_amount,0) as sum_credit_amount, ifnull(sum_bank_amount,0) as sum_bank_amount, ifnull(sum_coupon_amount,0) as sum_coupon_amount, ifnull(sum_deposit_amount,0) as sum_deposit_amount from user_access a  inner join shift b on a.user_id = b.cashier_id inner join pos_machine c on  b.machine_id = c.id inner join warehouse d on c.def_wh_id = d.id inner join warehouse_shelf e on c.def_shelf_id = e.id where   a.access_token = ? and year(b.open_time) = year(CURDATE()) and month(b.open_time) = month(CURDATE()) and day(b.open_time) = day(CURDATE()) and c.company_id = ? and c.branch_id = ? and b.status = 0"
	rs := db.QueryRow(lccommand, access_token, company_id, branch_id)
	err := rs.Scan(&s.MachineID, &s.MachineNo, &s.WhID, &s.ShiftUUid, &s.CashierID, &s.CashierName, &s.ChangeAmount, &s.Status, &s.Remark, &s.OpenBy, &s.OpenTime, &s.CloseBy, &s.CloseTime, &s.SumOfCashAmount, &s.SumOfCreditAmount, &s.SumOfBankAmount, &s.SumOfCouponAmount, &s.SumOfDepositAmount)
	//err := db.Get(&m, lccommand, access_token, company_id, branch_id)
	if err != nil {
		fmt.Println("SearchUserShift error = ", err.Error())
	}

	//fmt.Println("WHCode =", s.machineID, s.MachineNo, s.changeAmount)

	return
}

func (u *UserAccess) GetProfileByToken1(db *sqlx.DB, token string) {

	lcCommand := "select user_id,user_code,b.company_id,b.branch_id,b.branch_code,b.zone_id " +
		",b.name from " + dbname + ".user_access a inner join npdl.`user` b on a.user_id=b.id " +
		" where access_token='" + token + "'"
	//fmt.Println(lcCommand)
	rs := db.QueryRow(lcCommand)
	rs.Scan(&u.UserId, &u.UserCode, &u.CompanyID, &u.BranchID, &u.BranchCode, &u.ZoneID, &u.Name)
	return
}

func (sh *ShiftModel) MachineOpenState(db *sqlx.DB) (state int) {
	lcCommand := `select is_open from pos_machine where id = ?`
	rs := db.QueryRow(lcCommand, sh.MachineID)
	rs.Scan(&state)
	return
}

func (sh *ShiftModel) CashierIsOpenShift(db *sqlx.DB) (count int) {
	lcCommand := `select count(*)  from shift where cashier_id = ? and status=?`
	rs := db.QueryRow(lcCommand, sh.CashierID, shiftStatusOpen)
	rs.Scan(&count)
	return
}

func (sh *ShiftModel) Open(db *sqlx.DB) (newuid string, err error) {
	// open pos_machine is_open=0 only
	if sh.MachineOpenState(db) == machineOpen {
		return "", fmt.Errorf(" Machine is already open ")
	}

	// check cashier_id is open another shift?
	if sh.CashierIsOpenShift(db) > 0 {
		return "", fmt.Errorf(" This Cashier is already open another shift ")
	}

	//fmt.Println("WH_ID = ", sh.docDate, sh.companyID, sh.branchID, sh.machineID, sh.cashierID, sh.changeAmount, sh.whID, sh.status, sh.shiftUUid, sh.openBy, sh.openTime, sh.remark)

	// insert new shift
	lcCommand := `insert into shift (
		doc_date,company_id,branch_id,machine_id,cashier_id,
		change_amount,wh_id,status,shift_uid,open_by,
		open_time,remark)
		values (?,?,?,?,?,?,?,?,?,?,?,?)`
	rs, err := db.Exec(lcCommand, sh.DocDate, sh.CompanyID, sh.BranchID, sh.MachineID, sh.CashierID, sh.ChangeAmount, sh.WhID, sh.Status, sh.ShiftUUid, sh.OpenBy, sh.OpenTime, sh.Remark)
	if err != nil {
		return "", err
	}
	newID, err := rs.LastInsertId()
	// update machine status by id
	lcCommand = `update pos_machine set is_open=? where id=? `
	_, err = db.Exec(lcCommand, machineOpen, sh.MachineID)
	if err != nil {
		return "", fmt.Errorf("error update machine status ")
	}

	var newShiftUUID string
	uid := db.QueryRow(`select shift_uid from shift where id =? `, newID)
	uid.Scan(&newShiftUUID)

	return newShiftUUID, nil
}

func (sh *ShiftModel) Details(db *sqlx.DB, shift_uuid string) () {
	lccommand := `select id,doc_date,shift_uid,company_id,branch_id,machine_id,cashier_id,change_amount,wh_id,status,open_by,open_time,sum_cash_amount,sum_credit_amount,sum_bank_amount,sum_coupon_amount,sum_deposit_amount,ifnull(close_by,'') as close_by,close_time,ifnull(remark,'') as remark from shift where shift_uid =? `
	uid := db.QueryRow(lccommand, shift_uuid)
	err := uid.Scan(&sh.Id, &sh.DocDate, &sh.ShiftUUid, &sh.CompanyID, &sh.BranchID, &sh.MachineID, &sh.CashierID, &sh.ChangeAmount, &sh.WhID, &sh.Status, &sh.OpenBy, &sh.OpenTime, &sh.SumOfCashAmount, &sh.SumOfCreditAmount, &sh.SumOfBankAmount, &sh.SumOfCouponAmount, &sh.SumOfDepositAmount, &sh.CloseBy, &sh.CloseTime, &sh.Remark)
	if err != nil {
		fmt.Println("err = ", err.Error())
	}
	return
}

func (sh *ShiftModel) Update(db *sqlx.DB) {
	// todo : Update data  shift to db
}

func (sh *ShiftModel) Close(db *sqlx.DB) error {
	// todo : check shift status must by open
	if sh.GetShiftStatus(db, sh.ShiftUUid) != shiftStatusOpen {
		return fmt.Errorf("shift is not open status..")
	}

	// todo: update status to close & update case,credit,coupon,deposit,bank amount
	lcCommand := `update shift set status = ? ,
	 			sum_cash_amount = ?,
	 			sum_credit_amount = ?,
	 			sum_bank_amount = ?,
	 			sum_coupon_amount=?,
	 			sum_deposit_amount=?,
	 			close_by = ?,
	 			close_time = ?
	 			where shift_uid = ?`
	rs, err := db.Exec(lcCommand,
		shiftStatusClose,
		sh.SumOfCashAmount,
		sh.SumOfCreditAmount,
		sh.SumOfBankAmount,
		sh.SumOfCouponAmount,
		sh.SumOfDepositAmount,
		sh.CloseBy,
		sh.CloseTime,
		sh.ShiftUUid)
	if err != nil {
		fmt.Println("error when update table shift to close status ...")
		return fmt.Errorf("Error When Update Table Shift : ", err.Error())
	}

	// check row update must by 1 record ...
	rowUpdate, err := rs.RowsAffected()
	if rowUpdate != 1 {
		return fmt.Errorf("Update not Equal %v record ...", rowUpdate)
	}

	sh.Details(db, sh.ShiftUUid)

	lcCommand = `update pos_machine set is_open = 0 where id = ?`
	rs, err = db.Exec(lcCommand, sh.MachineID)
	if err != nil {
		fmt.Println("error when update table machine  to close isopen ...")
		return fmt.Errorf("Error When Update Table machine : ", err.Error())
	}

	// todo : add type of cash (further)

	return nil

}

func (sh *ShiftModel) GetShiftStatus(db *sqlx.DB, shiftUUID string) (state int) {
	rs := db.QueryRow(`select status from shift where shift_uid = ? `, shiftUUID)
	err := rs.Scan(&state)
	if err != nil {
		fmt.Println("err GetShiftStatus = ", err.Error())
	}

	fmt.Println("status =", state)
	return
}

func (sh *ShiftModel) Get(db *sqlx.DB, uuid string) error {
	// todo : get exists shift data
	//db.Select(sh, `select uuid from npdl.shift where uid = ?`, uuid)

	fmt.Println(sh)
	return nil
}
