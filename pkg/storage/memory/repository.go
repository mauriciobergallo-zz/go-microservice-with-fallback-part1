package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/adding"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/listing"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/storage"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/updating"
)

type Storage struct {
	users []storage.User
}

func (m *Storage) InsertUser(u adding.User) (adding.User, error) {
	u.ID = uuid.New()
	newUser := storage.User{
		ID:                    u.ID,
		FirstName:             u.FirstName,
		LastNam:               u.LastName,
		UserName:              u.UserName,
		CountryId:             u.CountryId,
		EMail:                 u.EMail,
		BirthDate:             u.BirthDate,
		AvatarShouldBeDeleted: false,
	}

	m.users = append(m.users, newUser)

	return u, nil
}

func (m *Storage) GetById(id uuid.UUID) (listing.User, error) {
	for _, u := range m.users {
		if u.ID == id && u.AvatarShouldBeDeleted == false {
			return listing.User{
				ID:        u.ID,
				FirstName: u.FirstName,
				LastName:  u.LastNam,
				UserName:  u.UserName,
				CountryId: u.CountryId,
				EMail:     u.EMail,
				BirthDate: u.BirthDate,
			}, nil
		}
	}

	return listing.User{}, errors.New("user not found")
}

func (m *Storage) DeleteUser(id uuid.UUID) error {
	for i, u := range m.users {
		if u.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)

			return nil
		}
	}

	return errors.New("user not found")
}

func (m *Storage) UpdateUser(u updating.User, avatarShouldBeDeleted bool) (updating.User, error) {
	for i := range m.users {
		if m.users[i].ID == u.ID {
			m.users[i].FirstName = u.FirstName
			m.users[i].LastNam = u.LastName
			m.users[i].UserName = u.UserName
			m.users[i].CountryId = u.CountryId
			m.users[i].EMail = u.EMail
			m.users[i].BirthDate = u.BirthDate
			m.users[i].AvatarShouldBeDeleted = avatarShouldBeDeleted

			return u, nil
		}
	}

	return u, errors.New("user not found")
}

func (m *Storage) GetUsersMarkedForDeletion() ([]listing.User, error) {
	var users []listing.User
	for i, u := range m.users {
		if u.AvatarShouldBeDeleted == true {
			found := listing.User{
				ID:        m.users[i].ID,
				FirstName: m.users[i].FirstName,
				LastName:  m.users[i].LastNam,
				UserName:  m.users[i].UserName,
				CountryId: m.users[i].CountryId,
				EMail:     m.users[i].EMail,
				BirthDate: m.users[i].BirthDate,
			}
			users = append(users, found)
		}
	}

	return users, nil
}
