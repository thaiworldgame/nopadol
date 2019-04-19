package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BCQuotation struct {
	Id                  int64            `db:"Id" ่ json:"id"`
	UUID                string           `db:"UUID" json:"uuid"`
	DocNo               string           `json:"doc_no" db:"DocNo"`
	DocDate             string           `db:"DocDate" json:"doc_date"`
	CompanyId           int64            `db:"CompanyId" json:"company_id"`
	BranchId            int64            `db:"BranchId" json:"branch_id"`
	DocType             int64            `db:"DocType" json:"doc_type"`
	ArId                int64            `db:"ArId" json:"ar_id"`
	ArCode              string           `db:"ArCode" json:"ar_code"`
	ArName              string           `json:"ar_name" db:"ArName" `
	ArBillAddress       string           `db:"ArBillAddress" json:"ar_bill_address"`
	ArTelephone         string           `db:"ArTelephone" json:"ar_telephone"`
	SaleId              int64            `db:"SaleId" json:"sale_id"`
	SaleCode            string           `db:"SaleCode" json:"sale_code"`
	SaleName            string           `db:"SaleName" json:"sale_name"`
	BillType            int64            `db:"BillType" json:"bill_type"`
	TaxType             int64            `db:"TaxType" json:"tax_type"`
	TaxRate             float64          `db:"TaxRate" json:"tax_rate"`
	DepartId            int64            `db:"DepartId" json:"depart_id"`
	DepartCode          string           `json:"depart_code" db:"DepartCode"`
	RefNo               string           `db:"RefNo" json:"ref_no"`
	JobId               string           `db:"JobId" json:"job_id"`
	IsConfirm           int64            `db:"IsConfirm" json:"is_confirm"`
	BillStatus          int64            `db:"BillStatus" json:"bill_status"`
	Validity            int64            `db:"Validity" json:"validity"`
	CreditDay           int64            `db:"CreditDay" json:"credit_day"`
	DueDate             string           `db:"DueDate" json:"due_date"`
	ExpireCredit        int64            `db:"ExpireCredit" json:"expire_credit"`
	ExpireDate          string           `db:"ExpireDate" json:"expire_date"`
	DeliveryDay         int64            `db:"DeliveryDay" json:"delivery_day"`
	DeliveryDate        string           `db:"DeliveryDate" json:"delivery_date"`
	AssertStatus        int64            `db:"AssertStatus" json:"assert_status"`
	IsConditionSend     int64            `db:"IsConditionSend" json:"is_condition_send"`
	MyDescription       string           `db:"MyDescription" json:"my_description"`
	SumOfItemAmount     float64          `db:"SumOfItemAmount" json:"sum_of_item_amount"`
	DiscountWord        string           `db:"DiscountWord" json:"discount_word"`
	DiscountAmount      float64          `db:"DiscountAmount" json:"discount_amount"`
	AfterDiscountAmount float64          `db:"AfterDiscountAmount" json:"after_discount_amount"`
	BeforeTaxAmount     float64          `db:"BeforeTaxAmount" json:"before_tax_amount"`
	TaxAmount           float64          `db:"TaxAmount" json:"tax_amount"`
	TotalAmount         float64          `db:"TotalAmount" json:"total_amount"`
	NetDebtAmount       float64          `db:"NetDebtAmount" json:"net_debt_amount"`
	ProjectId           int64            `db:"ProjectId" json:"project_id"`
	ProjectCode         string           `json:"project_code" db:"ProjectCode"`
	AllocateId          int64            `db:"AllocateId" json:"allocate_id"`
	AllocateCode        string           `json:"allocate_code" db:"AllocateCode"`
	IsCancel            int64            `db:"IsCancel" json:"is_cancel"`
	CreateBy            string           `db:"CreateBy" json:"create_by"`
	CreateTime          string           `db:"CreateTime" json:"create_time"`
	EditBy              string           `db:"EditBy" json:"edit_by"`
	EditTime            string           `db:"EditTime" json:"edit_time"`
	ConfirmBy           string           `db:"ConfirmBy" json:"confirm_by"`
	ConfirmTime         string           `db:"ConfirmTime" json:"confirm_time"`
	CancelBy            string           `db:"CancelBy" json:"cancel_by"`
	CancelTime          string           `db:"CancelTime" json:"cancel_time"`
	Subs                []BCQuotationSub `db:"subs" json:"subs"`
}

