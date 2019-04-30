package endpoint

//import (
//	"github.com/mrtomyum/nopadol/pos"
//	"context"
//	"fmt"
//)
//
//func New(s pos.Service) pos.Endpoint {
//	return &endpoint{s}
//}
//
//type endpoint struct {
//	s pos.Service
//}
//
//func (ep *endpoint) New(ctx context.Context, req pos.NewPosRequest) (resp *pos.NewPosResponse, err error) {
//	p := map_pos_request(req)
//	for _, subs := range req.PosSubs {
//		itemline := map_pos_sub_request(subs)
//		p.PosSubs = append(p.PosSubs, itemline)
//	}
//	for _, creditcards := range req.CreditCards {
//		creditcardline := map_pos_creditcard_request(creditcards)
//		p.CreditCards = append(p.CreditCards, creditcardline)
//	}
//	for _, chqins := range req.ChqIns {
//		chqline := map_pos_chq_request(chqins)
//		p.ChqIns = append(p.ChqIns, chqline)
//	}
//
//	posdata, err := ep.s.New(ctx, &pos.NewPosTemplate{
//		DocNo:           req.DocNo,
//		DocDate:         req.DocDate,
//		ArCode:          req.ArCode,
//		SaleCode:        req.SaleCode,
//		ShiftNo:         req.ShiftNo,
//		ShiftCode:       req.ShiftCode,
//		MachineNo:       req.MachineNo,
//		MachineCode:     req.MachineCode,
//		CashierCode:     req.CashierCode,
//		CoupongAmount:   req.CoupongAmount,
//		ChangeAmount:    req.ChangeAmount,
//		ChargeAmount:    req.ChargeAmount,
//		TaxType:         req.TaxType,
//		SumOfItemAmount: req.SumOfItemAmount,
//		DiscountWord:    req.DiscountWord,
//		AfterDiscount:   req.AfterDiscount,
//		TotalAmount:     req.TotalAmount,
//		SumCashAmount:   req.SumCashAmount,
//		SumChqAmount:    req.SumChqAmount,
//		SumCreditAmount: req.SumCreditAmount,
//		SumBankAmount:   req.SumBankAmount,
//		NetDebtAmount:   req.NetDebtAmount,
//		UserCode:        req.UserCode,
//		PosSubs:         p.PosSubs,
//		CreditCards:     p.CreditCards,
//		ChqIns:          p.ChqIns,
//	})
//
//	x := map_pos_response(posdata)
//
//	return &pos.NewPosResponse{Id: x.Id}, nil
//}
//
//func (ep *endpoint) SearchById(ctx context.Context, req *pos.SearchPosByIdRequest) (*pos.SearchPosByIdResponse, error) {
//	data, err := ep.s.SearchById(ctx, &pos.SearchPosByIdRequestTemplate{
//		Id: req.Id,
//	})
//	if err != nil {
//		fmt.Println("error = ", err.Error())
//		return nil, err
//	}
//
//	resp := map_search_pos_response(data)
//
//	fmt.Println("DocNo =", resp.DocNo)
//	return &resp, nil
//}

