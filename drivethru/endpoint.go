package drivethru

import (
	"fmt"
	"context"
	"time"

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

	NewPickupRequest struct {
		CarNumber   string `json:"car_number"`
		CarBrand    string `json:"car_brand"`
		AccessToken string `json:"access_token"`
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

		return resp, nil
	}
}
func makeShiftOpen(s Service) interface{} {
	type request struct {
		token        string  `json:"token"`
		machineID    int     `json:"machine_id"`
		changeAmount float64 `json:"change_amount"`
		cashierID    int     `json:"cashier_id"`
		whID         int     `json:"wh_id"`
		remark       string  `json:"remark"`
	}
	//maybe : use token to get user to open shift ?
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("start endpoint shift open ....")
		resp, err := s.ShiftOpen(&ShiftOpenRequest{
			Token:        req.token,
			ChangeAmount: req.changeAmount,
			MachineID:    req.machineID,
			CashierID:    req.cashierID,
			Remark:       req.remark,
			Created:      time.Now(),
			WhID:         req.whID,
		})

		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "Shift Open",
				"processDesc": "Success",
				"isSuccess":   true,
			},
			"shift_uuid": resp,
		}, nil
	}
}

func pickupNew(s Service) interface{} {
	return func(ctx context.Context, req *NewPickupRequest) (interface{}, error) {
		fmt.Println("start endpoint userlogin usercode is => ", req.CarNumber)
		resp, err := s.pickupNew(&NewPickupRequest{CarNumber:req.CarNumber,CarBrand:req.CarBrand,AccessToken:req.AccessToken})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func makeShiftClose(s Service) interface{} {
	type request struct {
		token            string  `json:"token"`
		shiftUUID        string  `json:"shift_uuid"`
		sumCashAmount    float64 `json:"sum_cash_amount"`
		sumCreditAmount  float64 `json:"sum_credit_amount"`
		sumBankAmount    float64 `json:"sum_bank_amount"`
		sumCouponAmount  float64 `json:"sum_coupon_amount"`
		sumDepositAmount float64 `json:"sum_deposit_amount"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("start endpoint shift open ....")
		resp, err := s.ShiftClose(&ShiftCloseRequest{
			Token:            req.token,
			ShiftUUID:        req.shiftUUID,
			SumCashAmount:    req.sumCashAmount,
			SumCreditAmount:  req.sumCreditAmount,
			SumBankAmount:    req.sumBankAmount,
			SumCouponAmount:  req.sumCouponAmount,
			SumDepositAmount: req.sumDepositAmount,
			ClosedAt:         time.Now(),
		})
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"response": map[string]interface{}{
				"process":     "Shift Close",
				"processDesc": "Success",
				"isSuccess":   true,
			},
			"data": resp,
		}, nil
	}
}
