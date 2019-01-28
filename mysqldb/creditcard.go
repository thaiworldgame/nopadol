package mysqldb

import "github.com/jmoiron/sqlx"

type CreditCard struct {
	CardNo       string  `json:"card_no"`
	ConfirmNo    string  `json:"confirm_no"`
	CreditType   string  `json:"credit_type"`
	BankCode     string  `json:"bank_code"`
	Amount       float64 `json:"amount"`
	ChargeAmount float64 `json:"charge_amount"`
}

func (c *CreditCard)CheckCreditCardUsed(db *sqlx.DB, card_no string, confirm_no string, credit_type string, bank_code string ,amount float64) bool {
	return true
}