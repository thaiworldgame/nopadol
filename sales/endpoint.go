package sales

import (
	"fmt"
	"context"
)

type (
	NewQuoRequest struct {
		Id                  int64               `json:"id"`
		DocNo               string              `json:"doc_no"`
		DocDate             string              `json:"doc_date"`
		CompanyId           int64               `json:"company_id"`
		BranchId            int64               `json:"branch_id"`
		DocType             int64               `json:"doc_type"`
		ArId                int64               `json:"ar_id"`
		ArCode              string              `json:"ar_code"`
		ArName              string              `json:"ar_name"`
		ArBillAddress       string              `json:"ar_bill_address"`
		ArTelephone         string              `json:"ar_telephone"`
		SaleId              int                 `json:"sale_id"`
		SaleCode            string              `json:"sale_code"`
		SaleName            string              `json:"sale_name"`
		BillType            int64               `json:"bill_type"`
		TaxType             int                 `json:"tax_type"`
		TaxRate             float64             `json:"tax_rate"`
		DepartId            int64               `json:"depart_id"`
		RefNo               string              `json:"ref_no"`
		JobId               string              `json:"job_id"`
		IsConfirm           int64               `json:"is_confirm"`
		BillStatus          int64               `json:"bill_status"`
		Validity            int64               `json:"validity"`
		CreditDay           int64               `json:"credit_day"`
		DueDate             string              `json:"due_date"`
		ExpireCredit        int64               `json:"expire_credit"`
		ExpireDate          string              `json:"expire_date"`
		DeliveryDay         int64               `json:"delivery_day"`
		DeliveryDate        string              `json:"delivery_date"`
		AssertStatus        int64               `json:"assert_status"`
		IsConditionSend     int64               `json:"is_condition_send"`
		MyDescription       string              `json:"my_description"`
		SumOfItemAmount     float64             `json:"sum_of_item_amount"`
		DiscountWord        string              `json:"discount_word"`
		DiscountAmount      float64             `json:"discount_amount"`
		AfterDiscountAmount float64             `json:"after_discount_amount"`
		BeforeTaxAmount     float64             `json:"before_tax_amount"`
		TaxAmount           float64             `json:"tax_amount"`
		TotalAmount         float64             `json:"total_amount"`
		NetDebtAmount       float64             `json:"net_debt_amount"`
		ProjectId           int64               `json:"project_id"`
		AllocateId          int64               `json:"allocate_id"`
		IsCancel            int64               `json:"is_cancel"`
		CreateBy            string              `json:"creator_by"`
		CreateTime          string              `json:"create_time"`
		EditBy              string              `json:"edit_by"`
		EditTime            string              `json:"edit_time"`
		CancelBy            string              `json:"cancel_by"`
		CancelTime          string              `json:"cancel_time"`
		Subs                []NewQuoItemRequest `json:"subs"`
	}

	NewQuoItemRequest struct {
		Id              int64   `json:"id"`
		QuoId           int64   `json:"quo_id"`
		ItemId          int64   `json:"item_id"`
		ItemCode        string  `json:"item_code"`
		BarCode         string  `json:"bar_code"`
		ItemName        string  `json:"item_name"`
		Qty             float64 `json:"qty"`
		RemainQty       float64 `json:"remain_qty"`
		Price           float64 `json:"price"`
		DiscountWord    string  `json:"discount_word"`
		DiscountAmount  float64 `json:"discount_amount"`
		UnitCode        string  `json:"unit_code"`
		ItemAmount      float64 `json:"item_amount"`
		ItemDescription string  `json:"item_description"`
		PackingRate1    float64 `json:"packing_rate_1"`
		IsCancel        int64   `json:"is_cancel"`
		LineNumber      int     `json:"line_number"`
	}

	NewSaleRequest struct {
		Id                  int64                `json:"id"`
		DocNo               string               `json:"doc_no"`
		DocDate             string               `json:"doc_date"`
		CompanyId           int64                `json:"company_id"`
		BranchId            int64                `json:"branch_id"`
		DocType             int64                `json:"doc_type"`
		ArId                int64                `json:"ar_id"`
		ArCode              string               `json:"ar_code"`
		ArName              string               `json:"ar_name"`
		ArBillAddress       string               `json:"ar_bill_address"`
		ArTelephone         string               `json:"ar_telephone"`
		SaleId              int                  `json:"sale_id"`
		SaleCode            string               `json:"sale_code"`
		SaleName            string               `json:"sale_name"`
		BillType            int64                `json:"bill_type"`
		TaxType             int                  `json:"tax_type"`
		TaxRate             float64              `json:"tax_rate"`
		DepartId            int64                `json:"depart_id"`
		RefNo               string               `json:"ref_no"`
		IsConfirm           int64                `json:"is_confirm"`
		BillStatus          int64                `json:"bill_status"`
		HoldingStatus       int64                `json:"holding_status"`
		CreditDay           int64                `json:"credit_day"`
		DueDate             string               `json:"due_date"`
		DeliveryDay         int64                `json:"delivery_day"`
		DeliveryDate        string               `json:"delivery_date"`
		IsConditionSend     int64                `json:"is_condition_send"`
		DeliveryAddressId   int64                `json:"delivery_address_id"`
		CarLicense          string               `json:"car_license"`
		PersonReceiveTel    string               `json:"person_receive_tel"`
		MyDescription       string               `json:"my_description"`
		SumOfItemAmount     float64              `json:"sum_of_item_amount"`
		DiscountWord        string               `json:"discount_word"`
		DiscountAmount      float64              `json:"discount_amount"`
		AfterDiscountAmount float64              `json:"after_discount_amount"`
		BeforeTaxAmount     float64              `json:"before_tax_amount"`
		TaxAmount           float64              `json:"tax_amount"`
		TotalAmount         float64              `json:"total_amount"`
		NetDebtAmount       float64              `json:"net_debt_amount"`
		ProjectId           int64                `json:"project_id"`
		AllocateId          int64                `json:"allocate_id"`
		JobId               string               `json:"job_id"`
		IsCancel            int64                `json:"is_cancel"`
		CreateBy            string               `json:"create_by"`
		CreateTime          string               `json:"create_time"`
		EditBy              string               `json:"edit_by"`
		EditTime            string               `json:"edit_time"`
		ConfirmBy           string               `json:"confirm_by"`
		ConfirmTime         string               `json:"confirm_time"`
		CancelBy            string               `json:"cancel_by"`
		CancelTime          string               `json:"cancel_time"`
		Subs                []NewSaleItemRequest `json:"subs"`
	}

	NewSaleItemRequest struct {
		Id              int64   `json:"id"`
		SOId            int64   `json:"so_id"`
		ItemId          int64   `json:"item_id"`
		ItemCode        string  `json:"item_code"`
		BarCode         string  `json:"bar_code"`
		ItemName        string  `json:"item_name"`
		WHCode          string  `json:"wh_code"`
		ShelfCode       string  `json:"shelf_code"`
		Qty             float64 `json:"qty"`
		RemainQty       float64 `json:"remain_qty"`
		Price           float64 `json:"price"`
		DiscountWord    string  `json:"discount_word"`
		DiscountAmount  float64 `json:"discount_amount"`
		UnitCode        string  `json:"unit_code"`
		ItemAmount      float64 `json:"item_amount"`
		ItemDescription string  `json:"item_description"`
		StockType       int64   `json:"stock_type"`
		AverageCost     float64 `json:"average_cost"`
		SumOfCost       float64 `json:"sum_of_cost"`
		PackingRate1    float64 `json:"packing_rate_1"`
		RefNo           string  `json:"ref_no"`
		QuoId           int64   `json:"quo_id"`
		LineNumber      int     `json:"line_number"`
		RefLineNumber   int64   `json:"ref_line_number"`
		IsCancel        int64   `json:"is_cancel"`
	}

	SearchByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchByKeywordRequest struct {
		SaleCode string `json:"sale_code"`
		Keyword  string `json:"keyword"`
	}

	SearchDocResponse struct {
		Id            int64   `json:"id"`
		DocNo         string  `json:"doc_no"`
		DocDate       string  `json:"doc_date"`
		Module        string  `json:"module"`
		ArCode        string  `json:"ar_code"`
		ArName        string  `json:"ar_name"`
		SaleCode      string  `json:"sale_code"`
		SaleName      string  `json:"sale_name"`
		MyDescription string  `json:"my_description"`
		TotalAmount   float64 `json:"total_amount"`
		IsCancel      int     `json:"is_cancel"`
		IsConfirm     int     `json:"is_confirm"`
	}

	NewDepositRequest struct {
		Id               int64                   `json:"id"`
		CompanyId        int64                   `json:"company_id"`
		BranchId         int64                   `json:"branch_id"`
		Uuid             string                  `json:"uuid"`
		DocNo            string                  `json:"doc_no"`
		TaxNo            string                  `json:"tax_no"`
		DocDate          string                  `json:"doc_date"`
		BillType         int64                   `json:"bill_type"`
		ArId             int64                   `json:"ar_id"`
		ArCode           string                  `json:"ar_code"`
		ArName           string                  `json:"ar_name"`
		ArBillAddress    string                  `json:"ar_bill_address"`
		ArTelephone      string                  `json:"ar_telephone"`
		SaleId           int64                   `json:"sale_id"`
		SaleCode         string                  `json:"sale_code"`
		SaleName         string                  `json:"sale_name"`
		TaxType          int64                   `json:"tax_type"`
		TaxRate          float64                 `json:"tax_rate"`
		RefNo            string                  `json:"ref_no"`
		CreditDay        int64                   `json:"credit_day"`
		DueDate          string                  `json:"due_date"`
		DepartId         int64                   `json:"depart_id"`
		AllocateId       int64                   `json:"allocate_id"`
		ProjectId        int64                   `json:"project_id"`
		MyDescription    string                  `json:"my_description"`
		BeforeTaxAmount  float64                 `json:"before_tax_amount"`
		TaxAmount        float64                 `json:"tax_amount"`
		TotalAmount      float64                 `json:"total_amount"`
		NetAmount        float64                 `json:"net_amount"`
		BillBalance      float64                 `json:"bill_balance"`
		CashAmount       float64                 `json:"cash_amount"`
		CreditcardAmount float64                 `json:"creditcard_amount"`
		ChqAmount        float64                 `json:"chq_amount"`
		BankAmount       float64                 `json:"bank_amount"`
		IsReturnMoney    int64                   `json:"is_return_money" `
		IsCancel         int64                   `json:"is_cancel"`
		IsConfirm        int64                   `json:"is_confirm"`
		ScgId            string                  `json:"scg_id"`
		JobNo            string                  `json:"job_no"`
		CreateBy         string                  `json:"create_by"`
		CreateTime       string                  `json:"create_time"`
		EditBy           string                  `json:"edit_by"`
		EditTime         string                  `json:"edit_time"`
		CancelBy         string                  `json:"cancel_by"`
		CancelTime       string                  `json:"cancel_time" `
		ConfirmBy        string                  `json:"confirm_by"`
		ConfirmTime      string                  `json:"confirm_time"`
		Subs             []NewDepositItemRequest `json:"subs"`
		CreditCard       []CreditCard            `json:"credit_card"`
		Chq              []ChqIn                 `json:"chq"`
	}

	NewDepositItemRequest struct {
		Id              int64   `json:"id"`
		SORefNo         string  `json:"so_ref_no"`
		SOId            int64   `json:"so_id"`
		ItemId          int64   `json:"item_id"`
		ItemCode        string  `json:"item_code"`
		BarCode         string  `json:"bar_code"`
		ItemName        string  `json:"item_name"`
		WHCode          string  `json:"wh_code"`
		ShelfCode       string  `json:"shelf_code"`
		Qty             float64 `json:"qty"`
		RemainQty       float64 `json:"remain_qty"`
		Price           float64 `json:"price"`
		DiscountWord    string  `json:"discount_word"`
		DiscountAmount  float64 `json:"discount_amount"`
		UnitCode        string  `json:"unit_code"`
		ItemAmount      float64 `json:"item_amount"`
		ItemDescription string  `json:"item_description"`
		PackingRate1    float64 `json:"packing_rate_1"`
		RefNo           string  `json:"ref_no"`
		QuoId           int64   `json:"quo_id"`
		LineNumber      int     `json:"line_number"`
		RefLineNumber   int64   `json:"ref_line_number"`
		IsCancel        int64   `json:"is_cancel"`
	}

	CreditCard struct {
		Id           int64   `json:"id"`
		RefId        int64   `json:"ref_id"`
		CreditCardNo string  `json:"credit_card_no"`
		CreditType   string  `json:"credit_type"`
		ConfirmNo    string  `json:"confirm_no"`
		Amount       float64 `json:"amount"`
		ChargeAmount float64 `json:"charge_amount"`
		Description  string  `json:"description"`
		BankId       int64   `json:"bank_id"`
		BankBranchId int64   `json:"bank_branch_id"`
		ReceiveDate  string  `json:"receive_date"`
		DueDate      string  `json:"due_date"`
		BookId       int64   `json:"book_id"`
	}

	ChqIn struct {
		Id           int64   `json:"id"`
		RefId        int64   `json:"ref_id"`
		ChqNumber    string  `json:"chq_number"`
		BankId       int64   `json:"bank_id"`
		BankBranchId int64   `json:"bank_branch_id"`
		ReceiveDate  string  `json:"receive_date"`
		DueDate      string  `json:"due_date"`
		BookId       int64   `json:"book_id"`
		ChqStatus    int64   `json:"chq_status"`
		ChqAmount    float64 `json:"chq_amount"`
		ChqBalance   float64 `json:"chq_balance"`
		Description  string  `json:"description"`
	}
)

