package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/dataprint"
	"fmt"
)

type PrintRequest struct {
	Module string `json:"module" db:"module"`
}

type DataPrint struct {
	Id           int    `json:"id" db:"id"`
	Module       string `json:"module" db:"module"`
	Doctype      string `json:"doctype" db:"doc_type"`
	FormType     string `json:"form_type" db:"form_type"`
	InsertedTime string `json:"inserted_time" db:"inserted_time"`
	IsPrinted    int    `json:"is_printed" db:"is_printed"`
	Data         string `json:"data" db:"data"`
}

type dataprintRepository struct{ db *sqlx.DB }

func NewDataPrintRepository(db *sqlx.DB) dataprint.Repository {
	return &dataprintRepository{db}
}

func (s *dataprintRepository) DataPrint() (resp interface{}, err error) {
	data_print := DataPrint{}
	resp, err = data_print.getWaitDataPrint(s.db)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (sl *DataPrint) getWaitDataPrint(db *sqlx.DB) (resp []DataPrint, err error) {
	sql := `select id,module,doc_type,form_type,inserted_time,is_printed,data from npdl.print_queue where is_printed = 0  order by inserted_time`
	fmt.Println(sql)
	//sync := syncLogs{}
	//
	//qt := DataPrint{}
	//qts := []DataPrint{}

	rs, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		err = rs.Scan(&sl.Id, &sl.Module, &sl.Doctype, &sl.FormType, &sl.InsertedTime, &sl.IsPrinted,&sl.Data)
		if err != nil {
			return nil, err
		}

		resp = nil//append(sl, qt)
	}
	return resp, nil
}

//func (sl *syncLogs) getWaitSaleOrder(db *sqlx.DB) (resp []BcSaleOrderSend, err error) {
//	sql := `select id,uuid,type,table_name,key_field,value from npdl.sync_logs where send_status=0 and table_name='SaleOrder' `
//	fmt.Println(sql)
//	sync := syncLogs{}
//
//	so := BcSaleOrderSend{}
//	sos := []BcSaleOrderSend{}
//
//	rs, err := db.Queryx(sql)
//	if err != nil {
//		return nil, err
//	}
//
//	for rs.Next() {
//		err = rs.Scan(&sync.id, &sync.uuid, &sync.module_type, &sync.table_name, &sync.key_field, &sync.value)
//		if err != nil {
//			return nil, err
//		}
//
//		so.LogUuid = sync.uuid
//		so.DocNo = sync.value
//
//		so.getSOByDocNo(db)
//		sos = append(sos, so)
//	}
//	return sos, nil
//}

//func (d DataPrint) DataPrint() (interface{}, error) {
//	sql := `select a.Id,a.CompanyId,a.BranchId,a.DocNo,a.DocDate,a.DocType,a.Validity,a.BillType,a.ArId,a.ArCode,a.ArName,a.SaleId,a.SaleCode,a.SaleName,ifnull(a.DepartId,0) as DepartId,ifnull(c.code,'') as DepartCode,ifnull(a.RefNo,'') as RefNo,ifnull(a.JobId,'') as JobId,a.TaxType,a.IsConfirm,a.BillStatus,a.CreditDay,ifnull(a.DueDate,'') as DueDate,a.ExpireCredit,ifnull(a.ExpireDate,'') as ExpireDate,a.DeliveryDay,ifnull(a.DeliveryDate,'') as DeliveryDate,a.AssertStatus,a.IsConditionSend,ifnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,ifnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.AfterDiscountAmount,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.NetDebtAmount,a.TaxRate,a.ProjectId,ifnull(d.code,'') as ProjectCode,a.AllocateId,ifnull(e.code,'') as AllocateCode,a.IsCancel,ifnull(a.CreateBy,'') as CreateBy,ifnull(a.CreateTime,'') as CreateTime,ifnull(a.EditBy,'') as EditBy,ifnull(a.EditTime,'') as EditTime,ifnull(a.CancelBy,'') as CancelBy,ifnull(a.CancelTime,'') as CancelTime,ifnull(b.address,'') as ArBillAddress,ifnull(b.telephone,'') as ArTelephone from Quotation a left join Customer b on a.ArId = b.id  left join Department c on a.DepartId = c.Id left join Project d on a.ProjectId = d.Id left join Allocate e on a.AllocateId = e.id   where a.DocNo = ?`
//	rs := db.QueryRow(sql, q.DocNo)
//	err := rs.Scan(&q.Id, &q.CompanyId, &q.BranchId, &q.DocNo, &q.DocDate, &q.DocType, &q.Validity, &q.BillType, &q.ArId, &q.ArCode, &q.ArName, &q.SaleId, &q.SaleCode, &q.SaleName, &q.DepartId, &q.DepartCode, &q.RefNo, &q.JobId, &q.TaxType, &q.IsConfirm, &q.BillStatus, &q.CreditDay, &q.DueDate, &q.ExpireCredit, &q.ExpireDate, &q.DeliveryDay, &q.DeliveryDate, &q.AssertStatus, &q.IsConditionSend, &q.MyDescription, &q.SumOfItemAmount, &q.DiscountWord, &q.DiscountAmount, &q.AfterDiscountAmount, &q.BeforeTaxAmount, &q.TaxAmount, &q.TotalAmount, &q.NetDebtAmount, &q.TaxRate, &q.ProjectId, &q.ProjectCode, &q.AllocateId, &q.AllocateCode, &q.IsCancel, &q.CreateBy, &q.CreateTime, &q.EditBy, &q.EditTime, &q.CancelBy, &q.CancelTime, &q.ArBillAddress, &q.ArTelephone)
//	if err != nil {
//		fmt.Println("err = ", err.Error())
//		return err
//	}
//
//	jstr := `{"key": "userKeyValue", "value": "userValueValue"}`
//
//	// declare a map that has a key string and value interface{} so that any values or
//	// types will be accepted;
//	jmap := make(map[string]interface{})
//
//	err := json.Unmarshal(jstr, &jmap)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for k, v := range jmap {
//		fmt.Printf("Key: %v, Value: %v\n", k, v)
//	}
//
//	return nil, nil
//}
