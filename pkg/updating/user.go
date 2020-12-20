package updating

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:id`
	FirstName string    `json:firstName`
	LastName  string    `json:lastName`
	UserName  string    `json:userName`
	CountryId string    `json:countryId`
	EMail     string    `json:email`
	BirthDate time.Time `json:birthDate`
}
