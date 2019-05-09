package mysqldb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/config"
	"github.com/mrtomyum/nopadol/gendocno"
	"github.com/mrtomyum/nopadol/sales"
)

type NewQuoModel struct {
	//Id                  int64             `db:"id"`
	//DocNo               string            `db:"doc_no"`
	//DocDate             string            `db:"doc_date"`
	//CompanyId           int64             `db:"company_id"`
	//BranchId            int64             `db:"branch_id"`
	//DocType             int64             `db:"doc_type"`
	//ArId                int64             `db:"ar_id"`
	//ArCode              string            `db:"ar_code"`
	//ArName              string            `db:"ar_name"`
	//ArBillAddress       string            `db:"ar_bill_address"`
	//ArTelephone         string            `db:"ar_telephone"`
	//SaleId              int64             `db:"sale_id"`
	//SaleCode            string            `db:"sale_code"`
	//SaleName            string            `db:"sale_name"`
	//BillType            int64             `db:"bill_type"`
	//TaxType             int64             `db:"tax_type"`
	//TaxRate             float64           `db:"tax_rate"`
	//DepartId            int64             `db:"depart_id"`
	//RefNo               string            `db:"ref_no"`
	//JobId               string            `db:"job_id"`
	//IsConfirm           int64             `db:"is_confirm"`
	//BillStatus          int64             `db:"bill_status"`
	//Validity            int64             `db:"validity"`
	//CreditDay           int64             `db:"credit_day"`
	//DueDate             string            `db:"due_date"`
	//ExpireCredit        int64             `db:"expire_credit"`
	//ExpireDate          string            `db:"expire_date"`
	//DeliveryDay         int64             `db:"delivery_day"`
	//DeliveryDate        string            `db:"delivery_date"`
	//AssertStatus        int64             `db:"assert_status"`
	//IsConditionSend     int64             `db:"is_condition_send"`
	//MyDescription       string            `db:"my_description"`
	//SumOfItemAmount     float64           `db:"sum_of_item_amount"`
	//DiscountWord        string            `db:"discount_word"`
	//DiscountAmount      float64           `db:"discount_amount"`
	//AfterDiscountAmount float64           `db:"after_discount_amount"`
	//BeforeTaxAmount     float64           `db:"before_tax_amount"`
	//TaxAmount           float64           `db:"tax_amount"`
	//TotalAmount         float64           `db:"total_amount"`
	//NetDebtAmount       float64           `db:"net_debt_amount"`
	//ProjectId           int64             `db:"project_id"`
	//AllocateId          int64             `db:"allocate_id"`
	//IsCancel            int64             `db:"is_cancel"`
	//CreateBy            string            `db:"create_by"`
	//CreateTime          string            `db:"create_time"`
	//EditBy              string            `db:"edit_by"`
	//EditTime            string            `db:"edit_time"`
	//ConfirmBy           string            `db:"confirm_by"`
	//ConfirmTime         string            `db:"confirm_time"`
	//CancelBy            string            `db:"cancel_by"`
	//CancelTime          string            `db:"cancel_time"`
	//Subs                []NewQuoItemModel `db:"subs"`

	Id                  int64             `db:"Id"`
	DocNo               string            `db:"DocNo"`
	DocDate             string            `db:"DocDate"`
	CompanyId           int64             `db:"CompanyId"`
	BranchId            int64             `db:"BranchId"`
	DocType             int64             `db:"DocType"`
	ArId                int64             `db:"ArId"`
	ArCode              string            `db:"ArCode"`
	ArName              string            `db:"ArName"`
	ArBillAddress       string            `db:"ArBillAddress"`
	ArTelephone         string            `db:"ArTelephone"`
	SaleId              int64             `db:"SaleId"`
	SaleCode            string            `db:"SaleCode"`
	SaleName            string            `db:"SaleName"`
	BillType            int64             `db:"BillType"`
	TaxType             int64             `db:"TaxType"`
	TaxRate             float64           `db:"TaxRate"`
	DepartId            int64             `db:"DepartId"`
	RefNo               string            `db:"RefNo"`
	JobId               string            `db:"JobId"`
	IsConfirm           int64             `db:"IsConfirm"`
	BillStatus          int64             `db:"BillStatus"`
	Validity            int64             `db:"Validity"`
	CreditDay           int64             `db:"CreditDay"`
	DueDate             string            `db:"DueDate"`
	ExpireCredit        int64             `db:"ExpireCredit"`
	ExpireDate          string            `db:"ExpireDate"`
	DeliveryDay         int64             `db:"DeliveryDay"`
	DeliveryDate        string            `db:"DeliveryDate"`
	AssertStatus        int64             `db:"AssertStatus"`
	IsConditionSend     int64             `db:"IsConditionSend"`
	MyDescription       string            `db:"MyDescription"`
	SumOfItemAmount     float64           `db:"SumOfItemAmount"`
	DiscountWord        string            `db:"DiscountWord"`
	DiscountAmount      float64           `db:"DiscountAmount"`
	AfterDiscountAmount float64           `db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64           `db:"BeforeTaxAmount"`
	TaxAmount           float64           `db:"TaxAmount"`
	TotalAmount         float64           `db:"TotalAmount"`
	NetDebtAmount       float64           `db:"NetDebtAmount"`
	ProjectId           int64             `db:"ProjectId"`
	AllocateId          int64             `db:"AllocateId"`
	IsCancel            int64             `db:"IsCancel"`
	CreateBy            string            `db:"CreateBy"`
	CreateTime          string            `db:"CreateTime"`
	EditBy              string            `db:"EditBy"`
	EditTime            string            `db:"EditTime"`
	ConfirmBy           string            `db:"ConfirmBy"`
	ConfirmTime         string            `db:"ConfirmTime"`
	CancelBy            string            `db:"CancelBy"`
	CancelTime          string            `db:"CancelTime"`
	ContactId           int64             `db:"ContactId"`
	Subs                []NewQuoItemModel `db:"subs"`
}

type NewQuoItemModel struct {
	//Id              int64   `db:"id"`
	//QuoId           int64   `db:"quo_id"`
	//ItemId          int64   `db:"item_id"`
	//ItemCode        string  `db:"item_code"`
	//BarCode         string  `db:"bar_code"`
	//ItemName        string  `db:"item_name"`
	//Qty             float64 `db:"qty"`
	//RemainQty       float64 `db:"remain_qty"`
	//Price           float64 `db:"price"`
	//DiscountWord    string  `db:"discount_word"`
	//DiscountAmount  float64 `db:"discount_amount"`
	//UnitCode        string  `db:"unit_code"`
	//ItemAmount      float64 `db:"item_amount"`
	//ItemDescription string  `db:"item_description"`
	//PackingRate1    float64 `db:"packing_rate1"`
	//IsCancel        int64   `db:"is_cancel"`
	//LineNumber      int     `db:"line_number"`
	Id              int64   `db:"Id"`
	QuoId           int64   `db:"QuoId"`
	ItemId          int64   `db:"ItemId"`
	ItemCode        string  `db:"ItemCode"`
	BarCode         string  `db:"BarCode"`
	ItemName        string  `db:"ItemName"`
	Qty             float64 `db:"Qty"`
	RemainQty       float64 `db:"RemainQty"`
	Price           float64 `db:"Price"`
	DiscountWord    string  `db:"DiscountWord"`
	DiscountAmount  float64 `db:"DiscountAmount"`
	UnitCode        string  `db:"UnitCode"`
	ItemAmount      float64 `db:"ItemAmount"`
	ItemDescription string  `db:"ItemDescription"`
	PackingRate1    float64 `db:"PackingRate1"`
	IsCancel        int64   `db:"IsCancel"`
	LineNumber      int     `db:"LineNumber"`
	WHCode          string  `db:"WhCode"`
	ShelfCode       string  `db:"ShelfCode"`
}

type NewSaleModel struct {
	Id                  int64              `db:"Id"`
	DocNo               string             `db:"DocNo"`
	DocDate             string             `db:"DocDate"`
	CompanyId           int64              `db:"CompanyId"`
	BranchId            int64              `db:"BranchId"`
	DocType             int64              `db:"DocType"`
	ArId                int64              `db:"ArId"`
	ArCode              string             `db:"ArCode"`
	ArName              string             `db:"ArName"`
	ArBillAddress       string             `db:"ArBillAddress"`
	ArTelephone         string             `db:"ArTelephone"`
	SaleId              int64              `db:"SaleId"`
	SaleCode            string             `db:"SaleCode"`
	SaleName            string             `db:"SaleName"`
	BillType            int64              `db:"BillType"`
	TaxType             int64              `db:"TaxType"`
	TaxRate             float64            `db:"TaxRate"`
	DepartId            int64              `db:"DepartId"`
	RefNo               string             `db:"RefNo"`
	IsConfirm           int64              `db:"IsConfirm"`
	BillStatus          int64              `db:"BillStatus"`
	HoldingStatus       int64              `db:"HoldingStatus"`
	CreditDay           int64              `db:"CreditDay"`
	DueDate             string             `db:"DueDate"`
	DeliveryDay         int64              `db:"DeliveryDay"`
	DeliveryDate        string             `db:"DeliveryDate"`
	IsConditionSend     int64              `db:"IsConditionSend"`
	DeliveryAddressId   int64              `db:"DeliveryAddressId"`
	CarLicense          string             `db:"CarLicense"`
	PersonReceiveTel    string             `db:"PersonReceiveTel"`
	MyDescription       string             `db:"MyDescription"`
	SumOfItemAmount     float64            `db:"SumOfItemAmount"`
	DiscountWord        string             `db:"DiscountWord"`
	DiscountAmount      float64            `db:"DiscountAmount"`
	AfterDiscountAmount float64            `db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64            `db:"BeforeTaxAmount"`
	TaxAmount           float64            `db:"TaxAmount"`
	TotalAmount         float64            `db:"TotalAmount"`
	NetDebtAmount       float64            `db:"NetDebtAmount"`
	ProjectId           int64              `db:"ProjectId"`
	AllocateId          int64              `db:"AllocateId"`
	JobId               string             `db:"JobId"`
	IsCancel            int64              `db:"IsCancel"`
	CreateBy            string             `db:"CreateBy"`
	CreateTime          string             `db:"CreateTime"`
	EditBy              string             `db:"EditBy"`
	EditTime            string             `db:"EditTime"`
	ConfirmBy           string             `db:"ConfirmBy"`
	ConfirmTime         string             `db:"ConfirmTime"`
	CancelBy            string             `db:"CancelBy"`
	CancelTime          string             `db:"CancelTime"`
	ContactId           int64              `db:"ContactId"`
	Subs                []NewSaleItemModel `db:"subs"`
}

type NewSaleItemModel struct {
	Id              int64   `db:"Id"`
	SOId            int64   `db:"SOId"`
	ItemId          int64   `db:"ItemId"`
	ItemCode        string  `db:"ItemCode"`
	BarCode         string  `db:"BarCode"`
	WHCode          string  `db:"WHCode"`
	ShelfCode       string  `db:"ShelfCode"`
	ItemName        string  `db:"ItemName"`
	Qty             float64 `db:"Qty"`
	RemainQty       float64 `db:"RemainQty"`
	Price           float64 `db:"Price"`
	DiscountWord    string  `db:"DiscountWord"`
	DiscountAmount  float64 `db:"DiscountAmount"`
	UnitCode        string  `db:"UnitCode"`
	ItemAmount      float64 `db:"ItemAmount"`
	ItemDescription string  `db:"ItemDescription"`
	StockType       int64   `db:"StockType"`
	AverageCost     float64 `db:"AverageCost"`
	SumOfCost       float64 `db:"SumOfCost"`
	PackingRate1    float64 `db:"PackingRate1"`
	RefNo           string  `db:"RefNo"`
	QuoId           int64   `db:"QuoId"`
	LineNumber      int     `db:"LineNumber"`
	RefLineNUmber   int     `db:"RefLineNUmber"`
	IsCancel        int64   `db:"IsCancel"`
}

type SearchDocModel struct {
	Id            int64   `db:"Id"`
	DocNo         string  `db:"DocNo"`
	DocDate       string  `db:"DocDate"`
	Module        string  `db:"Module"`
	ArCode        string  `db:"ArCode"`
	ArName        string  `db:"ArName"`
	SaleCode      string  `db:"SaleCode"`
	SaleName      string  `db:"SaleName"`
	MyDescription string  `db:"MyDescription"`
	TotalAmount   float64 `db:"TotalAmount"`
	IsCancel      int     `db:"IsCancel"`
	IsConfirm     int     `db:"IsConfirm"`
}

