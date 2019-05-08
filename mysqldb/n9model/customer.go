package n9model

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Customer struct {
	Id           int64     `db:"id"`
	Code         string    `db:"code"`
	Name         string    `db:"name"`
	Address      string    `db:"address"`
	Telephone    string    `db:"telephone"`
	BillCredit   int64     `db:"bill_credit"`
	DebtAmount   float64   `db:"debt_amount"`
	DebtLimit    float64   `db:"debt_limit"`
	Email        string    `db:"email"`
	CompanyID    int       `db:"company_id"`
	CreateBy     string    `db:"create_by"`
	CreateTime   time.Time `db:"create_time"`
	UpdateBy     string    `db:"update_by"`
	UpdateTime   time.Time `db:"update_time"`
	Fax          string    `db:"fax"`
	TaxNo        string    `db:"tax_no"`
	MemberID     string    `db:"member_id"`
	PointBalance float64   `db:"point_balance"`
}

func (c *Customer) GetByCode(db *sqlx.DB, code string) {

}

func (c *Customer) GetById(db *sqlx.DB, code string) {

}


func (c *Customer) Add(db *sqlx.DB) error {
	return nil
}

func (c *Customer) Update(db *sqlx.DB) error {
	return nil
}

func (c *Customer) Inactive(db *sqlx.DB, code string) error {
	return nil
}

func (c *Customer) ChangeCode(db *sqlx.DB,code string) error {
	return nil
}