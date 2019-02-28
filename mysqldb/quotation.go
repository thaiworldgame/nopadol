package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
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
	BillType            int64            `db:"BillType"`
	TaxType             int64            `db:"TaxType"`
	TaxRate             float64          `db:"TaxRate"`
	DepartId            int64            `db:"DepartId"`
	RefNo               string           `db:"RefNo"`
	JobId               string           `db:"JobId"`
	IsConfirm           int64            `db:"IsConfirm"`
	BillStatus          int64            `db:"BillStatus"`
	Validity            int64            `db:"Validity"`
	CreditDay           int64            `db:"CreditDay"`
	DueDate             string           `db:"DueDate"`
	ExpireCredit        int64            `db:"ExpireCredit"`
	ExpireDate          string           `db:"ExpireDate"`
	DeliveryDay         int64            `db:"DeliveryDay"`
	DeliveryDate        string           `db:"DeliveryDate"`
	AssertStatus        int64            `db:"AssertStatus"`
	IsConditionSend     int64            `db:"IsConditionSend"`
	MyDescription       string           `db:"MyDescription"`
	SumOfItemAmount     float64          `db:"SumOfItemAmount"`
	DiscountWord        string           `db:"DiscountWord"`
	DiscountAmount      float64          `db:"DiscountAmount"`
	AfterDiscountAmount float64          `db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64          `db:"BeforeTaxAmount"`
	TaxAmount           float64          `db:"TaxAmount"`
	TotalAmount         float64          `db:"TotalAmount"`
	NetDebtAmount       float64          `db:"NetDebtAmount"`
	ProjectId           int64            `db:"ProjectId"`
	AllocateId          int64            `db:"AllocateId"`
	IsCancel            int64            `db:"IsCancel"`
	CreateBy            string           `db:"CreateBy"`
	CreateTime          string           `db:"CreateTime"`
	EditBy              string           `db:"EditBy"`
	EditTime            string           `db:"EditTime"`
	ConfirmBy           string           `db:"ConfirmBy"`
	ConfirmTime         string           `db:"ConfirmTime"`
	CancelBy            string           `db:"CancelBy"`
	CancelTime          string           `db:"CancelTime"`
	Subs                []BCQuotationSub `db:"subs"`
}

type BCQuotationSub struct {
	Id              int64   `db:"Id"`
	QuoUUID         string  `db:"QuoUUID"`
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
}

func (q *BCQuotation) get(db *sqlx.DB) error {
	sql := "select a.uuid,a.Id,DocNo,DocDate,BillType,b.code as arcode,SumOfItemAmount " +
		"from Quotation a left join Customer b on a.ArId = b.id  where DocNo='" + q.DocNo + "'"
	fmt.Println(sql)
	rs := db.QueryRow(sql)

	err := rs.Scan(&q.UUID, &q.Id, &q.DocNo, &q.DocDate, &q.BillType, &q.ArCode, &q.SumOfItemAmount)

	fmt.Println("Quotation Head = ",q)
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
