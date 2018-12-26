package gendocno

type DocNoTemplate struct {
	BranchId     int64  `json:"branch_id" db:"BranchId"`
	TableCode    string `json:"table_code" db:"TableCode"`
	BillType     int64  `json:"bill_type" db:"BillType"`
	Header       string `json:"header" db:"Header"`
	UseYear      int64  `json:"use_year" db:"UseYear"`
	UseMonth     int64  `json:"use_month" db:"UseMonth"`
	UseDay       int64  `json:"use_day" db:"UseDay"`
	UseDash      int64  `json:"use_dash" db:"UseDash"`
	FormatNumber int64  `json:"format_number" db:"FormatNumber"`
	ActiveStatus int64  `json:"active_status" db:"ActiveStatus"`
}
