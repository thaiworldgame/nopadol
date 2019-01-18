package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserAccess struct {
	UserId     int64  `db:"id"`
	UserCode   string `db:"user_code"`
	CompanyID  int    `db:"company_id"`
	BranchID   int    `db:"branch_id`
	BranchCode string `db:"branch_code"`
	ZoneID     int    `db:"zone_id"`
	Name       string `db:"name"`
	//BranchName string `db:"branch_name"`

}

func (u *UserAccess) GetProfileByToken(db *sqlx.DB, token string) {
	lcCommand := "select user_id,user_code,b.company_id,b.branch_id,b.branch_code,b.zone_id " +
		",b.name from npdl.user_access a inner join npdl.`user` b on a.user_id=b.id " +
		" where access_token='" + token + "'"
	fmt.Println(lcCommand)
	rs := db.QueryRow(lcCommand)
	rs.Scan(&u.UserId,&u.UserCode,&u.CompanyID,&u.BranchID,&u.BranchCode,&u.ZoneID,&u.Name)
	return
}
