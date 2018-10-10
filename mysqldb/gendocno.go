package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
	"github.com/mrtomyum/nopadol/gendocno"
)

type DocNoModel struct {
	BranchId     int64  `db:"BranchId"`
	TableCode    string `db:"TableCode"`
	BillType     int64  `db:"BillType"`
	Header       string `db:"Header"`
	UseYear      int64    `db:"UseYear"`
	UseMonth     int64    `db:"UseMonth"`
	UseDay       int64    `db:"UseDay"`
	UseDash      int64    `db:"UseDash"`
	FormatNumber int64    `db:"FormatNumber"`
	ActiveStatus int64    `db:"ActiveStatus"`
}

type gendocnoRepository struct{ db *sqlx.DB }

func NewGenDocNoRepository(db *sqlx.DB) gendocno.Repository {
	return &gendocnoRepository{db}
}

func (repo *gendocnoRepository) Gen(req *gendocno.DocNoTemplate) (resp string, err error) {
	var last_number1 int
	var last_number string
	var snumber string
	var intyear int
	var vHeader string
	var vyear string

	var intmonth int
	var intmonth1 int
	var vmonth string
	var vmonth1 string
	var lenmonth int
	var vbranch_head string

	fmt.Println("Table = ", req.TableCode)
	fmt.Println("BillType = ", req.BillType)

	d := DocNoModel{}

	sql := `select BranchId,TableCode,BillType,Header,UseYear,UseMonth,UseDay,UseDash,FormatNumber,ActiveStatus from GenDocMaster where TableCode =? and BillType = ? and BranchId = ?`
	err = repo.db.Get(&d, sql, req.TableCode, req.BillType, req.BranchId)
	if err != nil {
		return "", err
	}

	last_number1, _ = GetLastDocNo(repo.db, req.BranchId, req.TableCode, d.FormatNumber, req.BillType)
	last_number = strconv.Itoa(last_number1)
	fmt.Println("Last No = ", last_number)
	if time.Now().Year() >= 2560 {
		intyear = time.Now().Year()
	} else {
		intyear = time.Now().Year() + 543
	}

	vyear = strconv.Itoa(intyear)
	vyear1 := vyear[2:len(vyear)]

	fmt.Println("year = ", vyear1)

	intmonth = int(time.Now().Month())
	intmonth1 = int(intmonth)
	vmonth = strconv.Itoa(intmonth1)

	fmt.Println("month =", vmonth)

	lenmonth = len(vmonth)

	if lenmonth == 1 {
		vmonth1 = "0" + vmonth
	} else {
		vmonth1 = vmonth
	}

	if len(string(last_number)) == 1 {
		snumber = "000" + last_number
	}
	if len(string(last_number)) == 2 {
		snumber = "00" + last_number
	}
	if len(string(last_number)) == 3 {
		snumber = "0" + last_number
	}
	if len(string(last_number)) == 4 {
		snumber = last_number
	}

	vHeader = d.Header

	if d.BillType == 0 {
		vbranch_head = "S01"
	} else {
		vbranch_head = "W01"
	}

	NewDocNo := vbranch_head + "-" + d.Header + vyear1 + vmonth1 + "-" + snumber
	fmt.Println(snumber)
	fmt.Println(vHeader)

	fmt.Println("NewDocNo = ", NewDocNo)

	return NewDocNo, nil
}

func GetLastDocNo(db *sqlx.DB, branch_id int64, table_code string, formatnum int64, bill_type int64) (last_no int, err error) {
	var sql string
	var sqlcase string

	switch table_code {
	case "QT":
		sqlcase = `select cast(right(ifnull(max(docno),0),?) as int)+1 maxno from Quotation where BranchId = ? and BillType = ? and year(DocDate) = year(CURDATE()) and month(DocDate) = month(CURDATE())`
	case "BO":
		sqlcase = `select cast(right(ifnull(max(docno),0),?) as int)+1 maxno from Quotation where BranchId = ? and BillType = ? and year(DocDate) = year(CURDATE()) and month(DocDate) = month(CURDATE())`
	case "SO":
		sqlcase = `select cast(right(ifnull(max(docno),0),?) as int)+1 maxno from SaleOrder where BranchId = ? and DocType = 1 and BillType = ? and year(DocDate) = year(CURDATE()) and month(DocDate) = month(CURDATE())`
	case "RO":
		sqlcase = `select cast(right(ifnull(max(docno),0),?) as int)+1 maxno from SaleOrder where BranchId = ? and DocType = 0 and BillType = ? and year(DocDate) = year(CURDATE()) and month(DocDate) = month(CURDATE())`
	}

	sql = sqlcase
	fmt.Println("Branch ID =",branch_id)
	fmt.Println("Query = ", sql)
	err = db.Get(&last_no, sql, formatnum, branch_id, bill_type)
	if err != nil {
		fmt.Println(err)
		return 1, err
	}

	fmt.Println("Last No = ", last_no)
	return last_no, nil
}
