package storage

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                    uuid.UUID
	FirstName             string
	LastNam               string
	UserName              string
	CountryId             string
	EMail                 string
	BirthDate             time.Time
	AvatarShouldBeDeleted bool
}
