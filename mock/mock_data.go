package mock

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

func setupPerson() ([]*Period, []*Person, []*Customer, []*CusRank, []*Title, []*Employee, []*Org) {
	pr := make([]*Period, 12)
	year := 2015
	for i, v := range pr {
		month := i+1
		v = &Period{Year: year, Month: month }
	}

	ps := make([]*Person, 5)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	ps[0] = &Person{FirstName: "เกษม", LastName: "อานนทวิลาศ", NickName: "Tom", BirthDate: time.Date(1974, time.October, 4, 0, 0, 0, 0, loc)}
	ps[1] = &Person{FirstName: "จิราภรณ์", LastName: "อานนทวิลาศ", NickName: "Jip", BirthDate: time.Date(1976, time.August, 12, 0, 0, 0, 0, time.Local)}
	ps[2] = &Person{FirstName: "ทนัฐพร", LastName: "อานนทวิลาศ", NickName: "Tim", BirthDate: time.Date(1976, time.August, 31, 0, 0, 0, 0, time.Local)}
	ps[3] = &Person{FirstName: "ธนันท์", LastName: "อานนทวิลาศ", NickName: "Tam", BirthDate: time.Date(1974, time.June, 7, 0, 0, 0, 0, time.Local)}
	ps[4] = &Person{FirstName: "สาธิต", LastName: "โฉมวัฒนา", NickName: "น้อย", BirthDate: time.Date(1974, time.June, 10, 0, 0, 0, 0, time.Local)}

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
	tt[1] = &Title{Parent: tt[0], ID: 1, TH: "รองประธานกรรมการ", EN: "Director"}
	tt[2] = &Title{Parent: tt[1], ID: 2, TH: "กรรมการผู้จัดการ", EN: "Managing Director"}
	tt[3] = &Title{Parent: tt[2], ID: 3, TH: "ผู้อำนวยการขาย", EN: "Sales Director"}
	tt[4] = &Title{Parent: tt[2], ID: 4, TH: "ผู้อำนวยการบริหารสินค้า", EN: "Merchandise Director"}
	tt[5] = &Title{Parent: tt[3], ID: 5, TH: "พนักงานขาย", EN: "Sales"}
	fmt.Println("Title")


	o := make([]*Org, 20)
	o[0] = &Org{TH: "สำนักกรรมการผู้จัดการ", EN: "Managing Director Office", Short: "MD"}
	o[1] = &Org{Parent: o[0], TH: "ขาย", EN: "Sales", Short: "SA"}
	o[2] = &Org{Parent: o[1], TH: "หน้าร้านสาขา 1", EN: "Retail Store 1", Short: "S1"}
	o[3] = &Org{Parent: o[1], TH: "หน้าร้านสาขา 2", EN: "Retail Store 2", Short: "S2"}
	o[4] = &Org{Parent: o[1], TH: "โครงการ 1", EN: "Project1", Short: "PJ1"}
	o[5] = &Org{Parent: o[1], TH: "โครงการ 2", EN: "Project2", Short: "PJ2"}
	o[6] = &Org{Parent: o[1], TH: "ค้าส่ง", EN: "Wholesales", Short: "WS"}
	o[7] = &Org{Parent: o[0], TH: "บริหารสินค้า", EN: "Merchandising", Short: "MC"}
	o[8] = &Org{Parent: o[7], TH: "แผนกสินค้า1", EN: "Category1", Short: "CAT1"}
	o[9] = &Org{Parent: o[7], TH: "แผนกสินค้า2", EN: "Category2", Short: "CAT2"}
	o[10] = &Org{Parent: o[7], TH: "แผนกสินค้า1", EN: "Category1", Short: "CAT3"}
	o[11] = &Org{Parent: o[7], TH: "แผนกสินค้า4", EN: "Category4", Short: "CAT4"}
	o[12] = &Org{Parent: o[7], TH: "การตลาด", EN: "Marketing", Short: "CAT4"}
	o[13] = &Org{Parent: o[7], TH: "พัฒนาระบบคุณภาพ", EN: "Quality Management Development", Short: "QMD"}
	o[14] = &Org{Parent: o[7], TH: "พัฒนาระบบสาขา", EN: "Business Development", Short: "QMD"}
	o[15] = &Org{Parent: o[7], TH: "พัฒนาระบบบริการลูกค้า", EN: "Customer Service Development", Short: "CSD"}
	o[16] = &Org{Parent: o[7], TH: "Information Technology", EN: "Information Technology", Short: "IT"}

	emp := make([]*Employee, 5)
	emp[0] = &Employee{Person: ps[0], PersonID: 1, Titles: []*Title{tt[2], tt[3]}, Code: "39001", salary: decimal.New(40000, 0)}
	//emp[1] = &Employee{Person: ps[1], PersonID: 2, Titles: &Title{tt[3]}, Code: "48001", salary: decimal.New(30000, 0)}
	//emp[2] = &Employee{Person: ps[2], PersonID: 3, Titles: &Title{tt[4]}, Code: "49001", salary: decimal.New(20000, 0)}
	//emp[3] = &Employee{Person: ps[3], PersonID: 4, Titles: &Title{tt[4]}, Code: "50001", salary: decimal.New(20000, 0)}
	//emp[4] = &Employee{Person: ps[4], PersonID: 5, Titles: &Title{tt[5]}, Code: "53001", salary: decimal.New(9000, 0)}
	fmt.Println("Employee")

	return pr, ps, c, cr, tt, emp, o
}

