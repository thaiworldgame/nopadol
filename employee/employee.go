package employee


type SearchByKeyword struct {
	keyword string `json:"keyword"`
}

type SearchById struct {
	Id int64 `json:"id"`
}

type EmployeeTemplate struct {
	Id int64 `json:"employee_id"`
	SaleCode string `json:"sale_code"`
	SaleName string `json:"sale_name"`
}
