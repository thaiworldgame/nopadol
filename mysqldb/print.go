package mysqldb

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/knq/escpos"
	"github.com/mrtomyum/nopadol/cloudprint"
	"github.com/mrtomyum/nopadol/hw"
)

func printslippos(req *cloudprint.CloudPrintRequest) {
	fmt.Println("Docno = ", req.Data, req.FormType)

	host := req.PrinterIp + ":" + req.PrinterPort
	f, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(f)

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
	pt.WriteStringLines("NO: " + req.Data + "          " + "DATE: " + DocDate + "\n")
	pt.SetAlign("center")
	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	pt.SetAlign("left")
	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + req.FormType + "\n")
	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	pt.SetAlign("left")
	//pt.WriteStringLines("CS : " + req.SaleCode + "/" + req.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////
	pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////

	var CountItem int64
	var CountQty float64
	//for _, subcount := range req.Subs {
	//	CountItem = CountItem + 1
	//	CountQty = CountQty + 1//subcount.Qty
	//}

	fmt.Println("CountItem =", CountItem, CountQty)
	///////////////////////////////////////////////////////////////////////////////////
	pt.SetAlign("left")
	for _, sub := range req.Data {
		var vDiffEmpty int
		var vItemPriceAmount string
		var vItemAmount float64

		pt.SetFont("A")

		//pt.WriteStringLines(" " + sub.ItemName + "\n")

		//vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)

		vItemPriceAmount = " " + "sub.Code" //+ strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
		vLen := len(vItemPriceAmount)
		vDiff := 25 - (vLen / 3)

		if vDiff < 0 {
			vDiffEmpty = 0
		} else {
			vDiffEmpty = vDiff
		}

		fmt.Println("ItemName=", sub)
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
	pt.WriteStringLines("                          ")
	//pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	//if s.CoupongAmount != 0 {
	//	pt.WriteStringLines("COUPON: " + "                         " + CommaFloat(s.CoupongAmount) + "\n")
	//}
	//if s.SumCashAmount != 0 {
	//	pt.WriteStringLines("CASH: " + "                           " + CommaFloat(s.SumCashAmount) + "\n")
	//}
	//if s.SumCreditAmount != 0 {
	//	pt.WriteStringLines("CREDIT: " + "                         " + CommaFloat(s.SumCreditAmount) + "\n")
	//}
	//if s.ChangeAmount != 0 {
	//	pt.WriteStringLines("CHANGE: " + "                         " + CommaFloat(s.ChangeAmount) + "\n")
	//}

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
}

func cashbill(req *cloudprint.CloudPrintRequest) {
	fmt.Println("Docno = ", req.Data, req.FormType)

	host := req.PrinterIp + ":" + req.PrinterPort
	f, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(f)

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
	pt.WriteStringLines("CASH INV (ABB)")
	pt.LineFeed()
	pt.SetTextSize(0, 0)
	pt.SetAlign("left")
	pt.WriteStringLines("NO: " + req.Data + "          " + "DATE: " + DocDate + "\n")
	pt.SetAlign("center")
	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	pt.SetAlign("left")
	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + req.FormType + "\n")
	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	pt.SetAlign("left")
	//pt.WriteStringLines("CS : " + req.SaleCode + "/" + req.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////
	pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////

	var CountItem int64
	var CountQty float64
	//for _, subcount := range req.Subs {
	//	CountItem = CountItem + 1
	//	CountQty = CountQty + 1//subcount.Qty
	//}

	fmt.Println("CountItem =", CountItem, CountQty)
	///////////////////////////////////////////////////////////////////////////////////
	pt.SetAlign("left")
	for _, sub := range req.Data {
		var vDiffEmpty int
		var vItemPriceAmount string
		var vItemAmount float64

		pt.SetFont("A")

		//pt.WriteStringLines(" " + sub.ItemName + "\n")

		//vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)

		vItemPriceAmount = " " + "sub.Code" //+ strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
		vLen := len(vItemPriceAmount)
		vDiff := 25 - (vLen / 3)

		if vDiff < 0 {
			vDiffEmpty = 0
		} else {
			vDiffEmpty = vDiff
		}

		fmt.Println("ItemName=", sub)
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
	pt.WriteStringLines("                          ")
	//pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	//if s.CoupongAmount != 0 {
	//	pt.WriteStringLines("COUPON: " + "                         " + CommaFloat(s.CoupongAmount) + "\n")
	//}
	//if s.SumCashAmount != 0 {
	//	pt.WriteStringLines("CASH: " + "                           " + CommaFloat(s.SumCashAmount) + "\n")
	//}
	//if s.SumCreditAmount != 0 {
	//	pt.WriteStringLines("CREDIT: " + "                         " + CommaFloat(s.SumCreditAmount) + "\n")
	//}
	//if s.ChangeAmount != 0 {
	//	pt.WriteStringLines("CHANGE: " + "                         " + CommaFloat(s.ChangeAmount) + "\n")
	//}

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
}

func creditbill(req *cloudprint.CloudPrintRequest) {
	fmt.Println("Docno = ", req.Data, req.FormType)

	host := req.PrinterIp + ":" + req.PrinterPort
	f, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(f)

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
	pt.WriteStringLines("CRD INV (ABB)")
	pt.LineFeed()
	pt.SetTextSize(0, 0)
	pt.SetAlign("left")
	pt.WriteStringLines("NO: " + req.Data + "          " + "DATE: " + DocDate + "\n")
	pt.SetAlign("center")
	pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	pt.SetAlign("left")
	pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + req.FormType + "\n")
	//pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	pt.SetAlign("left")
	//pt.WriteStringLines("CS : " + req.SaleCode + "/" + req.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	//pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////
	pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	makeline(pt)
	///////////////////////////////////////////////////////////////////////////////////

	var CountItem int64
	var CountQty float64
	//for _, subcount := range req.Subs {
	//	CountItem = CountItem + 1
	//	CountQty = CountQty + 1//subcount.Qty
	//}

	fmt.Println("CountItem =", CountItem, CountQty)
	///////////////////////////////////////////////////////////////////////////////////
	pt.SetAlign("left")
	for _, sub := range req.Data {
		var vDiffEmpty int
		var vItemPriceAmount string
		var vItemAmount float64

		pt.SetFont("A")

		//pt.WriteStringLines(" " + sub.ItemName + "\n")

		//vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)

		vItemPriceAmount = " " + "sub.Code" //+ strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
		vLen := len(vItemPriceAmount)
		vDiff := 25 - (vLen / 3)

		if vDiff < 0 {
			vDiffEmpty = 0
		} else {
			vDiffEmpty = vDiff
		}

		fmt.Println("ItemName=", sub)
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
	pt.WriteStringLines("                          ")
	//pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
	////////////////////////////////////////////////////////////////////////////////////
	pt.SetFont("A")
	//pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	//pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	//if s.CoupongAmount != 0 {
	//	pt.WriteStringLines("COUPON: " + "                         " + CommaFloat(s.CoupongAmount) + "\n")
	//}
	//if s.SumCashAmount != 0 {
	//	pt.WriteStringLines("CASH: " + "                           " + CommaFloat(s.SumCashAmount) + "\n")
	//}
	//if s.SumCreditAmount != 0 {
	//	pt.WriteStringLines("CREDIT: " + "                         " + CommaFloat(s.SumCreditAmount) + "\n")
	//}
	//if s.ChangeAmount != 0 {
	//	pt.WriteStringLines("CHANGE: " + "                         " + CommaFloat(s.ChangeAmount) + "\n")
	//}

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
}

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
