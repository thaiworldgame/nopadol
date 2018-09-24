package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/sales"
	"fmt"
	"github.com/mrtomyum/nopadol/config"
	"time"
	"strconv"
	"errors"
	"net/http"
	"bytes"
	"encoding/json"
	"log"
)

type NewQuoModel struct {
	Id                  int64             `db:"Id"`
	DocNo               string            `db:"DocNo"`
	DocDate             string            `db:"DocDate"`
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
	Subs                []NewQuoItemModel `db:"subs"`
}

type NewQuoItemModel struct {
	Id              int64   `db:"Id"`
	QuoId           int64   `db:"QuoId"`
	ArId            int64   `db:"ArId"`
	SaleId          int64   `db:"SaleId"`
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
	LineNumber      int     `db:"LineNumber"`
}

type NewSaleModel struct {
	Id                  int64              `db:"Id"`
	DocNo               string             `db:"DocNo"`
	DocDate             string             `db:"DocDate"`
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
	DueDate             string             `db:"DueDate"`
	ExpireDate          string             `db:"ExpireDate"`
	DeliveryDate        string             `db:"DeliveryDate"`
	AssertStatus        int64              `db:"AssertStatus"`
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
	IsCancel            int64              `db:"IsCancel"`
	CreateBy            string             `db:"CreateBy"`
	CreateTime          string             `db:"CreateTime"`
	EditBy              string             `db:"EditBy"`
	EditTime            string             `db:"EditTime"`
	CancelBy            string             `db:"CancelBy"`
	CancelTime          string             `db:"CancelTime"`
	Subs                []NewSaleItemModel `db:"subs"`
}

type NewSaleItemModel struct {
	Id              int64   `db:"Id"`
	QuoId           int64   `db:"QuoId"`
	ArId            int64   `db:"ArId"`
	SaleId          int64   `db:"SaleId"`
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
	LineNumber      int     `db:"LineNumber"`
}

type salesRepository struct{ db *sqlx.DB }

func NewSalesRepository(db *sqlx.DB) sales.Repository {
	return &salesRepository{db}
}

func (repo *salesRepository) CreateQuo(req *sales.NewQuoTemplate) (resp interface{}, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	//var tax_rate float64
	//var pos_tax_type int
	//
	//var is_complete_save int
	//var deposit_inc_tax int
	//var home_amount float64
	//var bill_balance float64
	//var pos_status int
	//
	//var line_number int
	//var item_amount float64
	//var my_type int
	//var cn_qty float64
	//var packing_rate_2 float64
	//var item_home_amount float64
	//var item_net_amount float64
	var new_doc_no string

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

	fmt.Println("DocDate =", req.DocDate)

	pos_machine_no := def.PosMachineNo
	fmt.Println("pos_machine_no", pos_machine_no)

	//tax_rate = def.TaxRateDefault
	//pos_tax_type = def.PosTaxType

	for _, sub_item := range req.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1

			item_discount_amount_sub, err := strconv.ParseFloat(sub_item.DiscountWord, 64)
			if err != nil {
				fmt.Println(err)
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

	sqlexist := `select count(DocNo) as check_exist from Quotation where DocNo = ?`
	fmt.Println("DocNo =", req.DocNo)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

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
		url := "http://localhost:8081/gendocno/v1/gen"
		var jsonStr []byte

		if req.BillType == 0 {
			jsonStr = []byte(`{"table_code":"QT","bill_type":0}`)
		} else {
			jsonStr = []byte(`{"table_code":"QT","bill_type":1}`)
		}

		reqs, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		reqs.Header.Set("X-Custom-Header", "myvalue")
		reqs.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(reqs)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&new_doc_no); err != nil {
			log.Println(err)
		}

		req.DocNo = new_doc_no

		fmt.Println("Docno =", req.DocNo, new_doc_no)
		switch {
		case req.DocNo == "":
			fmt.Println("error =", "Docno is null")
			return nil, errors.New("Docno is null")
		}
		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

		sql := `INSERT INTO Quotation(DocNo,DocDate,BillType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartCode,RefNo,TaxType,TaxRate,DueDate,ExpireDate,DeliveryDate,AssertStatus,IsConditionSend,MyDescription,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,ProjectId,CreateBy,CreateTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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
			req.CreateTime)

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
			sqlsub := `INSERT INTO QuotationSub(QuoId,ArId,SaleId,ItemId,ItemCode,ItemName,Qty,RemainQty,Price,DiscountWord,DiscountAmount,UnitCode,ItemAmount,ItemDescription,PackingRate1,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err := repo.db.Exec(sqlsub,
				req.Id,
				req.ArId,
				req.SaleId,
				sub.ItemId,
				sub.ItemCode,
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
				sub.LineNumber)

			fmt.Println("QuotationSub =", sql, sub.QuoId)
			if err != nil {
				return "Insert Quotation Not Success", err
			}
		}

	} //else {
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

