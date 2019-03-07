package sales

type NewQuoTemplate struct {
	Id                  int64                `json:"id"`
	DocNo               string               `json:"doc_no"`
	DocDate             string               `json:"doc_date"`
	CompanyId           int64                `json:"company_id"`
	BranchId            int64                `json:"branch_id"`
	DocType             int64                `json:"doc_type"`
	ArId                int64                `json:"ar_id"`
	ArCode              string               `json:"ar_code"`
	ArName              string               `json:"ar_name"`
	ArBillAddress       string               `json:"ar_bill_address"`
	ArTelephone         string               `json:"ar_telephone"`
	SaleId              int64                `json:"sale_id"`
	SaleCode            string               `json:"sale_code"`
	SaleName            string               `json:"sale_name"`
	BillType            int64                `json:"bill_type"`
	TaxType             int64                `json:"tax_type"`
	TaxRate             float64              `json:"tax_rate"`
	DepartId            int64                `json:"depart_id"`
	RefNo               string               `json:"ref_no"`
	JobId               string               `json:"job_id"`
	IsConfirm           int64                `json:"is_confirm"`
	BillStatus          int64                `json:"bill_status"`
	Validity            int64                `json:"validity"`
	CreditDay           int64                `json:"credit_day"`
	DueDate             string               `json:"due_date"`
	ExpireCredit        int64                `json:"expire_credit"`
	ExpireDate          string               `json:"expire_date"`
	DeliveryDay         int64                `json:"delivery_day"`
	DeliveryDate        string               `json:"delivery_date"`
	AssertStatus        int64                `json:"assert_status"`
	IsConditionSend     int64                `json:"is_condition_send"`
	MyDescription       string               `json:"my_description"`
	SumOfItemAmount     float64              `json:"sum_of_item_amount"`
	DiscountWord        string               `json:"discount_word"`
	DiscountAmount      float64              `json:"discount_amount"`
	AfterDiscountAmount float64              `json:"after_discount_amount"`
	BeforeTaxAmount     float64              `json:"before_tax_amount"`
	TaxAmount           float64              `json:"tax_amount"`
	TotalAmount         float64              `json:"total_amount"`
	NetDebtAmount       float64              `json:"net_debt_amount"`
	ProjectId           int64                `json:"project_id"`
	AllocateId          int64                `json:"allocate_id"`
	IsCancel            int64                `json:"is_cancel"`
	CreateBy            string               `json:"creator_by"`
	CreateTime          string               `json:"create_time"`
	EditBy              string               `json:"edit_by"`
	EditTime            string               `json:"edit_time"`
	ConfirmBy           string               `json:"confirm_by"`
	ConfirmTime         string               `json:"confirm_time"`
	CancelBy            string               `json:"cancel_by"`
	CancelTime          string               `json:"cancel_time"`
	Subs                []NewQuoItemTemplate `json:"subs"`
}

type NewQuoItemTemplate struct {
	Id              int64   `json:"id"`
	QuoId           int64   `json:"quo_id"`
	ItemId          int64   `json:"item_id"`
	ItemCode        string  `json:"item_code"`
	BarCode         string  `json:"bar_code"`
	ItemName        string  `json:"item_name"`
	Qty             float64 `json:"qty"`
	RemainQty       float64 `json:"remain_qty"`
	Price           float64 `json:"price"`
	DiscountWord    string  `json:"discount_word"`
	DiscountAmount  float64 `json:"discount_amount"`
	UnitCode        string  `json:"unit_code"`
	ItemAmount      float64 `json:"item_amount"`
	ItemDescription string  `json:"item_description"`
	PackingRate1    float64 `json:"packing_rate_1"`
	LineNumber      int     `json:"line_number"`
	IsCancel        int64   `json:"is_cancel"`
}

