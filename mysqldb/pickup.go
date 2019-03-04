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
	Id                        int                       `json:"id" db:"id"`
	UUID                      string                    `json:"uuid" db:"uuid"`
	QueueId                   int                       `json:"queue_id" db:"queue_id"`
	NumberOfItem              int                       `json:"number_of_item" db:"number_of_item"`
	TimeCreated               string                    `json:"time_created" db:"time_created"`
	Status                    int                       `json:"status" db:"status"`
	IsCancel                  int                       `json:"is_cancel" db:"is_cancel"`
	ArCode                    string                    `json:"ar_code" db:"ar_code"`
	ArName                    string                    `json:"ar_name" db:"ar_name"`
	SaleName                  string                    `json:"sale_name" db:"sale_name"`
	SaleCode                  string                    `json:"sale_code" db:"sale_code"`
	DocNo                     string                    `json:"doc_no" db:"doc_no"`
	Source                    int                       `json:"source" db:"source"`
	ReceiverName              string                    `json:"receiver_name" db:"receiver_name"`
	DocDate                   string                    `json:"doc_date" db:"doc_date"`
	PickupDateTime            string                    `json:"pickup_datetime" db:"pickup_datetime"`
	TotalAmount               float64                   `json:"total_amount" db:"total_amount"`
	IsLoaded                  int                       `json:"is_loaded" db:"is_loaded"`
	CarBrand                  string                    `json:"car_brand" db:"car_brand"`
	PlateNumber               string                    `json:"plate_number" db:"plate_number"`
	StatusForSaleOrderCurrent int                       `json:"status_for_saleorder_current" db:"status_for_saleorder_current"`
	TotalBeforeAmount         float64                   `json:"total_before_amount" db:"total_before_amount"`
	TotalAfterAmount          float64                   `json:"total_after_amount" db:"total_after_amount"`
	OTPPassword               string                    `json:"otp_password" db:"otp_password"`
	BillType                  int                       `json:"bill_type" db:"bill_type"`
	CancelRemark              string                    `json:"cancel_remark" db:"cancel_remark"`
	WhoCancel                 string                    `json:"who_cancel" db:"who_cancel"`
	SaleOrder                 string                    `json:"sale_order" db:"sale_order"`
	OwnerPhone                []OwnerPhoneModel         `json:"owner_phone" db:"owner_phone"`
	ReceiverPhone             []OwnerPhoneModel         `json:"receiver_phone" db:"receiver_phone"`
	StatusForSaleorderHistory []QueueStatusHistoryModel `json:"status_for_saleorder_history" db:"status_for_saleorder_history"`
	Item                      []QueueItem               `json:"item" db:"item"`
}

type OwnerPhoneModel struct {
	Phone_no string `json:"phone_no" db:"phone_no"`
}

type QueueStatusHistoryModel struct {
	StatusId           int    `json:"status_id"`
	StatusForSaleOrder int    `json:"status_for_sale_order"`
	CreateDateTime     string `json:"create_date_time"`
}

type QueueItem struct {
	Id               int     `json:"id" db:"id"`
	ItemBarCode      string  `json:"item_barcode" db:"item_bar_code"`
	FilePath         string  `json:"file_path" db:"file_path"`
	IsCancel         int     `json:"is_cancel" db:"is_cancel"`
	IsCheck          int     `json:"is_check" db:"is_check"`
	ItemId           int     `json:"item_id" db:"item_id"`
	ItemCode         string  `json:"item_code" db:"item_code"`
	ItemName         string  `json:"item_name" db:"item_name"`
	PickupStaffName  string  `json:"pickup_staff_name" db:"pickup_staff_name"`
	SaleCode         string  `json:"sale_code" db:"sale_code"`
	ItemPrice        float64 `json:"item_price" db:"item_price"`
	QtyAfter         float64 `json:"qty_after" db:"qty_after"`
	QtyBefore        float64 `json:"qty_before" db:"qty_before"`
	QtyLoad          float64 `json:"qty_load" db:"qty_load"`
	AverageCost      float64 `json:"average_cost" db:"average_cost"`
	Rate1            int     `json:"rate_1" db:"rate1"`
	TotalPriceAfter  float64 `json:"total_price_after" db:"total_price_after"`
	TotalPriceBefore float64 `json:"total_price_before" db:"total_price_before"`
	ItemUnitCode     string  `json:"item_unit_code" db:"item_unit_code"`
	RequestQty       float64 `json:"request_qty" db:"request_qty"`
	ItemQty          float64 `json:"item_qty" db:"item_qty"`
	PickZoneId       string  `json:"pick_zone_id" db:"pick_zone_id"`
	LineNumber       int     `json:"line_number" db:"line_number"`
}

func (q *ListQueueModel) SearchQueueList(db *sqlx.DB, req *drivethru.ListQueueRequest) (interface{}, error) {
	que := []ListQueueModel{}
	que_data := []ListQueueModel{}

	//lccommand := `select id, que_id as queue_id, car_brand, ref_number as plate_number,uuid, doc_date, number_of_item, create_time as time_created, status, is_cancel, '' as ar_code, '' as ar_name, '' as sale_name, '' as sale_code, doc_no, doc_type as source, '' as receiver_name, pickup_time as pickup_datetime, total_amount, 0 as is_loaded, 0 as status_for_saleorder_current, ifnull(sum_item_amount,0) as total_before_amount, ifnull(total_amount,0) as total_after_amount, '' as otp_password, 0 as bill_type, '' as cancel_remark, '' as who_cancel, '' as sale_order from basket where doc_date = CURRENT_DATE order by id`
	lccommand := `call USP_DT_SearchListQue(?, ?, ?, ?, ?, ?)`
	err := db.Select(&que, lccommand, req.PickDate, req.CreateDate, req.Status, req.Page, req.Keyword, req.QueId)
	//err := db.Select(&que, lccommand)
	if err != nil {
		return map[string]interface{}{
			"error":   true,
			"message": "Queue List Doc Error = " + err.Error(),
			"success": false,
			"order":   nil,
		}, nil
	}

	for _, qid := range que {

		fmt.Println("que item = ", qid.Id, qid.QueueId, qid.UUID)

		lccommand := `select id, item_id, item_code, item_name ,bar_code as item_bar_code, request_qty, pick_qty as qty_before, checkout_qty as qty_after, price as item_price, unit_code as item_unit_code, pick_amount as total_price_before, checkout_amount as total_price_after, rate1, '' as sale_code, average_cost, line_number, '' as pick_zone_id from basket_sub where basket_id = ? and que_id = ? and uuid = ? and doc_date = CURDATE() order by line_number`
		err := db.Select(&qid.Item, lccommand, qid.Id, qid.QueueId, qid.UUID)
		if err != nil {
			return map[string]interface{}{
				"error":   true,
				"message": "Queue List item Error = " + err.Error(),
				"success": false,
				"order":   nil,
			}, nil
		}

		//fmt.Println("ItemCode = ", qid.Item[0].ItemCode)

		lccommand1 := `select phone_no from owner_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
		err = db.Select(&qid.OwnerPhone, lccommand1, qid.Id, qid.QueueId, qid.UUID, qid.DocNo)
		if err != nil {
			return map[string]interface{}{
				"error":   true,
				"message": "Queue List phone Error = " + err.Error(),
				"success": false,
				"order":   nil,
			}, nil
		}

		lccommand2 := `select phone_no from order_trust_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
		err = db.Select(&qid.ReceiverPhone, lccommand2, qid.Id, qid.QueueId, qid.UUID, qid.DocNo)
		if err != nil {
			return map[string]interface{}{
				"error":   true,
				"message": "Queue List phone Error = " + err.Error(),
				"success": false,
				"order":   nil,
			}, nil
		}

		que_data = append(que_data, qid)
	}

	return map[string]interface{}{
		"error":   false,
		"message": "",
		"success": true,
		"order":   que_data,
	}, nil
}

