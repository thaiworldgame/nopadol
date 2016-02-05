package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Person struct {
	Fname string
	Lname string
	Nname string
	Bdate time.Time
}

type Employee struct {
	*Person
	*Title
	Code   string
	salary decimal.Decimal
}

type Title struct {
	TH string
	EN string
}
