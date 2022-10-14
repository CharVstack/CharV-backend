package host

import (
	"github.com/CharVstack/CharV-backend/openapi/v1"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetStorageInfo(getInfo host.Host) []openapi.StoragePool {
	var storageInfoPointer = make([]*openapi.StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := openapi.StoragePool{
			Name:      pool.Name,
			TotalSize: openapi.TypeUint64(pool.TotalSize),
			UsedSize:  openapi.TypeUint64(pool.UsedSize),
			Path:      pool.Path,
			Status:    string(pool.Status),
		}
		storageInfoPointer[getInfoIndex] = &getStoragePool
	}

	var storageInfo = make([]openapi.StoragePool, len(storageInfoPointer))

	for getInfoIndex, pool := range storageInfoPointer {
		storageInfo[getInfoIndex] = *pool
	}

	return storageInfo
}
