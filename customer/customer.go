package customer

import (

)


type SearchByKeyword struct {
	keyword string `json:"keyword"`
}

type SearchById struct {
	Id int64 `json:"id"`
}

type CustomerTemplate struct {
	CustomerId int64 `json:"customer_id"`
	CustomerCode string `json:"customer_code"`
	CustomerName string `json:"customer_name"`
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
