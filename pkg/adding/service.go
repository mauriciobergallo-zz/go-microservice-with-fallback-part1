package adding

import (
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/logging"
)

type Repository interface {
  InsertUser(User) (User, error)
}

type Service interface {
  AddUser(User) (User, error)
}

type service struct {
  r Repository
  l logging.Service
}

func NewService( r Repository, l logging.Service) Service {
  return &service{r, l}
}

func (s *service) AddUser(u User) (User, error) {
  s.l.Debug("AddingService", "AddUser() Called")

  iu, err := s.r.InsertUser(u)
  if err != nil {
    return User{}, err
  }

  return iu, nil
}
