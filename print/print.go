package print

type PosSlipRequestTemplate struct {
	DocNo string `json:"doc_no"`
}

type PosDriveThruSlipRequestTemplate struct {
	//DbHost string `json:"db_host"`
	//DbName string `json:"db_name"`
	//DbUser string `json:"db_user"`
	//DbPass string `json:"db_pass"`
	//HostIP string `json:"host_ip"`

	AccessToken string `json:"access_token"`
	FormId      string `json:"form_id"`
	PosId       string `json:"pos_id"`
	DocNo       string `json:"doc_no"`
}

type PosSlipResponseTemplate struct {
	DocNo string `json:"doc_no"`
}

type CompanyDataTemplate struct {
	CompanyId        int64  `json:"company_id"`
	CompanyName      string `json:"company_name"`
	CompanyAddress   string `json:"company_address"`
	CompanyTaxNo     string `json:"company_tax_no"`
	CompanyTelephone string `json:"company_telephone"`
	CompanyFax       string `json:"company_fax"`
	BranchId         int64  `json:"branch_id"`
	BranchName       string `json:"branch_name"`
	BranchAddress    string `json:"branch_address"`
	BranchTaxNo      string `json:"branch_tax_no"`
	BranchTelephone  string `json:"branch_telephone"`
	BranchFax        string `json:"branch_fax"`
}

type CustomerDataTemplate struct {
	ArCode        string `json:"ar_code"`
	ArName        string `json:"ar_name"`
	ArBillAddress string `json:"ArBillAddress"`
	ArTelephone   string `json:"ArTelephone"`
	ArFax         string `json:"ar_fax"`
	ArTaxNo       string `json:"ar_tax_no"`
}

type SaleDataTemplate struct {
	SaleCode string `json:"sale_code"`
	SaleName string `json:"sale_name"`
}

type EnvironmentTemplate struct {
	DepartCode   string `json:"depart_code"`
	DepartName   string `json:"depart_name"`
	AllocateCode string `json:"allocate_code"`
	AllocateName string `json:"allocate_name"`
	ProjectCode  string `json:"project_code"`
	ProjectName  string `json:"project_name"`
}

type InvoiceTemplate struct {
	company             CompanyData             `json:"company"`
	customer            CustomerData            `json:"customer"`
	sale                SaleData                `json:"sale"`
	Id                  int64                   `json:"id"`
	DocNo               string                  `json:"doc_no"`
	TaxNo               string                  `json:"tax_no"`
	BillType            int64                   `json:"bill_type"`
	DocDate             string                  `json:"doc_date"`
	PosMachineCode      string                  `json:"pos_machine_code"`
	CashierCode         int64                   `json:"cashier_code"`
	TaxType             int64                   `json:"tax_type"`
	TaxRate             float64                 `json:"tax_rate"`
	NumberOfItem        float64                 `json:"number_of_item"`
	PosStatus           int64                   `json:"pos_status"`
	CreditDay           int64                   `json:"credit_day"`
	DueDate             string                  `json:"due_date"`
	DeliveryDay         int64                   `json:"delivery_day"`
	DeliveryDate        string                  `json:"delivery_date"`
	IsConfirm           int64                   `json:"is_confirm"`
	IsConditionSend     int64                   `json:"is_condition_send"`
	MyDescription       string                  `json:"my_description"`
	SoRefNo             string                  `json:"so_ref_no"`
	ChangeAmount        float64                 `json:"change_amount"`
	SumCashAmount       float64                 `json:"sum_cash_amount"`
	SumCreditAmount     float64                 `json:"sum_credit_amount"`
	SumChqAmount        float64                 `json:"sum_chq_amount"`
	SumBankAmount       float64                 `json:"sum_bank_amount"`
	SumOfDeposit        float64                 `json:"sum_of_deposit"`
	SumOnLineAmount     float64                 `json:"sum_on_line_amount"`
	CouponAmount        float64                 `json:"coupon_amount"`
	SumOfItemAmount     float64                 `json:"sum_of_item_amount"`
	DiscountWord        string                  `json:"discount_word"`
	DiscountAmount      float64                 `json:"discount_amount"`
	AfterDiscountAmount float64                 `json:"after_discount_amount"`
	BeforeTaxAmount     float64                 `json:"before_tax_amount"`
	TaxAmount           float64                 `json:"tax_amount"`
	TotalAmount         float64                 `json:"total_amount"`
	NetDebtAmount       float64                 `json:"net_debt_amount"`
	BillBalance         float64                 `json:"bill_balance"`
	PayBillStatus       int64                   `json:"pay_bill_status"`
	PayBillAmount       float64                 `json:"pay_bill_amount"`
	DeliveryStatus      int64                   `json:"delivery_status"`
	ReceiveName         string                  `json:"receive_name"`
	ReceiveTel          string                  `json:"receive_tel"`
	CarLicense          string                  `json:"car_license"`
	IsCancel            int64                   `json:"is_cancel"`
	IsHold              int64                   `json:"is_hold"`
	JobId               string                  `json:"job_id"`
	JobNo               string                  `json:"job_no"`
	CouponNo            string                  `json:"coupon_no"`
	RedeemNo            string                  `json:"redeem_no"`
	ScgNumber           string                  `json:"scg_number"`
	ScgId               string                  `json:"scg_id"`
	CreateBy            string                  `json:"create_by"`
	CreateTime          string                  `json:"create_time"`
	EditBy              string                  `json:"edit_by"`
	EditTime            string                  `json:"edit_time"`
	ConfirmBy           string                  `json:"confirm_by"`
	ConfirmTime         string                  `json:"confirm_time"`
	CancelBy            string                  `json:"cancel_by"`
	CancelTime          string                  `json:"cancel_time"`
	CancelDescId        int64                   `json:"cancel_desc_id"`
	CancelDesc          string                  `json:"cancel_desc"`
	Subs                []NewInvoiceItemRequest `json:"subs"`
	RecMoney            []RecMoney              `json:"rec_money"`
	CreditCard          []CreditCard            `json:"credit_card"`
	Chq                 []ChqIn                 `json:"chq"`
}
