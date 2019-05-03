package drivethru

type ShiftOpenRequest struct {
	AccessToken  string  `json:"access_token"`
	MachineID    int     `json:"machine_id"`
	ChangeAmount float64 `json:"change_amount"`
	Remark       string  `json:"remark"`
	WhID         int     `json:"wh_id"`
}

type ShiftCloseRequest struct {
	AccessToken      string  `json:"access_token"`
	ShiftUUID        string  `json:"shift_uuid"`
	SumCashAmount    float64 `json:"sum_cash_amount"`
	SumCreditAmount  float64 `json:"sum_credit_amount"`
	SumBankAmount    float64 `json:"sum_bank_amount"`
	SumCouponAmount  float64 `json:"sum_coupon_amount"`
	SumDepositAmount float64 `json:"sum_deposit_amount"`
}
