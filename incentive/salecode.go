package incentive

type SaleCode struct {
	SaleCode  string `db:"SaleCode"`
	SaleName string `db:"SaleName"`
	Subs []*struct {
		EnYear int `db:"EnYear"`
		MonthOfYear int `db:"MonthOfYear"`
		ProfitCenter string `db:"ProfitCenter"`
		TeamStatus string `db:"TeamStatus"`
	}
}