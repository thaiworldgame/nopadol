package customer

import (
	"context"
	"fmt"
	"github.com/mrtomyum/nopadol/auth"
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
)

func SearchById(s Service) interface{} {
	return func(ctx context.Context, req *SearchCustomerByIdRequest) (interface{}, error) {
		resp, err := s.SearchById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"member": resp,
		}, nil
	}
}

func SearchByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error ", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"member": resp,
		}, nil
	}
}

func makeNewCustomer(s Service) interface{} {
	type request struct {
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		BillAddress  string  `json:"bill_address"`
		Telephone    string  `json:"telephone"`
		Email        string  `json:"email"`
		CreditAmount float64 `json:"credit_amount"`
		CompanyID    int     `json:"company_id"`
	}

	return func(ctx context.Context, req *request) (interface{}, error) {
		//companyID := auth.GetCompanyID(ctx)
		userID := auth.GetUserCode(ctx)
		ct := CustomerTemplate{
			Code:         req.Code,
			Name:         req.Name,
			Address:      req.BillAddress,
			Telephone:    req.Telephone,
			CreditAmount: req.CreditAmount,
			Email:        req.Email,
			CompanyID:    req.CompanyID,
			CreateBy:     userID,
		}

		fmt.Println("start endpoint store Customer with param , ",ct)
		resp, err := s.StoreCustomer(&ct)
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"result": resp,
		}, nil
	}
}

func makeUpdateCustomer(s Service) interface{} {
	type request struct {
		Id           int64   `json:"id"`
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		BillAddress  string  `json:"bill_address"`
		Telephone    string  `json:"telephone"`
		Email        string  `json:"email"`
		CreditAmount float64 `json:"credit_amount"`
		CompanyID    int     `json:"compayny_id"`
	}

	return func(ctx context.Context, req *request) (interface{}, error) {
		ct := CustomerTemplate{
			Code:         req.Code,
			Name:         req.Name,
			Address:      req.BillAddress,
			Telephone:    req.Telephone,
			CreditAmount: req.CreditAmount,
			Email:        req.Email,
			CompanyID:    req.CompanyID,
		}
		resp, err := s.StoreCustomer(&ct)
		fmt.Println("start endpoint store Customer")
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"result": resp,
		}, nil
	}
}
