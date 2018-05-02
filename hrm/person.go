package models

import (
	"github.com/shopspring/decimal"
	"time"
	//	"github.com/jinzhu/gorm"
)

type Person struct {
	//	gorm.Model
	FirstName string    `gorm:"column:FirstName"`
	LastName  string    `gorm:"column:LastName"`
	NickName  string    `gorm:"column:NickName"`
	BirthDate time.Time `gorm:"column:BirthDate"`
}

type Employee struct {
	*Person
	PersonID int
	Titles   []*Title
//	Orgs     []Org
	Code     string
	salary   decimal.Decimal
}

type Title struct {
	Parent *Title
	ID     int    `gorm:"primary_key"`
	TH     string `gorm:"column:TH"`
	EN     string `gorm:"column:EN"`
}

type Org struct {
	ID     int    `gorm:"column:ID;primary_key"`
	Parent *Org   `gorm:"column:Parent"`
	TH     string `gorm:"column:TH"`
	EN     string `gorm:"column:EN"`
	Short  string `gorm:"column:Short"`
}
