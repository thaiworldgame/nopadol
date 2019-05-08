package customer

type SearchByKeywordTemplate struct {
	Keyword string `json:"keyword"`
}

type SearchByIdTemplate struct {
	Id int64 `json:"id"`
}

type CustomerTemplate struct {
	Id           int64   `json:"id"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	Telephone    string  `json:"telephone"`
	BillCredit   int64   `json:"bill_credit"`
	Email        string  `json:"email"`
	CreditAmount float64 `json:"credit_amount"`
	CompanyID    int     `json:"company_id"`
	CreateBy     string  `json:"create_by"`
	Fax          string  `json:"fax"`
	TaxNo        string  `json:"tax_no"`
	DebtAmount   float64 `json:"debt_amount"`
	DebtLimit    float64 `json:"debt_limit"`
	MemberID     string  `json:"member_id"`
	PointBalance float64 `json:"point_balance"`
}
