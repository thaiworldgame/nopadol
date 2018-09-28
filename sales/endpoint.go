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
		ArId                int64               `json:"ar_id"`
		ArCode              string              `json:"ar_code"`
		ArName              string              `json:"ar_name"`
		SaleId              int                 `json:"sale_id"`
		SaleCode            string              `json:"sale_code"`
		SaleName            string              `json:"sale_name"`
		BillType            int64               `json:"bill_type"`
		TaxType             int                 `json:"tax_type"`
		TaxRate             float64             `json:"tax_rate"`
		DepartCode          string              `json:"depart_code"`
		RefNo               string              `json:"ref_no"`
		IsConfirm           int64               `json:"is_confirm"`
		BillStatus          int64               `json:"bill_status"`
		DueDate             string              `json:"due_date"`
		ExpireDate          string              `json:"expire_date"`
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
		ArId                int64                `json:"ar_id"`
		ArCode              string               `json:"ar_code"`
		ArName              string               `json:"ar_name"`
		SaleId              int                  `json:"sale_id"`
		SaleCode            string               `json:"sale_code"`
		SaleName            string               `json:"sale_name"`
		BillType            int64                `json:"bill_type"`
		TaxType             int                  `json:"tax_type"`
		TaxRate             float64              `json:"tax_rate"`
		DepartCode          string               `json:"depart_code"`
		RefNo               string               `json:"ref_no"`
		IsConfirm           int64                `json:"is_confirm"`
		BillStatus          int64                `json:"bill_status"`
		SoStatus            int64                `json:"so_status"`
		HoldingStatus       int64                `json:"holding_status"`
		CreditDay           int64                `json:"credit_day"`
		DueDate             string               `json:"due_date"`
		ExpireDate          string               `json:"expire_date"`
		DeliveryDate        string               `json:"delivery_date"`
		AssertStatus        int64                `json:"assert_status"`
		IsConditionSend     int64                `json:"is_condition_send"`
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
		Keyword string `json:"keyword"`
	}
)

////// Quotation /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateQuo(s Service) interface{} {
	return func(ctx context.Context, req *NewQuoRequest) (interface{}, error) {
		q := map_quo_request(req)

		fmt.Println("p =")

		for _, subs := range req.Subs {
			itemline := map_quo_sub_request(subs)
			q.Subs = append(q.Subs, itemline)
		}
		resp, err := s.CreateQuo(&NewQuoTemplate{
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
			DepartCode:          req.DepartCode,
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
		DepartCode:          x.DepartCode,
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

func CreateSale(s Service) interface{} {
	return func(ctx context.Context, req *NewSaleRequest) (interface{}, error) {
		so := map_sale_request(req)

		for _, subs := range req.Subs {
			itemline := map_sale_sub_request(subs)
			so.Subs = append(so.Subs, itemline)
		}
		resp, err := s.CreateSale(&NewSaleTemplate{
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
			DepartCode:          req.DepartCode,
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
		SaleId:              x.SaleId,
		SaleCode:            x.SaleCode,
		SaleName:            x.SaleName,
		TaxType:             x.TaxType,
		TaxRate:             x.TaxRate,
		RefNo:               x.RefNo,
		DepartCode:          x.DepartCode,
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

func SearchSaleById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchSaleById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
