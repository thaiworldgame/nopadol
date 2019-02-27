package mysqldb

import (
	"github.com/jmoiron/sqlx"
	//"fmt"
	"fmt"
	"github.com/mrtomyum/nopadol/sync"
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

type syncRepository struct{ db *sqlx.DB }

func NewSyncRepository(db *sqlx.DB) sync.Repository {
	return &syncRepository{db}
}

func (s *syncRepository) GetNewQoutaion() (resp interface{}, err error) {
	sync := syncLogs{}
	resp, err = sync.getWaitQuotation(s.db)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (sl *syncLogs) getWaitQuotation(db *sqlx.DB) (resp []BcQuotaionSend, err error) {
	sql := `select id,uuid,type,table_name,key_field,value from npdl.sync_logs
	where   send_status=0 and table_name='Quotation' `
	fmt.Println(sql)
	sync := syncLogs{}

	qt := BcQuotaionSend{}
	qts := []BcQuotaionSend{}
	//qtsub := []BCQuotationSub{}

	rs, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		err = rs.Scan(&sync.id, &sync.uuid, &sync.module_type, &sync.table_name,
			&sync.key_field, &sync.value)
		if err != nil {
			return nil, err
		}
		//fmt.Println(sync)
		qt.LogUuid = sync.uuid
		qt.DocNo = sync.value

		qt.get(db)
		qts = append(qts, qt)

		//syncs = append(syncs, sync)
	}
	//fmt.Print("sync object : ", syncs )
	return qts, nil

}

//func (sync *syncLogs) GetNewQuotation(db *sqlx.DB) (xs []BCQuotation, err error) {
//
//	syncList, err := sync.getWaitingQuotation(db, "Quotation")
//	if err != nil {
//		return nil, err
//	}
//
//	if len(syncList) == 0 {
//		return nil, nil
//	}
//
//	var y []BCQuotation
//	//var xx BCQuotation
//	if err != nil {
//		return nil, err
//	}
//
//	//todo : generate quotation /sub data
//	// todo : return quotation /sub
//
//	//for _, v := range syncList {
//	//	x := BCQuotation{}
//	//	sql := "select code,name1,defstkunitcode,stocktype from " + v. + " where " + v.keyField + "= '" + v.value + "'"
//	//	fmt.Println(sql)
//	//	rs := db.QueryRow(sql)
//	//	if err != nil {
//	//		fmt.Println("error query select ar")
//	//		return nil, err
//	//	}
//	//	rs.Scan(&x.Code, &x.Name, &x.DefStkUnitCode, &x.StockType)
//	//	xx.UID = v.Uid
//	//	xx.NewItem = x
//	//	y = append(y, xx)
//	//	//fmt.Println("xx :", x)
//	//}
//
//	return y, err
//}

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
