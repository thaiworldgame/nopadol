package sqldb

import (
	"context"
	"github.com/mrtomyum/nopadol/pos"
	"github.com/jmoiron/sqlx"
	"fmt"
	"math"
	"strconv"
	"errors"
	"github.com/mrtomyum/nopadol/config"
)

type NewPosModel struct {
	DocNo           string                `db:"DocNo"`
	DocDate         string                `db:"DocDate"`
	ArCode          string                `db:"ArCode"`
	SaleCode        string                `db:"SaleCode"`
	ShiftCode       string                `db:"ShiftCode"`
	CashierCode     string                `db:"CashierCode"`
	ShiftNo         int                   `db:"ShiftNo"`
	MachineNo       string                `db:"MachineNo"`
	MachineCode     string                `db:"MachineCode"`
	CoupongAmount   float64               `db:"CoupongAmount"`
	ChangeAmount    float64               `db:"ChangeAmount"`
	ChargeAmount    float64               `db:"ChargeAmount"`
	TaxType         int                   `db:"TaxType"`
	SumOfItemAmount float64               `db:"SumOfItemAmount"`
	DiscountWord    string                `db:"DiscountWord"`
	AfterDiscount   float64               `db:"AfterDiscount"`
	TotalAmount     float64               `db:"TotalAmount"`
	SumCashAmount   float64               `db:"SumCashAmount"`
	SumChqAmount    float64               `db:"SumChqAmount"`
	SumCreditAmount float64               `db:"SumCreditAmount"`
	SumBankAmount   float64               `db:"SumBankAmount"`
	BankNo          string                `db:"BankNo"`
	NetDebtAmount   float64               `db:"NetDebtAmount"`
	UserCode        string                `db:"UserCode"`
	ChqIns          []ListChqInModel      `db:"ChqIns"`
	CreditCards     []ListCreditCardModel `db:"CreditCards"`
	PosSubs         []NewPosItemModel     `db:"PosSubs"`
}

type NewPosItemModel struct {
	ItemCode     string  `db:"ItemCode"`
	ItemName     string  `db:"ItemName"`
	WHCode       string  `db:"WHCode"`
	ShelfCode    string  `db:"ShelfCode"`
	Qty          float64 `db:"Qty"`
	Price        float64 `db:"Price"`
	DiscountWord string  `db:"DiscountWord"`
	UnitCode     string  `db:"UnitCode"`
	LineNumber   int     `db:"LineNumber"`
	BarCode      string  `db:"BarCode"`
	AverageCost  float64 `db:"AverageCost"`
	PackingRate1 float64 `db:"PackingRate1"`
}

type ListChqInModel struct {
	ChqNumber      string  `db:"ChqNumber"`
	BankCode       string  `db:"BankCode"`
	BankBranchCode string  `db:"BankBranchCode"`
	BookNo         string  `db:"BookNo"`
	ReceiveDate    string  `db:"ReceiveDate"`
	DueDate        string  `db:"DueDate"`
	Status         int     `db:"Status"`
	Amount         float64 `db:"Amount"`
	Balance        float64 `db:"Balance"`
	RefChqRowOrder int     `db:"RefChqRowOrder"`
	StatusDate     string  `db:"StatusDate"`
	StatusDocNo    string  `db:"StatusDocNo"`
}

type ListCreditCardModel struct {
	BankCode       string  `db:"BankCode"`
	CreditCardNo   string  `db:"CreditCardNo"`
	ReceiveDate    string  `db:"ReceiveDate"`
	DueDate        string  `db:"DueDate"`
	BookNo         string  `db:"BookNo"`
	Status         int     `db:"Status"`
	StatusDate     string  `db:"StatusDate"`
	StatusDocNo    string  `db:"StatusDocNo"`
	BankBranchCode string  `db:"BankBranchCode"`
	Amount         float64 `db:"Amount"`
	MyDescription  string  `db:"MyDescription"`
	CreditType     string  `db:"CreditType"`
	ConfirmNo      string  `db:"ConfirmNo"`
	ChargeAmount   float64 `db:"ChargeAmount"`
}

type posRepository struct{ db *sqlx.DB }

func NewPosRepository(db *sqlx.DB) pos.Repository {
	return &posRepository{db}
}

