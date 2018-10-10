package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/customer"
	"fmt"
)

type CustomerModel struct {
	Id         int64  `db:"id"`
	Code       string `db:"code"`
	Name       string `db:"name"`
	Address    string `db:"address"`
	Telephone  string `db:"telephone"`
	BillCredit int64  `db:"bill_credit"`
}

type customerRepository struct{ db *sqlx.DB }

func NewCustomerRepository(db *sqlx.DB) customer.Repository {
	return &customerRepository{db}
}

func (cr *customerRepository) SearchById(req *customer.SearchByIdTemplate) (resp interface{}, err error) {
	cust := CustomerModel{}

	sql := `select id,code,name,ifnull(address,'') as address,ifnull(telephone,'') as telephone,bill_credit from Customer where id = ?`
	err = cr.db.Get(&cust, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	cust_resp := map_customer_template(cust)

	return map[string]interface{}{
		"id":        cust_resp.Id,
		"code":      cust_resp.Code,
		"name":      cust_resp.Name,
		"address":   cust_resp.Address,
		"telephone": cust_resp.Telephone,
	}, nil
}

func (cr *customerRepository) SearchByKeyword(req *customer.SearchByKeywordTemplate) (resp interface{}, err error) { //(doc_no like CONCAT("%",?,"%"))
	custs := []CustomerModel{}

	sql := `select id,code,name,ifnull(address,'') as address,ifnull(telephone,'') as telephone,bill_credit from Customer where (code like concat('%',?,'%') or name like concat('%',?,'%') or address like concat('%',?,'%')) order by code limit 20`
	err = cr.db.Select(&custs, sql, req.Keyword, req.Keyword, req.Keyword)
	fmt.Println(sql)
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
		Id:        x.Id,
		Code:      x.Code,
		Name:      x.Name,
		Address:   x.Address,
		Telephone: x.Telephone,
		BillCredit:x.BillCredit,
	}
}
