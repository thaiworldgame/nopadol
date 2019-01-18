package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
	"time"

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
	status int `db:"status"`
	openBy string `db:"open_by"`
	openTime mysql.NullTime `db:"open_time"`
	editBy string `db:"edit_by"`
	editTime mysql.NullTime
}



func (sh *ShiftModel) Open(db *sqlx.DB,token string,req drivethru.ShiftOpenRequest) ( newuid string ,err error) {
	// todo : create new Shift record in db
	uac := UserAccess{}
	uac.GetProfileByToken(db,token)
	sh.docDate = time.Now()
	sh.companyID = uac.CompanyID
	sh.branchID = uac.BranchID
	sh.cashierID = req.CashierCode
	return // new uuid
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