func (repo *posRepository) NewPos(ctx context.Context, req *pos.NewPosTemplate) (resp pos.NewPosResponseTemplate, err error) {
	var check_doc_exist int
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var sum_pay_amount float64
	var sum_remain_amount float64
	var tax_rate float64
	var pos_tax_type int

	var exchange_rate float64
	var save_form int
	var is_complete_save int
	var deposit_inc_tax int
	var home_amount float64
	var bill_balance float64
	var pos_status int
	var my_description_recmoney string

	var line_number int
	var item_amount float64
	var my_type int
	var cn_qty float64
	var packing_rate_2 float64
	var item_home_amount float64
	var item_net_amount float64
	var sum_of_cost float64

	def := config.Default{}
	def = config.LoadDefaultData("config/config.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)
	fmt.Println("DocDate = ", req.DocDate)
	count_item_qty = 0
	count_item_unit = 0

	tax_rate = def.TaxRateDefault
	pos_tax_type = def.PosTaxType

	sum_pay_amount = (req.SumCashAmount + req.SumCreditAmount + req.SumChqAmount + req.SumBankAmount + req.CoupongAmount)

	for _, sub_item := range req.PosSubs {
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

	sqlexist := `select count(docno) as check_exist from dbo.bcarinvoice where docno = ?`
	fmt.Println("DocNo =", req.DocNo)
	err = repo.db.Get(&check_doc_exist, sqlexist, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return resp, err
	}

	fmt.Println("Len Chq =", len(req.ChqIns))

	switch {
	case req.DocNo == "":
		fmt.Println("error =", "Docno is null")
		return resp, errors.New("Docno is null")
	case req.ArCode == "":
		fmt.Println("error =", "Arcode is null")
		return resp, errors.New("Arcode is null")
	case req.DocDate == "":
		fmt.Println("error =", "Docdate is null")
		return resp, errors.New("Docdate is null")
	case count_item == 0:
		fmt.Println("error =", "Docno is not have item")
		return resp, errors.New("Docno is not have item")
	case (req.SumCashAmount == 0 && req.SumCreditAmount == 0 && req.SumChqAmount == 0 && req.SumBankAmount == 0):
		fmt.Println("error =", "Docno not set money to another type payment")
		return resp, errors.New("Docno not set money to another type payment")
	case req.SumOfItemAmount == 0:
		fmt.Println("error =", "Sumofitemamount = 0")
		return resp, errors.New("Sumofitemamount = 0")
	case count_item_qty > 0:
		fmt.Println("error =", "Docno is null")
		return resp, errors.New("Item not have qty")
	case count_item_unit > 0:
		fmt.Println("error =", "Item not have qty")
		return resp, errors.New("Item not have unitcode")
	case sum_pay_amount > req.TotalAmount:
		fmt.Println("error =", "Rec money is over totalamount")
		return resp, errors.New("Rec money is over totalamount")
	case sum_item_amount != sum_pay_amount:
		fmt.Println("error =", "Rec money is less than totalamount")
		return resp, errors.New("Rec money is less than totalamount")
	case (req.MachineCode == "" || req.MachineNo == "" || req.ShiftNo == 0 || req.ShiftCode == "" || req.CashierCode == ""):
		fmt.Println("error =", "Docno not have pos data", req.MachineCode, req.MachineNo, req.ShiftNo, req.ShiftCode, req.CashierCode)
		return resp, errors.New("Docno not have pos data")
	case req.SumChqAmount != 0 && len(req.ChqIns) == 0:
		fmt.Println("error =", "Docno not have chq data")
		return resp, errors.New("Docno not have chq data")
	case req.SumCreditAmount != 0 && len(req.CreditCards) == 0:
		fmt.Println("error =", "Docno not have credit card data")
		return resp, errors.New("Docno not have credit card data")
	}

	before_tax_amount, tax_amount, total_amount := calcTaxItem(pos_tax_type, tax_rate, req.AfterDiscount)

	sum_remain_amount = total_amount - sum_pay_amount

	if sum_remain_amount != 0 {
		return resp, errors.New("Docno have remain money to paid")
	}

	exchange_rate = def.ExchangeRateDefault
	save_form = def.PosSaveForm

	is_complete_save = 1
	deposit_inc_tax = 1
	pos_status = 1

	req.NetDebtAmount = req.TotalAmount
	home_amount = req.TotalAmount
	bill_balance = total_amount
	depart_code := "S01-00-00"

	discount_amount, err := strconv.ParseFloat(req.DiscountWord, 64)
	if err != nil {
		fmt.Println("error =", err.Error())
	}

	fmt.Println("check_doc_exist = ", check_doc_exist)

	if (check_doc_exist == 0) {
		sql := `set dateformat dmy     insert into dbo.bcarinvoice(DocNo,DocDate,ArCode,TaxType,CashierCode,ShiftNo,MachineNo,MachineCode,PosStatus,BillTime,GrandTotal,CoupongAmount,ChangeAmount,DepartCode,SaleCode,TaxRate,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,NetDebtAmount,HomeAmount,BillBalance,ExchangeRate,IsCompleteSave,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,convert(varchar(10), GETDATE(), 108),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
		fmt.Println("sql insert = ", sql)
		id, err := repo.db.Exec(sql, req.DocNo, req.DocDate, req.ArCode, pos_tax_type, req.CashierCode, req.ShiftNo, req.MachineNo, req.MachineCode, pos_status, total_amount, req.CoupongAmount, req.ChangeAmount, depart_code, req.SaleCode, tax_rate, req.SumOfItemAmount, req.DiscountWord, discount_amount, req.AfterDiscount, before_tax_amount, tax_amount, req.TotalAmount, req.SumCashAmount, req.SumChqAmount, req.SumCreditAmount, req.SumBankAmount, deposit_inc_tax, req.NetDebtAmount, home_amount, bill_balance, exchange_rate, is_complete_save, req.UserCode)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}

		resp.Id, err = id.LastInsertId()
	} else {
		sql := `set dateformat dmy     update dbo.bcarinvoice set DocDate=?,ArCode=?,TaxType=?,CashierCode=?,ShiftNo=?,MachineNo=?,MachineCode=?,GrandTotal=?,CoupongAmount=?,ChangeAmount=?,SaleCode=?,TaxRate=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,NetDebtAmount=?,HomeAmount=?,BillBalance=?,LastEditorCode=?,LastEditDateT=getdate() where DocNo=?`
		fmt.Println("sql update = ", sql)
		id, err := repo.db.Exec(sql, req.DocDate, req.ArCode, pos_tax_type, req.CashierCode, req.ShiftNo, req.MachineNo, req.MachineCode, total_amount, req.CoupongAmount, req.ChangeAmount, req.SaleCode, tax_rate, req.SumOfItemAmount, req.DiscountWord, discount_amount, req.AfterDiscount, before_tax_amount, tax_amount, req.TotalAmount, req.SumCashAmount, req.SumChqAmount, req.SumCreditAmount, req.SumBankAmount, req.NetDebtAmount, home_amount, bill_balance, req.UserCode, req.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}

		resp.Id, err = id.LastInsertId()
	}

	sql_del_sub := `delete dbo.bcarinvoicesub where docno = ?`
	_, err = repo.db.Exec(sql_del_sub, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return resp, err
	}

	for _, item := range req.PosSubs {
		fmt.Println("ItemSub")
		item_discount_amount, err := strconv.ParseFloat(item.DiscountWord, 64)

		item_amount = item.Qty * (item.Price - item_discount_amount)

		my_type = def.PosMyType
		cn_qty = item.Qty
		item.LineNumber = line_number

		if (item.PackingRate1 == 0) {
			item.PackingRate1 = 1
		}
		packing_rate_2 = 1

		switch {
		case pos_tax_type == 0:
			item_home_amount = item_amount
			item_net_amount = item_amount
		case pos_tax_type == 1:
			taxamount := ToFixed(item_amount-((item_amount*100)/(100+float64(tax_rate))), 2)
			beforetaxamount := ToFixed(item_amount-taxamount, 2)
			item_home_amount = beforetaxamount
			item_net_amount = beforetaxamount
		case pos_tax_type == 2:
			item_home_amount = item_amount
			item_net_amount = item_amount
		}

		sum_of_cost = item.AverageCost * item.Qty

		sqlsub := `set dateformat dmy      insert into dbo.BCArInvoiceSub(MyType,DocNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, LineNumber, BarCode, POSSTATUS, AVERAGECOST, PackingRate1, PackingRate2) values(?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = repo.db.Exec(sqlsub, my_type, req.DocNo, pos_tax_type, item.ItemCode, req.DocDate, req.ArCode, depart_code, req.SaleCode, "MobileApp", item.ItemName, item.WHCode, item.ShelfCode, cn_qty, item.Qty, item.Price, item.DiscountWord, item_discount_amount, item_amount, item_net_amount, item_home_amount, sum_of_cost, item.UnitCode, item.LineNumber, item.BarCode, pos_status, item.AverageCost, item.PackingRate1, packing_rate_2)
		fmt.Println("sqlsub = ", sqlsub, my_type, req.DocNo, pos_tax_type, item.ItemCode, req.DocDate, req.ArCode, depart_code, req.SaleCode, "MobileApp", item.ItemName, item.WHCode, item.ShelfCode, cn_qty, item.Qty, item.Price, item.DiscountWord, item_discount_amount, item_amount, item_net_amount, item_home_amount, sum_of_cost, item.UnitCode, item.LineNumber, item.BarCode, pos_status, item.AverageCost, item.PackingRate1, packing_rate_2)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}

		sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
		_, err = repo.db.Exec(sqlprocess, item.ItemCode)
		fmt.Println("sqlprocess = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			fmt.Println(err.Error())
		}

		line_number = line_number + 1
	}

	sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
	_, err = repo.db.Exec(sqlrecdel, req.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return resp, err
	}

	my_description_recmoney = "ขายเงินสด"

	fmt.Println("RecMoney")
	var linenumber int

	if (req.SumCashAmount != 0) { //subs.PaymentType == 0:
		fmt.Println("SumCashAmount")
		sqlrec := `set dateformat dmy      insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?)`
		_, err = repo.db.Exec(sqlrec, req.DocNo, req.DocDate, req.ArCode, exchange_rate, req.SumCashAmount, 0, save_form, linenumber, depart_code, req.SaleCode, my_description_recmoney)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}
	}
	//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
	if (req.SumCreditAmount != 0) {
		var crd_credit_type string
		var crd_confirm_no string
		var crd_credit_no string
		var crd_bank_code string
		var crd_bank_branch_code string

		fmt.Println("SumCreditAmount")
		if (req.SumCashAmount != 0) {
			linenumber = 1
		} else {
			linenumber = 0
		}

		if len(req.CreditCards) != 0 {
			crd_credit_type = req.CreditCards[0].CreditType
			crd_confirm_no = req.CreditCards[0].ConfirmNo
			crd_credit_no = req.CreditCards[0].CreditCardNo
			crd_bank_code = req.CreditCards[0].BankCode
			crd_bank_branch_code = req.CreditCards[0].BankBranchCode

		}

		sqlrec := `set dateformat dmy      insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,SaveFrom,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,BankBranchCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = repo.db.Exec(sqlrec, req.DocNo, req.DocDate, req.ArCode, exchange_rate, req.SumCreditAmount, req.SumCreditAmount, 1, save_form, crd_credit_type, crd_confirm_no, linenumber, crd_credit_no, crd_bank_code, crd_bank_branch_code, depart_code, req.SaleCode, my_description_recmoney, req.DocDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}
	}

	//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
	if (req.SumChqAmount != 0) {
		var chq_book_no string
		var chq_bank_code string
		var chq_bank_branch_code string

		fmt.Println("SumChqAmount")
		if (req.SumCashAmount != 0 && req.SumCreditAmount != 0) {
			linenumber = 2
		} else if ((req.SumCashAmount != 0 && req.SumCreditAmount == 0)) {
			linenumber = 1
		} else if ((req.SumCashAmount == 0 && req.SumCreditAmount != 0)) {
			linenumber = 1
		} else if ((req.SumCashAmount == 0 && req.SumCreditAmount == 0)) {
			linenumber = 0
		}

		if len(req.ChqIns) != 0 {
			chq_book_no = req.ChqIns[0].BookNo
			chq_bank_code = req.ChqIns[0].BankCode
			chq_bank_branch_code = req.ChqIns[0].BankBranchCode

		}

		sqlrec := `set dateformat dmy      insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,BankCode,DepartCode,SaleCode,BankBranchCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = repo.db.Exec(sqlrec, req.DocNo, req.DocDate, req.ArCode, exchange_rate, req.SumChqAmount, 2, save_form, linenumber, chq_book_no, chq_bank_code, depart_code, req.SaleCode, chq_bank_branch_code, my_description_recmoney, req.DocDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}
	}

	//case dp.SumBankAmount != 0: //subs.PaymentType == 3:
	if (req.SumBankAmount != 0) {
		fmt.Println("SumBankAmount")
		if (req.SumCashAmount != 0 && req.SumCreditAmount != 0 && req.SumChqAmount != 0) {
			linenumber = 3
		} else if (req.SumCashAmount != 0 && req.SumCreditAmount == 0 && req.SumChqAmount != 0) {
			linenumber = 2
		} else if (req.SumCashAmount == 0 && req.SumCreditAmount != 0 && req.SumChqAmount != 0) {
			linenumber = 2
		} else if (req.SumCashAmount != 0 && req.SumCreditAmount != 0 && req.SumChqAmount == 0) {
			linenumber = 2
		} else if (req.SumCashAmount != 0 && req.SumCreditAmount != 0 && req.SumChqAmount == 0) {
			linenumber = 2
		} else if (req.SumCashAmount != 0 && req.SumCreditAmount == 0 && req.SumChqAmount == 0) {
			linenumber = 1
		} else if (req.SumCashAmount == 0 && req.SumCreditAmount != 0 && req.SumChqAmount == 0) {
			linenumber = 1
		} else if (req.SumCashAmount == 0 && req.SumCreditAmount == 0 && req.SumChqAmount != 0) {
			linenumber = 1
		} else if (req.SumCashAmount == 0 && req.SumCreditAmount == 0 && req.SumChqAmount == 0) {
			linenumber = 0
		}

		sqlrec := `set dateformat dmy      insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,DepartCode,SaleCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = repo.db.Exec(sqlrec, req.DocNo, req.DocDate, req.ArCode, exchange_rate, req.SumBankAmount, 3, save_form, linenumber, req.BankNo, depart_code, req.SaleCode, my_description_recmoney, req.DocDate, req.DocDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}
	}

	if (req.SumChqAmount != 0) {
		sqlchqdel := `delete dbo.BCChqIn where docno = ?`
		_, err = repo.db.Exec(sqlchqdel, req.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}

		for _, c := range req.ChqIns {
			if ((c.ReceiveDate == "") || (c.DueDate == "")) {
				c.ReceiveDate = req.DocDate
				c.DueDate = req.DocDate
			}

			sqldep := `set dateformat dmy      insert into dbo.bcchqin(BankCode,ChqNumber,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = repo.db.Exec(sqldep, c.BankCode, c.ChqNumber, req.DocNo, req.ArCode, req.SaleCode, c.ReceiveDate, c.DueDate, c.BookNo, c.Status, save_form, c.StatusDate, c.StatusDocNo, depart_code, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, exchange_rate, c.RefChqRowOrder)
			if err != nil {
				fmt.Println("Error = ", err.Error())
				return resp, err
			}
		}
	}

	if (req.SumCreditAmount != 0) {
		sqlcrddel := `delete dbo.BCCreditCard where docno = ?`
		_, err = repo.db.Exec(sqlcrddel, req.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return resp, err
		}

		for _, d := range req.CreditCards {
			fmt.Println("ReceiveDate =", d.ReceiveDate, d.DueDate)
			if (d.DueDate == "" || d.DueDate == "01/01/1900") {
				d.DueDate = req.DocDate
			}

			if (d.ReceiveDate == "" || d.ReceiveDate == "01/01/1900") {
				d.ReceiveDate = req.DocDate
			}

			if (d.StatusDocNo != "" && d.StatusDate == "") {
				d.StatusDate = req.DocDate
			}

			sqlcrd := `set dateformat dmy      insert into dbo.bccreditcard(BankCode,CreditCardNo,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,MyDescription,ExchangeRate,CreditType,ConfirmNo,ChargeAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
			fmt.Println("sqlcrd = ", sqlcrd)
			_, err = repo.db.Exec(sqlcrd, d.BankCode, d.CreditCardNo, req.DocNo, req.ArCode, req.SaleCode, d.ReceiveDate, d.DueDate, d.BookNo, d.Status, save_form, d.StatusDate, d.StatusDocNo, depart_code, d.BankBranchCode, d.Amount, my_description_recmoney, exchange_rate, d.CreditType, d.ConfirmNo, d.ChargeAmount, req.UserCode)
			if err != nil {
				fmt.Println("Error = ", err.Error())
				return resp, err
			}
		}

	}

	return resp, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func calcTaxItem(taxtype int, taxrate float64, afterdiscountamount float64) (beforetaxamount float64, taxamount float64, totalamount float64) {
	switch taxtype {
	case 0:
		beforetaxamount = ToFixed(afterdiscountamount, 2)
		taxamount = ToFixed(((afterdiscountamount*(100+float64(taxrate)))/(100))-afterdiscountamount, 2)
		totalamount = ToFixed(beforetaxamount+taxamount, 2)
	case 1:
		taxamount = ToFixed(afterdiscountamount-((afterdiscountamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = ToFixed(afterdiscountamount-taxamount, 2)
		totalamount = ToFixed(afterdiscountamount, 2)
	case 2:
		beforetaxamount = ToFixed(afterdiscountamount, 2)
		taxamount = 0
		totalamount = ToFixed(afterdiscountamount, 2)
	}

	fmt.Println("taxtype,taxrate,beforetaxamount,taxamount,totalamount", taxtype, taxrate, beforetaxamount, taxamount, totalamount)

	return beforetaxamount, taxamount, totalamount
}
