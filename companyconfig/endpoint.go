package companyconfig

import (
	"context"
	"fmt"
)

type (
	RequestConfig struct {
		Id           int    `json:"id"`
		ComSysId     int    `json:"com_sys_id"`
		CompanyName  string `json:"company_name"`
		NameEng      string `json:"name_eng"`
		Address      string `json:"address"`
		Telephone    string `json:"telephone"`
		Fax          string `json:"fax"`
		TaxNumber    string `json:"tax_number"`
		Email        string `json:"email"`
		WebSite      string `json:"web_site"`
		TaxRate      int    `json:"tax_rate"`
		StockStatus  int    `json:"stock_status"`
		SaleTaxType  int    `json:"sale_tax_type"`
		BuyTaxType   int    `json:"buy_tax_type"`
		LogoPath     string `json:"logo_path"`
		ActiveStatus int    `json:"active_status"`
		CreateBy     string `json:"create_by"`
		CreateTime   string `json:"create_time"`
		EditBy       string `json:"edit_by"`
		EditTime     string `json:"edit_time"`
	}

	SearchByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}
)

func Create(s Service) interface{} {
	return func(ctx context.Context, req *RequestConfig) (interface{}, error) {
		resp, err := s.Create(&RequestConfigTemplate{
			Id:           req.Id,
			ComSysId:     req.ComSysId,
			CompanyName:  req.CompanyName,
			NameEng:      req.NameEng,
			Address:      req.Address,
			Telephone:    req.Telephone,
			Fax:          req.Fax,
			TaxNumber:    req.TaxNumber,
			Email:        req.Email,
			WebSite:      req.WebSite,
			TaxRate:      req.TaxRate,
			StockStatus:  req.StockStatus,
			SaleTaxType:  req.SaleTaxType,
			BuyTaxType:   req.BuyTaxType,
			LogoPath:     req.LogoPath,
			ActiveStatus: req.ActiveStatus,
			CreateBy:     req.CreateBy,
		})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchById(&SearchByIdRequestTemplate{
			Id: req.Id,
		})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
