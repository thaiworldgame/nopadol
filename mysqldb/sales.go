package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/sales"
	"fmt"
	"github.com/mrtomyum/nopadol/config"
	"time"
	"strconv"
	"errors"
	//"net/http"
	//"bytes"
	//"encoding/json"
	//"log"
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
	DepartCode          string            `db:"DepartCode"`
	RefNo               string            `db:"RefNo"`
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
	SaleId              int                `db:"SaleId"`
	SaleCode            string             `db:"SaleCode"`
	SaleName            string             `db:"SaleName"`
	BillType            int64              `db:"BillType"`
	TaxType             int                `db:"TaxType"`
	TaxRate             float64            `db:"TaxRate"`
	DepartCode          string             `db:"DepartCode"`
	RefNo               string             `db:"RefNo"`
	IsConfirm           int64              `db:"IsConfirm"`
	BillStatus          int64              `db:"BillStatus"`
	SoStatus            int64              `db:"SoStatus"`
	HoldingStatus       int64              `db:"HoldingStatus"`
	CreditDay           int64              `db:"CreditDay"`
	DueDate             string             `db:"DueDate"`
	ExpireDate          string             `db:"ExpireDate"`
	DeliveryDate        string             `db:"DeliveryDate"`
	IsConditionSend     int64              `db:"IsConditionSend"`
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

type salesRepository struct{ db *sqlx.DB }

func NewSalesRepository(db *sqlx.DB) sales.Repository {
	return &salesRepository{db}
}

func (repo *salesRepository) CreateQuo(req *sales.NewQuoTemplate) (resp interface{}, err error) {
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

	req.DocDate = DocDate
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

		sql := `INSERT INTO Quotation(DocNo,DocDate,BillType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartCode,RefNo,TaxType,TaxRate,DueDate,ExpireDate,DeliveryDate,AssertStatus,IsConditionSend,MyDescription,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,ProjectId,CreateBy,CreateTime,Validity,CreditDay,ExpireCredit,DeliveryDay,AllocateId,DocType,BranchId,CompanyId) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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
			req.DepartCode,
			req.RefNo,
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

		sql := `Update Quotation set DocDate=?,BillType=?,ArId=?,ArCode=?,ArName=?,SaleId=?,SaleCode=?,SaleName=?,DepartCode=?,RefNo=?,TaxType=?,TaxRate=?,DueDate=?,ExpireDate=?,DeliveryDate=?,AssertStatus=?,IsConditionSend=?,MyDescription=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscountAmount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,NetDebtAmount=?,ProjectId=?,EditBy=?,EditTime=?,AllocateId=?,DocType=?,CompanyId=?,BranchId=?,Validity=?,CreditDay=?,ExpireCredit=?,DeliveryDay=? where Id=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.DocDate, req.BillType, req.ArId, req.ArCode, req.ArName, req.SaleId, req.SaleCode, req.SaleName, req.DepartCode, req.RefNo, req.TaxType, req.TaxRate, req.DueDate, req.ExpireDate, req.DeliveryDate, req.AssertStatus, req.IsConditionSend, req.MyDescription, req.SumOfItemAmount, req.DiscountWord, req.DiscountAmount, req.AfterDiscountAmount, req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount, req.NetDebtAmount, req.ProjectId, req.EditBy, req.EditTime, req.AllocateId, req.DocType, req.CompanyId, req.BranchId, req.Validity, req.CreditDay, req.ExpireCredit, req.DeliveryDay, req.Id)
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

	sql := `select   	a.Id,DocNo,DocDate,DocType,Validity,BillType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,ifnull(DepartCode,'') as DepartCode,ifnull(RefNo,'') as RefNo,TaxType,IsConfirm,BillStatus,CreditDay,ifnull(DueDate,'') as DueDate,ExpireCredit,ifnull(ExpireDate,'') as ExpireDate,DeliveryDay,ifnull(DeliveryDate,'') as DeliveryDate,AssertStatus,IsConditionSend,ifnull(MyDescription,'') as MyDescription,SumOfItemAmount,ifnull(DiscountWord,'') as DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,TaxRate,ProjectId,AllocateId,IsCancel,ifnull(CreateBy,'') as CreateBy,ifnull(CreateTime,'') as CreateTime,ifnull(EditBy,'') as EditBy,ifnull(EditTime,'') as EditTime,ifnull(CancelBy,'') as CancelBy,ifnull(CancelTime,'') as CancelTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from 	Quotation a left join Customer b on a.ArId = b.id  where a.Id = ?`
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
		SaleId:              x.SaleId,
		SaleCode:            x.SaleCode,
		SaleName:            x.SaleName,
		DepartCode:          x.DepartCode,
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

