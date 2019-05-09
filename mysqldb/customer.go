package mysqldb

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/customer"
	"github.com/mrtomyum/nopadol/mysqldb/n9model"
)

type CustomerModel struct {
	Id           int64     `db:"id"`
	Code         string    `db:"code"`
	Name         string    `db:"name"`
	Address      string    `db:"address"`
	Telephone    string    `db:"telephone"`
	BillCredit   float64   `db:"bill_credit"`
	DebtAmount   float64   `db:"debt_amount"`
	DebtLimit    float64   `db:"debt_limit"`
	Email        string    `db:"email"`
	CompanyID    int       `db:"company_id"`
	CreateBy     string    `db:"create_by"`
	CreateTime   time.Time `db:"create_time"`
	UpdateBy     string    `db:"update_by"`
	UpdateTime   time.Time `db:"update_time"`
	Fax          string    `db:"fax"`
	TaxNo        string    `db:"tax_no"`
	MemberID     string    `db:"member_id"`
	PointBalance float64   `db:"point_balance"`
}

type customerRepository struct{ db *sqlx.DB }

func NewCustomerRepository(db *sqlx.DB) customer.Repository {
	return &customerRepository{db}
}

func (cr *customerRepository) SearchById(req *customer.SearchByIdTemplate) (resp interface{}, err error) {
	//cust := CustomerModel{}

	//request id must <>0
	if req.Id ==0 {
		return nil,fmt.Errorf("empty Id ")
	}
	n9customer := n9model.Customer{Id: req.Id}
	err = n9customer.GetById(cr.db, req.Id)
	if err != nil {
		return nil, err
	}

	//sql := `select id,code,name,ifnull(address,'') as address,ifnull(telephone,'') as telephone,bill_credit from Customer where id = ?`
	//err = cr.db.Get(&cust, sql, req.Id)
	//if err != nil {
	//	fmt.Println("err = ", err.Error())
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cust_resp := map_customer_template(cust)
	cust_resp := map_from_n9Model_to_customer_template(n9customer)
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

func map_from_n9Model_to_customer_template(x n9model.Customer) customer.CustomerTemplate {
	return customer.CustomerTemplate{
		Id:           x.Id,
		Code:         x.Code,
		Name:         x.Name,
		Address:      x.Address,
		Telephone:    x.Telephone,
		BillCredit:   x.BillCredit,
		CreditAmount: x.BillCredit,
		CompanyID:    x.CompanyID,
		CreateBy:     x.CreateBy,
		MemberID:     x.MemberID,
		PointBalance: x.PointBalance,
		Fax:          x.Fax,
		DebtLimit:    x.DebtLimit,
		DebtAmount:   x.DebtAmount,
	}
}

func (cust *CustomerModel) Search(db *sqlx.DB, ar_code string) {
	sql := `select id,code,name,ifnull(address,'') as address,ifnull(telephone,'') as telephone,bill_credit from Customer where code = ?`
	rs := db.QueryRow(sql, ar_code)
	rs.Scan(&cust.Id, &cust.Code, &cust.Name, &cust.Address, &cust.Telephone, &cust.BillCredit)

	return
}

func (cr *customerRepository) StoreCustomer(req *customer.CustomerTemplate) (res interface{}, err error) {
	cus := CustomerModel{
		Id:           req.Id,
		Code:         req.Code,
		Name:         req.Name,
		Address:      req.Address,
		Telephone:    req.Telephone,
		Fax:          req.Fax,
		TaxNo:        req.TaxNo,
		BillCredit:   req.BillCredit,
		DebtAmount:   req.DebtAmount,
		DebtLimit:    req.DebtLimit,
		MemberID:     req.MemberID,
		PointBalance: req.PointBalance,
		Email:        req.Email,
		CompanyID:    req.CompanyID,
		CreateBy:     req.CreateBy,
		CreateTime:   time.Now(),
		UpdateBy:     req.CreateBy,
		UpdateTime:   time.Now(),
	}
	// check case insert & update  (0,1)
	var id int64

	cus.Id = id
	return cus.save(cr.db)
	//return nil,nil
}

func (c *CustomerModel) getIdByCode(db *sqlx.DB, code string) (int64, error) {
	sql := "select id from Customer where code='" + code + "'"
	fmt.Println(sql)
	var curID int64
	db.QueryRow(sql).Scan(&curID)
	fmt.Println("current id -> ", curID)
	return curID, nil
}

func (c *CustomerModel) save(db *sqlx.DB) (interface{}, error) {
	var curID int64
	fmt.Println("start customer.save ,  ", c)

	// return -1 if not exists by code
	curID, err := c.getIdByCode(db, c.Code)
	if err != nil {
		return nil, err
	}
	//validate id if empty -> insert
	switch {

	//new customer case
	case curID == 0:
		{
			sql := `insert into Customer (code,name,address,telephone,bill_credit,
						active_status,create_by,create_time)
		values (?,?,?,?,?,?,?,?)`
			fmt.Println(sql)
			rs, err := db.Exec(sql, c.Code, c.Name, c.Address, c.Telephone, c.BillCredit, 0, c.CreateBy, c.CreateTime)

			if err != nil {
				return nil, err
			}
			_, err = rs.LastInsertId()
			if err != nil {
				return nil, err
			}
			return "insert success", nil

		}
	// existing customer just edit record
	case curID != 0:
		{
			//new customer case
			sql := `update Customer
			set 	code = ?,name=?,address=?,telephone=?,bill_credit=?,
				active_status=?,edit_by=?,edit_time=?
			where id = ?`
			fmt.Println(sql)

			_, err := db.Exec(sql, c.Code, c.Name, c.Address,
				c.Telephone, c.BillCredit, 0,
				c.CreateBy, c.CreateTime, curID)

			if err != nil {
				return nil, err
			}
			//_ , err := rs.RowsAffected()
			//if err != nil {
			//	return nil, err
			//}
			return "update success", nil
		}
	}

	// update
	return "success", nil

}
