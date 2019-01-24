package drivethru

import "time"

type ShiftOpenRequest struct {
	Token        string    `json:"uuid"`
	MachineID    int       `json:"machine_id"`
	CashierID    int       `json:"cashier_id"`
	ChangeAmount float64   `json:"change_amount"`
	Remark       string    `json:"remark"`
	Created      time.Time `json:"created"`
	WhID         int       `json:"wh_id"`
}

type ShiftCloseRequest struct {
	Token            string
	ShiftUUID        string
	SumCashAmount    float64
	SumCreditAmount  float64
	SumBankAmount    float64
	SumCouponAmount  float64
	SumDepositAmount float64
	ClosedAt         time.Time
}