type NewSaleTemplate struct {
	Id                  int64                 `json:"id"`
	Uuid                string                `json:"uuid"`
	DocNo               string                `json:"doc_no"`
	DocDate             string                `json:"doc_date"`
	CompanyId           int64                 `json:"company_id"`
	BranchId            int64                 `json:"branch_id"`
	DocType             int64                 `json:"doc_type"`
	ArId                int64                 `json:"ar_id"`
	ArCode              string                `json:"ar_code"`
	ArName              string                `json:"ar_name"`
	ArBillAddress       string                `json:"ar_bill_address"`
	ArTelephone         string                `json:"ar_telephone"`
	SaleId              int64                 `json:"sale_id"`
	SaleCode            string                `json:"sale_code"`
	SaleName            string                `json:"sale_name"`
	BillType            int64                 `json:"bill_type"`
	TaxType             int64                 `json:"tax_type"`
	TaxRate             float64               `json:"tax_rate"`
	DepartId            int64                 `json:"depart_id"`
	RefNo               string                `json:"ref_no"`
	IsConfirm           int64                 `json:"is_confirm"`
	BillStatus          int64                 `json:"bill_status"`
	HoldingStatus       int64                 `json:"holding_status"`
	CreditDay           int64                 `json:"credit_day"`
	DueDate             string                `json:"due_date"`
	DeliveryDay         int64                 `json:"delivery_day"`
	DeliveryDate        string                `json:"delivery_date"`
	IsConditionSend     int64                 `json:"is_condition_send"`
	DeliveryAddressId   int64                 `json:"delivery_address_id"`
	CarLicense          string                `json:"car_license"`
	PersonReceiveTel    string                `json:"person_receive_tel"`
	MyDescription       string                `json:"my_description"`
	SumOfItemAmount     float64               `json:"sum_of_item_amount"`
	DiscountWord        string                `json:"discount_word"`
	DiscountAmount      float64               `json:"discount_amount"`
	AfterDiscountAmount float64               `json:"after_discount_amount"`
	BeforeTaxAmount     float64               `json:"before_tax_amount"`
	TaxAmount           float64               `json:"tax_amount"`
	TotalAmount         float64               `json:"total_amount"`
	NetDebtAmount       float64               `json:"net_debt_amount"`
	ProjectId           int64                 `json:"project_id"`
	AllocateId          int64                 `json:"allocate_id"`
	JobId               string                `json:"job_id"`
	IsCancel            int64                 `json:"is_cancel"`
	CreateBy            string                `json:"create_by"`
	CreateTime          string                `json:"create_time"`
	EditBy              string                `json:"edit_by"`
	EditTime            string                `json:"edit_time"`
	ConfirmBy           string                `json:"confirm_by"`
	ConfirmTime         string                `json:"confirm_time"`
	CancelBy            string                `json:"cancel_by"`
	CancelTime          string                `json:"cancel_time"`
	Subs                []NewSaleItemTemplate `json:"subs"`
}

type NewSaleItemTemplate struct {
	Id              int64   `json:"id"`
	SOId            int64   `json:"so_id"`
	ItemId          int64   `json:"item_id"`
	ItemCode        string  `json:"item_code"`
	BarCode         string  `json:"bar_code"`
	ItemName        string  `json:"item_name"`
	WHCode          string  `json:"wh_code"`
	ShelfCode       string  `json:"shelf_code"`
	Qty             float64 `json:"qty"`
	RemainQty       float64 `json:"remain_qty"`
	Price           float64 `json:"price"`
	DiscountWord    string  `json:"discount_word"`
	DiscountAmount  float64 `json:"discount_amount"`
	UnitCode        string  `json:"unit_code"`
	ItemAmount      float64 `json:"item_amount"`
	ItemDescription string  `json:"item_description"`
	StockType       int64   `json:"stock_type"`
	AverageCost     float64 `json:"average_cost"`
	SumOfCost       float64 `json:"sum_of_cost"`
	PackingRate1    float64 `json:"packing_rate_1"`
	RefNo           string  `json:"ref_no"`
	QuoId           int64   `json:"quo_id"`
	LineNumber      int     `json:"line_number"`
	RefLineNumber   int64   `json:"ref_line_number"`
	IsCancel        int64   `json:"is_cancel"`
}

