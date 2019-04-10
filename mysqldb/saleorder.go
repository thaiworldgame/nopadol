package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type BCSaleOrder struct {
	Id                  int64            `json:"id" db:"Id"`
	UUID                string           `json:"uuid" db:"UUID"`
	DocNo               string           `json:"doc_no" db:"DocNo"`
	DocDate             string           `json:"doc_date" db:"DocDate"`
	CompanyId           int64            `json:"company_id" db:"CompanyId"`
	BranchId            int64            `json:"branch_id" db:"BranchId"`
	DocType             int64            `json:"doc_type" db:"DocType"`
	ArId                int64            `json:"ar_id" db:"ArId"`
	ArCode              string           `json:"ar_code" db:"ArCode"`
	ArName              string           `json:"ar_name" db:"ArName"`
	ArBillAddress       string           `json:"ar_bill_address" db:"ArBillAddress"`
	ArTelephone         string           `json:"ar_telephone" db:"ArTelephone"`
	SaleId              int64            `json:"sale_id" db:"SaleId"`
	SaleCode            string           `json:"sale_code" db:"SaleCode"`
	SaleName            string           `json:"sale_name" db:"SaleName"`
	BillType            int64            `json:"bill_type" db:"BillType"`
	TaxType             int64            `json:"tax_type" db:"TaxType"`
	TaxRate             float64          `json:"tax_rate" db:"TaxRate"`
	DepartId            int64            `json:"depart_id" db:""DepartId`
	RefNo               string           `json:"ref_no" db:"RefNo"`
	IsConfirm           int64            `json:"is_confirm" db:"IsConfirm"`
	BillStatus          int64            `json:"bill_status" db:"BillStatus"`
	HoldingStatus       int64            `json:"holding_status" db:"HoldingStatus"`
	CreditDay           int64            `json:"credit_day" db:"CreditDay"`
	DueDate             string           `json:"due_date" db:"DueDate"`
	DeliveryDay         int64            `json:"delivery_day" db:"DeliveryDay"`
	DeliveryDate        string           `json:"delivery_date" db:"DeliveryDate"`
	IsConditionSend     int64            `json:"is_condition_send" db:"IsConditionSend"`
	DeliveryAddressId   int64            `json:"delivery_address_id" db:"DeliveryAddressId"`
	CarLicense          string           `json:"car_license" db:"CarLicense"`
	PersonReceiveTel    string           `json:"person_receive_tel" db:"PersonReceiveTel"`
	MyDescription       string           `json:"my_description" db:"MyDescription"`
	SumOfItemAmount     float64          `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	DiscountWord        string           `json:"discount_word" db:"DiscountWord"`
	DiscountAmount      float64          `json:"discount_amount" db:"DiscountAmount"`
	AfterDiscountAmount float64          `json:"after_discount_amount" db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64          `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount           float64          `json:"tax_amount" db:"TaxAmount"`
	TotalAmount         float64          `json:"total_amount" db:"TotalAmount"`
	NetDebtAmount       float64          `json:"net_debt_amount" db:"NetDebtAmount"`
	ProjectId           int64            `json:"project_id" db:"ProjectId"`
	ProjectCode         string           `json:"project_code"`
	AllocateId          int64            `json:"allocate_id" db:"AllocateId"`
	AllocateCode        string           `json:"allocate_code"`
	JobId               string           `json:"job_id" db:"JobId"`
	IsCancel            int64            `json:"is_cancel" db:"IsCancel"`
	CreateBy            string           `json:"create_by" db:"CreateBy"`
	CreateTime          string           `json:"create_time" db:"CreateTime"`
	EditBy              string           `json:"edit_by" db:"EditBy"`
	EditTime            string           `json:"edit_time" db:"EditTime"`
	ConfirmBy           string           `json:"confirm_by" db:"ConfirmBy"`
	ConfirmTime         string           `json:"confirm_time" db:"CancelBy"`
	ApproveBy           string           `json:"approve_by" db:"ApproveBy"`
	ApproveTime         string           `json:"approve_time" db:"ApproveTime"`
	CancelBy            string           `json:"cancel_by" db:"CancelBy"`
	CancelTime          string           `json:"cancel_time" db:"CancelTime"`
	Subs                []BCSaleOrderSub `json:"subs"`
}

type BCSaleOrderSub struct {
	Id              int64   `json:"id" db:"Id"`
	SoUUID          string  `json:"so_uuid" db:"SoUUID"`
	SOId            int64   `json:"so_id" db:"SOId"`
	ItemId          int64   `json:"item_id" db:"ItemId"`
	ItemCode        string  `json:"item_code" db:"ItemCode"`
	BarCode         string  `json:"bar_code" db:"BarCode"`
	ItemName        string  `json:"item_name" db:"ItemName"`
	WHCode          string  `json:"wh_code" db:"WHCode"`
	ShelfCode       string  `json:"shelf_code" db:"ShelfCode"`
	Qty             float64 `json:"qty" db:"Qty"`
	RemainQty       float64 `json:"remain_qty" db:"RemainQty"`
	Price           float64 `json:"price" db:"Price"`
	DiscountWord    string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64 `json:"discount_amount" db:"DiscountAmount"`
	UnitCode        string  `json:"unit_code" db:"UnitCode"`
	ItemAmount      float64 `json:"item_amount" db:"ItemAmount"`
	ItemDescription string  `json:"item_description" db:"ItemDescription"`
	StockType       int64   `json:"stock_type" db:"StockType"`
	AverageCost     float64 `json:"average_cost" db:"AverageCost"`
	SumOfCost       float64 `json:"sum_of_cost" db:"SumOfCost"`
	PackingRate1    float64 `json:"packing_rate_1" db:"PackingRate1"`
	RefNo           string  `json:"ref_no" db:"RefNo"`
	QuoId           int64   `json:"quo_id" db:"QuoId"`
	LineNumber      int     `json:"line_number" db:"LineNumber"`
	RefLineNumber   int64   `json:"ref_line_number" db:"RefLineNumber"`
	IsCancel        int64   `json:"is_cancel" db:"IsCancel"`
}

func (s *BCSaleOrder) getSOByDocNo(db *sqlx.DB) error {
	fmt.Println("SaleOrder DocNo = ", s.DocNo)
	sql := `select a.Id,ifnull(a.uuid,'') as UUID,a.DocNo,ifnull(a.DocDate,'') as DocDate,a.CompanyId,a.BranchId,a.DocType,a.BillType,a.TaxType,a.ArId,ifnull(a.ArCode,'') as ArCode,ifnull(a.ArName,'') as ArName,a.SaleId,ifnull(a.SaleCode,'') as SaleCode,ifnull(a.SaleName,'') as SaleName,a.DepartId,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.TaxRate,a.IsConfirm,ifnull(a.MyDescription,'') as MyDescription,a.BillStatus,a.HoldingStatus,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.IsCancel,a.IsConditionSend,a.DeliveryAddressId,ifnull(a.CarLicense,'') as CarLicense,ifnull(a.PersonReceiveTel,'') as PersonReceiveTel,ifnull(a.JobId,'') as JobId,a.ProjectId,ifnull(d.code,'') as ProjectCode,a.AllocateId,ifnull(e.code,'') as AllocateCode,ifnull(a.CreateBy,'') as CreateBy,a.CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime, ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime, ifnull(a.ConfirmBy,'') as ConfirmBy,ifnull(a.ConfirmTime,'') as ConfirmTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from SaleOrder a left join Customer b on a.ArId = b.id left join Department c on a.DepartId = c.Id left join Project d on a.ProjectId = d.Id left join Allocate e on a.AllocateId = e.id    where a.DocNo = ?`
	rs := db.QueryRow(sql, s.DocNo)
	err := rs.Scan(&s.Id, &s.UUID, &s.DocNo, &s.DocDate, &s.CompanyId, &s.BranchId, &s.DocType, &s.BillType, &s.TaxType, &s.ArId, &s.ArCode, &s.ArName, &s.SaleId, &s.SaleCode, &s.SaleName, &s.DepartId, &s.CreditDay, &s.DueDate, &s.DeliveryDay, &s.DeliveryDate, &s.TaxRate, &s.IsConfirm, &s.MyDescription, &s.BillStatus, &s.HoldingStatus, &s.SumOfItemAmount, &s.DiscountWord, &s.DiscountAmount, &s.AfterDiscountAmount, &s.BeforeTaxAmount, &s.TaxAmount, &s.TotalAmount, &s.NetDebtAmount, &s.IsCancel, &s.IsConditionSend, &s.DeliveryAddressId, &s.CarLicense, &s.PersonReceiveTel, &s.JobId, &s.ProjectId, &s.ProjectCode, &s.AllocateId, &s.AllocateCode, &s.CreateBy, &s.CreateTime, &s.EditBy, &s.EditTime, &s.CancelBy, &s.CancelTime, &s.ConfirmBy, &s.ConfirmTime, &s.ArBillAddress, &s.ArTelephone)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return err
	}

	sql_sub := `select Id,ifnull(so_uuid,'') as SoUUID,SOId,ItemId,ifnull(ItemCode,'') as ItemCode,ifnull(BarCode,'') as BarCode,ifnull(ItemName,'') as ItemName,ifnull(WhCode,'') as WHCode,ifnull(ShelfCode,'') as ShelfCode,Qty,RemainQty,ifnull(UnitCode,'') as UnitCode,Price,ifnull(DiscountWord,'') as DiscountWord,DiscountAmount,ItemAmount,ifnull(ItemDescription,'') as ItemDescription,StockType,AverageCost,SumOfCost,ifnull(RefNo,'') as RefNo,QuoId,IsCancel,PackingRate1,RefLineNumber,LineNumber from SaleOrderSub  where SoId = ? order by linenumber`
	err = db.Select(&s.Subs, sql_sub, s.Id)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return err
	}

	return nil
}