func setupItem() ([]*Category, []*Item, []*Location, []*Stock, []*Tran) {
	c := make([]*Category, 5)
	c[0] = &Category{Name: "root"}
	c[1] = &Category{Parent: c[0], Name: "Beverage"}
	c[2] = &Category{Parent: c[0], Name: "Snack"}
	c[3] = &Category{Parent: c[0], Name: "Fruit"}
	c[4] = &Category{Parent: c[1], Name: "Soda"}

	i := make([]*Item, 10)
	i[0] = &Item{Parent: c[0], Name: "Dummy Item"}
	i[1] = &Item{Parent: c[4], Name: "Coke"}
	i[2] = &Item{Parent: c[4], Name: "Pepsi"}
	i[3] = &Item{Parent: c[1], Name: "อิชิตันส้ม รสข้าวคั่ว"}
	i[4] = &Item{Parent: c[1], Name: "อิชิตันมะนาว"}
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
	s[0] = &Stock{item: i[1], loc: l[1], bal: 0}
	s[1] = &Stock{item: i[1], loc: l[2], bal: 0}
	s[2] = &Stock{item: i[2], loc: l[1], bal: 0}
	s[3] = &Stock{item: i[2], loc: l[2], bal: 0}
	s[4] = &Stock{item: i[2], loc: l[3], bal: 0}

	t := make([]*Tran, 4)
	t[0] = &Tran{ID: 1, item: i[1], locOut: l[0], locIn: l[1], qty: 100}
	t[1] = &Tran{ID: 2, item: i[1], locOut: l[1], locIn: l[2], qty: 10}
	t[2] = &Tran{ID: 3, item: i[1], locOut: l[1], locIn: l[3], qty: 20}
	t[3] = &Tran{ID: 4, item: i[1], locOut: l[1], locIn: l[4], qty: 50}
	//	t[4] = &Trans{ID: 5, item: i[1], locOut: l[4], locIn: l[1], qty: 10}
	//	t[5] = &Trans{ID: 6, item: i[1], locOut: l[1], locIn: l[3], qty: 10}
	//	t[6] = &Trans{ID: 7, item: i[1], locOut: l[1], locIn: l[4], qty: 10}
	//	t[7] = &Trans{ID: 8, item: i[1], locOut: l[3], locIn: l[5], qty: 1}
	//	t[8] = &Trans{ID: 9, item: i[1], locOut: l[4], locIn: l[1], qty: 1}
	//	t[9] = &Trans{ID: 10, item: i[1], locOut: l[3], locIn: l[1], qty: 1}

	return c, i, l, s, t
}