type SearchByIdTemplate struct {
	Id int64 `json:"id"`
}
type SearchcreditcardTamplate struct {
	Keyword string `json:"keyword"`
}
type SearchByKeywordTemplate struct {
	ArId     int64  `json:"ar_id"`
	SaleCode string `json:"sale_code"`
	Keyword  string `json:"keyword"`
}
type SearchByItemTemplate struct {
	ArId     string `json:"ar_id"`
	ArCode   string `json:"ar_code"`
	ItemCode string `json:"item_code"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Page     string `json:"page"`
}

type SearchDocTemplate struct {
	Id            int64   `json:"id"`
	DocNo         string  `json:"doc_no"`
	DocDate       string  `json:"doc_date"`
	Module        string  `json:"module"`
	ArCode        string  `json:"ar_code"`
	ArName        string  `json:"ar_name"`
	SaleCode      string  `json:"sale_code"`
	SaleName      string  `json:"sale_name"`
	MyDescription string  `json:"my_description"`
	TotalAmount   float64 `json:"total_amount"`
	IsCancel      int     `json:"is_cancel"`
	IsConfirm     int     `json:"is_confirm"`
}
type SearchInvTemplate struct {
	Id            int64   `json:"id"`
	DocNo         string  `json:"doc_no"`
	DocDate       string  `json:"doc_date"`
	Doctype       string  `json:"doc_type"`
	ArCode        string  `json:"ar_code"`
	ArName        string  `json:"ar_name"`
	SaleCode      string  `json:"sale_code"`
	SaleName      string  `json:"sale_name"`
	MyDescription string  `json:"my_description"`
	TotalAmount   float64 `json:"total_amount"`
	IsCancel      int     `json:"is_cancel"`
	IsConfirm     int     `json:"is_confirm"`
}
type CreditCardTypeTemplate struct {
	Id                 int64  `json:"id"`
	CreditCardTypeName string `json:"creditcardtype_name"`
}
type SearchIVDocTemplate struct {
	Id            int64   `json:"id"`
	DocNo         string  `json:"doc_no"`
	DocDate       string  `json:"doc_date"`
	ArCode        string  `json:"ar_code"`
	ArName        string  `json:"ar_name"`
	SaleCode      string  `json:"sale_code"`
	SaleName      string  `json:"sale_name"`
	MyDescription string  `json:"my_description"`
	TotalAmount   float64 `json:"total_amount"`
	IsCancel      int     `json:"is_cancel"`
	IsConfirm     int     `json:"is_confirm"`
}
type NewDepositTemplate struct {
	Id               int64                    `json:"id"`
	CompanyId        int64                    `json:"company_id"`
	BranchId         int64                    `json:"branch_id"`
	Uuid             string                   `json:"uuid"`
	DocNo            string                   `json:"doc_no"`
	TaxNo            string                   `json:"tax_no"`
	DocDate          string                   `json:"doc_date"`
	BillType         int64                    `json:"bill_type"`
	ArId             int64                    `json:"ar_id"`
	ArCode           string                   `json:"ar_code"`
	ArName           string                   `json:"ar_name"`
	ArBillAddress    string                   `json:"ar_bill_address"`
	ArTelephone      string                   `json:"ar_telephone"`
	SaleId           int64                    `json:"sale_id"`
	SaleCode         string                   `json:"sale_code"`
	SaleName         string                   `json:"sale_name"`
	TaxType          int64                    `json:"tax_type"`
	TaxRate          float64                  `json:"tax_rate"`
	RefNo            string                   `json:"ref_no"`
	CreditDay        int64                    `json:"credit_day"`
	DueDate          string                   `json:"due_date"`
	DepartId         int64                    `json:"depart_id"`
	AllocateId       int64                    `json:"allocate_id"`
	ProjectId        int64                    `json:"project_id"`
	MyDescription    string                   `json:"my_description"`
	BeforeTaxAmount  float64                  `json:"before_tax_amount"`
	TaxAmount        float64                  `json:"tax_amount"`
	TotalAmount      float64                  `json:"total_amount"`
	NetAmount        float64                  `json:"net_amount"`
	BillBalance      float64                  `json:"bill_balance"`
	CashAmount       float64                  `json:"cash_amount"`
	CreditcardAmount float64                  `json:"creditcard_amount"`
	ChqAmount        float64                  `json:"chq_amount"`
	BankAmount       float64                  `json:"bank_amount"`
	IsReturnMoney    int64                    `json:"is_return_money" `
	IsCancel         int64                    `json:"is_cancel"`
	IsConfirm        int64                    `json:"is_confirm"`
	ScgId            string                   `json:"scg_id"`
	JobNo            string                   `json:"job_no"`
	CreateBy         string                   `json:"create_by"`
	CreateTime       string                   `json:"create_time"`
	EditBy           string                   `json:"edit_by"`
	EditTime         string                   `json:"edit_time"`
	CancelBy         string                   `json:"cancel_by"`
	CancelTime       string                   `json:"cancel_time" `
	ConfirmBy        string                   `json:"confirm_by"`
	ConfirmTime      string                   `json:"confirm_time"`
	Subs             []NewDepositItemTemplate `json:"subs"`
	RecMoney         []RecMoney               `json:"rec_money"`
	CreditCard       []CreditCardTemplate     `json:"credit_card"`
	Chq              []ChqInTemplate          `json:"chq"`
}

type NewDepositItemTemplate struct {
	Id              int64   `json:"id"`
	SORefNo         string  `json:"so_ref_no"`
	SOId            int64   `json:"so_id"`
	ItemId          int64   `json:"item_id"`
	ItemCode        string  `json:"item_code"`
	BarCode         string  `json:"bar_code"`
	ItemName        string  `json:"item_name"`
	WHCode          string  `json:"wh_code"`
	ShelfCode       string  `json:"shelf_code"`
	Qty             float64 `json:"qty"`
	RemainQty       float64 `json:"remain_qty"`
	Price           float64 `json:"price"`
	DiscountWord    string  `json:"discount_word"`
	DiscountAmount  float64 `json:"discount_amount"`
	UnitCode        string  `json:"unit_code"`
	ItemAmount      float64 `json:"item_amount"`
	ItemDescription string  `json:"item_description"`
	PackingRate1    float64 `json:"packing_rate_1"`
	RefNo           string  `json:"ref_no"`
	QuoId           int64   `json:"quo_id"`
	LineNumber      int     `json:"line_number"`
	RefLineNumber   int64   `json:"ref_line_number"`
	IsCancel        int64   `json:"is_cancel"`
}

type CreditCardTemplate struct {
	Id           int64   `json:"id"`
	RefId        int64   `json:"ref_id"`
	CreditCardNo string  `json:"credit_card_no"`
	CreditType   string  `json:"credit_type"`
	ConfirmNo    string  `json:"confirm_no"`
	Amount       float64 `json:"amount"`
	ChargeAmount float64 `json:"charge_amount"`
	Description  string  `json:"description"`
	BankId       int64   `json:"bank_id"`
	BankBranchId int64   `json:"bank_branch_id"`
	ReceiveDate  string  `json:"receive_date"`
	DueDate      string  `json:"due_date"`
	BookId       int64   `json:"book_id"`
}

type ChqInTemplate struct {
	Id           int64   `json:"id"`
	RefId        int64   `json:"ref_id"`
	ChqNumber    string  `json:"chq_number"`
	BankId       int64   `json:"bank_id"`
	BankBranchId int64   `json:"bank_branch_id"`
	ReceiveDate  string  `json:"receive_date"`
	DueDate      string  `json:"due_date"`
	BookId       int64   `json:"book_id"`
	ChqStatus    int64   `json:"chq_status"`
	ChqAmount    float64 `json:"chq_amount"`
	ChqBalance   float64 `json:"chq_balance"`
	Description  string  `json:"description"`
}
type BankpayTemplate struct {
	Id           int64   `json:"id"`
	RefId        int64   `json:"ref_id"`
	BankAccount  string  `json:"bank_account"`
	BankName     string  `json:"bank_name"`
	BankAmount   float64 `json:"bank_amount"`
	Activestatus int64   `json:"active_status"`
	CreateBy     string  `json:"create_by"`
	EditBy       string  `json:"edit_by"`
}

type RecMoneyTemplate struct {
	Id             int64   `json:"id"`
	DocType        int64   `json:"doc_type"`
	RefId          int64   `json:"ref_id"`
	ArId           int64   `json:"ar_id"`
	PaymentType    int64   `json:"payment_type"`
	PayAmount      float64 `json:"pay_amount"`
	ChqTotalAmount float64 `json:"chq_total_amount"`
	CreditType     int64   `json:"credit_type"`
	ChargeAmount   float64 `json:"charge_amount"`
	ConfirmNo      string  `json:"confirm_no"`
	RefNo          string  `json:"ref_no"`
	BankCode       string  `json:"bank_code"`
	BankBranchCode string  `json:"bank_branch_code"`
	RefDate        string  `json:"ref_date"`
	BankTransDate  string  `json:"bank_trans_date"`
	LineNumber     int64   `json:"line_number"`
}

type NewInvoiceTemplate struct {
	Id                  int64   `json:"id"`
	CompanyId           int64   `json:"company_id"`
	BranchId            int64   `json:"branch_id"`
	ItemName            string  `json:"item_name"`
	ItemCode            string  `json:"item_code"`
	Uuid                string  `json:"uuid"`
	DocNo               string  `json:"doc_no"`
	TaxNo               string  `json:"tax_no"`
	BillType            int64   `json:"bill_type"`
	DocDate             string  `json:"doc_date"`
	DocType             int64   `json:"doc_type"`
	ArId                int64   `json:"ar_id"`
	ArCode              string  `json:"ar_code"`
	ArName              string  `json:"ar_name"`
	ArBillAddress       string  `json:"ArBillAddress"`
	ArTelephone         string  `json:"ArTelephone"`
	SaleId              int64   `json:"sale_id"`
	SaleCode            string  `json:"sale_code"`
	SaleName            string  `json:"sale_name"`
	PosMachineId        int64   `json:"pos_machine_id"`
	PeriodId            int64   `json:"period_id"`
	CashId              int64   `json:"cash_id"`
	TaxType             int64   `json:"tax_type"`
	TaxRate             float64 `json:"tax_rate"`
	NumberOfItem        float64 `json:"number_of_item"`
	DepartId            string  `json:"depart_id"`
	AllocateId          int64   `json:"allocate_id"`
	ProjectId           int64   `json:"project_id"`
	PosStatus           int64   `json:"pos_status"`
	CreditDay           int64   `json:"credit_day"`
	DueDate             string  `json:"due_date"`
	DeliveryDay         int64   `json:"delivery_day"`
	DeliveryDate        string  `json:"delivery_date"`
	IsConfirm           int64   `json:"is_confirm"`
	IsConditionSend     int64   `json:"is_condition_send"`
	MyDescription       string  `json:"my_description"`
	SoRefNo             string  `json:"so_ref_no"`
	ChangeAmount        float64 `json:"change_amount"`
	SumCashAmount       float64 `json:"sum_cash_amount"`
	SumCreditAmount     float64 `json:"sum_credit_amount"`
	SumChqAmount        float64 `json:"sum_chq_amount"`
	SumBankAmount       float64 `json:"sum_bank_amount"`
	SumOfDeposit        float64 `json:"sum_of_deposit"`
	SumOnLineAmount     float64 `json:"sum_on_line_amount"`
	CouponAmount        float64 `json:"coupon_amount"`
	SumOfItemAmount     float64 `json:"sum_of_item_amount"`
	DiscountWord        string  `json:"discount_word"`
	DiscountAmount      float64 `json:"discount_amount"`
	AfterDiscountAmount float64 `json:"after_discount_amount"`
	BeforeTaxAmount     float64 `json:"before_tax_amount"`
	TaxAmount           float64 `json:"tax_amount"`
	TotalAmount         float64 `json:"total_amount"`
	NetDebtAmount       float64 `json:"net_debt_amount"`
	BillBalance         float64 `json:"bill_balance"`
	PayBillStatus       int64   `json:"pay_bill_status"`
	PayBillAmount       float64 `json:"pay_bill_amount"`
	DeliveryStatus      int64   `json:"delivery_status"`
	ReceiveName         string  `json:"receive_name"`
	ReceiveTel          string  `json:"receive_tel"`
	CarLicense          string  `json:"car_license"`
	IsCancel            int64   `json:"is_cancel"`
	IsHold              int64   `json:"is_hold"`
	IsPosted            int64   `json:"is_posted"`
	IsCreditNote        int64   `json:"is_credit_note"`
	IsDebitNote         int64   `json:"is_debit_note"`
	GlStatus            int64   `json:"gl_status"`
	JobId               string  `json:"job_id"`
	JobNo               string  `json:"job_no"`
	CouponNo            string  `json:"coupon_no"`
	RedeemNo            string  `json:"redeem_no"`
	ScgNumber           string  `json:"scg_number"`
	ScgId               string  `json:"scg_id"`
	CancelDescId        int64   `json:"cancel_desc_id"`
	CancelDesc          string  `json:"cancel_desc"`
	CreateBy            string  `json:"create_by"`
	CreateTime          string  `json:"create_time"`
	EditBy              string  `json:"edit_by"`
	EditTime            string  `json:"edit_time"`
	ConfirmBy           string  `json:"confirm_by"`
	ConfirmTime         string  `json:"confirm_time"`
	CancelBy            string  `json:"cancel_by"`
	CancelTime          string  `json:"cancel_time"`

	Subs []NewInvoiceItemTemplate `json:"subs"`
	//RecMoney            []RecMoneyTemplate       `json:"rec_money"`
	CreditCard []CreditCardTemplate `json:"credit_card"`
	Chq        []ChqInTemplate      `json:"chq"`
	BankPay    []BankpayTemplate    `json:"bank"`
}

type NewInvoiceItemTemplate struct {
	Id    int64 `json:"id"`
	InvId int64 `json:"inv_id"`

	ItemCode        string  `json:"item_code"`
	Itemid          int64   `json:"item_id"`
	ItemName        string  `json:"item_name"`
	BarCode         string  `json:"bar_code"`
	WhId            int64   `json:"wh_id"`
	ShelfId         int64   `json:"shelf_id"`
	Price           float64 `json:"price"`
	UnitCode        string  `json:"unit_code"`
	Location        string  `json:"location"`
	Qty             float64 `json:"qty"`
	CnQty           float64 `json:"cn_qty"`
	DiscountWord    string  `json:"discount_word_sub"`
	DiscountAmount  float64 `json:"discount_amount_sub"`
	ItemAmount      float64 `json:"amount"`
	NetAmount       float64 `json:"net_amount"`
	Average_cost    float64 `json:"average_cost"`
	SumOfCost       float32 `json:"sum_of_cost"`
	ItemDescription string  `json:"item_description"`
	IsCancel        int64   `json:"is_cancel"`
	IsCreditNote    int64   `json:"is_credit_note"`
	IsDebitNote     int64   `json:"is_debit_note"`
	PackingRate1    int64   `json:"packing_rate_1"`
	PackingRate2    int64   `json:"packing_rate_2"`
	RefNo           string  `json:"ref_no"`
	RefLineNumber   int64   `json:"ref_line_number"`
	LineNumber      int64   `json:"line_number"`
}

type NewSearchItemTemplate struct {
	Id              int64   `json:"id"`
	DocDate         string  `json:"doc_date"`
	DocNo           string  `json:"doc_no"`
	ItemId          int64   `json:"item_id"`
	ItemCode        string  `json:"item_code"`
	ItemName        string  `json:"item_name"`
	BarCode         string  `json:"bar_code"`
	UnitICode       string  `json:"unit_code"`
	WhId            int64   `json:"wh_id"`
	ShelfId         int64   `json:"shelf_id"`
	Price           float64 `json:"price"`
	Qty             float64 `json:"qty"`
	CnQty           float64 `json:"cn_qty"`
	DiscountWord    string  `json:"discount_word_sub"`
	ItemDescription string  `json:"item_description"`
	IsCreditNote    int64   `json:"is_credit_note"`
	IsDebitNote     int64   `json:"is_debit_note"`
	PackingRate1    int64   `json:"packing_rate_1"`
	PackingRate2    int64   `json:"packing_rate_2"`
	SoRefNo         string  `json:"so_ref_no"`
	AverageCost     float64 `json:"average_cost"`
	SumOfCost       float64 `json:"sum_of_cost"`
	RefLineNumber   int64   `json:"ref_line_number"`
	LineNumber      int64   `json:"line_number"`
	ArName          string  `json:"ar_name"`
	ArCode          string  `json:"ar_code"`
	ArId            int64   `json:"ar_id"`
	Name            string  `json:"name"`
	NId             int64   `json:"Id"`
	NDocNo          string  `json:"DocNo"`
	NDocDate        string  `json:"DocDate"`
	NItemId         int64   `json:"ItemId"`
	NArId           int64   `json:"ArId"`
	NBarCode        string  `json:"BarCode"`
	NItemCode       string  `json:"ItemCode"`
	NItemName       string  `json:"ItemName"`
	NUnitCode       string  `json:"UnitCode"`
	NQty            float64 `json:"Qty"`
	NPrice          float64 `json:"Price"`
	NArName         string  `json:"ArName"`
	NDiscountWord   string  `json:"DiscountWord"`
}

type SearchHisCustomerTemplate struct {
	ArId     string `json:"ar_id"`
	ArCode   string `json:"ar_code"`
	ItemCode string `json:"item_code"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Page     string `json:"page"`
}

type NewSearchHisCustomerTemplate struct {
	Id           int64   `json:"id"`
	DocDate      string  `json:"doc_date"`
	DocNo        string  `json:"doc_no"`
	ArName       string  `json:"ar_name"`
	ArCode       string  `json:"ar_code"`
	ArId         int64   `json:"ar_id"`
	SaleName     string  `json:"sale_name"`
	TotalAmount  float64 `json:"total_amount"`
	NId          int64   `json:"Id"`
	NDocNo       string  `json:"DocNo"`
	NDocDate     string  `json:"DocDate"`
	NArId        int64   `json:"ArId"`
	NArName      string  `json:"ArName"`
	NSaleName    string  `json:"SaleName"`
	NTotalAmount int64   `json:"TotalAmount"`
}
