package pos

import (
	"context"
	"fmt"
)

//type Endpoint interface {
//	New(context.Context, NewPosRequest) (*NewPosResponse, error)
//	SearchById(context.Context, *SearchPosByIdRequest) (*SearchPosByIdResponse, error)
//}

type (
	NewPosRequest struct {
		DocNo           string                  `json:"doc_no"`
		DocDate         string                  `json:"doc_date"`
		ArCode          string                  `json:"ar_code"`
		SaleCode        string                  `json:"sale_code"`
		ShiftCode       string                  `json:"shiftcode"`
		CashierCode     string                  `json:"cashier_code"`
		ShiftNo         int                     `json:"shift_no"`
		MachineNo       string                  `json:"machine_no"`
		MachineCode     string                  `json:"machine_code"`
		CoupongAmount   float64                 `json:"coupong_amount"`
		ChangeAmount    float64                 `json:"change_amount"`
		ChargeAmount    float64                 `json:"charge_amount"`
		TaxType         int                     `json:"tax_type"`
		SumOfItemAmount float64                 `json:"sum_of_item_amount"`
		DiscountWord    string                  `json:"discount_word"`
		AfterDiscount   float64                 `json:"after_discount"`
		TotalAmount     float64                 `json:"total_amount"`
		SumCashAmount   float64                 `json:"sum_cash_amount"`
		SumChqAmount    float64                 `json:"sum_chq_amount"`
		SumCreditAmount float64                 `json:"sum_credit_amount"`
		SumBankAmount   float64                 `json:"sum_bank_amount"`
		BankNo          string                  `json:"bank_no"`
		NetDebtAmount   float64                 `json:"net_debt_amount"`
		UserCode        string                  `json:"user_code"`
		ChqIns          []ListChqInRequest      `json:"chq_ins"`
		CreditCards     []ListCreditCardRequest `json:"credit_cards"`
		PosSubs         []NewPosItemRequest     `json:"pos_subs"`
	}

	NewPosItemRequest struct {
		ItemCode     string  `json:"item_code"`
		ItemName     string  `json:"item_name"`
		WHCode       string  `json:"wh_code"`
		ShelfCode    string  `json:"shelf_code"`
		Qty          float64 `json:"qty"`
		Price        float64 `json:"price"`
		DiscountWord string  `json:"discount_word"`
		UnitCode     string  `json:"unit_code"`
		LineNumber   int     `json:"line_number"`
		BarCode      string  `json:"bar_code"`
		AverageCost  float64 `json:"averagecost"`
		PackingRate1 float64 `json:"packing_rate_1"`
	}

	ListChqInRequest struct {
		ChqNumber      string  `json:"chq_number"`
		BankCode       string  `json:"bank_code"`
		BankBranchCode string  `json:"bank_branch_code"`
		BookNo         string  `json:"book_no"`
		ReceiveDate    string  `json:"receive_date"`
		DueDate        string  `json:"due_date"`
		Status         int     `json:"status"`
		Amount         float64 `json:"amount"`
		Balance        float64 `json:"balance"`
		RefChqRowOrder int     `json:"ref_chq_row_order"`
		StatusDate     string  `json:"status_date"`
		StatusDocNo    string  `json:"status_doc_no"`
	}

	ListCreditCardRequest struct {
		BankCode       string  `json:"bank_code"`
		CreditCardNo   string  `json:"credit_card_no"`
		ReceiveDate    string  `json:"receive_date"`
		DueDate        string  `json:"due_date"`
		BookNo         string  `json:"book_no"`
		Status         int     `json:"status"`
		StatusDate     string  `json:"status_date"`
		StatusDocNo    string  `json:"status_doc_no"`
		BankBranchCode string  `json:"bank_branch_code"`
		Amount         float64 `json:"amount"`
		MyDescription  string  `json:"my_description"`
		CreditType     string  `json:"credit_type"`
		ConfirmNo      string  `json:"confirm_no"`
		ChargeAmount   float64 `json:"charge_amount"`
	}

	NewPosResponse struct {
		Id int64 `json:"id"`
	}

	SearchPosByIdResponse struct {
		Id              int                     `json:"id"`
		DocNo           string                  `json:"doc_no"`
		DocDate         string                  `json:"doc_date"`
		TaxNo           string                  `json:"tax_no"`
		TaxDate         string                  `json:"tax_date"`
		PosStatus       int                     `json:"pos_status"`
		ArCode          string                  `json:"ar_code"`
		ArName          string                  `json:"ar_name"`
		SaleCode        string                  `json:"sale_code"`
		SaleName        string                  `json:"sale_name"`
		ShiftCode       string                  `json:"shiftcode"`
		CashierCode     string                  `json:"cashier_code"`
		ShiftNo         int                     `json:"shift_no"`
		MachineNo       string                  `json:"machine_no"`
		MachineCode     string                  `json:"machine_code"`
		CoupongAmount   float64                 `json:"coupong_amount"`
		ChangeAmount    float64                 `json:"change_amount"`
		ChargeAmount    float64                 `json:"charge_amount"`
		TaxType         int                     `json:"tax_type"`
		SumOfItemAmount float64                 `json:"sum_of_item_amount"`
		DiscountWord    string                  `json:"discount_word"`
		AfterDiscount   float64                 `json:"after_discount"`
		BeforeTaxAmount float64                 `json:"before_tax_amount"`
		TaxAmount       float64                 `json:"tax_amount"`
		TotalAmount     float64                 `json:"total_amount"`
		SumCashAmount   float64                 `json:"sum_cash_amount"`
		SumChqAmount    float64                 `json:"sum_chq_amount"`
		SumCreditAmount float64                 `json:"sum_credit_amount"`
		SumBankAmount   float64                 `json:"sum_bank_amount"`
		BankNo          string                  `json:"bank_no"`
		NetDebtAmount   float64                 `json:"net_debt_amount"`
		IsCancel        int                     `json:"is_cancel"`
		IsConfirm       int                     `json:"is_confirm"`
		CreatorCode     string                  `json:"creator_code"`
		CreateDateTime  string                  `json:"create_date_time"`
		LastEditorCode  string                  `json:"last_editor_code"`
		LastEditDateT   string                  `json:"last_edit_date_t"`
		ChqIns          []ListChqInRequest      `json:"chq_ins"`
		CreditCards     []ListCreditCardRequest `json:"credit_cards"`
		PosSubs         []NewPosItemRequest     `json:"pos_subs"`
	}

	SearchPosByIdRequest struct {
		Id int64 `json:"id"`
	}

	SearchPosByKeywordRequest struct {
		Keyword string `json:"keyword"`
	}
)

