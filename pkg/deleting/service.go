package deleting

import (
	"github.com/google/uuid"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/listing"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/logging"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/updating"
)

type Repository interface {
	GetById(uuid.UUID) (listing.User, error)
	UpdateUser(updating.User, bool) (updating.User, error)
	DeleteUser(uuid.UUID) error
}

type FileServerAdapter interface {
	DeleteFile(uuid.UUID) error
}

type Service interface {
	RemoveUser(uuid.UUID) error
}

type service struct {
	r  Repository
	fs FileServerAdapter
	l  logging.Service
}

func NewService(r Repository, fs FileServerAdapter, l logging.Service) Service {
	return &service{r, fs, l}
}

func (s *service) RemoveUser(uid uuid.UUID) error {
	s.l.Debug("DeletingService", "RemoveUser() Called")
	u, err := s.r.GetById(uid)
	if err != nil {
		return err
	}

	// Delete from FileServer
	err = s.fs.DeleteFile(uid)
	if err != nil {
		s.l.Info("DeletingService", "RemoveUser() Can´t delete the image from FileServer, I will mark the entity for " +
			"being deleted in the fallback")

		uu := updating.User{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			UserName:  u.UserName,
			CountryId: u.CountryId,
			EMail:     u.EMail,
			BirthDate: u.BirthDate,
		}

		_, err = s.r.UpdateUser(uu, true)
		if err != nil {
			return err
		}

		return nil
	}

	err = s.r.DeleteUser(uid)
	if err != nil {
		return err
	}

	return nil
}
