package models

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

func setupPerson() ([]*Period, []*Person, []*Customer, []*CusRank, []*Employee) {
	pr := make([]*Period, 12)
	pr[0] = &Period{Year: 2015, Month: 1}
	pr[1] = &Period{Year: 2015, Month: 2}
	pr[2] = &Period{Year: 2015, Month: 3}
	pr[3] = &Period{Year: 2015, Month: 4}
	pr[4] = &Period{Year: 2015, Month: 5}
	pr[5] = &Period{Year: 2015, Month: 6}
	pr[6] = &Period{Year: 2015, Month: 7}
	pr[7] = &Period{Year: 2015, Month: 8}
	pr[8] = &Period{Year: 2015, Month: 9}
	pr[9] = &Period{Year: 2015, Month: 10}
	pr[10] = &Period{Year: 2015, Month: 11}
	pr[11] = &Period{Year: 2015, Month: 12}

	ps := make([]*Person, 5)
	ps[0] = &Person{Fname: "เกษม", Lname: "อานนทวิลาศ", Nname: "Tom", Bdate: time.Date(1974, time.October, 4, 0, 0, 0, 0, time.UTC)}
	ps[1] = &Person{Fname: "จิราภรณ์", Lname: "อานนทวิลาศ", Nname: "Jip", Bdate: time.Date(1976, time.August, 12, 0, 0, 0, 0, time.UTC)}
	ps[2] = &Person{Fname: "ทนัฐพร", Lname: "อานนทวิลาศ", Nname: "Tim", Bdate: time.Date(1976, time.August, 31, 0, 0, 0, 0, time.UTC)}
	ps[3] = &Person{Fname: "ธนันท์", Lname: "อานนทวิลาศ", Nname: "Tam", Bdate: time.Date(1974, time.June, 7, 0, 0, 0, 0, time.UTC)}
	ps[4] = &Person{Fname: "สาธิต", Lname: "โฉมวัฒนา", Nname: "น้อย", Bdate: time.Date(1974, time.June, 10, 0, 0, 0, 0, time.UTC)}

	c := make([]*Customer, 5)
	c[0] = &Customer{CusType: COMPANY, Contact: ps[0], Name: "Dummy Customer", debit: decimal.Zero, credit: decimal.Zero}
	c[1] = &Customer{CusType: COMPANY, Contact: nil, Name: "บจก.เชียงใหม่คอนสตรัคชั่น"}
	c[2] = &Customer{CusType: COMPANY, Contact: nil, Name: "บจก.ปาล์ม การ์เด้น เชียงใหม่"}
	c[3] = &Customer{CusType: COMPANY, Contact: nil, Name: "บริษัท ส.เต็งไตรรัตน์(น่าน) จำกัด"}

	cr := make([]*CusRank, 5)
	cr[0] = &CusRank{Period: pr[0], Customer: c[0], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[1] = &CusRank{Period: pr[0], Customer: c[1], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[2] = &CusRank{Period: pr[0], Customer: c[2], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[3] = &CusRank{Period: pr[0], Customer: c[3], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}
	cr[4] = &CusRank{Period: pr[0], Customer: c[4], Rank: A, KI1Continous: 4, KI2PaymentDue: 4, KI3Responsibility: 4, KI4Charactor: 4}

	tt := make([]*Title, 7)
	tt[0] = &Title{TH: "ประธานกรรมการ", EN: "Board of Director"}
	tt[1] = &Title{TH: "รองประธานกรรมการ", EN: "Director"}
	tt[2] = &Title{TH: "กรรมการผู้จัดการ", EN: "Managing Director"}
	tt[3] = &Title{TH: "ผู้อำนวยการขาย", EN: "Sales Director"}
	tt[4] = &Title{TH: "ผู้อำนวยการบริหารสินค้า", EN: "Merchandise Director"}
	tt[5] = &Title{TH: "พนักงานขาย", EN: "Sales"}
	fmt.Println("Title")

	emp := make([]*Employee, 5)
	emp[0] = &Employee{Person: ps[0], Title: tt[2], Code: "39001", salary: decimal.New(40000, 0)}
	emp[1] = &Employee{Person: ps[1], Title: tt[3], Code: "48001", salary: decimal.New(30000, 0)}
	emp[2] = &Employee{Person: ps[2], Title: tt[4], Code: "49001", salary: decimal.New(20000, 0)}
	emp[3] = &Employee{Person: ps[3], Title: tt[4], Code: "50001", salary: decimal.New(20000, 0)}
	emp[4] = &Employee{Person: ps[4], Title: tt[5], Code: "53001", salary: decimal.New(9000, 0)}
	fmt.Println("Employee")
	return pr, ps, c, cr, emp
}

func setupItem() ([]*Category, []*Item, []*Location, []*Stock, []*Trans) {
	c := make([]*Category, 5)
	c[0] = &Category{Name: "root"}
	c[1] = &Category{Parent: c[0], Name: "Beverage"}
	c[2] = &Category{Parent: c[0], Name: "Snack"}
	c[3] = &Category{Parent: c[0], Name: "Fruit"}
	c[4] = &Category{Parent: c[1], Name: "Soda"}

	i := make([]*Item, 10)
	i[0] = &Item{Parent: c[0], Name: "Dummy Item"}
	i[1] = &Item{Parent: c[4], Name: "Coke Can 325ml"}
	i[2] = &Item{Parent: c[4], Name: "Pepsi Can 325ml"}
	i[3] = &Item{Parent: c[1], Name: "Ichitan Can 325ml"}
	i[4] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[5] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[6] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[7] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[8] = &Item{Parent: c[1], Name: "Coke Can 325ml"}
	i[9] = &Item{Parent: c[1], Name: "Coke Can 325ml"}

	l := make([]*Location, 15)
	l[0] = &Location{LocType: ROOT, Code: "0"}
	l[1] = &Location{Parent: l[0], LocType: STORE, Code: "S1"}
	l[2] = &Location{Parent: l[1], LocType: BUY, Code: "B1"}
	l[3] = &Location{Parent: l[1], LocType: BUY, Code: "V2"}
	l[4] = &Location{Parent: l[1], LocType: SALE, Code: "SA1"}
	l[5] = &Location{Parent: l[1], LocType: SALE, Code: "SA2"}
	l[6] = &Location{Parent: l[1], LocType: VEHICLE, Code: "V1"}
	l[7] = &Location{Parent: l[1], LocType: VEHICLE, Code: "V2"}
	l[8] = &Location{Parent: l[1], LocType: MACHINE, Code: "C1"}
	l[9] = &Location{Parent: l[1], LocType: MACHINE, Code: "C2"}
	l[10] = &Location{Parent: l[8], LocType: COLUMN, Code: "01"} // Column in Machine
	l[11] = &Location{Parent: l[8], LocType: COLUMN, Code: "02"}
	l[12] = &Location{Parent: l[8], LocType: COLUMN, Code: "03"}
	l[13] = &Location{Parent: l[8], LocType: COLUMN, Code: "04"}
	l[14] = &Location{Parent: l[8], LocType: COLUMN, Code: "05"}

	s := make([]*Stock, 5)
	s[0] = &Stock{item: i[1], loc: l[1], bal: 0, digit: 0}
	s[1] = &Stock{item: i[1], loc: l[2], bal: 0, digit: 0}
	s[2] = &Stock{item: i[1], loc: l[3], bal: 0, digit: 0}
	s[3] = &Stock{item: i[2], loc: l[1], bal: 0, digit: 0}
	s[4] = &Stock{item: i[2], loc: l[2], bal: 0, digit: 0}

	t := make([]*Trans, 10)
	t[0] = &Trans{ID: 1, item: i[1], locOut: l[0], locIn: l[1], qty: 100}
	t[1] = &Trans{ID: 2, item: i[1], locOut: l[0], locIn: l[1], qty: 10}
	t[2] = &Trans{ID: 3, item: i[1], locOut: l[0], locIn: l[1], qty: 20}
	t[3] = &Trans{ID: 4, item: i[1], locOut: l[1], locIn: l[2], qty: 50}
	t[4] = &Trans{ID: 5, item: i[1], locOut: l[2], locIn: l[3], qty: 10}
	t[5] = &Trans{ID: 6, item: i[1], locOut: l[1], locIn: l[3], qty: 10}
	t[6] = &Trans{ID: 7, item: i[1], locOut: l[1], locIn: l[4], qty: 10}
	t[7] = &Trans{ID: 8, item: i[1], locOut: l[3], locIn: l[5], qty: 1}
	t[8] = &Trans{ID: 9, item: i[1], locOut: l[4], locIn: l[1], qty: 1}
	t[9] = &Trans{ID: 10, item: i[1], locOut: l[3], locIn: l[1], qty: 1}

	return c, i, l, s, t
}
