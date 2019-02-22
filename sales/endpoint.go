package sales

import (
	"context"
	"fmt"
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
		SaleId              int64               `json:"sale_id"`
		SaleCode            string              `json:"sale_code"`
		SaleName            string              `json:"sale_name"`
		BillType            int64               `json:"bill_type"`
		TaxType             int64               `json:"tax_type"`
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
		ConfirmBy           string              `json:confirm_by`
		ConfirmTime         string              `json:"confirm_time"`
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
		SaleId              int64                `json:"sale_id"`
		SaleCode            string               `json:"sale_code"`
		SaleName            string               `json:"sale_name"`
		BillType            int64                `json:"bill_type"`
		TaxType             int64                `json:"tax_type"`
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
		ArId     int64  `json:"ar_id"`
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
		RecMoney         []RecMoney              `json:"rec_money"`
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
	Bankpay struct {
		Id           int64   `json:"id"`
		RefId        int64   `json:"ref_id"`
		BankAccont   string  `json:"bank_account"`
		BankName     string  `json:"bank_name"`
		BankAmount   float64 `json:"bank_amount"`
		Activestatus int64   `json:"active_status"`
		CreateBy     string  `json:"create_by"`
		EditBy       string  `json:"edit_by"`
	}

	RecMoney struct {
		Id             int64   `json:"id"`
		DocType        int64   `json:"doc_type"`
		RefId          int64   `json:"ref_id"`
		ArId           int64   `json:"ar_id"`
		PaymentType    int64   `json:"payment_type"`
		PayAmount      float64 `json:"pay_amount"`
		ChqTotalAmount float64 `json:"chq_total_amount"`
		CreditType     int64   `json:"credit_type"`
		ChargeAmount   float64 `json:"charge_amount"`
		ConfirmNo      string  `json:"confirm_no"`
		RefNo          string  `json:"ref_no"`
		BankCode       string  `json:"bank_code"`
		BankBranchCode string  `json:"bank_branch_code"`
		RefDate        string  `json:"ref_date"`
		BankTransDate  string  `json:"bank_trans_date"`
		LineNumber     int64   `json:"line_number"`
	}

	NewInvoiceRequest struct {
		Id                  int64                   `json:"id"`
		CompanyId           int64                   `json:"company_id"`
		BranchId            int64                   `json:"branch_id"`
		ItemId              int64                   `json:"item_id"`
		ItemCode            string                  `json:"item_code"`
		Uuid                string                  `json:"uuid"`
		DocNo               string                  `json:"doc_no"`
		TaxNo               string                  `json:"tax_no"`
		BillType            int64                   `json:"bill_type"`
		DocDate             string                  `json:"doc_date"`
		ArId                int64                   `json:"ar_id"`
		ArCode              string                  `json:"ar_code"`
		ArName              string                  `json:"ar_name"`
		ArBillAddress       string                  `json:"ArBillAddress"`
		ArTelephone         string                  `json:"ArTelephone"`
		SaleId              int64                   `json:"sale_id"`
		SaleCode            string                  `json:"sale_code"`
		SaleName            string                  `json:"sale_name"`
		PosMachineId        int64                   `json:"pos_machine_id"`
		PeriodId            int64                   `json:"period_id"`
		CashId              int64                   `json:"cash_id"`
		TaxType             int64                   `json:"tax_type"`
		TaxRate             float64                 `json:"tax_rate"`
		NumberOfItem        float64                 `json:"number_of_item"`
		DepartId            string                  `json:"depart_id"`
		AllocateId          int64                   `json:"allocate_id"`
		ProjectId           int64                   `json:"project_id"`
		PosStatus           int64                   `json:"pos_status"`
		CreditDay           int64                   `json:"credit_day"`
		DueDate             string                  `json:"due_date"`
		DeliveryDay         int64                   `json:"delivery_day"`
		DeliveryDate        string                  `json:"delivery_date"`
		IsConfirm           int64                   `json:"is_confirm"`
		IsConditionSend     int64                   `json:"is_condition_send"`
		MyDescription       string                  `json:"my_description"`
		SoRefNo             string                  `json:"so_ref_no"`
		ChangeAmount        float64                 `json:"change_amount"`
		SumCashAmount       float64                 `json:"sum_cash_amount"`
		SumCreditAmount     float64                 `json:"sum_credit_amount"`
		SumChqAmount        float64                 `json:"sum_chq_amount"`
		SumBankAmount       float64                 `json:"sum_bank_amount"`
		SumOfDeposit        float64                 `json:"sum_of_deposit"`
		SumOnLineAmount     float64                 `json:"sum_on_line_amount"`
		CouponAmount        float64                 `json:"coupon_amount"`
		SumOfItemAmount     float64                 `json:"sum_of_item_amount"`
		DiscountWord        string                  `json:"discount_word"`
		DiscountAmount      float64                 `json:"discount_amount"`
		AfterDiscountAmount float64                 `json:"after_discount_amount"`
		BeforeTaxAmount     float64                 `json:"before_tax_amount"`
		TaxAmount           float64                 `json:"tax_amount"`
		TotalAmount         float64                 `json:"total_amount"`
		NetDebtAmount       float64                 `json:"net_debt_amount"`
		BillBalance         float64                 `json:"bill_balance"`
		PayBillStatus       int64                   `json:"pay_bill_status"`
		PayBillAmount       float64                 `json:"pay_bill_amount"`
		DeliveryStatus      int64                   `json:"delivery_status"`
		ReceiveName         string                  `json:"receive_name"`
		ReceiveTel          string                  `json:"receive_tel"`
		CarLicense          string                  `json:"car_license"`
		IsCancel            int64                   `json:"is_cancel"`
		IsHold              int64                   `json:"is_hold"`
		IsPosted            int64                   `json:"is_posted"`
		IsCreditNote        int64                   `json:"is_credit_note"`
		IsDebitNote         int64                   `json:"is_debit_note"`
		GlStatus            int64                   `json:"gl_status"`
		JobId               string                  `json:"job_id"`
		JobNo               string                  `json:"job_no"`
		CouponNo            string                  `json:"coupon_no"`
		RedeemNo            string                  `json:"redeem_no"`
		ScgNumber           string                  `json:"scg_number"`
		ScgId               string                  `json:"scg_id"`
		CreateBy            string                  `json:"create_by"`
		CreateTime          string                  `json:"create_time"`
		EditBy              string                  `json:"edit_by"`
		EditTime            string                  `json:"edit_time"`
		ConfirmBy           string                  `json:"confirm_by"`
		ConfirmTime         string                  `json:"confirm_time"`
		CancelBy            string                  `json:"cancel_by"`
		CancelTime          string                  `json:"cancel_time"`
		CancelDescId        int64                   `json:"cancel_desc_id"`
		CancelDesc          string                  `json:"cancel_desc"`
		Subs                []NewInvoiceItemRequest `json:"subs"`
		//RecMoney            []RecMoney              `json:"rec_money"`
		CreditCard []CreditCard `json:"credit_card"`
		Chq        []ChqIn      `json:"chq"`
		BankPay    []Bankpay    `json:"bank"`
	}

	NewInvoiceItemRequest struct {
		Id    int64 `json:"id"`
		InvId int64 `json:"inv_id"`

		ItemCode        string  `json:"item_code"`
		Itemid          int64   `json:"item_id"`
		ItemName        string  `json:"item_name"`
		BarCode         string  `json:"bar_code"`
		WhId            int64   `json:"wh_id"`
		ShelfId         int64   `json:"shelf_id"`
		Price           float64 `json:"price"`
		UnitCode        string  `json:"unit_code"`
		Location        string  `json:"location"`
		Qty             float64 `json:"qty"`
		CnQty           float64 `json:"cn_qty"`
		DiscountWord    float64 `json:"discount_word_sub"`
		DiscountAmount  float64 `json:"discount_amount_sub"`
		ItemAmount      float64 `json:"amount"`
		NetAmount       float64 `json:"net_amount"`
		Average_cost    float64 `json:"average_cost"`
		SumOfCost       float32 `json:"sum_of_cost"`
		ItemDescription string  `json:"item_description"`
		IsCancel        int64   `json:"is_cancel"`
		IsCreditNote    int64   `json:"is_credit_note"`
		IsDebitNote     int64   `json:"is_debit_note"`
		PackingRate1    int64   `json:"packing_rate_1"`
		PackingRate2    int64   `json:"packing_rate_2"`
		RefNo           string  `json:"ref_no"`
		RefLineNumber   int64   `json:"ref_line_number"`
		LineNumber      int64   `json:"line_number"`
	}
)