type SearchDocDetailsModel struct {
	Id                  int64             `db:"Id"`
	DocNo               string            `db:"DocNo"`
	DocDate             string            `db:"DocDate"`
	DocType             int64             `db:"DocType"`
	ArId                int64             `db:"ArId"`
	ArCode              string            `db:"ArCode"`
	ArName              string            `db:"ArName"`
	SaleId              int64             `db:"SaleId"`
	SaleCode            string            `db:"SaleCode"`
	SaleName            string            `db:"SaleName"`
	BillType            int64             `db:"BillType"`
	TaxType             int64             `db:"TaxType"`
	TaxRate             float64           `db:"TaxRate"`
	DepartCode          string            `db:"DepartCode"`
	RefNo               string            `db:"RefNo"`
	IsConfirm           int64             `db:"IsConfirm"`
	BillStatus          int64             `db:"BillStatus"`
	CreditDay           int64             `db:"CreditDay"`
	DueDate             string            `db:"DueDate"`
	ExpireDate          string            `db:"ExpireDate"`
	DeliveryDate        string            `db:"DeliveryDate"`
	AssertStatus        int64             `db:"AssertStatus"`
	IsConditionSend     int64             `db:"IsConditionSend"`
	MyDescription       string            `db:"MyDescription"`
	SumOfItemAmount     float64           `db:"SumOfItemAmount"`
	DiscountWord        string            `db:"DiscountWord"`
	DiscountAmount      float64           `db:"DiscountAmount"`
	AfterDiscountAmount float64           `db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64           `db:"BeforeTaxAmount"`
	TaxAmount           float64           `db:"TaxAmount"`
	TotalAmount         float64           `db:"TotalAmount"`
	NetDebtAmount       float64           `db:"NetDebtAmount"`
	ProjectId           int64             `db:"ProjectId"`
	IsCancel            int64             `db:"IsCancel"`
	CreateBy            string            `db:"CreateBy"`
	CreateTime          string            `db:"CreateTime"`
	EditBy              string            `db:"EditBy"`
	EditTime            string            `db:"EditTime"`
	CancelBy            string            `db:"CancelBy"`
	CancelTime          string            `db:"CancelTime"`
	SoStatus            int64             `db:"SoStatus"`
	HoldingStatus       int64             `db:"HoldingStatus"`
	AllocateId          int64             `db:"AllocateId"`
	JobId               string            `db:"JobId"`
	ConfirmBy           string            `db:"ConfirmBy"`
	ConfirmTime         string            `db:"ConfirmTime"`
	Subs                []NewQuoItemModel `db:"subs"`
}

type NewDepositModel struct {
	Id               int64                 `db:"id"`
	CompanyId        int64                 `db:"company_id"`
	BranchId         int64                 `db:"branch_id"`
	Uuid             string                `db:"uuid"`
	DocNo            string                `db:"doc_no"`
	TaxNo            string                `db:"tax_no"`
	DocDate          string                `db:"doc_date"`
	BillType         int64                 `db:"bill_type"`
	ArId             int64                 `db:"ar_id"`
	ArCode           string                `db:"ar_code"`
	ArName           string                `db:"ar_name"`
	ArBillAddress    string                `db:"ar_bill_address"`
	ArTelephone      string                `db:"ar_telephone"`
	SaleId           int64                 `db:"sale_id"`
	SaleCode         string                `db:"sale_code"`
	SaleName         string                `db:"sale_name"`
	TaxType          int64                 `db:"tax_type"`
	TaxRate          float64               `db:"tax_rate"`
	RefNo            string                `db:"ref_no"`
	CreditDay        int64                 `db:"credit_day"`
	DueDate          string                `db:"due_date"`
	DepartId         int64                 `db:"depart_id"`
	AllocateId       int64                 `db:"allocate_id"`
	ProjectId        int64                 `db:"project_id"`
	MyDescription    string                `db:"my_description"`
	BeforeTaxAmount  float64               `db:"before_tax_amount"`
	TaxAmount        float64               `db:"tax_amount"`
	TotalAmount      float64               `db:"total_amount"`
	NetAmount        float64               `db:"net_amount"`
	BillBalance      float64               `db:"bill_balance"`
	CashAmount       float64               `db:"cash_amount"`
	CreditcardAmount float64               `db:"creditcard_amount"`
	ChqAmount        float64               `db:"chq_amount"`
	BankAmount       float64               `db:"bank_amount"`
	IsReturnMoney    int64                 `db:"is_return_money" `
	IsCancel         int64                 `db:"is_cancel"`
	IsConfirm        int64                 `db:"is_confirm"`
	ScgId            string                `db:"scg_id"`
	JobNo            string                `db:"job_no"`
	CreateBy         string                `db:"create_by"`
	CreateTime       string                `db:"create_time"`
	EditBy           string                `db:"edit_by"`
	EditTime         string                `db:"edit_time"`
	CancelBy         string                `db:"cancel_by"`
	CancelTime       string                `db:"cancel_time" `
	ConfirmBy        string                `db:"confirm_by"`
	ConfirmTime      string                `db:"confirm_time"`
	Subs             []NewDepositItemModel `db:"subs"`
	RecMoney         []RecMoneyModel       `db:"rec_money"`
	CreditCard       []CreditCardModel     `db:"credit_card"`
	Chq              []ChqInModel          `db:"chq"`
}

type NewDepositItemModel struct {
	Id              int64   `db:"id"`
	SORefNo         string  `db:"so_ref_no"`
	SOId            int64   `db:"so_id"`
	ItemId          int64   `db:"item_id"`
	ItemCode        string  `db:"item_code"`
	BarCode         string  `db:"bar_code"`
	ItemName        string  `db:"item_name"`
	WHCode          string  `db:"wh_code"`
	ShelfCode       string  `db:"shelf_code"`
	Qty             float64 `db:"qty"`
	RemainQty       float64 `db:"remain_qty"`
	Price           float64 `db:"price"`
	DiscountWord    string  `db:"discount_word"`
	DiscountAmount  float64 `db:"discount_amount"`
	UnitCode        string  `db:"unit_code"`
	ItemAmount      float64 `db:"item_amount"`
	ItemDescription string  `db:"item_description"`
	PackingRate1    float64 `db:"packing_rate_1"`
	RefNo           string  `db:"ref_no"`
	QuoId           int64   `db:"quo_id"`
	LineNumber      int     `db:"line_number"`
	RefLineNumber   int64   `db:"ref_line_number"`
	IsCancel        int64   `db:"is_cancel"`
}

type CreditCardModel struct {
	Id           int64   `db:"id"`
	RefId        int64   `db:"ref_id"`
	CreditCardNo string  `db:"credit_card_no"`
	CreditType   string  `db:"credit_type"`
	ConfirmNo    string  `db:"confirm_no"`
	Amount       float64 `db:"amount"`
	ChargeAmount float64 `db:"charge_amount"`
	Description  string  `db:"description"`
	BankId       int64   `db:"bank_id"`
	BankBranchId int64   `db:"bank_branch_id"`
	ReceiveDate  string  `db:"receive_date"`
	DueDate      string  `db:"due_date"`
	BookId       int64   `db:"book_id"`
}

type ChqInModel struct {
	Id           int64   `db:"id"`
	RefId        int64   `db:"ref_id"`
	ChqNumber    string  `db:"chq_number"`
	BankId       int64   `db:"bank_id"`
	BankBranchId int64   `db:"bank_branch_id"`
	ReceiveDate  string  `db:"receive_date"`
	DueDate      string  `db:"due_date"`
	BookId       int64   `db:"book_id"`
	ChqStatus    int64   `db:"chq_status"`
	ChqAmount    float64 `db:"chq_amount"`
	ChqBalance   float64 `db:"chq_balance"`
	Description  string  `db:"description"`
}

type RecMoneyModel struct {
	Id             int64   `db:"id"`
	DocType        int64   `db:"doc_type"`
	RefId          int64   `db:"ref_id"`
	ArId           int64   `db:"ar_id"`
	PaymentType    int64   `db:"payment_type"`
	PayAmount      float64 `db:"pay_amount"`
	ChqTotalAmount float64 `db:"chq_total_amount"`
	CreditType     int64   `db:"credit_type"`
	ChargeAmount   float64 `db:"charge_amount"`
	ConfirmNo      string  `db:"confirm_no"`
	RefNo          string  `db:"ref_no"`
	BankCode       string  `db:"bank_code"`
	BankBranchCode string  `db:"bank_branch_code"`
	RefDate        string  `db:"ref_date"`
	BankTransDate  string  `db:"bank_trans_date"`
	LineNumber     int64   `db:"line_number"`
}
type CreditCardTypeModel struct {
	Id                 int64  `db:"id"`
	CreditcardTypeName string `db:"creditcardtype_name"`
}
type BankpayModel struct {
	Id           int64   `db:"id"`
	RefId        int64   `db:"ref_id"`
	BankAccount  string  `db:"bank_account"`
	BankName     string  `db:"bank_name"`
	BankAmount   float64 `db:"bank_amount"`
	Activestatus int64   `db:"active_status"`
	CreateBy     string  `db:"create_by"`
	CreateTime   string  `db:"create_time"`
	EditBy       string  `db:"edit_by"`
}

type NewInvoiceModel struct {
	Id                  int64                 `db:"id"`
	CompanyId           int64                 `db:"company_id"`
	BranchId            int64                 `db:"branch_id"`
	Uuid                string                `db:"uuid"`
	DocNo               string                `db:"doc_no"`
	TaxNo               string                `db:"tax_no"`
	BillType            int64                 `db:"bill_type"`
	DocDate             string                `db:"doc_date"`
	DocType             int64                 `db:"doc_type"`
	ArId                int64                 `db:"ar_id"`
	ArCode              string                `db:"ar_code"`
	Code                string                `db:"code"`
	ArName              string                `db:"ar_name"`
	ItemName            string                `db:"item_name"`
	ArBillAddress       string                `db:"ar_bill_address"`
	ArTelephone         string                `db:"ar_telephone"`
	SaleId              int64                 `db:"sale_id"`
	SaleCode            string                `db:"sale_code"`
	SaleName            string                `db:"sale_name"`
	PosMachineId        int64                 `db:"pos_machine_id"`
	shiftuid            string                `db:"shift_uuid"`
	PeriodId            int64                 `db:"period_id"`
	CashId              int64                 `db:"cash_id"`
	TaxType             int64                 `db:"tax_type"`
	TaxRate             float64               `db:"tax_rate"`
	NumberOfItem        float64               `db:"number_of_item"`
	DepartId            string                `db:"depart_id"`
	AllocateId          int64                 `db:"allocate_id"`
	ProjectId           int64                 `db:"project_id"`
	PosStatus           int64                 `db:"pos_status"`
	CreditDay           int64                 `db:"credit_day"`
	DueDate             string                `db:"due_date"`
	DeliveryDay         int64                 `db:"delivery_day"`
	DeliveryDate        string                `db:"delivery_date"`
	IsConfirm           int64                 `db:"is_confirm"`
	IsConditionSend     int64                 `db:"is_condition_send"`
	MyDescription       string                `db:"my_description"`
	SoRefNo             string                `db:"so_ref_no"`
	ChangeAmount        float64               `db:"change_amount"`
	SumCashAmount       float64               `db:"sum_cash_amount"`
	SumCreditAmount     float64               `db:"sum_credit_amount"`
	SumChqAmount        float64               `db:"sum_chq_amount"`
	SumBankAmount       float64               `db:"sum_bank_amount"`
	SumOfDeposit        float64               `db:"sum_of_deposit"`
	SumOnLineAmount     float64               `db:"sum_on_line_amount"`
	CouponAmount        float64               `db:"coupon_amount"`
	SumOfItemAmount     float64               `db:"sum_of_item_amount"`
	DiscountWord        string                `db:"discount_word"`
	DiscountAmount      float64               `db:"discount_amount"`
	AfterDiscountAmount float64               `db:"after_discount_amount"`
	BeforeTaxAmount     float64               `db:"before_tax_amount"`
	TaxAmount           float64               `db:"tax_amount"`
	TotalAmount         float64               `db:"total_amount"`
	NetDebtAmount       float64               `db:"net_debt_amount"`
	BillBalance         float64               `db:"bill_balance"`
	PayBillStatus       int64                 `db:"pay_bill_status"`
	PayBillAmount       float64               `db:"pay_bill_amount"`
	DeliveryStatus      int64                 `db:"delivery_status"`
	ReceiveName         string                `db:"receive_name"`
	ReceiveTel          string                `db:"receive_tel"`
	CarLicense          string                `db:"car_license"`
	IsCancel            int64                 `db:"is_cancel"`
	IsHold              int64                 `db:"is_hold"`
	IsPosted            int64                 `db:"is_posted"`
	IsCreditNote        int64                 `db:"is_credit_note"`
	IsDebitNote         int64                 `db:"is_debit_note"`
	GlStatus            int64                 `db:"gl_status"`
	JobId               string                `db:"job_id"`
	JobNo               string                `db:"job_no"`
	CouponNo            string                `db:"coupon_no"`
	RedeemNo            string                `db:"redeem_no"`
	ScgNumber           string                `db:"scg_number"`
	ScgId               string                `db:"scg_id"`
	CreateBy            string                `db:"create_by"`
	CreateTime          string                `db:"create_time"`
	EditBy              string                `db:"edit_by"`
	EditTime            string                `db:"edit_time"`
	ConfirmBy           string                `db:"confirm_by"`
	ConfirmTime         string                `db:"confirm_time"`
	CancelBy            string                `db:"cancel_by"`
	CancelTime          string                `db:"cancel_time"`
	CancelDescId        int64                 `db:"cancel_desc_id"`
	CancelDesc          string                `db:"cancel_desc"`
	ItemCode            string                `db:"item_code"`
	Subs                []NewInvoiceItemModel `db:"subs"`
	RecMoney            []RecMoneyModel       `db:"rec_money"`
	CreditCard          []CreditCardModel     `db:"credit_card"`
	Chq                 []ChqInModel          `db:"chq"`
	Bank                []BankpayModel        `db:"bank"`
}

type NewInvoiceItemModel struct {
	Id              int64   `db:"id"`
	InvId           int64   `db:"inv_id"`
	Itemid          int64   `db:"item_id"`
	ItemCode        string  `db:"item_code"`
	ItemName        string  `db:"item_name"`
	BarCode         string  `db:"bar_code"`
	WhId            int64   `db:"wh_id"`
	ShelfId         int64   `db:"shelf_id"`
	Price           float64 `db:"price"`
	UnitCode        string  `db:"unit_code"`
	Qty             float64 `db:"qty"`
	CnQty           float64 `db:"cn_qty"`
	DiscountWord    string  `db:"discount_word_sub"`
	DiscountAmount  float64 `db:"discount_amount_sub"`
	ItemAmount      float64 `db:"amount"`
	NetAmount       float64 `db:"net_amount"`
	Average_cost    float64 `db:"average_cost"`
	SumOfCost       float32 `db:"sum_of_cost"`
	ItemDescription string  `db:"item_decription"`
	IsCancel        int64   `db:"is_cancel"`
	IsCreditNote    int64   `db:"is_credit_note"`
	IsDebitNote     int64   `db:"is_debit_note"`
	PackingRate1    int64   `db:"packing_rate_1"`
	PackingRate2    int64   `db:"packing_rate_2"`
	RefNo           string  `db:"ref_no"`
	RefLineNumber   int64   `db:"ref_line_number"`
	LineNumber      int64   `db:"line_number"`
}
type NewSearchItemModel struct {
	Id              int64   `db:"id"`
	DocDate         string  `db:"doc_date"`
	DocNo           string  `db:"doc_no"`
	ItemId          int64   `db:"item_id"`
	ItemCode        string  `db:"item_code"`
	ItemName        string  `db:"item_name"`
	BarCode         string  `db:"bar_code"`
	UnitICode       string  `db:"unit_code"`
	WhId            int64   `db:"wh_id"`
	ShelfId         int64   `db:"shelf_id"`
	Price           float64 `db:"price"`
	Qty             float64 `db:"qty"`
	CnQty           float64 `db:"cn_qty"`
	DiscountWord    string  `db:"discount_word_sub"`
	ItemDescription string  `db:"item_description"`
	IsCreditNote    int64   `db:"is_credit_note"`
	IsDebitNote     int64   `db:"is_debit_note"`
	PackingRate1    int64   `db:"packing_rate_1"`
	PackingRate2    int64   `db:"packing_rate_2"`
	SoRefNo         string  `db:"so_ref_no"`
	AverageCost     float64 `db:"average_cost"`
	SumOfCost       float64 `db:"sum_of_cost"`
	RefLineNumber   int64   `db:"ref_line_number"`
	LineNumber      int64   `db:"line_number"`
	ArName          string  `db:"ar_name"`
	ArCode          string  `db:"ar_code"`
	ArId            int64   `db:"ar_id"`
	MyDescription   string  `db:"my_description"`
	Name            string  `db:"name"`
	NId             int64   `db:"Id"`
	NDocNo          string  `db:"DocNo"`
	NDocDate        string  `db:"DocDate"`
	NItemId         int64   `db:"ItemId"`
	NArId           int64   `db:"ArId"`
	NBarCode        string  `db:"BarCode"`
	NItemCode       string  `db:"ItemCode"`
	NItemName       string  `db:"ItemName"`
	NUnitCode       string  `db:"UnitCode"`
	NQty            float64 `db:"Qty"`
	NPrice          float64 `db:"Price"`
	NArName         string  `db:"ArName"`
	NDiscountWord   string  `db:"DiscountWord"`
	NMyDescription  string  `db:"MyDescription"`
}
type salesRepository struct{ db *sqlx.DB }

func NewSalesRepository(db *sqlx.DB) sales.Repository {
	return &salesRepository{db}
}

func (repo *salesRepository) CreateQuotation(req *sales.NewQuoTemplate) (resp interface{}, err error) {
	var check_doc_exist int64
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var uuid string
	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)
	fmt.Println("DocDate = ", req.DocDate)
	count_item_qty = 0
	count_item_unit = 0

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	if req.DocDate == "" {
		req.DocDate = DocDate
	}

	req.CreateTime = now.String()
	req.EditTime = now.String()
	req.CancelTime = now.String()

	fmt.Println("DocType = ", req.DocType)

	for _, sub_item := range req.Subs {
		if sub_item.Qty != 0 {
			count_item = count_item + 1

			fmt.Println("Count Item =", count_item)

			sum_item_amount = sum_item_amount + (sub_item.Qty * (sub_item.Price - sub_item.DiscountAmount))
		}
		if sub_item.ItemCode != "" && sub_item.Qty == 0 {
			count_item_qty = count_item_qty + 1
		}
		if sub_item.ItemCode != "" && sub_item.UnitCode == "" {
			count_item_unit = count_item_unit + 1
		}
	}

	switch {
	case req.DocNo == "":
		return nil, errors.New("Docno is null")
	}

	sqlexist := `select count(DocNo) as check_exist from Quotation where id = ?`
	fmt.Println("DocNo Id", req.Id)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	fmt.Println("check_doc_exist", check_doc_exist)

	if check_doc_exist == 0 {
		//API Call Get API
		//url := "http://localhost:8081/gendocno/v1/gen?table_code=QT&bill_type=0"
		//reqs, err := http.NewRequest("POST", url, nil)
		//if err != nil {
		//	log.Fatal("NewRequest: ", err)
		//	return nil, err
		//}
		//
		//client := &http.Client{}
		//
		//resp, err := client.Do(reqs)
		//if err != nil {
		//	log.Fatal("Do: ", err)
		//	return nil, err
		//}
		//
		//defer resp.Body.Close()
		//
		//if err := json.NewDecoder(resp.Body).Decode(&new_doc_no); err != nil {
		//	log.Println(err)
		//}
		//
		////API Get Post API
		//url := "http://venus.nopadol.com:8081/gendocno/v1/gen"
		//var jsonStr []byte
		//
		////append(jsonStr, "":"")
		//
		//if req.BillType == 0 {
		//	jsonStr = []byte(`{"table_code":"QT","bill_type":0}`)
		//} else {
		//	jsonStr = []byte(`{"table_code":"QT","bill_type":1}`)
		//}
		//
		//reqs, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		//reqs.Header.Set("X-Custom-Header", "myvalue")
		//reqs.Header.Set("Content-Type", "application/json")
		//
		//client := &http.Client{}
		//resp, err := client.Do(reqs)
		//if err != nil {
		//	panic(err)
		//}
		//defer resp.Body.Close()
		//
		//if err := json.NewDecoder(resp.Body).Decode(&new_doc_no); err != nil {
		//	log.Println(err)
		//}
		//
		//req.DocNo = new_doc_no
		//
		//fmt.Println("Docno =", req.DocNo, new_doc_no)

		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)
		req.NetDebtAmount = req.TotalAmount

		uuid = GetAccessToken()

		sql := `INSERT INTO Quotation(uuid,DocNo,DocDate,BillType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartId,RefNo,JobId,TaxType,TaxRate,DueDate,ExpireDate,DeliveryDate,AssertStatus,IsConditionSend,MyDescription,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,ProjectId,CreateBy,CreateTime,Validity,CreditDay,ExpireCredit,DeliveryDay,AllocateId,DocType,BranchId,CompanyId,ContactId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		res, err := repo.db.Exec(sql,
			uuid,
			req.DocNo,
			req.DocDate,
			req.BillType,
			req.ArId,
			req.ArCode,
			req.ArName,
			req.SaleId,
			req.SaleCode,
			req.SaleName,
			req.DepartId,
			req.RefNo,
			req.JobId,
			req.TaxType,
			req.TaxRate,
			req.DueDate,
			req.ExpireDate,
			req.DeliveryDate,
			req.AssertStatus,
			req.IsConditionSend,
			req.MyDescription,
			req.SumOfItemAmount,
			req.DiscountWord,
			req.DiscountAmount,
			req.AfterDiscountAmount,
			req.BeforeTaxAmount,
			req.TaxAmount,
			req.TotalAmount,
			req.NetDebtAmount,
			req.ProjectId,
			req.CreateBy,
			req.CreateTime,
			req.Validity,
			req.CreditDay,
			req.ExpireCredit,
			req.DeliveryDay,
			req.AllocateId,
			req.DocType,
			req.BranchId,
			req.CompanyId,
			req.ContactId)

		fmt.Println("query=", sql, "Hello")
		if err != nil {
			return "", err
		}

		id, _ := res.LastInsertId()
		req.Id = id
		fmt.Println("New Quotation", req.Id)

		for _, sub := range req.Subs {
			fmt.Println("ArId Sub = ", req.ArId)
			fmt.Println("SaleId Sub = ", req.SaleId)
			sqlsub := `INSERT INTO QuotationSub(quo_uuid,QuoId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,Qty,RemainQty,Price,DiscountWord,DiscountAmount,UnitCode,ItemAmount,ItemDescription,PackingRate1,LineNumber,IsCancel,WhCode,ShelfCode) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err := repo.db.Exec(sqlsub,
				uuid,
				req.Id,
				req.ArId,
				req.SaleId,
				sub.ItemId,
				sub.ItemCode,
				sub.BarCode,
				sub.ItemName,
				sub.Qty,
				sub.Qty,
				sub.Price,
				sub.DiscountWord,
				sub.DiscountAmount,
				sub.UnitCode,
				sub.ItemAmount,
				sub.ItemDescription,
				sub.PackingRate1,
				sub.LineNumber,
				sub.IsCancel,
				sub.WHCode,
				sub.ShelfCode)

			fmt.Println("QuotationSub =", sql, sub.QuoId)
			if err != nil {
				return "Insert Quotation Not Success", err
			}
		}

	} else {
		fmt.Println("Update")

		switch {
		case req.DocNo == "":
			fmt.Println("error =", "Docno is null")
			return nil, errors.New("Docno is null")
		case req.BillStatus != 0:
			return nil, errors.New("เอกสารโดนอ้างนำไปใช้งานแล้ว")
		case req.IsConfirm == 1:
			return nil, errors.New("เอกสารโดนอ้างนำไปใช้งานแล้ว")
		case req.IsCancel == 1:
			return nil, errors.New("เอกสารถุกยกเลิกไปแล้ว")
		}

		req.EditBy = req.CreateBy

		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)
		req.NetDebtAmount = req.TotalAmount

		sql := `Update Quotation set DocDate=?,BillType=?,ArId=?,ArCode=?,ArName=?,SaleId=?,SaleCode=?,SaleName=?,DepartId=?,RefNo=?,JobId=?,TaxType=?,TaxRate=?,DueDate=?,ExpireDate=?,DeliveryDate=?,AssertStatus=?,IsConditionSend=?,MyDescription=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscountAmount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,NetDebtAmount=?,ProjectId=?,EditBy=?,EditTime=?,AllocateId=?,DocType=?,CompanyId=?,BranchId=?,Validity=?,CreditDay=?,ExpireCredit=?,DeliveryDay=?,ContactId=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.DocDate, req.BillType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.DepartId, req.RefNo, req.JobId, req.TaxType, req.TaxRate, req.DueDate, req.ExpireDate, req.DeliveryDate, req.AssertStatus, req.IsConditionSend, req.MyDescription, req.SumOfItemAmount, req.DiscountWord, req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetDebtAmount, req.ProjectId, req.EditBy, req.EditTime, req.AllocateId, req.DocType, req.CompanyId, req.BranchId, req.Validity, req.CreditDay, req.ExpireCredit, req.DeliveryDay, req.ContactId, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)
	}

	fmt.Println("ReqID=", req.Id)

	sql_del_sub := `delete from QuotationSub where QuoId = ?`
	_, err = repo.db.Exec(sql_del_sub, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	var line_number int

	for _, sub := range req.Subs {
		sub.LineNumber = line_number
		sqlsub := `INSERT INTO QuotationSub(QuoId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,Qty,RemainQty,Price,DiscountWord,DiscountAmount,UnitCode,ItemAmount,ItemDescription,PackingRate1,LineNumber,IsCancel,WhCode,ShelfCode) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := repo.db.Exec(sqlsub,
			req.Id,
			req.ArId,
			req.SaleId,
			sub.ItemId,
			sub.ItemCode,
			sub.BarCode,
			sub.ItemName,
			sub.Qty,
			sub.Qty,
			sub.Price,
			sub.DiscountWord,
			sub.DiscountAmount,
			sub.UnitCode,
			sub.ItemAmount,
			sub.ItemDescription,
			sub.PackingRate1,
			sub.LineNumber,
			sub.IsCancel,
			sub.WHCode,
			sub.ShelfCode)
		if err != nil {
			return nil, err
		}

		line_number = line_number + 1
	}

	return map[string]interface{}{
		"id":       req.Id,
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
		"ar_code":  req.ArCode,
	}, nil
}

func (repo *salesRepository) CancelQuotation(req *sales.NewQuoTemplate) (resp interface{}, err error) {
	var check_doc_exist int64

	now := time.Now()

	req.CancelTime = now.String()

	switch {
	case req.CancelBy == "":
		return nil, errors.New("ไม่ได้ระบุผู้ยกเลิก")
	case req.IsConfirm == 1:
		return nil, errors.New("เอกสารถูกอ้างอิงไปแล้ว ไม่สามารถยกเลิกได้")
	case req.IsCancel == 1:
		return nil, errors.New("เอกสารถูกยกเลิกไปแล้ว ไม่สามารถยกเลิกได้")
	}

	sqlexist := `select count(DocNo) as check_exist from Quotation where id = ?`
	fmt.Println("DocNo Id", req.Id)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	fmt.Println("check_doc_exist", check_doc_exist)

	if check_doc_exist != 0 {
		fmt.Println("Cancel")

		sql := `Update Quotation set IsCancel=1,CancelBy=?,CancelTime=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.CancelBy, req.CancelTime, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		fmt.Println("ReqID=", req.Id)

		sql_del_sub := `update QuotationSub set IsCancel = 1 where QuoId = ?`
		_, err = repo.db.Exec(sql_del_sub, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}

	return map[string]interface{}{
		"id":       req.Id,
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
		"ar_code":  req.ArCode,
	}, nil
}

func (repo *salesRepository) ConfirmQuotation(req *sales.NewQuoTemplate) (resp interface{}, err error) {
	var check_doc_exist int64
	now := time.Now()

	req.ConfirmTime = now.String()

	switch {
	case req.ConfirmBy == "":
		return nil, errors.New("ไม่ได้ระบุผู้อนุมัติ")
	case req.IsConfirm == 1:
		return nil, errors.New("เอกสารถูกอ้างอิงไปแล้ว ไม่สามารถอนุมัติได้")
	case req.IsCancel == 1:
		return nil, errors.New("เอกสารถูกยกเลิกไปแล้ว ไม่สามารถอนุมัติได้")
	case req.AssertStatus == 0:
		return nil, errors.New("เอกสารยังไม่ได้ตอบกลับ ไม่สามารถอนุมัติได้")
		//case req.ExpireDate <= now:
		//	return nil, errors.New("เอกสารยังไม่ได้ตอบกลับ ไม่สามารถอนุมัติได้")
	}

	sqlexist := `select count(DocNo) as check_exist from Quotation where id = ?`
	fmt.Println("DocNo Id", req.Id)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	fmt.Println("check_doc_exist", check_doc_exist)
	if check_doc_exist != 0 {
		fmt.Println("Confirm")

		sql := `Update Quotation set IsConfirm=1,ConfirmBy=?,ConfirmTime=? where Id=?`
		fmt.Println("sql confirm = ", sql)
		id, err := repo.db.Exec(sql, req.ConfirmBy, req.ConfirmTime, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		fmt.Println("ReqID=", req.Id)
	}

	return map[string]interface{}{
		"id":       req.Id,
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
		"ar_code":  req.ArCode,
	}, nil
}

func (repo *salesRepository) SearchQuoById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	q := NewQuoModel{}

	sql := `select a.Id,a.DocNo,a.DocDate,a.DocType,a.Validity,a.BillType,a.ArId,a.ArCode,a.ArName,a.SaleId,a.SaleCode,a.SaleName,ifnull(a.DepartId,0) as DepartId,ifnull(a.RefNo,'') as RefNo,ifnull(a.JobId,'') as JobId,a.TaxType,a.IsConfirm,a.BillStatus,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.ExpireCredit,ifnull(a.ExpireDate,'') as ExpireDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.AssertStatus,a.IsConditionSend,ifnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.TaxRate,a.ProjectId,a.AllocateId,a.IsCancel,ifnull(a.CreateBy,'') as CreateBy,ifnull(a.CreateTime,'') as CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime,ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone ,ifnull(a.ConfirmBy,'') as ConfirmBy,ifnull(a.ConfirmTime,'') as ConfirmTime, ifnull(a.ContactId,'') as ContactId
	from Quotation a left join Customer b on a.ArId = b.id  
	where a.Id = ?`
	err = repo.db.Get(&q, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	qt_resp := map_quo_template(q)

	subs := []NewQuoItemModel{}

	sql_sub := `select a.Id,a.QuoId,a.ItemId,a.ItemCode,a.ItemName,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel,a.WhCode,a.ShelfCode
	from QuotationSub a  where QuoId = ? order by a.linenumber`
	err = repo.db.Select(&subs, sql_sub, q.Id)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	for _, sub := range subs {
		subline := map_quo_subs_template(sub)
		qt_resp.Subs = append(qt_resp.Subs, subline)
	}

	return qt_resp, nil
}

func (repo *salesRepository) SearchQuoByKeyword(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {

	d := []SearchDocModel{}

	if req.Keyword == "" {
		sql := `select a.Id,a.DocNo,a.DocDate, case when a.DocType = 0 then 'BO' else 'QT' end as Module,a.ArCode,a.ArName,a.SaleCode,a.SaleName,ifnull(a.MyDescription,'') as MyDescription,a.TotalAmount, a.IsCancel, a.IsConfirm from Quotation a Where a.SaleCode = ? order by Id desc limit 30`
		err = repo.db.Select(&d, sql, req.SaleCode)
	} else {
		sql := `select a.Id,a.DocNo,a.DocDate, case when a.DocType = 0 then 'BO' else 'QT' end as Module,a.ArCode,a.ArName,a.SaleCode,a.SaleName,ifnull(a.MyDescription,'') as MyDescription,a.TotalAmount, a.IsCancel, a.IsConfirm from Quotation a Where (a.DocNo like CONCAT("%",?,"%") or a.ArCode like CONCAT("%",?,"%") or a.ArName like CONCAT("%",?,"%") or a.SaleCode like CONCAT("%",?,"%") or a.SaleName like CONCAT("%",?,"%")) order by Id desc limit 30`
		err = repo.db.Select(&d, sql, req.Keyword, req.Keyword, req.Keyword, req.Keyword, req.Keyword)
	}

	//sql := `select a.Id,a.DocNo,a.DocDate, case 'QT' as Module,a.ArCode,a.ArName,a.SaleCode,a.SaleName,ifnull(a.MyDescription,'') as MyDescription,a.TotalAmount, a.IsCancel, a.IsConfirm from Quotation a where a.SaleCode = ? and (a.DocNo like CONCAT("%",?,"%") or a.ArCode like CONCAT("%",?,"%") or a.ArName like CONCAT("%",?,"%") or a.SaleCode like CONCAT("%",?,"%") or a.SaleName like CONCAT("%",?,"%")) order by Id desc limit 30`
	//err = repo.db.Select(&d, sql)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	doc := []sales.SearchDocTemplate{}

	for _, c := range d {

		docline := map_doc_template(c)
		doc = append(doc, docline)
	}

	return doc, nil
}

func (repo *salesRepository) SearchDocById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {
	doc := SearchDocDetailsModel{}

	return doc, nil
}

func (repo *salesRepository) SearchDocByKeyword(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {

	d := []SearchDocModel{}

	sql := `call USP_SO_SearchDoc (?,?)`

	err = repo.db.Select(&d, sql, req.SaleCode, req.Keyword)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	doc := []sales.SearchDocTemplate{}

	for _, c := range d {

		docline := map_doc_template(c)
		doc = append(doc, docline)
	}

	return doc, nil
}

func (repo *salesRepository) QuotationToSaleOrder(req *sales.SearchByIdTemplate) (resp interface{}, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var item_discount_amount_sub float64
	var new_doc_no string
	var uuid string
	var so_id int64

	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)
	count_item_qty = 0
	count_item_unit = 0

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	doc_date := now.AddDate(0, 0, 0).Format("2006-01-02")

	q := NewQuoModel{}

	sql := `select a.Id,a.CompanyId,a.BranchId,a.DocNo,a.DocDate,a.DocType,a.Validity,a.BillType,a.ArId,a.ArCode,a.ArName,a.SaleId,a.SaleCode,a.SaleName,ifnull(a.DepartId,0) as DepartId,ifnull(a.RefNo,'') as RefNo,ifnull(a.JobId,'') as JobId,a.TaxType,a.IsConfirm,a.BillStatus,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.ExpireCredit,ifnull(a.ExpireDate,'') as ExpireDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.AssertStatus,a.IsConditionSend,ifnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.TaxRate,a.ProjectId,a.AllocateId,a.IsCancel,ifnull(a.CreateBy,'') as CreateBy,ifnull(a.CreateTime,'') as CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime,ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from Quotation a left join Customer b on a.ArId = b.id  where a.Id = ?`
	err = repo.db.Get(&q, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	qt_resp := map_quo_template(q)

	subs := []NewQuoItemModel{}

	sql_sub := `select a.Id,a.QuoId,a.ItemId,a.ItemCode,a.ItemName,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel,a.WhCode,a.ShelfCode from QuotationSub a  where QuoId = ? order by a.linenumber`
	err = repo.db.Select(&subs, sql_sub, req.Id)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	for _, sub := range subs {
		fmt.Println("sub = ", subs[0].ItemName)
		subline := map_quo_subs_template(sub)
		qt_resp.Subs = append(qt_resp.Subs, subline)
	}

	if qt_resp.DocDate == "" {
		qt_resp.DocDate = doc_date
	}

	create_time := now.String()

	fmt.Println("DocDate = ", q.DocDate)

	for _, sub_item := range qt_resp.Subs {
		if sub_item.Qty != 0 {

			count_item = count_item + 1

			if sub_item.DiscountWord != "" {
				item_discount_amount_sub, err = strconv.ParseFloat(sub_item.DiscountWord, 64)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				item_discount_amount_sub = 0
			}

			sum_item_amount = sum_item_amount + (sub_item.Qty * (sub_item.Price - item_discount_amount_sub))
		}
		if sub_item.ItemCode != "" && sub_item.Qty == 0 {
			count_item_qty = count_item_qty + 1
		}
		if sub_item.ItemCode != "" && sub_item.UnitCode == "" {
			count_item_unit = count_item_unit + 1
		}
	}

	switch {
	case qt_resp.AssertStatus == 0:
		fmt.Println("error =", "Docno is not aready to saleorder")
		return nil, errors.New("Docno is not aready to saleorder assert status not prompt")
	case qt_resp.IsCancel == 1:
		return nil, errors.New("เอกสารถูกยกเลิกไปแล้ว ไม่สามารถทำใบสั่งขายได้")
	case qt_resp.BillStatus == 1:
		return nil, errors.New("เอกสารถูกอ้างอิงไปแล้ว ไม่สามารถทำใบสั่งขายได้")
		//case req.ExpireDate <= now:
		//	return nil, errors.New("เอกสารยังไม่ได้ตอบกลับ ไม่สามารถอนุมัติได้")
	}

	d := gendocno.DocNoTemplate{}
	d.BranchId = qt_resp.BranchId
	d.BillType = qt_resp.BillType
	d.TableCode = "SO"

	//API Get Post API
	url := "https://n9.nopadol.com/gendocno/v1/gen"
	var jsonStr []byte

	//append(jsonStr, "":"")

	if d.BillType == 0 && qt_resp.BranchId == 1 {
		jsonStr = []byte(`{"table_code":"SO","bill_type":0, "branch_id":1}`)
	} else if d.BillType == 1 && qt_resp.BranchId == 1 {
		jsonStr = []byte(`{"table_code":"SO","bill_type":1, "branch_id":1}`)
	} else if d.BillType == 0 && qt_resp.BranchId == 2 {
		jsonStr = []byte(`{"table_code":"SO","bill_type":0, "branch_id":2}`)
	} else if d.BillType == 1 && qt_resp.BranchId == 2 {
		jsonStr = []byte(`{"table_code":"SO","bill_type":1, "branch_id":2}`)
	}

	reqs, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	reqs.Header.Set("X-Custom-Header", "myvalue")
	reqs.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp_docno, err := client.Do(reqs)
	if err != nil {
		panic(err)
	}
	defer resp_docno.Body.Close()

	if err := json.NewDecoder(resp_docno.Body).Decode(&new_doc_no); err != nil {
		log.Println(err)
	}

	fmt.Println("new doc_no = ", new_doc_no)

	doc_no := new_doc_no

	sqlexist := `select count(DocNo) as check_exist from SaleOrder where DocNo = ?`
	err = repo.db.Get(&check_doc_exist, sqlexist, doc_no)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	var HoldingStatus int
	var DeliveryAddressId int
	var CarLicense string
	var PersonReceiveTel string
	var create_by string
	var wh_code string
	var shelf_code string

	wh_code = "S1-A"
	shelf_code = "-"

	var credit_day int
	var delivery_day int

	credit_day = int(qt_resp.CreditDay)
	delivery_day = int(qt_resp.DeliveryDay)

	due_date := now.AddDate(0, 0, credit_day).Format("2006-01-02") //strconv.Itoa(97)
	delivery_date := now.AddDate(0, 0, delivery_day).Format("2006-01-02")

	if check_doc_exist == 0 {

		q.BeforeTaxAmount, q.TaxAmount, q.TotalAmount = config.CalcTaxItem(q.TaxType, q.TaxRate, q.AfterDiscountAmount)

		uuid = GetAccessToken()

		sql := `INSERT INTO SaleOrder(uuid,DocNo,DocDate,CompanyId,BranchId,DocType,BillType,TaxType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartId,CreditDay,DueDate,DeliveryDay,DeliveryDate,TaxRate,IsConfirm,MyDescription,BillStatus,HoldingStatus,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,IsCancel,IsConditionSend,DeliveryAddressId,CarLicense,PersonReceiveTel,JobId,ProjectId,AllocateId,CreateBy,CreateTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		res, err := repo.db.Exec(sql,
			uuid,
			doc_no,
			doc_date,
			qt_resp.CompanyId,
			qt_resp.BranchId,
			qt_resp.DocType,
			qt_resp.BillType,
			qt_resp.TaxType,
			qt_resp.ArId,
			qt_resp.ArCode,
			qt_resp.ArName,
			qt_resp.SaleId,
			qt_resp.SaleCode,
			qt_resp.SaleName,
			qt_resp.DepartId,
			qt_resp.CreditDay,
			due_date,
			qt_resp.DeliveryDay,
			delivery_date,
			qt_resp.TaxRate,
			qt_resp.IsConfirm,
			qt_resp.MyDescription,
			qt_resp.BillStatus,
			HoldingStatus,
			qt_resp.SumOfItemAmount,
			qt_resp.DiscountWord,
			qt_resp.DiscountAmount,
			qt_resp.AfterDiscountAmount,
			qt_resp.BeforeTaxAmount,
			qt_resp.TaxAmount,
			qt_resp.TotalAmount,
			qt_resp.NetDebtAmount,
			qt_resp.IsCancel,
			qt_resp.IsConditionSend,
			DeliveryAddressId,
			CarLicense,
			PersonReceiveTel,
			qt_resp.JobId,
			qt_resp.ProjectId,
			qt_resp.AllocateId,
			create_by,
			create_time)

		//fmt.Println("query=", sql, "Hello")
		if err != nil {
			return "", err
		}

		so_id, _ = res.LastInsertId()
		fmt.Println("SaleOrder Id =", so_id)

	}

	var vLineNumber int
	vLineNumber = 0

	for _, sub := range qt_resp.Subs {
		sqlsub := `INSERT INTO SaleOrderSub(so_uuid,SOId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,WhCode,ShelfCode,Qty,RemainQty,UnitCode,Price,DiscountWord,DiscountAmount,ItemAmount,ItemDescription,StockType,AverageCost,SumOfCost,RefNo,QuoId,IsCancel,PackingRate1,RefLineNumber,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := repo.db.Exec(sqlsub,
			uuid,
			so_id,
			qt_resp.ArId,
			qt_resp.SaleId,
			sub.ItemId,
			sub.ItemCode,
			sub.BarCode,
			sub.ItemName,
			wh_code,
			shelf_code,
			sub.Qty,
			sub.RemainQty,
			sub.UnitCode,
			sub.Price,
			sub.DiscountWord,
			sub.DiscountAmount,
			sub.ItemAmount,
			sub.ItemDescription,
			0,
			0,
			0,
			"",
			req.Id,
			sub.IsCancel,
			sub.PackingRate1,
			0,
			sub.LineNumber)

		vLineNumber = vLineNumber + 1
		if err != nil {
			return "Insert SaleOrder Not Success", err
		}

		sql_sub := `update QuotationSub set RemainQty = 0 where Id = ? and QuoId = ?`
		_, err = repo.db.Exec(sql_sub, sub.Id, req.Id)
		if err != nil {
			return "Insert SaleOrder Not Success", err
		}

	}

	sql_confirm := `update Quotation set BillStatus = 1 where id = ?`
	rs, err := repo.db.Exec(sql_confirm, req.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rs)

	return map[string]interface{}{
		"id":       so_id,
		"doc_no":   doc_no,
		"doc_date": doc_date,
	}, nil
}

func map_doc_template(x SearchDocModel) sales.SearchDocTemplate {
	return sales.SearchDocTemplate{
		Id:            x.Id,
		DocNo:         x.DocNo,
		DocDate:       x.DocDate,
		ArCode:        x.ArCode,
		ArName:        x.ArName,
		SaleCode:      x.SaleCode,
		SaleName:      x.SaleName,
		TotalAmount:   x.TotalAmount,
		MyDescription: x.MyDescription,
		Module:        x.Module,
		IsCancel:      x.IsCancel,
		IsConfirm:     x.IsConfirm,
	}
}

func map_quo_template(x NewQuoModel) sales.NewQuoTemplate {
	return sales.NewQuoTemplate{
		Id:                  x.Id,
		DocType:             x.DocType,
		DocNo:               x.DocNo,
		DocDate:             x.DocDate,
		BillType:            x.BillType,
		ArId:                x.ArId,
		ArCode:              x.ArCode,
		ArName:              x.ArName,
		ArBillAddress:       x.ArBillAddress,
		ArTelephone:         x.ArTelephone,
		SaleId:              x.SaleId,
		SaleCode:            x.SaleCode,
		SaleName:            x.SaleName,
		DepartId:            x.DepartId,
		RefNo:               x.RefNo,
		TaxType:             x.TaxType,
		TaxRate:             x.TaxRate,
		Validity:            x.Validity,
		CreditDay:           x.CreditDay,
		DueDate:             x.DueDate,
		ExpireCredit:        x.ExpireCredit,
		ExpireDate:          x.ExpireDate,
		DeliveryDay:         x.DeliveryDay,
		DeliveryDate:        x.DeliveryDate,
		AssertStatus:        x.AssertStatus,
		IsConditionSend:     x.IsConditionSend,
		MyDescription:       x.MyDescription,
		SumOfItemAmount:     x.SumOfItemAmount,
		DiscountWord:        x.DiscountWord,
		DiscountAmount:      x.DiscountAmount,
		AfterDiscountAmount: x.AfterDiscountAmount,
		BeforeTaxAmount:     x.BeforeTaxAmount,
		TaxAmount:           x.TaxAmount,
		TotalAmount:         x.TotalAmount,
		NetDebtAmount:       x.NetDebtAmount,
		ProjectId:           x.ProjectId,
		AllocateId:          x.AllocateId,
		CreateBy:            x.CreateBy,
		CreateTime:          x.CreateTime,
		EditBy:              x.EditBy,
		EditTime:            x.EditTime,
		CancelBy:            x.CancelBy,
		CancelTime:          x.CancelTime,
		BillStatus:          x.BillStatus,
		BranchId:            x.BranchId,
		CompanyId:           x.CompanyId,
		ConfirmBy:           x.ConfirmBy,
		ConfirmTime:         x.ConfirmTime,
		IsCancel:            x.IsCancel,
		IsConfirm:           x.IsConfirm,
		JobId:               x.JobId,
		ContactId:           x.ContactId,
	}
}

func map_quo_subs_template(x NewQuoItemModel) sales.NewQuoItemTemplate {
	return sales.NewQuoItemTemplate{
		Id:              x.Id,
		QuoId:           x.QuoId,
		ItemId:          x.ItemId,
		ItemCode:        x.ItemCode,
		BarCode:         x.BarCode,
		ItemName:        x.ItemName,
		Qty:             x.Qty,
		RemainQty:       x.RemainQty,
		Price:           x.Price,
		DiscountWord:    x.DiscountWord,
		DiscountAmount:  x.DiscountAmount,
		UnitCode:        x.UnitCode,
		ItemAmount:      x.ItemAmount,
		ItemDescription: x.ItemDescription,
		PackingRate1:    x.PackingRate1,
		LineNumber:      x.LineNumber,
		IsCancel:        x.IsCancel,
		WHCode:          x.WHCode,
		ShelfCode:       x.ShelfCode,
	}
}

func (repo *salesRepository) CreateSaleOrder(req *sales.NewSaleTemplate) (resp interface{}, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var item_discount_amount_sub float64
	var credit_balance int64
	var uuid string

	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)
	count_item_qty = 0
	count_item_unit = 0

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	if req.DocDate == "" {
		req.DocDate = DocDate
	}

	req.CreateTime = now.String()
	req.EditTime = now.String()
	req.CancelTime = now.String()

	fmt.Println("DocDate = ", req.DocDate)

	for _, sub_item := range req.Subs {
		if sub_item.Qty != 0 {
			count_item = count_item + 1

			if sub_item.DiscountWord != "" {
				item_discount_amount_sub, err = strconv.ParseFloat(sub_item.DiscountWord, 64)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				item_discount_amount_sub = 0
			}

			sum_item_amount = sum_item_amount + (sub_item.Qty * (sub_item.Price - item_discount_amount_sub))
		}
		if sub_item.ItemCode != "" && sub_item.Qty == 0 {
			count_item_qty = count_item_qty + 1
		}
		if sub_item.ItemCode != "" && sub_item.UnitCode == "" {
			count_item_unit = count_item_unit + 1
		}
	}

	switch {
	case req.DocNo == "":
		fmt.Println("error =", "Docno is null")
		return nil, errors.New("Docno is null")
	}

	fmt.Println("DocNo =", req.DocNo)

	sqlexist := `select count(DocNo) as check_exist from SaleOrder where Id = ?`
	err = repo.db.Get(&check_doc_exist, sqlexist, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	//CheckCredit
	fmt.Println("BillType = ", req.BillType)
	if req.BillType == 1 {
		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)
		check_balance := `select sum(debt_limit - (debt_amount+?)) as check_credit from Customer where code = ?`
		err = repo.db.Get(&credit_balance, check_balance, req.TotalAmount, req.ArCode)
		fmt.Println("This Value =", req.TotalAmount)
		fmt.Println("ArCode = ", req.ArCode)
		fmt.Println("check_balance = ", check_balance)
		if err != nil {
			return "", err
		}

		if credit_balance > 0 {
			fmt.Println("credit enough")
		} else {
			ins_credit := `update SaleOrder set HoldingStatus=1 where DocNo=? `
			_, err := repo.db.Exec(ins_credit, req.DocNo)
			fmt.Println("ins_credit =", ins_credit)
			fmt.Println("This Value =", req.DocNo)
			if err != nil {
				return "", err
			}
			fmt.Println("credit not enough")
		}
	}

	if check_doc_exist == 0 {
		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

		uuid = GetAccessToken()

		sql := `INSERT INTO SaleOrder(uuid,DocNo,DocDate,CompanyId,BranchId,DocType,BillType,TaxType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartId,CreditDay,DueDate,DeliveryDay,DeliveryDate,TaxRate,IsConfirm,MyDescription,BillStatus,HoldingStatus,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,IsCancel,IsConditionSend,DeliveryAddressId,CarLicense,PersonReceiveTel,JobId,RefNo,ProjectId,AllocateId,CreateBy,CreateTime,ContactId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		res, err := repo.db.Exec(sql,
			uuid,
			req.DocNo,
			req.DocDate,
			req.CompanyId,
			req.BranchId,
			req.DocType,
			req.BillType,
			req.TaxType,
			req.ArId,
			req.ArCode,
			req.ArName,
			req.SaleId,
			req.SaleCode,
			req.SaleName,
			req.DepartId,
			req.CreditDay,
			req.DueDate,
			req.DeliveryDay,
			req.DeliveryDate,
			req.TaxRate,
			req.IsConfirm,
			req.MyDescription,
			req.BillStatus,
			req.HoldingStatus,
			req.SumOfItemAmount,
			req.DiscountWord,
			req.DiscountAmount,
			req.AfterDiscountAmount,
			req.BeforeTaxAmount,
			req.TaxAmount,
			req.TotalAmount,
			req.NetDebtAmount,
			req.IsCancel,
			req.IsConditionSend,
			req.DeliveryAddressId,
			req.CarLicense,
			req.PersonReceiveTel,
			req.JobId,
			req.RefNo,
			req.ProjectId,
			req.AllocateId,
			req.CreateBy,
			req.CreateTime,
			req.ContactId)
		//fmt.Println("This Value =", req.TotalAmount)
		//fmt.Println("query=", sql, "Hello")
		if err != nil {
			return "", err
		}

		id, _ := res.LastInsertId()
		req.Id = id

	} else {
		switch {
		case req.DocNo == "":
			fmt.Println("error =", "Docno is null")
			return nil, errors.New("Docno is null")
		case req.BillStatus != 0:
			return nil, errors.New("เอกสารโดนอ้างนำไปใช้งานแล้ว")
		case req.IsConfirm == 1:
			return nil, errors.New("เอกสารโดนอ้างนำไปใช้งานแล้ว")
		case req.IsCancel == 1:
			return nil, errors.New("เอกสารถุกยกเลิกไปแล้ว")
		}

		fmt.Println("Update")
		req.EditBy = req.CreateBy

		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

		sql := `Update SaleOrder set DocNo=?,DocDate=?,CompanyId=?,BranchId=?,DocType=?,BillType=?,TaxType=?,ArId=?,ArCode=?,ArName=?,SaleId=?,SaleCode=?,SaleName=?,DepartId=?,CreditDay=?,DueDate=?,DeliveryDay=?,DeliveryDate=?,TaxRate=?,IsConfirm=?,MyDescription=?,BillStatus=?,HoldingStatus=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscountAmount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,NetDebtAmount=?,IsCancel=?,IsConditionSend=?,DeliveryAddressId=?,CarLicense=?,PersonReceiveTel=?,JobId=?,RefNo=?,ProjectId=?,AllocateId=?,EditBy=?,EditTime=?,ContactId=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.DocNo, req.DocDate, req.CompanyId, req.BranchId, req.DocType, req.BillType, req.TaxType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.DepartId, req.CreditDay, req.DueDate, req.DeliveryDay, req.DeliveryDate, req.TaxRate, req.IsConfirm, req.MyDescription, req.BillStatus, req.HoldingStatus, req.SumOfItemAmount, req.DiscountWord, req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetDebtAmount, req.IsCancel, req.IsConditionSend, req.DeliveryAddressId, req.CarLicense, req.PersonReceiveTel, req.JobId, req.RefNo, req.ProjectId, req.AllocateId, req.EditBy, req.EditTime, req.ContactId, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		fmt.Println("ReqID=", req.Id)

		sql_del_sub := `delete from SaleOrderSub where SOId = ?`
		_, err = repo.db.Exec(sql_del_sub, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

	}

	var vLineNumber int
	vLineNumber = 0

	for _, sub := range req.Subs {
		sqlsub := `INSERT INTO SaleOrderSub(so_uuid,SOId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,WhCode,ShelfCode,Qty,RemainQty,UnitCode,Price,DiscountWord,DiscountAmount,ItemAmount,ItemDescription,StockType,AverageCost,SumOfCost,RefNo,QuoId,IsCancel,PackingRate1,RefLineNumber,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := repo.db.Exec(sqlsub,
			uuid,
			req.Id,
			req.ArId,
			req.SaleId,
			sub.ItemId,
			sub.ItemCode,
			sub.BarCode,
			sub.ItemName,
			sub.WHCode,
			sub.ShelfCode,
			sub.Qty,
			sub.RemainQty,
			sub.UnitCode,
			sub.Price,
			sub.DiscountWord,
			sub.DiscountAmount,
			sub.ItemAmount,
			sub.ItemDescription,
			sub.StockType,
			sub.AverageCost,
			sub.SumOfCost,
			sub.RefNo,
			sub.QuoId,
			sub.IsCancel,
			sub.PackingRate1,
			sub.RefLineNumber,
			sub.LineNumber)

		vLineNumber = vLineNumber + 1
		if err != nil {
			return "Insert SaleOrder Not Success", err
		}
	}

	return map[string]interface{}{
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
	}, nil
}

func (repo *salesRepository) SearchSaleOrderById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	s := NewSaleModel{}

	sql := `select a.Id,a.DocNo,ifnull(a.DocDate,'') as DocDate,a.CompanyId,a.BranchId,a.DocType,a.BillType,a.TaxType,a.ArId,ifnull(a.ArCode,'') as ArCode,ifnull(a.ArName,'') as ArName,a.SaleId,ifnull(a.SaleCode,'') as SaleCode,ifnull(a.SaleName,'') as SaleName,a.DepartId,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.TaxRate,a.IsConfirm,ifnull(a.MyDescription,'') as MyDescription,a.BillStatus,a.HoldingStatus,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.IsCancel,a.IsConditionSend,a.DeliveryAddressId,ifnull(a.CarLicense,'') as CarLicense,ifnull(a.PersonReceiveTel,'') as PersonReceiveTel,ifnull(a.JobId,'') as JobId, ifnull(a.RefNo,'') as RefNo,a.ProjectId,a.AllocateId,ifnull(a.CreateBy,'') as CreateBy,a.CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime, ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime, ifnull(a.ConfirmBy,'') as ConfirmBy,ifnull(a.ConfirmTime,'') as ConfirmTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone, ifnull(a.ContactId,'') as ContactId
	from SaleOrder a left join Customer b on a.ArId = b.id 
	where a.id=?`
	err = repo.db.Get(&s, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	so_resp := map_saleorder_template(s)

	subs := []NewSaleItemModel{}

	fmt.Println("s.Id =", s.Id)

	sql_sub := `select a.Id,a.SOId,a.ItemId,a.ItemCode,a.ItemName,ifnull(a.WHCode,'') as WHCode,ifnull(a.ShelfCode,'') as ShelfCode,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.StockType,a.AverageCost,a.SumOfCost,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel 
	from SaleOrderSub a  
	where SOId = ? order by a.linenumber`
	err = repo.db.Select(&subs, sql_sub, s.Id)
	fmt.Println("sql_sub = ", sql_sub)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	for _, sub := range subs {
		subline := map_sale_subs_template(sub)
		so_resp.Subs = append(so_resp.Subs, subline)
	}

	return so_resp, nil
}

func (repo *salesRepository) SearchSaleOrderByKeyword(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {

	d := []SearchDocModel{}

	if req.Keyword == "" {
		sql := `select a.Id,a.DocNo,a.DocDate, case when a.DocType = 0 then 'RO' else 'SO' end as Module,a.ArCode,a.ArName,a.SaleCode,a.SaleName,ifnull(a.MyDescription,'') as MyDescription,a.TotalAmount, a.IsCancel, a.IsConfirm from SaleOrder a Where a.SaleCode = ? order by Id desc limit 30`
		err = repo.db.Select(&d, sql, req.SaleCode)
		fmt.Println("sale order sql empty = ", sql)
	} else {
		sql := `select a.Id,a.DocNo,a.DocDate, case when a.DocType = 0 then 'RO' else 'SO' end as Module,a.ArCode,a.ArName,a.SaleCode,a.SaleName,ifnull(a.MyDescription,'') as MyDescription,a.TotalAmount, a.IsCancel, a.IsConfirm from SaleOrder a Where (a.DocNo like CONCAT("%",?,"%") or a.ArCode like CONCAT("%",?,"%") or a.ArName like CONCAT("%",?,"%") or a.SaleCode like CONCAT("%",?,"%") or a.SaleName like CONCAT("%",?,"%")) order by Id desc limit 30`
		err = repo.db.Select(&d, sql, req.Keyword, req.Keyword, req.Keyword, req.Keyword, req.Keyword)
		fmt.Println("sale order sql = ", sql)
	}

	//sql := `select a.Id,a.DocNo,a.DocDate, case 'QT' as Module,a.ArCode,a.ArName,a.SaleCode,a.SaleName,ifnull(a.MyDescription,'') as MyDescription,a.TotalAmount, a.IsCancel, a.IsConfirm from Quotation a where a.SaleCode = ? and (a.DocNo like CONCAT("%",?,"%") or a.ArCode like CONCAT("%",?,"%") or a.ArName like CONCAT("%",?,"%") or a.SaleCode like CONCAT("%",?,"%") or a.SaleName like CONCAT("%",?,"%")) order by Id desc limit 30`
	//err = repo.db.Select(&d, sql)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	doc := []sales.SearchDocTemplate{}

	for _, c := range d {

		docline := map_doc_template(c)
		doc = append(doc, docline)
	}

	return doc, nil
}

func map_saleorder_template(x NewSaleModel) sales.NewSaleTemplate {
	return sales.NewSaleTemplate{
		AllocateId:          x.AllocateId,
		ArCode:              x.ArCode,
		ArId:                x.ArId,
		AfterDiscountAmount: x.AfterDiscountAmount,
		ArTelephone:         x.ArTelephone,
		ArBillAddress:       x.ArBillAddress,
		ArName:              x.ArName,
		BillType:            x.BillType,
		BranchId:            x.BranchId,
		BeforeTaxAmount:     x.BeforeTaxAmount,
		BillStatus:          x.BillStatus,
		CreditDay:           x.CreditDay,
		CreateTime:          x.CreateTime,
		CreateBy:            x.CreateBy,
		CompanyId:           x.CompanyId,
		CarLicense:          x.CarLicense,
		CancelTime:          x.CancelTime,
		CancelBy:            x.CancelBy,
		ConfirmBy:           x.ConfirmBy,
		ConfirmTime:         x.ConfirmTime,
		DueDate:             x.DueDate,
		DepartId:            x.DepartId,
		DocDate:             x.DocDate,
		DocNo:               x.DocNo,
		DocType:             x.DocType,
		DiscountAmount:      x.DiscountAmount,
		DiscountWord:        x.DiscountWord,
		DeliveryDay:         x.DeliveryDay,
		DeliveryAddressId:   x.DeliveryAddressId,
		DeliveryDate:        x.DeliveryDate,
		EditBy:              x.EditBy,
		EditTime:            x.EditTime,
		HoldingStatus:       x.HoldingStatus,
		Id:                  x.Id,
		IsConfirm:           x.IsConfirm,
		IsCancel:            x.IsCancel,
		IsConditionSend:     x.IsConditionSend,
		JobId:               x.JobId,
		MyDescription:       x.MyDescription,
		NetDebtAmount:       x.NetDebtAmount,
		ProjectId:           x.ProjectId,
		PersonReceiveTel:    x.PersonReceiveTel,
		RefNo:               x.RefNo,
		SaleCode:            x.SaleCode,
		SaleId:              x.SaleId,
		SaleName:            x.SaleName,
		SumOfItemAmount:     x.SumOfItemAmount,
		TotalAmount:         x.TotalAmount,
		TaxAmount:           x.TaxAmount,
		TaxRate:             x.TaxRate,
		TaxType:             x.TaxType,
		ContactId:           x.ContactId,
	}
}

func map_invoice_sub_template(x NewInvoiceItemModel) sales.NewInvoiceItemTemplate {
	fmt.Println("endpoint x", x)
	return sales.NewInvoiceItemTemplate{
		Id:              x.Id,
		InvId:           x.InvId,
		ItemCode:        x.ItemCode,
		Itemid:          x.Itemid,
		BarCode:         x.BarCode,
		ItemName:        x.ItemName,
		WhId:            x.WhId,
		ShelfId:         x.ShelfId,
		Qty:             x.Qty,
		Price:           x.Price,
		DiscountWord:    x.DiscountWord,
		DiscountAmount:  x.DiscountAmount,
		UnitCode:        x.UnitCode,
		ItemAmount:      x.ItemAmount,
		ItemDescription: x.ItemDescription,
		Average_cost:    x.Average_cost,
		SumOfCost:       x.SumOfCost,
		PackingRate1:    x.PackingRate1,
		LineNumber:      x.LineNumber,
		IsCancel:        x.IsCancel,
	}
}
func map_sale_subs_template(x NewSaleItemModel) sales.NewSaleItemTemplate {
	return sales.NewSaleItemTemplate{
		Id:              x.Id,
		QuoId:           x.QuoId,
		ItemId:          x.ItemId,
		ItemCode:        x.ItemCode,
		BarCode:         x.BarCode,
		ItemName:        x.ItemName,
		Qty:             x.Qty,
		RemainQty:       x.RemainQty,
		WHCode:          x.WHCode,
		ShelfCode:       x.ShelfCode,
		Price:           x.Price,
		DiscountWord:    x.DiscountWord,
		DiscountAmount:  x.DiscountAmount,
		UnitCode:        x.UnitCode,
		ItemAmount:      x.ItemAmount,
		ItemDescription: x.ItemDescription,
		StockType:       x.StockType,
		AverageCost:     x.AverageCost,
		SumOfCost:       x.SumOfCost,
		PackingRate1:    x.PackingRate1,
		LineNumber:      x.LineNumber,
		IsCancel:        x.IsCancel,
	}
}

func (repo *salesRepository) CreateDeposit(req *sales.NewDepositTemplate) (interface{}, error) {
	var check_doc_exist int64

	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")
	fmt.Println("tax rate = ", def.TaxRateDefault)

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	sqlexist := `select count(doc_no) as check_exist from ar_deposit where id = ? or doc_no = ?`
	err := repo.db.Get(&check_doc_exist, sqlexist, req.Id, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	uuid := GenUUID()

	if req.DocDate == "" {
		req.DocDate = DocDate
	}

	due_date := now.AddDate(0, 0, int(req.CreditDay)).Format("2006-01-02")

	req.DueDate = due_date

	fmt.Println("duedate =", req.DueDate)

	req.CreateTime = now.String()
	req.EditTime = now.String()
	req.CancelTime = now.String()
	//req.TaxRate = def.TaxRateDefault

	if req.Uuid == "" {
		req.Uuid = uuid
	}

	fmt.Println("Doc UUID = ", req.Uuid)

	if req.TotalAmount != 0 {
		req.BeforeTaxAmount, req.TaxAmount = config.CalcTaxTotalAmount(req.TaxType, req.TaxRate, req.TotalAmount)
		req.NetAmount = req.TotalAmount
	}

	fmt.Println("check_doc_exist = ", check_doc_exist, sqlexist, req.Id)

	switch {
	case req.TotalAmount != (req.CashAmount + req.CreditcardAmount + req.ChqAmount + req.BankAmount):
		return nil, errors.New("มูลค่ารวมทั้งหมดไม่ตรงกับมูลค่ารับชำระ")
	}

	if check_doc_exist == 0 {
		req.BillBalance = req.NetAmount
		sql := `insert into ar_deposit(company_id, branch_id, uuid, doc_no, tax_no, doc_date, bill_type, ar_id, ar_code, ar_name, sale_id, sale_code, sale_name, tax_type, tax_rate, ref_no, credit_day, due_date, depart_id, allocate_id, project_id, my_description, before_tax_amount, tax_amount, total_amount, net_amount ,bill_balance ,cash_amount ,creditcard_amount, chq_amount, bank_amount, is_return_money, is_cancel, is_confirm, scg_id, job_no, create_by, create_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,? ,? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		resp, err := repo.db.Exec(sql, req.CompanyId, req.BranchId, req.Uuid, req.DocNo, req.TaxNo, req.DocDate, req.BillType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.TaxType, req.TaxRate, req.RefNo, req.CreditDay, req.DueDate, req.DepartId, req.AllocateId, req.ProjectId, req.MyDescription, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetAmount, req.BillBalance, req.CashAmount, req.CreditcardAmount, req.ChqAmount, req.BankAmount, req.IsReturnMoney, req.IsCancel, req.IsConfirm, req.ScgId, req.JobNo, req.CreateBy, req.CreateTime)
		if err != nil {
			fmt.Println("error = ", err.Error())
		}
		fmt.Println("sql = ", sql)

		id, _ := resp.LastInsertId()

		req.Id = id
	} else {

		switch {
		case req.IsConfirm == 1:
			return nil, errors.New("เอกสารโดนอ้างนำไปใช้งานแล้ว")
		case req.IsCancel == 1:
			return nil, errors.New("เอกสารถุกยกเลิกไปแล้ว")
		}

		sql := `update ar_deposit set company_id=?, branch_id=?, uuid=?, doc_no=?,tax_no=?, doc_date=?, bill_type=?, ar_id=?, ar_code=?, ar_name=?, sale_id=?, sale_code=?, sale_name=?, tax_type=?, tax_rate=?, ref_no=?, credit_day=?, due_date=?, depart_id=?, allocate_id=?, project_id=?, my_description=?, before_tax_amount=?, tax_amount=?, total_amount=?, net_amount=?, bill_balance=?, cash_amount=? ,creditcard_amount=?, chq_amount=?, bank_amount=?, is_return_money=?, is_cancel=?, is_confirm=?, scg_id=?, job_no=?, edit_by=?, edit_time=?  where id = ?`
		resp, err := repo.db.Exec(sql, req.CompanyId, req.BranchId, req.Uuid, req.DocNo, req.TaxNo, req.DocDate, req.BillType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.TaxType, req.TaxRate, req.RefNo, req.CreditDay, req.DueDate, req.DepartId, req.AllocateId, req.ProjectId, req.MyDescription, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetAmount, req.BillBalance, req.CashAmount, req.CreditcardAmount, req.ChqAmount, req.BankAmount, req.IsReturnMoney, req.IsCancel, req.IsConfirm, req.ScgId, req.JobNo, req.EditBy, req.EditTime, req.Id)
		if err != nil {
			fmt.Println("error = ", err.Error())
		}
		fmt.Println("sql = ", sql)

		rowAffect, err := resp.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		fmt.Println("ReqID=", req.Id)

	}

	var count_crd_err int64
	var sum_crd_amount float64

	fmt.Println("UUID1 = ", req.Uuid)

	sql_del_crd := `delete from credit_card where uuid=? and ref_id=? and company_id=? and branch_id=? `
	_, err = repo.db.Exec(sql_del_crd, req.Uuid, req.Id, req.CompanyId, req.BranchId)
	if err != nil {
		fmt.Println("sql_del = ", err.Error())
	}

	count_crd_err = 0
	if req.CreditcardAmount != 0 {
		for _, crd := range req.CreditCard {

			sql_del := `delete from credit_card where uuid=? and ref_id=? and company_id=? and branch_id=? and credit_card_no=? and confirm_no=? and bank_id=?`
			crd_del, _ := repo.db.Exec(sql_del, req.Uuid, req.Id, req.CompanyId, req.BranchId, crd.CreditCardNo, crd.ConfirmNo, crd.BankId)
			if err != nil {
				fmt.Println("sql_del = ", err.Error())
			}
			fmt.Println(crd_del.RowsAffected())

			verify_crd, _ := verify_creditcard(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, crd.CreditCardNo, crd.ConfirmNo, crd.BankId)
			fmt.Println("verify_crd = ", verify_crd)
			if verify_crd == false {
				count_crd_err = count_crd_err + 1
			}

			sum_crd_amount = sum_crd_amount + crd.Amount
		}

		fmt.Println("count_crd_err", count_crd_err)
		fmt.Println("เช็คดิต")
		switch {
		case count_crd_err != 0:
			return nil, errors.New("ข้อมูลบัตรเครดิต มีการใช้ไปแล้ว")
		case sum_crd_amount != req.CreditcardAmount:
			return nil, errors.New("มูลค่าบัตรเครดิต ไม่เท่ากับ มูลค่ารายการบัตรเครดิต")
		}

		for _, i_crd := range req.CreditCard {
			i_crd.Description = "รับเงินมัดจำ"
			sql_crd := `insert into credit_card (company_id, branch_id, uuid, ref_id, ar_id, doc_no, doc_date, credit_card_no, credit_type, confirm_no, amount, charge_amount, description, bank_id, bank_branch_id,receive_date,due_date,book_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
			crd, _ := repo.db.Exec(sql_crd, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, i_crd.CreditCardNo, i_crd.CreditType, i_crd.ConfirmNo, i_crd.Amount, i_crd.ChargeAmount, i_crd.Description, i_crd.BankId, i_crd.BankBranchId, i_crd.ReceiveDate, i_crd.DueDate, i_crd.BookId)
			if err != nil {
				return nil, err
			}
			crdRowAffect, err := crd.RowsAffected()
			if err != nil {
				return nil, err
			}
			fmt.Println("Row Affect = ", crdRowAffect)

			i_crd.Id = crdRowAffect
		}
	}

	var count_chq_err int64
	var sum_chq_amount float64

	sql_del_chq := `delete from chq_in where uuid = ? and ref_id=? and company_id=? and branch_id=?`
	_, err = repo.db.Exec(sql_del_chq, req.Uuid, req.Id, req.CompanyId, req.BranchId)
	if err != nil {
		fmt.Println("sql_del = ", err.Error())
	}

	count_chq_err = 0
	if req.ChqAmount != 0 {
		for _, chq := range req.Chq {
			fmt.Println("UUID Chq= ", req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber)

			sql_del := `delete from chq_in where uuid = ? and ref_id=? and company_id=? and branch_id=? and bank_id=? and chq_number=?`
			chq_del, _ := repo.db.Exec(sql_del, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.BankId, chq.ChqNumber)
			fmt.Println(sql_del)
			if err != nil {
				fmt.Println("sql_del = ", err.Error())
			}
			fmt.Println(chq_del)

			verify_chq, _ := verify_chq_in(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber, chq.BankId)
			fmt.Println("verify_chq = ", verify_chq)
			if verify_chq == false {
				count_chq_err = count_chq_err + 1
			}

			sum_chq_amount = sum_chq_amount + chq.ChqAmount
		}

		switch {
		case count_chq_err != 0:
			return nil, errors.New("ข้อมูลเลขที่เช็ค มีอยู่แล้ว")
		case sum_chq_amount != req.ChqAmount:
			return nil, errors.New("มูลค่าเช็ค ไม่เท่ากับ มูลค่ารายการเช็ค")
		}

		for _, i_chq := range req.Chq {
			i_chq.Description = "รับเงินมัดจำ"
			sql_chq := `insert into chq_in (company_id,branch_id,uuid,ref_id,ar_id,doc_no,doc_date,chq_number,bank_id,bank_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			chq, _ := repo.db.Exec(sql_chq, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, i_chq.ChqNumber, i_chq.BankId, i_chq.BankBranchId, i_chq.ReceiveDate, i_chq.DueDate, i_chq.BookId, i_chq.ChqStatus, i_chq.ChqAmount, i_chq.ChqAmount, i_chq.Description, req.CreateBy, req.CreateTime)
			if err != nil {
				return nil, err
			}
			chqRowAffect, err := chq.RowsAffected()
			if err != nil {
				return nil, err
			}
			fmt.Println("Row Affect = ", chqRowAffect)

			i_chq.Id = chqRowAffect
		}
	}

	return map[string]interface{}{
		"id":       req.Id,
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
		"ar_code":  req.ArCode,
	}, nil
}

func verify_creditcard(db *sqlx.DB, Uuid string, RefId int64, CompanyId int64, BranchId int64, CreditCardNo string, ConfirmNo string, BankId int64) (bool, error) {
	var exist int64

	sql_verify_crd := `select ifnull(count(doc_no),0) as vcount from credit_card where credit_card_no = ? and confirm_no = ? and bank_id = ?`
	err := db.Get(&exist, sql_verify_crd, CreditCardNo, ConfirmNo, BankId)
	fmt.Println("sql_verify_crd = ", sql_verify_crd, CreditCardNo, ConfirmNo, BankId)
	if err != nil {
		fmt.Println("sql_verify_crd err = ", err.Error())
		return false, err
	}

	fmt.Println("exist = ", exist)

	if exist != 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func verify_chq_in(db *sqlx.DB, Uuid string, RefId int64, CompanyId int64, BranchId int64, ChqNumber string, BankId int64) (bool, error) {
	var exist int64

	sql_verify_chq := `select ifnull(count(doc_no),0) as vcount from chq_in where uuid = ? and ref_id = ? and company_id = ? and branch_id = ? and chq_number = ? and bank_id = ?`
	err := db.Get(&exist, sql_verify_chq, Uuid, RefId, CompanyId, BranchId, ChqNumber, BankId)
	fmt.Println("sql_verify_chq = ", sql_verify_chq, Uuid, RefId, CompanyId, BranchId, ChqNumber, BankId)
	if err != nil {
		fmt.Println("sql_verify_chq err = ", err.Error())
		return false, err
	}

	fmt.Println("exist = ", exist)

	if exist != 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (repo *salesRepository) SearchDepositById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	d := NewDepositModel{}

	sql := `select a.id, a.company_id, a.branch_id, ifnull(a.uuid,'') as uuid, ifnull(a.doc_no,'') as doc_no, ifnull(a.tax_no,'') as tax_no, ifnull(a.doc_date,'') as doc_date, a.bill_type, a.ar_id, ifnull(a.ar_code,'') as ar_code, ifnull(a.ar_name,'') as ar_name,a.sale_id, ifnull(a.sale_code,'') as sale_code, ifnull(a.sale_name,'') as sale_name,a.tax_type, a.tax_rate, ifnull(a.ref_no,'') as ref_no, a.credit_day, ifnull(a.due_date,'') as due_date, a.depart_id, a.allocate_id, a.project_id, ifnull(a.my_description,'') as my_description, a.before_tax_amount, a.tax_amount, a.total_amount, a.net_amount ,a.bill_balance ,a.cash_amount ,a.creditcard_amount, a.chq_amount, a.bank_amount, a.is_return_money, a.is_cancel, a.is_confirm, ifnull(a.scg_id,'') as scg_id, ifnull(a.job_no,'') as job_no, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, ifnull(a.edit_by,'') as edit_by, ifnull(a.edit_time,'') as edit_time, ifnull(a.cancel_by,'') as cancel_by, ifnull(a.cancel_time,'') as cancel_time, ifnull(a.confirm_by,'') as confirm_by, ifnull(a.confirm_time,'') as confirm_time,ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone 
	from ar_deposit a left join Customer b on a.ar_id = b.id 
	where a.id=?`
	err = repo.db.Get(&d, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	dp_resp := map_deposit_template(d)

	fmt.Println("CompanyId,BranchId,Uuid", d.CompanyId, d.BranchId, d.Uuid)

	crds := []CreditCardModel{}
	sql_crd := `select id, ref_id, credit_card_no, credit_type, confirm_no, amount, charge_amount, ifnull(description,'') as description, bank_id, bank_branch_id,receive_date,due_date,book_id from credit_card where company_id = ? and branch_id = ? and ref_id=?`
	err = repo.db.Select(&crds, sql_crd, d.CompanyId, d.BranchId, req.Id)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	for _, crd := range crds {
		crd_line := map_deposit_crd_template(crd)
		dp_resp.CreditCard = append(dp_resp.CreditCard, crd_line)
	}

	chqs := []ChqInModel{}
	sql_chq := `select id,ref_id,chq_number,bank_id,bank_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description from chq_in where company_id = ? and branch_id = ? and ref_id = ? `
	err = repo.db.Select(&chqs, sql_chq, d.CompanyId, d.BranchId, d.Id)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	for _, chq := range chqs {
		chq_line := map_deposit_chq_template(chq)
		dp_resp.Chq = append(dp_resp.Chq, chq_line)
	}

	return dp_resp, nil
}

func (repo *salesRepository) SearchDepositByKeyword(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {
	var sql string

	d := []NewDepositModel{}

	if req.Keyword == "" {
		sql = `select a.id, a.company_id, a.branch_id, ifnull(a.uuid,'') as uuid, ifnull(a.doc_no,'') as doc_no, ifnull(a.tax_no,'') as tax_no, ifnull(a.doc_date,'') as doc_date, a.bill_type, a.ar_id, ifnull(a.ar_code,'') as ar_code, ifnull(a.ar_name,'') as ar_name,a.sale_id, ifnull(a.sale_code,'') as sale_code, ifnull(a.sale_name,'') as sale_name,a.tax_type, a.tax_rate, ifnull(a.ref_no,'') as ref_no, a.credit_day, ifnull(a.due_date,'') as due_date, a.depart_id, a.allocate_id, a.project_id, ifnull(a.my_description,'') as my_description, a.before_tax_amount, a.tax_amount, a.total_amount, a.net_amount ,a.bill_balance ,a.cash_amount ,a.creditcard_amount, a.chq_amount, a.bank_amount, a.is_return_money, a.is_cancel, a.is_confirm, ifnull(a.scg_id,'') as scg_id, ifnull(a.job_no,'') as job_no, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, ifnull(a.edit_by,'') as edit_by, ifnull(a.edit_time,'') as edit_time, ifnull(a.cancel_by,'') as cancel_by, ifnull(a.cancel_time,'') as cancel_time, ifnull(a.confirm_by,'') as confirm_by, ifnull(a.confirm_time,'') as confirm_time,
		ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone 
		from ar_deposit a left join Customer b on a.ar_id = b.id  
		order by a.id desc limit 30`
		err = repo.db.Select(&d, sql)
	} else {
		sql = `select a.id, a.company_id, a.branch_id, ifnull(a.uuid,'') as uuid, ifnull(a.doc_no,'') as doc_no, ifnull(a.tax_no,'') as tax_no, ifnull(a.doc_date,'') as doc_date, a.bill_type, a.ar_id, ifnull(a.ar_code,'') as ar_code, ifnull(a.ar_name,'') as ar_name,a.sale_id, ifnull(a.sale_code,'') as sale_code, ifnull(a.sale_name,'') as sale_name,a.tax_type, a.tax_rate, ifnull(a.ref_no,'') as ref_no, a.credit_day, ifnull(a.due_date,'') as due_date, a.depart_id, a.allocate_id, a.project_id, ifnull(a.my_description,'') as my_description, a.before_tax_amount, a.tax_amount, a.total_amount, a.net_amount ,a.bill_balance ,a.cash_amount ,a.creditcard_amount, a.chq_amount, a.bank_amount, a.is_return_money, a.is_cancel, a.is_confirm, ifnull(a.scg_id,'') as scg_id, ifnull(a.job_no,'') as job_no, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, ifnull(a.edit_by,'') as edit_by, ifnull(a.edit_time,'') as edit_time, ifnull(a.cancel_by,'') as cancel_by, ifnull(a.cancel_time,'') as cancel_time, ifnull(a.confirm_by,'') as confirm_by, ifnull(a.confirm_time,'') as confirm_time,ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone 
		from ar_deposit a left join Customer b on a.ar_id = b.id  
		where a.doc_no like  concat(?,'%') or a.ar_code like  concat(?,'%') or a.ar_name like  concat(?,'%') 
		order by a.id desc limit 30`
		err = repo.db.Select(&d, sql, req.Keyword, req.Keyword, req.Keyword)
	}

	fmt.Println("sql = ", sql, req.Keyword)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	dp := []sales.NewDepositTemplate{}

	for _, dep := range d {
		dpline := map_deposit_template(dep)
		dp = append(dp, dpline)
	}

	return dp, nil
}

func (repo *salesRepository) SearchReserveToDeposit(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {
	var sql string

	rsv := []NewSaleModel{}

	if req.Keyword == "" {
		sql = `select a.Id,a.DocNo,ifnull(a.DocDate,'') as DocDate,a.CompanyId,a.BranchId,a.DocType,a.BillType,a.TaxType,a.ArId,ifnull(a.ArCode,'') as ArCode,ifnull(a.ArName,'') as ArName,a.SaleId,ifnull(a.SaleCode,'') as SaleCode,ifnull(a.SaleName,'') as SaleName,a.IsConfirm,ifnull(a.MyDescription,'') as MyDescription,a.BillStatus,a.HoldingStatus,a.TotalAmount,a.NetDebtAmount,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from SaleOrder a left join Customer on a.ArId = b.id where DocType = 0 and BillStatus = 0 and IsCancel = 0 and ar_id = ? order by Id desc limit 30`
		err = repo.db.Select(&rsv, sql, req.ArId)
	} else {
		sql = `select a.Id,a.DocNo,ifnull(a.DocDate,'') as DocDate,a.CompanyId,a.BranchId,a.DocType,a.BillType,a.TaxType,a.ArId,ifnull(a.ArCode,'') as ArCode,ifnull(a.ArName,'') as ArName,a.SaleId,ifnull(a.SaleCode,'') as SaleCode,ifnull(a.SaleName,'') as SaleName,a.IsConfirm,ifnull(a.MyDescription,'') as MyDescription,a.BillStatus,a.HoldingStatus,a.TotalAmount,a.NetDebtAmount,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone 
		from SaleOrder a left join Customer on a.ArId = b.id 
		where DocType = 0 and BillStatus = 0 and IsCancel = 0 and ar_id = ? and (DocNo like  concat('%',?,'%') or MyDescription like  concat('%',?,'%') )order by Id desc limit 30`
		err = repo.db.Select(&rsv, sql, req.ArId, req.Keyword, req.Keyword)
	}
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	ro := []sales.NewSaleTemplate{}

	for _, l := range rsv {

		roline := map_ro_template(l)
		ro = append(ro, roline)
	}

	return ro, nil
}

func map_ro_template(x NewSaleModel) sales.NewSaleTemplate {
	return sales.NewSaleTemplate{
		Id:            x.Id,
		DocNo:         x.DocNo,
		DocDate:       x.DocDate,
		ArCode:        x.ArCode,
		ArName:        x.ArName,
		SaleCode:      x.SaleCode,
		SaleName:      x.SaleName,
		TotalAmount:   x.TotalAmount,
		MyDescription: x.MyDescription,
		IsCancel:      x.IsCancel,
		IsConfirm:     x.IsConfirm,
	}
}

func map_deposit_template(x NewDepositModel) sales.NewDepositTemplate {
	return sales.NewDepositTemplate{
		AllocateId:       x.AllocateId,
		ArCode:           x.ArCode,
		ArId:             x.ArId,
		ArName:           x.ArName,
		ArBillAddress:    x.ArBillAddress,
		ArTelephone:      x.ArTelephone,
		BeforeTaxAmount:  x.BeforeTaxAmount,
		BranchId:         x.BranchId,
		BillType:         x.BillType,
		BankAmount:       x.BankAmount,
		BillBalance:      x.BillBalance,
		ConfirmTime:      x.ConfirmTime,
		ConfirmBy:        x.ConfirmBy,
		CancelBy:         x.CancelBy,
		CancelTime:       x.CancelTime,
		CompanyId:        x.CompanyId,
		CreateBy:         x.CreateBy,
		CreateTime:       x.CreateTime,
		CreditDay:        x.CreditDay,
		ChqAmount:        x.ChqAmount,
		CreditcardAmount: x.CreditcardAmount,
		CashAmount:       x.CashAmount,
		DocNo:            x.DocNo,
		DocDate:          x.DocDate,
		DepartId:         x.DepartId,
		DueDate:          x.DueDate,
		EditTime:         x.EditTime,
		EditBy:           x.EditBy,
		Id:               x.Id,
		IsCancel:         x.IsCancel,
		IsConfirm:        x.IsConfirm,
		IsReturnMoney:    x.IsReturnMoney,
		JobNo:            x.JobNo,
		MyDescription:    x.MyDescription,
		NetAmount:        x.NetAmount,
		ProjectId:        x.ProjectId,
		RefNo:            x.RefNo,
		SaleName:         x.SaleName,
		SaleId:           x.SaleId,
		SaleCode:         x.SaleCode,
		ScgId:            x.ScgId,
		TaxType:          x.TaxType,
		TaxRate:          x.TaxRate,
		TaxAmount:        x.TaxAmount,
		TotalAmount:      x.TotalAmount,
		TaxNo:            x.TaxNo,
		Uuid:             x.Uuid,
	}
}

func map_deposit_subs_template(x NewDepositItemModel) sales.NewDepositItemTemplate {
	return sales.NewDepositItemTemplate{
		BarCode:         x.BarCode,
		DiscountWord:    x.DiscountWord,
		DiscountAmount:  x.DiscountAmount,
		IsCancel:        x.IsCancel,
		Id:              x.Id,
		ItemCode:        x.ItemCode,
		ItemName:        x.ItemName,
		ItemId:          x.ItemId,
		ItemAmount:      x.ItemAmount,
		ItemDescription: x.ItemDescription,
		LineNumber:      x.LineNumber,
		Price:           x.Price,
		PackingRate1:    x.PackingRate1,
		Qty:             x.Qty,
		QuoId:           x.QuoId,
		RefNo:           x.RefNo,
		RemainQty:       x.RemainQty,
		RefLineNumber:   x.RefLineNumber,
		ShelfCode:       x.ShelfCode,
		SOId:            x.SOId,
		SORefNo:         x.SORefNo,
		UnitCode:        x.UnitCode,
		WHCode:          x.WHCode,
	}
}
func map_bank_template(x BankpayModel) sales.BankpayTemplate {
	return sales.BankpayTemplate{
		Id:           x.Id,
		RefId:        x.RefId,
		BankAccount:  x.BankAccount,
		BankName:     x.BankName,
		BankAmount:   x.BankAmount,
		Activestatus: x.Activestatus,
		CreateBy:     x.CreateBy,
		EditBy:       x.EditBy,
	}
}
func map_crd_template(x CreditCardTypeModel) sales.CreditCardTypeTemplate {
	return sales.CreditCardTypeTemplate{
		Id:                 x.Id,
		CreditCardTypeName: x.CreditcardTypeName,
	}
}
func map_deposit_crd_template(x CreditCardModel) sales.CreditCardTemplate {
	return sales.CreditCardTemplate{
		Amount:       x.Amount,
		BookId:       x.BookId,
		BankBranchId: x.BankBranchId,
		BankId:       x.BankId,
		ConfirmNo:    x.ConfirmNo,
		CreditType:   x.CreditType,
		ChargeAmount: x.ChargeAmount,
		CreditCardNo: x.CreditCardNo,
		Description:  x.Description,
		DueDate:      x.DueDate,
		Id:           x.Id,
		ReceiveDate:  x.ReceiveDate,
		RefId:        x.RefId,
	}
}

func map_deposit_chq_template(x ChqInModel) sales.ChqInTemplate {
	return sales.ChqInTemplate{
		BookId:       x.BookId,
		BankBranchId: x.BankBranchId,
		BankId:       x.BankId,
		ChqAmount:    x.ChqAmount,
		ChqNumber:    x.ChqNumber,
		ChqStatus:    x.ChqStatus,
		ChqBalance:   x.ChqBalance,
		Description:  x.Description,
		DueDate:      x.DueDate,
		Id:           x.Id,
		ReceiveDate:  x.ReceiveDate,
		RefId:        x.RefId,
	}
}
func (repo *salesRepository) CancelInvoice(req *sales.NewInvoiceTemplate) (resp interface{}, err error) {
	var check_doc_exist int64
	var check_item_strock int64

	now := time.Now()

	req.CancelTime = now.String()

	switch {
	case req.CancelBy == "":
		return nil, errors.New("ไม่ได้ระบุผู้ยกเลิก")
	case req.IsConfirm == 1:
		return nil, errors.New("เอกสารถูกอ้างอิงไปแล้ว ไม่สามารถยกเลิกได้")
	case req.IsCancel == 1:
		return nil, errors.New("เอกสารถูกยกเลิกไปแล้ว ไม่สามารถยกเลิกได้")
	}

	sqlexist := `select count(doc_no) as check_exist from ar_invoice where id = ?`
	fmt.Println("DocNo Id", req.Id)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	fmt.Println("check_doc_exist", check_doc_exist)

	if check_doc_exist != 0 {
		fmt.Println("Cancel")

		sql := `Update ar_invoice set is_cancel=1,cancel_by=?,cancel_time=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.CancelBy, req.CancelTime, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
		fmt.Println("Cancel", req.SumOfDeposit)
		updatedep := `update Customer_copy1 set debt_amount = debt_amount-? where code = ?`
		_, err = repo.db.Exec(updatedep, req.SumOfDeposit, req.ArCode)
		if err != nil {
			fmt.Println("error update Customer = ", err.Error())
		}

		rowAffect, err := id.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		fmt.Println("ReqID=", req.Id)

		sql_del_sub := `update ar_invoice_sub set is_cancel = 1 where inv_id = ?`
		_, err = repo.db.Exec(sql_del_sub, req.Id)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}
	sqlitems := `select count(doc_no) as check_exist from ar_invoice_sub where inv_id = ? or doc_no =?`
	errs := repo.db.Get(&check_item_strock, sqlitems, req.Id, req.DocNo)
	if errs != nil {
		fmt.Println("Error = ", errs.Error())
		return nil, errs
	}
	if check_item_strock != 0 {
		var itemcode_ string
		var qty float64
		var qty_item float64
		var wh_id int64
		var wh_code string
		var qty_location float64
		sqlrow, err := repo.db.Query("select item_code,wh_id,qty from ar_invoice_sub where doc_no = ?", req.DocNo)
		defer sqlrow.Close()
		for sqlrow.Next() {
			err = sqlrow.Scan(&itemcode_, &wh_id, &qty)
			if err != nil {
				fmt.Println("Error = ", errs.Error())
				return nil, errs
			}
			fmt.Println(qty, "---")
			if wh_id == 1 {
				wh_code = "S1-A"
			} else if wh_id == 2 {
				wh_code = "S1-B"
			} else if wh_id == 3 {
				wh_code = "S2-A"
			} else if wh_id == 4 {
				wh_code = "S2-B"
			} else {
				wh_code = ""
			}
			//คืน stock
			qelstrock := `select stock_qty from Item_copy1 where code = ?`
			errs := repo.db.Get(&qty_item, qelstrock, &itemcode_)
			if errs != nil {
				fmt.Println("Error = ", errs.Error())
				return nil, errs
			}

			fmt.Println("ค้นหาสิค้า")
			fmt.Println(qty_item, "asd")
			qty_item += qty

			fmt.Println(" คืน stock ", qty_item)
			returnstock := `update Item_copy1 set stock_qty = ? where  code = ?`
			_, errss := repo.db.Exec(returnstock, &qty_item, &itemcode_)
			if errss != nil {
				fmt.Println("Error = ", err.Error())
				return nil, errss
			}
			if wh_code != "" {
				fmt.Println("ค้นหาสิค้า sk")
				sqllocation := `select qty from StockLocation_copy1 where item_code =? And wh_code = ? `
				errs = repo.db.Get(&qty_location, sqllocation, &itemcode_, &wh_code)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}

				qty_location += qty
				fmt.Println("จำนวยสินค้า ", qty_location, qty)
				updatestock := `update StockLocation_copy1 set qty = ? where item_code = ? And wh_code = ?`
				_, errs = repo.db.Exec(updatestock, &qty_location, &itemcode_, &wh_code)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}

			}
		}
	}
	return map[string]interface{}{
		"id":       req.Id,
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
		"ar_code":  req.ArCode,
	}, nil
}
func (repo *salesRepository) Searchcreditcard(req *sales.SearchcreditcardTamplate) (resp interface{}, err error) {
	crd := []CreditCardTypeModel{}
	sql := `select id,creditcardtype_name from creditcard_type`

	fmt.Println(sql)
	err = repo.db.Select(&crd, sql)
	if err != nil {
		fmt.Print("Error", err.Error())
		return nil, err
	}

	ro := []sales.CreditCardTypeTemplate{}

	for _, l := range crd {

		roline := map_crd_template(l)
		ro = append(ro, roline)
	}

	return ro, nil
}
func (repo *salesRepository) CreateInvoice(req *sales.NewInvoiceTemplate) (interface{}, error) {
	var check_doc_exist int64
	var check_item_strock int64
	var inv_id int64
	var deplimit int64
	var sumdep float64

	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")
	fmt.Println("tax rate = ", def.TaxRateDefault)

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	sqlexist := `select count(doc_no) as check_exist from ar_invoice where id = ? or doc_no = ?`
	err := repo.db.Get(&check_doc_exist, sqlexist, req.Id, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	uuid := GenUUID()

	if req.DocDate == "" {
		req.DocDate = DocDate
	}

	due_date := now.AddDate(0, 0, int(req.CreditDay)).Format("2006-01-02")

	req.DueDate = due_date

	fmt.Println("duedate =", req.DueDate)
	fmt.Println("SumOfDeposit", req.SumOfDeposit)
	req.CreateTime = now.String()
	req.EditTime = now.String()
	req.CancelTime = now.String()
	//req.TaxRate = def.TaxRateDefault

	if req.Uuid == "" {
		req.Uuid = uuid
	}

	fmt.Println("Doc UUID = ", req.Uuid)

	if req.TotalAmount != 0 {
		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)
		req.NetDebtAmount = req.TotalAmount
	}

	fmt.Println("check_doc_exist = ", check_doc_exist, sqlexist, req.Id)

	switch {
	case req.BillType == 0 && (req.TotalAmount != (req.SumCashAmount + req.SumCreditAmount + req.SumChqAmount + req.SumBankAmount + req.SumOfDeposit + req.CouponAmount)):
		return nil, errors.New("มูลค่ารวมทั้งหมดไม่ตรงกับมูลค่ารับชำระ")
	}
	if req.BillType == 1 {
		//	เช็ควงเงิน
		sqlcheckdebt := `select count(debt_limit) from Customer_copy1 where code = ? and  debt_limit >= ?`
		err := repo.db.Get(&deplimit, sqlcheckdebt, req.ArCode, req.SumOfDeposit)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}

		if deplimit != 0 {
			if check_doc_exist != 0 {
				sumdep = 0
				sqlgetdespo := `select sum_of_deposit from ar_invoice where doc_no = ?`
				err := repo.db.Get(&sumdep, sqlgetdespo, req.DocNo)
				if err != nil {
					fmt.Println("Error = ", err.Error())
					return nil, err
				}
				fmt.Println(sumdep, "sumofdepo")
				updatedepre := `update Customer_copy1 set debt_amount = debt_amount-? where code = ?`
				_, err = repo.db.Exec(updatedepre, &sumdep, req.ArCode)
				if err != nil {
					fmt.Println("error update Customer = ", err.Error())
				}
			}

			updatedep := `update Customer_copy1 set debt_amount = debt_amount+? where code = ?`
			_, err = repo.db.Exec(updatedep, req.SumOfDeposit, req.ArCode)
			if err != nil {
				fmt.Println("error update Customer = ", err.Error())
			}
		} else {
			return nil, err
		}

	}

	if check_doc_exist == 0 {
		//sql := `insert into ar_invoice(company_id,branch_id,uuid,doc_no,tax_no,bill_type,doc_date,ar_id,ar_code,ar_name,sale_id,sale_code,sale_name,pos_machine_id,period_id,cash_id,tax_type,tax_rate,number_of_item,depart_id,allocate_id,project_id,pos_status,credit_day,due_date,delivery_day,delivery_date,is_confirm,is_condition_send,my_description,so_ref_no,change_amount,sum_cash_amount,sum_credit_amount,sum_chq_amount,sum_bank_amount,sum_of_deposit,sum_on_line_amount,coupon_amount,sum_of_item_amount,discount_word,discount_amount,after_discount_amount,before_tax_amount,tax_amount,total_amount,net_debt_amount,bill_balance,pay_bill_status,pay_bill_amount,delivery_status,receive_name,receive_tel,car_license,is_cancel,is_hold,is_posted,is_credit_note,is_debit_note,gl_status,job_id,job_no,coupon_no,redeem_no,scg_number,scg_id,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		// sql := `insert into ar_invoice(company_id,branch_id,uuid,doc_no,tax_no,bill_type,doc_date,ar_id,ar_code,ar_name,sale_id,sale_code,sale_name,tax_type
		// 	,tax_rate,my_description,so_ref_no,change_amount,sum_cash_amount,sum_credit_amount,sum_chq_amount,sum_bank_amount,sum_of_deposit,sum_on_line_amount
		// 	,coupon_amount,sum_of_item_amount,discount_word,discount_amount,after_discount_amount,before_tax_amount
		// 	,tax_amount,total_amount,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

		sql := `insert into ar_invoice(company_id,branch_id,uuid,doc_no,tax_no,bill_type,doc_date,ar_id,ar_code,ar_name,sale_id,sale_code,sale_name,credit_day,due_date,tax_type,tax_rate,my_description,so_ref_no,change_amount,sum_cash_amount,sum_credit_amount,sum_chq_amount,sum_bank_amount,sum_of_deposit,sum_on_line_amount,coupon_amount,sum_of_item_amount,discount_word,discount_amount,after_discount_amount,before_tax_amount,tax_amount,total_amount,create_by,create_time,job_id) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		fmt.Println(sql)
		//resp, err := repo.db.Exec(sql, req.CompanyId,req.BranchId,req.Uuid,req.DocNo,req.TaxNo,req.BillType,req.DocDate,req.ArId,req.ArCode,req.ArName,req.SaleId,req.SaleCode,req.SaleName,req.PosMachineId,req.PeriodId,req.CashId,req.TaxType,req.TaxRate,req.NumberOfItem,req.DepartId,req.AllocateId,req.ProjectId,req.PosStatus,req.CreditDay,req.DueDate,req.DeliveryDay,req.DeliveryDate,req.IsConfirm,req.IsConditionSend,req.MyDescription,req.SoRefNo,req.ChangeAmount,req.SumCashAmount,req.SumCreditAmount,req.SumChqAmount,req.SumBankAmount,req.SumOfDeposit,req.SumOnLineAmount,req.CouponAmount,req.SumOfItemAmount,req.DiscountWord,req.DiscountAmount,req.AfterDiscountAmount,req.BeforeTaxAmount,req.TaxAmount,req.TotalAmount,req.BillBalance,req.PayBillStatus,req.PayBillAmount,req.DeliveryStatus,req.ReceiveName,req.ReceiveTel,req.CarLicense,req.IsCancel,req.IsHold,req.IsPosted,req.IsCreditNote,req.IsDebitNote,req.GlStatus,req.JobId,req.JobNo,req.CouponNo,req.RedeemNo,req.ScgNumber,req.ScgId,req.CreateBy,req.CreateTime)
		resp, err := repo.db.Exec(sql, req.CompanyId, req.BranchId, req.Uuid, req.DocNo, req.TaxNo, req.BillType, req.DocDate, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.CreditDay, req.DueDate, req.TaxType, req.TaxRate, req.MyDescription, req.SoRefNo, req.ChangeAmount, req.SumCashAmount, req.SumCreditAmount, req.SumChqAmount, req.SumBankAmount, req.SumOfDeposit, req.SumOnLineAmount, req.CouponAmount, req.SumOfItemAmount, req.DiscountWord, req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.CreateBy, req.CreateTime,req.JobId)
		fmt.Println(resp, "เพิ่มส่วนhead")
		if err != nil {
			fmt.Println("error = ", err.Error())
		}
		fmt.Println("sql = ", sql)

		id, _ := resp.LastInsertId()

		req.Id = id
		fmt.Println("เพิ่มสินค้านะ", req.Subs)
		fmt.Println("insert credit_card")

		// if len(req.CreditCard) != 0 {
		// 	for _, crd := range req.CreditCard {
		// 		sqlsub_card := `insert into credit_card(company_id, branch_id,ref_uuid, ref_id,ar_id,doc_no, doc_date, credit_card_no, credit_type, confirm_no, amount, charge_amount, description,receive_date, due_date) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		// 		_, err = repo.db.Exec(sqlsub_card, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, crd.CreditCardNo, crd.CreditType, crd.ConfirmNo, crd.Amount, crd.ChargeAmount, "invoice pos", req.DocDate, req.DocDate)
		// 		if err != nil {
		// 			fmt.Println("error insert credit card = ", err.Error())
		// 		}
		// 	}
		// }
		// if len(req.Chq) != 0 {
		// 	for _, chq := range req.Chq {
		// 		sql_chq := `insert into chq_in (company_id,branch_id,uuid,ref_id,ar_id,doc_no,doc_date,chq_number,bank_id,bank_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		// 		_, err := repo.db.Exec(sql_chq, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, chq.ChqNumber, chq.BankId, chq.BankBranchId, chq.ReceiveDate, chq.DueDate, chq.BookId, chq.ChqStatus, chq.ChqAmount, chq.ChqAmount, chq.Description, req.CreateBy, req.CreateTime)
		// 		if err != nil {
		// 			fmt.Println("error insert check = ", err.Error())
		// 		}
		// 	}
		// }
	} else {
		sqlgetid := `select id from ar_invoice where doc_no = ?`
		err := repo.db.Get(&inv_id, sqlgetid, req.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
		req.Id = inv_id

		//	sql :=	`update ar_invoice company_id=?,branch_id=?,uuid=?,doc_no=?,tax_no=?,bill_type=?,doc_date=?,ar_id=?,ar_code=?,ar_name=?,sale_id=?,sale_code=?,sale_name=?,tax_type=?,tax_rate=?,my_description=?,so_ref_no=?,change_amount=?,sum_cash_amount=?,sum_credit_amount=?,sum_chq_amount=?,sum_bank_amount=?,sum_of_deposit=?,sum_on_line_amount=?,coupon_amount=?,sum_of_item_amount=?,discount_word=?,discount_amount=?,after_discount_amount=?,before_tax_amount=?,tax_amount=?,total_amount=?,create_by=?,create_time=?
		sql := `update ar_invoice set company_id=?,branch_id=?,
		uuid=?,doc_no=?,tax_no=?,bill_type=?,doc_date=?,ar_id=?,
		ar_code=?,ar_name=?,sale_id=?,sale_code=?,sale_name=?,
		tax_type=?,tax_rate=?,my_description=?,so_ref_no=?,change_amount=?,
		sum_cash_amount=?,sum_credit_amount=?,sum_chq_amount=?,sum_bank_amount=?,
		sum_of_deposit=?,sum_on_line_amount=?,coupon_amount=?,sum_of_item_amount=?,
		discount_word=?,discount_amount=?,after_discount_amount=?,before_tax_amount=?,
		tax_amount=?,total_amount=?,edit_by=?,edit_time=?,job_id=? where id = ?`
		resp, err := repo.db.Exec(sql, req.CompanyId,
			req.BranchId, req.Uuid, req.DocNo,
			req.TaxNo, req.BillType, req.DocDate, req.ArId, req.ArCode,
			req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.TaxType,
			req.TaxRate, req.MyDescription, req.SoRefNo, req.ChangeAmount,
			req.SumCashAmount, req.SumCreditAmount, req.SumChqAmount,
			req.SumBankAmount, req.SumOfDeposit, req.SumOnLineAmount,
			req.CouponAmount, req.SumOfItemAmount, req.DiscountWord,
			req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount,
			req.TaxAmount, req.TotalAmount, req.CreateBy, req.EditTime, req.JobId, req.Id)

		if err != nil {
			fmt.Println("error = ", err.Error())
		}
		fmt.Println("sql = ", sql)
		//
		rowAffect, err := resp.RowsAffected()
		fmt.Println("Row Affect = ", rowAffect)

		//rowAffect, err := resp.RowsAffected()
		//fmt.Println("Row Affect = ", rowAffect)

	}
	fmt.Println("ค้นหาสิค้า")
	sqlitems := `select count(doc_no) as check_exist from ar_invoice_sub where inv_id = ? or doc_no =?`
	errs := repo.db.Get(&check_item_strock, sqlitems, req.Id, req.DocNo)
	if errs != nil {
		fmt.Println("Error = ", errs.Error())
		return nil, errs
	}
	if check_item_strock == 0 {

		for _, sub := range req.Subs {
			fmt.Println("ค้นหาสิค้า", sub.Location)
			if sub.Location != "" {
				var qty float64
				fmt.Println("ค้นหาสิค้า", req.DocNo)
				qelstrock := `select qty from StockLocation_copy1 where item_code = ? And wh_code = ?`
				errs := repo.db.Get(&qty, qelstrock, sub.ItemCode, sub.Location)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}
				fmt.Println("ค้นหาสิค้า")
				fmt.Println(qty, "asd")
				qty -= sub.Qty
				// ตัด stock ไหม่
				fmt.Println(" ตัด stock ไหม่", qty)
				Updatenew := `update StockLocation_copy1 set qty = ? where item_code = ? And wh_code = ?`
				_, errss := repo.db.Exec(Updatenew, &qty, sub.ItemCode, sub.Location)
				if errss != nil {
					fmt.Println("Error = ", err.Error())
					return nil, errss
				}
				fmt.Println(" ตัด stock ไหม่นะ")
			} else {
				var qty float64
				fmt.Println("ค้นหาสิค้า", req.DocNo)
				qelstrock := `select stock_qty from Item_copy1 where code = ?`
				errs := repo.db.Get(&qty, qelstrock, sub.ItemCode)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}
				fmt.Println("ค้นหาสิค้า")
				fmt.Println(qty, "asd")
				qty -= sub.Qty
				// ตัด stock ไหม่
				fmt.Println(" ตัด stock ไหม่", qty)
				Updatenew := `update Item_copy1 set stock_qty = ? where  code = ?`
				_, errss := repo.db.Exec(Updatenew, &qty, sub.ItemCode)
				if errss != nil {
					fmt.Println("Error = ", err.Error())
					return nil, errss
				}
				fmt.Println(" ตัด stock ไหม่นะ")
			}

		}
	} else {
		var itemcode_ string
		var qty float64
		var qty_item float64
		var wh_id int64
		var wh_code string
		var qty_location float64
		sqlrow, err := repo.db.Query("select item_code,wh_id,qty from ar_invoice_sub where doc_no = ?", req.DocNo)
		defer sqlrow.Close()
		for sqlrow.Next() {
			err = sqlrow.Scan(&itemcode_, &wh_id, &qty)
			if err != nil {
				fmt.Println("Error = ", errs.Error())
				return nil, errs
			}
			fmt.Println(qty, "---")
			if wh_id == 1 {
				wh_code = "S1-A"
			} else if wh_id == 2 {
				wh_code = "S1-B"
			} else if wh_id == 3 {
				wh_code = "S2-A"
			} else if wh_id == 4 {
				wh_code = "S2-B"
			} else {
				wh_code = ""
			}
			//คืน stock
			qelstrock := `select stock_qty from Item_copy1 where code = ?`
			errs := repo.db.Get(&qty_item, qelstrock, &itemcode_)
			if errs != nil {
				fmt.Println("Error = ", errs.Error())
				return nil, errs
			}

			fmt.Println("ค้นหาสิค้า")
			fmt.Println(qty_item, "asd")
			qty_item += qty

			fmt.Println(" คืน stock ", qty_item)
			returnstock := `update Item_copy1 set stock_qty = ? where  code = ?`
			_, errss := repo.db.Exec(returnstock, &qty_item, &itemcode_)
			if errss != nil {
				fmt.Println("Error = ", err.Error())
				return nil, errss
			}
			if wh_code != "" {
				fmt.Println("ค้นหาสิค้า sk")
				sqllocation := `select qty from StockLocation_copy1 where item_code =? And wh_code = ? `
				errs = repo.db.Get(&qty_location, sqllocation, &itemcode_, &wh_code)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}

				qty_location += qty
				fmt.Println("จำนวยสินค้า ", qty_location, qty)
				updatestock := `update StockLocation_copy1 set qty = ? where item_code = ? And wh_code = ?`
				_, errs = repo.db.Exec(updatestock, &qty_location, &itemcode_, &wh_code)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}

			}

		}
		for _, sub := range req.Subs {

			//ค้นหาสินค้าที่มี และคืน stock

			// ตัด stock ไหม่
			//ค้นหาสินค้าไหม่

			fmt.Println("ค้นหาสิค้า", req.DocNo)
			sqlstock := `select stock_qty from Item_copy1 where code = ?`
			errs := repo.db.Get(&qty, sqlstock, sub.ItemCode)
			if errs != nil {
				fmt.Println("Error = ", errs.Error())
				return nil, errs
			}
			fmt.Println("ค้นหาสิค้า")
			fmt.Println(qty, "asd")
			qty -= sub.Qty
			// ตัด stock ไหม่
			fmt.Println(" ตัด stock ไหม่", qty)
			Updatestock := `update Item_copy1 set stock_qty = ? where  code = ?`
			_, errss := repo.db.Exec(Updatestock, &qty, sub.ItemCode)
			if errss != nil {
				fmt.Println("Error = ", err.Error())
				return nil, errss
			}
			fmt.Println(" ตัด stock ไหม่นะ")
			if sub.Location != "" {
				var qtys float64
				fmt.Println("ค้นหาสิค้า", req.DocNo)
				qelstrock := `select qty from StockLocation_copy1 where item_code = ? And wh_code = ?`
				errs := repo.db.Get(&qtys, qelstrock, sub.ItemCode, sub.Location)
				if errs != nil {
					fmt.Println("Error = ", errs.Error())
					return nil, errs
				}
				fmt.Println("ค้นหาสิค้า")
				fmt.Println(qty, "asd")
				qtys -= sub.Qty
				// ตัด stock ไหม่
				fmt.Println(" ตัด stock ไหม่", qty)
				Updatenew := `update StockLocation_copy1 set qty = ? where item_code = ? And wh_code = ?`
				_, errss := repo.db.Exec(Updatenew, &qtys, sub.ItemCode, sub.Location)
				if errss != nil {
					fmt.Println("Error = ", err.Error())
					return nil, errss
				}
				fmt.Println(" ตัด stock ไหม่นะ")
			}

		}

	}
	sql_del_sub := `delete from ar_invoice_sub where inv_id = ?`
	_, err = repo.db.Exec(sql_del_sub, req.Id)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	fmt.Println("ReqID=", req.Id)
	var vLineNumber = 0

	fmt.Println(req.Subs)
	for _, sub := range req.Subs {
		var wh_id int64
		if sub.Location == "S1-A" {
			wh_id = 1
		} else if sub.Location == "S1-B" {
			wh_id = 2
		} else if sub.Location == "S2-A" {
			wh_id = 3
		} else if sub.Location == "S2-B" {
			wh_id = 4
		} else {
			wh_id = 0
		}
		fmt.Println("เพิ่มสินค้านะ", sub, sub.ItemCode)
		sqlsub := `INSERT INTO ar_invoice_sub(company_id,branch_id,uuid,inv_id,doc_no,doc_date,ar_id,sale_id,item_id,item_code,
			bar_code,item_name,wh_id,shelf_id,qty,cn_qty,unit_code,price,discount_word_sub,discount_amount_sub,amount,net_amount,average_cost,sum_of_cost,item_decription,is_cancel,is_credit_note,is_debit_note,packing_rate_1,packing_rate_2,ref_no,ref_line_number,line_number) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := repo.db.Exec(sqlsub,
			req.CompanyId,
			req.BranchId,
			req.Uuid,
			req.Id,
			req.DocNo,
			req.DocDate,
			req.ArId,
			req.SaleId,
			sub.Itemid,
			sub.ItemCode,
			sub.BarCode,
			sub.ItemName,
			wh_id,
			sub.ShelfId,
			sub.Qty,
			sub.CnQty,
			sub.UnitCode,
			sub.Price,
			sub.DiscountWord,
			sub.DiscountAmount,
			sub.ItemAmount,
			sub.NetAmount,
			sub.Average_cost,
			sub.SumOfCost,
			sub.ItemDescription,
			sub.IsCancel,
			sub.IsCreditNote,
			sub.IsDebitNote,
			sub.PackingRate1,
			sub.PackingRate2,
			sub.RefNo,
			sub.RefLineNumber,
			sub.LineNumber)

		vLineNumber = vLineNumber + 1
		if err != nil {
			return "Insert SaleOrder Not Success", err
		}
	}

	sql_del_crd := `delete from credit_card where ref_id=? and doc_no=? and company_id=? and branch_id=? `
	_, err = repo.db.Exec(sql_del_crd, req.Id, req.DocNo, req.CompanyId, req.BranchId)
	if err != nil {
		fmt.Println("sql_del = ", err.Error())
	}

	var count_crd_err int64
	var sum_crd_amount float64
	if len(req.CreditCard) != 0 {
		for _, crd := range req.CreditCard {

			sql_del := `delete from credit_card where doc_no=? and ref_id=? and company_id=? and branch_id=? and credit_card_no=? and confirm_no=? and bank_id=?`
			crd_del, _ := repo.db.Exec(sql_del, req.DocNo, req.Id, req.CompanyId, req.BranchId, crd.CreditCardNo, crd.ConfirmNo, crd.BankId)
			if err != nil {
				fmt.Println("sql_del = ", err.Error())
			}
			fmt.Println(crd_del.RowsAffected())

			verify_crd, _ := verify_creditcard(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, crd.CreditCardNo, crd.ConfirmNo, crd.BankId)
			fmt.Println("verify_crd = ", verify_crd)
			if verify_crd == false {
				count_crd_err = count_crd_err + 1
			}

			sum_crd_amount = sum_crd_amount + crd.Amount
		}
		for _, crd := range req.CreditCard {
			sqlsub_card := `insert into credit_card(company_id, branch_id,ref_uuid, ref_id,ar_id,doc_no, doc_date, credit_card_no, credit_type, confirm_no, amount, charge_amount, description,receive_date, due_date) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
			_, err = repo.db.Exec(sqlsub_card, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, crd.CreditCardNo, crd.CreditType, crd.ConfirmNo, crd.Amount, crd.ChargeAmount, "invoice pos", req.DocDate, req.DocDate)
			if err != nil {
				fmt.Println("error insert credit card = ", err.Error())
			}
		}
	}
	var count_chq_err int64
	var sum_chq_amount float64
	sql_del_chq := `delete from chq_in where doc_no = ? and ref_id=? and company_id=? and branch_id=?`
	_, err = repo.db.Exec(sql_del_chq, req.DocNo, req.Id, req.CompanyId, req.BranchId)
	if err != nil {
		fmt.Println("sql_del = ", err.Error())
	}
	fmt.Println("req ID:", req.Id)
	if len(req.Chq) != 0 {
		for _, chq := range req.Chq {
			fmt.Println("UUID Chq= ", req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber)

			sql_del := `delete from chq_in where doc_no = ? and ref_id=? and company_id=? and branch_id=? and bank_id=? and chq_number=?`
			chq_del, _ := repo.db.Exec(sql_del, req.DocNo, req.Id, req.CompanyId, req.BranchId, chq.BankId, chq.ChqNumber)
			fmt.Println(sql_del)
			if err != nil {
				fmt.Println("sql_del = ", err.Error())
			}
			fmt.Println(chq_del)

			verify_chq, _ := verify_chq_in(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber, chq.BankId)
			fmt.Println("verify_chq = ", verify_chq)
			if verify_chq == false {
				count_chq_err = count_chq_err + 1
			}

			sum_chq_amount = sum_chq_amount + chq.ChqAmount
		}

		for _, i_chq := range req.Chq {
			i_chq.Description = "ออกบิลขาย"
			sql_chq := `insert into chq_in (company_id,branch_id,uuid,ref_id,ar_id,doc_no,doc_date,chq_number,bank_id,bank_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			chq, _ := repo.db.Exec(sql_chq, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, i_chq.ChqNumber, i_chq.BankId, i_chq.BankBranchId, i_chq.ReceiveDate, i_chq.DueDate, i_chq.BookId, i_chq.ChqStatus, i_chq.ChqAmount, i_chq.ChqAmount, i_chq.Description, req.CreateBy, req.CreateTime)
			if err != nil {
				return nil, err
			}
			chqRowAffect, err := chq.RowsAffected()
			if err != nil {
				return nil, err
			}
			fmt.Println("Row Affect = ", chqRowAffect)

			i_chq.Id = chqRowAffect
		}
	}

	sql_del_bnk := `delete from bank_transfer where doc_no = ? and ref_id=? and company_id=? and branch_id=?`
	_, err = repo.db.Exec(sql_del_bnk, req.DocNo, req.Id, req.CompanyId, req.BranchId)
	if err != nil {
		fmt.Println("sql_del = ", err.Error())
	}
	fmt.Println("req ID:", req.Id)
	if len(req.BankPay) != 0 {
		for _, bnk := range req.BankPay {
			//fmt.Println("UUID Chq= ", req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber)

			sql_del := `delete from bank_transfer where doc_no = ? and ref_id=? and company_id=? and branch_id=? and bank_account =? `
			chq_del, _ := repo.db.Exec(sql_del, req.DocNo, req.Id, req.CompanyId, req.BranchId, bnk.BankAccount)
			fmt.Println(chq_del)
			if err != nil {
				fmt.Println("sql_del = ", err.Error())
			}
			fmt.Println(chq_del, "deleted")

		}

		for _, bnk := range req.BankPay {
			fmt.Println(bnk, "insert")
			sql_chq := `insert into bank_transfer (company_id,branch_id,uuid,ref_id,doc_no,doc_date,bank_account,bank_name,bank_amount,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?)`
			bnks, err := repo.db.Exec(sql_chq, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.DocNo, req.DocDate, bnk.BankAccount, bnk.BankName, bnk.BankAmount, req.CreateBy, req.CreateTime)
			if err != nil {
				return nil, err
			}
			chqRowAffect, err := bnks.RowsAffected()
			if err != nil {
				return nil, err
			}
			fmt.Println("Row Affect = ", chqRowAffect)

			bnk.Id = chqRowAffect
		}
	}
	// var count_crd_err int64
	// var sum_crd_amount float64

	// fmt.Println("UUID1 = ", req.Uuid)

	// sql_del_crd := `delete from credit_card where ref_uuid=? and ref_id=? and company_id=? and branch_id=? `
	// _, err = repo.db.Exec(sql_del_crd, req.Uuid, req.Id, req.CompanyId, req.BranchId)
	// if err != nil {
	// 	fmt.Println("sql_del = ", err.Error())
	// }

	// count_crd_err = 0
	// if req.SumCreditAmount != 0 {
	// 	for _, crd := range req.CreditCard {

	// 		sql_del := `delete from credit_card where ref_uuid=? and ref_id=? and company_id=? and branch_id=? and credit_card_no=? and confirm_no=? and bank_id=?`
	// 		crd_del, _ := repo.db.Exec(sql_del, req.Uuid, req.Id, req.CompanyId, req.BranchId, crd.CreditCardNo, crd.ConfirmNo, crd.BankId)
	// 		if err != nil {
	// 			fmt.Println("sql_del = ", err.Error())
	// 		}
	// 		fmt.Println(crd_del.RowsAffected())

	// 		verify_crd, _ := verify_creditcard(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, crd.CreditCardNo, crd.ConfirmNo, crd.BankId)
	// 		fmt.Println("verify_crd = ", verify_crd)
	// 		if verify_crd == false {
	// 			count_crd_err = count_crd_err + 1
	// 		}

	// 		sum_crd_amount = sum_crd_amount + crd.Amount
	// 	}

	// 	fmt.Println("count_crd_err", count_crd_err)

	// 	switch {
	// 	case count_crd_err != 0:
	// 		return nil, errors.New("ข้อมูลบัตรเครดิต มีการใช้ไปแล้ว")
	// 	case sum_crd_amount != req.SumCreditAmount:
	// 		return nil, errors.New("มูลค่าบัตรเครดิต ไม่เท่ากับ มูลค่ารายการบัตรเครดิต")
	// 	}

	// 	for _, i_crd := range req.CreditCard {
	// 		i_crd.Description = "ขายสินค้า"
	// 		sql_crd := `insert into credit_card (company_id, branch_id, uuid, ref_id, ar_id, doc_no, doc_date, credit_card_no, credit_type, confirm_no, amount, charge_amount, description, bank_id, bank_branch_id,receive_date,due_date,book_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	// 		crd, _ := repo.db.Exec(sql_crd, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, i_crd.CreditCardNo, i_crd.CreditType, i_crd.ConfirmNo, i_crd.Amount, i_crd.ChargeAmount, i_crd.Description, i_crd.BankId, i_crd.BankBranchId, i_crd.ReceiveDate, i_crd.DueDate, i_crd.BookId)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		crdRowAffect, err := crd.RowsAffected()
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		fmt.Println("Row Affect = ", crdRowAffect)

	// 		i_crd.Id = crdRowAffect
	// 	}
	// }

	// var count_chq_err int64
	// var sum_chq_amount float64

	// sql_del_chq := `delete from chq_in where uuid = ? and ref_id=? and company_id=? and branch_id=?`
	// _, err = repo.db.Exec(sql_del_chq, req.Uuid, req.Id, req.CompanyId, req.BranchId)
	// if err != nil {
	// 	fmt.Println("sql_del = ", err.Error())
	// }

	// count_chq_err = 0
	// if req.SumChqAmount != 0 {
	// 	for _, chq := range req.Chq {
	// 		fmt.Println("UUID Chq= ", req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber)

	// 		sql_del := `delete from chq_in where uuid = ? and ref_id=? and company_id=? and branch_id=? and bank_id=? and chq_number=?`
	// 		chq_del, _ := repo.db.Exec(sql_del, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.BankId, chq.ChqNumber)
	// 		fmt.Println(sql_del)
	// 		if err != nil {
	// 			fmt.Println("sql_del = ", err.Error())
	// 		}
	// 		fmt.Println(chq_del)

	// 		verify_chq, _ := verify_chq_in(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber, chq.BankId)
	// 		fmt.Println("verify_chq = ", verify_chq)
	// 		if verify_chq == false {
	// 			count_chq_err = count_chq_err + 1
	// 		}

	// 		sum_chq_amount = sum_chq_amount + chq.ChqAmount
	// 	}

	// 	switch {
	// 	case count_chq_err != 0:
	// 		return nil, errors.New("ข้อมูลเลขที่เช็ค มีอยู่แล้ว")
	// 	case sum_chq_amount != req.SumChqAmount:
	// 		return nil, errors.New("มูลค่าเช็ค ไม่เท่ากับ มูลค่ารายการเช็ค")
	// 	}

	// 	for _, i_chq := range req.Chq {
	// 		i_chq.Description = "ขายสินค้า"
	// 		sql_chq := `insert into chq_in (company_id,branch_id,uuid,ref_id,ar_id,doc_no,doc_date,chq_number,bank_id,bank_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	// 		chq, _ := repo.db.Exec(sql_chq, req.CompanyId, req.BranchId, req.Uuid, req.Id, req.ArId, req.DocNo, req.DocDate, i_chq.ChqNumber, i_chq.BankId, i_chq.BankBranchId, i_chq.ReceiveDate, i_chq.DueDate, i_chq.BookId, i_chq.ChqStatus, i_chq.ChqAmount, i_chq.ChqAmount, i_chq.Description, req.CreateBy, req.CreateTime)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		chqRowAffect, err := chq.RowsAffected()
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		fmt.Println("Row Affect = ", chqRowAffect)

	// 		i_chq.Id = chqRowAffect
	// 	}
	// }

	return map[string]interface{}{
		"id":       req.Id,
		"doc_no":   req.DocNo,
		"doc_date": req.DocDate,
		"ar_code":  req.ArCode,
	}, nil
}

func (repo *salesRepository) SearchInvoiceById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	i := NewInvoiceModel{}

	sql := `select a.id,a.company_id,a.branch_id,ifnull(a.uuid,'') as uuid,a.doc_no,ifnull(a.tax_no,'') as tax_no,a.bill_type,a.doc_date,a.ar_id,a.ar_code,ifnull(b.name,'') as ar_name,ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone,a.sale_id,a.sale_code,ifnull(a.sale_name,'') sale_name,a.pos_machine_id,a.cash_id,a.tax_type,a.tax_rate,a.number_of_item,ifnull(a.depart_id,'') as depart_id,a.allocate_id,a.project_id,a.pos_status,a.credit_day,ifnull(a.due_date,'') as due_date,a.delivery_day,ifnull(a.delivery_date,'') as delivery_date,a.is_confirm,a.is_condition_send,ifnull(a.my_description,'') as my_description,ifnull(a.so_ref_no,'') as so_ref_no,a.change_amount,a.sum_cash_amount,a.sum_credit_amount,a.sum_chq_amount,a.sum_bank_amount,a.sum_of_deposit,a.sum_on_line_amount,a.coupon_amount,a.sum_of_item_amount,ifnull(a.discount_word,'') as discount_word,a.discount_amount,a.after_discount_amount,a.before_tax_amount,a.tax_amount,a.total_amount,a.net_debt_amount,a.bill_balance,a.pay_bill_status,a.pay_bill_amount,a.delivery_status,ifnull(a.receive_name,'') as receive_name,ifnull(a.receive_tel,'') as receive_tel,ifnull(a.car_license,'') as car_license,a.is_cancel,a.is_hold,a.is_posted,a.is_credit_note,a.is_debit_note,a.gl_status,ifnull(a.job_id,'') as job_id,ifnull(a.job_no,'') as job_no,ifnull(a.coupon_no,'') as coupon_no,ifnull(a.redeem_no,'') as redeem_no,ifnull(a.scg_number,'') as scg_number,ifnull(a.scg_id,'') as scg_id,a.create_by,a.create_time,ifnull(a.edit_by,'') as edit_by,ifnull(a.edit_time,'') as edit_time,ifnull(a.confirm_by,'') as confirm_by,ifnull(a.confirm_time,'') as confirm_time,ifnull(a.cancel_by,'') as cancel_by,ifnull(a.cancel_time,'') as cancel_time,a.cancel_desc_id,ifnull(a.cancel_desc,'') as cancel_desc     from ar_invoice a left join Customer b on a.ar_id = b.id  where a.id=?`
	err = repo.db.Get(&i, sql, req.Id)
	fmt.Println("sql = ", sql)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	inv_resp := map_invoice_template(i)

	subs := []NewInvoiceItemModel{}

	fmt.Println("s.Id =", i.Id)

	sql_sub := `select a.id,a.inv_id,a.item_id,a.item_code,a.item_name,ifnull(a.wh_id,'') as wh_id,ifnull(a.shelf_id,'') as shelf_id,a.qty,a.cn_qty,a.price,discount_word_sub,discount_amount_sub,ifnull(a.unit_code,'') as unit_code,ifnull(a.bar_code,'') as bar_code,ifnull(a.item_decription,'') as item_decription,a.average_cost,a.sum_of_cost,a.amount,a.packing_rate_1,a.line_number,a.is_cancel from ar_invoice_sub a  where inv_id = ? order by a.line_number`
	err = repo.db.Select(&subs, sql_sub, i.Id)
	// fmt.Println("sql_sub = ", sql_sub)
	if err != nil {
		// 	fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	for _, sub := range subs {
		subline := map_invoice_sub_template(sub)
		inv_resp.Subs = append(inv_resp.Subs, subline)
	}

	// return so_resp, nil

	//subs := []NewInvoiceItemModel{}
	//
	//sql_sub := `select id,inv_id,item_id,ifnull(item_code,'') as item_code,ifnull(bar_code,'') as bar_code,ifnull(item_name,'') as item_name,wh_id,shelf_id,qty,cn_qty,ifnull(unit_code,'') as unit_code,price,ifnull(discount_word_sub,'') as discount_word_sub,discount_amount_sub,amount,net_amount,average_cost,sum_of_cost,ifnull(item_decription,'') as item_decription,is_cancel,is_credit_note,is_debit_note,packing_rate_1,packing_rate_2,ifnull(ref_no,'') as ref_no,ref_line_number,line_number from ar_invoice_sub where inv_id = ?`
	//err = repo.db.Select(&subs, sql_sub, req.Id)
	//if err != nil {
	//	fmt.Println("err = ", err.Error())
	//	return resp, err
	//}

	crds := []CreditCardModel{}
	sql_crd := `select id, ref_id, credit_card_no, credit_type, confirm_no, amount, charge_amount, ifnull(description,'') as description, bank_id, bank_branch_id,receive_date,due_date,book_id from credit_card where company_id = ? and branch_id = ? and ref_id=?`
	err = repo.db.Select(&crds, sql_crd, i.CompanyId, i.BranchId, req.Id)
	if err != nil {
		fmt.Println("err sub=1 ", err.Error())
		return resp, err
	}

	for _, crd := range crds {
		crd_line := map_deposit_crd_template(crd)
		inv_resp.CreditCard = append(inv_resp.CreditCard, crd_line)
	}

	chqs := []ChqInModel{}
	sql_chq := `select id,ref_id,chq_number,bank_id,bank_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description from chq_in where company_id = ? and branch_id = ? and ref_id = ? `
	err = repo.db.Select(&chqs, sql_chq, i.CompanyId, i.BranchId, i.Id)
	if err != nil {
		fmt.Println("err sub=2 ", err.Error())
		return resp, err
	}

	for _, chq := range chqs {
		chq_line := map_deposit_chq_template(chq)
		inv_resp.Chq = append(inv_resp.Chq, chq_line)
	}
	bnks := []BankpayModel{}
	sql_bnk := `select id,ref_id,bank_account,bank_name,bank_amount,create_by,create_time from bank_transfer where company_id = ? and branch_id = ? and ref_id = ?`
	err = repo.db.Select(&bnks, sql_bnk, i.CompanyId, i.BranchId, i.Id)
	if err != nil {
		fmt.Println("err sub=3 ", err.Error())
		return resp, err
	}
	for _, bnk := range bnks {
		bnk_line := map_bank_template(bnk)
		inv_resp.BankPay = append(inv_resp.BankPay, bnk_line)
	}
	return inv_resp, nil
}

func map_invoice_template(x NewInvoiceModel) sales.NewInvoiceTemplate {
	return sales.NewInvoiceTemplate{
		AllocateId:          x.AllocateId,
		ArCode:              x.ArCode,
		ArId:                x.ArId,
		ArName:              x.ArName,
		ItemName:            x.ItemName,
		ItemCode:            x.ItemCode,
		ArBillAddress:       x.ArBillAddress,
		ArTelephone:         x.ArTelephone,
		BeforeTaxAmount:     x.BeforeTaxAmount,
		BranchId:            x.BranchId,
		BillType:            x.BillType,
		DocType:             x.DocType,
		SumBankAmount:       x.SumBankAmount,
		BillBalance:         x.BillBalance,
		ConfirmTime:         x.ConfirmTime,
		ConfirmBy:           x.ConfirmBy,
		CancelBy:            x.CancelBy,
		CancelTime:          x.CancelTime,
		CompanyId:           x.CompanyId,
		CreateBy:            x.CreateBy,
		CreateTime:          x.CreateTime,
		CreditDay:           x.CreditDay,
		SumChqAmount:        x.SumChqAmount,
		SumCreditAmount:     x.SumCreditAmount,
		SumCashAmount:       x.SumCashAmount,
		DocNo:               x.DocNo,
		DocDate:             x.DocDate,
		DepartId:            x.DepartId,
		DueDate:             x.DueDate,
		EditTime:            x.EditTime,
		EditBy:              x.EditBy,
		Id:                  x.Id,
		IsCancel:            x.IsCancel,
		IsConfirm:           x.IsConfirm,
		IsCreditNote:        x.IsCreditNote,
		IsDebitNote:         x.IsDebitNote,
		JobNo:               x.JobNo,
		MyDescription:       x.MyDescription,
		NetDebtAmount:       x.NetDebtAmount,
		ProjectId:           x.ProjectId,
		SoRefNo:             x.SoRefNo,
		SaleName:            x.SaleName,
		SaleId:              x.SaleId,
		SaleCode:            x.SaleCode,
		ScgId:               x.ScgId,
		TaxType:             x.TaxType,
		TaxRate:             x.TaxRate,
		TaxAmount:           x.TaxAmount,
		TotalAmount:         x.TotalAmount,
		TaxNo:               x.TaxNo,
		Uuid:                x.Uuid,
		AfterDiscountAmount: x.AfterDiscountAmount,
		CarLicense:          x.CarLicense,
		CashId:              x.CashId,
		ChangeAmount:        x.ChangeAmount,
		CouponAmount:        x.CouponAmount,
		CancelDesc:          x.CancelDesc,
		CancelDescId:        x.CancelDescId,
		CouponNo:            x.CouponNo,
		DiscountAmount:      x.DiscountAmount,
		DiscountWord:        x.DiscountWord,
		DeliveryDate:        x.DeliveryDate,
		DeliveryDay:         x.DeliveryDay,
		DeliveryStatus:      x.DeliveryStatus,
		GlStatus:            x.GlStatus,
		IsConditionSend:     x.IsConditionSend,
		IsPosted:            x.IsPosted,
		IsHold:              x.IsHold,
		JobId:               x.JobId,
		NumberOfItem:        x.NumberOfItem,
		PosMachineId:        x.PosMachineId,
		PosStatus:           x.PosStatus,
		PeriodId:            x.PeriodId,
		PayBillAmount:       x.PayBillAmount,
		PayBillStatus:       x.PayBillStatus,
		ReceiveName:         x.ReceiveName,
		ReceiveTel:          x.ReceiveTel,
		RedeemNo:            x.RedeemNo,
		SumOfItemAmount:     x.SumOfItemAmount,
		ScgNumber:           x.ScgNumber,
		SumOfDeposit:        x.SumOfDeposit,
		SumOnLineAmount:     x.SumOnLineAmount,
		//Subs:                subs,
		//RecMoney:            recmoneys,
		//CreditCard:          crds,
		//Chq:                 chqs,
	}
}

// searchinvoice by keyword

type SearchInvModel struct {
	Id            int64   `db:"id"`
	DocNo         string  `db:"doc_no"`
	DocDate       string  `db:"doc_date"`
	Doctype       string  `db:"doc_type"`
	ArCode        string  `db:"ar_code"`
	ArName        string  `db:"ar_name"`
	SaleCode      string  `db:"sale_code"`
	SaleName      string  `db:"sale_name"`
	MyDescription string  `db:"my_description"`
	TotalAmount   float64 `db:"total_amount"`
	IsCancel      int     `db:"is_cancel"`
	IsConfirm     int     `db:"is_confirm"`
}

func map_doc_invoic_template(x SearchInvModel) sales.SearchInvTemplate {
	return sales.SearchInvTemplate{
		Id:            x.Id,
		DocNo:         x.DocNo,
		DocDate:       x.DocDate,
		ArCode:        x.ArCode,
		ArName:        x.ArName,
		SaleCode:      x.SaleCode,
		SaleName:      x.SaleName,
		TotalAmount:   x.TotalAmount,
		MyDescription: x.MyDescription,
		Doctype:       x.Doctype,
		IsCancel:      x.IsCancel,
		IsConfirm:     x.IsConfirm,
	}
}

func (repo *salesRepository) SearchInvoiceByKeyword(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {

	var sql string

	d := []SearchInvModel{}

	sql = `select a.id,a.doc_no,a.doc_date,a.doc_type,a.ar_code,a.ar_name,a.sale_code,a.sale_name,ifnull(a.my_description,'') as my_description,
		a.total_amount,a.is_cancel,a.is_confirm 
		from ar_invoice a 
		where doc_type = 0`
	err = repo.db.Select(&d, sql)

	fmt.Println("sql = ", sql, req.Keyword)
	if err != nil {
		fmt.Println("errsss = ", err.Error())
		return resp, err
	}

	dp := []sales.SearchInvTemplate{}

	for _, dep := range d {
		dpline := map_doc_invoic_template(dep)
		dp = append(dp, dpline)
	}

	return dp, nil
}

func map_searchitem_template(x NewSearchItemModel) sales.NewSearchItemTemplate {
	return sales.NewSearchItemTemplate{
		Id:              x.Id,
		DocDate:         x.DocDate,
		DocNo:           x.DocNo,
		ItemId:          x.ItemId,
		ItemCode:        x.ItemCode,
		ItemName:        x.ItemName,
		BarCode:         x.BarCode,
		UnitICode:       x.UnitICode,
		WhId:            x.WhId,
		ShelfId:         x.ShelfId,
		Price:           x.Price,
		Qty:             x.Qty,
		CnQty:           x.CnQty,
		ItemDescription: x.ItemDescription,
		IsCreditNote:    x.IsCreditNote,
		IsDebitNote:     x.IsDebitNote,
		PackingRate1:    x.PackingRate1,
		PackingRate2:    x.PackingRate2,
		SoRefNo:         x.SoRefNo,
		AverageCost:     x.AverageCost,
		SumOfCost:       x.SumOfCost,
		RefLineNumber:   x.RefLineNumber,
		LineNumber:      x.LineNumber,
		ArName:          x.ArName,
		ArCode:          x.ArCode,
		ArId:            x.ArId,
		Name:            x.Name,
		MyDescription:   x.MyDescription,
		NId:             x.NId,
		NDocNo:          x.NDocNo,
		NDocDate:        x.NDocDate,
		NItemId:         x.NItemId,
		NArId:           x.NArId,
		NBarCode:        x.NBarCode,
		NItemCode:       x.NItemCode,
		NItemName:       x.NItemName,
		NUnitCode:       x.NUnitCode,
		NQty:            x.NQty,
		NPrice:          x.NPrice,
		NArName:         x.NArName,
		DiscountWord:    x.DiscountWord,
		NDiscountWord:   x.NDiscountWord,
		NMyDescription:  x.NMyDescription,
	}

}

func (repo *salesRepository) SearchSaleByItem(req *sales.SearchByItemTemplate) (resp interface{}, err error) {
	var sql string
	d := []NewSearchItemModel{}
	if req.Name == "" && req.ItemCode == "" {
		fmt.Println("No Data")
	} else {
		switch {
		case req.Page == "invoice":

			sql = `select a.id, ifnull(a.doc_date,'') as doc_date, ifnull(a.doc_no,'') as doc_no,a.ar_id, a.ar_name,a.my_description,
			b.unit_code, b.qty, b.price, ifnull(b.item_code,'') as item_code, b.ar_id, ifnull(b.item_name,'') as item_name,b.ar_id, ifnull(b.discount_word_sub,'') as discount_word_sub
			from ar_invoice a left join ar_invoice_sub b on a.ar_id = b.ar_id
			where a.ar_name like concat(?) and b.item_code like concat(?)
			order by a.Id desc limit 20`
			err = repo.db.Select(&d, sql, req.Name, req.ItemCode)
		case req.Page == "quotation":
			sql = `select a.Id, ifnull(a.DocDate,'') as DocDate, ifnull(a.DocNo,'') as DocNo, a.ArId, a.ArName, a.MyDescription,
			b.UnitCode, b.Qty, b.Price, ifnull(b.ItemCode,'') as ItemCode, b.ArId, ifnull(b.ItemName,'') as ItemName,b.ArId, ifnull(b.DiscountWord,'') as DiscountWord
			from Quotation a left join QuotationSub b on a.ArId = b.ArId
			where a.ArName like concat(?) and b.ItemCode like concat(?) 
			order by a.Id desc limit 20`
			err = repo.db.Select(&d, sql, req.Name, req.ItemCode)
		case req.Page == "saleorder":
			sql = `select a.Id, ifnull(a.DocDate,'') as DocDate, ifnull(a.DocNo,'') as DocNo, a.ArId, a.ArName, a.MyDescription,
			b.UnitCode, b.Qty, b.Price, ifnull(b.ItemCode,'') as ItemCode, b.ArId, ifnull(b.ItemName,'') as ItemName,b.ArId, ifnull(b.DiscountWord,'') as DiscountWord
			from SaleOrder a left join SaleOrderSub b on a.ArId = b.ArId
			where a.ArName like concat(?) and b.ItemCode like concat(?) 
			order by a.Id desc limit 20`
			err = repo.db.Select(&d, sql, req.Name, req.ItemCode)
		}
	}
	fmt.Println("sql = ", sql, req.ItemCode)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	dp := []sales.NewSearchItemTemplate{}
	for _, dep := range d {
		dpline := map_searchitem_template(dep)
		dp = append(dp, dpline)
	}

	return dp, nil
}

func (repo *salesRepository) SearchHisByKeyword(req *sales.SearchByKeywordTemplate) (resp interface{}, err error) {
	var sql string
	d := []SearchInvModel{}
	if req.Keyword == "" {
		sql = `select a.id, a.doc_no,a.doc_date, a.doc_type ,a.ar_code,a.ar_name,a.sale_code,a.sale_name, ifnull(a.my_description,'') as my_description,a.total_amount, a.is_cancel,a.is_confirm 
		from ar_invoice a
		order by a.id desc limit 30`
		err = repo.db.Select(&d, sql)
	} else {
		sql = `select a.id, a.doc_no,a.doc_date, a.doc_type ,a.ar_code,a.ar_name,a.sale_code,a.sale_name, ifnull(a.my_description,'') as my_description,a.total_amount, a.is_cancel,a.is_confirm 
		from ar_invoice a
		where a.doc_no like  concat(?,'%') or a.ar_code like  concat(?,'%') or a.ar_name like  concat(?,'%') 
		order by a.id desc limit 30`
		err = repo.db.Select(&d, sql, req.Keyword, req.Keyword, req.Keyword)
	}
	fmt.Println("sql = ", sql, req.Keyword)
	if err != nil {
		fmt.Println("errsss = ", err.Error())
		return resp, err
	}

	dp := []sales.SearchInvTemplate{}

	for _, dep := range d {
		dpline := map_doc_invoic_template(dep)
		dp = append(dp, dpline)
	}
	return dp, nil
}

func map_hiscustomer_template(x NewSearchHisCustomerModel) sales.NewSearchHisCustomerTemplate {
	return sales.NewSearchHisCustomerTemplate{
		Id:           x.Id,
		DocDate:      x.DocDate,
		DocNo:        x.DocNo,
		ArName:       x.ArName,
		ArCode:       x.ArCode,
		ArId:         x.ArId,
		SaleName:     x.SaleName,
		TotalAmount:  x.TotalAmount,
		NId:          x.NId,
		NDocNo:       x.NDocNo,
		NDocDate:     x.NDocDate,
		NArId:        x.NArId,
		NArName:      x.NArName,
		NSaleName:    x.NSaleName,
		NTotalAmount: x.NTotalAmount,
	}

}

type NewSearchHisCustomerModel struct {
	Id           int64   `db:"id"`
	DocDate      string  `db:"doc_date"`
	DocNo        string  `db:"doc_no"`
	ArName       string  `db:"ar_name"`
	ArCode       string  `db:"ar_code"`
	ArId         int64   `db:"ar_id"`
	SaleName     string  `db:"sale_name"`
	TotalAmount  float64 `db:"total_amount"`
	NId          int64   `db:"Id"`
	NDocNo       string  `db:"DocNo"`
	NDocDate     string  `db:"DocDate"`
	NArId        int64   `db:"ArId"`
	NArName      string  `db:"ArName"`
	NSaleName    string  `db:"SaleName"`
	NTotalAmount int64   `db:"TotalAmount"`
}

func (repo *salesRepository) SearchHisCustomer(req *sales.SearchHisCustomerTemplate) (resp interface{}, err error) {
	var sql string
	d := []NewSearchHisCustomerModel{}
	switch {
	case req.Page == "invoice":
		sql = `select a.id, ifnull(a.doc_date,'') as doc_date, ifnull(a.doc_no,'') as doc_no, a.ar_id, a.ar_name, a.sale_name , a.total_amount 
		from ar_invoice a 
		where a.ar_code like concat(?) 
		order by a.id desc limit 20`
		err = repo.db.Select(&d, sql, req.ArCode)
	case req.Page == "quotation" || req.Page == "saleorder":
		sql = `select a.Id, ifnull(a.DocDate,'') as DocDate, ifnull(a.DocNo,'') as DocNo, a.ArId, a.ArName, a.SaleName , a.TotalAmount 
		from SaleOrder a 
		where a.ArCode like concat(?) 
		order by a.Id desc limit 20`
		err = repo.db.Select(&d, sql, req.ArCode)
	}

	fmt.Println("sql = ", sql, req.ArCode)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	dp := []sales.NewSearchHisCustomerTemplate{}
	for _, dep := range d {
		dpline := map_hiscustomer_template(dep)
		dp = append(dp, dpline)
	}

	return dp, nil
}
