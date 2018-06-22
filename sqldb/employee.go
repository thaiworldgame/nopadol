package sqldb

import (
	"github.com/jmoiron/sqlx"
	"context"
	"github.com/mrtomyum/nopadol/employee"
	"fmt"
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

func (em *employeeRepository) SearchEmployeeById(ctx context.Context, req *employee.SearchById) (resp employee.EmployeeTemplate, err error) {
	emp := EmployeeModel{}
	sql := `select RowOrder as Id,code as SaleCode, name as SaleName from dbo.bcsale where code = '56163'`
	err = em.db.Get(&emp, sql)
	if err != nil {
		fmt.Println("error =", err.Error())
		return resp, err
	}

	Resp := map_employee_template(emp)
	return Resp, nil
}

func map_employee_template(x EmployeeModel) employee.EmployeeTemplate {
	return employee.EmployeeTemplate{
		Id:       x.Id,
		SaleCode: x.SaleCode,
		SaleName: x.SaleName,
	}
}
