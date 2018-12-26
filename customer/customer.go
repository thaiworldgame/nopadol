package customer

type SearchByKeywordTemplate struct {
	Keyword string `json:"keyword"`
}

type SearchByIdTemplate struct {
	Id int64 `json:"id"`
}

type CustomerTemplate struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
	BillCredit int64  `json:"bill_credit"`
}
