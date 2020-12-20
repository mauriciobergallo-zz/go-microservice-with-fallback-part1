package updating

import (
	"github.com/google/uuid"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/listing"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/logging"
)

type Repository interface {
	GetById(uuid.UUID) (listing.User, error)
	UpdateUser(User, bool) (User, error)
}

type Service interface {
	UpdateUser(User) (User, error)
}

type service struct {
	r Repository
	l logging.Service
}

func NewService (r Repository, l logging.Service) Service {
	return &service{r, l}
}

func (s *service) UpdateUser(u User) (User, error) {
	s.l.Debug("UpdatingService", "UpdateUser() Called")

	_, err := s.r.GetById(u.ID)
	if err != nil {
		return User{}, err
	}

	uu, err := s.r.UpdateUser(u, false)
	if err != nil {
		return User{}, err
	}

	return uu, nil
}
