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
	CreateQuotation(req *NewQuoTemplate) (interface{}, error)
	SearchQueById(req *SearchByIdTemplate) (interface{}, error)
	CreateSaleOrder(req *NewSaleTemplate) (interface{}, error)
	SearchSaleOrderById(req *SearchByIdTemplate) (interface{}, error)
	SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	CreateDeposit(req *NewDepositTemplate) (interface{}, error)
	SearchDepositById(req *SearchByIdTemplate) (interface{}, error)
	SearchDepositByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchReserveToDeposit(req *SearchByKeywordTemplate) (interface{}, error)
	CreateInvoice(req *NewInvoiceTemplate) (interface{}, error)
	SearchInvoiceById(req *SearchByIdTemplate) (interface{}, error)
}

func (s *service) CreateQuotation(req *NewQuoTemplate) (interface{}, error) {
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
		return nil, errors.New("Quotation not have salecode")
	case sum_item_amount != req.SumOfItemAmount:
		return nil, errors.New("ItemAmountSub not equa SumOfItemAmount")
	}

	resp, err := s.repo.CreateQuotation(req)
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

func (s *service) CreateSaleOrder(req *NewSaleTemplate) (interface{}, error) {
	var count_item int
	var count_item_qty int
	var count_item_unit int
	var sum_item_amount float64

	fmt.Println("Service Sale")
	var item_discount_amount_sub float64
	var err error
	for _, sub_item := range req.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
			if sub_item.DiscountWord != "" {
				item_discount_amount_sub, err = strconv.ParseFloat(sub_item.DiscountWord, 64)
				if err != nil {
					fmt.Println(err)
				}
			} else {
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
		return nil, errors.New("เอกสารไม่ได้ระบุ ลูกค้า")
	case count_item == 0:
		return nil, errors.New("เอกสารไม่มีรายการสินค้า")
	case req.SumOfItemAmount == 0:
		return nil, errors.New("เอกสารไม่มีมูลค่าสินค้า")
	case count_item_qty > 0:
		return nil, errors.New("รายการสินค้าไม่ได้ระบุ จำนวน")
	case count_item_unit > 0:
		return nil, errors.New("รายการสินค้าไม่ได้ระบุ หน่วยนับ")
	case req.SaleCode == "":
		return nil, errors.New("เอกสารไม่ได้ระบุ พนักงานขาย")
	case sum_item_amount != req.SumOfItemAmount:
		return nil, errors.New("มูลค่ารวมรายการสินค้าไม่เท่ากับมูลค่าสินค้า")
	}

	resp, err := s.repo.CreateSaleOrder(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchSaleOrderById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchSaleOrderById(req)
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

func (s *service) CreateDeposit(req *NewDepositTemplate) (interface{}, error) {
	var sum_pay_all float64

	sum_pay_all = req.CashAmount + req.CreditcardAmount + req.ChqAmount + req.BankAmount
	switch {
	case req.ArId == 0:
		return nil, errors.New("เอกสารไม่ได้ระบุ ลูกค้า")
	case req.TotalAmount == 0:
		return nil, errors.New("มูลค่าเอกสาร = 0")
	case req.SaleId == 0:
		return nil, errors.New("เอกสารไม่ได้ระบุ พนักงานขาย")
	case req.CashAmount == 0 && req.CreditcardAmount ==0 && req.ChqAmount == 0 && req.BankAmount == 0:
		return nil, errors.New("เอกสารไม่ได้ระบุยอดชำระ")
	case sum_pay_all != req.TotalAmount:
		return nil,errors.New("ยอดชำระไม่เท่ากับมูลค่าเอกสาร")
	}


	resp, err := s.repo.CreateDeposit(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchDepositById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchDepositById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchDepositByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchDepositByKeyword(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchReserveToDeposit(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchReserveToDeposit(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) CreateInvoice(req *NewInvoiceTemplate) (interface{}, error) {
	var sum_pay_all float64

	sum_pay_all = req.SumCashAmount + req.SumCreditAmount + req.SumChqAmount + req.SumBankAmount+req.SumOfDeposit+req.CouponAmount
	switch {
	case req.ArId == 0:
		return nil, errors.New("เอกสารไม่ได้ระบุ ลูกค้า")
	case req.TotalAmount == 0:
		return nil, errors.New("มูลค่าเอกสาร = 0")
	case req.SaleId == 0:
		return nil, errors.New("เอกสารไม่ได้ระบุ พนักงานขาย")
	case req.BillType == 0 && req.SumCashAmount == 0 && req.SumCreditAmount ==0 && req.SumChqAmount == 0 && req.SumBankAmount == 0 && req.SumOfDeposit == 0 && req.CouponAmount==0:
		return nil, errors.New("เอกสารไม่ได้ระบุยอดชำระ")
	case req.BillType == 0 && sum_pay_all != req.TotalAmount:
		return nil,errors.New("ยอดชำระไม่เท่ากับมูลค่าเอกสาร")
	}


	resp, err := s.repo.CreateInvoice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchInvoiceById(req *SearchByIdTemplate) (interface{}, error){
	resp, err := s.repo.SearchInvoiceById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}