func (repo *salesRepository) SearchQuoById() (resp interface{}, err error) {

	q := NewQuoModel{}

	sql := `select isnull(CompanyName,'') as CompanyName,isnull(CompanyAddress,'') as CompanyAddress,isnull(Telephone,'') as Telephone,isnull(TaxId,'') as TaxId,isnull(ArCode,'') as ArCode,isnull(PosId,'') as PosId,isnull(WhCode,'') as WhCode,isnull(ShelfCode,'') as ShelfCode,isnull(PrinterPosIp,'') as PrinterPosIp,isnull(PrinterCopyIp,'') as PrinterCopyIp,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode from posconfig`
	err = repo.db.Get(&q, sql)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	qt_resp := map_quo_template(q)

	return qt_resp, nil
}

func map_quo_template(x NewQuoModel) sales.NewQuoTemplate {
	return sales.NewQuoTemplate{
		ArCode: x.ArCode,
	}
}

func (repo *salesRepository) CreateSale(req *sales.NewSaleTemplate) (resp interface{}, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	//var tax_rate float64
	//var pos_tax_type int
	//
	//var is_complete_save int
	//var deposit_inc_tax int
	//var home_amount float64
	//var bill_balance float64
	//var pos_status int
	//
	//var line_number int
	//var item_amount float64
	//var my_type int
	//var cn_qty float64
	//var packing_rate_2 float64
	//var item_home_amount float64
	//var item_net_amount float64
	var new_doc_no string

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

	fmt.Println("DocDate =", req.DocDate)

	pos_machine_no := def.PosMachineNo
	fmt.Println("pos_machine_no", pos_machine_no)

	//tax_rate = def.TaxRateDefault
	//pos_tax_type = def.PosTaxType

	for _, sub_item := range req.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1

			item_discount_amount_sub, err := strconv.ParseFloat(sub_item.DiscountWord, 64)
			if err != nil {
				fmt.Println(err)
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

	sqlexist := `select count(DocNo) as check_exist from Quotation where DocNo = ?`
	fmt.Println("DocNo =", req.DocNo)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

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
		url := "http://localhost:8081/gendocno/v1/gen"
		var jsonStr []byte

		if req.BillType == 0 {
			jsonStr = []byte(`{"table_code":"QT","bill_type":0}`)
		} else {
			jsonStr = []byte(`{"table_code":"QT","bill_type":1}`)
		}

		reqs, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		reqs.Header.Set("X-Custom-Header", "myvalue")
		reqs.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(reqs)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&new_doc_no); err != nil {
			log.Println(err)
		}

		req.DocNo = new_doc_no

		fmt.Println("Docno =", req.DocNo, new_doc_no)
		switch {
		case req.DocNo == "":
			fmt.Println("error =", "Docno is null")
			return nil, errors.New("Docno is null")
		}
		req.BeforeTaxAmount, req.TaxAmount, req.TotalAmount = config.CalcTaxItem(req.TaxType, req.TaxRate, req.AfterDiscountAmount)

		sql := `INSERT INTO Quotation(DocNo,DocDate,BillType,ArId,ArCode,ArName,SaleId,SaleCode,SaleName,DepartCode,RefNo,TaxType,TaxRate,DueDate,ExpireDate,DeliveryDate,AssertStatus,IsConditionSend,MyDescription,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscountAmount,BeforeTaxAmount,TaxAmount,TotalAmount,NetDebtAmount,ProjectId,CreateBy,CreateTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
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
			req.CreateTime)

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
			sqlsub := `INSERT INTO QuotationSub(QuoId,ArId,SaleId,ItemId,ItemCode,ItemName,Qty,RemainQty,Price,DiscountWord,DiscountAmount,UnitCode,ItemAmount,ItemDescription,PackingRate1,LineNumber) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err := repo.db.Exec(sqlsub,
				req.Id,
				req.ArId,
				req.SaleId,
				sub.ItemId,
				sub.ItemCode,
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
				sub.LineNumber)

			fmt.Println("QuotationSub =", sql, sub.QuoId)
			if err != nil {
				return "Insert Quotation Not Success", err
			}
		}

	} //else {
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

func (repo *salesRepository) SearchSaleById() (resp interface{}, err error) {

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
