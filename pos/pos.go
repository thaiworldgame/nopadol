package pos




type NewPosResponseTemplate struct {
	Id int64 `json:"id"`
}

type NewPosTemplate struct {
	DocNo           string                   `json:"doc_no"`
	DocDate         string                   `json:"doc_date"`
	ArCode          string                   `json:"ar_code"`
	BookCode        string                   `json:"book_code"`
	SaleCode        string                   `json:"sale_code"`
	ShiftCode       string                   `json:"shiftcode"`
	CashierCode     string                   `json:"cashier_code"`
	ShiftNo         int                      `json:"shift_no"`
	MachineNo       string                   `json:"machine_no"`
	MachineCode     string                   `json:"machine_code"`
	CoupongAmount   float64                  `json:"coupong_amount"`
	ChangeAmount    float64                  `json:"change_amount"`
	ChargeAmount    float64                  `json:"charge_amount"`
	TaxType         int                      `json:"tax_type"`
	SumOfItemAmount float64                  `json:"sum_of_item_amount"`
	DiscountWord    string                   `json:"discount_word"`
	DiscountAmount  float64                  `json:"discount_amount"`
	AfterDiscount   float64                  `json:"after_discount"`
	TotalAmount     float64                  `json:"total_amount"`
	SumCashAmount   float64                  `json:"sum_cash_amount"`
	SumChqAmount    float64                  `json:"sum_chq_amount"`
	SumCreditAmount float64                  `json:"sum_credit_amount"`
	SumBankAmount   float64                  `json:"sum_bank_amount"`
	BankNo          string                   `json:"bank_no"`
	NetDebtAmount   float64                  `json:"net_debt_amount"`
	UserCode        string                   `json:"user_code"`
	ChqIns          []ListChqInTemplate      `json:"chq_ins"`
	CreditCards     []ListCreditCardTemplate `json:"credit_cards"`
	PosSubs         []NewPosItemTemplate     `json:"pos_subs"`
}

type NewPosItemTemplate struct {
	ItemCode       string  `json:"item_code"`
	ItemName       string  `json:"item_name"`
	WHCode         string  `json:"wh_code"`
	ShelfCode      string  `json:"shelf_code"`
	Qty            float64 `json:"qty"`
	Price          float64 `json:"price"`
	DiscountWord   string  `json:"discount_word"`
	DiscountAmount float64 `json:"discount_amount"`
	UnitCode       string  `json:"unit_code"`
	LineNumber     int     `json:"line_number"`
	BarCode        string  `json:"bar_code"`
	AverageCost    float64 `json:"averagecost"`
	PackingRate1   float64 `json:"packing_rate_1"`
}

type ListChqInTemplate struct {
	ChqNumber      string  `json:"chq_number"`
	BankCode       string  `json:"bank_code"`
	BankBranchCode string  `json:"bank_branch_code"`
	BookNo         string  `json:"book_no"`
	ReceiveDate    string  `json:"receive_date"`
	DueDate        string  `json:"due_date"`
	Status         int     `json:"status"`
	Amount         float64 `json:"amount"`
	Balance        float64 `json:"balance"`
	RefChqRowOrder int     `json:"ref_chq_row_order"`
	StatusDate     string  `json:"status_date"`
	StatusDocNo    string  `json:"status_doc_no"`
}

type ListCreditCardTemplate struct {
	BankCode       string  `json:"bank_code"`
	CreditCardNo   string  `json:"credit_card_no"`
	ReceiveDate    string  `json:"receive_date"`
	DueDate        string  `json:"due_date"`
	BookNo         string  `json:"book_no"`
	Status         int     `json:"status"`
	StatusDate     string  `json:"status_date"`
	StatusDocNo    string  `json:"status_doc_no"`
	BankBranchCode string  `json:"bank_branch_code"`
	Amount         float64 `json:"amount"`
	MyDescription  string  `json:"my_description"`
	CreditType     string  `json:"credit_type"`
	ConfirmNo      string  `json:"confirm_no"`
	ChargeAmount   float64 `json:"charge_amount"`
}

type SearchPosByIdResponseTemplate struct {
	Id              int                      `json:"id"`
	DocNo           string                   `json:"doc_no"`
	DocDate         string                   `json:"doc_date"`
	TaxNo           string                   `json:"tax_no"`
	TaxDate         string                   `json:"tax_date"`
	PosStatus       int                      `json:"pos_status"`
	ArCode          string                   `json:"ar_code"`
	ArName          string                   `json:"ar_name"`
	SaleCode        string                   `json:"sale_code"`
	SaleName        string                   `json:"sale_name"`
	ShiftCode       string                   `json:"shiftcode"`
	CashierCode     string                   `json:"cashier_code"`
	ShiftNo         int                      `json:"shift_no"`
	MachineNo       string                   `json:"machine_no"`
	MachineCode     string                   `json:"machine_code"`
	CoupongAmount   float64                  `json:"coupong_amount"`
	ChangeAmount    float64                  `json:"change_amount"`
	ChargeAmount    float64                  `json:"charge_amount"`
	TaxType         int                      `json:"tax_type"`
	SumOfItemAmount float64                  `json:"sum_of_item_amount"`
	DiscountWord    string                   `json:"discount_word"`
	DiscountAmount  float64                  `json:"discount_amount"`
	AfterDiscount   float64                  `json:"after_discount"`
	BeforeTaxAmount float64                  `json:"before_tax_amount"`
	TaxAmount       float64                  `json:"tax_amount"`
	TotalAmount     float64                  `json:"total_amount"`
	SumCashAmount   float64                  `json:"sum_cash_amount"`
	SumChqAmount    float64                  `json:"sum_chq_amount"`
	SumCreditAmount float64                  `json:"sum_credit_amount"`
	SumBankAmount   float64                  `json:"sum_bank_amount"`
	BankNo          string                   `json:"bank_no"`
	NetDebtAmount   float64                  `json:"net_debt_amount"`
	IsCancel        int                      `json:"is_cancel"`
	IsConfirm       int                      `json:"is_confirm"`
	CreatorCode     string                   `json:"creator_code"`
	CreateDateTime  string                   `json:"create_date_time"`
	LastEditorCode  string                   `json:"last_editor_code"`
	LastEditDateT   string                   `json:"last_edit_date_t"`
	ChqIns          []ListChqInTemplate      `json:"chq_ins"`
	CreditCards     []ListCreditCardTemplate `json:"credit_cards"`
	PosSubs         []NewPosItemTemplate     `json:"pos_subs"`
}

type SearchPosByIdRequestTemplate struct {
	Id int64 `json:"id"`
}

type SearchPosByKeywordRequestTemplate struct {
	Keyword string `json:"keyword"`
}
