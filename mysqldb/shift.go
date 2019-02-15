package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	machineOpen      int = 1
	shiftStatusOpen  int = 0
	shiftStatusClose int = 1
)

type ShiftModel struct {
	id                 int64           `db:"id"`
	docDate            mysql.NullTime  `db:"doc_date"`
	companyID          int             `db:"company_id"`
	branchID           int             `db:"branch_id"`
	machineID          int             `db:"machine_id"`
	cashierID          int             `db:"cashier_id"`
	changeAmount       sql.NullFloat64 `db:"change_amount"`
	whID               int             `db:"wh_id"`
	status             int             `db:"status"`
	shiftUUid          string          `db:"shift_uid"`
	openBy             string          `db:"open_by"`
	openTime           mysql.NullTime  `db:"open_time"`
	editBy             string          `db:"edit_by"`
	editTime           mysql.NullTime  `db:"edit_time"`
	closeTime          mysql.NullTime
	closeBy            string
	sumOfCashAmount    float64
	sumOfCreditAmount  float64
	sumOfBankAmount    float64
	sumOfCouponAmount  float64
	sumOfDepositAmount float64
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

func (m *Machine) SearchMachineNo(db *sqlx.DB, company_id int, branch_id int, access_token string) {
	fmt.Println("Search Machine UUID = ", access_token, company_id, branch_id)
	lccommand := "	select ifnull(c.id,0) as machine_id, ifnull(machine_no, '') as machine_no, ifnull(machine_code, '') as machine_code, def_wh_id, def_shelf_id, ifnull(d.wh_code, '') as wh_code, ifnull(e.shelf_code, '') as shelf_code, ifnull(b.shift_uid,'') as shift_uuid,ifnull(b.cashier_id,0) as cashier_id from        user_access a 	inner join shift b on a.user_id = b.cashier_id inner join pos_machine c on  b.machine_id = c.id inner join warehouse d on c.def_wh_id = d.id inner join warehouse_shelf e on c.def_shelf_id = e.id where   a.access_token = ? and year(b.open_time) = year(CURDATE()) and month(b.open_time) = month(CURDATE()) and day(b.open_time) = day(CURDATE()) and c.company_id = ? and c.branch_id = ?"

	rs := db.QueryRow(lccommand, access_token, company_id, branch_id)
	rs.Scan(&m.MachineId, &m.MachineNo, &m.MachineCode, &m.DefWhId, &m.DefShelfId, &m.WHCode, &m.ShelfCode, &m.ShiftUUID, &m.CashierID)
	//err := db.Get(&m, lccommand, company_id, branch_id, access_token)
	//if err != nil {
	//	fmt.Println("machine error = ",err.Error())
	//}

	fmt.Println("Machine No =", m.MachineNo)

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
	rs := db.QueryRow(lcCommand, sh.machineID)
	rs.Scan(&state)
	return
}

func (sh *ShiftModel) CashierIsOpenShift(db *sqlx.DB) (count int) {
	lcCommand := `select count(*)  from shift where cashier_id = ? and status=?`
	rs := db.QueryRow(lcCommand, sh.cashierID, shiftStatusOpen)
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

	// insert new shift
	lcCommand := `insert into shift (
		doc_date,company_id,branch_id,machine_id,cashier_id,
		change_amount,wh_id,status,shift_uid,open_by,
		open_time)
		values (?,?,?,?,?,?,?,?,?,?,?)`

	rs, err := db.Exec(lcCommand, sh.docDate.Time, sh.companyID, sh.branchID, sh.machineID, sh.cashierID,
		sh.changeAmount.Float64, sh.whID, sh.status, sh.shiftUUid, sh.openBy, sh.openTime.Time)
	if err != nil {
		return "", err
	}
	newID, err := rs.LastInsertId()
	// update machine status by id
	lcCommand = `update pos_machine set is_open=? where id=? `
	_, err = db.Exec(lcCommand, machineOpen, sh.machineID)
	if err != nil {
		return "", fmt.Errorf("error update machine status ")
	}

	var newShiftUUID string
	uid := db.QueryRow(`select shift_uid from shift where id =? `, newID)
	uid.Scan(&newShiftUUID)

	return newShiftUUID, nil
}

func (sh *ShiftModel) Update(db *sqlx.DB) {
	// todo : Update data  shift to db
}

func (sh *ShiftModel) Close(db *sqlx.DB) error {
	// todo : check shift status must by open
	if sh.GetShiftStatus(db, sh.shiftUUid) != shiftStatusOpen {
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
		sh.sumOfCashAmount,
		sh.sumOfCreditAmount,
		sh.sumOfBankAmount,
		sh.sumOfCouponAmount,
		sh.sumOfDepositAmount,
		sh.closeBy,
		sh.closeTime.Time,
		sh.shiftUUid)
	if err != nil {
		fmt.Println("error when update table shift to close status ...")
		return fmt.Errorf("Error When Update Table Shift : ", err.Error())
	}

	// check row update must by 1 record ...
	rowUpdate, err := rs.RowsAffected()
	if rowUpdate != 1 {
		return fmt.Errorf("Update not Equal %v record ...", rowUpdate)
	}

	// todo : add type of cash (further)

	return nil

}

func (sh *ShiftModel) GetShiftStatus(db *sqlx.DB, shiftUUID string) (state int) {
	rs := db.QueryRow(`select status from shift where shift_uid = ? `, shiftUUID)
	rs.Scan(&state)
	return
}

func (sh *ShiftModel) Get(db *sqlx.DB, uuid string) error {
	// todo : get exists shift data
	//db.Select(sh, `select uuid from npdl.shift where uid = ?`, uuid)

	fmt.Println(sh)
	return nil
}
