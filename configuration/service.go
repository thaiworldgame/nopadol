package configuration

type Service interface {
	ConfigSetting(req *RequestSettingTemplate) (interface{}, error)
	SearchSettingById(req *SearchByIdRequestTemplate) (interface{}, error)
	SearchSettingByKeyword(req *SearchByKeywordRequestTemplate) (interface{}, error)
}

func New(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repo Repository
}

func (s *service) ConfigSetting(req *RequestSettingTemplate) (interface{}, error) {
	resp, err := s.repo.ConfigSetting(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchSettingById(req *SearchByIdRequestTemplate) (interface{}, error) {
	resp, err := s.repo.SearchSettingById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SearchSettingByKeyword(req *SearchByKeywordRequestTemplate) (interface{}, error) {
	resp, err := s.repo.SearchSettingByKeyword(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
