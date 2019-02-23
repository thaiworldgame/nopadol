package mysqldb

import (
	"github.com/jmoiron/sqlx"
	//"fmt"
)

type BcItem struct {
	Code           string `json:"code" db:"code"`
	Name           string `json:"name" db:"name1" db:"name1"`
	DefStkUnitCode string `json:"def_stk_unit_code" db:"defstkunitcode"`
	StockType      int    `json:"stock_type" db:"stocktype"`
}

type BcItemSend struct {
	NewItem BcItem
	UID     string
}

type syncLogs struct {
	id         int64  `db:"id"`
	uuid        string `db:"uid"`
	module_type string `db:"type"`
	table_name  string `db:"tablename"`
	action     string `db:"action"`
	key_field   string `db:"key_field"`
	value      string `db:"value"`
}

type BCQuotation struct {
	Id                  int64            `db:"Id"`
	UUID                string           `db:"UUID"`
	DocNo               string           `db:"DocNo"`
	DocDate             string           `db:"DocDate"`
	CompanyId           int64            `db:"CompanyId"`
	BranchId            int64            `db:"BranchId"`
	DocType             int64            `db:"DocType"`
	ArId                int64            `db:"ArId"`
	ArCode              string           `db:"ArCode"`
	ArName              string           `db:"ArName"`
	ArBillAddress       string           `db:"ArBillAddress"`
	ArTelephone         string           `db:"ArTelephone"`
	SaleId              int64            `db:"SaleId"`
	SaleCode            string           `db:"SaleCode"`
	SaleName            string           `db:"SaleName"`
	BillType            int64            `db:"BillType"`
	TaxType             int64            `db:"TaxType"`
	TaxRate             float64          `db:"TaxRate"`
	DepartId            int64            `db:"DepartId"`
	RefNo               string           `db:"RefNo"`
	JobId               string           `db:"JobId"`
	IsConfirm           int64            `db:"IsConfirm"`
	BillStatus          int64            `db:"BillStatus"`
	Validity            int64            `db:"Validity"`
	CreditDay           int64            `db:"CreditDay"`
	DueDate             string           `db:"DueDate"`
	ExpireCredit        int64            `db:"ExpireCredit"`
	ExpireDate          string           `db:"ExpireDate"`
	DeliveryDay         int64            `db:"DeliveryDay"`
	DeliveryDate        string           `db:"DeliveryDate"`
	AssertStatus        int64            `db:"AssertStatus"`
	IsConditionSend     int64            `db:"IsConditionSend"`
	MyDescription       string           `db:"MyDescription"`
	SumOfItemAmount     float64          `db:"SumOfItemAmount"`
	DiscountWord        string           `db:"DiscountWord"`
	DiscountAmount      float64          `db:"DiscountAmount"`
	AfterDiscountAmount float64          `db:"AfterDiscountAmount"`
	BeforeTaxAmount     float64          `db:"BeforeTaxAmount"`
	TaxAmount           float64          `db:"TaxAmount"`
	TotalAmount         float64          `db:"TotalAmount"`
	NetDebtAmount       float64          `db:"NetDebtAmount"`
	ProjectId           int64            `db:"ProjectId"`
	AllocateId          int64            `db:"AllocateId"`
	IsCancel            int64            `db:"IsCancel"`
	CreateBy            string           `db:"CreateBy"`
	CreateTime          string           `db:"CreateTime"`
	EditBy              string           `db:"EditBy"`
	EditTime            string           `db:"EditTime"`
	ConfirmBy           string           `db:"ConfirmBy"`
	ConfirmTime         string           `db:"ConfirmTime"`
	CancelBy            string           `db:"CancelBy"`
	CancelTime          string           `db:"CancelTime"`
	Subs                []BCQuotationSub `db:"subs"`
}

type BCQuotationSub struct {
	Id              int64   `db:"Id"`
	QuoUUID         string  `db:"QuoUUID"`
	QuoId           int64   `db:"QuoId"`
	ItemId          int64   `db:"ItemId"`
	ItemCode        string  `db:"ItemCode"`
	BarCode         string  `db:"BarCode"`
	ItemName        string  `db:"ItemName"`
	Qty             float64 `db:"Qty"`
	RemainQty       float64 `db:"RemainQty"`
	Price           float64 `db:"Price"`
	DiscountWord    string  `db:"DiscountWord"`
	DiscountAmount  float64 `db:"DiscountAmount"`
	UnitCode        string  `db:"UnitCode"`
	ItemAmount      float64 `db:"ItemAmount"`
	ItemDescription string  `db:"ItemDescription"`
	PackingRate1    float64 `db:"PackingRate1"`
	IsCancel        int64   `db:"IsCancel"`
	LineNumber      int     `db:"LineNumber"`
}

func (ar *BCQuotation) GetNewQuotation(db *sqlx.DB) (xs []BCQuotation, err error) {

	sync := syncLogs{}
	syncList, err := sync.getWaitingQuotation(db, "Quotation")
	if err != nil {
		return nil, err
	}

	if len(syncList) == 0 {
		return nil, nil
	}

	var y []BCQuotation
	//var xx BCQuotation
	if err != nil {
		return nil, err
	}

	//for _, v := range syncList {
	//	x := BCQuotation{}
	//	sql := "select code,name1,defstkunitcode,stocktype from " + v. + " where " + v.keyField + "= '" + v.value + "'"
	//	fmt.Println(sql)
	//	rs := db.QueryRow(sql)
	//	if err != nil {
	//		fmt.Println("error query select ar")
	//		return nil, err
	//	}
	//	rs.Scan(&x.Code, &x.Name, &x.DefStkUnitCode, &x.StockType)
	//	xx.UID = v.Uid
	//	xx.NewItem = x
	//	y = append(y, xx)
	//	//fmt.Println("xx :", x)
	//}

	return y, err
}

func (syn *syncLogs) getWaitingQuotation(db *sqlx.DB, moduleType string) ([]syncLogs, error) {
	sql := `select id,convert(nvarchar(50),uuid) as uuid,type,table_name,action,key_field,value
	 	from sync_logs
	 	where   send_status=0
	 		and type='item' `
	sync := syncLogs{}
	syncs := []syncLogs{}
	rs, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		err = rs.Scan(&sync.id, &sync.uuid, &sync.module_type, &sync.table_name, &sync.action, &sync.key_field, &sync.value)
		if err != nil {
			return nil, err
		}
		//fmt.Println(sync)
		syncs = append(syncs, sync)
	}
	//fmt.Print("sync object : ", syncs )
	return syncs, nil
}

//func (syn *syncLogs) getWaitingItem(db *sqlx.DB, moduleType string) ([]syncLogs, error) {
//	sql := `select id,convert(nvarchar(50),uid) as uid,type,tablename,roworder,action,key_field,value
//	 	from sync_logs
//	 	where   send_status=0
//	 		and type='item' `
//	sync := syncLogs{}
//	syncs := []syncLogs{}
//	rs, err := db.Queryx(sql)
//	if err != nil {
//		return nil, err
//	}
//
//	for rs.Next() {
//		err = rs.Scan(&sync.Id, &sync.Uid, &sync.ModuleType, &sync.TableName, &sync.roworder, &sync.action, &sync.keyField, &sync.value)
//		if err != nil {
//			return nil, err
//		}
//		//fmt.Println(sync)
//		syncs = append(syncs, sync)
//	}
//	//fmt.Print("sync object : ", syncs )
//	return syncs, nil
//}
