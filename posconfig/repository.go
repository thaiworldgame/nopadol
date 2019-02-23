package posconfig

type Repository interface {
	Create(req *PosConfigTemplate) (interface{}, error)
	SearchById () (interface{}, error)
}
