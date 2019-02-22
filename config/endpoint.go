package config

import (
	"context"
	"fmt"
)

//type Endpoint interface {
//	SearchCustomerById(context.Context, *SearchCustomerByIdRequest) (*SearchCustomerResponse, error)
//}

type (
	SearchCustomerByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchCustomerResponse struct {
		Id         int64  `json:"id"`
		Code       string `json:"code"`
		Name       string `json:"name"`
		Address    string `json:"address"`
		Telephone  string `json:"telephone"`
		BillCredit int64  `json:"bill_credit"`
	}
	SearchSettingTemplate struct {
		Id             int64  `json:"company_id"`
		BranchId       int64  `json:"branch_id"`
		TaxType        int64  `json:"tax_type"`
		TaxRate        int64  `json:"tax_rate"`
		LogoPath       string `json:"logo_path"`
		DepartId       int64  `json:"depart_id"`
		DefSaleWhId    int64  `json:"def_sale_wh_id"`
		DefSaleShelfId int64  `json:"def_sale_shelf_id"`
		DefBuyWhId     int64  `json:"def_buy_wh_id"`
		DefBuyShelfId  int64  `json:"def_buy_shelf_id"`
		StockStatus    int64  `json:"stock_status"`
		SaleTaxType    int64  `json:"sale_tax_type"`
		BuyTaxType     int64  `json:"buy_tax_type"`
		SaleBillType   int64  `json:"sale_bill_type"`
		BuyBillType    int64  `json:"buy_bill_type"`
		UseAddress     string `json:"use_address"`
		PosDefCustId   int64  `json:"pos_def_cust_id"`
		PosDefStock    int64  `json:"pos_def_stock"`
		DefCustId      int64  `json:"def_cust_id"`
		CreateBy       string `json:"create_by"`
		CreateTime     string `json:"create_time"`
		EditBy         string `json:"edit_by"`
		EditTime       string `json:"edit_time"`
	}
)

func SettingSys(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdTemplate) (interface{}, error) {
		resp, err := s.SettingSys(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
