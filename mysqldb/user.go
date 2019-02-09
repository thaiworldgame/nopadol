package mysqldb

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrtomyum/nopadol/drivethru"
	"fmt"
	"time"
)

type userLogInModel struct {
	Id               string `json:"id"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	Password         string `json:"password"`
	ImageFileName    string `json:"image_filename"`
	Role             int    `json:"role"`
	ActiveStatus     int    `json:"activeStatus"`
	IsConfirm        int    `json:"isConfirm"`
	CreatorCode      string `json:"creatorCode"`
	CreateDateTime   string `json:"createdateTime"`
	LastEditorCode   string `json:"lasteditorCode"`
	LastEditDateTime string `json:"lasteditdateTime"`
	BranchCode       string `json:"branchCode"`
	Remark           string `json:"remark"`
	LoginZone        string `json:"loginZone"`
	CompanyId        int    `json:"company_id"`
	BranchId         int    `json:"branch_id"`
	SaleId           int    `json:"sale_id"`
	SaleCode         string `json:"sale_code"`
}

type loginModel struct {
	employeeCode  string `json:"employee_code"`
	branchId      int    `json:"branch_id"`
	employeeName  string `json:"employee_name"`
	server_name   string `json:"server_name"`
	database_name string `json:"database_name"`
}

func (u *userLogInModel) Userlogin(db *sqlx.DB, req *drivethru.UserLogInRequest) (interface{}, error) {
	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02 00:00:00"))
	//DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	now.String()

	fmt.Println("mysql recived param req -> ", req.UserCode, req.BranchId)

	branch_code := getBranch(db, req.BranchId)

	lccommand := "select id,code,name,ifnull(pwd,'') as password,ifnull(image_path,'') as image_filename,role,active_status as activesTatus,is_confirm as isConfirm,ifnull(create_by,'') as creatorCode,ifnull(create_time,'') as createdateTime,ifnull(edit_by,'') as lasteditorCode,ifnull(edit_time,'') as lasteditdateTime,ifnull(branch_code,'') as branchCode,'' as remark,ifnull(zone_id,'') as loginZone,ifnull(company_id,1) as company_id,ifnull(branch_id,1) as branch_id from user where code = ? and branch_code = ? and pwd = ?"
	fmt.Println(lccommand, req.UserCode, branch_code)
	rs := db.QueryRow(lccommand, req.UserCode, branch_code, &req.Password)

	user := userLogInModel{}
	err := rs.Scan(&user.Id, &user.Code, &user.Name, &user.Password, &user.ImageFileName, &user.Role, &user.ActiveStatus, &user.IsConfirm, &user.CreatorCode, &user.CreateDateTime, &user.LastEditorCode, &user.LastEditDateTime, &user.BranchCode, &user.Remark, &user.LoginZone, &user.CompanyId, &user.BranchId)
	if err != nil || user.Code == "" {
		fmt.Println("error = ", err.Error())
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "login",
				"processDesc": "login fail",
				"isSuccess":   false,
			},
		}, nil
	}
	fmt.Println("before mysql return -> ", user)

	var check_exist int
	var uuid string

	lccommand_check := `select count(*) as vCount from user_access where user_id = ? and user_code = ? and CONVERT_TZ(CURRENT_TIMESTAMP,'+00:00','+7:00') < expire_time`
	err = db.Get(&check_exist, lccommand_check, user.Id, req.UserCode)
	if err != nil {
		fmt.Println(err.Error())
	}

	if check_exist == 0 {
		uuid = GetAccessToken()
	} else {
		lccommand_check := `select access_token from user_access where user_id = ? and user_code = ? and expire_time > CURRENT_TIMESTAMP`
		err = db.Get(&uuid, lccommand_check, user.Id, req.UserCode)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(check_exist)
	var expire_date string

	expire := now.AddDate(0, 0, int(1)).Format("2006-01-02 15:04:05")

	expire_date = expire
	if check_exist == 0 {
		lccommand = `START TRANSACTION`
		_, err := db.Exec(lccommand)

		lccommand = "insert user_access(user_id,user_code,access_token,company_id,branch_id,branch_code,zone_id,create_time,expire_time) values(?,?,?,?,?,?,?,?,?)"
		_, err = db.Exec(lccommand, user.Id, user.Code, uuid, user.CompanyId, user.BranchId, branch_code, user.LoginZone, now.String(), expire_date)
		if err != nil {
			lccommand = `ROLLBACK`
			_, err = db.Exec(lccommand)
			fmt.Println("error = ", err.Error())
			return map[string]interface{}{
				"response": map[string]interface{}{
					"process":     "login",
					"processDesc": err.Error(),
					"isSuccess":   false,
				},
			}, nil
		}
		lccommand = `COMMIT`
		_, err = db.Exec(lccommand)

	} else {
		lccommand = "update user_access set last_login_time = ? where user_id = ? and user_code = ?"
		ins, err := db.Exec(lccommand, now.String(), user.Id, user.Code)
		if err != nil {
			fmt.Println("error = ", err.Error())
			return map[string]interface{}{
				"response": map[string]interface{}{
					"process":     "login",
					"processDesc": err.Error(),
					"isSuccess":   false,
				},
			}, nil
		}
		fmt.Println(ins.RowsAffected())
	}

	return map[string]interface{}{
		"response": map[string]interface{}{
			"process":     "login",
			"processDesc": "successful",
			"isSuccess":   true,
		},
		"accessToken":    uuid,
		"accessDatetime": now.String(),
		"pathPHPUpload":  "http://qserver.nopadol.com/drivethru/upload.php",
		"pathFile":       "http://qserver.nopadol.com/drivethru/tmp/",
		"user":           user,
	}, nil
}

func (u *userLogInModel) login(db *sqlx.DB, req *loginModel) (interface{}, error) {
	var uuid string

	now := time.Now()
	fmt.Println("yyyy-mm-dd date format : ", now.AddDate(0, 0, 0).Format("2006-01-02 00:00:00"))
	//DocDate := now.AddDate(0, 0, 0).Format("2006-01-02")

	now.String()
	branch_code := getBranch(db, req.branchId)

	lccommand := "select id,code,name,ifnull(pwd,'') as password,ifnull(image_path,'') as image_filename,role,active_status as activesTatus,is_confirm as isConfirm,ifnull(create_by,'') as creatorCode,ifnull(create_time,'') as createdateTime,ifnull(edit_by,'') as lasteditorCode,ifnull(edit_time,'') as lasteditdateTime,ifnull(branch_code,'') as branchCode,'' as remark,ifnull(zone_id,'') as loginZone,ifnull(company_id,1) as company_id,ifnull(branch_id,1) as branch_id from user where code = ? and branch_code = ? "
	fmt.Println(lccommand, req.employeeCode, branch_code)
	rs := db.QueryRow(lccommand, req.employeeCode, branch_code)

	user := userLogInModel{}
	err := rs.Scan(&user.Id, &user.Code, &user.Name, &user.Password, &user.ImageFileName, &user.Role, &user.ActiveStatus, &user.IsConfirm, &user.CreatorCode, &user.CreateDateTime, &user.LastEditorCode, &user.LastEditDateTime, &user.BranchCode, &user.Remark, &user.LoginZone, &user.CompanyId, &user.BranchId)
	if err != nil || user.Code == "" {
		fmt.Println("error = ", err.Error())
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "login",
				"processDesc": "login fail",
				"isSuccess":   false,
			},
		}, nil
	}
	fmt.Println("before mysql return -> ", user)

	if user.Name != "" {
		var check_exist int

		lccommand_check := `select count(*) as vCount from user_access where user_id = ? and user_code = ? and CONVERT_TZ(CURRENT_TIMESTAMP,'+00:00','+7:00') < expire_time`
		err = db.Get(&check_exist, lccommand_check, user.Id, req.employeeCode)
		if err != nil {
			fmt.Println(err.Error())
		}

		if check_exist == 0 {
			uuid = GetAccessToken()
		} else {
			lccommand_check := `select access_token from user_access where user_id = ? and user_code = ? and expire_time > CURRENT_TIMESTAMP`
			err = db.Get(&uuid, lccommand_check, user.Id, req.employeeCode)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		fmt.Println(check_exist)
		var expire_date string

		expire := now.AddDate(0, 0, int(1)).Format("2006-01-02 15:04:05")

		expire_date = expire
		if check_exist == 0 {
			lccommand = `START TRANSACTION`
			_, err := db.Exec(lccommand)

			lccommand = "insert user_access(user_id,user_code,access_token,company_id,branch_id,branch_code,zone_id,create_time,expire_time) values(?,?,?,?,?,?,?,?,?)"
			_, err = db.Exec(lccommand, user.Id, user.Code, uuid, user.CompanyId, user.BranchId, branch_code, user.LoginZone, now.String(), expire_date)
			if err != nil {
				lccommand = `ROLLBACK`
				_, err = db.Exec(lccommand)
				fmt.Println("error = ", err.Error())
				return map[string]interface{}{
					"response": map[string]interface{}{
						"process":     "login",
						"processDesc": err.Error(),
						"isSuccess":   false,
					},
				}, nil
			}
			lccommand = `COMMIT`
			_, err = db.Exec(lccommand)

		} else {
			lccommand = "update user_access set last_login_time = ? where user_id = ? and user_code = ?"
			ins, err := db.Exec(lccommand, now.String(), user.Id, user.Code)
			if err != nil {
				fmt.Println("error = ", err.Error())
				return map[string]interface{}{
					"response": map[string]interface{}{
						"process":     "login",
						"processDesc": err.Error(),
						"isSuccess":   false,
					},
				}, nil
			}
			fmt.Println(ins.RowsAffected())
		}
	}

	return map[string]interface{}{
		"response": map[string]interface{}{
			"process":     "login",
			"processDesc": "successful",
			"isSuccess":   true,
		},
		"accessToken":    uuid,
		"accessDatetime": now.String(),
		"pathPHPUpload":  "http://qserver.nopadol.com/drivethru/upload.php",
		"pathFile":       "http://qserver.nopadol.com/drivethru/tmp/",
		"user":           user,
	}, nil
}
