package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/cloudprint"
	"fmt"
	"github.com/mrtomyum/nopadol/drivethru"
	"strconv"
	"encoding/json"
)

type cloudprintRepository struct{ db *sqlx.DB }

func NewCloudPrintRepository(db *sqlx.DB) cloudprint.Repository {
	return &cloudprintRepository{db}
}

type ListPrint struct {
	PrinterId   string `json:"printer_id" db:"printer_id"`
	PrinterName string `json:"printer_name" db:"printer_name"`
}

func (p *ListPrint) ListPrinter(db *sqlx.DB, access_token string) (resp interface{}, err error) {
	user := UserAccess{}
	user.GetProfileByToken(db, access_token)
	printer := []*ListPrint{}

	fmt.Println("Company = ", user.CompanyID)

	lccommand := `select name as printer_name,concat(ip_address,':',port) as printer_id from printer where active_status = 1 and company_id = ? and branch_id = ? order by ip_address`
	err = db.Select(&printer, lccommand, user.CompanyID, user.BranchID)
	if err != nil {
		return map[string]interface{}{
			"error":   true,
			"message": "List Printer Error = " + err.Error(),
			"success": false,
			"printer": nil,
		}, nil
	}

	return map[string]interface{}{
		"error":   false,
		"message": "",
		"success": true,
		"printer": printer,
	}, nil
}

func PrintSubmit(db *sqlx.DB, req *drivethru.PrintSubmitRequest) (resp interface{}, err error) {
	user := UserAccess{}
	user.GetProfileByToken(db, req.AccessToken)
	que_id, err := strconv.Atoi(req.QueueId)

	queue := ListQueueModel{}

	request := drivethru.QueueProductRequest{AccessToken: req.AccessToken, QueueId: que_id}

	q, err := queue.QueueData(db, &request)

	fmt.Println("queue = ", q)

	var str []byte

	if str, err = json.Marshal(q); err != nil {
		fmt.Println("error marshal : ", err)
	}

	str1 := string(str)

	fmt.Println("string json = ", str1)

	module := "INV"
	doc_type := "ArInvoice"

	lccommand := `insert into print_queue(company_id,branch_id,module,doc_type,form_type,printer_id,inserted_time,data) value(?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP ,?)`
	_, err = db.Exec(lccommand, user.CompanyID, user.BranchID, module, doc_type, req.FormType, req.PrinterId, str1)
	if err != nil {
		fmt.Println("error = ", err.Error())
	}

	return map[string]interface{}{
		"error":   false,
		"message": "",
		"success": true,
		"data": q,
	}, nil
}

