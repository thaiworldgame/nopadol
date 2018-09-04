package sqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/print"
	"strconv"
	"bufio"
	"github.com/knq/escpos"
	"fmt"
	"strings"
	"time"
	"github.com/mrtomyum/nopadol/hw"
	"bytes"
	"github.com/mrtomyum/nopadol/pos"
	//"os"
	"net"
)

var sql_dbc *sqlx.DB

type PosSlipModel struct {
	Id              int                `db:"Id"`
	TaxId           string             `db:"TaxId"`
	DocNo           string             `db:"DocNo"`
	DocDate         string             `db:"DocDate"`
	TaxNo           string             `db:"TaxNo"`
	TaxDate         string             `db:"TaxDate"`
	PosStatus       int                `db:"PosStatus"`
	ArCode          string             `db:"ArCode"`
	ArName          string             `db:"ArName"`
	SaleCode        string             `db:"SaleCode"`
	SaleName        string             `db:"SaleName"`
	ShiftCode       string             `db:"ShiftCode"`
	CashierCode     string             `db:"CashierCode"`
	ShiftNo         int                `db:"ShiftNo"`
	MachineNo       string             `db:"MachineNo"`
	MachineCode     string             `db:"MachineCode"`
	CoupongAmount   float64            `db:"CoupongAmount"`
	ChangeAmount    float64            `db:"ChangeAmount"`
	ChargeAmount    float64            `db:"ChargeAmount"`
	TaxType         int                `db:"TaxType"`
	SumOfItemAmount float64            `db:"SumOfItemAmount"`
	DiscountWord    string             `db:"DiscountWord"`
	AfterDiscount   float64            `db:"AfterDiscount"`
	BeforeTaxAmount float64            `db:"BeforeTaxAmount"`
	TaxAmount       float64            `db:"TaxAmount"`
	TotalAmount     float64            `db:"TotalAmount"`
	SumCashAmount   float64            `db:"SumCashAmount"`
	SumChqAmount    float64            `db:"SumChqAmount"`
	SumCreditAmount float64            `db:"SumCreditAmount"`
	SumBankAmount   float64            `db:"SumBankAmount"`
	BankNo          string             `db:"BankNo"`
	NetDebtAmount   float64            `db:"NetDebtAmount"`
	IsCancel        int                `db:"IsCancel"`
	IsConfirm       int                `db:"IsConfirm"`
	CreatorCode     string             `db:"CreatorCode"`
	CreateDateTime  string             `db:"CreateDateTime"`
	LastEditorCode  string             `db:"LastEditorCode"`
	LastEditDateT   string             `db:"LastEditDateT"`
	PosSubs         []PosItemSlipModel `db:"PosSubs"`
}

type PosItemSlipModel struct {
	ItemCode       string  `db:"ItemCode"`
	ItemName       string  `db:"ItemName"`
	WHCode         string  `db:"WHCode"`
	ShelfCode      string  `db:"ShelfCode"`
	Qty            float64 `db:"Qty"`
	Price          float64 `db:"Price"`
	DiscountWord   string  `db:"DiscountWord"`
	DiscountAmount float64 `db:"DiscountAmount"`
	UnitCode       string  `db:"UnitCode"`
	LineNumber     int     `db:"LineNumber"`
	BarCode        string  `db:"BarCode"`
	AverageCost    float64 `db:"AverageCost"`
	PackingRate1   float64 `db:"PackingRate1"`
}

//type PosSlipResponseModel struct {
//	DocNo string `json:"doc_no"`
//}

type printRepository struct{ db *sqlx.DB }

func SqlConn(db_host string, db_name string, db_user string, db_pass string) (msdb *sqlx.DB, err error) {
	port := "1433"
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	fmt.Println("SqlConn", dsn)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error ")
	}

	return msdb, nil
}

func NewPrintRepository(db *sqlx.DB) print.Repository {
	return &printRepository{db}
}

