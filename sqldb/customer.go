package sqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/customer"
	"fmt"
)

type CustomerModel struct {
	Id     int64  `db:"Id"`
	ArCode string `db:"ArCode"`
	ArName string `db:"ArName"`
}

type customerRepository struct{ db *sqlx.DB }

func NewCustomerRepository(db *sqlx.DB) customer.Repository {
	return &customerRepository{db}
}

func (cr *customerRepository) SearchById(req *customer.SearchByIdTemplate) (resp interface{}, err error) {
	cust := CustomerModel{}

	sql := `select roworder as Id,code as ArCode,name1 as ArName from dbo.bcar where roworder = ?`
	err = cr.db.Get(&cust, sql, req.Id)
	if err != nil {
		fmt.Println("err = ",err.Error())
		return resp, fmt.Errorf(err.Error())
		//return resp, errors.New(err.Error())
	}

	//fmt.Println("customer = ", cust)

	cust_resp := map_customer_template(cust)

	//fmt.Println("customer Resp= ", cust_resp)

	return map[string]interface{}{
		"customer_id" : cust_resp.CustomerId,
		"customer_code" : cust_resp.CustomerCode,
		"customer_name" : cust_resp.CustomerName,
	}, nil
}


func map_customer_template(x CustomerModel) customer.CustomerTemplate {
	return customer.CustomerTemplate{
		CustomerId:   x.Id,
		CustomerCode: x.ArCode,
		CustomerName:  x.ArName,
	}
}