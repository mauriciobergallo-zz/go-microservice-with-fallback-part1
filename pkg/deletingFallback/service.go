package deletingFallback

import (
	"github.com/google/uuid"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/listing"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/logging"
)

type Repository interface {
	GetUsersMarkedForDeletion() ([]listing.User, error)
	DeleteUser(uuid.UUID) error
}

type FileServerAdapter interface {
	DeleteFile(uuid.UUID) error
}

type Service interface {
	RemoveUsersFallback() error
}

type service struct {
	r  Repository
	fs FileServerAdapter
	l  logging.Service
}

func NewService(r Repository, fs FileServerAdapter, l logging.Service) Service {
	return &service{r, fs, l}
}

func (s *service) RemoveUsersFallback() error {
	s.l.Debug("DeletingFallbackService", "RemoveUsersFallback() Called")
	uf, err := s.r.GetUsersMarkedForDeletion()
	if err != nil {
		s.l.Error("DeletingFallbackService", err.Error())
		return err
	}

	for _, u := range uf {
		err = s.fs.DeleteFile(u.ID)
		if err != nil {
			s.l.Error("DeletingFallbackService", err.Error())
			// We will not do anything.
			continue
		}

		err = s.r.DeleteUser(u.ID)
		if err != nil {
			s.l.Error("DeletingFallbackService", err.Error())
			return err
		}
	}

	return nil
}