func (repo *printRepository) PosSlip(req *print.PosSlipRequestTemplate) (resp interface{}, err error) {
	host := "192.168.0.247:9100"

	f, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(w)

	s := PosSlipModel{}

	sql := `select a.roworder as Id,a.DocNo,a.DocDate,isnull(a.TaxNo,'') as TaxNo,isnull(a.docdate,'') as TaxDate,a.PosStatus,a.ArCode,isnull(b.name1,'') as ArName,a.SaleCode,isnull(c.name,'') as SaleName,isnull(ShiftCode,'') as ShiftCode,CashierCode,isnull(ShiftNo,'') as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,CoupongAmount,ChangeAmount,ChargeAmount,a.TaxType,SumOfItemAmount,isnull(a.DiscountWord,'') as DiscountWord,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount ,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,'' as BankNo,NetDebtAmount,IsCancel,IsConfirm,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT from dbo.bcarinvoice a  left join dbo.bcar b on a.arcode = b.code left join dbo.bcsale c  on a.salecode = c.code where a.docno = ?`
	err = repo.db.Get(&s, sql, req.DocNo)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return resp, err
	}

	sql_sub := `select a.ItemCode,a.ItemName,a.WHCode,a.ShelfCode,a.Qty,a.Price,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.UnitCode,isnull(a.BarCode,'') as BarCode,isnull(a.AverageCost,0) as AverageCost,a.PackingRate1,a.LineNumber from dbo.bcarinvoicesub a left join dbo.bcitem b on a.itemcode = b.code where a.docno = ? order by a.linenumber`
	err = repo.db.Select(&s.PosSubs, sql_sub, s.DocNo)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return resp, err
	}

	pt := hw.PosPrinter{p, w}
	pt.Init()
	pt.SetLeftMargin(20)

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)

	//now := time.Now()
	//fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	//DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")
	DocDate := now.Format("02-01-2006 ")

	//////////////////////////////////////////////////////////////////////////////////////
	pt.SetCharaterCode(26)
	pt.SetAlign("center")
	pt.SetTextSize(0, 0)
	pt.WriteStringLines("TAX INV (ABB)")
	pt.LineFeed()
	pt.SetTextSize(0, 0)
	pt.SetAlign("left")
	pt.WriteStringLines("NO: " + s.DocNo + "          " + "DATE: " + DocDate + "\n")
	pt.SetAlign("center")
	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	pt.SetAlign("left")
	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + s.ShiftCode + "\n")
	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	pt.SetAlign("left")
	pt.WriteStringLines("CS : " + s.SaleCode + "/" + s.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////
	pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////

	var CountItem int64
	var CountQty float64
	for _, subcount := range s.PosSubs {
		CountItem = CountItem + 1
		CountQty = CountQty + subcount.Qty
	}

	fmt.Println("CountItem =", CountItem, CountQty)
	///////////////////////////////////////////////////////////////////////////////////
	pt.SetAlign("left")
	for _, sub := range s.PosSubs {
		var vDiffEmpty int
		var vItemPriceAmount string
		var vItemAmount float64

		pt.SetFont("A")

		pt.WriteStringLines(" " + sub.ItemName + "\n")

		vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)

		//vItemPriceAmount = " " + strconv.FormatFloat(sub.Price, 'f', -1, 64) + " X " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode
		vItemPriceAmount = " " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
		vLen := len(vItemPriceAmount)
		vDiff := 25 - (vLen / 3)

		if (vDiff < 0) {
			vDiffEmpty = 0
		} else {
			vDiffEmpty = vDiff
		}

		fmt.Println("ItemName=", sub.ItemName)
		fmt.Println("Len", vLen/3)
		fmt.Println("Diff ", vDiff)
		pt.WriteStringLines(vItemPriceAmount + strings.Repeat(" ", vDiffEmpty))

		pt.WriteStringLines("   ")
		pt.WriteStringLines(CommaFloat(vItemAmount) + "\n\n")
		//pt.FormfeedN(3)
	}
	makeline(pt)
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	pt.WriteStringLines(" " + strconv.Itoa(int(CountItem)) + " รายการ " + strconv.Itoa(int(CountQty)) + " ชิ้น\n")
	pt.WriteStringLines("TOTAL: ")
	pt.WriteStringLines("                           ")
	//pt.WriteStringLines(strconv.FormatFloat(s.TotalAmount, 'f', 2, 64)+"\n")
	pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	if (s.CoupongAmount != 0) {
		pt.WriteStringLines("COUPON: " + "                          " + CommaFloat(s.CoupongAmount) + "\n")
	}
	if (s.SumCashAmount != 0) {
		pt.WriteStringLines("CASH: " + "                            " + CommaFloat(s.SumCashAmount) + "\n")
	}
	if (s.SumCreditAmount != 0) {
		pt.WriteStringLines("CREDIT: " + "                          " + CommaFloat(s.SumCreditAmount) + "\n")
	}
	if (s.ChangeAmount != 0) {
		pt.WriteStringLines("CHANGE: " + "                            " + CommaFloat(s.ChangeAmount) + "\n")
	}

	pt.SetFont("A")
	pt.SetAlign("center")
	pt.WriteStringLines("ขอบคุณที่มาใช้บริการ" + "\n")
	pt.WriteStringLines("เปลี่ยนสินค้าภายใน 30 วัน (รหัส [P] ไม่รับคืน)" + "\n")
	pt.WriteStringLines("สอบถามโทร. 0-5324-0377 8.00-17.00 ทุกวัน" + "\n")
	makeline(pt)

	pt.Formfeed()
	pt.Cut()
	pt.OpenCashBox()
	pt.End()

	return req.DocNo, err
}