////// Quotation /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateQuotation(s Service) interface{} {
	return func(ctx context.Context, req *NewQuoRequest) (interface{}, error) {
		q := map_quo_request(req)

		fmt.Println("p =")

		for _, subs := range req.Subs {
			itemline := map_quo_sub_request(subs)
			q.Subs = append(q.Subs, itemline)
		}
		resp, err := s.CreateQuotation(&NewQuoTemplate{
			Id:                  req.Id,
			DocType:             req.DocType,
			CompanyId:           req.CompanyId,
			BranchId:            req.BranchId,
			AssertStatus:        req.AssertStatus,
			BillStatus:          req.BillStatus,
			IsCancel:            req.IsCancel,
			IsConfirm:           req.IsConfirm,
			NetDebtAmount:       req.NetDebtAmount,
			ProjectId:           req.ProjectId,
			AllocateId:          req.AllocateId,
			Validity:            req.Validity,
			CreditDay:           req.CreditDay,
			ExpireCredit:        req.ExpireCredit,
			DeliveryDay:         req.DeliveryDay,
			DocNo:               req.DocNo,
			DocDate:             req.DocDate,
			BillType:            req.BillType,
			ArId:                req.ArId,
			ArName:              req.ArName,
			ArCode:              req.ArCode,
			SaleId:              req.SaleId,
			SaleCode:            req.SaleCode,
			SaleName:            req.SaleName,
			TaxType:             req.TaxType,
			TaxRate:             req.TaxRate,
			RefNo:               req.RefNo,
			DepartId:            req.DepartId,
			DueDate:             req.DueDate,
			ExpireDate:          req.ExpireDate,
			DeliveryDate:        req.DeliveryDate,
			IsConditionSend:     req.IsConditionSend,
			MyDescription:       req.MyDescription,
			SumOfItemAmount:     req.SumOfItemAmount,
			DiscountWord:        req.DiscountWord,
			DiscountAmount:      req.DiscountAmount,
			AfterDiscountAmount: req.AfterDiscountAmount,
			BeforeTaxAmount:     req.BeforeTaxAmount,
			TaxAmount:           req.TaxAmount,
			TotalAmount:         req.TotalAmount,
			CreateBy:            req.CreateBy,
			Subs:                q.Subs,
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

func map_quo_request(x *NewQuoRequest) NewQuoTemplate {
	var subs []NewQuoItemTemplate

	return NewQuoTemplate{
		AllocateId:          x.AllocateId,
		AssertStatus:        x.AssertStatus,
		CreditDay:           x.CreditDay,
		CreateTime:          x.CreateTime,
		DeliveryDay:         x.DeliveryDay,
		CompanyId:           x.CompanyId,
		BranchId:            x.BranchId,
		DocType:             x.DocType,
		ExpireCredit:        x.ExpireCredit,
		Id:                  x.Id,
		IsConfirm:           x.IsConfirm,
		IsCancel:            x.IsCancel,
		NetDebtAmount:       x.NetDebtAmount,
		ProjectId:           x.ProjectId,
		Validity:            x.Validity,
		DocNo:               x.DocNo,
		DocDate:             x.DocDate,
		BillType:            x.BillType,
		ArId:                x.ArId,
		ArName:              x.ArName,
		ArCode:              x.ArCode,
		SaleId:              x.SaleId,
		SaleCode:            x.SaleCode,
		SaleName:            x.SaleName,
		TaxType:             x.TaxType,
		TaxRate:             x.TaxRate,
		RefNo:               x.RefNo,
		DepartId:            x.DepartId,
		DueDate:             x.DueDate,
		ExpireDate:          x.ExpireDate,
		DeliveryDate:        x.DeliveryDate,
		IsConditionSend:     x.IsConditionSend,
		MyDescription:       x.MyDescription,
		SumOfItemAmount:     x.SumOfItemAmount,
		DiscountWord:        x.DiscountWord,
		DiscountAmount:      x.DiscountAmount,
		AfterDiscountAmount: x.AfterDiscountAmount,
		BeforeTaxAmount:     x.BeforeTaxAmount,
		TaxAmount:           x.TaxAmount,
		TotalAmount:         x.TotalAmount,
		CreateBy:            x.CreateBy,
		Subs:                subs,
	}
}

func map_quo_sub_request(x NewQuoItemRequest) NewQuoItemTemplate {
	return NewQuoItemTemplate{
		Id:              x.Id,
		ItemId:          x.ItemId,
		ItemCode:        x.ItemCode,
		BarCode:         x.BarCode,
		ItemName:        x.ItemName,
		Qty:             x.Qty,
		Price:           x.Price,
		DiscountWord:    x.DiscountWord,
		DiscountAmount:  x.DiscountAmount,
		UnitCode:        x.UnitCode,
		ItemAmount:      x.ItemAmount,
		ItemDescription: x.ItemDescription,
		PackingRate1:    x.PackingRate1,
		LineNumber:      x.LineNumber,
		IsCancel:        x.IsCancel,
	}
}

func SearchQuoById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchQueById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

//
//////// Sale Order /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateSaleOrder(s Service) interface{} {
	return func(ctx context.Context, req *NewSaleRequest) (interface{}, error) {
		so := map_sale_request(req)

		for _, subs := range req.Subs {
			itemline := map_sale_sub_request(subs)
			so.Subs = append(so.Subs, itemline)
		}
		resp, err := s.CreateSaleOrder(&NewSaleTemplate{
			Id:                  req.Id,
			DocType:             req.DocType,
			CompanyId:           req.CompanyId,
			BranchId:            req.BranchId,
			IsConfirm:           req.IsConfirm,
			IsCancel:            req.IsCancel,
			JobId:               req.JobId,
			DocNo:               req.DocNo,
			DocDate:             req.DocDate,
			BillType:            req.BillType,
			HoldingStatus:       req.HoldingStatus,
			AllocateId:          req.AllocateId,
			ArId:                req.ArId,
			ArName:              req.ArName,
			ArCode:              req.ArCode,
			BillStatus:          req.BillStatus,
			NetDebtAmount:       req.NetDebtAmount,
			ProjectId:           req.ProjectId,
			SaleId:              req.SaleId,
			SaleCode:            req.SaleCode,
			SaleName:            req.SaleName,
			TaxType:             req.TaxType,
			TaxRate:             req.TaxRate,
			RefNo:               req.RefNo,
			CarLicense:          req.CarLicense,
			CreditDay:           req.CreditDay,
			DeliveryAddressId:   req.DeliveryAddressId,
			DeliveryDay:         req.DeliveryDay,
			PersonReceiveTel:    req.PersonReceiveTel,
			DepartId:            req.DepartId,
			DueDate:             req.DueDate,
			DeliveryDate:        req.DeliveryDate,
			IsConditionSend:     req.IsConditionSend,
			MyDescription:       req.MyDescription,
			SumOfItemAmount:     req.SumOfItemAmount,
			DiscountWord:        req.DiscountWord,
			DiscountAmount:      req.DiscountAmount,
			AfterDiscountAmount: req.AfterDiscountAmount,
			BeforeTaxAmount:     req.BeforeTaxAmount,
			TaxAmount:           req.TaxAmount,
			TotalAmount:         req.TotalAmount,
			CreateBy:            req.CreateBy,
			Subs:                so.Subs,
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

func map_sale_request(x *NewSaleRequest) NewSaleTemplate {
	var subs []NewSaleItemTemplate

	return NewSaleTemplate{
		DocNo:               x.DocNo,
		DocDate:             x.DocDate,
		BillType:            x.BillType,
		ArId:                x.ArId,
		ArName:              x.ArName,
		ArCode:              x.ArCode,
		ArBillAddress:       x.ArBillAddress,
		ArTelephone:         x.ArTelephone,
		SaleId:              x.SaleId,
		SaleCode:            x.SaleCode,
		SaleName:            x.SaleName,
		TaxType:             x.TaxType,
		TaxRate:             x.TaxRate,
		RefNo:               x.RefNo,
		DepartId:            x.DepartId,
		DueDate:             x.DueDate,
		DeliveryDate:        x.DeliveryDate,
		IsConditionSend:     x.IsConditionSend,
		MyDescription:       x.MyDescription,
		SumOfItemAmount:     x.SumOfItemAmount,
		DiscountWord:        x.DiscountWord,
		DiscountAmount:      x.DiscountAmount,
		AfterDiscountAmount: x.AfterDiscountAmount,
		BeforeTaxAmount:     x.BeforeTaxAmount,
		TaxAmount:           x.TaxAmount,
		TotalAmount:         x.TotalAmount,
		CreateBy:            x.CreateBy,
		Subs:                subs,
	}
}

func map_sale_sub_request(x NewSaleItemRequest) NewSaleItemTemplate {
	return NewSaleItemTemplate{
		ItemCode:        x.ItemCode,
		BarCode:         x.BarCode,
		ItemName:        x.ItemName,
		WHCode:          x.WHCode,
		ShelfCode:       x.ShelfCode,
		Qty:             x.Qty,
		Price:           x.Price,
		DiscountWord:    x.DiscountWord,
		DiscountAmount:  x.DiscountAmount,
		UnitCode:        x.UnitCode,
		ItemAmount:      x.ItemAmount,
		ItemDescription: x.ItemDescription,
		PackingRate1:    x.PackingRate1,
		LineNumber:      x.LineNumber,
		IsCancel:        x.IsCancel,
	}
}

func SearchSaleOrderById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchSaleOrderById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchDocByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchDocByKeyword(&SearchByKeywordTemplate{SaleCode: req.SaleCode, Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

////// Deposit /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateDeposit(s Service) interface{} {
	return func(ctx context.Context, req *NewDepositRequest) (interface{}, error) {

		dp := map_deposit_request(req)

		for _, crds := range req.CreditCard {
			crdline := map_creditcard_request(crds)
			dp.CreditCard = append(dp.CreditCard, crdline)
		}

		for _, chqs := range req.Chq {
			chqline := map_chq_request(chqs)
			dp.Chq = append(dp.Chq, chqline)
		}

		resp, err := s.CreateDeposit(&NewDepositTemplate{
			Id:               req.Id,
			CompanyId:        req.CompanyId,
			BranchId:         req.BranchId,
			DocNo:            req.DocNo,
			TaxNo:            req.TaxNo,
			DocDate:          req.DocDate,
			BillType:         req.BillType,
			ArId:             req.ArId,
			ArCode:           req.ArCode,
			ArName:           req.ArName,
			ArBillAddress:    req.ArBillAddress,
			ArTelephone:      req.ArTelephone,
			SaleId:           req.SaleId,
			SaleCode:         req.SaleCode,
			SaleName:         req.SaleName,
			TaxType:          req.TaxType,
			TaxRate:          req.TaxRate,
			RefNo:            req.RefNo,
			CreditDay:        req.CreditDay,
			DueDate:          req.DueDate,
			DepartId:         req.DepartId,
			AllocateId:       req.AllocateId,
			ProjectId:        req.ProjectId,
			MyDescription:    req.MyDescription,
			BeforeTaxAmount:  req.BeforeTaxAmount,
			TaxAmount:        req.TaxAmount,
			TotalAmount:      req.TotalAmount,
			NetAmount:        req.NetAmount,
			BillBalance:      req.BillBalance,
			CashAmount:       req.CashAmount,
			CreditcardAmount: req.CreditcardAmount,
			ChqAmount:        req.ChqAmount,
			BankAmount:       req.BankAmount,
			IsReturnMoney:    req.IsReturnMoney,
			IsCancel:         req.IsCancel,
			IsConfirm:        req.IsConfirm,
			ScgId:            req.ScgId,
			JobNo:            req.JobNo,
			CreateBy:         req.CreateBy,
			CreateTime:       req.CreateTime,
			EditBy:           req.EditBy,
			Uuid:             req.Uuid,
			CreditCard:       dp.CreditCard,
			Chq:              dp.Chq,
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

func map_deposit_request(x *NewDepositRequest) NewDepositTemplate {
	var subs []NewDepositItemRequest
	var credit_cards []CreditCardTemplate
	var chqs []ChqInTemplate

	return NewDepositTemplate{
		Id:               x.Id,
		DocNo:            x.DocNo,
		DocDate:          x.DocDate,
		ArId:             x.ArId,
		ArTelephone:      x.ArTelephone,
		ArBillAddress:    x.ArBillAddress,
		ArName:           x.ArName,
		ArCode:           x.ArCode,
		AllocateId:       x.AllocateId,
		BranchId:         x.BranchId,
		BillBalance:      x.BillBalance,
		BankAmount:       x.BankAmount,
		BillType:         x.BillType,
		BeforeTaxAmount:  x.BeforeTaxAmount,
		CompanyId:        x.CompanyId,
		CreditcardAmount: x.CreditcardAmount,
		ChqAmount:        x.ChqAmount,
		CashAmount:       x.CashAmount,
		CreditDay:        x.CreditDay,
		CreateTime:       x.CreateTime,
		CreateBy:         x.CreateBy,
		CancelTime:       x.CancelTime,
		CancelBy:         x.CancelBy,
		ConfirmBy:        x.ConfirmBy,
		ConfirmTime:      x.ConfirmTime,
		DueDate:          x.DueDate,
		DepartId:         x.DepartId,
		EditBy:           x.EditBy,
		EditTime:         x.EditTime,
		IsReturnMoney:    x.IsReturnMoney,
		IsConfirm:        x.IsConfirm,
		IsCancel:         x.IsCancel,
		JobNo:            x.JobNo,
		MyDescription:    x.MyDescription,
		NetAmount:        x.NetAmount,
		ProjectId:        x.ProjectId,
		RefNo:            x.RefNo,
		SaleName:         x.SaleName,
		SaleCode:         x.SaleCode,
		SaleId:           x.SaleId,
		ScgId:            x.ScgId,
		TaxNo:            x.TaxNo,
		TotalAmount:      x.TotalAmount,
		TaxAmount:        x.TaxAmount,
		TaxRate:          x.TaxRate,
		TaxType:          x.TaxType,
		Uuid:             x.Uuid,
		Subs:             subs,
		CreditCard:       credit_cards,
		Chq:              chqs,
	}
}

func map_creditcard_request(x CreditCard) CreditCardTemplate {

	return CreditCardTemplate{
		BankBranchId: x.BankBranchId,
		BankId:       x.BankId,
		BookId:       x.BookId,
		CreditCardNo: x.CreditCardNo,
		Amount:       x.Amount,
		ChargeAmount: x.ChargeAmount,
		CreditType:   x.CreditType,
		ConfirmNo:    x.ConfirmNo,
		Description:  x.Description,
		DueDate:      x.DueDate,
		Id:           x.Id,
		ReceiveDate:  x.ReceiveDate,
		RefId:        x.RefId,
	}
}

func map_chq_request(x ChqIn) ChqInTemplate {
	return ChqInTemplate{
		BookId:       x.BookId,
		BankId:       x.BankId,
		BankBranchId: x.BankBranchId,
		ChqAmount:    x.ChqAmount,
		ChqBalance:   x.ChqBalance,
		ChqNumber:    x.ChqNumber,
		ChqStatus:    x.ChqStatus,
		DueDate:      x.DueDate,
		Id:           x.Id,
		RefId:        x.RefId,
		ReceiveDate:  x.ReceiveDate,
		Description:  x.Description,
	}
}

func SearchDepositById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchDepositById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchDepositByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchDepositByKeyword(&SearchByKeywordTemplate{Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
