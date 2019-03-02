package configuration

import (
	"context"
	"fmt"
)

type (
	RequestConfig struct {
		Id             int64  `json:"id"`
		CompanyId      int64  `json:"company_id"`
		BranchId       int64  `json:"branch_id"`
		TaxType        string `json:"tax_type"`
		TaxRate        string `json:"tax_rate"`
		LogoPath       string `json:"logo_path"`
		DepartId       string `json:"depart_id"`
		DefSaleWhId    string `json:"def_sale_wh_id"`
		DefSaleShelfId string `json:"def_sale_shelf_id"`
		DefBuyWhId     string `json:"def_buy_wh_id"`
		DefBuyShelfId  string `json:"def_buy_shelf_id"`
		SrockStatus    int    `json:"stock_status"`
		SaleTaxType    string `json:"sale_tax_type"`
		BuyTaxType     string `json:"buy_tax_type"`
		SaleBillType   string `json:"sale_bill_type"`
		BuyBillType    string `json:"buy_bill_type"`
		UseAddress     int    `json:"use_address"`
		PosDefCustId   int    `json:"pos_def_cust_id"`
		PosDefStock    int    `json:"pos_def_stock"`
		DefCustId      string `json:"def_cust_id"`
		CreateBy       string `json:"create_by"`
		CreateTime     string `json:"create_time"`
		EditBy         string `json:"edit_by"`
		EditTime       string `json:"edit_time"`
		BranchName     string `json:"branch_name"`
		Address        string `json:"address"`
		Telephone      string `json:"telephone"`
		Fax            string `json:"fax"`
		CompanyName    string `json:"company_name"`
	}

	SearchByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}
)

func ConfigSetting(s Service) interface{} {
	return func(ctx context.Context, req *RequestConfig) (interface{}, error) {
		resp, err := s.ConfigSetting(&RequestSettingTemplate{
			Id:             req.Id,
			CompanyId:      req.CompanyId,
			BranchId:       req.BranchId,
			TaxType:        req.TaxType,
			TaxRate:        req.TaxRate,
			LogoPath:       req.LogoPath,
			DepartId:       req.DepartId,
			DefSaleWhId:    req.DefSaleWhId,
			DefSaleShelfId: req.DefSaleShelfId,
			DefBuyWhId:     req.DefBuyWhId,
			DefBuyShelfId:  req.DefBuyShelfId,
			SrockStatus:    req.SrockStatus,
			SaleTaxType:    req.SaleTaxType,
			BuyTaxType:     req.BuyTaxType,
			SaleBillType:   req.SaleBillType,
			BuyBillType:    req.BuyBillType,
			UseAddress:     req.UseAddress,
			PosDefCustId:   req.PosDefCustId,
			PosDefStock:    req.PosDefStock,
			DefCustId:      req.DefCustId,
			CreateBy:       req.CreateBy,
			CreateTime:     req.CreateTime,
			EditBy:         req.EditBy,
			EditTime:       req.EditTime,
			BranchName:     req.BranchName,
			Address:        req.Address,
			Telephone:      req.Telephone,
			Fax:            req.Fax,
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

func SearchSettingById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchSettingById(&SearchByIdRequestTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchSettingByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchSettingByKeyword(&SearchByKeywordRequestTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