//func map_search_pos_response(x pos.SearchPosByIdResponseTemplate) pos.SearchPosByIdResponse {
//	return pos.SearchPosByIdResponse{
//		Id:              x.Id,
//		DocNo:           x.DocNo,
//		DocDate:         x.DocDate,
//		TaxNo:           x.TaxNo,
//		PosStatus:       x.PosStatus,
//		ArCode:          x.ArCode,
//		ArName:          x.ArName,
//		SaleCode:        x.SaleCode,
//		SaleName:        x.SaleName,
//		CashierCode:     x.CashierCode,
//		ShiftNo:         x.ShiftNo,
//		ShiftCode:       x.ShiftCode,
//		MachineNo:       x.MachineNo,
//		MachineCode:     x.MachineCode,
//		CoupongAmount:   x.CoupongAmount,
//		ChangeAmount:    x.ChangeAmount,
//		ChargeAmount:    x.ChargeAmount,
//		TaxType:         x.TaxType,
//		SumOfItemAmount: x.SumOfItemAmount,
//		DiscountWord:    x.DiscountWord,
//		AfterDiscount:   x.AfterDiscount,
//		BeforeTaxAmount: x.BeforeTaxAmount,
//		TaxAmount:       x.TaxAmount,
//		TotalAmount:     x.TotalAmount,
//		SumCashAmount:   x.SumCashAmount,
//		SumChqAmount:    x.SumChqAmount,
//		SumCreditAmount: x.SumCreditAmount,
//		SumBankAmount:   x.SumBankAmount,
//		NetDebtAmount:   x.NetDebtAmount,
//		IsCancel:        x.IsCancel,
//		IsConfirm:       x.IsConfirm,
//		CreatorCode:     x.CreatorCode,
//		CreateDateTime:  x.CreateDateTime,
//		LastEditorCode:  x.LastEditorCode,
//		LastEditDateT:   x.LastEditDateT,
//	}
//}
//
//func map_pos_request(x pos.NewPosRequest) pos.NewPosTemplate {
//	var subs []pos.NewPosItemTemplate
//	var creditcards []pos.ListCreditCardTemplate
//	var chqs []pos.ListChqInTemplate
//	return pos.NewPosTemplate{
//		DocNo:           x.DocNo,
//		DocDate:         x.DocDate,
//		ArCode:          x.ArCode,
//		SaleCode:        x.SaleCode,
//		ShiftNo:         x.ShiftNo,
//		ShiftCode:       x.ShiftCode,
//		MachineNo:       x.MachineNo,
//		MachineCode:     x.MachineCode,
//		CoupongAmount:   x.CoupongAmount,
//		ChangeAmount:    x.ChangeAmount,
//		ChargeAmount:    x.ChargeAmount,
//		TaxType:         x.TaxType,
//		SumOfItemAmount: x.SumOfItemAmount,
//		DiscountWord:    x.DiscountWord,
//		AfterDiscount:   x.AfterDiscount,
//		TotalAmount:     x.TotalAmount,
//		SumCashAmount:   x.SumCashAmount,
//		SumChqAmount:    x.SumChqAmount,
//		SumCreditAmount: x.SumCreditAmount,
//		SumBankAmount:   x.SumBankAmount,
//		NetDebtAmount:   x.NetDebtAmount,
//		UserCode:        x.UserCode,
//		PosSubs:         subs,
//		CreditCards:     creditcards,
//		ChqIns:          chqs,
//	}
//}
//
//func map_pos_sub_request(x pos.NewPosItemRequest) pos.NewPosItemTemplate {
//	return pos.NewPosItemTemplate{
//		ItemCode:     x.ItemCode,
//		ItemName:     x.ItemName,
//		WHCode:       x.WHCode,
//		ShelfCode:    x.ShelfCode,
//		Qty:          x.Qty,
//		Price:        x.Price,
//		DiscountWord: x.DiscountWord,
//		UnitCode:     x.UnitCode,
//		BarCode:      x.BarCode,
//		AverageCost:  x.AverageCost,
//		PackingRate1: x.PackingRate1,
//		LineNumber:   x.LineNumber,
//	}
//}
//
//func map_pos_creditcard_request(x pos.ListCreditCardRequest) pos.ListCreditCardTemplate {
//	return pos.ListCreditCardTemplate{
//		BankCode:       x.BankCode,
//		CreditCardNo:   x.CreditCardNo,
//		ReceiveDate:    x.ReceiveDate,
//		DueDate:        x.DueDate,
//		BookNo:         x.BookNo,
//		Status:         x.Status,
//		StatusDate:     x.StatusDate,
//		StatusDocNo:    x.StatusDocNo,
//		BankBranchCode: x.BankBranchCode,
//		Amount:         x.Amount,
//		MyDescription:  x.MyDescription,
//		CreditType:     x.CreditType,
//		ConfirmNo:      x.ConfirmNo,
//		ChargeAmount:   x.ChargeAmount,
//	}
//}
//
//func map_pos_chq_request(x pos.ListChqInRequest) pos.ListChqInTemplate {
//	return pos.ListChqInTemplate{
//		ChqNumber:      x.ChqNumber,
//		BankCode:       x.BankCode,
//		BankBranchCode: x.BankBranchCode,
//		BookNo:         x.BookNo,
//		ReceiveDate:    x.ReceiveDate,
//		DueDate:        x.DueDate,
//		Status:         x.Status,
//		Amount:         x.Amount,
//		Balance:        x.Balance,
//		RefChqRowOrder: x.RefChqRowOrder,
//		StatusDate:     x.StatusDate,
//		StatusDocNo:    x.StatusDocNo,
//	}
//}
//
//func map_pos_response(x pos.NewPosResponseTemplate) pos.NewPosResponse {
//	return pos.NewPosResponse{Id: x.Id}
//}