func (repo *cloudprintRepository) CloudPrint(req *cloudprint.CloudPrintRequest) (resp interface{}, err error) {
	fmt.Println("Docno = ", req.FormType)

	switch req.FormType {
	case "short_form":
		printslippos(req)
	case "full_form":
		printslippos(req)
	case "cash_bill":
		cashbill(req)
	case "credit_bill":
		creditbill(req)
	}

	//host := "192.168.1.40:9100"
	//
	//f, err := net.Dial("tcp", host)
	//if err != nil {
	//	return nil, err
	//}
	//defer f.Close()
	//
	//w := bufio.NewWriter(f)
	//p := escpos.New(f)
	//
	//pt := hw.PosPrinter{p, w}
	//pt.Init()
	//pt.SetLeftMargin(20)
	//
	//loc, _ := time.LoadLocation("Asia/Bangkok")
	//now := time.Now().In(loc)
	//
	////now := time.Now()
	////fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	////DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")
	//DocDate := now.Format("02-01-2006 ")
	//
	////////////////////////////////////////////////////////////////////////////////////////
	//pt.SetCharaterCode(26)
	//pt.SetAlign("center")
	//pt.SetTextSize(0, 0)
	//pt.WriteStringLines("TAX INV (ABB)")
	//pt.LineFeed()
	//pt.SetTextSize(0, 0)
	//pt.SetAlign("left")
	//pt.WriteStringLines("NO: " + req.DocNo + "          " + "DATE: " + DocDate + "\n")
	//pt.SetAlign("center")
	//pt.WriteStringLines("บริษัท นพดลพานิช จำกัด" + "\n")
	//pt.SetAlign("left")
	//pt.WriteStringLines("TAX ID: 0505533999157" + "      " + "POS: " + req.FormType + "\n")
	////pt.WriteStringLines("ใบกำกับภาษีอย่างย่อ\n")
	//pt.SetAlign("left")
	////pt.WriteStringLines("CS : " + req.SaleCode + "/" + req.SaleName + "      " + "TIME:" + now.Format("15:04:05") + "\n")
	////pt.WriteStringLines(" พนักงาน : "+s.CreateBy+"\n")
	//makeline(pt)
	/////////////////////////////////////////////////////////////////////////////////////
	//pt.WriteStringLines(" รายการ " + "                        " + "มูลค่า" + "\n")
	//makeline(pt)
	/////////////////////////////////////////////////////////////////////////////////////
	//
	//var CountItem int64
	//var CountQty float64
	////for _, subcount := range req.Subs {
	////	CountItem = CountItem + 1
	////	CountQty = CountQty + 1//subcount.Qty
	////}
	//
	//fmt.Println("CountItem =", CountItem, CountQty)
	/////////////////////////////////////////////////////////////////////////////////////
	//pt.SetAlign("left")
	////for _, sub := range req.Subs {
	////	var vDiffEmpty int
	////	var vItemPriceAmount string
	////	var vItemAmount float64
	////
	////	pt.SetFont("A")
	////
	////	//pt.WriteStringLines(" " + sub.ItemName + "\n")
	////
	////	//vItemAmount = sub.Qty * (sub.Price - sub.DiscountAmount)
	////
	////	vItemPriceAmount = " " //+ strconv.Itoa(int(sub.Qty)) + " " + sub.UnitCode + " X " + strconv.FormatFloat(sub.Price, 'f', -1, 64)
	////	vLen := len(vItemPriceAmount)
	////	vDiff := 25 - (vLen / 3)
	////
	////	if vDiff < 0 {
	////		vDiffEmpty = 0
	////	} else {
	////		vDiffEmpty = vDiff
	////	}
	////
	////	//fmt.Println("ItemName=", sub.ItemName)
	////	fmt.Println("Len", vLen/3)
	////	fmt.Println("Diff ", vDiff)
	////	pt.WriteStringLines(vItemPriceAmount + strings.Repeat(" ", vDiffEmpty))
	////
	////	pt.WriteStringLines("   ")
	////	pt.WriteStringLines(CommaFloat(vItemAmount) + "\n\n")
	////	//pt.FormfeedN(3)
	////}
	//makeline(pt)
	//////////////////////////////////////////////////////////////////////////////////////
	//pt.SetFont("A")
	//pt.WriteStringLines(" " + strconv.Itoa(int(CountItem)) + " รายการ " + strconv.Itoa(int(CountQty)) + " ชิ้น\n")
	//pt.WriteStringLines("TOTAL: ")
	//pt.WriteStringLines("                          ")
	////pt.WriteStringLines(CommaFloat(s.NetDebtAmount) + "\n")
	//////////////////////////////////////////////////////////////////////////////////////
	//pt.SetFont("A")
	////pt.WriteStringLines(" มูลค่าสินค้ามีภาษีมูลค่าเพิ่ม"+"                       "+Commaf(vBeforeTaxAmount)+"\n")
	////pt.WriteStringLines(" ภาษีมูลค่าเพิ่ม"+strconv.Itoa(c.TaxRate)+"%"+"                                "+Commaf(vTaxAmount)+"\n")
	////if s.CoupongAmount != 0 {
	////	pt.WriteStringLines("COUPON: " + "                         " + CommaFloat(s.CoupongAmount) + "\n")
	////}
	////if s.SumCashAmount != 0 {
	////	pt.WriteStringLines("CASH: " + "                           " + CommaFloat(s.SumCashAmount) + "\n")
	////}
	////if s.SumCreditAmount != 0 {
	////	pt.WriteStringLines("CREDIT: " + "                         " + CommaFloat(s.SumCreditAmount) + "\n")
	////}
	////if s.ChangeAmount != 0 {
	////	pt.WriteStringLines("CHANGE: " + "                         " + CommaFloat(s.ChangeAmount) + "\n")
	////}
	//
	//pt.SetFont("A")
	//pt.SetAlign("center")
	//pt.WriteStringLines("ขอบคุณที่มาใช้บริการ" + "\n")
	//pt.WriteStringLines("เปลี่ยนสินค้าภายใน 30 วัน (รหัส [P] ไม่รับคืน)" + "\n")
	//pt.WriteStringLines("สอบถามโทร. 0-5324-0377 8.00-17.00 ทุกวัน" + "\n")
	//makeline(pt)
	//
	//pt.Formfeed()
	//pt.Cut()
	//pt.OpenCashBox()
	//pt.End()

	return req.Data, nil
}
