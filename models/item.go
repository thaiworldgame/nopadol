package models

import (
	"fmt"
	"strings"
)

type Item struct {
	Parent *Category
	Name   string
	UOM    *Unit
}

type Unit struct {
	Name  string
	Ratio float32
}

func (i *Item) MoveCat(newCat *Category) {
	i.Parent = newCat
}

func (i *Item) Stock(s []*Stock) []*Stock {
	r := make([]*Stock, 0)
	for _, v := range s {
		if v.item == i {
			r = append(r, v)
		}
	}
	//	fmt.Printf("r=", r)
	return r
}

func (i *Item) Calc(st []*Stock, tx []*Trans) {
	for _, t := range tx {
		if t.item == i {
			for k, s := range st {
				// ถ้า Trans ไม่มี Location ใน Stock ให้เพิ่มรายการใน Stock ก่อน
				if s.item == i && s.loc == t.locIn {
					s.bal += t.qty
					fmt.Printf("Item %v Loc: %v qty: %v bal: %v\n", s.item.Name, s.loc.Code, t.qty, s.bal)
				} else {
					st = append(st, &Stock{item: i, loc: t.locIn, bal: t.qty})
					fmt.Println(">>Append Stock from Loc In:", st[k].item.Name, st[k].loc.Code, st[k].bal)
				}
				if s.item == i && s.loc == t.locOut {
					s.bal -= t.qty
					fmt.Printf("Item %v Loc: %v qty:-%v bal: %v\n", s.item.Name, s.loc.Code, t.qty, s.bal)
				} else {
					st = append(st, &Stock{item: i, loc: t.locOut, bal: t.qty})
					fmt.Println(">>Append Stock from Loc Out:", st[k].item.Name, st[k].loc.Code, st[k].bal)
				}
			}
		}
	}
}

// ==== Category ==== //
type Category struct {
	Parent *Category
	Name   string
}

func (c *Category) New() {
	c.Name = "New Cat"
}

func FindByName(cats []*Category, n string) *Category {
	for i, c := range cats {
		if strings.Contains(c.Name, n) {
			return cats[i]
		}
	}
	return nil
}

func (c *Category) MoveCat(newCat *Category) {
	// ต้องเพิ่มการป้องกันการอ้าง Child มาเป็น Parent ของโหนด
	// โดยจะต้องทวนสอบย้อนกลับว่าโหนดแม่ที่ย้ายมาจะไม่อยู่ใต้โหนดของลูกตัวเอง
	c.Parent = newCat
}

// ==== Location ==== //
type LocType int

const (
	ROOT LocType = iota
	BUY
	STORE
	VEHICLE
	MACHINE
	COLUMN
	SALE
)

type Location struct {
	Parent *Location
	LocType
	Code string
}

// ==== Stock ==== //
type Stock struct {
	item  *Item
	loc   *Location
	bal   int64
	digit int8 // ขนาดของทศนิยมที่จะใช้กับสินค้านี้
}

func (s *Stock) Move(newLoc *Location) {
	s.loc = newLoc
}

// ==== Transaction ==== //
type Trans struct {
	ID     uint64
	item   *Item
	locOut *Location
	locIn  *Location
	qty    int64
}
