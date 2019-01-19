package drivethru

import "time"

type ShiftOpenRequest struct {
	Token         string    `json:"uuid"`
	CashierCode   string    `json:"cashier_code"`
	ChangeAmounnt float64   `json:"change_amounnt"`
	Remark        string    `json:"remark"`
	Created       time.Time `json:"created"`
}
