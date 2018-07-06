package mysqldb

//Pos///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type OutPutTax struct {
	TaxNo    string `json:"tax_no"`
	TaxDate  string `json:"tax_date"`
	BookCode string `json:"book_code"`
}

type Customer struct {
	ArCode string `json:"ar_code"`
}

type SaleMan struct {
	SaleCode string `json:"sale_code"`
}

type NewPosModel struct {
	SaveFrom        int               `json:"save_from"`
	Source          int               `json:"source"`
	DocNo           string            `json:"doc_no"`
	DocDate         string            `json:"doc_date"`
	OutPutTax
	Customer
	SaleMan
	ShiftCode       string            `json:"shiftcode"`
	CashierCode     string            `json:"cashier_code"`
	ShiftNo         string            `json:"shift_no"`
	MachineNo       string            `json:"machine_no"`
	MachineCode     string            `json:"machine_code"`
	CoupongAmount   float64           `json:"coupong_amount"`
	ChangeAmount    float64           `json:"change_amount"`
	ChargeAmount    float64           `json:"charge_amount"`
	TaxType         int               `json:"tax_type"`
	MyDescription   string            `json:"my_description"`
	SumOfItemAmount float64           `json:"sum_of_item_amount"`
	DiscountWord    string            `json:"discount_word"`
	AfterDiscount   float64           `json:"after_discount"`
	TotalAmount     float64           `json:"total_amount"`
	SumCashAmount   float64           `json:"sum_cash_amount"`
	SumChqAmount    float64           `json:"sum_chq_amount"`
	SumCreditAmount float64           `json:"sum_credit_amount"`
	SumBankAmount   float64           `json:"sum_bank_amount"`
	NetDebtAmount   float64           `json:"net_debt_amount"`
	UserCode        string            `json:"user_code"`
	PosSubs         []NewPosItemModel `json:"pos_subs"`
}

type NewPosItemModel struct {
	ItemCode     string  `json:"item_code"`
	ItemName     string  `json:"item_name"`
	WHCode       string  `json:"wh_code"`
	ShelfCode    string  `json:"shelf_code"`
	Qty          float64 `json:"qty"`
	Price        float64 `json:"price"`
	DiscountWord string  `json:"discount_word"`
	UnitCode     string  `json:"unit_code"`
	LineNumber   int     `json:"line_number"`
	BarCode      string  `json:"bar_code"`
	AverageCost  float64 `json:"averagecost"`
	PackingRate1 float64 `json:"packing_rate_1"`
}
