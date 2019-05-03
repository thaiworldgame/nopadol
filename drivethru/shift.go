package drivethru

type ShiftOpenRequest struct {
	AccessToken  string  `json:"access_token"`
	MachineID    string     `json:"machine_id"`
	ChangeAmount string `json:"change_amount"`
	Remark       string  `json:"remark"`
	WhID         string     `json:"wh_id"`
}

type ShiftCloseRequest struct {
	AccessToken      string  `json:"access_token"`
	ShiftUUID        string  `json:"shift_uuid"`
	SumCashAmount    string `json:"sum_cash_amount"`
	SumCreditAmount  string `json:"sum_credit_amount"`
	SumBankAmount    string `json:"sum_bank_amount"`
	SumCouponAmount  string `json:"sum_coupon_amount"`
	SumDepositAmount string `json:"sum_deposit_amount"`
}
