package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/mrtomyum/nopadol/employee"
)

type EmployeeModel struct {
	Id       int64  `db:"Id"`
	SaleCode string `db:"SaleCode"`
	SaleName string `db:"SaleName"`
}

type employeeRepository struct{ db *sqlx.DB }

func NewEmployeeRepository(db *sqlx.DB) employee.Repository {
	return &employeeRepository{db}
}

func (em *employeeRepository) SearchById(req *employee.SearchByIdTemplate) (resp interface{}, err error) {
	emp := EmployeeModel{}
	sql := `select Id, SaleCode, SaleName from Sale where Activestatus =1 and id = ?`
	err = em.db.Get(&emp, sql, req.Id)
	if err != nil {
		fmt.Println("error =", err.Error())
		return resp, err
	}

	emp_resp := map_employee_template(emp)

	return map[string]interface{}{
		"sale_id":   emp_resp.Id,
		"sale_code": emp_resp.SaleCode,
		"sale_name": emp_resp.SaleName,
	}, nil
}

func (em *employeeRepository) SearchByKeyword(req *employee.SearchByKeywordTemplate) (resp interface{}, err error) {
	emps := []EmployeeModel{}

	sql := `select Id,SaleCode, SaleName from Sale where Activestatus =1 and (SaleCode like concat('%',?,'%') or SaleName like concat('%',?,'%')) order by SaleCode limit 20`
	err = em.db.Select(&emps, sql, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("error =", err.Error())
		return resp, err
	}

	employee := []employee.EmployeeTemplate{}

	for _, e := range emps {
		empline := map_employee_template(e)
		employee = append(employee, empline)
	}

	return employee, nil
}

func map_employee_template(x EmployeeModel) employee.EmployeeTemplate {
	return employee.EmployeeTemplate{
		Id:       x.Id,
		SaleCode: x.SaleCode,
		SaleName: x.SaleName,
	}
}
