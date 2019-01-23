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

type ListQueueModel struct {
	Id                        int                     `json:"id"`
	UUID                      string                  `json:"uuid"`
	QueueId                   int                     `json:"queue_id"`
	NumberOfItem              int                     `json:"number_of_item"`
	TimeCreated               string                  `json:"time_created"`
	Status                    int                     `json:"status"`
	IsCancel                  int                     `json:"is_cancel"`
	ArCode                    string                  `json:"ar_code"`
	ArName                    string                  `json:"ar_name"`
	SaleName                  string                  `json:"sale_name"`
	SaleCode                  string                  `json:"sale_code"`
	DocNo                     string                  `json:"doc_no"`
	Source                    int                     `json:"source"`
	ReceiverName              string                  `json:"receiver_name"`
	DocDate                   string                  `json:"doc_date"`
	PickupDateTime            string                  `json:"pickup_date_time"`
	TotalAmount               float64                 `json:"total_amount"`
	IsLoaded                  int                     `json:"is_loaded"`
	CarBrand                  string                  `json:"car_brand"`
	PlateNumber               string                  `json:"plate_number"`
	StatusForSaleOrderCurrent int                     `json:"status_for_saleorder_current"`
	TotalBeforeAmount         float64                 `json:"total_before_amount"`
	TotalAfterAmount          float64                 `json:"total_after_amount"`
	OTPPassword               string                  `json:"otp_password"`
	BillType                  int                     `json:"bill_type"`
	CancelRemark              string                  `json:"cancel_remark"`
	WhoCancel                 string                  `json:"who_cancel"`
	SaleOrder                 string                  `json:"sale_order"`
	OwnerPhone                OwnerPhoneModel         `json:"owner_phone"`
	ReceiverPhone             OwnerPhoneModel         `json:"receiver_phone"`
	StatusForSaleorderHistory QueueStatusHistoryModel `json:"status_for_saleorder_history"`
	Item                      QueueItem               `json:"item"`
}

type OwnerPhoneModel struct {
	phone_no string `json:"phone_no"`
}

type QueueStatusHistoryModel struct {
	StatusId           int    `json:"status_id"`
	StatusForSaleOrder int    `json:"status_for_sale_order"`
	CreateDateTime     string `json:"create_date_time"`
}

type QueueItem struct {
	ItemBarCode      string  `json:"item_bar_code"`
	FilePath         string  `json:"file_path"`
	IsCancel         int     `json:"is_cancel"`
	ISCheck          int     `json:"is_check"`
	ItemCode         string  `json:"item_code"`
	ItemName         string  `json:"item_name"`
	PickupStaffName  string  `json:"pickup_staff_name"`
	SaleCode         string  `json:"sale_code"`
	ItemPrice        float64 `json:"item_price"`
	QtyAfter         float64 `json:"qty_after"`
	QtyBefore        float64 `json:"qty_before"`
	QtyLoad          float64 `json:"qty_load"`
	TotalPriceAfter  float64 `json:"total_price_after"`
	TotalPriceBefore float64 `json:"total_price_before"`
	ItemUnitCode     string  `json:"item_unit_code"`
	RequestQty       float64 `json:"request_qty"`
	ItemQty          float64 `json:"item_qty"`
	PickZoneId       string  `json:"pick_zone_id"`
	LineNumber       int     `json:"line_number"`
}

