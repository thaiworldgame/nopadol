package mysqldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/customer"
)

type CustomerModel struct {
	Id         int64  `db:"id"`
	Code       string `db:"code"`
	Name       string `db:"name"`
	Address    string `db:"address"`
	Telephone  string `db:"telephone"`
	BillCredit int64  `db:"bill_credit"`
	Email      string `json:"email"`
	CompanyID  int    `json:"company_id"`
	CreateBy string `json:"create_by"`

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
		"id":          cust_resp.Id,
		"code":        cust_resp.Code,
		"name":        cust_resp.Name,
		"address":     cust_resp.Address,
		"telephone":   cust_resp.Telephone,
		"bill_credit": cust_resp.BillCredit,
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
		Id:         x.Id,
		Code:       x.Code,
		Name:       x.Name,
		Address:    x.Address,
		Telephone:  x.Telephone,
		BillCredit: x.BillCredit,
	}
}

func (cust *CustomerModel) Search(db *sqlx.DB, ar_code string) {
	sql := `select id,code,name,ifnull(address,'') as address,ifnull(telephone,'') as telephone,bill_credit from Customer where code = ?`
	rs := db.QueryRow(sql, ar_code)
	rs.Scan(&cust.Id, &cust.Code, &cust.Name, &cust.Address, &cust.Telephone, &cust.BillCredit)

	return
}

func (cr *customerRepository) StoreCustomer(req *customer.CustomerTemplate) (interface{}, error) {
	cus := CustomerModel{
		Id:         req.Id,
		Code:       req.Code,
		Name:       req.Name,
		Address:    req.Address,
		Telephone:  req.Telephone,
		BillCredit: req.BillCredit,
		Email:      req.Email,
		CompanyID:  req.CompanyID,
	}

	return cus.save(cr.db)
	//return nil,nil
}

func (c CustomerModel) save(db *sqlx.DB) (interface{}, error) {
	// check id = 0 to new customer
	// check code , name , telephone exists
	// update
	// insert
	sql := `insert into customer (code,name,address,telephone,bill_credit,active_status,create_by)
		valuse (?,?,?,?,?,?,?)`
	rs,err := db.Exec(sql,c.Code,c.Name,c.Address,c.Telephone,c.BillCredit,0,c.CreateBy)
	if err != nil {
		return nil,err
	}
	newID ,err := rs.LastInsertId()
	if err != nil {
		return nil,err
	}
	return newID, nil
}
