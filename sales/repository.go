package sales

type Repository interface {
	CreateQuotation(req *NewQuoTemplate) (interface{}, error)
	SearchQuoById(req *SearchByIdTemplate) (interface{}, error)
	CreateSaleOrder(req *NewSaleTemplate) (interface{}, error)
	SearchSaleById(req *SearchByIdTemplate) (interface{}, error)
	SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	CreateDeposit(req *NewDepositTemplate) (interface{}, error)
}
