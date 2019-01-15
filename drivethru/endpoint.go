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
		BranchId     int64  `json:"branch_id"`
		ServerName   string `json:"server_name"`
		DataBaseName string `json:"data_base_name"`
	}
)

func SearchListCompany(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.SearchListCompany()
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}


