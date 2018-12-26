package mysqldb

import (
	"github.com/mrtomyum/nopadol/p9"
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
)

type BasketModel struct {
	Id                  int64            `db:"id" db:"id"`
	CompanyId           int64            `db:"company_id"`
	BranchId            int64            `db:"branch_id"`
	Uuid                string           `db:"uuid"`
	InvoiceNo           string           `db:"invoice_no"`
	TaxNo               string           `db:"tax_no"`
	QueId               int64            `db:"que_id"`
	DocType             int64            `db:"doc_type"`
	DocDate             string           `db:"doc_date"`
	ArId                int64            `db:"ar_id"`
	SaleId              int64            `db:"sale_id"`
	PosMachineId        int64            `db:"pos_machine_id"`
	PeriodId            int64            `db:"period_id"`
	CashId              int64            `db:"cash_id"`
	PosStatus           int64            `db:"pos_status"`
	TaxType             int64            `db:"tax_type"`
	TaxRate             int64            `db:"tax_rate"`
	NumberOfItem        int64            `db:"number_of_item"`
	ChangeAmount        float64          `db:"change_amount"`
	CashAmount          float64          `db:"cash_amount"`
	CreditCardAmount    float64          `db:"credit_card_amount"`
	ChqAmount           float64          `db:"chq_amount"`
	BankAmount          float64          `db:"bank_amount"`
	DepositAmount       float64          `db:"deposit_amount"`
	OnlineAmount        float64          `db:"online_amount"`
	CouponAmount        float64          `db:"coupon_amount"`
	CreditAmount        float64          `db:"credit_amount"`
	SumItemAmount       float64          `db:"sum_item_amount"`
	DiscountWord        string           `db:"discount_word"`
	DiscountAmount      float64          `db:"discount_amount"`
	AfterDiscountAmount float64          `db:"after_discount_amount"`
	BeforeTaxAmount     float64          `db:"before_tax_amount"`
	TaxAmount           float64          `db:"tax_amount"`
	TotalAmount         float64          `db:"total_amount"`
	NetAmount           float64          `db:"net_amount"`
	BillBalance         float64          `db:"bill_balance"`
	OtpPassword         string           `db:"otp_password"`
	Status              int64            `db:"status"`
	PickStatus          int64            `db:"pick_status"`
	HoldingStatus       int64            `db:"holding_status"`
	DeliveryStatus      int64            `db:"delivery_status"`
	ReceiveName         string           `db:"receive_name"`
	ReceiveTel          string           `db:"receive_tel"`
	IsPosted            int64            `db:"is_posted"`
	IsReturn            int64            `db:"is_return"`
	GLStatus            int64            `db:"gl_status"`
	ScgId               string           `db:"scg_id"`
	CreateBy            string           `db:"create_by"`
	CreateTime          string           `db:"create_time"`
	EditBy              string           `db:"edit_by"`
	EditTime            string           `db:"edit_time"`
	ConfirmBy           string           `db:"confirm_by"`
	ConfirmTime         string           `db:"confirm_time"`
	CancelBy            string           `db:"cancel_by"`
	CancelTime          string           `db:"cancel_time"`
	CancelDescId        int64            `db:"cancel_desc_id"`
	CancelDesc          string           `db:"cancel_desc"`
	Sub                 []BasketSubModel `db:"sub"`
}

type BasketSubModel struct {
	Id              int64   `db:"id"`
	PosId           int64   `db:"pos_id"`
	Uuid            string  `db:"uuid"`
	QueId           int64   `db:"que_id"`
	DocDate         string  `db:"doc_date"`
	ItemId          int64   `db:"item_id"`
	ItemCode        string  `db:"item_code"`
	ItemName        string  `db:"item_name"`
	BarCode         string  `db:"bar_code"`
	WhId            int64   `db:"wh_id"`
	ShelfId         int64   `db:"shelf_id"`
	RequestQty      float64 `db:"request_qty"`
	PickQty         float64 `db:"pick_qty"`
	CheckoutQty     float64 `db:"checkout_qty"`
	Price           float64 `db:"price"`
	UnitId          int64   `db:"unit_id"`
	PickAmount      float64 `db:"pick_amount"`
	CheckoutAmount  float64 `json:"checkout_amount"`
	Qty             float64 `db:"qty"`
	RemainQty       float64 `db:"remain_qty"`
	IsReturn        int64   `db:"is_return"`
	Rate1           int64   `db:"rate_1"`
	RefNo           string  `db:"ref_no"`
	SaleId          int64   `db:"sale_id"`
	AverageCost     float64 `db:"average_cost"`
	SumOfCost       float64 `db:"sum_of_cost"`
	DeliveryOrderId int64   `db:"delivery_order_id"`
	RefLineNumber   int64   `db:"ref_line_number"`
	LineNumber      int64   `db:"line_number"`
	RequestBy       string  `db:"request_by"`
	RequestTime     string  `db:"request_time"`
	PickBy          string  `db:"pick_by"`
	PickTime        string  `db:"pick_time"`
	CheckoutBy      string  `db:"checkout_by"`
	CheckoutTime    string  `db:"checkout_time"`
}

