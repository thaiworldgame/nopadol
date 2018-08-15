package print

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	PosDriveThruSlip(req *PosDriveThruSlipRequestTemplate) (interface{}, error)
	PosSlip(req *PosSlipRequestTemplate) (interface{}, error)
}

func (s *service) PosSlip(req *PosSlipRequestTemplate) (interface{}, error) {
	resp, err := s.repo.PosSlip(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) PosDriveThruSlip(req *PosDriveThruSlipRequestTemplate) (interface{}, error) {
	resp, err := s.repo.PosDriveThruSlip(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
