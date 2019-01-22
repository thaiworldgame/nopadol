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

	cust_resp := map_data_template(d)

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
		depline := map_data_template(d)
		departments = append(departments, depline)
	}

	return departments, nil
}


func (repo *envRepository) SearchProjectById(req *environment.SearchByIdTemplate) (resp interface{}, err error) {

	p := DepartmentModel{}

	sql := `select id,code,name from Department where id = ? and active_status =1`
	err = repo.db.Get(&p, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	cust_resp := map_data_template(p)

	return map[string]interface{}{
		"id":        cust_resp.Id,
		"code":      cust_resp.Code,
		"name":      cust_resp.Name,
	}, nil
}


func (repo *envRepository) SearchProjectByKeyword(req *environment.SearchByKeywordTemplate) (resp interface{}, err error) {

	projs := []DepartmentModel{}

	sql := `select id,code,name from Project where (code like concat('%',?,'%')  or name like concat('%',?,'%'))  and active_status =1 order by code`
	err = repo.db.Select(&projs, sql, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	projects := []environment.DepartmentTemplate{}
	for _, d := range projs {
		depline := map_data_template(d)
		projects = append(projects, depline)
	}

	return projects, nil
}


func (repo *envRepository) SearchAllocateById(req *environment.SearchByIdTemplate) (resp interface{}, err error) {

	a := DepartmentModel{}

	sql := `select id,code,name from Allocate where id = ? and active_status =1`
	err = repo.db.Get(&a, sql, req.Id)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	cust_resp := map_data_template(a)

	return map[string]interface{}{
		"id":        cust_resp.Id,
		"code":      cust_resp.Code,
		"name":      cust_resp.Name,
	}, nil
}


func (repo *envRepository) SearchAllocateByKeyword(req *environment.SearchByKeywordTemplate) (resp interface{}, err error) {

	allos := []DepartmentModel{}

	sql := `select id,code,name from Allocate where (code like concat('%',?,'%')  or name like concat('%',?,'%'))  and active_status =1 order by code`
	err = repo.db.Select(&allos, sql, req.Keyword, req.Keyword)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	allocates := []environment.DepartmentTemplate{}
	for _, d := range allos {
		depline := map_data_template(d)
		allocates = append(allocates, depline)
	}

	return allocates, nil
}


func map_data_template(x DepartmentModel) environment.DepartmentTemplate {
	fmt.Println("Code =",x.Code)
	return environment.DepartmentTemplate{
		Id:        x.Id,
		Code:      x.Code,
		Name:      x.Name,
	}
}
