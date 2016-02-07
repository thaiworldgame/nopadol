package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Person struct {
	FName string
	LName string
	NName string
	BDate time.Time
}

type Employee struct {
	*Person
	*Title
	Code   string
	salary decimal.Decimal
}

type Title struct {
	Parent *Title
	TH     string
	EN     string
}

type Org struct {
	Parent *Org
	TH     string
	EN     string
	Short  string
}
