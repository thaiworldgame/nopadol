package n9model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"time"
)

const customerInactive  = 1

type Customer struct {
	Id         int64   `db:"id"`
	Code       string  `db:"code"`
	Name       string  `db:"name"`
	Address    string  `db:"address"`
	Telephone  string  `db:"telephone"`
	BillCredit float64 `db:"bill_credit"`
	DebtAmount float64 `db:"debt_amount"`
	DebtLimit  float64 `db:"debt_limit"`
	//Email        string    `db:"email"`
	CompanyID    int       `db:"company_id"`
	CreateBy     string    `db:"create_by"`
	CreateTime   time.Time `db:"create_time"`
	EditBy       string    `db:"edit_by"`
	EditTime     time.Time `db:"edit_time"`
	Fax          string    `db:"fax"`
	TaxNo        string    `db:"tax_no"`
	MemberID     string    `db:"member_id"`
	PointBalance float64   `db:"point_balance"`
	ActiveStatus int64     `db:"active_status"`
}

func (c *Customer) GetIdByCode(db *sqlx.DB, code string) (int64, error) {
	sql := "select id from Customer where code='" + code + "'"
	fmt.Println(sql)
	var curID int64
	db.Get(&curID, sql)
	fmt.Println("current id -> ", curID)
	return curID, nil
}

func (c *Customer) Get(db *sqlx.DB, code string) error {
	sql := "select id,code,name,address,isnull(telephone) as telephone," +
		"isnull(bill_credit) as bill_credit," +
		"isnull(debt_amount) as debt_amount," +
		"isnull(debt_limit) as debt_limit," +
		"isnull(company_id) as company_id," +
		"create_by," +
		"create_time," +
		"edit_by," +
		"edit_time," +
		"fax,tax_no," +
		"isnull(member_id) as member_id," +
		"isnull(point_balance) as point_balance ," +
		"active_status " +
		" from Customer " +
		"where code='" + code + "'"
	fmt.Println(sql)
	err := db.Get(c, sql)
	if err != nil {
		log.Println("error ", err.Error())
		return err
	}
	return nil
}

func (c *Customer) GetById(db *sqlx.DB, id int64) error {
	sql := "select id,code,name,address,isnull(telephone) as telephone," +
		"isnull(bill_credit) as bill_credit," +
		"isnull(debt_amount) as debt_amount," +
		"isnull(debt_limit) as debt_limit," +
		"isnull(company_id) as company_id," +
		"create_by," +
		"create_time," +
		"edit_by," +
		"edit_time," +
		"fax,tax_no," +
		"isnull(member_id) as member_id," +
		"isnull(point_balance) as point_balance, " +
		"active_status " +
		" from Customer " +
		"where id=" + strconv.FormatInt(id, 10)
	fmt.Println(sql)
	err := db.Get(c, sql)
	if err != nil {
		log.Println("error ", err.Error())
		return err
	}
	return nil
}

func (c *Customer) Add(db *sqlx.DB) (err error) {
	if c.CheckExistById(db, c.Id) {
		return fmt.Errorf("invalid duplicate data : %v", c.Name)
	}
	// todo : implement insert to db

	return nil
}

func (c *Customer) Update(db *sqlx.DB) error {
	// if not exits data return error
	if c.CheckExistById(db, c.Id) == false {
		return fmt.Errorf("can not find customer id  : %v", c.Id)
	}
	//todo : update customer data

	return nil
}

func (c *Customer) Inactive(db *sqlx.DB, id  int64) (row int64 ,err error ){
	sql := `update Customer set active_status = ? where id = ?`
	rs,err := db.Exec(sql,customerInactive,id)
	if err != nil {
		return 0,fmt.Errorf("error inactive ")
	}
	val,_ := rs.RowsAffected()
	return val,nil
}

func (c *Customer) ChangeCode(db *sqlx.DB, code string) error {
	// todo : implement changeCode
	return nil
}

func (c *Customer) CheckExistById(db *sqlx.DB, id int64) bool {

	sql := "select id from Customer where id=" + strconv.FormatInt(id, 10)
	fmt.Println(sql)
	var curID int64
	db.Get(&curID, sql)
	fmt.Println("current id -> ", curID)
	if curID == 0 {
		return false
	}
	return true
}
func (c *Customer) CheckExistByCode(db *sqlx.DB, code string) bool {

	sql := "select id from Customer where code='"+code+"'"
	fmt.Println(sql)
	var curID int64
	db.Get(&curID, sql)
	fmt.Println("current id -> ", curID)
	if curID == 0 {
		return false
	}
	return true
}


