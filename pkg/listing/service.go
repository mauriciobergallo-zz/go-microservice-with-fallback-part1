package listing

import (
	"github.com/google/uuid"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/logging"
)

type Repository interface {
	GetById(uuid.UUID) (User, error)
}

type Service interface {
	ObtainUserById(uuid.UUID) (User, error)
}

type service struct {
	r Repository
	l logging.Service
}

func NewService(r Repository, l logging.Service) Service {
	return &service{r, l}
}

func (s *service) ObtainUserById(id uuid.UUID) (User, error) {
	s.l.Debug("ListingService", "ObtainUserById() Called")
	return User{}, nil
}