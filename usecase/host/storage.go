package host

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetStorageInfo(getInfo host.Host) []models.StoragePool {
	var storageInfoPointer = make([]*models.StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := models.StoragePool{
			Name:      pool.Name,
			TotalSize: models.TypeUint64(pool.TotalSize),
			UsedSize:  models.TypeUint64(pool.UsedSize),
			Path:      pool.Path,
			Status:    string(pool.Status),
		}
		storageInfoPointer[getInfoIndex] = &getStoragePool
	}

	var storageInfo = make([]models.StoragePool, len(storageInfoPointer))

	for getInfoIndex, pool := range storageInfoPointer {
		storageInfo[getInfoIndex] = *pool
	}

	return storageInfo
}