func (q *ListQueueModel) QueueProduct(db *sqlx.DB, req *drivethru.QueueProductRequest) (interface{}, error) {
	que := ListQueueModel{}

	fmt.Println("Q", req.QueueId)
	lccommand := `select a.id, que_id as queue_id, car_brand, ref_number as plate_number,uuid, doc_date, number_of_item, a.create_time as time_created, status, a.is_cancel, ifnull(b.code,'') as ar_code, ifnull(b.name,'') as ar_name, ifnull(c.SaleName,'') as sale_name, ifnull(c.SaleCode,'') as sale_code, doc_no, doc_type as source, '' as receiver_name, pickup_time as pickup_datetime, total_amount, 0 as is_loaded, 0 as status_for_saleorder_current, ifnull(sum_item_amount,0) as total_before_amount, ifnull(total_amount,0) as total_after_amount, '' as otp_password, 0 as bill_type, '' as cancel_remark, '' as who_cancel, '' as sale_order from basket a left join Customer b on a.ar_id = b.id left join Sale c on a.sale_id = c.id where que_id = ? and doc_date = CURDATE() order by id`
	err := db.Get(&que, lccommand, req.QueueId)
	if err != nil {
		return map[string]interface{}{
				"error":   true,
				"message": "Queue List Doc Error = " + err.Error(),
				"success": false,
				"queue":   nil,
		}, nil
	}

	lccommand1 := `select id, item_id, item_code, item_name ,bar_code as item_bar_code, request_qty, pick_qty as qty_before, checkout_qty as qty_after, price as item_price, unit_code as item_unit_code, pick_amount as total_price_before, checkout_amount as total_price_after, rate1, '' as sale_code, average_cost, line_number, '' as pick_zone_id from basket_sub where basket_id = ? and que_id = ? and uuid = ? and doc_date = CURDATE() order by line_number`
	err = db.Select(&que.Item, lccommand1, que.Id, que.QueueId, que.UUID)
	if err != nil {
		return map[string]interface{}{
				"error":   true,
				"message": "Queue List item Error = " + err.Error(),
				"success": false,
				"queue":   nil,
		}, nil
	}

	lccommand2 := `select phone_no from owner_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
	err = db.Select(&que.OwnerPhone, lccommand2, que.Id, que.QueueId, que.UUID, que.DocNo)
	if err != nil {
		return map[string]interface{}{
				"error":   true,
				"message": "Queue List phone Error = " + err.Error(),
				"success": false,
				"queue":   nil,
		}, nil
	}

	lccommand3 := `select phone_no from order_trust_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
	err = db.Select(&que.ReceiverPhone, lccommand3, que.Id, que.QueueId, que.UUID, que.DocNo)
	if err != nil {
		return map[string]interface{}{
				"error":   true,
				"message": "Queue List phone Error = " + err.Error(),
				"success": false,
				"queue":   nil,
		}, nil
	}

	if que.Item == nil {
		que.Item = []QueueItem{}
	}

	if que.OwnerPhone == nil {
		que.OwnerPhone = []OwnerPhoneModel{}
	}

	if que.ReceiverPhone == nil {
		que.ReceiverPhone = []OwnerPhoneModel{}
	}

	if que.StatusForSaleorderHistory == nil {
		que.StatusForSaleorderHistory = []QueueStatusHistoryModel{}
	}

	return map[string]interface{}{
			"error":   false,
			"message": "",
			"success": true,
			"queue":   que,
	}, nil
}

func (q *ListQueueModel) QueueDetails(db *sqlx.DB, que_id int, access_token string) (interface{}, error) {
	if que_id == 0 {
		return map[string]interface{}{
				"process":     "queue list",
				"processDesc": "Queue Id = 0",
				"isSuccess":   false,
		}, nil
	}

	lccommand := `select id, que_id as queue_id, car_brand, ref_number as plate_number,uuid, doc_date, number_of_item, create_time as time_created, status, is_cancel, '' as ar_code, '' as ar_name, '' as sale_name, '' as sale_code, doc_no, doc_type as source, '' as receiver_name, pickup_time as pickup_datetime, total_amount, 0 as is_loaded, 0 as status_for_saleorder_current, ifnull(sum_item_amount,0) as total_before_amount, ifnull(total_amount,0) as total_after_amount, '' as otp_password, 0 as bill_type, '' as cancel_remark, '' as who_cancel, '' as sale_order from basket where que_id = ? and doc_date = CURRENT_DATE order by id`
	err := db.Get(&q, lccommand, que_id)
	if err != nil {
		return map[string]interface{}{
				"process":     "queue list",
				"processDesc": "Queue List Doc Error = " + err.Error(),
				"isSuccess":   false,
		}, nil
	}

	lccommand1 := `select id, item_id, item_code, item_name ,bar_code as item_bar_code, request_qty, pick_qty as qty_before, checkout_qty as qty_after, price as item_price, unit_code as item_unit_code, pick_amount as total_price_before, checkout_amount as total_price_after, rate1, '' as sale_code, average_cost, line_number, '' as pick_zone_id from basket_sub where basket_id = ? and que_id = ? and uuid = ? and doc_date = CURDATE() order by line_number`
	err = db.Select(&q.Item, lccommand1, q.Id, q.QueueId, q.UUID)
	if err != nil {
		return map[string]interface{}{
				"process":     "queue list item",
				"processDesc": "Queue List item Error = " + err.Error(),
				"isSuccess":   false,
		}, nil
	}

	lccommand2 := `select phone_no from owner_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
	err = db.Select(&q.OwnerPhone, lccommand2, q.Id, q.QueueId, q.UUID, q.DocNo)
	if err != nil {
		return map[string]interface{}{
				"process":     "queue list phone",
				"processDesc": "Queue List phone Error = " + err.Error(),
				"isSuccess":   false,
		}, nil
	}

	lccommand3 := `select phone_no from order_trust_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
	err = db.Select(&q.ReceiverPhone, lccommand3, q.Id, q.QueueId, q.UUID, q.DocNo)
	if err != nil {
		return map[string]interface{}{
				"process":     "queue list phone",
				"processDesc": "Queue List phone Error = " + err.Error(),
				"isSuccess":   false,
		}, nil
	}

	return q, nil
}

