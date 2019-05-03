package cloudprint

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	CloudPrint(req *CloudPrintRequest) (interface{}, error)
}

func (s *service) CloudPrint(req *CloudPrintRequest) (interface{}, error) {
	resp, err := s.repo.CloudPrint(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