func Create(s Service) interface{} {
	fmt.Println("Endpoint")
	return func(ctx context.Context, req *NewPosRequest) (interface{}, error) {
		p := map_pos_request(req)

		fmt.Println("p =")

		for _, subs := range req.PosSubs {
			itemline := map_pos_sub_request(subs)
			p.PosSubs = append(p.PosSubs, itemline)
		}
		for _, creditcards := range req.CreditCards {
			creditcardline := map_pos_creditcard_request(creditcards)
			p.CreditCards = append(p.CreditCards, creditcardline)
		}
		for _, chqins := range req.ChqIns {
			chqline := map_pos_chq_request(chqins)
			p.ChqIns = append(p.ChqIns, chqline)
		}

		resp, err := s.Create(&NewPosTemplate{
			DocNo:           req.DocNo,
			DocDate:         req.DocDate,
			ArCode:          req.ArCode,
			SaleCode:        req.SaleCode,
			ShiftNo:         req.ShiftNo,
			ShiftCode:       req.ShiftCode,
			MachineNo:       req.MachineNo,
			MachineCode:     req.MachineCode,
			CashierCode:     req.CashierCode,
			CoupongAmount:   req.CoupongAmount,
			ChangeAmount:    req.ChangeAmount,
			ChargeAmount:    req.ChargeAmount,
			TaxType:         req.TaxType,
			SumOfItemAmount: req.SumOfItemAmount,
			DiscountWord:    req.DiscountWord,
			AfterDiscount:   req.AfterDiscount,
			TotalAmount:     req.TotalAmount,
			SumCashAmount:   req.SumCashAmount,
			SumChqAmount:    req.SumChqAmount,
			SumCreditAmount: req.SumCreditAmount,
			SumBankAmount:   req.SumBankAmount,
			NetDebtAmount:   req.NetDebtAmount,
			UserCode:        req.UserCode,
			PosSubs:         p.PosSubs,
			CreditCards:     p.CreditCards,
			ChqIns:          p.ChqIns,
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
	fmt.Println("EndPoint")
	return func(ctx context.Context, req *SearchPosByIdRequest)(interface{}, error){
		resp , err := s.SearchById(&SearchPosByIdRequestTemplate{Id:req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string] interface{}{
			"data": resp,
		},nil
	}
}

func map_search_pos_response(x SearchPosByIdResponseTemplate) SearchPosByIdResponse {
	return SearchPosByIdResponse{
		Id:              x.Id,
		DocNo:           x.DocNo,
		DocDate:         x.DocDate,
		TaxNo:           x.TaxNo,
		PosStatus:       x.PosStatus,
		ArCode:          x.ArCode,
		ArName:          x.ArName,
		SaleCode:        x.SaleCode,
		SaleName:        x.SaleName,
		CashierCode:     x.CashierCode,
		ShiftNo:         x.ShiftNo,
		ShiftCode:       x.ShiftCode,
		MachineNo:       x.MachineNo,
		MachineCode:     x.MachineCode,
		CoupongAmount:   x.CoupongAmount,
		ChangeAmount:    x.ChangeAmount,
		ChargeAmount:    x.ChargeAmount,
		TaxType:         x.TaxType,
		SumOfItemAmount: x.SumOfItemAmount,
		DiscountWord:    x.DiscountWord,
		AfterDiscount:   x.AfterDiscount,
		BeforeTaxAmount: x.BeforeTaxAmount,
		TaxAmount:       x.TaxAmount,
		TotalAmount:     x.TotalAmount,
		SumCashAmount:   x.SumCashAmount,
		SumChqAmount:    x.SumChqAmount,
		SumCreditAmount: x.SumCreditAmount,
		SumBankAmount:   x.SumBankAmount,
		NetDebtAmount:   x.NetDebtAmount,
		IsCancel:        x.IsCancel,
		IsConfirm:       x.IsConfirm,
		CreatorCode:     x.CreatorCode,
		CreateDateTime:  x.CreateDateTime,
		LastEditorCode:  x.LastEditorCode,
		LastEditDateT:   x.LastEditDateT,
	}
}

func map_pos_request(x *NewPosRequest) NewPosTemplate {
	fmt.Println("p1 =")
	var subs []NewPosItemTemplate
	var creditcards []ListCreditCardTemplate
	var chqs []ListChqInTemplate
	return NewPosTemplate{
		DocNo:           x.DocNo,
		DocDate:         x.DocDate,
		ArCode:          x.ArCode,
		SaleCode:        x.SaleCode,
		ShiftNo:         x.ShiftNo,
		ShiftCode:       x.ShiftCode,
		MachineNo:       x.MachineNo,
		MachineCode:     x.MachineCode,
		CoupongAmount:   x.CoupongAmount,
		ChangeAmount:    x.ChangeAmount,
		ChargeAmount:    x.ChargeAmount,
		TaxType:         x.TaxType,
		SumOfItemAmount: x.SumOfItemAmount,
		DiscountWord:    x.DiscountWord,
		AfterDiscount:   x.AfterDiscount,
		TotalAmount:     x.TotalAmount,
		SumCashAmount:   x.SumCashAmount,
		SumChqAmount:    x.SumChqAmount,
		SumCreditAmount: x.SumCreditAmount,
		SumBankAmount:   x.SumBankAmount,
		NetDebtAmount:   x.NetDebtAmount,
		UserCode:        x.UserCode,
		PosSubs:         subs,
		CreditCards:     creditcards,
		ChqIns:          chqs,
	}
}

func map_pos_sub_request(x NewPosItemRequest) NewPosItemTemplate {
	return NewPosItemTemplate{
		ItemCode:     x.ItemCode,
		ItemName:     x.ItemName,
		WHCode:       x.WHCode,
		ShelfCode:    x.ShelfCode,
		Qty:          x.Qty,
		Price:        x.Price,
		DiscountWord: x.DiscountWord,
		UnitCode:     x.UnitCode,
		BarCode:      x.BarCode,
		AverageCost:  x.AverageCost,
		PackingRate1: x.PackingRate1,
		LineNumber:   x.LineNumber,
	}
}

func map_pos_creditcard_request(x ListCreditCardRequest) ListCreditCardTemplate {
	return ListCreditCardTemplate{
		BankCode:       x.BankCode,
		CreditCardNo:   x.CreditCardNo,
		ReceiveDate:    x.ReceiveDate,
		DueDate:        x.DueDate,
		BookNo:         x.BookNo,
		Status:         x.Status,
		StatusDate:     x.StatusDate,
		StatusDocNo:    x.StatusDocNo,
		BankBranchCode: x.BankBranchCode,
		Amount:         x.Amount,
		MyDescription:  x.MyDescription,
		CreditType:     x.CreditType,
		ConfirmNo:      x.ConfirmNo,
		ChargeAmount:   x.ChargeAmount,
	}
}

func map_pos_chq_request(x ListChqInRequest) ListChqInTemplate {
	return ListChqInTemplate{
		ChqNumber:      x.ChqNumber,
		BankCode:       x.BankCode,
		BankBranchCode: x.BankBranchCode,
		BookNo:         x.BookNo,
		ReceiveDate:    x.ReceiveDate,
		DueDate:        x.DueDate,
		Status:         x.Status,
		Amount:         x.Amount,
		Balance:        x.Balance,
		RefChqRowOrder: x.RefChqRowOrder,
		StatusDate:     x.StatusDate,
		StatusDocNo:    x.StatusDocNo,
	}
}

func map_pos_response(x NewPosResponseTemplate) NewPosResponse {
	return NewPosResponse{Id: x.Id}
}