type BCQuotationSub struct {
	Id              int64   `json:"id"db:"Id"`
	QuoUUID         string  `json:"quo_uuid"db:"QuoUUID"`
	QuoId           int64   `json:"quo_id"db:"QuoId"`
	ItemId          int64   `json:"item_id"db:"ItemId"`
	ItemCode        string  `json:"item_code"db:"ItemCode"`
	BarCode         string  `json:"bar_code"db:"BarCode"`
	ItemName        string  `json:"item_name"db:"ItemName"`
	Qty             float64 `json:"qty"db:"Qty"`
	RemainQty       float64 `json:"remain_qty"db:"RemainQty"`
	Price           float64 `json:"price"db:"Price"`
	DiscountWord    string  `json:"discount_word"db:"DiscountWord"`
	DiscountAmount  float64 `json:"discount_amount"db:"DiscountAmount"`
	UnitCode        string  `json:"unit_code"db:"UnitCode"`
	ItemAmount      float64 `json:"item_amount"db:"ItemAmount"`
	ItemDescription string  `json:"item_description"db:"ItemDescription"`
	PackingRate1    float64 `json:"packing_rate_1"db:"PackingRate1"`
	IsCancel        int64   `json:"is_cancel"db:"IsCancel"`
	LineNumber      int     `json:"line_number"db:"LineNumber"`
}

func (q *BCQuotation) getByDocNo(db *sqlx.DB) error {
	fmt.Println("QueDocNo = ", q.DocNo)
	sql := `select a.Id,a.CompanyId,a.BranchId,a.DocNo,a.DocDate,a.DocType,a.Validity,a.BillType,a.ArId,a.ArCode,a.ArName,a.SaleId,a.SaleCode,a.SaleName,ifnull(a.DepartId,0) as DepartId,ifnull(c.code,'') as DepartCode,ifnull(a.RefNo,'') as RefNo,ifnull(a.JobId,'') as JobId,a.TaxType,a.IsConfirm,a.BillStatus,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.ExpireCredit,ifnull(a.ExpireDate,'') as ExpireDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.AssertStatus,a.IsConditionSend,ifnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.TaxRate,a.ProjectId,ifnull(d.code,'') as ProjectCode,a.AllocateId,ifnull(e.code,'') as AllocateCode,a.IsCancel,ifnull(a.CreateBy,'') as CreateBy,ifnull(a.CreateTime,'') as CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime,ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from Quotation a left join Customer b on a.ArId = b.id  left join Department c on a.DepartId = c.Id left join Project d on a.ProjectId = d.Id left join Allocate e on a.AllocateId = e.id   where a.DocNo = ?`
	rs := db.QueryRow(sql, q.DocNo)
	err := rs.Scan(&q.Id, &q.CompanyId, &q.BranchId, &q.DocNo, &q.DocDate, &q.DocType, &q.Validity, &q.BillType, &q.ArId, &q.ArCode, &q.ArName, &q.SaleId, &q.SaleCode, &q.SaleName, &q.DepartId, &q.DepartCode, &q.RefNo, &q.JobId, &q.TaxType, &q.IsConfirm, &q.BillStatus, &q.CreditDay, &q.DueDate, &q.ExpireCredit, &q.ExpireDate, &q.DeliveryDay, &q.DeliveryDate, &q.AssertStatus, &q.IsConditionSend, &q.MyDescription, &q.SumOfItemAmount, &q.DiscountWord, &q.DiscountAmount, &q.AfterDiscountAmount, &q.BeforeTaxAmount, &q.TaxAmount, &q.TotalAmount, &q.NetDebtAmount, &q.TaxRate, &q.ProjectId, &q.ProjectCode, &q.AllocateId, &q.AllocateCode, &q.IsCancel, &q.CreateBy, &q.CreateTime, &q.EditBy, &q.EditTime, &q.CancelBy, &q.CancelTime, &q.ArBillAddress, &q.ArTelephone)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return err
	}

	sql_sub := `select a.Id,a.QuoId,a.ItemId,a.ItemCode,a.ItemName,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel from QuotationSub a  where QuoId = ? order by a.linenumber`
	err = db.Select(&q.Subs, sql_sub, q.Id)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return err
	}

	return nil
}

//func (qs *BCQuotationSub) getSub(db *sqlx.DB) (resp []BCQuotationSub, err error) {
//	strID := string(qs.QuoId)
//	fmt.Println("strID = ", strID)
//	sql := "select ItemCode,ItemName,Qty,Price,ItemAmount,UnitCode from QuotationSub where QuoId =" + strconv.FormatInt(qs.QuoId, 10) + ""
//
//	fmt.Println("get qtsub query --> ", sql)
//	x := BCQuotationSub{}
//
//	rs, err := db.Query(sql)
//
//	if err != nil {
//		return nil, err
//	}
//
//	for rs.Next() {
//		err = rs.Scan(&x.ItemCode, &x.ItemName, &x.Qty, &x.Price, &x.ItemAmount, &x.UnitCode)
//		if err != nil {
//			log.Fatalf("error rs.scan obj %v", err.Error())
//		}
//		fmt.Println("\n fetch sub---> ", x)
//		resp = append(resp, x)
//	}
//
//	return resp, nil
//}
