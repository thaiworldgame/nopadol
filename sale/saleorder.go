package sale

type SaleOrder struct {
	DocNo  string `db:"DocNo"`
	ArCode string `db:"ArCode"`
	ArName string `db:"ArName"`
	Subs []*struct {
		ItemCode string `db:"ItemCode"`
		ItemName string  `db:"ItemName"`
		Qty float64 `db:"Qty"`
		UnitCode string `db:"UnitCode"`
	}
}

