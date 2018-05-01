package places

// Service implements business logic for places.
type Service struct {
	repo Repo
}

// NewService creates new service.
func NewService(repo Repo) *Service {
	return &Service{repo}
}

// GetByName find places by name (or by part of name).
func (s *Service) GetByName(name string) ([]*Place, error) {
	return s.repo.GetByName(name)
}
