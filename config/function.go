package config

import (
	"fmt"
	"math"
)

func CalcTaxItem(taxtype int, taxrate float64, afterdiscountamount float64) (beforetaxamount float64, taxamount float64, totalamount float64) {
	switch taxtype {
	case 0:
		beforetaxamount = toFixed(afterdiscountamount, 2)
		taxamount = toFixed(((afterdiscountamount*(100+float64(taxrate)))/(100))-afterdiscountamount, 2)
		totalamount = toFixed(beforetaxamount+taxamount, 2)
	case 1:
		taxamount = toFixed(afterdiscountamount-((afterdiscountamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = toFixed(afterdiscountamount-taxamount, 2)
		totalamount = toFixed(afterdiscountamount, 2)
	case 2:
		beforetaxamount = toFixed(afterdiscountamount, 2)
		taxamount = 0
		totalamount = toFixed(afterdiscountamount, 2)
	}

	fmt.Println("taxtype,taxrate,beforetaxamount,taxamount,totalamount", taxtype, taxrate, beforetaxamount, taxamount, totalamount)

	return beforetaxamount, taxamount, totalamount
}

func CalcTaxTotalAmount(taxtype int64, taxrate float64, totalamount float64) (beforetaxamount float64, taxamount float64) {
	switch taxtype {
	case 0:
		beforetaxamount = toFixed(totalamount, 2)
		taxamount = toFixed(((totalamount*(100+float64(taxrate)))/(100))-totalamount, 2)
	case 1:
		taxamount = toFixed(totalamount-((totalamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = toFixed(totalamount-taxamount, 2)
	case 2:
		beforetaxamount = toFixed(totalamount, 2)
		taxamount = 0
	}

	fmt.Println("taxtype,taxrate,beforetaxamount,taxamount,totalamount", taxtype, taxrate, beforetaxamount, taxamount)

	return beforetaxamount, taxamount
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