func (repo *printRepository) PosDriveThruSlip(req *print.PosDriveThruSlipRequestTemplate) (resp interface{}, err error) {
	sql_db, err := SqlConn("192.168.2.100", "POS_Test", "vbuser", "132")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	sql_dbc = sql_db

	host := "192.168.0.247:9100"

	f, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(w)

	s := PosSlipModel{}

	sql := `select a.roworder as Id,a.DocNo,a.DocDate,isnull(a.TaxNo,'') as TaxNo,isnull(a.docdate,'') as TaxDate,a.PosStatus,a.ArCode,isnull(b.name1,'') as ArName,a.SaleCode,isnull(c.name,'') as SaleName,isnull(ShiftCode,'') as ShiftCode,CashierCode,isnull(ShiftNo,'') as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,CoupongAmount,ChangeAmount,ChargeAmount,a.TaxType,SumOfItemAmount,isnull(a.DiscountWord,'') as DiscountWord,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount ,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,'' as BankNo,NetDebtAmount,IsCancel,IsConfirm,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT from dbo.bcarinvoice a  left join dbo.bcar b on a.arcode = b.code left join dbo.bcsale c  on a.salecode = c.code where a.docno = ?`
	err = sql_dbc.Get(&s, sql, req.DocNo)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, err
	}

	sql_sub := `select a.ItemCode,a.ItemName,a.WHCode,a.ShelfCode,a.Qty,a.Price,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.UnitCode,isnull(a.BarCode,'') as BarCode,isnull(a.AverageCost,0) as AverageCost,a.PackingRate1,a.LineNumber from dbo.bcarinvoicesub a left join dbo.bcitem b on a.itemcode = b.code where a.docno = ? order by a.linenumber`
	err = sql_dbc.Select(&s.PosSubs, sql_sub, s.DocNo)
	if err != nil {
		fmt.Println("err sub= ", err.Error())
		return nil, err
	}

	pt := hw.PosPrinter{p, w}
	pt.Init()
	pt.SetLeftMargin(20)

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	//now := time.Now()
	//fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	//DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")
	DocDate := now.Format("02-01-2006 ")

	//////////////////////////////////////////////////////////////////////////////////////
	pt.SetCharaterCode(26)
	pt.SetAlign("center")
	pt.SetTextSize(0, 0)
	pt.WriteStringLines("TAX INV (ABB)")
	pt.LineFeed()
	pt.SetTextSize(0, 0)
	pt.SetAlign("left")
	pt.WriteStringLines("NO: " + s.DocNo + "          " + "DATE: " + DocDate + "\n")
	pt.SetAlign("center")
	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	pt.SetAlign("left")
	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + s.ShiftCode + "\n")
	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	pt.SetAlign("left")
	//pt.WriteStringLines("CS : " + s.SaleCode + "/" + s.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////
	pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////

	var CountItem int64
	var CountQty float64
	for _, subcount := range s.PosSubs {
		CountItem = CountItem + 1
		CountQty = CountQty + subcount.Qty
	}

	fmt.Println("CountItem =", CountItem, CountQty)
	///////////////////////////////////////////////////////////////////////////////////
	pt.SetAlign("left")
	for _, sub := range s.PosSubs {
		var vDiffEmpty int
		var vItemPriceAmount string
		var vItemAmount float64

		pt.SetFont("A")

		pt.WriteStringLines(" " + sub.ItemName + "\n")

		vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)

		//vItemPriceAmount = " " + strconv.FormatFloat(sub.Price, 'f', -1, 64) + " X " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode
		vItemPriceAmount = " " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
		vLen := len(vItemPriceAmount)
		vDiff := 25 - (vLen / 3)

		if (vDiff < 0) {
			vDiffEmpty = 0
		} else {
			vDiffEmpty = vDiff
		}

		fmt.Println("ItemName=", sub.ItemName)
		fmt.Println("Len", vLen/3)
		fmt.Println("Diff ", vDiff)
		pt.WriteStringLines(vItemPriceAmount + strings.Repeat(" ", vDiffEmpty))

		pt.WriteStringLines("   ")
		pt.WriteStringLines(CommaFloat(vItemAmount) + "\n\n")
		//pt.FormfeedN(3)
	}
	makeline(pt)
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	pt.WriteStringLines(" " + strconv.Itoa(int(CountItem)) + " รายการ " + strconv.Itoa(int(CountQty)) + " ชิ้น\n")
	pt.WriteStringLines("TOTAL: ")
	pt.WriteStringLines("                           ")
	//pt.WriteStringLines(strconv.FormatFloat(s.TotalAmount, 'f', 2, 64)+"\n")
	pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	if (s.CoupongAmount != 0) {
		pt.WriteStringLines("COUPON: " + "                          " + CommaFloat(s.CoupongAmount) + "\n")
	}
	if (s.SumCashAmount != 0) {
		pt.WriteStringLines("CASH: " + "                            " + CommaFloat(s.SumCashAmount) + "\n")
	}
	if (s.SumCreditAmount != 0) {
		pt.WriteStringLines("CREDIT: " + "                          " + CommaFloat(s.SumCreditAmount) + "\n")
	}
	if (s.ChangeAmount != 0) {
		pt.WriteStringLines("CHANGE: " + "                            " + CommaFloat(s.ChangeAmount) + "\n")
	}

	pt.SetFont("A")
	pt.SetAlign("center")
	pt.WriteStringLines("ขอบคุณที่มาใช้บริการ" + "\n")
	pt.WriteStringLines("เปลี่ยนสินค้าภายใน 30 วัน (รหัส [P] ไม่รับคืน)" + "\n")
	pt.WriteStringLines("สอบถามโทร. 0-5324-0377 8.00-17.00 ทุกวัน" + "\n")
	makeline(pt)

	pt.Formfeed()
	pt.Cut()
	pt.OpenCashBox()
	pt.End()

	repo.PosDriveThruSlipCopy(s)

	return req.DocNo, err
}

