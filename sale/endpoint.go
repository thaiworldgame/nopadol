package sale

import "context"

// Endpoint is the sale endpoint
type Endpoint interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Search(context.Context, *SearchSaleOrderRequest) (*SearchSaleOrderResponse, error)
	NewSaleOrder(context.Context, NewSaleOrderRequest) (*NewSaleOrderResponse, error)
	Search(context.Context, *SearchSaleOrderRequest) (*SearchSaleOrderResponse, error)

	//POS /////////////////////////////////////////////////////////////////////////////////////////////
	//NewPos(context.Context, NewPosRequest) (*NewResponse, error)
}

type (
	// CreateRequest is the request for create endpoint
	CreateRequest struct {
		Field1 string `json:"field1"`
	}

	// CreateResponse is the response for create endpoint
	CreateResponse struct {
		ID string `json:"id"`
	}

	//Sale Order ///////////////////////////////////////////////////////////////////////////////////////////////////////

	// CreateRequest SaleOrder Struct
	NewSaleOrderRequest struct {
		DocNo   string                    `json:"doc_no"`
		DocDate string                    `json:"doc_date"`
		ArCode  string                    `json:"ar_code"`
		ArName  string                    `json:"ar_name"`
		Subs    []NewSubsSaleOrderRequest `json:"subs"`
	}

	// CreateRequest SaleOrder Details Struct
	NewSubsSaleOrderRequest struct {
		ItemCode string  `json:"item_code"`
		ItemName string  `json:"item_name"`
		Qty      float64 `json:"qty"`
		UnitCode string  `json:"unit_code"`
	}

	NewSaleOrderResponse struct {
		Id int64 `json:"doc_no"`
	}

	SearchSaleOrderRequest struct {
		Keyword string `json:"keyword"`
	}

	SearchSaleOrderResponse struct {
		DocNo   string                  `json:"doc_no"`
		DocDate string                  `json:"doc_date"`
		ArCode  string                  `json:"ar_code"`
		ArName  string                  `json:"ar_name"`
		Subs    []SubsSaleOrderResponse `json:"subs"`
	}

	SubsSaleOrderResponse struct {
		ItemCode string  `json:"item_code"`
		ItemName string  `json:"item_name"`
		Qty      float64 `json:"qty"`
		UnitCode string  `json:"unit_code"`
	}

	//Center ///////////////////////////////////////////////////////////////////////////////////////////////////////////
	NewResponse struct {
		Id int64 `json:"id"`
	}

	//POS //////////////////////////////////////////////////////////////////////////////////////////////////////////////

	OutPutTax struct {
		TaxNo    string `json:"tax_no"`
		TaxDate  string `json:"tax_date"`
		BookCode string `json:"book_code"`
	}

	Customer struct {
		ArCode string `json:"ar_code"`
	}

	SaleMan struct {
		SaleCode string `json:"sale_code"`
	}

	NewPosRequest struct {
		SaveFrom        int              `json:"save_from"`
		Source          int              `json:"source"`
		DocNo           string           `json:"doc_no"`
		DocDate         string           `json:"doc_date"`
		OutPutTax
		Customer
		SaleMan
		ShiftCode       string           `json:"shiftcode"`
		CashierCode     string           `json:"cashier_code"`
		ShiftNo         string           `json:"shift_no"`
		MachineNo       string           `json:"machine_no"`
		MachineCode     string           `json:"machine_code"`
		CoupongAmount   float64          `json:"coupong_amount"`
		ChangeAmount    float64          `json:"change_amount"`
		ChargeAmount    float64          `json:"charge_amount"`
		TaxType         int              `json:"tax_type"`
		MyDescription   string           `json:"my_description"`
		SumOfItemAmount float64          `json:"sum_of_item_amount"`
		DiscountWord    string           `json:"discount_word"`
		AfterDiscount   float64          `json:"after_discount"`
		TotalAmount     float64          `json:"total_amount"`
		SumCashAmount   float64          `json:"sum_cash_amount"`
		SumChqAmount    float64          `json:"sum_chq_amount"`
		SumCreditAmount float64          `json:"sum_credit_amount"`
		SumBankAmount   float64          `json:"sum_bank_amount"`
		NetDebtAmount   float64          `json:"net_debt_amount"`
		UserCode        string           `json:"user_code"`
		PosSubs         []NewPosItemRequest `json:"pos_subs"`
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
)
