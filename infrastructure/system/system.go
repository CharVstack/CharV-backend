package system

import (
	"os"
	"strings"
)

type target int

const (
	IMAGES target = iota
	GUESTS
	StoragePools
	QMP
	VNC
)

// No need to check for the existence of paths.
// Checks are in `main.go` .

type Paths struct {
	Images       string
	Guests       string
	StoragePools string
	QMP          string
	VNC          string
}

func (s Paths) Search(dir target, name string) bool {
	path := s.switchTarget(dir)

	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}

	for _, file := range files {
		if strings.Contains(file.Name(), name) {
			return true
		}
	}

	return false
}

func (s Paths) List(t target) ([]string, error) {
	path := s.switchTarget(t)

	files, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	names := make([]string, len(files))
	for i, file := range files {
		names[i] = file.Name()
	}

	return names, nil
}

func (s Paths) switchTarget(t target) string {
	switch t {
	case IMAGES:
		return s.Images
	case GUESTS:
		return s.Guests
	case StoragePools:
		return s.StoragePools
	case QMP:
		return s.QMP
	case VNC:
		return s.VNC
	}

	// NOT reach here or bug
	return ""
}