func (repo *printRepository) PosDriveThruSlipCopy(req PosSlipModel) () {

	host := "192.168.0.247:9100"

	f, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(w)

	pt := hw.PosPrinter{p, w}
	pt.Init()
	pt.SetLeftMargin(20)

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	DocDate := now.Format("02-01-2006 ")

	//////////////////////////////////////////////////////////////////////////////////////
	pt.SetCharaterCode(26)
	pt.SetAlign("center")
	pt.SetTextSize(0, 0)
	pt.WriteStringLines("TAX INV (ABB)     (สำเนา)")
	pt.LineFeed()
	pt.SetTextSize(0, 0)
	pt.SetAlign("left")
	pt.WriteStringLines("NO: " + req.DocNo + "          " + "DATE: " + DocDate + "\n")
	pt.SetAlign("center")
	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	pt.SetAlign("left")
	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + req.ShiftCode + "\n")
	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	pt.SetAlign("left")
	pt.WriteStringLines("CS : " + req.SaleCode + "/" + req.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////
	pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////

	var CountItem int64
	var CountQty float64
	for _, subcount := range req.PosSubs {
		CountItem = CountItem + 1
		CountQty = CountQty + subcount.Qty
	}

	fmt.Println("CountItem =", CountItem, CountQty)
	///////////////////////////////////////////////////////////////////////////////////
	pt.SetAlign("left")
	for _, sub := range req.PosSubs {
		var vDiffEmpty int
		var vItemPriceAmount string
		var vItemAmount float64

		pt.SetFont("A")

		pt.WriteStringLines(" " + sub.ItemName + "\n")

		vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)

		//vItemPriceAmount = " " + strconv.FormatFloat(sub.Price, 'f', -1, 64) + " X " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode
		vItemPriceAmount = " " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
		vLen := len(vItemPriceAmount)
		vDiff := 25 - (vLen / 3)

		if (vDiff < 0) {
			vDiffEmpty = 0
		} else {
			vDiffEmpty = vDiff
		}

		fmt.Println("ItemName=", sub.ItemName)
		fmt.Println("Len", vLen/3)
		fmt.Println("Diff ", vDiff)
		pt.WriteStringLines(vItemPriceAmount + strings.Repeat(" ", vDiffEmpty))

		pt.WriteStringLines("   ")
		pt.WriteStringLines(CommaFloat(vItemAmount) + "\n\n")
		//pt.FormfeedN(3)
	}
	makeline(pt)
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	pt.WriteStringLines(" " + strconv.Itoa(int(CountItem)) + " รายการ " + strconv.Itoa(int(CountQty)) + " ชิ้น\n")
	pt.WriteStringLines("TOTAL: ")
	pt.WriteStringLines("                           ")
	//pt.WriteStringLines(strconv.FormatFloat(s.TotalAmount, 'f', 2, 64)+"\n")
	pt.WriteStringLines(CommaFloat(req.NetDebtAmount) + "\n")
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	if (req.CoupongAmount != 0) {
		pt.WriteStringLines("COUPON: " + "                          " + CommaFloat(req.CoupongAmount) + "\n")
	}
	if (req.SumCashAmount != 0) {
		pt.WriteStringLines("CASH: " + "                            " + CommaFloat(req.SumCashAmount) + "\n")
	}
	if (req.SumCreditAmount != 0) {
		pt.WriteStringLines("CREDIT: " + "                          " + CommaFloat(req.SumCreditAmount) + "\n")
	}
	if (req.ChangeAmount != 0) {
		pt.WriteStringLines("CHANGE: " + "                            " + CommaFloat(req.ChangeAmount) + "\n")
	}

	pt.SetFont("A")
	pt.SetAlign("center")
	pt.WriteStringLines("ขอบคุณที่มาใช้บริการ" + "\n")
	pt.WriteStringLines("เปลี่ยนสินค้าภายใน 30 วัน (รหัส [P] ไม่รับคืน)" + "\n")
	pt.WriteStringLines("สอบถามโทร. 0-5324-0377 8.00-17.00 ทุกวัน" + "\n")
	makeline(pt)

	pt.Formfeed()
	pt.End()

}

