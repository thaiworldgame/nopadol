package pos

import (
	"context"
)

type Endpoint interface {
	NewPos(context.Context, NewPosRequest) (*NewPosResponse, error)
}

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
)
