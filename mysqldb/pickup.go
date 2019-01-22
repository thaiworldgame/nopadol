package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
	"fmt"
	"time"
	"strconv"
)

type pickupModel struct {
	QueUUID   string `json:"que_uuid"`
	QueId     int    `json:"que_id"`
	DocNo     string `json:"doc_no"`
	DocDate   string `json:"doc_date"`
	SaleId    int    `json:"sale_id"`
	ArId      int    `json:"ar_id"`
	CarBrand  string `json:"car_brand"`
	CarNumber string `json:"car_number"`
}

func (p *pickupModel) PickupNew(db *sqlx.DB, req *drivethru.NewPickupRequest) (interface{}, error) { //ขอโดมแก้หน้ากาก ตอนทำ Pickup
	user := UserAccess{}
	user.GetProfileByToken(db, req.AccessToken)

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	doc_date := now.AddDate(0, 0, 0).Format("2006-01-02")

	qId, err := getQueId(db, user.CompanyID, user.BranchID)
	if err != nil {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "pickup new",
				"processDesc": err.Error(),
				"isSuccess":   false,
			},
		}, nil
	}

	uuid := GetAccessToken()

	var doc_type int

	if (len(req.CarNumber) <= 4) {
		doc_type = 0
	} else {
		doc_type = 1
	}

	doc_no, err := getDocNo(db, user.CompanyID, user.BranchID, doc_type)
	ar_id := 0

	fmt.Println(qId, doc_type, doc_date, doc_no, user.UserCode, uuid)
	p.QueId = qId

	lccommand := `insert basket(company_id, branch_id, uuid, doc_no, que_id, doc_type, doc_date, car_brand, ref_number, ar_id, sale_id, create_by, create_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	fmt.Println("insert basket =", lccommand)
	resp, err := db.Exec(lccommand, user.CompanyID, user.BranchID, uuid, doc_no, qId, doc_type, doc_date, req.CarBrand, req.CarNumber, ar_id, user.UserId, user.UserCode, now.String())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp.LastInsertId())

	p.QueId = qId

	return map[string]interface{}{
		"response": map[string]interface{}{
			"process":     "pickup new",
			"processDesc": "successful",
			"isSuccess":   true,
		},
		"queid": p.QueId,
	}, nil
}

func getQueId(db *sqlx.DB, company_id int, branch_id int) (int, error) {
	var qId int

	lccommand := `select ifnull(MAX(que_id),0)+1 as qId from basket where  company_id = ? and branch_id =  ? and year(create_time) = year(CURRENT_DATE) and month(create_time) = month(CURRENT_DATE) and day(create_time) = day(CURRENT_DATE)`
	err := db.Get(&qId, lccommand, company_id, branch_id)
	if err != nil {
		fmt.Println("error gen qid =", err.Error())
		return 0, err
	}
	return qId, nil
}

func getDocNo(db *sqlx.DB, company_id int, branch_id int, doc_type int) (string, error) {
	var last_number1 int
	var last_number string
	var snumber string
	var intyear int
	var vHeader string
	var branch_header string
	var header string
	var vyear string

	var intmonth int
	var intmonth1 int
	var vmonth string
	var vmonth1 string
	var lenmonth int

	last_number1, _ = getLastDocNo(db, company_id, branch_id, doc_type)
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

	if branch_id == 1 {
		branch_header = "S01"
	} else {
		branch_header = "S02"
	}

	if doc_type == 0 {
		header = "QUE"
	} else {
		header = "BSK"
	}

	vHeader = branch_header + "-" + header

	NewDocNo := vHeader + vyear1 + vmonth1 + "-" + snumber
	fmt.Println(snumber)
	fmt.Println(vHeader)

	fmt.Println("NewDocNo = ", NewDocNo)

	return NewDocNo, nil
}

func getLastDocNo(db *sqlx.DB, company_id int, branch_id int, doc_type int) (last_no int, err error) {
	sql := `select cast(right(ifnull(max(doc_no),0),4) as int)+1 maxno from basket where company_id = ? and branch_id = ? and doc_type = ? and year(DocDate) = year(CURDATE()) and month(DocDate) = month(CURDATE())`
	fmt.Println("Branch ID =", branch_id)
	fmt.Println("Query = ", sql)
	err = db.Get(&last_no, sql, company_id, branch_id, doc_type)
	if err != nil {
		//fmt.Println("Last No Error = ",err)
		return 1, nil
	}

	fmt.Println("Last No = ", last_no)
	return last_no, nil
}
