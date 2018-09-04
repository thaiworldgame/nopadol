package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mrtomyum/nopadol/delivery"
	"log"
	//"google.golang.org/genproto/googleapis/type/date"
	"fmt"
	//"time"
	//"time"
	//"github.com/mrtomyum/nopadol/sale"
	//"github.com/satit13/bcsync/app"
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
		remark      string
		invoiceno   string
		carlicense  string
		salecode    string
		saleman     string
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
		Remark      string  `json:"remark"`
		Invoice     string  `json:"invoice"`
		CarLicense  string  `json:"car_license"`
		SaleCode    string  `json:"sale_code"`
		Saleman     string  `json:"saleman"`
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
		err := rs.Scan(&_do.id, &_do.doDocno, &_do.soNo, &_do.confirmDate,
			&_do.doDate, &_do.diffDate, &_do.description, &_do.arName,
			&_do.itemAmount, &_do.itemGroup, &_do.remark, &_do.invoiceno,
			&_do.carlicense, &_do.salecode, &_do.saleman)
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
		_doResponse.Remark = _do.remark
		_doResponse.Invoice = _do.invoiceno
		_doResponse.CarLicense = _do.carlicense
		_doResponse.SaleCode = _do.salecode
		_doResponse.Saleman = _do.saleman

		_dos = append(_dos, _doResponse)
		fmt.Println(_do)

	}

	fmt.Println(_do)
	return map[string]interface{}{
		"data": _dos,
	}, nil

	//return nil, fmt.Errorf("error make mannual")
}
func (d *deliveryRepository) GetTeam() (interface{}, error) {
	lcCommand := "select distinct branch from sm_do.tb_do_queue where COALESCE(upper(branch),'') <> ''"

	fmt.Printf("start postgres.getTeam() \n Query : %s \n ", lcCommand)
	rs, err := d.db.Query(lcCommand)

	type teamModel struct {
		ProfitCenter string `db:"branch" json:"profit_code"`
	}
	t := teamModel{}
	Ts := []teamModel{}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}
	for rs.Next() {
		rs.Scan(&t.ProfitCenter)
		Ts = append(Ts, t)
	}
	return Ts, nil
}

func (d *deliveryRepository) GetSale(_profit string) (interface{}, error) {
	fmt.Println("param _profit = ", _profit)
	lcCommand := "select distinct salecode,salename from sm_do.saleteam " +
		"where COALESCE(profitcenter,'') <> '' 	and activestatus=1 and salecode in " +
		"(select distinct salecode from sm_do.tb_do_queue	where COALESCE(upper(branch),'') = '" + _profit + "'" +
		"		and date_trunc('year',createdate)=date_trunc('year',now())  )order by salecode  "

	fmt.Printf("start postgres.getSale() \n Query : %s \n ", lcCommand)
	rs, err := d.db.Query(lcCommand)

	type saleModel struct {
		SaleCode string `db:"salecode" json:"sale_code"`
		SaleName string `db:"salename" json:"sale_name"`
	}
	s := saleModel{}
	Ss := []saleModel{}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}
	for rs.Next() {
		rs.Scan(&s.SaleCode, &s.SaleName)
		Ss = append(Ss, s)
	}
	return Ss, nil
}
