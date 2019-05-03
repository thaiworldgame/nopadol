package mysqldb

import (
	"github.com/jmoiron/sqlx"
	//"fmt"
	"fmt"
	"github.com/mrtomyum/nopadol/sync"
	//"encoding/json"
	//"bytes"
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
	id          int64  `db:"id"`
	uuid        string `db:"uid"`
	module_type string `db:"type"`
	table_name  string `db:"tablename"`
	action      string `db:"action"`
	key_field   string `db:"key_field"`
	value       string `db:"value"`
}

type BcQuotaionSend struct {
	BCQuotation
	LogUuid string `json:"log_uuid"`
}

type BcSaleOrderSend struct {
	BCSaleOrder
	LogUuid string `json:"log_uuid"`
}

type syncRepository struct{ db *sqlx.DB }

func NewSyncRepository(db *sqlx.DB) sync.Repository {
	return &syncRepository{db}
}

func (s *syncRepository) GetNewQuotaion() (resp interface{}, err error) {
	sync := syncLogs{}
	resp, err = sync.getWaitQuotation(s.db)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (sl *syncLogs) getWaitQuotation(db *sqlx.DB) (resp []BcQuotaionSend, err error) {
	sql := `select id,uuid,type,table_name,key_field,value from npdl.sync_logs where send_status=0 and table_name='Quotation' `
	fmt.Println(sql)
	sync := syncLogs{}

	qt := BcQuotaionSend{}
	qts := []BcQuotaionSend{}

	rs, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		err = rs.Scan(&sync.id, &sync.uuid, &sync.module_type, &sync.table_name, &sync.key_field, &sync.value)
		if err != nil {
			return nil, err
		}

		qt.LogUuid = sync.uuid
		qt.DocNo = sync.value

		qt.getQTByDocNo(db)
		qts = append(qts, qt)
	}
	return qts, nil
}

func (sl *syncLogs) getWaitSaleOrder(db *sqlx.DB) (resp []BcSaleOrderSend, err error) {
	sql := `select id,uuid,type,table_name,key_field,value from npdl.sync_logs where send_status=0 and table_name='SaleOrder' `
	fmt.Println(sql)
	sync := syncLogs{}

	so := BcSaleOrderSend{}
	sos := []BcSaleOrderSend{}

	rs, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		err = rs.Scan(&sync.id, &sync.uuid, &sync.module_type, &sync.table_name, &sync.key_field, &sync.value)
		if err != nil {
			return nil, err
		}

		so.LogUuid = sync.uuid
		so.DocNo = sync.value

		so.getSOByDocNo(db)
		sos = append(sos, so)
	}
	return sos, nil
}

func (s *syncRepository) GetNewSaleOrder() (resp interface{}, err error) {
	sync := syncLogs{}
	resp, err = sync.getWaitSaleOrder(s.db)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *syncRepository) ConfirmTransfer(req *sync.Logs) (status interface{}, err error) {
	fmt.Println("sync.Logs = ", req)

	for _, l := range req.LogsUUID {
		fmt.Println("l=", l.LogUUID)
		sql := `update npdl.sync_logs set send_status = 1 where uuid = ?`
		_, err = s.db.Exec(sql, l.LogUUID)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return map[string]interface{}{
		"status": "success",
	}, nil
}

