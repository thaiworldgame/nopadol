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
		ZoneId   string `json:"pick_zone_id"`
		ZoneName string `json:"name"`
	}

	UserLogInRequest struct {
		UserCode     string `json:"user_code"`
		Password     string `json:"password"`
		BranchId     string `json:"branch_id"`
		ServerName   string `json:"server_name"`
		DataBaseName string `json:"data_base_name"`
	}

	LoginRequest struct {
		EmployeeCode string `json:"employee_code"`
		BranchId     string `json:"branch_id"`
		EmployeeName string `json:"employee_name"`
		ServerName   string `json:"server_name"`
		DatabaseName string `json:"database_name"`
	}

	NewPickupRequest struct {
		CarNumber   string `json:"carNumber"`
		CarBrand    string `json:"carBrand"`
		DocType     string `json:"doc_type"`//0 drivethru 1 pos 2 saleorder
		AccessToken string `json:"access_token"`
	}

	ListQueueRequest struct {
		AccessToken string `json:"access_token"`
		CreateDate  string `json:"create_date"`
		PickDate    string `json:"pick_date"`
		Status      int    `json:"status"`
		Page        string `json:"page"`
		Keyword     string `json:"keyword"`
		QueId       int    `json:"que_id"`
	}

	ManagePickupRequest struct {
		AccessToken string  `json:"access_token"`
		QueueId     int     `json:"queue_id"`
		ItemBarcode string  `json:"item_barcode"`
		QtyBefore   float64 `json:"qty_before"`
		IsCancel    int     `json:"is_cancel"`
		LineNumber  int     `json:"line_number"`
	}

	ManageCheckoutRequest struct {
		AccessToken string  `json:"access_token"`
		QueueId     int     `json:"queue_id"`
		ItemBarcode string  `json:"item_barcode"`
		QtyAfter    float64 `json:"qty_after"`
		IsCancel    int     `json:"is_cancel"`
		LineNumber  int     `json:"line_number"`
	}

	QueueEditRequest struct {
		AccessToken string `json:"access_token"`
		CarBrand    string `json:"car_brand"`
		QueueId     int    `json:"queue_id"`
		Status      int    `json:"status"`
		SaleCode    string `json:"sale_code"`
		PlateNumber string `json:"plate_number"`
	}

	PickupEditRequest struct {
		AccessToken string `json:"access_token"`
		CarBrand    string `json:"carBrand"`
		QId     int    `json:"qId"`
		Status      int    `json:"status"`
		SaleCode    string `json:"saleCode"`
		CarNumber string `json:"carNumber"`
	}

	QueueStatusRequest struct {
		AccessToken               string `json:"access_token"`
		QueueId                   int    `json:"queue_id"`
		StatusForSaleorderCurrent int    `json:"status_for_saleorder_current"`
		IsLoad                    int    `json:"is_load"`
		CancelRemark              string `json:"cancel_remark"`
	}

	QueueProductRequest struct {
		AccessToken string `json:"access_token"`
		QueueId     int    `json:"queue_id"`
	}

	AccessTokenRequest struct {
		AccessToken string `json:"access_token"`
	}

	BillingDoneRequest struct {
		AccessToken   string        `json:"access_token"`
		ArCode        string        `json:"ar_code"`
		Confirm       int           `json:"confirm"`
		QueueId       int           `json:"queue_id"`
		Cash          float64       `json:"cash"`
		ScgId         string        `json:"scg_id"`
		CreditCard    []*CreditCard `json:"credit_card"`
		CouponCode    []*Coupon     `json:"coupon_code"`
		DepositAmount []*Deposit    `json:"deposit_amount"`
	}

	CreditCard struct {
		CardNo       string  `json:"card_no"`
		ConfirmNo    string  `json:"confirm_no"`
		CreditType   string  `json:"credit_type"`
		BankCode     string  `json:"bank_code"`
		Amount       float64 `json:"amount"`
		ChargeAmount float64 `json:"charge_amount"`
	}

	Coupon struct {
		CouponCode string  `json:"coupon_code"`
		Amount     float64 `json:"amount"`
	}

	Deposit struct {
		DepositId string  `json:"deposit_id"`
		Amount    float64 `json:"amount"`
	}
)

func makeListCompany(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resp, err := s.SearchListCompany()
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}

		return resp, nil
	}
}

