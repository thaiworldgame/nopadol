package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/sales"
	"fmt"
	"github.com/mrtomyum/nopadol/config"
	"time"
	"strconv"
	"errors"
)

type NewQuoModel struct {
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
	SaleId              int               `db:"SaleId"`
	SaleCode            string            `db:"SaleCode"`
	SaleName            string            `db:"SaleName"`
	BillType            int64             `db:"BillType"`
	TaxType             int               `db:"TaxType"`
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
	CancelBy            string            `db:"CancelBy"`
	CancelTime          string            `db:"CancelTime"`
	Subs                []NewQuoItemModel `db:"subs"`
}

type NewQuoItemModel struct {
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
	ArBillAddress       string             `db:"ar_bill_address"`
	ArTelephone         string             `db:"ar_telephone"`
	SaleId              int                `db:"SaleId"`
	SaleCode            string             `db:"SaleCode"`
	SaleName            string             `db:"SaleName"`
	BillType            int64              `db:"BillType"`
	TaxType             int                `db:"TaxType"`
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
	SaleId              int               `db:"SaleId"`
	SaleCode            string            `db:"SaleCode"`
	SaleName            string            `db:"SaleName"`
	BillType            int64             `db:"BillType"`
	TaxType             int               `db:"TaxType"`
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
		if (sub_item.Qty != 0) {
			count_item = count_item + 1

			fmt.Println("Count Item =", count_item)

			sum_item_amount = sum_item_amount + (sub_item.Qty * (sub_item.Price - sub_item.DiscountAmount))
		}
		if (sub_item.ItemCode != "" && sub_item.Qty == 0) {
			count_item_qty = count_item_qty + 1
		}
		if (sub_item.ItemCode != "" && sub_item.UnitCode == "") {
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

	if (check_doc_exist == 0) {
		//API Call Get API
		//url := "http://localhost:8081/gendocno/v1/gen?table_code=QT&bill_type=0"
		//reqs, err := http.NewRequest("POST", url, nil)
		//if err != nil {
		//	log.Fatal("NewRequest: ", err)
		//	return nil, err
		//}

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

		//API Get Post API
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

		sql := `INSERT INTO Quotation(DocNo,DocDate,BillType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartId,RefNo,JobId,TaxType,TaxRate,DueDate,ExpireDate,DeliveryDate,AssertStatus,IsConditionSend,MyDescription,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,ProjectId,CreateBy,CreateTime,Validity,CreditDay,ExpireCredit,DeliveryDay,AllocateId,DocType,BranchId,CompanyId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		res, err := repo.db.Exec(sql,
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
			req.CompanyId)

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
			sqlsub := `INSERT INTO QuotationSub(QuoId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,Qty,RemainQty,Price,DiscountWord,DiscountAmount,UnitCode,ItemAmount,ItemDescription,PackingRate1,LineNumber,IsCancel) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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
				sub.IsCancel)

			fmt.Println("QuotationSub =", sql, sub.QuoId)
			if err != nil {
				return "Insert Quotation Not Success", err
			}
		}

	} else {
		fmt.Println("Update")
		req.EditBy = req.CreateBy

		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)
		req.NetDebtAmount = req.TotalAmount

		sql := `Update Quotation set DocDate=?,BillType=?,ArId=?,ArCode=?,ArName=?,SaleId=?,SaleCode=?,SaleName=?,DepartId=?,RefNo=?,JobId=?,TaxType=?,TaxRate=?,DueDate=?,ExpireDate=?,DeliveryDate=?,AssertStatus=?,IsConditionSend=?,MyDescription=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscountAmount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,NetDebtAmount=?,ProjectId=?,EditBy=?,EditTime=?,AllocateId=?,DocType=?,CompanyId=?,BranchId=?,Validity=?,CreditDay=?,ExpireCredit=?,DeliveryDay=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.DocDate, req.BillType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.DepartId, req.RefNo, req.JobId, req.TaxType, req.TaxRate, req.DueDate, req.ExpireDate, req.DeliveryDate, req.AssertStatus, req.IsConditionSend, req.MyDescription, req.SumOfItemAmount, req.DiscountWord, req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetDebtAmount, req.ProjectId, req.EditBy, req.EditTime, req.AllocateId, req.DocType, req.CompanyId, req.BranchId, req.Validity, req.CreditDay, req.ExpireCredit, req.DeliveryDay, req.Id)
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
		sqlsub := `INSERT INTO QuotationSub(QuoId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,Qty,RemainQty,Price,DiscountWord,DiscountAmount,UnitCode,ItemAmount,ItemDescription,PackingRate1,LineNumber,IsCancel) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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
			sub.IsCancel)
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

