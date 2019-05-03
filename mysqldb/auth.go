package mysqldb

import (
	//"fmt"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/auth"
	"github.com/satori/go.uuid"
)

type UserAccess struct {
	Id         int  `db:"id"`
	UserId     int  `db:"user_id"`
	UserCode   string `db:"user_code"`
	CompanyID  int    `db:"company_id"`
	BranchID   int    `db:"branch_id`
	BranchCode string `db:"branch_code"`
	ZoneID     string    `db:"zone_id"`
	Name       string `db:"name"`
	//BranchName string `db:"branch_name"`
}

//var dbname string = "demo"
var dbname string = "npdl"

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepository{db}
}

type authRepository struct {
	db *sqlx.DB
}

func (repo *authRepository) GetToken(tokenID string) (*auth.Token, error) {
	fmt.Println("start mysqldb.auth.GetToken ... with token ", tokenID)
	var m struct {
		ID         sql.NullInt64  `db:"id"`
		CompanyID  sql.NullInt64  `db:"company_id"`
		UserID     sql.NullInt64  `db:"user_id"`
		UserCode   sql.NullString `db:"user_code"`
		UserName   sql.NullString `db:"user_name"`
		BranchID   sql.NullInt64  `db:"branch_id"`
		BranchCode sql.NullString `db:"branch_code"`
		ZoneID     sql.NullString `db:"zone_id"`
		TokenID    sql.NullString `db:"token"`
		Created    time.Time      `db:"created"`
	}
	//fmt.Println("start repo *authRepository.GetToken")
	sqlcommand := "select b.id," +
		"user_id," +
		"user_code," +
		"b.company_id," +
		"b.branch_id," +
		"b.branch_code," +
		"b.zone_id, " +
		"b.name ," +
		"a.create_time " +
		"from " + dbname + ".user_access a inner join npdl.`user` b on a.user_id=b.id " +
		" where access_token='" + tokenID + "'"
	fmt.Println(sqlcommand)
	rows := repo.db.QueryRow(sqlcommand)
	//err := rows.Scan(
	rows.Scan(
		&m.ID,
		&m.UserID,
		&m.UserCode,
		&m.CompanyID,
		&m.BranchID,
		&m.BranchCode,
		&m.ZoneID,
		&m.UserName,
		&m.Created)

	// remark ชั่วคราว
	//if err == sql.ErrNoRows {
	//	return nil, auth.ErrTokenNotFound
	//}
	//if err != nil {
	//	return nil, err
	//}
	//expireTime := time.Now().Add(-(365 * 24 * time.Hour))
	//if m.Created.Before(expireTime) {
	//	return nil, auth.ErrTokenExpired
	//}

	tk := auth.Token{ID: tokenID}
	//fmt.Println("postgres.auth.go -> auth.Token.ID = ", tokenID)
	if m.CompanyID.Valid {
		tk.CompanyID = int(m.CompanyID.Int64)
	} else {
		tk.CompanyID = -1
	}
	if m.UserID.Valid {
		tk.UserID = m.UserID.Int64
	} else {
		tk.UserID = -1
	}
	if m.BranchID.Valid {
		tk.BranchID = m.BranchID.Int64
	} else {
		tk.BranchID = -1
	}
	if m.ZoneID.Valid {
		tk.ZoneID = m.ZoneID.String
	}
	if m.TokenID.Valid {
		tk.TokenID = m.TokenID.String
	}
	if m.UserName.Valid {
		tk.UserName = m.UserName.String
	}

	if m.UserCode.Valid {
		tk.UserCode = m.UserName.String
	}
	//fmt.Println("return postgres.auth.GetToken : ", tk)
	return &tk, nil
}

func (u *UserAccess) GetProfileByToken(db *sqlx.DB, token string) {

	lcCommand := "select b.id,user_id,user_code,b.company_id,b.branch_id,b.branch_code,b.zone_id ,ifnull(b.name,'') as name from user_access a inner join npdl.user b on a.user_id=b.id where access_token= ?"
	fmt.Println("user sql = ", lcCommand)
	rs := db.QueryRow(lcCommand, token)
	err := rs.Scan(&u.Id, &u.UserId, &u.UserCode, &u.CompanyID, &u.BranchID, &u.BranchCode, &u.ZoneID, &u.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Cashier Name = ", u.UserCode, u.CompanyID, u.BranchCode, u.Name)

	return
}

func GetAccessToken() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Get UUID Error = ", err)
		return err.Error()
	}
	fmt.Printf("UUIDv4: %s\n", uuid)

	return uuid.String()
}

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Get UUID Error = ", err)
		return err.Error()
	}
	fmt.Printf("UUIDv4: %s\n", uuid)

	return uuid.String()
}
