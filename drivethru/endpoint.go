package drivethru

import (
	"fmt"
	"context"
)

type (
	SearchById struct {
		Id int64 `json:"keyword"`
	}

	CompanyList struct {
		ListCompany []Company `json:"list_company"`
	}

	Company struct {
		CompanyId   string `json:"company_id"`
		CompanyName string `json:"company_name"`
		ListZone    []Zone `json:"list_zone"`
	}

	Zone struct {
		ZoneId   string `json:"zone_id"`
		ZoneName string `json:"zone_name"`
	}

	UserLogInRequest struct {
		UserCode     string `json:"user_code"`
		Password     string `json:"password"`
		BranchId     int    `json:"branch_id"`
		ServerName   string `json:"server_name"`
		DataBaseName string `json:"data_base_name"`
	}
)

func makeListCompany(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.SearchListCompany()
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}

		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "Search Brand",
				"processDesc": "Success",
				"isSuccess":   true,
			},
			"data": resp,
		}, nil
	}
}

func makeListMachine(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.SearchListMachine()
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "List Pos Machine",
				"processDesc": "Success",
				"isSuccess":   true,
			},
			"data": resp,
		}, nil
	}
}

func makeSearchCarBranch(s Service) interface{} {
	type request struct {
		Keyword string `json:"keyword"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("endpoint keyword is =>", req.Keyword)
		resp, err := s.SearchCarBrand(req.Keyword)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "Search car brand",
				"processDesc": "Success",
				"isSuccess":   true,
			},
			"data": resp,
		}, nil
	}
}

func makeSearchCustomer(s Service) interface{} {
	type request struct {
		Keyword string `json:"keyword"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("endpoint keyword is =>", req.Keyword)
		resp, err := s.SearchCustomer(req.Keyword)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "Search customer",
				"processDesc": "Success",
				"isSuccess":   true,
			},
			"data": resp,
		}, nil
	}
}

func makeItemSearch(s Service) interface{} {
	type request struct {
		Keyword string `json:"keyword"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("start endpoint item search keyword is => ", req.Keyword)
		resp, err := s.SearchItem(req.Keyword)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"response": map[string]interface{}{
				"success": true,
				"error":   false,
				"message": "",
			},
			"item": resp,
		}, nil
	}
}

func userLogIn(s Service) interface{} {
	return func(ctx context.Context, req *UserLogInRequest) (interface{}, error) {
		fmt.Println("start endpoint userlogin usercode is => ", req.UserCode)
		resp, err := s.UserLogIn(&UserLogInRequest{BranchId: req.BranchId, UserCode: req.UserCode, Password: req.Password})
		if err != nil {
			return nil, err
		}
		//return map[string]interface{}{
		//	"response": map[string]interface{}{
		//		"process":     "login",
		//		"processDesc": "successful",
		//		"isSuccess": resp,
		//	},
		//	"accessToken":    "",
		//	"accessDatetime": "",
		//	"pathPHPUpload":  "http://qserver.nopadol.com/drivethru/upload.php",
		//	"pathFile":       "http://qserver.nopadol.com/drivethru/tmp/",
		//	"user":           resp,
		//}, nil

		return resp, nil
	}
}
