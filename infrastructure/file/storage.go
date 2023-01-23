package file

import (
	"encoding/json"
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"os"
	"path/filepath"
)

type hostStorageAccess struct {
	path system.Paths
}

func NewStorageAccess(p system.Paths) models.StorageAccess {
	return &hostStorageAccess{
		path: p,
	}
}

func (h hostStorageAccess) Browse() ([]models.Storage, error) {
	files, err := os.ReadDir(h.path.StoragePools)
	if err != nil {
		return []models.Storage{}, err
	}

	storages := make([]models.Storage, len(files))

	for i, file := range files {
		storage, err := h.Read(file.Name())
		if err != nil {
			return []models.Storage{}, err
		}

		storages[i] = storage
	}

	return storages, nil
}

// Read StoragePool config
// name: xxx.json
func (h hostStorageAccess) Read(name string) (models.Storage, error) {
	var storage models.Storage

	// read storage config file
	path := filepath.Join(h.path.StoragePools + name)
	data, err := os.ReadFile(path) // path must be abs
	if err != nil {
		return models.Storage{}, err
	}

	err = json.Unmarshal(data, &storage)
	if err != nil {
		return models.Storage{}, err
	}

	return storage, nil
}
