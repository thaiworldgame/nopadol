package pos

import (
	"fmt"
	"errors"
	"strconv"
)

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	Create(req *NewPosTemplate) (interface{}, error)
	SearchById(req *SearchPosByIdRequestTemplate) (interface{}, error)
}

func (s *service)Create(req *NewPosTemplate) (interface{}, error){
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var sum_pay_amount float64


	fmt.Println("Service")
	sum_pay_amount = (req.SumCashAmount + req.SumCreditAmount + req.SumChqAmount + req.SumBankAmount + req.CoupongAmount)

	for _, sub_item := range req.PosSubs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1

			item_discount_amount_sub, err := strconv.ParseFloat(sub_item.DiscountWord, 64)
			if err != nil {
				fmt.Println(err)
			}
			sum_item_amount = sum_item_amount + (sub_item.Qty * (sub_item.Price - item_discount_amount_sub))
		}
		if (sub_item.ItemCode != "" && sub_item.Qty == 0) {
			count_item_qty = count_item_qty + 1
		}
		if (sub_item.ItemCode != "" && sub_item.UnitCode == "") {
			count_item_unit = count_item_unit + 1
		}
	}


	fmt.Println( " sum_item_amount,sum_pay_amount",sum_item_amount,sum_pay_amount)

	switch {
	case req.ArCode == "":
		return nil, errors.New("Arcode is null")
	case count_item == 0:
		return nil, errors.New("Docno is not have item")
	case (req.SumCashAmount == 0 && req.SumCreditAmount == 0 && req.SumChqAmount == 0 && req.SumBankAmount == 0):
		return nil, errors.New("Docno not set money to another type payment")
	case req.SumOfItemAmount == 0:
		return nil, errors.New("Sumofitemamount = 0")
	case count_item_qty > 0:
		return nil, errors.New("Item not have qty")
	case count_item_unit > 0:
		return nil, errors.New("Item not have unitcode")
	case sum_pay_amount > req.TotalAmount:
		return nil, errors.New("Rec money is over totalamount")
	case sum_item_amount != sum_pay_amount:
		return nil, errors.New("Rec money is less than totalamount")
	case (req.MachineCode == "" || req.ShiftNo == 0 || req.ShiftCode == "" || req.CashierCode == ""):
		return nil, errors.New("Docno not have pos data")
	case req.SumChqAmount != 0 && len(req.ChqIns) == 0:
		return nil, errors.New("Docno not have chq data")
	case req.SumCreditAmount != 0 && len(req.CreditCards) == 0:
		return nil, errors.New("Docno not have credit card data")
	}

	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service)SearchById(req *SearchPosByIdRequestTemplate)(interface{}, error){
	resp, err := s.repo.SearchById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
