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
	SaleId              int                  `json:"sale_id"`
	SaleCode            string               `json:"sale_code"`
	SaleName            string               `json:"sale_name"`
	BillType            int64                `json:"bill_type"`
	TaxType             int                  `json:"tax_type"`
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
	SaleId              int                   `json:"sale_id"`
	SaleCode            string                `json:"sale_code"`
	SaleName            string                `json:"sale_name"`
	BillType            int64                 `json:"bill_type"`
	TaxType             int                   `json:"tax_type"`
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

type SearchByKeywordTemplate struct {
	SaleCode string `json:"sale_code"`
	Keyword  string `json:"keyword"`
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

type NewDepositTemplate struct {
	Id               int64   `json:"id"`
	CompanyId        int64   `json:"company_id"`
	BranchId         int64   `json:"branch_id"`
	Uuid             string  `json:"uuid"`
	DocNo            string  `json:"doc_no"`
	TaxNo            string  `json:"tax_no"`
	DocDate          string  `json:"doc_date"`
	BillType         int64   `json:"bill_type"`
	ArId             int64   `json:"ar_id"`
	ArCode           string  `json:"ar_code"`
	SaleId           int64   `json:"sale_id"`
	SaleCode         string  `json:"sale_code"`
	TaxType          int64   `json:"tax_type"`
	TaxRate          float64 `json:"tax_rate"`
	RefNo            string  `json:"ref_no"`
	CreditDay        int64   `json:"credit_day"`
	DueDate          string  `json:"due_date"`
	DepartId         int64   `json:"depart_id"`
	AllocateId       int64   `json:"allocate_id"`
	ProjectId        int64   `json:"project_id"`
	MyDescription    string  `json:"my_description"`
	BeforeTaxAmount  float64 `json:"before_tax_amount"`
	TaxAmount        float64 `json:"tax_amount"`
	TotalAmount      float64 `json:"total_amount"`
	NetAmount        float64 `json:"net_amount"`
	BillBalance      float64 `json:"bill_balance"`
	CashAmount       float64 `json:"cash_amount"`
	CreditcardAmount float64 `json:"creditcard_amount"`
	ChqAmount        float64 `json:"chq_amount"`
	BankAmount       float64 `json:"bank_amount"`
	IsReturnMoney    int64   `json:"is_return_money" `
	IsCancel         int64   `json:"is_cancel"`
	IsConfirm        int64   `json:"is_confirm"`
	ScgId            string  `json:"scg_id"`
	JobNo            string  `json:"job_no"`
	CreateBy         string  `json:"create_by"`
	CreateTime       string  `json:"create_time"`
	EditBy           string  `json:"edit_by"`
	EditTime         string  `json:"edit_time"`
	CancelBy         string  `json:"cancel_by"`
	CancelTime       string  `json:"cancel_time" `
	ConfirmBy        string  `json:"confirm_by"`
	ConfirmTime      string  `json:"confirm_time"`
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
