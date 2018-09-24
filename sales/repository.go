package sales

type Repository interface {
	CreateQuo(req *NewQuoTemplate) (interface{}, error)
	SearchQuoById () (interface{}, error)
	CreateSale(req *NewSaleTemplate) (interface{}, error)
	SearchSaleById () (interface{}, error)
}