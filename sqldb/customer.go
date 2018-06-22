package sqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/customer"
	"context"
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

func (cr *customerRepository) SearchCustomerById(ctx context.Context, req *customer.SearchById) (resp customer.CustomerTemplate, err error) {
	cust := CustomerModel{}
	sql := `select roworder as Id,code as ArCode,name1 as ArName from dbo.bcar where code = '45040'`
	err = cr.db.Get(&cust, sql)
	if err != nil {
		fmt.Println("err = ",err.Error())
		return resp, err
	}

	fmt.Println("customer = ", cust)

	Resp := map_customer_template(cust)

	return Resp, nil
}


func map_customer_template(x CustomerModel) customer.CustomerTemplate {
	return customer.CustomerTemplate{
		CustomerId:   x.Id,
		CustomerCode: x.ArCode,
		CustomerName:  x.ArName,
	}
}