package customer

type SearchByKeywordTemplate struct {
	Keyword string `json:"keyword"`
}

type SearchByIdTemplate struct {
	Id int64 `json:"id"`
}

type CustomerTemplate struct {
	CustomerId        int64  `json:"customer_id"`
	CustomerCode      string `json:"customer_code"`
	CustomerName      string `json:"customer_name"`
	CustomerAddress   string `json:"customer_address"`
	CustomerTelephone string `json:"customer_telephone"`
}