// search item by keywork

////// Quotation /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreateQuotation(s Service) interface{} {
	return func(ctx context.Context, req *NewQuoRequest) (interface{}, error) {
		q := map_quo_request(req)

		fmt.Println("p =")

		for _, subs := range req.Subs {
			fmt.Println(subs, "asdajdlskal;kas;dk;aks")
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

func ConfirmQuotation(s Service) interface{} {
	return func(ctx context.Context, req *NewQuoRequest) (interface{}, error) {
		resp, err := s.ConfirmQuotation(&NewQuoTemplate{Id: req.Id, AssertStatus: req.AssertStatus, IsConfirm: req.IsConfirm, IsCancel: req.IsCancel, ConfirmBy: req.ConfirmBy, ConfirmTime: req.ConfirmTime})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func CancelQuotation(s Service) interface{} {
	return func(ctx context.Context, req *NewQuoRequest) (interface{}, error) {
		resp, err := s.CancelQuotation(&NewQuoTemplate{Id: req.Id, AssertStatus: req.AssertStatus, IsConfirm: req.IsConfirm, IsCancel: req.IsCancel, CancelBy: req.CancelBy, CancelTime: req.CancelTime})
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

func map_invoice_sub_request(x NewInvoiceItemRequest) NewInvoiceItemTemplate {
	fmt.Println("endpoint x", x)
	return NewInvoiceItemTemplate{

		ItemCode:        x.ItemCode,
		Itemid:          x.Itemid,
		BarCode:         x.BarCode,
		ItemName:        x.ItemName,
		WhId:            x.WhId,
		ShelfId:         x.ShelfId,
		Qty:             x.Qty,
		Location:        x.Location,
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
	fmt.Println("Invoicelist 222")
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {

		resp, err := s.SearchDocByKeyword(&SearchByKeywordTemplate{SaleCode: req.SaleCode, Keyword: req.Keyword})
		fmt.Println(resp, "99999999999999999999999999999999999999999999999999999999999999999999999999999999")
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
	//var subs []NewDepositItemRequest
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
		//Subs:             subs,
		CreditCard: credit_cards,
		Chq:        chqs,
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

func map_bank_request(x Bankpay) BankpayTemplate {
	return BankpayTemplate{
		Id:           x.Id,
		RefId:        x.RefId,
		BankAccount:  x.BankAccont,
		BankName:     x.BankName,
		BankAmount:   x.BankAmount,
		Activestatus: x.Activestatus,
		CreateBy:     x.CreateBy,
		EditBy:       x.EditBy,
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

func SearchReserveToDeposit(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchDepositByKeyword(&SearchByKeywordTemplate{ArId: req.ArId, Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

////// Invoice /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Invoicelist(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		fmt.Println("invoicelist 522")
		resp, err := s.Invoicelist(&SearchByKeywordTemplate{SaleCode: req.SaleCode, Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
func CancelInvoice(s Service) interface{} {
	return func(ctx context.Context, req *NewInvoiceRequest) (interface{}, error) {
		resp, err := s.CancelInvoice(&NewInvoiceTemplate{Id: req.Id, DocNo: req.DocNo, IsConfirm: req.IsConfirm, IsCancel: req.IsCancel, CancelBy: req.CancelBy, CancelTime: req.CancelTime})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func CreateInvoice(s Service) interface{} {
	fmt.Println("endpoint 1")
	return func(ctx context.Context, req *NewInvoiceRequest) (interface{}, error) {

		iv := map_invoice_request(req)

		for _, crds := range req.CreditCard {
			fmt.Println(crds, "caditcard2")
			crdline := map_creditcard_request(crds)
			iv.CreditCard = append(iv.CreditCard, crdline)
		}

		for _, chqs := range req.Chq {
			fmt.Println(chqs, "caditcard")
			chqline := map_chq_request(chqs)
			iv.Chq = append(iv.Chq, chqline)
		}
		for _, bnk := range req.BankPay {
			fmt.Println(bnk, "bank")
			bnkline := map_bank_request(bnk)
			iv.BankPay = append(iv.BankPay, bnkline)
		}
		for _, subs := range req.Subs {

			itemline := map_invoice_sub_request(subs)
			iv.Subs = append(iv.Subs, itemline)
		}
		fmt.Println(iv.Subs, "12312423534958309285083")
		resp, err := s.CreateInvoice(&NewInvoiceTemplate{
			Id:                  req.Id,
			CompanyId:           req.CompanyId,
			BranchId:            req.BranchId,
			DocNo:               req.DocNo,
			TaxNo:               req.TaxNo,
			DocDate:             req.DocDate,
			BillType:            req.BillType,
			ArId:                req.ArId,
			ArCode:              req.ArCode,
			ArName:              req.ArName,
			ArBillAddress:       req.ArBillAddress,
			ArTelephone:         req.ArTelephone,
			SaleId:              req.SaleId,
			SaleCode:            req.SaleCode,
			SaleName:            req.SaleName,
			TaxType:             req.TaxType,
			TaxRate:             req.TaxRate,
			CreditDay:           req.CreditDay,
			DueDate:             req.DueDate,
			DepartId:            req.DepartId,
			AllocateId:          req.AllocateId,
			ProjectId:           req.ProjectId,
			MyDescription:       req.MyDescription,
			BeforeTaxAmount:     req.BeforeTaxAmount,
			TaxAmount:           req.TaxAmount,
			TotalAmount:         req.TotalAmount,
			BillBalance:         req.BillBalance,
			IsCancel:            req.IsCancel,
			IsConfirm:           req.IsConfirm,
			ScgId:               req.ScgId,
			JobNo:               req.JobNo,
			CreateBy:            req.CreateBy,
			CreateTime:          req.CreateTime,
			EditBy:              req.EditBy,
			Uuid:                req.Uuid,
			AfterDiscountAmount: req.AfterDiscountAmount,
			CarLicense:          req.CarLicense,
			CashId:              req.CashId,
			ChangeAmount:        req.ChangeAmount,
			CouponAmount:        req.CouponAmount,
			CancelDesc:          req.CancelDesc,
			CancelDescId:        req.CancelDescId,
			CouponNo:            req.CouponNo,
			ConfirmTime:         req.ConfirmTime,
			ConfirmBy:           req.ConfirmBy,
			CancelBy:            req.CancelBy,
			CancelTime:          req.CancelTime,
			DiscountWord:        req.DiscountWord,
			DeliveryDate:        req.DeliveryDate,
			DeliveryDay:         req.DeliveryDay,
			DeliveryStatus:      req.DeliveryStatus,
			DiscountAmount:      req.DiscountAmount,
			EditTime:            req.EditTime,
			GlStatus:            req.GlStatus,
			IsCreditNote:        req.IsCreditNote,
			IsDebitNote:         req.IsDebitNote,
			IsConditionSend:     req.IsConditionSend,
			IsPosted:            req.IsPosted,
			IsHold:              req.IsHold,
			JobId:               req.JobId,
			NetDebtAmount:       req.NetDebtAmount,
			NumberOfItem:        req.NumberOfItem,
			PosMachineId:        req.PosMachineId,
			PosStatus:           req.PosStatus,
			PeriodId:            req.PeriodId,
			PayBillAmount:       req.PayBillAmount,
			PayBillStatus:       req.PayBillStatus,
			ReceiveName:         req.ReceiveName,
			ReceiveTel:          req.ReceiveTel,
			RedeemNo:            req.RedeemNo,
			SumBankAmount:       req.SumBankAmount,
			SumChqAmount:        req.SumChqAmount,
			SumCashAmount:       req.SumCashAmount,
			SumCreditAmount:     req.SumCreditAmount,
			SoRefNo:             req.SoRefNo,
			ScgNumber:           req.ScgNumber,
			SumOfDeposit:        req.SumOfDeposit,
			SumOnLineAmount:     req.SumOnLineAmount,
			SumOfItemAmount:     req.SumOfItemAmount,
			Subs:                iv.Subs,
			//		RecMoney:            iv.RecMoney,
			CreditCard: iv.CreditCard,
			Chq:        iv.Chq,
			BankPay:    iv.BankPay,
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

func map_invoice_request(x *NewInvoiceRequest) NewInvoiceTemplate {
	fmt.Println("endpoint3")
	var subs []NewInvoiceItemTemplate
	var credit_cards []CreditCardTemplate
	var chqs []ChqInTemplate
	var banks []BankpayTemplate
	//	var rec_moneys []RecMoneyTemplate

	return NewInvoiceTemplate{
		Id:                  x.Id,
		DocNo:               x.DocNo,
		DocDate:             x.DocDate,
		ArId:                x.ArId,
		ArTelephone:         x.ArTelephone,
		ArBillAddress:       x.ArBillAddress,
		ArName:              x.ArName,
		ArCode:              x.ArCode,
		AllocateId:          x.AllocateId,
		BranchId:            x.BranchId,
		BillBalance:         x.BillBalance,
		BillType:            x.BillType,
		BeforeTaxAmount:     x.BeforeTaxAmount,
		CompanyId:           x.CompanyId,
		CreditDay:           x.CreditDay,
		CreateTime:          x.CreateTime,
		CreateBy:            x.CreateBy,
		CancelTime:          x.CancelTime,
		CancelBy:            x.CancelBy,
		ConfirmBy:           x.ConfirmBy,
		ConfirmTime:         x.ConfirmTime,
		DueDate:             x.DueDate,
		DepartId:            x.DepartId,
		EditBy:              x.EditBy,
		EditTime:            x.EditTime,
		IsConfirm:           x.IsConfirm,
		IsCancel:            x.IsCancel,
		JobNo:               x.JobNo,
		MyDescription:       x.MyDescription,
		ProjectId:           x.ProjectId,
		SaleName:            x.SaleName,
		SaleCode:            x.SaleCode,
		SaleId:              x.SaleId,
		ScgId:               x.ScgId,
		TaxNo:               x.TaxNo,
		TotalAmount:         x.TotalAmount,
		TaxAmount:           x.TaxAmount,
		TaxRate:             x.TaxRate,
		TaxType:             x.TaxType,
		Uuid:                x.Uuid,
		AfterDiscountAmount: x.AfterDiscountAmount,
		CouponNo:            x.CouponNo,
		CancelDescId:        x.CancelDescId,
		CancelDesc:          x.CancelDesc,
		CouponAmount:        x.CouponAmount,
		ChangeAmount:        x.ChangeAmount,
		CashId:              x.CashId,
		CarLicense:          x.CarLicense,
		DiscountAmount:      x.DiscountAmount,
		DeliveryStatus:      x.DeliveryStatus,
		DeliveryDay:         x.DeliveryDay,
		DeliveryDate:        x.DeliveryDate,
		DiscountWord:        x.DiscountWord,
		IsHold:              x.IsHold,
		IsPosted:            x.IsPosted,
		IsConditionSend:     x.IsConditionSend,
		IsDebitNote:         x.IsDebitNote,
		IsCreditNote:        x.IsCreditNote,
		JobId:               x.JobId,
		NumberOfItem:        x.NumberOfItem,
		NetDebtAmount:       x.NetDebtAmount,
		PayBillStatus:       x.PayBillStatus,
		PayBillAmount:       x.PayBillAmount,
		PeriodId:            x.PeriodId,
		PosStatus:           x.PosStatus,
		PosMachineId:        x.PosMachineId,
		RedeemNo:            x.RedeemNo,
		ReceiveTel:          x.ReceiveTel,
		ReceiveName:         x.ReceiveName,
		SumOfItemAmount:     x.SumOfItemAmount,
		SumOnLineAmount:     x.SumOnLineAmount,
		SumOfDeposit:        x.SumOfDeposit,
		ScgNumber:           x.ScgNumber,
		SumCreditAmount:     x.SumCreditAmount,
		SumCashAmount:       x.SumCashAmount,
		SoRefNo:             x.SoRefNo,
		SumBankAmount:       x.SumBankAmount,
		SumChqAmount:        x.SumChqAmount,
		Subs:                subs,
		//	RecMoney:            rec_moneys,
		CreditCard: credit_cards,
		Chq:        chqs,
		BankPay:    banks,
	}
}

func SearchInvoiceById(s Service) interface{} {
	return func(ctx context.Context, req *SearchByIdRequest) (interface{}, error) {
		resp, err := s.SearchInvoiceById(&SearchByIdTemplate{Id: req.Id})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchInvoiceByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchInvoiceByKeyword(&SearchByKeywordTemplate{SaleCode: req.SaleCode, Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchSaleByItem(s Service) interface{} {
	return func(ctx context.Context, req *SearchByItemTemplate) (interface{}, error) {
		resp, err := s.SearchSaleByItem(&SearchByItemTemplate{Name: req.Name, ItemCode: req.ItemCode, Page: req.Page})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}

func SearchHisByKeyword(s Service) interface{} {
	return func(ctx context.Context, req *SearchByKeywordRequest) (interface{}, error) {
		resp, err := s.SearchHisByKeyword(&SearchByKeywordTemplate{SaleCode: req.SaleCode, Keyword: req.Keyword})
		if err != nil {
			fmt.Println("endpoint error =", err.Error())
			return nil, fmt.Errorf(err.Error())
		}
		return map[string]interface{}{
			"data": resp,
		}, nil
	}
}
