package postgres

import (
	"github.com/mrtomyum/nopadol/delivery"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	//"google.golang.org/genproto/googleapis/type/date"
	"fmt"
)

func NewDeliveryRepository(db *sql.DB) delivery.Repository {
	return &deliveryRepository{db}
}

type deliveryRepository struct{ db *sql.DB }

func (d *deliveryRepository)ReportDaily() (interface{}, error) {
	log.Println("begin postgres.ReportDaily")
	type doModel struct {
		Docno    string
		Docdate  sql.NullString `db:"docdate"`
		Datesend sql.NullString `db:"datesend"`
		Timesend sql.NullString `db:"timesend"`
	}

	_do := doModel{}

	lccommand := "select docno,docdate,datesend,timesend from sm_do.tb_do_delivery limit 1 "
	rs := d.db.QueryRow(lccommand)
	//_docno := ""
	err := rs.Scan(&_do.Docno, &_do.Docdate, &_do.Datesend, &_do.Timesend)
	//err := rs.Scan(&_do.docno)
	if err != nil {
		return nil, err

	}
	fmt.Println(_do)
	return map[string]interface{}{
		"doc_no" : _do.Docno,
		"doc_date" : _do.Docdate.String,
		"date_send": _do.Datesend.String,
		"time_send": _do.Timesend.String,
	}, nil
	//return nil, fmt.Errorf("error make mannual")
}
