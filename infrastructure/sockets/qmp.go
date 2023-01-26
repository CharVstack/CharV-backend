package sockets

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/interfaces/errors"
	usecase "github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/digitalocean/go-qemu/qmp"
	"github.com/google/uuid"
)

type qmpSocket struct {
	path system.Paths
}

func NewQMPSocket(p system.Paths) usecase.Socket {
	return qmpSocket{
		path: p,
	}
}

func (s qmpSocket) Create(id uuid.UUID) (any, error) {
	addr := filepath.Join(s.path.QMP, id.String()+".sock")
	sock, err := qmp.NewSocketMonitor("unix", addr, 2*time.Second)
	if err != nil {
		return nil, err
	}

	return sock, nil
}

func (s qmpSocket) List() ([]string, error) {
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

func (s qmpSocket) SearchFor(id uuid.UUID) bool {
	l, _ := s.List()
	for _, s := range l {
		if id.String() == s {
			return true
		}
	}

	return false
}

func (s qmpSocket) Connect() error {
	return nil
}

func (s qmpSocket) Send(data []byte) error {
	return nil
}

func (s qmpSocket) Delete(id uuid.UUID) error {
	path := filepath.Join(s.path.QMP, id.String()+".sock")
	if _, err := os.Stat(path); err != nil {
		return errors.NotFound.Wrap(err, fmt.Sprintf("%s is not found", path))
	}
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
