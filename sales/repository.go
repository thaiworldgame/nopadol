package sales

type Repository interface {
	Create(req *PosConfigTemplate1) (interface{}, error)
	SearchById () (interface{}, error)
}