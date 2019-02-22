package sales

type Repository interface {
	CreateQuotation(req *NewQuoTemplate) (interface{}, error)
	SearchQuoById(req *SearchByIdTemplate) (interface{}, error)
	ConfirmQuotation(req *NewQuoTemplate) (interface{}, error)
	CancelQuotation(req *NewQuoTemplate) (interface{}, error)
	CreateSaleOrder(req *NewSaleTemplate) (interface{}, error)
	SearchSaleOrderById(req *SearchByIdTemplate) (interface{}, error)
	SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	CreateDeposit(req *NewDepositTemplate) (interface{}, error)
	SearchDepositById(req *SearchByIdTemplate) (interface{}, error)
	SearchDepositByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchReserveToDeposit(req *SearchByKeywordTemplate) (interface{}, error)
	CreateInvoice(req *NewInvoiceTemplate) (interface{}, error)
	SearchInvoiceById(req *SearchByIdTemplate) (interface{}, error)
	SearchInvoiceByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchSaleByItem(req *SearchByItemTemplate) (interface{}, error)
	SearchHisByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	CancelInvoice(req *NewInvoiceTemplate) (interface{}, error)
}
