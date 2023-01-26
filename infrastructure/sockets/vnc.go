package sockets

import (
	"os"
	"path/filepath"

	"github.com/CharVstack/CharV-backend/infrastructure/system"
	usecase "github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/google/uuid"
)

type vncSocket struct {
	path system.Paths
}

func NewVNCSocket(p system.Paths) usecase.Socket {
	return &vncSocket{
		path: p,
	}
}

func (s vncSocket) Create(id uuid.UUID) (any, error) {
	return nil, nil
}

func (s vncSocket) List() ([]string, error) {
	files, err := os.ReadDir(s.path.QMP)
	if err != nil {
		return nil, err
	}

	l := make([]string, len(files))
	for i, f := range files {
		if len(f.Name()) != 41 { // uuid + .json
			continue
		}

		l[i] = f.Name()[:36]
	}

	return l, nil
}

func (s vncSocket) SearchFor(id uuid.UUID) bool {
	l, _ := s.List()
	for _, s := range l {
		if id.String() == s {
			return true
		}
	}

	return false
}

func (s vncSocket) Connect() error {
	return nil
}

func (s vncSocket) Send(data []byte) error {
	return nil
}

func (s vncSocket) Delete(id uuid.UUID) error {
	path := filepath.Join(s.path.VNC, id.String()+".sock")

	if _, err := os.Stat(path); err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
