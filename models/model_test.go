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
		fmt.Println("Employee:", emp[e].FName, emp[e].Title, emp[e].BDate.Format("01/02/2006"), emp[e].salary)
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

func TestItem_CalcBalancefromStockTrans(t *testing.T) {
	_, items, l, stocks, trans := setupItem()

	i := items[1]
	fmt.Println("Start Stock Calc:", stocks)
	iStock := i.Calc(stocks, trans)
	fmt.Println("Check Assert balance:")
	for _, v := range iStock {
		if v.loc == l[0] && v.bal != -100 {
			t.Error("Missing Calculation in", v.item.Name, v.loc.Code, v.bal)
		}
		if v.loc == l[1] && v.bal != 20 {
			t.Error("Missing Calculation in", v.item.Name, v.loc.Code, v.bal)
		}
		if v.loc == l[2] && v.bal != 10 {
			t.Error("Missing Calculation in", v.item.Name, v.loc.Code, v.bal)
		}
		fmt.Println(">>", v.item.Name, v.loc.Code, "Balance=", v.bal)
	}
}

//func TestStock_CalcBalanceByItem(t *testing.T) {
//	_, items, _, stocks, trans := setupItem()
//
//	i := items[1]
//	for k, _ := range stocks {
//		if stocks[k].item == i {
//			err := stocks[k].Calc(trans)
//			if err != nil {
//				t.Error("Error in Calc(i)>> ",err)
//			}
//		}
//		if stocks[k].bal != 10 { //Expected value here!
//			t.Error("Calculate stock balance of", stocks[k].item.Name ,"at location=", stocks[k].loc.Code,"Expected 'bal' = 10 but = ", stocks[k].bal)
//		}
//	}
//}

func TestCat_FindByName(t *testing.T) {
	cats, _, _, _, _ := setupItem()
	c := FindByName(cats, "Bev")
	if c != cats[1] {
		t.Error("Expected return index 1 but =", c)
	}

	x := FindByName(cats, "Alcohol")
	if x != nil {
		t.Error("Expected no matched category but =", x)
	}
	fmt.Println("Pass! TestCat_FindByName return index =", c)
}
