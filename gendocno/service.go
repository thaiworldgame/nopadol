package gendocno

func New(repo Repository) (Service) {
	return &service{repo}
}

type service struct {
	repo Repository
}

type Service interface {
	Gen(req *DocNoTemplate) (string, error)
}

func (s *service) Gen(req *DocNoTemplate) (string, error) {
	resp, err := s.repo.Gen(req)
	if err != nil {
		return "", err
	}
	return resp, nil
}
