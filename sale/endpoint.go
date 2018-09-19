package sale

import (
	"context"
	"fmt"
)

type (

	NewQTRequest struct {
		Id                  int64              `json:"id"`
		DocNo               string             `json:"doc_no"`
		DocDate             string             `json:"doc_date"`
		ArId                int64              `json:"ar_id"`
		ArCode              string             `json:"ar_code"`
		ArName              string             `json:"ar_name"`
		SaleId              int                `json:"sale_id"`
		SaleCode            string             `json:"sale_code"`
		SaleName            string             `json:"sale_name"`
		BillType            int64              `json:"bill_type"`
		TaxType             int64              `json:"tax_type"`
		TaxRate             int64              `json:"tax_rate"`
		DepartCode          string             `json:"depart_code"`
		RefNo               string             `json:"ref_no"`
		IsConfirm           int64              `json:"is_confirm"`
		BillStatus          int64              `json:"bill_status"`
		DueDate             string             `json:"due_date"`
		ExpireDate          string             `json:"expire_date"`
		DeliveryDate        string             `json:"delivery_date"`
		AssertStatus        int64              `json:"assert_status"`
		IsConditionSend     int64              `json:"is_condition_send"`
		MyDescription       string             `json:"my_description"`
		SumItemAmount       float64            `json:"sum_item_amount"`
		DiscountWord        string             `json:"discount_word"`
		DiscountAmount      float64            `json:"discount_amount"`
		AfterDiscountAmount float64            `json:"after_discount"`
		BeforeTaxAmount     float64            `json:"before_tax_amount"`
		TaxAmount           float64            `json:"tax_amount"`
		TotalAmount         float64            `json:"total_amount"`
		NetAmount           float64            `json:"net_debt_amount"`
		ProjectId           int64              `json:"project_id"`
		ProjectCode         string             `json:"project_code"`
		IsCancel            int64              `json:"is_cancel"`
		CreateBy            string             `json:"creator_by"`
		CreateTime          string             `json:"create_time"`
		EditBy              string             `json:"edit_by"`
		EditTime            string             `json:"edit_time"`
		CancelBy            string             `json:"cancel_by"`
		CancelTime          string             `json:"cancel_time"`
		Subs                []NewQTItemRequest `json:"subs"`
	}

	NewQTItemRequest struct {
		Id              int64   `json:"id"`
		QTId            int64   `json:"qt_id"`
		ArId            int64   `json:"ar_id"`
		SaleId          int64   `json:"sale_id"`
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
		LineNumber      int     `json:"line_number"`
	}
)

func Create(s Service) interface{} {
	return func(ctx *context.Context, req *NewQTRequest) (interface{}, error) {
		q := map_qt_request(req)

		for _, sub := range req.Subs {
			itemline := map_qt_sub_request(sub)
			q.Subs = append(q.Subs, itemline)
		}

		resp, err := s.Create(&NewQTTemplate{
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
			SumItemAmount:       req.SumItemAmount,
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

func map_qt_request(x *NewQTRequest) NewQTTemplate {
	var subs []NewQTItemTemplate

	return NewQTTemplate{
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
		SumItemAmount:       x.SumItemAmount,
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

func map_qt_sub_request(x NewQTItemRequest) NewQTItemTemplate {
	return NewQTItemTemplate{
		BarCode:         x.BarCode,
		ItemCode:        x.ItemCode,
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
	}
}
