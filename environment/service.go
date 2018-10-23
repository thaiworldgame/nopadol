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
	SearchProjectById(req *SearchByIdTemplate) (interface{}, error)
	SearchProjectByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
	SearchAllocateById(req *SearchByIdTemplate) (interface{}, error)
	SearchAllocateByKeyword(req *SearchByKeywordTemplate) (interface{}, error)
}

//Department
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

//Project
func (s *service) SearchProjectById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchProjectById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchProjectByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchProjectByKeyword(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//Allocate
func (s *service) SearchAllocateById(req *SearchByIdTemplate) (interface{}, error) {
	resp, err := s.repo.SearchAllocateById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchAllocateByKeyword(req *SearchByKeywordTemplate) (interface{}, error) {
	resp, err := s.repo.SearchAllocateByKeyword(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
