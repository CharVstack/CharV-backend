package host

import (
	"github.com/CharVstack/CharV-backend/adapters"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetStorageInfo(getInfo host.Host) []adapters.StoragePool {
	var storageInfoPointer = make([]*adapters.StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := adapters.StoragePool{
			Name:      pool.Name,
			TotalSize: adapters.TypeUint64(pool.TotalSize),
			UsedSize:  adapters.TypeUint64(pool.UsedSize),
			Path:      pool.Path,
			Status:    string(pool.Status),
		}
		storageInfoPointer[getInfoIndex] = &getStoragePool
	}

	var storageInfo = make([]adapters.StoragePool, len(storageInfoPointer))

	for getInfoIndex, pool := range storageInfoPointer {
		storageInfo[getInfoIndex] = *pool
	}

	return storageInfo
}
