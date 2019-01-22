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

//type CusType int
//
//const (
//	CONTRACTOR CusType = iota
//	RETAILER
//	PERSON
//	COMPANY
//	GOVERNMENT
//	ACADEMY
//)
//
//type Customer struct {
//	CusType
//	Contact *Person
//	Name    string
//	debit   decimal.Decimal
//	credit  decimal.Decimal
//}
//
//type Rank int
//
//const (
//	A Rank = iota
//	B
//	C
//	D
//)
//
//// CreditScore of customer
//type CreditScore struct {
//	*Period
//	*Customer
//	Rank
//	KI1Continous      int
//	KI2PaymentDue     int
//	KI3Responsibility int
//	KI4Charactor      int
//}