func makeListZone(s Service) interface{} {
	return func(ctx context.Context, req *AccessTokenRequest) (interface{}, error) {
		resp, err := s.SearchListZone(req.AccessToken)
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}

		return resp, nil
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
			"brand": resp,
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

func logIn(s Service) interface{} {
	return func(ctx context.Context, req *LoginRequest) (interface{}, error) {
		fmt.Println("start endpoint userlogin usercode is => ", req.EmployeeCode)
		resp, err := s.LogIn(&LoginRequest{BranchId: req.BranchId, EmployeeCode: req.EmployeeCode, EmployeeName: req.EmployeeName})
		if err != nil {
			return nil, err
		}

		return resp, nil
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
		accessToken  string  `json:"accessToken"`
		MachineID    int     `json:"machine_id"`
		ChangeAmount float64 `json:"change_amount"`
		CashierID    int     `json:"cashier_id"`
		WhID         int     `json:"wh_id"`
		Remark       string  `json:"remark"`
	}
	//maybe : use token to get user to open shift ?
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("start endpoint shift open ....")

		//validate request data
		if req.accessToken == "" {
			return nil, fmt.Errorf("access token is require..")
		}

		resp, err := s.ShiftOpen(&ShiftOpenRequest{
			AccessToken:  req.accessToken,
			ChangeAmount: req.ChangeAmount,
			MachineID:    req.MachineID,
			CashierID:    req.CashierID,
			Remark:       req.Remark,
			Created:      time.Now(),
			WhID:         req.WhID,
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

func pickupNew(s Service) interface{} {//API
	return func(ctx context.Context, req *NewPickupRequest) (interface{}, error) {
		fmt.Println("start endpoint pickupnew car number is => ", req.CarNumber)
		resp, err := s.PickupNew(&NewPickupRequest{CarNumber: req.CarNumber, CarBrand: req.CarBrand, AccessToken: req.AccessToken, DocType:req.DocType})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func managePickup(s Service) interface{} {
	return func(ctx context.Context, req *ManagePickupRequest) (interface{}, error) {
		fmt.Println("start endpoint mange pickup que id is => ", req.QueueId)
		resp, err := s.ManagePickup(&ManagePickupRequest{AccessToken: req.AccessToken, QueueId: req.QueueId, ItemBarcode: req.ItemBarcode, QtyBefore: req.QtyBefore, IsCancel: req.IsCancel})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func cancelQueue(s Service) interface{} {
	return func(ctx context.Context, req *QueueStatusRequest) (interface{}, error) {
		fmt.Println("start endpoint mange pickup que id is => ", req.QueueId)
		resp, err := s.CancelQueue(&QueueStatusRequest{AccessToken: req.AccessToken, QueueId: req.QueueId, CancelRemark: req.CancelRemark, IsLoad: req.IsLoad, StatusForSaleorderCurrent: req.StatusForSaleorderCurrent})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func manageCheckout(s Service) interface{} {
	return func(ctx context.Context, req *ManageCheckoutRequest) (interface{}, error) {
		fmt.Println("start endpoint mange checkout que id is => ", req.QueueId)
		resp, err := s.ManageCheckout(&ManageCheckoutRequest{AccessToken: req.AccessToken, QueueId: req.QueueId, ItemBarcode: req.ItemBarcode, QtyAfter: req.QtyAfter, IsCancel: req.IsCancel})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func makeSearchListQueue(s Service) interface{} {
	return func(ctx context.Context, req *ListQueueRequest) (interface{}, error) {
		fmt.Println("start endpoint list queue createdate is => ", req.CreateDate)
		resp, err := s.ListQueue(&ListQueueRequest{CreateDate: req.CreateDate, PickDate: req.PickDate, Status: req.Status, Page: req.Page, Keyword: req.Keyword, QueId: req.QueId, AccessToken: req.AccessToken})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func queueProduct(s Service) interface{} {
	return func(ctx context.Context, req *QueueProductRequest) (interface{}, error) {
		fmt.Println("start endpoint queue product queue is => ", req.QueueId)
		resp, err := s.QueueProduct(&QueueProductRequest{AccessToken: req.AccessToken, QueueId: req.QueueId})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func pickupEdit(s Service) interface{} {
	return func(ctx context.Context, req *PickupEditRequest) (interface{}, error) {
		fmt.Println("start endpoint list queue edit is => ", req.QId)
		resp, err := s.PickupEdit(&PickupEditRequest{QId: req.QId, CarBrand: req.CarBrand, CarNumber: req.CarNumber, SaleCode: req.SaleCode, Status: req.Status, AccessToken: req.AccessToken})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func queueEdit(s Service) interface{} {
	return func(ctx context.Context, req *QueueEditRequest) (interface{}, error) {
		fmt.Println("start endpoint list queue edit is => ", req.QueueId)
		resp, err := s.QueueEdit(&QueueEditRequest{QueueId: req.QueueId, CarBrand: req.CarBrand, PlateNumber: req.PlateNumber, SaleCode: req.SaleCode, Status: req.Status, AccessToken: req.AccessToken})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func queueStatus(s Service) interface{} {
	return func(ctx context.Context, req *QueueStatusRequest) (interface{}, error) {
		fmt.Println("start endpoint list queue status is => ", req.QueueId)
		fmt.Println("start endpoint list queue edit is => ", req.QueueId)
		resp, err := s.QueueStatus(&QueueStatusRequest{AccessToken: req.AccessToken, QueueId: req.QueueId, StatusForSaleorderCurrent: req.StatusForSaleorderCurrent, IsLoad: req.IsLoad})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func billingDone(s Service) interface{} {
	return func(ctx context.Context, req *BillingDoneRequest) (interface{}, error) {
		resp, err := s.BillingDone(&BillingDoneRequest{QueueId: req.QueueId, Confirm: req.Confirm, ArCode: req.ArCode, AccessToken: req.AccessToken, CouponCode: req.CouponCode, Cash: req.Cash, CreditCard: req.CreditCard, DepositAmount: req.DepositAmount})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func makeShiftClose(s Service) interface{} {
	type request struct {
		Token            string  `json:"token"`
		ShiftUUID        string  `json:"shift_uuid"`
		SumCashAmount    float64 `json:"sum_cash_amount"`
		SumCreditAmount  float64 `json:"sum_credit_amount"`
		SumBankAmount    float64 `json:"sum_bank_amount"`
		SumCouponAmount  float64 `json:"sum_coupon_amount"`
		SumDepositAmount float64 `json:"sum_deposit_amount"`
	}
	return func(ctx context.Context, req *request) (interface{}, error) {
		fmt.Println("start endpoint shift close ..request -> ", req)
		//validate request data
		if req.ShiftUUID == "" || req.Token == "" {
			return nil, fmt.Errorf("shift id is empty value")
		}

		resp, err := s.ShiftClose(&ShiftCloseRequest{
			Token:            req.Token,
			ShiftUUID:        req.ShiftUUID,
			SumCashAmount:    req.SumCashAmount,
			SumCreditAmount:  req.SumCreditAmount,
			SumBankAmount:    req.SumBankAmount,
			SumCouponAmount:  req.SumCouponAmount,
			SumDepositAmount: req.SumDepositAmount,
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