func (p *pickupModel) PickupNew(db *sqlx.DB, req *drivethru.NewPickupRequest) (interface{}, error) { //ขอโดมแก้หน้ากาก ตอนทำ Pickup
	user := UserAccess{}
	user.GetProfileByToken(db, req.AccessToken)

	fmt.Println("Company Branch = ", user.CompanyID, user.BranchID)

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	doc_date := now.AddDate(0, 0, 0).Format("2006-01-02")

	qId, err := getQueId(db, user.CompanyID, user.BranchID)
	if err != nil {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "newPickup",
				"processDesc": err.Error(),
				"isSuccess":   false,
			}, "qId": nil,
		}, nil
	}

	if qId == 0 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "newPickup",
				"processDesc": "Queue not gen Id",
				"isSuccess":   false,
			}, "qId": nil,
		}, nil
	}

	var doc_type int

	doc_type, _ = strconv.Atoi(req.DocType)

	if doc_type == 1 {
		req.CarNumber = user.UserCode
		req.CarBrand = "Basket"
	}

	if req.CarNumber == "" {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "newPickup",
				"processDesc": "Queue not have car number",
				"isSuccess":   false,
			}, "qId": nil,
		}, nil
	}

	if req.CarBrand == "" {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "newPickup",
				"processDesc": "Queue not have car brand",
				"isSuccess":   false,
			}, "qId": nil,
		}, nil
	}

	if req.AccessToken == "" {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "newPickup",
				"processDesc": "Queue not have access_token",
				"isSuccess":   false,
			}, "qId": nil,
		}, nil
	}

	uuid := GetAccessToken()
	config := RequestConfigModel{}

	config.Search(db, user.CompanyID, user.BranchID)

	doc_no, err := getBasketNo(db, user.CompanyID, user.BranchID, doc_type)
	ar_id := config.DefCustId

	fmt.Println(qId, doc_type, doc_date, doc_no, user.UserCode, uuid)
	p.QueId = qId

	lccommand := `insert basket(company_id, branch_id, uuid, doc_no, que_id, doc_type, doc_date, car_brand, ref_number, ar_id, sale_id, create_by, create_time, pick_by, pickup_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	fmt.Println("insert basket =", lccommand)
	resp, err := db.Exec(lccommand, user.CompanyID, user.BranchID, uuid, doc_no, qId, doc_type, doc_date, req.CarBrand, req.CarNumber, ar_id, user.UserId, user.UserCode, now.String(), user.UserCode, now.String())
	if err != nil {
		fmt.Println(err.Error())
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "newPickup",
				"processDesc": "err = " + err.Error(),
				"isSuccess":   false,
			}, "qId": nil,
		}, nil
	}

	fmt.Println(resp.LastInsertId())

	p.QueId = qId

	return map[string]interface{}{
		"response": map[string]interface{}{
			"process":     "newPickup",
			"processDesc": "Successful",
			"isSuccess":   true,
		}, "qId": p.QueId,
	}, nil
}

func (item *QueueItem) ManagePickup(db *sqlx.DB, req *drivethru.ManagePickupRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))

	if req.AccessToken == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Not Have Access Token",
			"item":    nil,
		}, nil
	}

	if req.QueueId == 0 {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Id Not Assign",
			"item":    nil,
		}, nil
	}

	if req.ItemBarcode == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Not Have Barcode",
			"item":    nil,
		}, nil
	}

	//if req.QtyBefore == 0 {
	//	return map[string]interface{}{
	//		"response": map[string]interface{}{
	//			"success": false,
	//			"error":   true,
	//			"message": "Queue Not Have Qty Pickup",
	//		},
	//		"queid": ""}, nil
	//}

	q := ListQueueModel{}
	q.Search(db, req.QueueId)

	fmt.Println("uuid =", q.UUID)

	p := ProductModel{}
	p.SearchByBarcode(db, req.ItemBarcode)

	if p.ItemCode == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "BarCode Not Have Data",
			"item":    nil,
		}, nil
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
						"success": false,
						"error":   true,
						"message": err.Error(),
						"item":    nil,
					}, nil
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
							"success": false,
							"error":   true,
							"message": err.Error(),
							"item":    nil,
						}, nil
					}
					fmt.Println(resp.LastInsertId())
				} else {
					lccommand := `update basket_sub set pick_qty=0, pick_amount=0, qty=0, remain_qty=0, is_cancel = 1, cancel_by = ?, cancel_time = ? where basket_id = ? and uuid = ? and que_id = ? and item_code = ? and unit_code = ? and doc_date = CURDATE() `
					resp, err := db.Exec(lccommand, u.UserCode, now.String(), q.Id, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)
					if err != nil {
						return map[string]interface{}{
							"success": false,
							"error":   true,
							"message": err.Error(),
							"item":    nil,
						}, nil
					}
					fmt.Println(resp.LastInsertId())
				}

			}

			lccommand := `update basket set number_of_item = (select count(*) as vcount from basket_sub where basket_id = ? and uuid = ? and que_id = ? and doc_date = CURDATE()),sum_item_amount = (select sum(pick_amount) as sumamount from basket_sub where basket_id = ? and uuid = ? and que_id = ? and doc_date = CURDATE() and is_cancel = 0) where id = ? and uuid = ? and que_id = ? and doc_date = CURDATE()`
			_, err := db.Exec(lccommand, q.Id, q.UUID, req.QueueId, q.Id, q.UUID, req.QueueId, q.Id, q.UUID, req.QueueId)
			if err != nil {
				return map[string]interface{}{
					"success": false,
					"error":   true,
					"message": err.Error(),
					"item":    nil,
				}, nil
			}

			item.SearchQueueItem(db, req.QueueId, item.ItemCode, item.ItemUnitCode, req.LineNumber)

			return map[string]interface{}{
				"success": true,
				"error":   true,
				"message": "",
				"item": map[string]interface{}{
					"item_barcode":       p.BarCode,
					"file_path":          p.PicPath1,
					"is_cancel":          item.IsCancel,
					"is_check":           item.IsCheck,
					"item_code":          p.ItemCode,
					"item_name":          p.ItemName,
					"pickup_staff_name":  s.SaleName,
					"sale_code":          s.SaleCode + "/" + s.SaleName,
					"item_price":         p.SalePrice1,
					"qty_after":          req.QtyBefore,
					"qty_before":         item.QtyBefore,
					"qty_load":           item.QtyAfter,
					"total_price_after":  item.TotalPriceAfter,
					"total_price_before": p.SalePrice1 * req.QtyBefore,
					"item_unit_code":     p.UnitCode,
					"request_qty":        item.RequestQty,
					"item_qty":           req.QtyBefore,
					"pick_zone_id":       item.PickZoneId,
					"line_number":        req.LineNumber,
				},
			}, nil
		} else {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue is ref used",
				"item":    nil,
			}, nil
		}

	} else {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue is cancel",
			"item":    nil,
		}, nil
	}

}

func (item *QueueItem) ManageCheckOut(db *sqlx.DB, req *drivethru.ManageCheckoutRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))

	u := UserAccess{}
	u.GetProfileByToken(db, req.AccessToken)

	s := EmployeeModel{}
	s.SearchBySaleCode(db, u.UserCode)

	if req.AccessToken == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Not Have Access Token",
			"item":    nil,
		}, nil
	}

	if req.QueueId == 0 {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Id Not Assign",
			"item":    nil,
		}, nil
	}

	if req.ItemBarcode == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Not Have Barcode",
			"item":    nil,
		}, nil
	}

	//if req.QtyAfter == 0 {
	//	return map[string]interface{}{
	//		"response": map[string]interface{}{
	//			"success": false,
	//			"error":   true,
	//			"message": "Queue Not Have Qty CheckOut",
	//		},
	//		"queid": ""}, nil
	//}

	q := ListQueueModel{}
	q.Search(db, req.QueueId)

	p := ProductModel{}
	p.SearchByBarcode(db, req.ItemBarcode)

	if p.ItemCode == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "BarCode Not Have Data",
			"item":    ""}, nil
	}

	fmt.Println("Car Number", q.PlateNumber)

	item_exist := QueCheckItemExist(db, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)
	fmt.Println("QueCheckItemExist", q.UUID, req.QueueId, p.ItemCode, p.UnitCode)

	fmt.Println(item_exist)
	if q.IsCancel == 0 {
		if q.Status < 2 {
			if item_exist == 0 {
				fmt.Println("Insert")
				lccommand := `insert basket_sub(basket_id, uuid, que_id, doc_date, item_id, item_code, item_name ,bar_code, checkout_qty, price, unit_id, unit_code, checkout_amount, qty, rate1, ref_no, sale_id, average_cost, delivery_order_id , line_number, checkout_by, checkout_time) values(?, ?, ?, ?, ?, ?, ? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? , ?, ?, ?)`
				resp, err := db.Exec(lccommand, q.Id, q.UUID, req.QueueId, q.DocDate, p.Id, p.ItemCode, p.ItemName, req.ItemBarcode, req.QtyAfter, p.SalePrice1, 0, p.UnitCode, req.QtyAfter*p.SalePrice1, req.QtyAfter, p.Rate1, q.PlateNumber, s.Id, p.AverageCost, 0, req.LineNumber, u.UserCode, now.String())
				if err != nil {
					return map[string]interface{}{
						"success": false,
						"error":   true,
						"message": err.Error(),
						"item":    nil,
					}, nil
				}

				fmt.Println(resp.LastInsertId())
			} else {
				fmt.Println("Update")
				fmt.Println("UUID =", q.UUID, q.Id, req.QueueId, p.ItemCode, p.UnitCode, req.QtyAfter)
				if req.IsCancel == 0 {
					lccommand := `update basket_sub set checkout_qty=?, checkout_amount=?, qty=?, remain_qty=pick_qty - ? where basket_id = ? and uuid = ? and que_id = ? and item_code = ? and unit_code = ? and doc_date = CURDATE() `
					resp, err := db.Exec(lccommand, req.QtyAfter, req.QtyAfter*p.SalePrice1, req.QtyAfter, req.QtyAfter, q.Id, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)
					if err != nil {
						return map[string]interface{}{
							"success": false,
							"error":   true,
							"message": err.Error(),
							"item":    nil,
						}, nil
					}
					fmt.Println(resp.LastInsertId())
				} else {
					lccommand := `update basket_sub set pick_qty=0, pick_amount=0, qty=0, remain_qty=0,checkout_qty = 0, checkout_amount = 0, is_cancel = 1, cancel_by = ?, cancel_time = ? where basket_id = ? and uuid = ? and que_id = ? and item_code = ? and unit_code = ? and doc_date = CURDATE() `
					resp, err := db.Exec(lccommand, u.UserCode, now.String(), q.Id, q.UUID, req.QueueId, p.ItemCode, p.UnitCode)
					if err != nil {
						return map[string]interface{}{
							"success": false,
							"error":   true,
							"message": err.Error(),
							"item":    nil,
						}, nil
					}
					fmt.Println(resp.LastInsertId())
				}
			}

			lccommand := `update basket set status = 1, pick_status=1, is_check_out = 1, number_of_item = (select count(*) as vcount from basket_sub where basket_id = ? and uuid = ? and que_id = ? and doc_date = CURDATE()),total_amount = (select sum(checkout_amount) as sumamount from basket_sub where basket_id = ? and uuid = ? and que_id = ? and doc_date = CURDATE() and is_cancel = 0), checkout_by = ?, checkout_time = ? where id = ? and uuid = ? and que_id = ? and doc_date = CURDATE()`
			_, err := db.Exec(lccommand, q.Id, q.UUID, req.QueueId, q.Id, q.UUID, req.QueueId, u.UserId, now.String(), q.Id, q.UUID, req.QueueId)
			if err != nil {
				return map[string]interface{}{
					"success": false,
					"error":   true,
					"message": err.Error(),
					"item":    nil,
				}, nil
			}

			item.SearchQueueItem(db, req.QueueId, item.ItemCode, item.ItemUnitCode, req.LineNumber)

			return map[string]interface{}{
				"success": true,
				"error":   true,
				"message": "",
				"item": map[string]interface{}{
					"item_barcode":       p.BarCode,
					"file_path":          p.PicPath1,
					"is_cancel":          item.IsCancel,
					"is_check":           item.IsCheck,
					"item_code":          p.ItemCode,
					"item_name":          p.ItemName,
					"pickup_staff_name":  s.SaleName,
					"sale_code":          s.SaleCode + "/" + s.SaleName,
					"item_price":         p.SalePrice1,
					"qty_after":          req.QtyAfter,
					"qty_before":         item.QtyBefore,
					"qty_load":           item.QtyAfter,
					"total_price_after":  item.TotalPriceAfter,
					"total_price_before": p.SalePrice1 * req.QtyAfter,
					"item_unit_code":     p.UnitCode,
					"request_qty":        item.RequestQty,
					"item_qty":           req.QtyAfter,
					"pick_zone_id":       item.PickZoneId,
					"line_number":        req.LineNumber,
				},

				//"queid": map[string]interface{}{
				//	"item_barcode":       p.BarCode,
				//	"file_path":          p.PicPath1,
				//	"is_cancel":          item.IsCancel,
				//	"is_check":           item.IsCheck,
				//	"item_code":          p.ItemCode,
				//	"item_name":          p.ItemName,
				//	"pickup_staff_name":  s.SaleName,
				//	"sale_code":          s.SaleCode + "/" + s.SaleName,
				//	"item_price":         p.SalePrice1,
				//	"qty_after":          req.QtyAfter,
				//	"qty_before":         item.QtyBefore,
				//	"qty_load":           item.QtyAfter,
				//	"total_price_after":  item.TotalPriceAfter,
				//	"total_price_before": p.SalePrice1 * req.QtyAfter,
				//	"item_unit_code":     p.UnitCode,
				//	"request_qty":        item.RequestQty,
				//	"item_qty":           req.QtyAfter,
				//	"pick_zone_id":       item.PickZoneId,
				//	"line_number":        req.LineNumber,
			}, nil
		} else {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue is ref used",
				"item":    nil,
			}, nil
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

func PickupEdit(db *sqlx.DB, req *drivethru.PickupEditRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))

	if req.AccessToken == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Not Have Access Token",
		}, nil
	}

	if req.QId == 0 {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Id Not Assign",
		}, nil
	}

	u := UserAccess{}
	u.GetProfileByToken(db, req.AccessToken)

	q := ListQueueModel{}
	q.Search(db, req.QId)

	if q.Status >= 2 {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue can not edit",
		}, nil
	}

	if q.IsCancel == 1 {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue is cancel",
		}, nil
	}

	lccommand := `update basket set car_brand = ?, ref_number = ?, sale_id = ?, status = ?, edit_by = ?, edit_time = ? where que_id = ? and doc_date = CURDATE()`
	_, err := db.Exec(lccommand, req.CarBrand, req.CarNumber, u.UserId, req.Status, u.UserCode, now.String(), req.QId)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": err.Error(),
		}, nil
	}

	return map[string]interface{}{
		"success": true,
		"error":   false,
		"message": "",
	}, nil
}

func QueueEdit(db *sqlx.DB, req *drivethru.QueueEditRequest) (interface{}, error) {
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

	u := UserAccess{}
	u.GetProfileByToken(db, req.AccessToken)

	q := ListQueueModel{}
	q.Search(db, req.QueueId)

	if q.Status >= 2 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue can not edit",
			},
			"queid": ""}, nil
	}

	if q.IsCancel == 1 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": "Queue is cancel",
			},
			"queid": ""}, nil
	}

	lccommand := `update basket set car_brand = ?, ref_number = ?, sale_id = ?, status = ?, edit_by = ?, edit_time = ? where que_id = ? and doc_date = CURDATE()`
	_, err := db.Exec(lccommand, req.CarBrand, req.PlateNumber, u.UserId, req.Status, u.UserCode, now.String(), req.QueueId)
	if err != nil {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
			},
			"queid": ""}, nil
	}

	return map[string]interface{}{
		"response": map[string]interface{}{
			"success": true,
			"error":   false,
			"message": "",
		},
		"queid": ""}, nil
}

func (q *ListQueueModel) QueueStatus(db *sqlx.DB, req *drivethru.QueueStatusRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))

	if req.AccessToken == "" {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Not Have Access Token",
			"queue":   nil,
		}, nil
	}

	if req.QueueId == 0 {
		return map[string]interface{}{
			"success": false,
			"error":   true,
			"message": "Queue Id Not Assign",
			"queue":   nil,
		}, nil
	}

	u := UserAccess{}
	u.GetProfileByToken(db, req.AccessToken)

	q.Search(db, req.QueueId)

	if req.StatusForSaleorderCurrent == 1 && q.IsCancel == 0 {
		lccommand := `update basket set status = ?, pick_status = ? where que_id = ? and doc_date = CURDATE()`
		_, err := db.Exec(lccommand, req.StatusForSaleorderCurrent, req.StatusForSaleorderCurrent, req.QueueId)
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}

		lccommand1 := `insert basket_status(uuid, basket_id, que_id, doc_no, status, create_time) values(?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(lccommand1, q.UUID, req.QueueId, q.DocNo, req.StatusForSaleorderCurrent, now.String())
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}
	}

	if req.StatusForSaleorderCurrent == 2 && q.IsCancel == 0 {
		lccommand := `update basket set status = 0, pick_status = ? where que_id = ? and doc_date = CURDATE()`
		_, err := db.Exec(lccommand, req.StatusForSaleorderCurrent, req.QueueId)
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}

		lccommand1 := `insert basket_status(uuid, basket_id, que_id, doc_no, status, create_time) values(?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(lccommand1, q.UUID, req.QueueId, q.DocNo, req.StatusForSaleorderCurrent, now.String())
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}
	}

	if req.StatusForSaleorderCurrent == 3 && q.IsCancel == 0 {
		lccommand := `update basket set status = 0, pick_status = ? where que_id = ? and doc_date = CURDATE()`
		_, err := db.Exec(lccommand, req.StatusForSaleorderCurrent, req.QueueId)
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}

		lccommand1 := `insert basket_status(uuid, basket_id, que_id, doc_no, status, create_time) values(?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(lccommand1, q.UUID, req.QueueId, q.DocNo, req.StatusForSaleorderCurrent, now.String())
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}
	}

	if req.StatusForSaleorderCurrent == 4 && q.IsCancel == 0 {
		lccommand := `update basket set status = 0, pick_status = ?, is_cancel = 1, cancel_desc = ?,cancel_id = ?, cancel_time = ?  where que_id = ? and doc_date = CURDATE()`
		_, err := db.Exec(lccommand, req.StatusForSaleorderCurrent, req.CancelRemark, u.UserId, now.String(), req.QueueId)
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}

		lccommand1 := `insert basket_status(uuid, basket_id, que_id, doc_no, status, create_time) values(?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(lccommand1, q.UUID, req.QueueId, q.DocNo, req.StatusForSaleorderCurrent, now.String())
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   true,
				"message": err.Error(),
				"queue":   nil,
			}, nil
		}
	}

	q.QueueDetails(db, req.QueueId, req.AccessToken)

	return map[string]interface{}{
		"success": true,
		"error":   false,
		"message": "",
		"queue":   q,
	}, nil
}

func (q *ListQueueModel) BillingDone(db *sqlx.DB, req *drivethru.BillingDoneRequest) (interface{}, error) {
	var change_amount float64
	var crd_amount float64
	var cou_amount float64
	var dep_amount float64

	var remain_amount float64
	var remain_all_amount float64
	var sum_remain float64

	var item_amount float64

	var sqlcommand string
	var pos_no string

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02"))
	doc_date := now.AddDate(0, 0, 0).Format("2006-01-02")

	u := UserAccess{}
	u.GetProfileByToken(db, req.AccessToken)

	q.Search(db, req.QueueId)

	s := EmployeeModel{}
	s.SearchBySaleCode(db, u.UserCode)

	config := RequestConfigModel{}
	config.Search(db, u.CompanyID, u.BranchID)

	fmt.Println("TaxRate = ", config.TaxRate)

	fmt.Println("status", q.Status, q.IsCancel, q.TotalAfterAmount, config.TaxRate)

	if q.Status < 2 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success":              false,
				"error":                true,
				"message":              "Queue status not billing done",
				"total_amount":         0,
				"invoice":              nil,
				"is_print_short_form":  0,
				"is_print_cash_form":   0,
				"is_print_credit_form": 0,
			},
		}, nil
	}

	if q.Status == 3 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success":              false,
				"error":                true,
				"message":              "Queue is invoiced",
				"total_amount":         0,
				"invoice":              nil,
				"is_print_short_form":  0,
				"is_print_cash_form":   0,
				"is_print_credit_form": 0,
			},
		}, nil
	}

	if q.IsCancel == 1 {
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success":              false,
				"error":                true,
				"message":              "Queue is cancel",
				"total_amount":         0,
				"invoice":              nil,
				"is_print_short_form":  0,
				"is_print_cash_form":   0,
				"is_print_credit_form": 0,
			},
		}, nil
	}

	if q.Status == 2 {
		if req.Confirm == 0 {

			for _, i := range q.Item {
				item_amount = item_amount + (i.QtyAfter * i.ItemPrice)
			}

			if item_amount != q.TotalAfterAmount {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "ItemAmount not equal total_amount",
						"total_amount": item_amount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Not get invoice no",
							"cash_amount":    req.Cash,
							"change_amount":  change_amount,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}
			/////////////////////////////////////////////////////////////////////////////////////////////////////////////
			if len(req.CreditCard) != 0 {
				fmt.Println("CreditCard = ", len(req.CreditCard))
				for _, c := range req.CreditCard {
					fmt.Println("Credit Amount =", c.Amount)
					if c.Amount == 0 {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "coupon not have amount",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					if c.CardNo == "" || c.ConfirmNo == "" || c.CreditType == "" {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "credit card not have cardno or confirm no or credit type",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					var same_credit int

					for _, crd_card := range req.CreditCard {
						no := crd_card.ConfirmNo
						card := crd_card.CardNo

						same_credit = 0
						for _, cd := range req.CreditCard {

							if no == cd.ConfirmNo && card == cd.CardNo {
								same_credit = same_credit + 1
							}

							fmt.Println("same = ", same_credit)

							if same_credit > 1 {
								return map[string]interface{}{
									"response": map[string]interface{}{
										"success":      false,
										"error":        true,
										"message":      "confirm no have deplicate",
										"total_amount": item_amount,
										"invoice": map[string]interface{}{
											"invoice_no":     "Not get invoice no",
											"cash_amount":    req.Cash,
											"change_amount":  change_amount,
											"credit_amount":  crd_amount,
											"coupong_amount": cou_amount,
											"deposit_amount": dep_amount,
											"remain_amount":  sum_remain,
										},
										"is_print_short_form":  0,
										"is_print_cash_form":   0,
										"is_print_credit_form": 0,
									},
								}, nil
							}
						}
					}

					cd := CreditCard{}
					chkCrdUsed, msg := cd.CheckCreditCardUsed(db, c.CardNo, c.ConfirmNo)
					if chkCrdUsed == false {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      msg,
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					crd_amount = crd_amount + c.Amount
				}
			}
			/////////////////////////////////////////////////////////////////////////////////////////////////////////////
			if len(req.CouponCode) != 0 {

				var same_coupon int

				for _, c := range req.CouponCode {
					code := c.CouponCode
					same_coupon = 0
					for _, cp := range req.CouponCode {

						if code == cp.CouponCode {
							same_coupon = same_coupon + 1
						}

						fmt.Println("same = ", same_coupon)

						if same_coupon > 1 {
							return map[string]interface{}{
								"response": map[string]interface{}{
									"success":      false,
									"error":        true,
									"message":      "coupon have deplicate",
									"total_amount": item_amount,
									"invoice": map[string]interface{}{
										"invoice_no":     "Not get invoice no",
										"cash_amount":    req.Cash,
										"change_amount":  change_amount,
										"credit_amount":  crd_amount,
										"coupong_amount": cou_amount,
										"deposit_amount": dep_amount,
										"remain_amount":  sum_remain,
									},
									"is_print_short_form":  0,
									"is_print_cash_form":   0,
									"is_print_credit_form": 0,
								},
							}, nil
						}
					}
				}

				for _, p := range req.CouponCode {

					if p.CouponCode == "" {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "coupon not have code",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					if p.Amount == 0 {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "coupon not have amount",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					cp := Coupon{}
					chkCouUsed, msg := cp.CheckCouponUsed(db, p.CouponCode, p.Amount)
					if chkCouUsed == false {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      msg,
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					cou_amount = cou_amount + p.Amount
				}
			}
			/////////////////////////////////////////////////////////////////////////////////////////////////////////////
			if len(req.DepositAmount) != 0 {
				var same_deposit int

				for _, d := range req.DepositAmount {
					dep_no := d.DepositId
					same_deposit = 0
					for _, dp := range req.DepositAmount {

						if dep_no == dp.DepositId {
							same_deposit = same_deposit + 1
						}

						fmt.Println("same = ", same_deposit)

						if same_deposit > 1 {
							return map[string]interface{}{
								"response": map[string]interface{}{
									"success":      false,
									"error":        true,
									"message":      "deposit have deplicate",
									"total_amount": item_amount,
									"invoice": map[string]interface{}{
										"invoice_no":     "Not get invoice no",
										"cash_amount":    req.Cash,
										"change_amount":  change_amount,
										"credit_amount":  crd_amount,
										"coupong_amount": cou_amount,
										"deposit_amount": dep_amount,
										"remain_amount":  sum_remain,
									},
									"is_print_short_form":  0,
									"is_print_cash_form":   0,
									"is_print_credit_form": 0,
								},
							}, nil
						}
					}

					dp := Deposit{}
					chkDepUsed, msg := dp.CheckArDepositUsed(db, req.ArCode, d.DepositId, d.Amount)
					if chkDepUsed == false {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      msg,
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}
					dep_amount = dep_amount + d.Amount
				}
			}

			remain_amount = (((q.TotalAfterAmount - crd_amount) - cou_amount) - dep_amount);
			fmt.Println("remain amount =", remain_amount)

			if (remain_amount < 0 && req.Cash != 0) {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "Payment cash over remain",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Can not save bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}

			if ((remain_amount > 0 && req.Cash > 0 && (remain_amount-req.Cash < 0)) || (remain_amount == 0)) {
				change_amount = -1 * (remain_amount - req.Cash);
			} else {
				change_amount = 0;
			}
			fmt.Println("change amount =", change_amount)

			if remain_amount < 0 {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "Payment over netamount",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Can not save bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}

			remain_all_amount = remain_amount - req.Cash + change_amount;
			if (remain_all_amount > 0) {
				sum_remain = remain_all_amount;
			} else {
				sum_remain = 0;
			}

			fmt.Println("total", q.TotalAfterAmount, req.Cash, crd_amount, cou_amount, dep_amount, remain_all_amount, sum_remain)

			if sum_remain != 0 {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "Payment have remain",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Can not save bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			} else {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      true,
						"error":        false,
						"message":      "",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Queue is aready for bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}

		} else { //IsConfirm == 1

			for _, i := range q.Item {
				item_amount = item_amount + (i.QtyAfter * i.ItemPrice)
			}

			fmt.Println("total_amount, item_amount = ", q.TotalAfterAmount, item_amount)

			if item_amount != q.TotalAfterAmount {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "ItemAmount not equal total_amount",
						"total_amount": item_amount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Not get invoice no",
							"cash_amount":    req.Cash,
							"change_amount":  change_amount,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}

			/////////////////////////////////////////////////////////////////////////////////////////////////////////////
			if len(req.CreditCard) != 0 {
				fmt.Println("CreditCard = ", len(req.CreditCard))
				for _, c := range req.CreditCard {
					fmt.Println("Credit Amount =", c.Amount)
					if c.Amount == 0 {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "coupon not have amount",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					if c.CardNo == "" || c.ConfirmNo == "" || c.CreditType == "" {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "credit card not have cardno or confirm no or credit type",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					var same_credit int

					for _, crd_card := range req.CreditCard {
						no := crd_card.ConfirmNo
						card := crd_card.CardNo

						same_credit = 0
						for _, cd := range req.CreditCard {

							if no == cd.ConfirmNo && card == cd.CardNo {
								same_credit = same_credit + 1
							}

							fmt.Println("same = ", same_credit)

							if same_credit > 1 {
								return map[string]interface{}{
									"response": map[string]interface{}{
										"success":      false,
										"error":        true,
										"message":      "confirm no have deplicate",
										"total_amount": item_amount,
										"invoice": map[string]interface{}{
											"invoice_no":     "Not get invoice no",
											"cash_amount":    req.Cash,
											"change_amount":  change_amount,
											"credit_amount":  crd_amount,
											"coupong_amount": cou_amount,
											"deposit_amount": dep_amount,
											"remain_amount":  sum_remain,
										},
										"is_print_short_form":  0,
										"is_print_cash_form":   0,
										"is_print_credit_form": 0,
									},
								}, nil
							}
						}
					}

					cd := CreditCard{}
					chkCrdUsed, msg := cd.CheckCreditCardUsed(db, c.CardNo, c.ConfirmNo)
					if chkCrdUsed == false {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      msg,
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					crd_amount = crd_amount + c.Amount
				}
			}
			/////////////////////////////////////////////////////////////////////////////////////////////////////////////
			if len(req.CouponCode) != 0 {

				var same_coupon int

				for _, c := range req.CouponCode {
					code := c.CouponCode
					same_coupon = 0
					for _, cp := range req.CouponCode {

						if code == cp.CouponCode {
							same_coupon = same_coupon + 1
						}

						fmt.Println("same = ", same_coupon)

						if same_coupon > 1 {
							return map[string]interface{}{
								"response": map[string]interface{}{
									"success":      false,
									"error":        true,
									"message":      "coupon have deplicate",
									"total_amount": item_amount,
									"invoice": map[string]interface{}{
										"invoice_no":     "Not get invoice no",
										"cash_amount":    req.Cash,
										"change_amount":  change_amount,
										"credit_amount":  crd_amount,
										"coupong_amount": cou_amount,
										"deposit_amount": dep_amount,
										"remain_amount":  sum_remain,
									},
									"is_print_short_form":  0,
									"is_print_cash_form":   0,
									"is_print_credit_form": 0,
								},
							}, nil
						}
					}
				}

				for _, p := range req.CouponCode {

					if p.CouponCode == "" {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "coupon not have code",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					if p.Amount == 0 {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "coupon not have amount",
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					cp := Coupon{}
					chkCouUsed, msg := cp.CheckCouponUsed(db, p.CouponCode, p.Amount)
					if chkCouUsed == false {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      msg,
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}

					cou_amount = cou_amount + p.Amount
				}
			}
			/////////////////////////////////////////////////////////////////////////////////////////////////////////////
			if len(req.DepositAmount) != 0 {
				var same_deposit int
				for _, d := range req.DepositAmount {
					dep_no := d.DepositId
					same_deposit = 0
					for _, dp := range req.DepositAmount {

						if dep_no == dp.DepositId {
							same_deposit = same_deposit + 1
						}

						fmt.Println("same = ", same_deposit)

						if same_deposit > 1 {
							return map[string]interface{}{
								"response": map[string]interface{}{
									"success":      false,
									"error":        true,
									"message":      "deposit have deplicate",
									"total_amount": item_amount,
									"invoice": map[string]interface{}{
										"invoice_no":     "Not get invoice no",
										"cash_amount":    req.Cash,
										"change_amount":  change_amount,
										"credit_amount":  crd_amount,
										"coupong_amount": cou_amount,
										"deposit_amount": dep_amount,
										"remain_amount":  sum_remain,
									},
									"is_print_short_form":  0,
									"is_print_cash_form":   0,
									"is_print_credit_form": 0,
								},
							}, nil
						}
					}

					dp := Deposit{}
					chkDepUsed, msg := dp.CheckArDepositUsed(db, req.ArCode, d.DepositId, d.Amount)
					if chkDepUsed == false {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      msg,
								"total_amount": item_amount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Not get invoice no",
									"cash_amount":    req.Cash,
									"change_amount":  change_amount,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}
					dep_amount = dep_amount + d.Amount
					fmt.Println("dep_amount =", dep_amount)
				}
			}

			remain_amount = (((q.TotalAfterAmount - crd_amount) - cou_amount) - dep_amount);
			fmt.Println("remain amount =", remain_amount)

			if (remain_amount < 0 && req.Cash != 0) {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "Payment cash over remain",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Can not save bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}

			if ((remain_amount > 0 && req.Cash > 0 && (remain_amount-req.Cash < 0)) || (remain_amount == 0)) {
				change_amount = -1 * (remain_amount - req.Cash);
			} else {
				change_amount = 0;
			}
			fmt.Println("change amount =", change_amount)

			if remain_amount < 0 {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "Payment over netamount",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Can not save bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			}

			remain_all_amount = remain_amount - req.Cash + change_amount;
			if (remain_all_amount > 0) {
				sum_remain = remain_all_amount;
			} else {
				sum_remain = 0;
			}

			fmt.Println("total", q.TotalAfterAmount, req.Cash, crd_amount, cou_amount, dep_amount, remain_all_amount, sum_remain)

			if sum_remain != 0 {
				return map[string]interface{}{
					"response": map[string]interface{}{
						"success":      false,
						"error":        true,
						"message":      "Payment have remain",
						"total_amount": q.TotalAfterAmount,
						"invoice": map[string]interface{}{
							"invoice_no":     "Can not save bill",
							"cash_amount":    req.Cash,
							"credit_amount":  crd_amount,
							"coupong_amount": cou_amount,
							"deposit_amount": dep_amount,
							"remain_amount":  sum_remain,
							"change_amount":  change_amount,
						},
						"is_print_short_form":  0,
						"is_print_cash_form":   0,
						"is_print_credit_form": 0,
					},
				}, nil
			} else {

				var before_tax_amount float64
				var tax_amount float64

				uuid := GetAccessToken()
				m := Machine{}
				m.SearchMachineNo(db, u.CompanyID, u.BranchID, req.AccessToken)
				fmt.Println("machine = ", m.ShiftUUID, m.MachineId)

				var err error
				var access_token string

				access_token = req.AccessToken

				pos_no, err = getPosNo(db, u.CompanyID, u.BranchID, access_token)
				if err != nil {
					fmt.Println(err.Error())
				}

				cust := CustomerModel{}
				cust.Search(db, req.ArCode)

				fmt.Println("TaxRate = ", config.TaxRate)

				ar_id := cust.Id
				total_amount := q.TotalAfterAmount
				before_tax_amount = (q.TotalAfterAmount * 100) / float64(config.TaxRate)
				tax_amount = q.TotalAfterAmount - before_tax_amount

				fmt.Println("pos_no, total_amount,before_tax_amount,tax_amount, uuid, ar_id", pos_no, total_amount, before_tax_amount, tax_amount, uuid, ar_id)
				bill_type := 0
				tax_type := config.SaleTaxType
				allocate_id := 0
				project_id := 0
				depart_id := 0
				pos_status := 1
				my_description := "Pos Drivethru"
				discount_word := ""
				discount_amount := 0
				chq_amount := 0
				bnk_amount := 0
				onl_amount := 0

				//sqlcommand = `START TRANSACTION`
				//_, err = db.Exec(sqlcommand)

				sqlcommand = `insert into ar_invoice(company_id,branch_id,uuid,doc_no,tax_no,doc_date,doc_type,bill_type ,tax_type ,tax_rate,pos_machine_id,shift_uuid,cash_id,number_of_item,ar_id,ar_code,ar_name,sale_id,sale_code,sale_name,depart_id,allocate_id,project_id,pos_status,my_description,so_ref_no ,sum_of_item_amount,discount_word,discount_amount,after_discount_amount,before_tax_amount,tax_amount,total_amount,change_amount,coupon_amount,sum_cash_amount,sum_chq_amount,sum_credit_amount ,sum_bank_amount,sum_of_deposit,sum_on_line_amount,net_debt_amount,car_license,create_by ,create_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
				rs, err := db.Exec(sqlcommand, u.CompanyID, u.BranchID, uuid, pos_no, pos_no, doc_date, 1, bill_type, tax_type, config.TaxRate, m.MachineId, m.ShiftUUID, m.CashierID, q.NumberOfItem, ar_id, req.ArCode, cust.Name, s.Id, s.SaleCode, s.SaleName, depart_id, allocate_id, project_id, pos_status, my_description, q.DocNo, q.TotalAfterAmount, discount_word, discount_amount, q.TotalAfterAmount, before_tax_amount, tax_amount, total_amount, change_amount, cou_amount, req.Cash, chq_amount, crd_amount, bnk_amount, dep_amount, onl_amount, total_amount, q.PlateNumber, u.UserCode, now.String())
				if err != nil {
					return map[string]interface{}{
						"response": map[string]interface{}{
							"success":      false,
							"error":        true,
							"message":      "error arinvoice = " + err.Error(),
							"total_amount": q.TotalAfterAmount,
							"invoice": map[string]interface{}{
								"invoice_no":     "Can not save bill",
								"cash_amount":    req.Cash,
								"credit_amount":  crd_amount,
								"coupong_amount": cou_amount,
								"deposit_amount": dep_amount,
								"remain_amount":  sum_remain,
								"change_amount":  change_amount,
							},
							"is_print_short_form":  0,
							"is_print_cash_form":   0,
							"is_print_credit_form": 0,
						},
					}, nil
				}

				id, _ := rs.LastInsertId()

				fmt.Println("ar_invoice_id = ", int(id))

				var discount_word_sub string
				var discount_amount_sub float64
				var sum_of_cost float64

				for _, item := range q.Item {

					fmt.Println("item", item.ItemBarCode)
					discount_word_sub = ""
					discount_amount_sub = 0

					sum_of_cost = item.QtyAfter * item.AverageCost

					sqlcommand = `insert into ar_invoice_sub(company_id, branch_id, uuid, inv_id, doc_no, doc_date, ar_id, sale_id, item_id, item_code, bar_code, item_name, wh_id, shelf_id, qty, cn_qty, unit_code, price, discount_word_sub, discount_amount_sub, amount, net_amount, average_cost, sum_of_cost, item_decription, packing_rate_1, ref_no, ref_line_number, line_number) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
					rs_sub, err := db.Exec(sqlcommand, u.CompanyID, u.BranchID, uuid, q.Id, pos_no, doc_date, ar_id, u.UserId, item.ItemId, item.ItemCode, item.ItemBarCode, item.ItemName, m.DefWhId, m.DefShelfId, item.QtyAfter, item.QtyAfter, item.ItemUnitCode, item.ItemPrice, discount_word_sub, discount_amount_sub, item.TotalPriceAfter, item.TotalPriceAfter, item.AverageCost, sum_of_cost, q.PlateNumber, item.Rate1, q.DocNo, item.LineNumber, item.LineNumber)
					if err != nil {
						return map[string]interface{}{
							"response": map[string]interface{}{
								"success":      false,
								"error":        true,
								"message":      "error arinvoice_sub = " + err.Error(),
								"total_amount": q.TotalAfterAmount,
								"invoice": map[string]interface{}{
									"invoice_no":     "Can not save bill",
									"cash_amount":    req.Cash,
									"credit_amount":  crd_amount,
									"coupong_amount": cou_amount,
									"deposit_amount": dep_amount,
									"remain_amount":  sum_remain,
									"change_amount":  change_amount,
								},
								"is_print_short_form":  0,
								"is_print_cash_form":   0,
								"is_print_credit_form": 0,
							},
						}, nil
					}
					item_id, _ := rs_sub.LastInsertId()

					item.Id = int(item_id)
				}

				if req.Cash != 0 {
					lccommand_rec := `insert into rec_money(company_id, branch_id, uuid, doc_type, ref_id, ar_id , doc_date, payment_type, pay_amount, line_number) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
					_, err := db.Exec(lccommand_rec, u.CompanyID, u.BranchID, uuid, 1, q.Id, ar_id, doc_date, 0, req.Cash, 0)
					if err != nil {
						fmt.Println("error insert recmoney cash = ", err.Error())
					}

				}

				var crd_line_number int

				if req.Cash != 0 {
					crd_line_number = 1
				} else {
					crd_line_number = 0
				}

				if len(req.CreditCard) != 0 {
					for _, crd := range req.CreditCard {
						lccommand_crd := `insert into credit_card(company_id, branch_id,ref_uuid, ref_id,ar_id,doc_no, doc_date, credit_card_no, credit_type, confirm_no, amount, charge_amount, description,receive_date, due_date) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
						_, err = db.Exec(lccommand_crd, u.CompanyID, u.BranchID, uuid, q.Id, ar_id, pos_no, doc_date, crd.CardNo, crd.CreditType, crd.ConfirmNo, crd.Amount, crd.ChargeAmount, "Drivethru Pos", doc_date, doc_date)
						if err != nil {
							fmt.Println("error insert credit card = ", err.Error())
						}

						//lccommand_rec := `insert into rec_money(company_id, branch_id, uuid, doc_type, ref_id, ar_id , doc_date, payment_type, pay_amount, chq_total_amount, credit_type, charge_amount, confirm_no, ref_no, bank_code, ref_date, line_number) values(?, ?, ?, ?, ?, ? , ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
						//_, err := db.Exec(lccommand_rec, u.CompanyID, u.BranchID, uuid, 1, q.Id, ar_id, doc_date, 1, crd.Amount, crd.Amount, crd.CreditType, crd.ChargeAmount, crd.ConfirmNo, crd.CardNo, crd.BankCode, doc_date, crd_line_number)
						//if err != nil {
						//	fmt.Println("error insert recmoney creditcard = ", err.Error())
						//}
					}
					crd_line_number = crd_line_number + 1
				}

				var line_number_coupon int
				if len(req.CouponCode) != 0 {
					for _, cou := range req.CouponCode {
						line_number_coupon = 0
						lccommand_cou := `insert into coupon_receive(company_id, branch_id, coupon_code, coupon_type, ref_doc_no, ref_uuid, coupon_value, line_number) values(?, ?, ?, ?, ?, ?, ?, ?)`
						_, err = db.Exec(lccommand_cou, u.CompanyID, u.BranchID, cou.CouponCode, 1, q.DocNo, uuid, cou.Amount, line_number_coupon)
						if err != nil {
							fmt.Println("error insert credit card = ", err.Error())
						}

						line_number_coupon = line_number_coupon + 1
					}
				}

				var line_number_deposit int
				if len(req.DepositAmount) != 0 {
					for _, dep := range req.DepositAmount {
						line_number_coupon = 0
						lccommand_cou := `insert into deposit_use(company_id, branch_id, deposit_no, ref_doc_no, ref_uuid, balance, amount, net_amount, line_number) values(?, ?, ?, ?, ?, ?, ?, ?, ?)`
						rs, err = db.Exec(lccommand_cou, u.CompanyID, u.BranchID, dep.DepositId, pos_no, uuid, dep.Amount, dep.Amount, dep.Amount, line_number_coupon)
						if err != nil {
							fmt.Println("error insert deposit = ", err.Error())
						}

						dep_id, err := rs.LastInsertId()
						if err != nil {
							fmt.Println("error insert dep use = " + err.Error())
						}

						line_number_deposit = line_number_deposit + 1

						if dep_id != 0 {
							lccommand := `update ar_deposit set balance = balance-? where doc_no=? and company_id = ? and branch_id = ?`
							_, err = db.Exec(lccommand, dep.Amount, dep.DepositId, u.CompanyID, u.BranchID)
							if err != nil {
								fmt.Println("error update deposit = ", err.Error())
							}
						}
					}
				}

				fmt.Println("update basket = ", pos_no, u.UserId, q.Id, req.QueueId)
				lccommand := `update basket set invoice_no = ?,status = 3,confirm_by=?, confirm_time=? where id = ? and que_id = ? `
				_, err = db.Exec(lccommand, pos_no, u.UserId, now.String(), q.Id, req.QueueId)
				if err != nil {
					fmt.Println("error update basket = ", err.Error())
				}

				//
				//sqlcommand = `COMMIT`
				//_, err = db.Exec(sqlcommand)
			}

		}
	}
	return map[string]interface{}{
		"response": map[string]interface{}{
			"success":      true,
			"error":        false,
			"message":      "",
			"total_amount": q.TotalAfterAmount,
			"invoice": map[string]interface{}{
				"invoice_no":     pos_no,
				"cash_amount":    req.Cash,
				"credit_amount":  crd_amount,
				"coupong_amount": cou_amount,
				"deposit_amount": dep_amount,
				"remain_amount":  sum_remain,
				"change_amount":  change_amount,
			},
			"is_print_short_form":  0,
			"is_print_cash_form":   0,
			"is_print_credit_form": 0,
		},
	}, nil
}

