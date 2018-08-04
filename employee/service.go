package employee

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}
type Service interface {
	//	SearchEmployeeById(ctx context.Context, req *SearchByIdTemplate) (emp EmployeeTemplate, err error)
	SearchById(req *SearchByIdTemplate) (interface{}, error)
	SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}

func (s *service) SearchById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchById(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) SearchByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchByKeyword(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
