package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type CreditCard struct {
	CardNo       string  `json:"card_no"`
	ConfirmNo    string  `json:"confirm_no"`
	CreditType   string  `json:"credit_type"`
	BankCode     string  `json:"bank_code"`
	Amount       float64 `json:"amount"`
	ChargeAmount float64 `json:"charge_amount"`
}

func (c *CreditCard) CheckCreditCardUsed(db *sqlx.DB, card_no string, confirm_no string) (bool,string) {
	var exist int

	lccommand := `select count(*) as vCount from credit_card where credit_card_no = ? and confirm_no = ?`
	err := db.Get(&exist, lccommand, card_no, confirm_no)
	if err != nil {
		fmt.Println("error check credit card used = ", err.Error())
		return false, err.Error()
	}
	if exist == 0 {
		return true, ""
	} else {
		return false, "creditcard is used"
	}

}
