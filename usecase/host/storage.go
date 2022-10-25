package host

import (
	backendModels "github.com/CharVstack/CharV-backend/domain/models"
	libModels "github.com/CharVstack/CharV-lib/domain/models"
)

func GetStorageInfo(getInfo libModels.Host) []backendModels.StoragePool {
	var storageInfoPointer = make([]*backendModels.StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := backendModels.StoragePool{
			Name:      pool.Name,
			TotalSize: backendModels.TypeUint64(pool.TotalSize),
			UsedSize:  backendModels.TypeUint64(pool.UsedSize),
			Path:      pool.Path,
			Status:    string(pool.Status),
		}
		storageInfoPointer[getInfoIndex] = &getStoragePool
	}

	var storageInfo = make([]backendModels.StoragePool, len(storageInfoPointer))

	for getInfoIndex, pool := range storageInfoPointer {
		storageInfo[getInfoIndex] = *pool
	}

	return storageInfo
}
