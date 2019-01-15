package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/mrtomyum/nopadol/pointofsale"
	"time"
	"github.com/satori/go.uuid"
	"errors"
)

type BasketModel struct {
	Id             int64            `db:"id" db:"id"`
	CompanyId      int64            `db:"company_id"`
	BranchId       int64            `db:"branch_id"`
	UUID           string           `db:"uuid"`
	InvoiceNo      string           `db:"invoice_no"`
	QueId          int64            `db:"que_id"`
	DocType        int64            `db:"doc_type"`
	DocDate        string           `db:"doc_date"`
	ArId           int64            `db:"ar_id"`
	SaleId         int64            `db:"sale_id"`
	TaxType        int64            `db:"tax_type"`
	TaxRate        int64            `db:"tax_rate"`
	NumberOfItem   int64            `db:"number_of_item"`
	SumItemAmount  float64          `db:"sum_item_amount"`
	TotalAmount    float64          `db:"total_amount"`
	NetAmount      float64          `db:"net_amount"`
	OtpPassword    string           `db:"otp_password"`
	Status         int64            `db:"status"`
	PickStatus     int64            `db:"pick_status"`
	DeliveryStatus int64            `db:"delivery_status"`
	ReceiveName    string           `db:"receive_name"`
	ReceiveTel     string           `db:"receive_tel"`
	IsCancel       int64            `db:"is_cancel"`
	IsConfirm      int64            `db:"is_confirm"`
	CreateBy       string           `db:"create_by"`
	CreateTime     string           `db:"create_time"`
	EditBy         string           `db:"edit_by"`
	EditTime       string           `db:"edit_time"`
	ConfirmBy      string           `db:"confirm_by"`
	ConfirmTime    string           `db:"confirm_time"`
	CancelBy       string           `db:"cancel_by"`
	CancelTime     string           `db:"cancel_time"`
	CancelDescId   int64            `db:"cancel_desc_id"`
	CancelDesc     string           `db:"cancel_desc"`
	Sub            []BasketSubModel `db:"sub"`
}

type BasketSubModel struct {
	Id              int64   `db:"id"`
	BasketId        int64   `db:"basket_id"`
	Uuid            string  `db:"uuid"`
	QueId           int64   `db:"que_id"`
	DocDate         string  `db:"doc_date"`
	ItemId          int64   `db:"item_id"`
	ItemCode        string  `db:"item_code"`
	ItemName        string  `db:"item_name"`
	BarCode         string  `db:"bar_code"`
	RequestQty      float64 `db:"request_qty"`
	PickQty         float64 `db:"pick_qty"`
	CheckoutQty     float64 `db:"checkout_qty"`
	Price           float64 `db:"price"`
	UnitId          int64   `db:"unit_id"`
	PickAmount      float64 `db:"pick_amount"`
	CheckoutAmount  float64 `json:"checkout_amount"`
	Qty             float64 `db:"qty"`
	RemainQty       float64 `db:"remain_qty"`
	Rate1           int64   `db:"rate_1"`
	RefNo           string  `db:"ref_no"`
	SaleId          int64   `db:"sale_id"`
	AverageCost     float64 `db:"average_cost"`
	DeliveryOrderId int64   `db:"delivery_order_id"`
	LineNumber      int64   `db:"line_number"`
	RequestBy       string  `db:"request_by"`
	RequestTime     string  `db:"request_time"`
	PickBy          string  `db:"pick_by"`
	PickTime        string  `db:"pick_time"`
	CheckoutBy      string  `db:"checkout_by"`
	CheckoutTime    string  `db:"checkout_time"`
}

type ConfigurationModel struct {
	CompanyId      int64  `db:"company_id"`
	BranchId       int64  `db:"branch_id"`
	TaxType        int64  `db:"tax_type"`
	TaxRate        int64  `db:"tax_rate"`
	LogoPath       string `db:"logo_path"`
	DepartId       int64  `db:"depart_id"`
	DefSaleWHId    int64  `db:"def_sale_wh_id"`
	DefSaleShelfId int64  `db:"def_sale_shelf_id"`
	DefBuyWHId     int64  `db:"def_buy_wh_id"`
	DefBuyShelfId  int64  `db:"def_buy_shelf_id"`
	StockStatus    int64  `db:"stock_status"`
	SaleTaxType    int64  `db:"sale_tax_type"`
	BuyTaxType     int64  `db:"buy_tax_type"`
	SaleBillType   int64  `db:"sale_bill_type"`
	BuyBillType    int64  `db:"buy_bill_type"`
	UseAddress     int64  `db:"use_address"`
}

