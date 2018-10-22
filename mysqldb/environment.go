package mysqldb

import (
	"fmt"
	"github.com/mrtomyum/nopadol/environment"
	"github.com/jmoiron/sqlx"
)

type DepartmentModel struct {
	Id           int64  `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	ActiveStatus int64  `json:"active_status"`
	CreateBy     string `json:"create_by"`
	CreateTime   string `json:"create_time"`
	EditBy       string `json:"edit_by"`
	EditTime     string `json:"edit_time"`
}

type envRepository struct{ db *sqlx.DB }

func NewEnvironmentRepository(db *sqlx.DB) environment.Repository {
	return &envRepository{db}
}

func (repo *envRepository) SearchDepartmentById(req *environment.SearchByIdTemplate) (resp interface{}, err error) {

	d := DepartmentModel{}

	sql := `select id,code,name from Department where id = ? and active_status =1`
	err = repo.db.Get(&d, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	cust_resp := map_department_template(d)

	return map[string]interface{}{
		"id":        cust_resp.Id,
		"code":      cust_resp.Code,
		"name":      cust_resp.Name,
	}, nil
}

func (repo *envRepository) SearchDepartmentByKeyword(req *environment.SearchByKeywordTemplate) (resp interface{}, err error) {

	deps := []DepartmentModel{}

	sql := `select id,code,name from Department where (code like concat('%',?,'%')  or name like concat('%',?,'%'))  and active_status =1 order by code`
	err = repo.db.Select(&deps, sql, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	departments := []environment.DepartmentTemplate{}
	for _, d := range deps {
		depline := map_department_template(d)
		departments = append(departments, depline)
	}

	return departments, nil
}


func map_department_template(x DepartmentModel) environment.DepartmentTemplate {
	fmt.Println("Code =",x.Code)
	return environment.DepartmentTemplate{
		Id:        x.Id,
		Code:      x.Code,
		Name:      x.Name,
	}
}
