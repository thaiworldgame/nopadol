package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mrtomyum/nopadol/delivery"

	//"google.golang.org/genproto/googleapis/type/date"
	"fmt"
	//"time"
	//"time"
)

func NewDeliveryRepository(db *sql.DB) delivery.Repository {
	return &deliveryRepository{db}
}

type deliveryRepository struct{ db *sql.DB }

func (d *deliveryRepository) ReportDaily(req string) (interface{}, error) {
	log.Println("begin postgres.ReportDaily")
	fmt.Println("param req -> ", req)
	type doModel struct {
		id          sql.NullInt64
		doDocno     sql.NullString `json:"do_docno"`
		soNo        sql.NullString `json:"so_no"`
		confirmDate sql.NullString `json:"confirm_date"`
		doDate      sql.NullString `json:"do_date"`
		diffDate    int64          `json:"diff_date"`
		description sql.NullString `json:"description"`
		arName      sql.NullString `json:"ar_name"`
		itemAmount  float64
		itemGroup   string
		ReMark      string `json:"remark"`
		InvoiceNo   string `json:"invoiceno"`
		Carlicense  string `json:"carlicense"`
		SaleOrder   string `json:"saleorder"`
		SaleMan     string `json:"saleman"`
	}

	type doResponse struct {
		Id          int64   `json:"id"`
		DoDocno     string  `json:"do_no"`
		SoNo        string  `json:"so_no"`
		ConfirmDate string  `json:"confirm_date"`
		DoDate      string  `json:"do_date"`
		DiffDate    int64   `json:"diff_date"`
		Description string  `json:"description"`
		ArName      string  `json:"ar_name"`
		ItemAmount  float64 `json:"item_amount"`
		ItemGroup   string  `json:"item_group"`
		ReMark      string  `json:"remark"`
		InvoiceNo   string  `json:"invoiceno"`
		Carlicense  string  `json:"carlicense"`
		SaleOrder   string  `json:"saleorder"`
		SaleMan     string  `json:"saleman"`
	}
	_do := doModel{}
	_dos := []doResponse{}
	_doResponse := doResponse{}

	lccommand := "select * from sm_do.doreport('" + req + "')"
	//lccommand = "select * from sm_do.doreport('2018-07-06')"

	fmt.Println(lccommand)
	rs, err := d.db.Query(lccommand)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rs.Next() {
		err := rs.Scan(&_do.id, &_do.doDocno, &_do.soNo, &_do.confirmDate, &_do.doDate, &_do.diffDate, &_do.description, &_do.arName, &_do.itemAmount, &_do.itemGroup, &_do.ReMark, &_do.InvoiceNo, &_do.Carlicense, &_do.SaleOrder, &_do.SaleMan)
		if err != nil {
			return nil, err
		}

		_doResponse.Id = _do.id.Int64
		_doResponse.DoDocno = _do.doDocno.String
		_doResponse.SoNo = _do.soNo.String

		//layout := "2006-01-02"
		//str := _do.confirmDate.String
		//t, err := time.Parse(layout, str)
		//
		//if err != nil {
		//	fmt.Println(err)
		//}
		////fmt.Println(t)
		//_doResponse.ConfirmDate = t
		//
		//str = _do.doDate.String
		//t, err = time.Parse(layout, str)
		_doResponse.ConfirmDate = _do.confirmDate.String
		_doResponse.DoDate = _do.doDate.String
		_doResponse.DiffDate = _do.diffDate
		_doResponse.Description = _do.description.String
		_doResponse.ArName = _do.arName.String
		_doResponse.ItemAmount = _do.itemAmount
		_doResponse.ItemGroup = _do.itemGroup
		_doResponse.ReMark = _do.ReMark
		_doResponse.InvoiceNo = _do.InvoiceNo
		_doResponse.Carlicense = _do.Carlicense
		_doResponse.SaleOrder = _do.SaleOrder

		_doResponse.SaleMan = _do.SaleMan

		_dos = append(_dos, _doResponse)
		fmt.Println(_do)
	}

	fmt.Println(_do)
	return map[string]interface{}{
		"data": _dos,
	}, nil

	//return nil, fmt.Errorf("error make mannual")
}
