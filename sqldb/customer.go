package sqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/customer"
	"fmt"
)

type CustomerModel struct {
	Id                int64  `db:"Id"`
	ArCode            string `db:"ArCode"`
	ArName            string `db:"ArName"`
	CustomerAddress   string `db:"CustomerAddress"`
	CustomerTelephone string `db:"CustomerTelephone"`
}

type customerRepository struct{ db *sqlx.DB }

func NewCustomerRepository(db *sqlx.DB) customer.Repository {
	return &customerRepository{db}
}

func (cr *customerRepository) SearchById(req *customer.SearchByIdTemplate) (resp interface{}, err error) {
	cust := CustomerModel{}

	sql := `select roworder as Id,code as ArCode,name1 as ArName,isnull(billaddress,'') as CustomerAddress,isnull(telephone,'') as CustomerTelephone from dbo.bcar where roworder = ?`
	err = cr.db.Get(&cust, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	cust_resp := map_customer_template(cust)

	return map[string]interface{}{
		"customer_id":        cust_resp.CustomerId,
		"customer_code":      cust_resp.CustomerCode,
		"customer_name":      cust_resp.CustomerName,
		"customer_address":   cust_resp.CustomerAddress,
		"customer_telephone": cust_resp.CustomerTelephone,
	}, nil
}

func (cr *customerRepository) SearchByKeyword(req *customer.SearchByKeywordTemplate) (resp interface{}, err error) {
	custs := []CustomerModel{}

	sql := `select roworder as Id,code as ArCode,name1 as ArName,isnull(billaddress,'') as CustomerAddress,isnull(telephone,'') as CustomerTelephone from dbo.bcar where (code like '%'+?+'%' or name1 like '%'+?+'%' or billaddress like '%'+?+'%')`
	err = cr.db.Select(&custs, sql, req.Keyword, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	customer := []customer.CustomerTemplate{}

	for _, c := range custs {

		custline := map_customer_template(c)
		customer = append(customer, custline)
	}

	fmt.Println("List customer =", custs)

	return customer, nil
}

func map_customer_template(x CustomerModel) customer.CustomerTemplate {
	return customer.CustomerTemplate{
		CustomerId:        x.Id,
		CustomerCode:      x.ArCode,
		CustomerName:      x.ArName,
		CustomerAddress:   x.CustomerAddress,
		CustomerTelephone: x.CustomerTelephone,
	}
}