func (repo *salesRepository) SearchQuoById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	q := NewQuoModel{}

	sql := `select a.Id,a.DocNo,a.DocDate,a.DocType,a.Validity,a.BillType,a.ArId,a.ArCode,a.ArName,a.SaleId,a.SaleCode,a.SaleName,ifnull(a.DepartId,0) as DepartId,ifnull(a.RefNo,'') as RefNo,ifnull(a.JobId,'') as JobId,a.TaxType,a.IsConfirm,a.BillStatus,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.ExpireCredit,ifnull(a.ExpireDate,'') as ExpireDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.AssertStatus,a.IsConditionSend,ifnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.TaxRate,a.ProjectId,a.AllocateId,a.IsCancel,ifnull(a.CreateBy,'') as CreateBy,ifnull(a.CreateTime,'') as CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime,ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from Quotation a left join Customer b on a.ArId = b.id  where a.Id = ?`
	err = repo.db.Get(&q, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	qt_resp := map_quo_template(q)

	subs := []NewQuoItemModel{}

	sql_sub := `select a.Id,a.QuoId,a.ItemId,a.ItemCode,a.ItemName,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel from QuotationSub a  where QuoId = ? order by a.linenumber`
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
	}
}

func (repo *salesRepository) CreateSaleOrder(req *sales.NewSaleTemplate) (resp interface{}, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var item_discount_amount_sub float64

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
		if (sub_item.Qty != 0) {
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
		if (sub_item.ItemCode != "" && sub_item.Qty == 0) {
			count_item_qty = count_item_qty + 1
		}
		if (sub_item.ItemCode != "" && sub_item.UnitCode == "") {
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

	//switch {
	//case check_doc_exist != 0:
	//	fmt.Println("error =", "Docno is exist")
	//	return nil, errors.New("Docno is exist")
	//}

	if (check_doc_exist == 0) {

		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

		sql := `INSERT INTO SaleOrder(DocNo,DocDate,CompanyId,BranchId,DocType,BillType,TaxType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartId,CreditDay,DueDate,DeliveryDay,DeliveryDate,TaxRate,IsConfirm,MyDescription,BillStatus,HoldingStatus,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,IsCancel,IsConditionSend,DeliveryAddressId,CarLicense,PersonReceiveTel,JobId,ProjectId,AllocateId,CreateBy,CreateTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		res, err := repo.db.Exec(sql,
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
			req.ProjectId,
			req.AllocateId,
			req.CreateBy,
			req.CreateTime)

		//fmt.Println("query=", sql, "Hello")
		if err != nil {
			return "", err
		}

		id, _ := res.LastInsertId()
		req.Id = id

		//var vLineNumber int
		//vLineNumber = 0
		//
		//for _, sub := range req.Subs {
		//	fmt.Println("ArId Sub = ", req.ArId)
		//	fmt.Println("SaleId Sub = ", req.SaleId)
		//	sub.LineNumber = vLineNumber
		//	sub.SumOfCost = sub.AverageCost * sub.Qty
		//
		//	sqlsub := `INSERT INTO SaleOrderSub(SOId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,WhCode,ShelfCode,Qty,RemainQty,UnitCode,Price,DiscountWord,DiscountAmount,ItemAmount,ItemDescription,StockType,AverageCost,SumOfCost,RefNo,QuoId,IsCancel,PackingRate1,RefLineNumber,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		//	_, err := repo.db.Exec(sqlsub,
		//		req.Id,
		//		req.ArId,
		//		req.SaleId,
		//		sub.ItemId,
		//		sub.ItemCode,
		//		sub.BarCode,
		//		sub.ItemName,
		//		sub.WHCode,
		//		sub.ShelfCode,
		//		sub.Qty,
		//		sub.RemainQty,
		//		sub.UnitCode,
		//		sub.Price,
		//		sub.DiscountWord,
		//		sub.DiscountAmount,
		//		sub.ItemAmount,
		//		sub.ItemDescription,
		//		sub.StockType,
		//		sub.AverageCost,
		//		sub.SumOfCost,
		//		sub.RefNo,
		//		sub.QuoId,
		//		sub.IsCancel,
		//		sub.PackingRate1,
		//		sub.RefLineNumber,
		//		sub.LineNumber)
		//
		//	vLineNumber = vLineNumber + 1
		//	if err != nil {
		//		return "Insert SaleOrder Not Success", err
		//	}
		//}

	} else {
		switch {
		case req.DocNo == "":
			fmt.Println("error =", "Docno is null")
			return nil, errors.New("Docno is null")
		}

		fmt.Println("Update")
		req.EditBy = req.CreateBy

		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

		sql := `Update SaleOrder set DocNo=?,DocDate=?,CompanyId=?,BranchId=?,DocType=?,BillType=?,TaxType=?,ArId=?,ArCode=?,ArName=?,SaleId=?,SaleCode=?,SaleName=?,DepartId=?,CreditDay=?,DueDate=?,DeliveryDay=?,DeliveryDate=?,TaxRate=?,IsConfirm=?,MyDescription=?,BillStatus=?,HoldingStatus=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscountAmount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,NetDebtAmount=?,IsCancel=?,IsConditionSend=?,DeliveryAddressId=?,CarLicense=?,PersonReceiveTel=?,JobId=?,ProjectId=?,AllocateId=?,EditBy=?,EditTime=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.DocNo, req.DocDate, req.CompanyId, req.BranchId, req.DocType, req.BillType, req.TaxType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.DepartId, req.CreditDay, req.DueDate, req.DeliveryDay, req.DeliveryDate, req.TaxRate, req.IsConfirm, req.MyDescription, req.BillStatus, req.HoldingStatus, req.SumOfItemAmount, req.DiscountWord, req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetDebtAmount, req.IsCancel, req.IsConditionSend, req.DeliveryAddressId, req.CarLicense, req.PersonReceiveTel, req.JobId, req.ProjectId, req.AllocateId, req.EditBy, req.EditTime, req.Id)
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
		sqlsub := `INSERT INTO SaleOrderSub(SOId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,WhCode,ShelfCode,Qty,RemainQty,UnitCode,Price,DiscountWord,DiscountAmount,ItemAmount,ItemDescription,StockType,AverageCost,SumOfCost,RefNo,QuoId,IsCancel,PackingRate1,RefLineNumber,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := repo.db.Exec(sqlsub,
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

	sql := `select a.Id,a.DocNo,ifnull(a.DocDate,'') as DocDate,a.CompanyId,a.BranchId,a.DocType,a.BillType,a.TaxType,a.ArId,ifnull(a.ArCode,'') as ArCode,ifnull(a.ArName,'') as ArName,a.SaleId,ifnull(a.SaleCode,'') as SaleCode,ifnull(a.SaleName,'') as SaleName,a.DepartId,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.TaxRate,a.IsConfirm,ifnull(a.MyDescription,'') as MyDescription,a.BillStatus,a.HoldingStatus,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.IsCancel,a.IsConditionSend,a.DeliveryAddressId,ifnull(a.CarLicense,'') as CarLicense,ifnull(a.PersonReceiveTel,'') as PersonReceiveTel,ifnull(a.JobId,'') as JobId,a.ProjectId,a.AllocateId,ifnull(a.CreateBy,'') as CreateBy,a.CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime, ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime, ifnull(a.ConfirmBy,'') as ConfirmBy,ifnull(a.ConfirmTime,'') as ConfirmTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from SaleOrder a left join Customer on a.ArId = b.id where a.id=?`
	err = repo.db.Get(&s, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	so_resp := map_saleorder_template(s)

	subs := []NewSaleItemModel{}

	fmt.Println("s.Id =", s.Id)

	sql_sub := `select a.Id,a.SOId,a.ItemId,a.ItemCode,a.ItemName,ifnull(a.WHCode,'') as WHCode,ifnull(a.ShelfCode,'') as ShelfCode,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.StockType,a.AverageCost,a.SumOfCost,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel from SaleOrderSub a  where SOId = ? order by a.linenumber`
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

	if (check_doc_exist == 0) {
		sql := `insert into ar_deposit(company_id, branch_id, uuid, doc_no, tax_no, doc_date, bill_type, ar_id, ar_code, ar_name, sale_id, sale_code, sale_name, tax_type, tax_rate, ref_no, credit_day, due_date, depart_id, allocate_id, project_id, my_description, before_tax_amount, tax_amount, total_amount, net_amount ,bill_balance ,cash_amount ,creditcard_amount, chq_amount, bank_amount, is_return_money, is_cancel, is_confirm, scg_id, job_no, create_by, create_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,? ,? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		resp, err := repo.db.Exec(sql, req.CompanyId, req.BranchId, req.Uuid, req.DocNo, req.TaxNo, req.DocDate, req.BillType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.TaxType, req.TaxRate, req.RefNo, req.CreditDay, req.DueDate, req.DepartId, req.AllocateId, req.ProjectId, req.MyDescription, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetAmount, req.BillBalance, req.CashAmount, req.CreditcardAmount, req.ChqAmount, req.BankAmount, req.IsReturnMoney, req.IsCancel, req.IsConfirm, req.ScgId, req.JobNo, req.CreateBy, req.CreateTime)
		if err != nil {
			fmt.Println("error = ", err.Error())
		}
		fmt.Println("sql = ", sql)

		id, _ := resp.LastInsertId()

		req.Id = id
	} else {
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

	count_chq_err = 0
	if req.ChqAmount != 0 {
		for _, chq := range req.Chq {

			sql_del := `delete from chq_in where uuid=? and ref_id=? and company_id=? and branch_id=? and chq_number=?`
			chq_del, _ := repo.db.Exec(sql_del, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber, chq.BankId)
			fmt.Println(sql_del)
			if err != nil {
				fmt.Println("sql_del = ", err.Error())
			}
			fmt.Println(chq_del.RowsAffected())

			verify_chq, _ := verify_chq_in(repo.db, req.Uuid, req.Id, req.CompanyId, req.BranchId, chq.ChqNumber, chq.BankId)
			fmt.Println("verify_chq = ", verify_chq)
			if verify_chq == false {
				count_chq_err = count_chq_err + 1
			}

			sum_chq_amount = sum_chq_amount + chq.ChqAmount
		}

		switch {
		case count_chq_err != 0:
			return nil, errors.New("ข้อมูลเลขที่เช็ค มีการใช้ไปแล้ว")
		case sum_chq_amount != req.ChqAmount:
			return nil, errors.New("มูลค่าเช็ค ไม่เท่ากับ มูลค่ารายการเช็ค")
		}

		for _, i_chq := range req.Chq {
			i_chq.Description = "รับเงินมัดจำ"
			sql_chq := `insert into chq_in (company_id,branch_id,uuid,ref_id,ar_id,doc_no,doc_date,chq_number,bank_id,back_branch_id,receive_date,due_date,book_id,chq_status,chq_amount,chq_balance,description,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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

	sql_verify_crd := `select ifnull(count(doc_no),0) as vcount from credit_card where uuid = ? and ref_id = ? and company_id = ? and branch_id = ? and credit_card_no = ? and confirm_no = ? and bank_id = ?`
	err := db.Get(&exist, sql_verify_crd, Uuid, RefId, CompanyId, BranchId, CreditCardNo, ConfirmNo, BankId)
	fmt.Println("sql_verify_crd = ", sql_verify_crd, Uuid, RefId, CompanyId, BranchId, CreditCardNo, ConfirmNo, BankId)
	if err != nil {
		fmt.Println("sql_verify_crd err = ", err.Error())
		return false, err
	}

	fmt.Println("exist = ", exist)

	if (exist != 0) {
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

	if (exist != 0) {
		return false, nil
	} else {
		return true, nil
	}

}

func (repo *salesRepository) SearchDepositById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	d := NewDepositModel{}

	sql := `select a.id, a.company_id, a.branch_id, ifnull(a.uuid,'') as uuid, ifnull(a.doc_no,'') as doc_no, ifnull(a.tax_no,'') as tax_no, ifnull(a.doc_date,'') as doc_date, a.bill_type, a.ar_id, ifnull(a.ar_code,'') as ar_code, ifnull(a.ar_name,'') as ar_name,a.sale_id, ifnull(a.sale_code,'') as sale_code, ifnull(a.sale_name,'') as sale_name,a.tax_type, a.tax_rate, ifnull(a.ref_no,'') as ref_no, a.credit_day, ifnull(a.due_date,'') as due_date, a.depart_id, a.allocate_id, a.project_id, ifnull(a.my_description,'') as my_description, a.before_tax_amount, a.tax_amount, a.total_amount, a.net_amount ,a.bill_balance ,a.cash_amount ,a.creditcard_amount, a.chq_amount, a.bank_amount, a.is_return_money, a.is_cancel, a.is_confirm, ifnull(a.scg_id,'') as scg_id, ifnull(a.job_no,'') as job_no, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, ifnull(a.edit_by,'') as edit_by, ifnull(a.edit_time,'') as edit_time, ifnull(a.cancel_by,'') as cancel_by, ifnull(a.cancel_time,'') as cancel_time, ifnull(a.confirm_by,'') as confirm_by, ifnull(a.confirm_time,'') as confirm_time,ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone from ar_deposit a left join Customer b on a.ar_id = b.id where a.id=?`
	err = repo.db.Get(&d, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	dp_resp := map_deposit_template(d)

	fmt.Println("CompanyId,BranchId,Uuid", d.CompanyId, d.BranchId, d.Uuid)

	//subs := []NewDepositItemModel{}
	//
	//sql_sub := `select a.Id,a.SOId,a.ItemId,a.ItemCode,a.ItemName,ifnull(a.WHCode,'') as WHCode,ifnull(a.ShelfCode,'') as ShelfCode,a.Qty,a.RemainQty,a.Price,ifnull(a.DiscountWord,'') as DiscountWord,DiscountAmount,ifnull(a.UnitCode,'') as UnitCode,ifnull(a.BarCode,'') as BarCode,ifnull(a.ItemDescription,'') as ItemDescription,a.StockType,a.AverageCost,a.SumOfCost,a.ItemAmount,a.PackingRate1,a.LineNumber,a.IsCancel from SaleOrderSub a  where SOId = ? order by a.linenumber`
	//err = repo.db.Select(&subs, sql_sub, d.RefNo)
	//fmt.Println("sql_sub = ", sql_sub)
	//if err != nil {
	//	fmt.Println("err sub= ", err.Error())
	//	return resp, err
	//}
	//
	//for _, sub := range subs {
	//	subline := map_deposit_subs_template(sub)
	//	dp_resp.Subs = append(dp_resp.Subs, subline)
	//}

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
		sql = `select a.id, a.company_id, a.branch_id, ifnull(a.uuid,'') as uuid, ifnull(a.doc_no,'') as doc_no, ifnull(a.tax_no,'') as tax_no, ifnull(a.doc_date,'') as doc_date, a.bill_type, a.ar_id, ifnull(a.ar_code,'') as ar_code, ifnull(a.ar_name,'') as ar_name,a.sale_id, ifnull(a.sale_code,'') as sale_code, ifnull(a.sale_name,'') as sale_name,a.tax_type, a.tax_rate, ifnull(a.ref_no,'') as ref_no, a.credit_day, ifnull(a.due_date,'') as due_date, a.depart_id, a.allocate_id, a.project_id, ifnull(a.my_description,'') as my_description, a.before_tax_amount, a.tax_amount, a.total_amount, a.net_amount ,a.bill_balance ,a.cash_amount ,a.creditcard_amount, a.chq_amount, a.bank_amount, a.is_return_money, a.is_cancel, a.is_confirm, ifnull(a.scg_id,'') as scg_id, ifnull(a.job_no,'') as job_no, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, ifnull(a.edit_by,'') as edit_by, ifnull(a.edit_time,'') as edit_time, ifnull(a.cancel_by,'') as cancel_by, ifnull(a.cancel_time,'') as cancel_time, ifnull(a.confirm_by,'') as confirm_by, ifnull(a.confirm_time,'') as confirm_time,ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone from ar_deposit a left join Customer b on a.ar_id = b.id  order by a.id desc limit 30`
		err = repo.db.Select(&d, sql)
	} else {
		sql = `select a.id, a.company_id, a.branch_id, ifnull(a.uuid,'') as uuid, ifnull(a.doc_no,'') as doc_no, ifnull(a.tax_no,'') as tax_no, ifnull(a.doc_date,'') as doc_date, a.bill_type, a.ar_id, ifnull(a.ar_code,'') as ar_code, ifnull(a.ar_name,'') as ar_name,a.sale_id, ifnull(a.sale_code,'') as sale_code, ifnull(a.sale_name,'') as sale_name,a.tax_type, a.tax_rate, ifnull(a.ref_no,'') as ref_no, a.credit_day, ifnull(a.due_date,'') as due_date, a.depart_id, a.allocate_id, a.project_id, ifnull(a.my_description,'') as my_description, a.before_tax_amount, a.tax_amount, a.total_amount, a.net_amount ,a.bill_balance ,a.cash_amount ,a.creditcard_amount, a.chq_amount, a.bank_amount, a.is_return_money, a.is_cancel, a.is_confirm, ifnull(a.scg_id,'') as scg_id, ifnull(a.job_no,'') as job_no, ifnull(a.create_by,'') as create_by, ifnull(a.create_time,'') as create_time, ifnull(a.edit_by,'') as edit_by, ifnull(a.edit_time,'') as edit_time, ifnull(a.cancel_by,'') as cancel_by, ifnull(a.cancel_time,'') as cancel_time, ifnull(a.confirm_by,'') as confirm_by, ifnull(a.confirm_time,'') as confirm_time,ifnull(b.address,'') as ar_bill_address,ifnull(b.telephone,'') as ar_telephone from ar_deposit a left join Customer b on a.ar_id = b.id  where a.doc_no like  concat(?,'%') or a.ar_code like  concat(?,'%') or a.ar_name like  concat(?,'%') order by a.id desc limit 30`
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
