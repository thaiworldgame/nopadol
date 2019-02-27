package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type BCQuotation struct {
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
	TaxType             int64            `json:"tax_type"  db:"TaxType"`
	TaxRate             float64          `json:"tax_rate"  db:"TaxRate"`
	DepartId            int64            `json:"depart_id" db:"DepartId"`
	RefNo               string           `json:"ref_no" db:"RefNo"`
	JobId               string           `json:"job_id" db:"JobId"`
	IsConfirm           int64            `json:"is_confirm" db:"IsConfirm"`
	BillStatus          int64            `json:"bill_status" db:"BillStatus"`
	Validity            int64            `json:"validity"db:"Validity"`
	CreditDay           int64            `json:"credit_day" db:"CreditDay"`
	DueDate             string           `json:"due_date"db:"DueDate"`
	ExpireCredit        int64            `json:"expire_credit"db:"ExpireCredit"`
	ExpireDate          string           `json:"expire_date"db:"ExpireDate"`
	DeliveryDay         int64            `json:"delivery_day"db:"DeliveryDay"`
	DeliveryDate        string           `json:"delivery_date"db:"DeliveryDate"`
	AssertStatus        int64            `json:"assert_status"db:"AssertStatus"`
	IsConditionSend     int64            `json:"is_condition_send"db:"IsConditionSend"`
	MyDescription       string           `json:"my_description"db:"MyDescription"`
	SumOfItemAmount     float64          `json:"sum_of_item_amount"db:"SumOfItemAmount"`
	DiscountWord        string           `json:"discount_word"db:"DiscountWord"`
	DiscountAmount      float64          `json:"discount_amount"db:"DiscountAmount"`
	AfterDiscountAmount float64          `json:"after_discount_amount"db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64          `json:"before_tax_amount"db:"BeforeTaxAmount"`
	TaxAmount           float64          `json:"tax_amount"db:"TaxAmount"`
	TotalAmount         float64          `json:"total_amount"db:"TotalAmount"`
	NetDebtAmount       float64          `json:"net_debt_amount"db:"NetDebtAmount"`
	ProjectId           int64            `json:"project_id"db:"ProjectId"`
	AllocateId          int64            `json:"allocate_id"db:"AllocateId"`
	IsCancel            int64            `json:"is_cancel"db:"IsCancel"`
	CreateBy            string           `json:"create_by"db:"CreateBy"`
	CreateTime          string           `json:"create_time"db:"CreateTime"`
	EditBy              string           `json:"edit_by"db:"EditBy"`
	EditTime            string           `json:"edit_time"db:"EditTime"`
	ConfirmBy           string           `json:"confirm_by"db:"ConfirmBy"`
	ConfirmTime         string           `json:"confirm_time"db:"ConfirmTime"`
	CancelBy            string           `json:"cancel_by"db:"CancelBy"`
	CancelTime          string           `json:"cancel_time"db:"CancelTime"`
	Subs                []BCQuotationSub `json:"subs" db:"subs"`
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

func (q *BCQuotation) get(db *sqlx.DB) error {
	sql := "select a.uuid,a.Id,DocNo,DocDate,BillType,b.code as arcode,SumOfItemAmount " +
		"from Quotation a left join Customer b on a.ArId = b.id  where DocNo='" + q.DocNo + "'"
	fmt.Println(sql)
	rs := db.QueryRow(sql)

	err := rs.Scan(&q.UUID, &q.Id, &q.DocNo, &q.DocDate, &q.BillType, &q.ArCode, &q.SumOfItemAmount)

	fmt.Println(q)
	if err != nil {
		return err
	}

	qs := BCQuotationSub{
		QuoId:   q.Id,
		QuoUUID: q.UUID,
	}

	fmt.Printf("\n\n quotation.id = %v \n", q.Id)

	subs, err := qs.getSub(db)
	if err != nil {
		return err
	}

	fmt.Println("receive sub data --->", subs)
	q.Subs = subs

	return nil
}

func (qs *BCQuotationSub) getSub(db *sqlx.DB) (resp []BCQuotationSub, err error) {
	strID := string(qs.QuoId)
	fmt.Println("strID = ", strID)
	sql := "select ItemCode,Qty,Price,ItemAmount,UnitCode from QuotationSub where QuoId =" + strconv.FormatInt(qs.QuoId, 10) + ""

	fmt.Println("get qtsub query --> ", sql)
	x := BCQuotationSub{}

	rs, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	for rs.Next() {
		err = rs.Scan(&x.ItemCode, &x.Qty, &x.Price, &x.ItemAmount, &x.UnitCode)
		if err != nil {
			log.Fatalf("error rs.scan obj %v", err.Error())
		}
		fmt.Println("\n fetch sub---> ", x)
		resp = append(resp, x)
	}

	return resp, nil
}
