package environment

type DepartmentTemplate struct {
	Id           int64  `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	ActiveStatus int64  `json:"active_status"`
	CreateBy     string `json:"create_by"`
	CreateTime   string `json:"create_time"`
	EditBy       string `json:"edit_by"`
	EditTime     string `json:"edit_time"`
}

type SearchByIdTemplate struct {
	Id int64 `json:"id"`
}

type SearchByKeywordTemplate struct {
	Keyword string `json:"keyword"`
}

type FindCustContactModel struct {
	Id          int64  `json:"id" db:"id"`
	ArId        int64  `json:"ar_id" db:"ar_id"`
	ContactCode string `json:"contact_code" db:"contact_code""`
	ContactName string `json:"contact_name" db:"contact_name"`
	Address     string `json:"address" db:"address"`
	Telephone   string `json:"telephone" db:"telephone"`
	Email       string `json:"email" db:"email"`
	LineId      string `json:"line_id" db:"line_id"`
	LineNumber  int64  `json:"line_number" db:"line_number"`
}
