package models

import (
	"github.com/shopspring/decimal"
)

type CusType int

const (
	CONTRACTOR CusType = iota
	RETAILER
	PERSON
	COMPANY
	GOVERNMENT
	ACADEMY
)

type Customer struct {
	CusType
	Contact *Person
	Name    string
	debit   decimal.Decimal
	credit  decimal.Decimal
}

type Rank int

const (
	A Rank = iota
	B
	C
	D
)

// CreditScore of customer
type CreditScore struct {
	*Period
	*Customer
	Rank
	KI1Continous      int
	KI2PaymentDue     int
	KI3Responsibility int
	KI4Charactor      int
}
