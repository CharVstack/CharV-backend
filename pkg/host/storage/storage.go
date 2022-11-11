package storage

import (
	"github.com/CharVstack/CharV-backend/domain/models"
)

func GetInfo(getInfo models.Host) []models.StoragePool {
	var storageInfoPointer = make([]*models.StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := models.StoragePool{
			Name:      pool.Name,
			TotalSize: pool.TotalSize,
			UsedSize:  pool.UsedSize,
			Path:      pool.Path,
			Status:    pool.Status,
		}
		storageInfoPointer[getInfoIndex] = &getStoragePool
	}

	var storageInfo = make([]models.StoragePool, len(storageInfoPointer))

	for getInfoIndex, pool := range storageInfoPointer {
		storageInfo[getInfoIndex] = *pool
	}

	return storageInfo
}
