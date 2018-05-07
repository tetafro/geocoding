package places

// Service implements business logic for places.
type Service struct {
	repo Repo
}

// NewService creates new service.
func NewService(repo Repo) *Service {
	return &Service{repo}
}

// GetByFullname find places by fullname (or by part of fullname).
func (s *Service) GetByFullname(fullname string) ([]*Place, error) {
	return s.repo.GetByFullname(fullname)
}
