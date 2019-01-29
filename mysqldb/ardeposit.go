package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Deposit struct {
	DepositId string  `json:"deposit_id"`
	Amount    float64 `json:"amount"`
}

func (c *Deposit) CheckArDepositUsed(db *sqlx.DB, ar_code string, deposit_no string, amount float64) (bool, string) {
	var exist int
	var remain int

	lccommand := `select count(*) as vCount from ar_deposit where ar_code = ? and doc_no = ?`
	err := db.Get(&exist, lccommand, ar_code, deposit_no)
	if err != nil {
		fmt.Println("error check deposit used = ", err.Error())
		return false, err.Error()
	}

	lccommand1 := `select count(*) as vCount from ar_deposit where ar_code = ? and doc_no = ? and bill_balance >= ? `
	err = db.Get(&remain, lccommand1, ar_code, deposit_no, amount)
	if err != nil {
		fmt.Println("error check deposit bill balancec = ", err.Error())
		return false, err.Error()
	}

	if exist != 0 && remain != 0 {
		return true, ""
	}else if exist == 0 {
		return false, "deposit is not exist"
	}else if remain == 0 {
		return false, "deposit amount is over than bill balance"
	}else{
		return true, ""
	}

}
