package print

type Repository interface {
	PosSlip(req *PosSlipRequestTemplate) (interface{}, error)
}
