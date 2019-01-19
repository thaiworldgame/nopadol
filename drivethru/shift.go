package drivethru

import "time"

type ShiftOpenRequest struct {
	Token        string    `json:"uuid"`
	MachineID    int     `json:"machine_id"`
	CashierID    int     `json:"cashier_id"`
	ChangeAmount float64   `json:"change_amount"`
	Remark       string    `json:"remark"`
	Created      time.Time `json:"created"`
	WhID         int     `json:"wh_id"`
}
