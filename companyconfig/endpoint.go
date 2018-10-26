package companyconfig

import (
	"context"
	"fmt"
)

type (
	RequestConfig struct {
		Id              int    `json:"id"`
		ComId           int64  `json:"com_id"`
		ComSysId        int64  `json:"com_sys_id"`
		CompanyName     string `json:"company_name"`
		NameEng         string `json:"name_eng"`
		Address         string `json:"address"`
		Telephone       string `json:"telephone"`
		Fax             string `json:"fax"`
		TaxNumber       string `json:"tax_number"`
		Email           string `json:"email"`
		WebSite         string `json:"web_site"`
		TaxRate         int    `json:"tax_rate"`
		BranchName      string `json:"branch_name"`
		BranchAddress   string `json:"branch_address"`
		BranchTelephone string `json:"branch_telephone"`
		BranchFax       string `json:"branch_fax"`
		StockStatus     int    `json:"stock_status"`
		SaleTaxType     int    `json:"sale_tax_type"`
		BuyTaxType      int    `json:"buy_tax_type"`
		DefSaleWh       string `json:"def_sale_wh"`
		DefSaleShelf    string `json:"def_sale_shelf"`
		DefBuyWh        string `json:"def_buy_wh"`
		DefBuyShelf     string `json:"def_buy_shelf"`
		SaleBillType    int64  `json:"sale_bill_type"`
		BuyBillType     int64  `json:"buy_bill_type"`
		LogoPath        string `json:"logo_path"`
		ActiveStatus    int    `json:"active_status"`
		CreateBy        string `json:"create_by"`
		CreateTime      string `json:"create_time"`
		EditBy          string `json:"edit_by"`
		EditTime        string `json:"edit_time"`
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
			Id:              req.Id,
			ComSysId:        req.ComSysId,
			CompanyName:     req.CompanyName,
			NameEng:         req.NameEng,
			Address:         req.Address,
			Telephone:       req.Telephone,
			Fax:             req.Fax,
			TaxNumber:       req.TaxNumber,
			Email:           req.Email,
			WebSite:         req.WebSite,
			TaxRate:         req.TaxRate,
			BranchName:      req.BranchName,
			BranchAddress:   req.BranchAddress,
			BranchTelephone: req.BranchTelephone,
			BranchFax:       req.BranchFax,
			StockStatus:     req.StockStatus,
			SaleTaxType:     req.SaleTaxType,
			BuyTaxType:      req.BuyTaxType,
			DefSaleWh:       req.DefSaleWh,
			DefSaleShelf:    req.DefSaleShelf,
			DefBuyWh:        req.DefSaleWh,
			DefBuyShelf:     req.DefBuyShelf,
			SaleBillType:    req.SaleBillType,
			BuyBillType:     req.BuyBillType,
			LogoPath:        req.LogoPath,
			ActiveStatus:    req.ActiveStatus,
			CreateBy:        req.CreateBy,
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
