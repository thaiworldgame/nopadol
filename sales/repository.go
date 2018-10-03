package sales

type Repository interface {
	CreateQuo(req *NewQuoTemplate) (interface{}, error)
	SearchQuoById (req *SearchByIdTemplate) (interface{}, error)
	CreateSale(req *NewSaleTemplate) (interface{}, error)
	SearchSaleById (req *SearchByIdTemplate) (interface{}, error)
	SearchDocByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}