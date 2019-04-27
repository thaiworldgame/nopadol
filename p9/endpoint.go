package p9
//
//import (
//	"golang.org/x/net/context"
//	"fmt"
//)
//
//type (
//	Basket struct {
//		Id                  int64       `json:"id" db:"id"`
//		CompanyId           int64       `json:"company_id"`
//		BranchId            int64       `json:"branch_id"`
//		Uuid                string      `json:"uuid"`
//		InvoiceNo           string      `json:"invoice_no"`
//		TaxNo               string      `json:"tax_no"`
//		QueId               int64       `json:"que_id"`
//		DocType             int64       `json:"doc_type"`
//		DocDate             string      `json:"doc_date"`
//		ArId                int64       `json:"ar_id"`
//		SaleId              int64       `json:"sale_id"`
//		PosMachineId        int64       `json:"pos_machine_id"`
//		PeriodId            int64       `json:"period_id"`
//		CashId              int64       `json:"cash_id"`
//		PosStatus           int64       `json:"pos_status"`
//		TaxType             int64       `json:"tax_type"`
//		TaxRate             int64       `json:"tax_rate"`
//		NumberOfItem        int64       `json:"number_of_item"`
//		ChangeAmount        float64     `json:"change_amount"`
//		CashAmount          float64     `json:"cash_amount"`
//		CreditCardAmount    float64     `json:"credit_card_amount"`
//		ChqAmount           float64     `json:"chq_amount"`
//		BankAmount          float64     `json:"bank_amount"`
//		DepositAmount       float64     `json:"deposit_amount"`
//		OnlineAmount        float64     `json:"online_amount"`
//		CouponAmount        float64     `json:"coupon_amount"`
//		CreditAmount        float64     `json:"credit_amount"`
//		SumItemAmount       float64     `json:"sum_item_amount"`
//		DiscountWord        string      `json:"discount_word"`
//		DiscountAmount      float64     `json:"discount_amount"`
//		AfterDiscountAmount float64     `json:"after_discount_amount"`
//		BeforeTaxAmount     float64     `json:"before_tax_amount"`
//		TaxAmount           float64     `json:"tax_amount"`
//		TotalAmount         float64     `json:"total_amount"`
//		NetAmount           float64     `json:"net_amount"`
//		BillBalance         float64     `json:"bill_balance"`
//		OtpPassword         string      `json:"otp_password"`
//		Status              int64       `json:"status"`
//		PickStatus          int64       `json:"pick_status"`
//		HoldingStatus       int64       `json:"holding_status"`
//		DeliveryStatus      int64       `json:"delivery_status"`
//		ReceiveName         string      `json:"receive_name"`
//		ReceiveTel          string      `json:"receive_tel"`
//		IsPosted            int64       `json:"is_posted"`
//		IsReturn            int64       `json:"is_return"`
//		GLStatus            int64       `json:"gl_status"`
//		ScgId               string      `json:"scg_id"`
//		CreateBy            string      `json:"create_by"`
//		CreateTime          string      `json:"create_time"`
//		EditBy              string      `json:"edit_by"`
//		EditTime            string      `json:"edit_time"`
//		ConfirmBy           string      `json:"confirm_by"`
//		ConfirmTime         string      `json:"confirm_time"`
//		CancelBy            string      `json:"cancel_by"`
//		CancelTime          string      `json:"cancel_time"`
//		CancelDescId        int64       `json:"cancel_desc_id"`
//		CancelDesc          string      `json:"cancel_desc"`
//		Sub                 []BasketSub `json:"sub"`
//	}
//
//	BasketSub struct {
//		Id              int64   `json:"id"`
//		PosId           int64   `json:"pos_id"`
//		Uuid            string  `json:"uuid"`
//		QueId           int64   `json:"que_id"`
//		DocDate         string  `json:"doc_date"`
//		ItemId          int64   `json:"item_id"`
//		ItemCode        string  `json:"item_code"`
//		ItemName        string  `json:"item_name"`
//		BarCode         string  `json:"bar_code"`
//		WhId            int64   `json:"wh_id"`
//		ShelfId         int64   `json:"shelf_id"`
//		RequestQty      float64 `json:"request_qty"`
//		PickQty         float64 `json:"pick_qty"`
//		CheckoutQty     float64 `json:"checkout_qty"`
//		Price           float64 `json:"price"`
//		UnitId          int64   `json:"unit_id"`
//		PickAmount      float64 `json:"pick_amount"`
//		CheckoutAmount  float64 `json:"checkout_amount"`
//		Qty             float64 `json:"qty"`
//		RemainQty       float64 `json:"remain_qty"`
//		IsReturn        int64   `json:"is_return"`
//		Rate1           int64   `json:"rate_1"`
//		RefNo           string  `json:"ref_no"`
//		SaleId          int64   `json:"sale_id"`
//		AverageCost     float64 `json:"average_cost"`
//		SumOfCost       float64 `json:"sum_of_cost"`
//		DeliveryOrderId int64   `json:"delivery_order_id"`
//		RefLineNumber   int64   `json:"ref_line_number"`
//		LineNumber      int64   `json:"line_number"`
//		RequestBy       string  `json:"request_by"`
//		RequestTime     string  `json:"request_time"`
//		PickBy          string  `json:"pick_by"`
//		PickTime        string  `json:"pick_time"`
//		CheckoutBy      string  `json:"checkout_by"`
//		CheckoutTime    string  `json:"checkout_time"`
//	}
//)
//
//func Create(s Service) interface{} {
//	return func(ctx context.Context, req *Basket) (interface{}, error) {
//		b := map_basket_request(req)
//
//		for _, subs := range req.Sub {
//			itemline := map_basketsub_request(subs)
//			b.Sub = append(b.Sub, itemline)
//		}
//		resp, err := s.Create(&BasketTemplate{
//			Id:                  req.Id,
//			AfterDiscountAmount: req.AfterDiscountAmount,
//			ArId:                req.ArId,
//			BeforeTaxAmount:     req.BeforeTaxAmount,
//			BankAmount:          req.BankAmount,
//			BillBalance:         req.BillBalance,
//			BranchId:            req.BranchId,
//			CashId:              req.CashId,
//			ChangeAmount:        req.ChangeAmount,
//			CashAmount:          req.CashAmount,
//			ChqAmount:           req.ChqAmount,
//			CompanyId:           req.CompanyId,
//			CouponAmount:        req.CouponAmount,
//			CreditAmount:        req.CreditAmount,
//			CreditCardAmount:    req.CreditCardAmount,
//			CreateBy:            req.CreateBy,
//			DocType:             req.DocType,
//			DocDate:             req.DocDate,
//			DiscountAmount:      req.DiscountAmount,
//			DiscountWord:        req.DiscountWord,
//			DeliveryStatus:      req.DeliveryStatus,
//			DepositAmount:       req.DepositAmount,
//			GLStatus:            req.GLStatus,
//			HoldingStatus:       req.HoldingStatus,
//			InvoiceNo:           req.InvoiceNo,
//			IsPosted:            req.IsPosted,
//			IsReturn:            req.IsReturn,
//			NetAmount:           req.NetAmount,
//			NumberOfItem:        req.NumberOfItem,
//			OnlineAmount:        req.OnlineAmount,
//			OtpPassword:         req.OtpPassword,
//			PickStatus:          req.PickStatus,
//			PosStatus:           req.PosStatus,
//			PeriodId:            req.PeriodId,
//			PosMachineId:        req.PosMachineId,
//			QueId:               req.QueId,
//			ReceiveName:         req.ReceiveName,
//			ReceiveTel:          req.ReceiveTel,
//			SaleId:              req.SaleId,
//			Status:              req.Status,
//			ScgId:               req.ScgId,
//			SumItemAmount:       req.SumItemAmount,
//			TaxRate:             req.TaxRate,
//			TaxType:             req.TaxType,
//			TaxNo:               req.TaxNo,
//			TotalAmount:         req.TotalAmount,
//			TaxAmount:           req.TaxAmount,
//			Uuid:                req.Uuid,
//			Sub:                 b.Sub,
//		})
//		if err != nil {
//			fmt.Println("endpoint error =", err.Error())
//			return nil, fmt.Errorf(err.Error())
//		}
//
//		return map[string]interface{}{
//			"data": resp,
//		}, nil
//	}
//}
//
//func ManageBasket(s Service) interface{} {
//	return func(ctx context.Context, req *BasketSub) (interface{}, error) {
//
//		resp, err := s.ManageBasket(&BasketSubTemplate{
//			Id:              req.Id,
//			AverageCost:     req.AverageCost,
//			BarCode:         req.BarCode,
//			CheckoutAmount:  req.CheckoutAmount,
//			CheckoutBy:      req.CheckoutBy,
//			CheckoutQty:     req.CheckoutQty,
//			CheckoutTime:    req.CheckoutTime,
//			DeliveryOrderId: req.DeliveryOrderId,
//			DocDate:         req.DocDate,
//			IsReturn:        req.IsReturn,
//			ItemName:        req.ItemName,
//			ItemCode:        req.ItemCode,
//			ItemId:          req.ItemId,
//			LineNumber:      req.LineNumber,
//			PickBy:          req.PickBy,
//			PosId:           req.PosId,
//			PickAmount:      req.PickAmount,
//			Price:           req.Price,
//			PickQty:         req.PickQty,
//			PickTime:        req.PickTime,
//			Qty:             req.Qty,
//			QueId:           req.QueId,
//			Rate1:           req.Rate1,
//			RemainQty:       req.RemainQty,
//			RefNo:           req.RefNo,
//			RefLineNumber:   req.RefLineNumber,
//			RequestBy:       req.RequestBy,
//			RequestQty:      req.RequestQty,
//			RequestTime:     req.RequestTime,
//			SaleId:          req.SaleId,
//			SumOfCost:       req.SumOfCost,
//			ShelfId:         req.ShelfId,
//			Uuid:            req.Uuid,
//			WhId:            req.WhId,
//		})
//		if err != nil {
//			fmt.Println("endpoint error =", err.Error())
//			return nil, fmt.Errorf(err.Error())
//		}
//
//		return map[string]interface{}{
//			"data": resp,
//		}, nil
//	}
//}
//
//func map_basket_request(x *Basket) BasketTemplate {
//	var subs []BasketSubTemplate
//	return BasketTemplate{
//		Id:                  x.Id,
//		ArId:                x.ArId,
//		AfterDiscountAmount: x.AfterDiscountAmount,
//		BillBalance:         x.BillBalance,
//		BankAmount:          x.BankAmount,
//		BeforeTaxAmount:     x.BeforeTaxAmount,
//		BranchId:            x.BranchId,
//		CreditAmount:        x.CreditAmount,
//		CouponAmount:        x.CouponAmount,
//		ChqAmount:           x.ChqAmount,
//		CashAmount:          x.CashAmount,
//		ChangeAmount:        x.ChangeAmount,
//		CashId:              x.CashId,
//		CompanyId:           x.CompanyId,
//		CreateBy:            x.CreateBy,
//		CreditCardAmount:    x.CreditCardAmount,
//		DepositAmount:       x.DepositAmount,
//		DeliveryStatus:      x.DeliveryStatus,
//		DiscountWord:        x.DiscountWord,
//		DiscountAmount:      x.DiscountAmount,
//		DocDate:             x.DocDate,
//		DocType:             x.DocType,
//		GLStatus:            x.GLStatus,
//		HoldingStatus:       x.HoldingStatus,
//		IsReturn:            x.IsReturn,
//		IsPosted:            x.IsPosted,
//		InvoiceNo:           x.InvoiceNo,
//		NumberOfItem:        x.NumberOfItem,
//		NetAmount:           x.NetAmount,
//		OtpPassword:         x.OtpPassword,
//		OnlineAmount:        x.OnlineAmount,
//		PosMachineId:        x.PosMachineId,
//		PeriodId:            x.PeriodId,
//		PosStatus:           x.PosStatus,
//		PickStatus:          x.PickStatus,
//		QueId:               x.QueId,
//		ReceiveTel:          x.ReceiveTel,
//		ReceiveName:         x.ReceiveName,
//		SaleId:              x.SaleId,
//		SumItemAmount:       x.SumItemAmount,
//		ScgId:               x.ScgId,
//		Status:              x.Status,
//		TaxAmount:           x.TaxAmount,
//		TotalAmount:         x.TotalAmount,
//		TaxNo:               x.TaxNo,
//		TaxType:             x.TaxType,
//		TaxRate:             x.TaxRate,
//		Uuid:                x.Uuid,
//		Sub:                 subs,
//	}
//}
//
//func map_basketsub_request(x BasketSub) BasketSubTemplate {
//	return BasketSubTemplate{
//		AverageCost:     x.AverageCost,
//		CheckoutTime:    x.CheckoutTime,
//		BarCode:         x.BarCode,
//		CheckoutQty:     x.CheckoutQty,
//		CheckoutBy:      x.CheckoutBy,
//		CheckoutAmount:  x.CheckoutAmount,
//		DocDate:         x.DocDate,
//		DeliveryOrderId: x.DeliveryOrderId,
//		ItemId:          x.ItemId,
//		ItemName:        x.ItemName,
//		IsReturn:        x.IsReturn,
//		Id:              x.Id,
//		ItemCode:        x.ItemCode,
//		LineNumber:      x.LineNumber,
//		PickTime:        x.PickTime,
//		PickQty:         x.PickQty,
//		Price:           x.Price,
//		PickAmount:      x.PickAmount,
//		PosId:           x.PosId,
//		PickBy:          x.PickBy,
//		QueId:           x.QueId,
//		Qty:             x.Qty,
//		RequestTime:     x.RequestTime,
//		RequestQty:      x.RequestQty,
//		RequestBy:       x.RequestBy,
//		RefLineNumber:   x.RefLineNumber,
//		RefNo:           x.RefNo,
//		RemainQty:       x.RemainQty,
//		Rate1:           x.Rate1,
//		ShelfId:         x.ShelfId,
//		SumOfCost:       x.SumOfCost,
//		SaleId:          x.SaleId,
//		Uuid:            x.Uuid,
//		WhId:            x.WhId,
//	}
//}
