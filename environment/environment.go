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