//
//func (repo *printRepository) DriveThruPosSlip(req1 interface{}) (resp interface{}, err error) {
//	host := "192.168.0.247:9100"
//
//
//	req.Header.Add("Accept","application/json")
//	req.Body = bytes.NewBufferString(MyJsonAsAString)
//	res,err:=client.Do(req)
//
//
//	req.Body = bytes.NewBufferString({"namespaceadmins":"asj", "datalist":["avx","bdx","xnf"], "attributes":{ "name":"aj", "listedvalues":["12","14","13"], }, "id":12, "groups":["ab","kl"] })
//
//
//	req1 := PosSlipModel{}
//
//	req1.MachineCode =
//
//	f, err := net.Dial("tcp", host)
//	if err != nil {
//		return nil, err
//	}
//	defer f.Close()
//
//	w := bufio.NewWriter(f)
//	p := escpos.New(w)
//
//	//s := PosSlipModel{}
//	//
//	//sql := `select a.roworder as Id,a.DocNo,a.DocDate,isnull(a.TaxNo,'') as TaxNo,isnull(a.docdate,'') as TaxDate,a.PosStatus,a.ArCode,isnull(b.name1,'') as ArName,a.SaleCode,isnull(c.name,'') as SaleName,isnull(ShiftCode,'') as ShiftCode,CashierCode,isnull(ShiftNo,'') as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,CoupongAmount,ChangeAmount,ChargeAmount,a.TaxType,SumOfItemAmount,isnull(a.DiscountWord,'') as DiscountWord,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount ,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,'' as BankNo,NetDebtAmount,IsCancel,IsConfirm,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT from dbo.bcarinvoice a  left join dbo.bcar b on a.arcode = b.code left join dbo.bcsale c  on a.salecode = c.code where a.docno = ?`
//	//err = sql_dbc.Get(&s, sql, req)
//	//if err != nil {
//	//	fmt.Println("err = ", err.Error())
//	//	return resp, err
//	//}
//	//
//	//sql_sub := `select a.ItemCode,a.ItemName,a.WHCode,a.ShelfCode,a.Qty,a.Price,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.UnitCode,isnull(a.BarCode,'') as BarCode,isnull(a.AverageCost,0) as AverageCost,a.PackingRate1,a.LineNumber from dbo.bcarinvoicesub a left join dbo.bcitem b on a.itemcode = b.code where a.docno = ? order by a.linenumber`
//	//err = sql_dbc.Select(&s.PosSubs, sql_sub, s.DocNo)
//	//if err != nil {
//	//	fmt.Println("err sub= ", err.Error())
//	//	return resp, err
//	//}
//
//	fmt.Println("Item =", req)
//	layout := "2014-09-12"
//
//	doc_date, err := time.Parse(layout, s.DocDate)
//
//	fmt.Println("doc_date =", doc_date)
//
//	pt := hw.PosPrinter{p, w}
//	pt.Init()
//	pt.SetLeftMargin(20)
//
//	//loc, _ := time.LoadLocation("Asia/Bangkok")
//	//now := time.Now().In(loc)
//
//	now := time.Now()
//	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
//	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")
//
//	//////////////////////////////////////////////////////////////////////////////////////
//	pt.SetCharaterCode(26)
//	pt.SetAlign("center")
//	pt.SetTextSize(0, 0)
//	pt.WriteStringLines("TAX INV (ABB)")
//	pt.LineFeed()
//	pt.SetTextSize(0, 0)
//	pt.SetAlign("left")
//	pt.WriteStringLines("NO: " + s.DocNo + "          " + "DATE: " + DocDate + "\n")
//	pt.SetAlign("center")
//	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
//	pt.SetAlign("left")
//	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + s.ShiftCode + "\n")
//	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
//	pt.SetAlign("left")
//	pt.WriteStringLines("CS : " + s.SaleCode + "/" + s.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
//	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
//	makeline(pt)
//	///////////////////////////////////////////////////////////////////////////////////
//	var CountItem int64
//	var CountQty float64
//	for _, subcount := range s.PosSubs {
//		CountItem = CountItem + 1
//		CountQty = CountQty + subcount.Qty
//	}
//
//	fmt.Println("CountItem =", CountItem, CountQty)
//	///////////////////////////////////////////////////////////////////////////////////
//	pt.SetAlign("left")
//	for _, sub := range s.PosSubs {
//		var vDiffEmpty int
//		var vItemPriceAmount string
//		var vItemAmount float64
//
//		pt.SetFont("A")
//
//		pt.WriteStringLines(" " + sub.ItemName + "\n")
//
//		vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)
//
//		vItemPriceAmount = " " + strconv.FormatFloat(sub.Price, 'f', -1, 64) + " X " + strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode
//
//		vLen := len(vItemPriceAmount)
//		vDiff := 25 - (vLen / 3)
//
//		if (vDiff < 0) {
//			vDiffEmpty = 0
//		} else {
//			vDiffEmpty = vDiff
//		}
//
//		fmt.Println("ItemName=", sub.ItemName)
//		fmt.Println("Len", vLen/3)
//		fmt.Println("Diff ", vDiff)
//		pt.WriteStringLines(vItemPriceAmount + strings.Repeat(" ", vDiffEmpty))
//
//		pt.WriteStringLines("   ")
//		pt.WriteStringLines(CommaFloat(vItemAmount) + "\n\n")
//		//pt.FormfeedN(3)
//	}
//	makeline(pt)
//	////////////////////////////////////////////////////////////////////////////////////
//	pt.SetFont("A")
//	pt.WriteStringLines(" " + strconv.Itoa(int(CountItem)) + " รายการ " + strconv.Itoa(int(CountQty)) + " ชิ้น\n")
//	pt.WriteStringLines("TOTAL: ")
//	pt.WriteStringLines("                              ")
//	//pt.WriteStringLines(strconv.FormatFloat(s.TotalAmount, 'f', 2, 64)+"\n")
//	pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
//	////////////////////////////////////////////////////////////////////////////////////
//	pt.SetFont("A")
//	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
//	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
//	if (s.CoupongAmount != 0) {
//		pt.WriteStringLines("COUPON: " + "                              " + CommaFloat(s.CoupongAmount) + "\n")
//	}
//	if (s.SumCashAmount != 0) {
//		pt.WriteStringLines("CASH: " + "                               " + CommaFloat(s.SumCashAmount) + "\n")
//	}
//	if (s.SumCreditAmount != 0) {
//		pt.WriteStringLines("CREDIT: " + "                             " + CommaFloat(s.SumCreditAmount) + "\n")
//	}
//	if (s.ChangeAmount != 0) {
//		pt.WriteStringLines("CHANGE: " + "                               " + CommaFloat(s.ChangeAmount) + "\n")
//	}
//
//	pt.SetFont("A")
//	pt.SetAlign("center")
//	pt.WriteStringLines("ขอบคุณที่มาใช้บริการ" + "\n")
//	pt.WriteStringLines("เปลี่ยนสินค้าภายใน 30 วัน (รหัส [P] ไม่รับคืน)" + "\n")
//	pt.WriteStringLines("สอบถามโทร. 0-5324-0377 8.00-17.00 ทุกวัน" + "\n")
//	makeline(pt)
//
//	pt.Formfeed()
//	pt.Cut()
//	pt.OpenCashBox()
//	pt.End()
//
//	return "Print OK", err
//}