func (q *ListQueueModel) SearchQueueList(db *sqlx.DB, req *drivethru.ListQueueRequest) (interface{}, error) {
	que := []ListQueueModel{}

	lccommand := `select id, que_id as queue_id, car_brand, ref_number as plate_number,uuid, doc_date, number_of_item, create_time as time_created, status, is_cancel, '' as ar_code, '' as ar_name, '' as sale_name, '' as sale_code, doc_no, doc_type as source, '' as receiver_name, pickup_time as pickup_date_time, total_amount, 0 as is_loaded, 0 as status_for_saleorder_current, 0 as total_before_amount, 0 as total_after_amount, '' as otp_password, 0 as bill_type, '' as cancel_remark, '' as who_cancel, '' as sale_order from basket where doc_date = CURRENT_DATE `
	err := db.Select(&que, lccommand)
	if err != nil {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "pickup new",
				"processDesc": err.Error(),
				"isSuccess":   false,
			},
		}, nil
	}

	//for queue, _ := range que{
	//
	//}
	fmt.Println("q CarBrand = ", q.Id, q.QueueId, q.CarBrand, q.PlateNumber)
	return que, nil
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
	ar_id := 58672

	fmt.Println(qId, doc_type, doc_date, doc_no, user.UserCode, uuid)
	p.QueId = qId

	lccommand := `insert basket(company_id, branch_id, uuid, doc_no, que_id, doc_type, doc_date, car_brand, ref_number, ar_id, sale_id, create_by, create_time, pick_by, pickup_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	fmt.Println("insert basket =", lccommand)
	resp, err := db.Exec(lccommand, user.CompanyID, user.BranchID, uuid, doc_no, qId, doc_type, doc_date, req.CarBrand, req.CarNumber, ar_id, user.UserId, user.UserCode, now.String(), user.UserCode, now.String())
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

func (item *QueueItem) ManagePickup(db *sqlx.DB, req *drivethru.ManagePickupRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))

	if req.AccessToken == "" {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue Not Have Access Token",
			},
			"queid": ""}, nil
	}

	if req.QueueId == 0 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue Id Not Assign",
			},
			"queid": ""}, nil
	}

	if req.ItemBarcode == "" {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue Not Have Barcode",
			},
			"queid": ""}, nil
	}

	if req.QtyBefore == 0 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue Not Have Qty Pickup",
			},
			"queid": ""}, nil
	}

	q := ListQueueModel{}
	q.Search(db, req.QueueId)

	p := ProductModel{}
	p.SearchByBarcode(db, req.ItemBarcode)

	if p.ItemCode == "" {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "BarCode Not Have Data",
			},
			"queid": ""}, nil
	}

	u := UserAccess{}
	u.GetProfileByToken(db, req.AccessToken)

	s := EmployeeModel{}
	s.SearchBySaleCode(db, u.UserCode)

	fmt.Println("Car Number", q.PlateNumber)

	item_exist := QueCheckItemExist(db, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)

	fmt.Println(item_exist)
	if q.IsCancel == 0 {
		if q.Status < 2 {
			if item_exist == 0 {
				fmt.Println("Insert")
				lccommand := `insert basket_sub(basket_id, uuid, que_id, doc_date, item_id, item_code, item_name ,bar_code, request_qty, pick_qty, checkout_qty, price, unit_id, unit_code, pick_amount, checkout_amount, qty, remain_qty, rate1, ref_no, sale_id, average_cost, delivery_order_id , line_number, request_by, request_time, pick_by, pick_time) values(?, ?, ?, ?, ?, ?, ? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? , ?, ?, ?, ?, ?)`
				resp, err := db.Exec(lccommand, q.Id, q.UUID, req.QueueId, q.DocDate, p.Id, p.ItemCode, p.ItemName, req.ItemBarcode, req.QtyBefore, req.QtyBefore, 0, p.SalePrice1, 0, p.UnitCode, req.QtyBefore*p.SalePrice1, 0, req.QtyBefore, req.QtyBefore, p.Rate1, q.PlateNumber, s.Id, p.AverageCost, 0, req.LineNumber, u.UserCode, now.String(), u.UserCode, now.String())
				if err != nil {
					return map[string]interface{}{
						"response": map[string]interface{}{
							"success": false,
							"error":   true,
							"message": err.Error(),
						},
						"queid": ""}, nil
				}

				fmt.Println(resp.LastInsertId())
			} else {
				fmt.Println("Update")
				fmt.Println("UUID =", q.UUID, q.Id, req.QueueId, p.ItemCode, p.UnitCode, req.QtyBefore)
				if req.IsCancel == 0 {
					lccommand := `update basket_sub set request_qty=?, pick_qty=?, pick_amount=?, qty=?, remain_qty=? where basket_id = ? and uuid = ? and que_id = ? and item_code = ? and unit_code = ? and doc_date = CURDATE() `
					resp, err := db.Exec(lccommand, req.QtyBefore, req.QtyBefore, req.QtyBefore*p.SalePrice1, req.QtyBefore, req.QtyBefore, q.Id, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)
					if err != nil {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success": false,
								"error":   true,
								"message": err.Error(),
							},
							"queid": ""}, nil
					}
					fmt.Println(resp.LastInsertId())
				} else {
					lccommand := `update basket_sub set pick_qty=0, pick_amount=0, qty=0, remain_qty=0, is_cancel = 1, cancel_by = ?, cancel_time = ? where basket_id = ? and uuid = ? and que_id = ? and item_code = ? and unit_code = ? and doc_date = CURDATE() `
					resp, err := db.Exec(lccommand, u.UserCode, now.String(), q.Id, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)
					if err != nil {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success": false,
								"error":   true,
								"message": err.Error(),
							},
							"queid": ""}, nil
					}
					fmt.Println(resp.LastInsertId())
				}

			}

			lccommand := `update basket set number_of_item = (select count(*) as vcount from basket_sub where basket_id = ? and uuid = ? and que_id = ? and doc_date = CURDATE()),sum_item_amount = (select sum(pick_amount) as sumamount from basket_sub where basket_id = ? and uuid = ? and que_id = ? and doc_date = CURDATE() and is_cancel = 0) where id = ? and uuid = ? and que_id = ? and doc_date = CURDATE()`
			_, err := db.Exec(lccommand, q.Id, q.UUID, req.QueueId, q.Id, q.UUID, req.QueueId, q.Id, q.UUID, req.QueueId)
			if err != nil {
				fmt.Println(err.Error())
			}

			return map[string]interface{}{
				"response": map[string]interface{}{
					"success": true,
					"error":   true,
					"message": "",
				},
				"queid": map[string]interface{}{
					"item_barcode":       p.BarCode,
					"file_path":          p.PicPath1,
					"is_cancel":          0,
					"is_check":           0,
					"item_code":          p.ItemCode,
					"item_name":          p.ItemName,
					"pickup_staff_name":  s.SaleName,
					"sale_code":          s.SaleCode + "/" + s.SaleName,
					"item_price":         p.SalePrice1,
					"qty_after":          0,
					"qty_before":         req.QtyBefore,
					"qty_load":           0,
					"total_price_after":  0,
					"total_price_before": p.SalePrice1 * req.QtyBefore,
					"item_unit_code":     p.UnitCode,
					"request_qty":        0,
					"item_qty":           req.QtyBefore,
					"pick_zone_id":       "B",
					"line_number":        req.LineNumber,
				},
			}, nil
		} else {
			return map[string]interface{}{
				"response": map[string]interface{}{
					"success": false,
					"error":   true,
					"message": "Queue is ref used",
				},
				"queid": ""}, nil
		}

	} else {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue is cancel",
			},
			"queid": ""}, nil
	}

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

	var intday int
	var intday1 int
	var vday string
	var vday1 string
	var lenday int

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

	intday = int(time.Now().Day())
	intday1 = int(intday)
	vday = strconv.Itoa(intday1)

	fmt.Println("day =", vday)

	lenmonth = len(vmonth)

	if lenmonth == 1 {
		vmonth1 = "0" + vmonth
	} else {
		vmonth1 = vmonth
	}

	if lenday == 1 {
		vday1 = "0" + vday
	} else {
		vday1 = vday
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

	NewDocNo := vHeader + vyear1 + vmonth1 + vday1 + "-" + snumber
	fmt.Println(snumber)
	fmt.Println(vHeader)

	fmt.Println("NewDocNo = ", NewDocNo)

	return NewDocNo, nil
}

func getLastDocNo(db *sqlx.DB, company_id int, branch_id int, doc_type int) (last_no int, err error) {
	sql := `select cast(right(ifnull(max(doc_no),0),4) as int)+1 maxno from basket where company_id = ? and branch_id = ? and doc_type = ? and year(doc_date) = year(CURDATE()) and month(doc_date) = month(CURDATE()) and day(doc_date) = day(CURDATE())`
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

func (q *ListQueueModel) Search(db *sqlx.DB, queue_id int) {
	fmt.Println("q = ", queue_id)

	lccommand := `select id, que_id as queue_id, car_brand, ref_number as plate_number,uuid, doc_date, number_of_item, create_time as time_created, status, is_cancel, '' as ar_code, '' as ar_name, '' as sale_name, '' as sale_code, doc_no, doc_type as source, '' as receiver_name, pickup_time as pickup_date_time, total_amount, 0 as is_loaded, 0 as status_for_saleorder_current, 0 as total_before_amount, 0 as total_after_amount, '' as otp_password, 0 as bill_type, '' as cancel_remark, '' as who_cancel, '' as sale_order from basket where que_id=? and doc_date = CURRENT_DATE `
	rs := db.QueryRow(lccommand, queue_id)
	rs.Scan(&q.Id, &q.QueueId, &q.CarBrand, &q.PlateNumber, &q.UUID, &q.DocDate, &q.NumberOfItem, &q.TimeCreated, &q.Status, &q.IsCancel, &q.ArCode, &q.ArName, &q.SaleName, &q.SaleCode, &q.DocNo, &q.Source, &q.ReceiverName, &q.PickupDateTime, &q.TotalAmount, &q.IsLoaded, &q.StatusForSaleOrderCurrent, &q.TotalBeforeAmount, &q.TotalAfterAmount, &q.OTPPassword, &q.BillType, &q.CancelRemark, &q.WhoCancel, &q.SaleOrder)
	fmt.Println("q CarBrand = ", q.Id, q.QueueId, q.CarBrand, q.PlateNumber)
	return
}

func QueCheckItemExist(db *sqlx.DB, que_uuid string, que_id int, item_code string, unit_code string) int {
	var exist int
	lcCommand := `select count(*) as vCount from basket_sub where uuid = ? and que_id = ? and item_code = ? and unit_code = ? and doc_date = CURRENT_DATE`
	err := db.Get(&exist, lcCommand, que_uuid, que_id, item_code, unit_code)
	if err != nil {
		fmt.Println(err.Error())
	}

	return exist
}
