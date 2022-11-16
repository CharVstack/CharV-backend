package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/CharVstack/CharV-backend/domain/models"
)

func IsPoolExists(path string) bool {
	f, _ := os.Stat(path)
	return f.IsDir()
}

func GetPoolFiles(storageDir string) ([]string, error) {
	entries, err := os.ReadDir(storageDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files, nil
}

func GetPoolInfo(file string, storageDir string) (models.StoragePool, error) {
	path, _ := filepath.Abs(filepath.Join(storageDir, file))
	bytes, err := os.ReadFile(path)
	if err != nil {
		return models.StoragePool{}, err
	}

	var storagePool models.StoragePool
	err = json.Unmarshal(bytes, &storagePool)
	if err != nil {
		return models.StoragePool{}, err
	}
	return storagePool, nil
}
