package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mrtomyum/nopadol/delivery"
	"log"
	//"google.golang.org/genproto/googleapis/type/date"
	"fmt"
	//"time"
	"time"
)

func NewDeliveryRepository(db *sql.DB) delivery.Repository {
	return &deliveryRepository{db}
}

type deliveryRepository struct{ db *sql.DB }

func (d *deliveryRepository) ReportDaily() (interface{}, error) {
	log.Println("begin postgres.ReportDaily")
	type doModel struct {
		Id          sql.NullInt64
		DoDocno     sql.NullString `json:"do_docno"`
		SoNo        sql.NullString `json:"so_no"`
		ConfirmDate sql.NullString `json:"confirm_date"`
		DoDate      sql.NullString `json:"do_date"`
		DiffDate    int64          `json:"diff_date"`
		Description sql.NullString `json:"description"`
		ArName      sql.NullString `json:"ar_name"`
	}

	type doResponse struct {
		Id          int64
		DoDocno     string `json:"do_docno"`
		SoNo        string `json:"so_no"`
		ConfirmDate time.Time `json:"confirm_date"`
		DoDate      time.Time `json:"do_date"`
		DiffDate    int64          `json:"diff_date"`
		Description string `json:"description"`
		ArName      string `json:"ar_name"`
	}
	_do := doModel{}
	_dos := []doResponse{}
	_doResponse := doResponse{}

	lccommand := `select * from sm_do.doreport('2018-07-05')`

	rs, err := d.db.Query(lccommand)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rs.Next() {
		err := rs.Scan(&_do.Id, &_do.DoDocno, &_do.SoNo, &_do.ConfirmDate, &_do.DoDate, &_do.DiffDate, &_do.Description, &_do.ArName)
		if err != nil {
			return nil, err
		}

		_doResponse.Id = _do.Id.Int64
		_doResponse.DoDocno = _do.DoDocno.String
		_doResponse.SoNo = _do.SoNo.String

		layout := "2006-01-02T15:04:05.000Z"
		str := _do.ConfirmDate.String
		t, err := time.Parse(layout, str)

		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(t)
		_doResponse.ConfirmDate = t

		str = _do.DoDate.String
		t, err = time.Parse(layout, str)
		_doResponse.DoDate = t
		_doResponse.DiffDate = _do.DiffDate
		_doResponse.Description = _do.Description.String
		_doResponse.ArName = _do.ArName.String
		_dos = append(_dos, _doResponse)
		fmt.Println(_doResponse)
	}

	fmt.Println(_do)
	return map[string]interface{}{
		"data":_dos,
	}, nil

	//return nil, fmt.Errorf("error make mannual")
}
