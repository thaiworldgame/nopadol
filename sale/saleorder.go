package sale

type SaleOrder struct {
	DocNo   string `json:"doc_no" db:"DocNo"`
	DocDate string `json:"doc_date" db:"DocDate"`
	ArCode  string `json:"ar_code" db:"ArCode"`
	ArName  string `json:"ar_name" db:"ArName"`
	Subs []*struct {
		ItemCode string  `json:"item_code" db:"ItemCode"`
		ItemName string  `json:"item_name" db:"ItemName"`
		Qty      float64 `json:"qty" db:"Qty"`
		UnitCode string  `json:"unit_code" db:"UnitCode"`
	}
}