type MachineModel struct {
	cash_id      int64  `db:"cash_id"`
	machine_id   int64  `db:"machine_id"`
	machine_code string `db:"machine_code"`
	machine_no   string `db:"machine_no"`
}

type ItemData struct {
	ItemId      int64   `json:"item_id" db:"ItemId"`
	ItemCode    string  `json:"item_code" db:"ItemCode"`
	ItemName    string  `json:"item_name" db:"ItemName"`
	BarCode     string  `json:"bar_code" db:"BarCode"`
	ItemPrice   float64 `json:"item_price" db:"ItemPrice"`
	UnitId      int64   `json:"unit_id" db:"UnitId"`
	UnitCode    string  `json:"unit_code" db:"UnitCode"`
	Rate1       int64   `json:"rate_1" db:"Rate1"`
	AverageCost float64 `json:"average_cost" db:"AverageCost"`
}

type pointofsaleRepository struct{ db *sqlx.DB }

var que_id int64

func NewPointOfSaleRepository(db *sqlx.DB) pointofsale.Repository {
	return &pointofsaleRepository{db}
}

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err.Error()
	}
	fmt.Printf("UUIDv4: %s\n", uuid)

	return uuid.String()
}

func GetBranchConfig(db *sqlx.DB) ConfigurationModel {
	config := ConfigurationModel{}
	sql := `select tax_type,tax_rate,ifnull(logo_path,'') as logo_path,depart_id,def_sale_wh_id,def_sale_shelf_id,def_buy_wh_id,def_buy_shelf_id,stock_status,sale_tax_type,buy_tax_type,sale_bill_type,buy_bill_type,use_address from configuration where company_id=1 and branch_id = 1`
	//sql := `select * from company a inner join branch b on a.id = b.company_id left join configuration c on a.id = c.company_id and b.id = c.branch_id inner join pos_machine d on a.id = d.company_id and b.id = d.branch_id`
	err := db.Get(&config, sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}

func GenQueId(db *sqlx.DB) int64 {
	sql := `select ifnull(max(que_id),0)+1 as maxQue from basket where doc_date = CURRENT_DATE()`
	err := db.Get(&que_id, sql)
	if err != nil {
		fmt.Println(err.Error())
	}

	return que_id
}

func (repo *pointofsaleRepository) Create(req *pointofsale.BasketTemplate) (interface{}, error) {
	que_id := GenQueId(repo.db)
	uuid := GetUUID()

	fmt.Println("Que Id = ", que_id, uuid)
	req.UUID = uuid
	req.QueId = que_id

	config := ConfigurationModel{}
	config = GetBranchConfig(repo.db)

	req.TaxType = config.TaxType
	req.TaxRate = config.TaxRate
	req.CompanyId = 1
	req.BranchId = 1
	req.IsConfirm = 0
	req.IsCancel = 0
	req.DeliveryStatus = 0
	req.TotalAmount = 0
	req.NumberOfItem = 0
	req.NetAmount = 0
	req.Status = 0
	req.SumItemAmount = 0
	req.PickStatus = 0
	req.DocType = 1

	fmt.Println("DocDate =", req.DocDate)

	now := time.Now()

	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	if req.DocDate == "" {
		req.DocDate = DocDate
		fmt.Println("DocDate Default =", req.DocDate)
	}

	sql := `insert into basket(company_id, branch_id, uuid, que_id, doc_type, doc_date, ar_id, sale_id)values(?, ?, ?, ?, ?, ?, ?, ?)`//, tax_type, tax_rate, number_of_item, sum_item_amount, total_amount, net_amount, status, pick_status, delivery_status, receive_name, receive_tel, is_cancel, is_confirm, create_by, create_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,CURRENT_TIMESTAMP())`
	resp, err := repo.db.Exec(sql, req.CompanyId, req.BranchId, req.UUID, req.QueId, req.DocType, req.DocDate, req.ArId, req.SaleId)//, req.TaxType, req.TaxRate, req.NumberOfItem, req.SumItemAmount, req.TotalAmount, req.NetAmount, req.Status, req.PickStatus, req.DeliveryStatus, req.ReceiveName, req.ReceiveTel, req.IsCancel, req.IsConfirm, req.CreateBy)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, _ := resp.LastInsertId()
	req.Id = id

	return map[string]interface{}{
		"status": "successfull",
		"uuid":   req.UUID,
		"que_id": req.QueId,
	}, nil
}

func CheckItemData(db *sqlx.DB, bar_code string) ItemData {
	item := ItemData{}

	sql := `select c.id as ItemId,c.code as ItemCode,c.item_name as ItemName,a.bar_code as BarCode,b.rate1 as Rate1, ifnull(d.sale_price_1,0) as ItemPrice,a.unit_code as UnitCode,ifnull(b.rate1,1)*c.average_cost as AverageCost from Barcode a inner join ItemRate b on a.item_code = b.item_code and a.unit_code = b.unit_code inner join Item c on a.item_code = c.code left join Price d on a.item_code = d.item_code and a.unit_code = d.unit_code where a.bar_code = ?`
	//sql := `select c.id as ItemId,c.code as ItemCode from Barcode a inner join ItemRate b on a.item_code = b.item_code and a.unit_code = b.unit_code inner join Item c on a.item_code = c.code left join Price d on a.item_code = d.item_code and a.unit_code = d.unit_code where a.bar_code = ? limit 1`
	err := db.Get(&item, sql, bar_code)
	if err != nil {
		fmt.Println("Error = ", err.Error())
	}

	return item
}

func (repo *pointofsaleRepository) ManageBasket(req *pointofsale.BasketTemplate) (interface{}, error) {
	var line_number int64

	now := time.Now()

	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	item := CheckItemData(repo.db, req.Sub[0].BarCode)
	req.Sub[0].Qty = req.Sub[0].PickQty
	req.Sub[0].RequestQty = req.Sub[0].PickQty
	req.Sub[0].PickAmount = req.Sub[0].PickQty * item.ItemPrice
	req.Sub[0].ItemCode = item.ItemCode
	req.Sub[0].LineNumber = 0
	req.Sub[0].PickBy = req.CreateBy
	req.DocDate = DocDate

	switch {
	case req.Id == 0:
		return nil, errors.New("คิวนี้ไม่ได้ระบุ ID ต้องส่ง ID มาด้วย")
	case req.UUID == "":
		return nil, errors.New("คิวนี้ไม่ได้ระบุ UUID ต้องส่ง UUID มาด้วย")
	case req.Sub[0].ItemCode == "":
		return nil, errors.New("ไม่ได้ระบุ รายการสินค้า")
	case req.Sub[0].PickQty == 0:
		return nil, errors.New("ไม่ได้ระบุ จำนวนที่ต้องการใส่ตะกร้า")
	case req.SaleId == 0:
		return nil, errors.New("ไม่ได้ระบุ ID พนักงานขาย")
	case req.ArId == 0:
		return nil, errors.New("ไม่ได้ระบุ ID ลูกค้า")
	}



	sql_sub := `insert into basket_sub(basket_id,uuid, que_id, doc_date, item_id, item_code, item_name, bar_code, checkount_qty, price, unit_id, checkout_amount, qty, remain_qty, rate1, ref_no, sale_id, average_cost, delivery_order_id, line_number, pick_by, pick_time) values(?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,CURRENT_TIMESTAMP())`
	_, err := repo.db.Exec(sql_sub, req.Id, req.UUID, req.QueId, req.DocDate, item.ItemId, item.ItemCode, item.ItemName, item.BarCode, req.Sub[0].PickQty, item.ItemPrice, item.UnitId, req.Sub[0].PickAmount, req.Sub[0].Qty, req.Sub[0].RemainQty, item.Rate1, req.Sub[0].RefNo, req.Sub[0].SaleId, item.AverageCost, req.Sub[0].DeliveryOrderId, req.Sub[0].LineNumber, req.Sub[0].PickBy)

	//sql_sub := `insert into basket_sub(basket_id,uuid, que_id, doc_date, item_id, item_code, item_name, bar_code, pick_qty, price, unit_id) values(?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`//, pick_amount, qty, remain_qty, rate_1, ref_no, sale_id, average_cost, delivery_order_id, line_number, request_by, request_time, pick_by, pick_time) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,CURRENT_TIMESTAMP())`
	//_, err := repo.db.Exec(sql_sub, req.Id, req.UUID, req.QueId, req.DocDate, item.ItemId, item.ItemCode, item.ItemName, item.BarCode, req.Sub[0].PickQty, item.ItemPrice, item.UnitId)//, req.Sub[0].PickAmount, req.Sub[0].Qty, req.Sub[0].RemainQty, item.Rate1, req.Sub[0].RefNo, req.Sub[0].SaleId, item.AverageCost, req.Sub[0].DeliveryOrderId, req.Sub[0].LineNumber, req.Sub[0].PickBy)

	if err != nil {
		return "Insert Quotation Not Success", err
	}

	line_number = line_number + 1

	return map[string]interface{}{
		"status": "successfull",
		"item":   item,
	}, nil
}