func makeline(pt hw.PosPrinter) {
	pt.SetTextSize(0, 0)
	pt.SetFont("A")
	pt.WriteStringLines("----------------------------------------------\n")
}

func CommaFloat(v float64) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

func map_posslip_template(x PosModel) pos.SearchPosByIdResponseTemplate {
	var subs []pos.NewPosItemTemplate
	var crds []pos.ListCreditCardTemplate
	return pos.SearchPosByIdResponseTemplate{
		Id:              x.Id,
		DocNo:           x.DocNo,
		DocDate:         x.DocDate,
		TaxNo:           x.TaxNo,
		TaxDate:         x.TaxDate,
		PosStatus:       x.PosStatus,
		ArCode:          x.ArCode,
		ArName:          x.ArName,
		SaleCode:        x.SaleCode,
		SaleName:        x.SaleName,
		ShiftCode:       x.ShiftCode,
		CashierCode:     x.CashierCode,
		ShiftNo:         x.ShiftNo,
		MachineNo:       x.MachineNo,
		MachineCode:     x.MachineCode,
		CoupongAmount:   x.CoupongAmount,
		ChangeAmount:    x.ChangeAmount,
		ChargeAmount:    x.ChargeAmount,
		TaxType:         x.TaxType,
		SumOfItemAmount: x.SumOfItemAmount,
		DiscountWord:    x.DiscountWord,
		AfterDiscount:   x.AfterDiscount,
		BeforeTaxAmount: x.BeforeTaxAmount,
		TaxAmount:       x.TaxAmount,
		TotalAmount:     x.TotalAmount,
		SumCashAmount:   x.SumCashAmount,
		SumChqAmount:    x.SumChqAmount,
		SumCreditAmount: x.SumCreditAmount,
		SumBankAmount:   x.SumBankAmount,
		BankNo:          x.BankNo,
		NetDebtAmount:   x.NetDebtAmount,
		IsCancel:        x.IsCancel,
		IsConfirm:       x.IsConfirm,
		CreatorCode:     x.CreatorCode,
		CreateDateTime:  x.CreateDateTime,
		LastEditorCode:  x.LastEditorCode,
		LastEditDateT:   x.LastEditDateT,
		PosSubs:         subs,
		CreditCards:     crds,
	}
}

func map_posslip_subs_template(x NewPosItemModel) pos.NewPosItemTemplate {
	return pos.NewPosItemTemplate{
		ItemCode:     x.ItemCode,
		ItemName:     x.ItemName,
		WHCode:       x.WHCode,
		ShelfCode:    x.ShelfCode,
		Qty:          x.Qty,
		Price:        x.Price,
		DiscountWord: x.DiscountWord,
		UnitCode:     x.UnitCode,
		BarCode:      x.BarCode,
		LineNumber:   x.LineNumber,
		AverageCost:  x.AverageCost,
		PackingRate1: x.PackingRate1,
	}
}
