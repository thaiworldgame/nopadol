package pointofsale

type BasketTemplate struct {
	Id             int64               `json:"id"`
	CompanyId      int64               `json:"company_id"`
	BranchId       int64               `json:"branch_id"`
	UUID           string              `json:"uuid"`
	InvoiceNo      string              `json:"invoice_no"`
	QueId          int64               `json:"que_id"`
	DocType        int64               `json:"doc_type"`
	DocDate        string              `json:"doc_date"`
	ArId           int64               `json:"ar_id"`
	SaleId         int64               `json:"sale_id"`
	TaxType        int64               `json:"tax_type"`
	TaxRate        int64               `json:"tax_rate"`
	NumberOfItem   int64               `json:"number_of_item"`
	SumItemAmount  float64             `json:"sum_item_amount"`
	TotalAmount    float64             `json:"total_amount"`
	NetAmount      float64             `json:"net_amount"`
	OtpPassword    string              `json:"otp_password"`
	Status         int64               `json:"status"`
	PickStatus     int64               `json:"pick_status"`
	DeliveryStatus int64               `json:"delivery_status"`
	ReceiveName    string              `json:"receive_name"`
	ReceiveTel     string              `json:"receive_tel"`
	IsCancel       int64               `json:"is_cancel"`
	IsConfirm      int64               `json:"is_confirm"`
	CreateBy       string              `json:"create_by"`
	CreateTime     string              `json:"create_time"`
	EditBy         string              `json:"edit_by"`
	EditTime       string              `json:"edit_time"`
	ConfirmBy      string              `json:"confirm_by"`
	ConfirmTime    string              `json:"confirm_time"`
	CancelBy       string              `json:"cancel_by"`
	CancelTime     string              `json:"cancel_time"`
	CancelDescId   int64               `json:"cancel_desc_id"`
	CancelDesc     string              `json:"cancel_desc"`
	Sub            []BasketSubTemplate `json:"sub"`
}

type BasketSubTemplate struct {
	Id              int64   `json:"id"`
	BasketId        int64   `json:"basket_id"`
	Uuid            string  `json:"uuid"`
	QueId           int64   `json:"que_id"`
	DocDate         string  `json:"doc_date"`
	ItemId          int64   `json:"item_id"`
	ItemCode        string  `json:"item_code"`
	ItemName        string  `json:"item_name"`
	BarCode         string  `json:"bar_code"`
	RequestQty      float64 `json:"request_qty"`
	PickQty         float64 `json:"pick_qty"`
	CheckoutQty     float64 `json:"checkout_qty"`
	Price           float64 `json:"price"`
	UnitId          int64   `json:"unit_id"`
	PickAmount      float64 `json:"pick_amount"`
	CheckoutAmount  float64 `json:"checkout_amount"`
	Qty             float64 `json:"qty"`
	RemainQty       float64 `json:"remain_qty"`
	Rate1           int64   `json:"rate_1"`
	RefNo           string  `json:"ref_no"`
	SaleId          int64   `json:"sale_id"`
	AverageCost     float64 `json:"average_cost"`
	DeliveryOrderId int64   `json:"delivery_order_id"`
	LineNumber      int64   `json:"line_number"`
	RequestBy       string  `json:"request_by"`
	RequestTime     string  `json:"request_time"`
	PickBy          string  `json:"pick_by"`
	PickTime        string  `json:"pick_time"`
	CheckoutBy      string  `json:"checkout_by"`
	CheckoutTime    string  `json:"checkout_time"`
}
