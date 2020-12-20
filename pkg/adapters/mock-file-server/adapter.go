package mockfileserver

import (
	"errors"
	"github.com/google/uuid"
)

type FileServer struct {

}

func (fs *FileServer) DeleteFile(uuid.UUID) error {
	return errors.New("Error mocked")
}