type Configuration struct {
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

type Machine struct {
	cash_id      int64  `db:"cash_id"`
	machine_id   int64  `db:"machine_id"`
	machine_code string `db:"machine_code"`
	machine_no   string `db:"machine_no"`
}

type p9Repository struct{ db *sqlx.DB }

func NewP9Repository(db *sqlx.DB) p9.Repository {
	return &p9Repository{db}
}

var que_id int64

func GenQueId(db *sqlx.DB) int64 {
	sql := `select ifnull(max(que_id),0)+1 as maxQue from sale_bill where doc_date = CURRENT_DATE()`
	err := db.Get(&que_id, sql)
	if err != nil {
		fmt.Println(err.Error())
	}

	return que_id
}

func GetBranchConfig(db *sqlx.DB) Configuration {
	config := Configuration{}
	sql := `select tax_type,tax_rate,ifnull(logo_path,'') as logo_path,depart_id,def_sale_wh_id,def_sale_shelf_id,def_buy_wh_id,def_buy_shelf_id,stock_status,sale_tax_type,buy_tax_type,sale_bill_type,buy_bill_type,use_address from configuration where company_id=1 and branch_id = 1`
	err := db.Get(&config, sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}

func GetMachine(db *sqlx.DB) Machine {
	m := Machine{}

	sql := `select a.cash_id,a.machine_id,b.machine_code,b.machine_no from pos_open a inner join pos_machine b on a.machine_id = b.id where company_id = 1 and branch_id=1`
	err := db.Get(m, sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	return m
}
//
//func GenUUID() string {
//	//u1 := uuid.Must(uuid.NewV4())
//	//fmt.Printf("UUIDv4: %s\n", u1)
//
//	// or error handling
//	uuid, err := uuid.NewV4()
//	if err != nil {
//		fmt.Printf("Something went wrong: %s", err)
//		return err.Error()
//	}
//	fmt.Printf("UUIDv4: %s\n", uuid)
//
//	// Parsing UUID from string input
//	//u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
//	//if err != nil {
//	//	fmt.Printf("Something went wrong: %s", err)
//	//	return err.Error()
//	//}
//	//fmt.Printf("Successfully parsed: %s", u2)
//
//	return uuid.String()
//}

func GetDataQue(db *sqlx.DB, uuid string) (BasketModel, error) {
	b := BasketModel{}

	sql := `select id,que_id,doc_date from sale_bill where uuid = ?`
	err := db.Get(&b, sql, uuid)
	if err != nil {
		return b, err
	}
	return b,nil
}

func (repo *p9Repository) Create(req *p9.BasketTemplate) (interface{}, error) {
	que_id := GenQueId(repo.db)
	uuid := GenUUID()

	fmt.Println("Que Id = ", que_id, uuid)
	req.Uuid = uuid
	req.QueId = que_id

	config := Configuration{}
	config = GetBranchConfig(repo.db)

	req.TaxType = config.TaxType
	req.TaxRate = config.TaxRate
	req.CompanyId = 1
	req.BranchId = 1
	req.DocType = 0
	req.CashId = 0

	fmt.Println("DocDate =", req.DocDate)

	now := time.Now()

	DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	if req.DocDate == "" {
		req.DocDate = DocDate
		fmt.Println("DocDate Default =", req.DocDate)
	}

	//machine := GetMachine(repo.db)

	//req.PosMachineId = machine.machine_id

	sql := `insert into sale_bill(company_id,branch_id,uuid,que_id,doc_date,ar_id,sale_id,tax_type,tax_rate,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP())`
	resp, err := repo.db.Exec(sql, req.CompanyId, req.BranchId, req.Uuid, req.QueId, req.DocDate, req.ArId, req.SaleId, req.TaxType, req.TaxRate, req.CreateBy)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, _ := resp.LastInsertId()
	req.Id = id

	var line_number int64

	for _, subs := range req.Sub {
		subs.LineNumber = line_number

		sql_sub := `insert into sale_bill_sub(pos_id,uuid,item_id,item_code,item_name,bar_code,wh_id,shelf_id,pick_qty,price,unit_id,pick_amount,qty,remain_qty,rate1,sale_id,sum_of_cost,line_number,pick_by,pick_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP())`
		_, err := repo.db.Exec(sql_sub, req.Id, req.Uuid, subs.ItemId, subs.ItemCode, subs.ItemName, subs.BarCode, subs.WhId, subs.ShelfId, subs.PickQty, subs.Price, subs.UnitId, subs.PickAmount, subs.Qty, subs.RemainQty, subs.Rate1, req.SaleId, subs.SumOfCost, subs.LineNumber, req.CreateBy)
		if err != nil {
			return "Insert Quotation Not Success", err
		}

		line_number = line_number + 1
	}

	return map[string]interface{}{
		"status": "successfull",
		"uuid":   req.Uuid,
		"que_id": req.QueId,
	}, nil
}

func (repo *p9Repository) ManageBasket(req *p9.BasketSubTemplate) (interface{}, error) {
	i := BasketModel{}
	i, err := GetDataQue(repo.db, req.Uuid)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Itemcode = ", req.ItemCode)
	req.QueId = i.QueId
	req.DocDate = i.DocDate
	req.PosId = i.Id
	sql_sub := `insert into sale_bill_sub(que_id,doc_date,pos_id,uuid,item_id,item_code,item_name,bar_code,wh_id,shelf_id,pick_qty,price,unit_id,pick_amount,qty,remain_qty,rate1,sale_id,sum_of_cost,line_number,pick_by,pick_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP())`
	_, err = repo.db.Exec(sql_sub, req.QueId, req.DocDate, req.PosId, req.Uuid, req.ItemId, req.ItemCode, req.ItemName, req.BarCode, req.WhId, req.ShelfId, req.PickQty, req.Price, req.UnitId, req.PickAmount, req.Qty, req.RemainQty, req.Rate1, req.SaleId, req.SumOfCost, req.LineNumber, req.PickBy)
	if err != nil {
		return "Insert Quotation Not Success", err
	}
	return map[string]interface{}{
		"status": "successfull",
		"uuid":   req.Uuid,
	}, nil
}
