package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const ( machineOpen int = 1
	shiftStatusOpen  int = 0
	)

type ShiftModel struct {
	id           int64           `db:"id"`
	docDate      mysql.NullTime  `db:"doc_date"`
	companyID    int             `db:"company_id"`
	branchID     int             `db:"branch_id"`
	machineID    int             `db:"machine_id"`
	cashierID    int             `db:"cashier_id"`
	changeAmount sql.NullFloat64 `db:"change_amount"`
	whID         int             `db:"wh_id"`
	status       int             `db:"status"`
	shiftUUid    string          `db:"shift_uid"`
	openBy       string          `db:"open_by"`
	openTime     mysql.NullTime  `db:"open_time"`
	editBy       string          `db:"edit_by"`
	editTime     mysql.NullTime  `db:"edit_time"`
}
func (sh *ShiftModel)MachineOpenState(db *sqlx.DB)(state int){
	lcCommand := `select is_open from pos_machine where id = ?`
	rs := db.QueryRow(lcCommand,sh.machineID)
	rs.Scan(&state)
	return
}

func (sh *ShiftModel)CashierIsOpenShift(db *sqlx.DB)(count int){
	lcCommand := `select count(*)  from shift where cashier_id = ? and status=?`
	rs := db.QueryRow(lcCommand,sh.cashierID,shiftStatusOpen)
	rs.Scan(&count)
	return
}
func (sh *ShiftModel) Open(db *sqlx.DB) (newuid string, err error) {
	// open pos_machine is_open=0 only
	if sh.MachineOpenState(db)==machineOpen {
		return "",fmt.Errorf(" Machine is already open ")
	}

	// check cashier_id is open another shift?
	if sh.CashierIsOpenShift(db) >0 {
		return "",fmt.Errorf(" This Cashier is already open another shift ")
	}

	// insert new shift
	lcCommand := `insert into shift (
		doc_date,company_id,branch_id,machine_id,cashier_id,
		change_amount,wh_id,status,shift_uid,open_by,
		open_time)
		values (?,?,?,?,?,?,?,?,?,?,?)`

	rs,err := db.Exec(lcCommand,sh.docDate.Time,sh.companyID,sh.branchID,sh.machineID,sh.cashierID,
		sh.changeAmount.Float64,sh.whID,sh.status,sh.shiftUUid,sh.openBy,sh.openTime.Time)
	if err != nil {
		return "",err
	}
	newID,err := rs.LastInsertId()
	// update machine status by id
	lcCommand = `update pos_machine set is_open=? where id=? `
	_,err = db.Exec(lcCommand,machineOpen,sh.machineID)
	if err != nil {
		return "",fmt.Errorf("error update machine status ")
	}


	var newShiftUUID string
	uid := db.QueryRow(`select shift_uid from shift where id =? `,newID)
	uid.Scan(&newShiftUUID)

	return newShiftUUID,nil
}

func (sh *ShiftModel) Save(db *sqlx.DB) {
	// todo : save shift to db
}

func (sh *ShiftModel) Close(db *sqlx.DB) {
	// todo : close shift in db
}

func (sh *ShiftModel) Get(db *sqlx.DB, uuid string) error {
	// todo : get exists shift data
	db.Select(sh, `select uuid from npdl.shift where uid = ?`, uuid)
	fmt.Println(sh)
	return nil
}