func (q *ListQueueModel) CancelQueue(db *sqlx.DB, req *drivethru.QueueStatusRequest) (interface{}, error) {

	if (req.QueueId != 0) {
		q.Search(db, req.QueueId)

		if (q.IsCancel == 0) {
			u := UserAccess{}
			u.GetProfileByToken(db, req.AccessToken)

			if (q.Status != 2) {
				lccommand := "update basket set status = 0,pick_status=4,is_cancel=1,cancel_desc=?,cancel_by = ?,cancel_time= CURRENT_TIMESTAMP() where que_id = ? and company_id = ? and branch_id = ? and uuid = ? and doc_date = curdate()";
				_, err := db.Exec(lccommand, req.CancelRemark, u.UserCode, req.QueueId, u.CompanyID, u.BranchID, q.UUID)
				if err != nil {
					return map[string]interface{}{
						"response": map[string]interface{}{
							"success": false,
							"error":   true,
							"message": err.Error(),
						},
					}, nil
				}
			}
		}
	}
	return map[string]interface{}{
		"response": map[string]interface{}{
			"success": true,
			"error":   false,
			"message": "",
		},
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

func getBasketNo(db *sqlx.DB, company_id int, branch_id int, doc_type int) (string, error) {
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

	last_number1, _ = getLastBasketNo(db, company_id, branch_id, doc_type)
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

	intday = int(time.Now().Day())
	intday1 = int(intday)
	vday = strconv.Itoa(intday1)

	fmt.Println("day =", vday)

	lenday = len(vday)

	fmt.Println("len day =", lenday)

	if lenday == 1 {
		vday1 = "0" + vday
	} else {
		vday1 = vday
	}

	fmt.Println("vDay = ", vday1)

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

func getLastBasketNo(db *sqlx.DB, company_id int, branch_id int, doc_type int) (last_no int, err error) {
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

func getPosNo(db *sqlx.DB, company_id int, branch_id int, access_token string) (string, error) {
	var last_number1 int
	var last_number string
	var snumber string
	var intyear int
	var vHeader string
	//var branch_header string
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

	m := Machine{}
	m.SearchMachineNo(db, company_id, branch_id, access_token)

	last_number1, _ = getLastPosNo(db, company_id, branch_id, m.MachineId)
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

	intday = int(time.Now().Day())
	intday1 = int(intday)
	vday = strconv.Itoa(intday1)

	fmt.Println("day =", vday)

	lenday = len(vday)

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

	//if branch_id == 1 {
	//	branch_header = "S01"
	//} else {
	//	branch_header = "S02"
	//}

	fmt.Println("POS NO Machine No = ", m.MachineNo)

	header = "D" + m.MachineNo

	vHeader = header

	NewDocNo := vHeader + vyear1 + vmonth1 + vday1 + "-" + snumber
	fmt.Println(snumber)
	fmt.Println(vHeader)

	fmt.Println("NewDocNo = ", NewDocNo)

	return NewDocNo, nil
}

func getLastPosNo(db *sqlx.DB, company_id int, branch_id int, machine_id int) (last_no int, err error) {

	fmt.Println("Pos Last MachineId =", machine_id)

	sql := `select cast(right(ifnull(max(doc_no),0),4) as int)+1 maxno from ar_invoice where company_id = ? and branch_id = ? and pos_machine_id = ? and year(doc_date) = year(CURDATE()) and month(doc_date) = month(CURDATE()) and day(doc_date) = day(CURDATE())`
	fmt.Println("Branch ID =", branch_id)
	fmt.Println("Query = ", sql)
	err = db.Get(&last_no, sql, company_id, branch_id, machine_id)
	if err != nil {
		//fmt.Println("Last No Error = ",err)
		return 1, nil
	}

	fmt.Println("Last No = ", last_no)
	return last_no, nil
}

func (q *ListQueueModel) Search(db *sqlx.DB, queue_id int) {
	fmt.Println("q = ", queue_id)

	lccommand := "select id, que_id as queue_id, car_brand, ref_number as plate_number,uuid, doc_date, number_of_item, create_time as time_created, status, is_cancel, '' as ar_code, '' as ar_name, '' as sale_name, '' as sale_code, doc_no, doc_type as source, '' as receiver_name, pickup_time as pickup_datetime, 0 as is_loaded, status as status_for_saleorder_current," +
		" ifnull((select sum(pick_amount) from basket_sub where basket_id = a.id and que_id = a.que_id and uuid = a.uuid GROUP BY basket_id,que_id),0) as total_amount, " +
		" ifnull((select sum(pick_amount) from basket_sub where basket_id = a.id and que_id = a.que_id and uuid = a.uuid GROUP BY basket_id,que_id),0) as total_before_amount, " +
		" ifnull((select sum(checkout_amount) from basket_sub where basket_id = a.id and que_id = a.que_id and uuid = a.uuid GROUP BY basket_id,que_id),0) as total_after_amount, '' as otp_password, 0 as bill_type, '' as cancel_remark, '' as who_cancel, '' as sale_order from basket a where que_id=? and doc_date = CURRENT_DATE "
	rs := db.QueryRow(lccommand, queue_id)
	rs.Scan(&q.Id, &q.QueueId, &q.CarBrand, &q.PlateNumber, &q.UUID, &q.DocDate, &q.NumberOfItem, &q.TimeCreated, &q.Status, &q.IsCancel, &q.ArCode, &q.ArName, &q.SaleName, &q.SaleCode, &q.DocNo, &q.Source, &q.ReceiverName, &q.PickupDateTime, &q.TotalAmount, &q.IsLoaded, &q.StatusForSaleOrderCurrent, &q.TotalBeforeAmount, &q.TotalAfterAmount, &q.OTPPassword, &q.BillType, &q.CancelRemark, &q.WhoCancel, &q.SaleOrder)
	fmt.Println("q CarBrand = ", q.Id, q.QueueId, q.CarBrand, q.PlateNumber)

	lccommand1 := `select id, item_id, item_code, item_name ,bar_code as item_bar_code, request_qty, pick_qty as qty_before, checkout_qty as qty_after, price as item_price, unit_code as item_unit_code, pick_amount as total_price_before, checkout_amount as total_price_after, rate1, '' as sale_code, average_cost, line_number, '' as pick_zone_id from basket_sub where basket_id = ? and que_id = ? and uuid = ? and doc_date = CURDATE() order by line_number`
	err := db.Select(&q.Item, lccommand1, q.Id, q.QueueId, q.UUID)
	if err != nil {
		fmt.Println("error item = ", err.Error())
	}

	lccommand2 := `select phone_no from owner_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
	err = db.Select(&q.OwnerPhone, lccommand2, q.Id, q.QueueId, q.UUID, q.DocNo)
	if err != nil {
		fmt.Println("error owner phone = ", err.Error())
	}

	lccommand3 := `select phone_no from order_trust_phone where basket_id = ? and que_id = ? and uuid = ? and doc_no = ?  order by id`
	err = db.Select(&q.ReceiverPhone, lccommand3, q.Id, q.QueueId, q.UUID, q.DocNo)
	if err != nil {
		fmt.Println("error receive phone = ", err.Error())
	}

	return
}

func (itm *QueueItem) SearchQueueItem(db *sqlx.DB, queue_id int, item_code string, unit_code string, line_number int) {
	fmt.Println("q = ", queue_id)

	lccommand := `select id, item_id, item_code, item_name ,bar_code as item_bar_code, request_qty, pick_qty as qty_before, checkout_qty as qty_after, price as item_price, unit_code as item_unit_code, pick_amount as total_price_before, checkout_amount as total_price_after, rate1, '' as sale_code, average_cost, line_number, '' as pick_zone_id, ifnull(b.SaleName,'') as PickupStaffName, is_check_out as IsCheck from basket_sub a left join Sale b on a.pick_by = b.id where que_id=? and item_code = ? and unit_code = ? and doc_date = CURRENT_DATE `
	rs := db.QueryRow(lccommand, queue_id, item_code, unit_code)
	rs.Scan(&itm.Id, &itm.Id, &itm.ItemId, &itm.ItemCode, &itm.ItemName, &itm.ItemBarCode, &itm.RequestQty, &itm.QtyBefore, &itm.QtyAfter, &itm.ItemPrice, &itm.ItemUnitCode, &itm.TotalPriceBefore, &itm.TotalPriceAfter, &itm.Rate1, &itm.SaleCode, &itm.AverageCost, &itm.LineNumber, &itm.PickZoneId, &itm.PickupStaffName, &itm.IsCheck)
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
