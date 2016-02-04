package models

import (
	"fmt"
	"testing"
)

func TestPrintModelsPerson(t *testing.T) {
	pr, ps, cs, _, emp := setupPerson()
	for x := range pr {
		fmt.Println("Period:", pr[x].Year, "/", pr[x].Month)
	}
	for y := range cs {
		fmt.Println("Customer:", cs[y])
	}
	for p := range ps {
		fmt.Println("Person:", ps[p])
	}
	for e := range emp {
		fmt.Println("Employee:", emp[e].Fname, emp[e].Title, emp[e].Bdate.Format("01/02/2006"), emp[e].salary)
	}
	//	for z := range cr {
	//		fmt.Println("Customr Rank:", cr[z].Period, cr[z].Customer.Name, cr[z].Rank, cr[z].KI1Continous, cr[z].KI2PaymentDue, cr[z].KI3Responsibility, cr[z].KI4Charactor )
	//	}
}

func TestPrintModelsItem(t *testing.T) {
	cat, item, loc, stock, trans := setupItem()
	for c := range cat {
		fmt.Println("Category:", cat[c])
	}
	for i := range item {
		fmt.Println("Items:", item[i])
	}
	for l := range loc {
		fmt.Println("Location:", loc[l])
	}
	for s := range stock {
		fmt.Println("Stock:", stock[s].item.Name, stock[s].loc.Code, stock[s].bal, stock[s].digit)
	}
	for t := range trans {
		fmt.Println("Transaction:", trans[t].ID, trans[t].item.Name, trans[t].locIn.Code, trans[t].locOut.Code, trans[t].qty)
	}
}

func TestStock_CalcStockBalanceFromTrans(t *testing.T) {
	_, items, _, stocks, trans := setupItem()

	i := items[1]
	fmt.Println("A Stock:", stocks)
	a := i.Stock(stocks)
	fmt.Println("B Stock:", i.Stock(a))

	i.Calc(stocks, trans)
	for _, s := range a {
		fmt.Println(">>", s.item.Name, s.loc.Code, "Balance=", s.bal)
	}
}
