package sales

type Repository interface {
	CreateQuotation(req *NewQuoTemplate) (interface{}, error)
	SearchQuoById(req *SearchByIdTemplate) (interface{}, error)
	SearchQuoByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	ConfirmQuotation(req *NewQuoTemplate) (interface{}, error)
	CancelQuotation(req *NewQuoTemplate) (interface{}, error)
	QuotationToSaleOrder(req *SearchByIdTemplate) (interface{}, error)

	CreateSaleOrder(req *NewSaleTemplate) (interface{}, error)
	SearchSaleOrderById(req *SearchByIdTemplate) (interface{}, error)
	SearchSaleOrderByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	CreateDeposit(req *NewDepositTemplate) (interface{}, error)
	SearchDepositById(req *SearchByIdTemplate) (interface{}, error)
	SearchDepositByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchReserveToDeposit(req *SearchByKeywordTemplate) (interface{}, error)

	CreateInvoice(req *NewInvoiceTemplate) (interface{}, error)
	SearchInvoiceById(req *SearchByIdTemplate) (interface{}, error)
	SearchInvoiceByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}
