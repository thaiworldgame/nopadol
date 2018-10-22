package sales

import (
	"fmt"
	"strconv"
	"errors"
)

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	CreateQuo(req *NewQuoTemplate) (interface{}, error)
	SearchQueById(req *SearchByIdTemplate) (interface{}, error)
	CreateSale(req *NewSaleTemplate) (interface{}, error)
	SearchSaleById(req *SearchByIdTemplate) (interface{}, error)
	SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}

func (s *service) CreateQuo(req *NewQuoTemplate) (interface{}, error) {
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64
	var err error


	fmt.Println("Service Quo")
	for _, sub_item := range req.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1

			sum_item_amount = sum_item_amount + (sub_item.Qty * (sub_item.Price - sub_item.DiscountAmount))
		}
		if (sub_item.ItemCode != "" && sub_item.Qty == 0) {
			count_item_qty = count_item_qty + 1
		}
		if (sub_item.ItemCode != "" && sub_item.UnitCode == "") {
			count_item_unit = count_item_unit + 1
		}
	}

	fmt.Println("Count Item", count_item)

	switch {
	case req.ArCode == "":
		return nil, errors.New("Arcode is null")
	case count_item == 0:
		return nil, errors.New("Docno is not have item")
	case req.SumOfItemAmount == 0:
		return nil, errors.New("SumOfItemAmount = 0")
	case count_item_qty > 0:
		return nil, errors.New("Item not have qty")
	case count_item_unit > 0:
		return nil, errors.New("Item not have unitcode")
	case req.SaleCode == "":
		return nil,errors.New("Quotation not have salecode")
	case sum_item_amount != req.SumOfItemAmount:
		return nil,errors.New("ItemAmountSub not equa SumOfItemAmount")
	}

	resp, err := s.repo.CreateQuo(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchQueById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchQuoById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) CreateSale(req *NewSaleTemplate) (interface{}, error) {
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64


	fmt.Println("Service Sale")
	var item_discount_amount_sub float64
	var  err error
	for _, sub_item := range req.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
			if sub_item.DiscountWord != "" {
				item_discount_amount_sub, err = strconv.ParseFloat(sub_item.DiscountWord, 64)
				if err != nil {
					fmt.Println(err)
				}
			}else{
				item_discount_amount_sub = 0
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

	switch {
	case req.ArCode == "":
		return nil, errors.New("Arcode is null")
	case count_item == 0:
		return nil, errors.New("Docno is not have item")
	case req.SumOfItemAmount == 0:
		return nil, errors.New("SumOfItemAmount = 0")
	case count_item_qty > 0:
		return nil, errors.New("Item not have qty")
	case count_item_unit > 0:
		return nil, errors.New("Item not have unitcode")
	case req.SaleCode == "":
		return nil,errors.New("Quotation not have salecode")
	case sum_item_amount != req.SumOfItemAmount:
		return nil,errors.New("ItemAmountSub not equa SumOfItemAmount")
	}

	resp, err := s.repo.CreateSale(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchSaleById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchSaleById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchDocByKeyword(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
