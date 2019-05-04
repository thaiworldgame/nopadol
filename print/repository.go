package print

type Repository interface {
	PosDriveThruSlip(req *PosDriveThruSlipRequestTemplate) (interface{}, error)
	PosSlip(req *PosSlipRequestTemplate) (interface{}, error)
	//AutoPrint(req *AutoPrintRequest) (interface{}, error)
}
