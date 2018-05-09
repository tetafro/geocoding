package places

import (
	"fmt"
)

// Service implements business logic for places.
type Service struct {
	repo Repo
}

// NewService creates new service.
func NewService(repo Repo) *Service {
	return &Service{repo}
}

// Get find places by criterea.
func (s *Service) Get(criterea *Place) ([]*Place, error) {
	if criterea.Fullname != nil {
		return s.repo.GetByFullname(*criterea.Fullname)
	}
	return nil, fmt.Errorf("wrong search criterea")
}
