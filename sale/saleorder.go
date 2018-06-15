package sale

type SaleOrderTemplate struct {
	DocNo   string `json:"doc_no"`
	DocDate string `json:"doc_date"`
	ArCode  string `json:"ar_code"`
	ArName  string `json:"ar_name"`
	Subs    []SubsTemplate `json:"subs"`
}
type SubsTemplate struct {
	ItemCode string `json:"item_code"`
	ItemName string `json:"item_name"`
	Qty      float64 `json:"qty"`
	UnitCode string `json:"unit_code"`
}
