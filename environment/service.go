package environment


func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	SearchDepartmentById(req *SearchByIdTemplate) (interface{}, error)
	SearchDepartmentByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}

func (s *service) SearchDepartmentById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchDepartmentById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchDepartmentByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchDepartmentByKeyword(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