func (repo *salesRepository) CreateSale(req *sales.NewSaleTemplate) (resp interface{}, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var new_doc_no string
	var item_discount_amount_sub float64

	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)
	count_item_qty = 0
	count_item_unit = 0

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	req.DocDate = DocDate
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

	sqlexist := `select count(DocNo) as check_exist from SaleOrder where DocNo = ?`
	err = repo.db.Get(&check_doc_exist, sqlexist, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	switch {
	case check_doc_exist == 0:
		fmt.Println("error =", "Docno is exist")
		return nil, errors.New("Docno is exist")
	}

	fmt.Println("SOStatus =", req.SoStatus, new_doc_no)

	//if (check_doc_exist == 0) {
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
	//if req.BillType == 0 {
	//	jsonStr = []byte(`{"table_code":"SO","bill_type":0}`)
	//} else {
	//	jsonStr = []byte(`{"table_code":"SO","bill_type":1}`)
	//}
	//
	//reqs, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//reqs.Header.Set("X-Custom-Header", "myvalue")
	//reqs.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//resp1, err := client.Do(reqs)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp1.Body.Close()
	//
	//if err := json.NewDecoder(resp1.Body).Decode(&new_doc_no); err != nil {
	//	log.Println(err)
	//}

	//req.DocNo = new_doc_no

	req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

	sql := `INSERT INTO SaleOrder(DocNo,DocDate,BillType,TaxType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartCode,CreditDay,DueDate,DeliveryDate,TaxRate,IsConfirm,MyDescription,BillStatus,SoStatus,HoldingStatus,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,IsCancel,IsConditionSend,JobId,ProjectId,AllocateId,CreateBy,CreateTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	res, err := repo.db.Exec(sql,
		req.DocNo,
		req.DocDate,
		req.BillType,
		req.TaxType,
		req.ArId,
		req.ArCode,
		req.ArName,
		req.SaleId,
		req.SaleCode,
		req.SaleName,
		req.DepartCode,
		req.CreditDay,
		req.DueDate,
		req.DeliveryDate,
		req.TaxRate,
		req.IsConfirm,
		req.MyDescription,
		req.BillStatus,
		req.SoStatus,
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

	var vLineNumber int
	vLineNumber = 0

	for _, sub := range req.Subs {
		fmt.Println("ArId Sub = ", req.ArId)
		fmt.Println("SaleId Sub = ", req.SaleId)
		sub.LineNumber = vLineNumber
		sqlsub := `INSERT INTO SaleOrderSub(SOId,ArId,SaleId,ItemId,ItemCode,BarCode,ItemName,WhCode,ShelfCode,Qty,RemainQty,UnitCode,Price,DiscountWord,DiscountAmount,ItemAmount,ItemDescription,RefNo,QuoId,IsCancel,PackingRate1,RefLineNumber,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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

	//} //else {
	//	switch {
	//	case req.DocNo == "":
	//		fmt.Println("error =", "Docno is null")
	//		return nil, errors.New("Docno is null")
	//	}
	//
	//	sql := `set dateformat dmy     update dbo.bcarinvoice set DocDate=?,ArCode=?,TaxType=?,CashierCode=?,ShiftNo=?,MachineNo=?,MachineCode=?,GrandTotal=?,CoupongAmount=?,ChangeAmount=?,SaleCode=?,TaxRate=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,NetDebtAmount=?,HomeAmount=?,BillBalance=?,LastEditorCode=?,LastEditDateT=getdate() where DocNo=?`
	//	fmt.Println("sql update = ", sql)
	//	id, err := repo.db.Exec(sql, req.DocDate, req.ArCode, pos_tax_type, req.CashierCode, req.ShiftNo, req.MachineNo, req.MachineCode, total_amount, req.CoupongAmount, req.ChangeAmount, req.SaleCode, tax_rate, req.SumOfItemAmount, req.DiscountWord, discount_amount, req.AfterDiscount, before_tax_amount, tax_amount, req.TotalAmount, req.SumCashAmount, req.SumChqAmount, req.SumCreditAmount, req.SumBankAmount, req.NetDebtAmount, home_amount, bill_balance, req.UserCode, req.DocNo)
	//	if err != nil {
	//		fmt.Println("Error = ", err.Error())
	//		return nil, err
	//	}
	//
	//	lastId, err = id.LastInsertId()
	//}
	//
	//sql_del_sub := `delete dbo.bcarinvoicesub where docno = ?`
	//_, err = repo.db.Exec(sql_del_sub, req.DocNo)
	//if err != nil {
	//	fmt.Println("Error = ", err.Error())
	//	return nil, err
	//}
	//
	//for _, item := range req.PosSubs {
	//	fmt.Println("ItemSub")
	//	item_discount_amount, err := strconv.ParseFloat(item.DiscountWord, 64)
	//
	//	item_amount = item.Qty * (item.Price - item_discount_amount)
	//
	//	my_type = def.PosMyType
	//	cn_qty = item.Qty
	//	item.LineNumber = line_number
	//
	//	if (item.PackingRate1 == 0) {
	//		item.PackingRate1 = 1
	//	}
	//	packing_rate_2 = 1
	//
	//	switch {
	//	case pos_tax_type == 0:
	//		item_home_amount = item_amount
	//		item_net_amount = item_amount
	//	case pos_tax_type == 1:
	//		taxamount := toFixed(item_amount-((item_amount*100)/(100+float64(tax_rate))), 2)
	//		beforetaxamount := toFixed(item_amount-taxamount, 2)
	//		item_home_amount = beforetaxamount
	//		item_net_amount = beforetaxamount
	//	case pos_tax_type == 2:
	//		item_home_amount = item_amount
	//		item_net_amount = item_amount
	//	}
	//
	//	sum_of_cost = item.AverageCost * item.Qty
	//
	//	sqlsub := `set dateformat dmy      insert into dbo.BCArInvoiceSub(MyType,DocNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, LineNumber, BarCode, POSSTATUS, AVERAGECOST, PackingRate1, PackingRate2) values(?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	//	_, err = repo.db.Exec(sqlsub, my_type, req.DocNo, pos_tax_type, item.ItemCode, req.DocDate, req.ArCode, depart_code, req.SaleCode, "MobileApp", item.ItemName, item.WHCode, item.ShelfCode, cn_qty, item.Qty, item.Price, item.DiscountWord, item_discount_amount, item_amount, item_net_amount, item_home_amount, sum_of_cost, item.UnitCode, item.LineNumber, item.BarCode, pos_status, item.AverageCost, item.PackingRate1, packing_rate_2)
	//	fmt.Println("sqlsub = ", sqlsub, my_type, req.DocNo, pos_tax_type, item.ItemCode, req.DocDate, req.ArCode, depart_code, req.SaleCode, "MobileApp", item.ItemName, item.WHCode, item.ShelfCode, cn_qty, item.Qty, item.Price, item.DiscountWord, item_discount_amount, item_amount, item_net_amount, item_home_amount, sum_of_cost, item.UnitCode, item.LineNumber, item.BarCode, pos_status, item.AverageCost, item.PackingRate1, packing_rate_2)
	//	if err != nil {
	//		fmt.Println("Error = ", err.Error())
	//		return nil, err
	//	}
	//
	//	sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
	//	_, err = repo.db.Exec(sqlprocess, item.ItemCode)
	//	fmt.Println("sqlprocess = ", sqlsub)
	//	if err != nil {
	//		fmt.Println("Error = ", err.Error())
	//		fmt.Println(err.Error())
	//	}
	//
	//	line_number = line_number + 1
	//}

	return map[string]interface{}{
		"company_name":    req.DocNo,
		"company_address": req.DocDate,
	}, nil
}

func (repo *salesRepository) SearchSaleById(req *sales.SearchByIdTemplate) (resp interface{}, err error) {

	q := NewSaleModel{}

	sql := `select isnull(CompanyName,'') as CompanyName,isnull(CompanyAddress,'') as CompanyAddress,isnull(Telephone,'') as Telephone,isnull(TaxId,'') as TaxId,isnull(ArCode,'') as ArCode,isnull(PosId,'') as PosId,isnull(WhCode,'') as WhCode,isnull(ShelfCode,'') as ShelfCode,isnull(PrinterPosIp,'') as PrinterPosIp,isnull(PrinterCopyIp,'') as PrinterCopyIp,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode from posconfig`
	err = repo.db.Get(&q, sql)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	qt_resp := map_sale_template(q)

	return qt_resp, nil
}

func map_sale_template(x NewSaleModel) sales.NewSaleTemplate {
	return sales.NewSaleTemplate{
		ArCode: x.ArCode,
	}
}
