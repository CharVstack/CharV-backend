package host

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-backend/pkg/host/cpu"
	"github.com/CharVstack/CharV-backend/pkg/host/memory"
	"github.com/CharVstack/CharV-backend/pkg/host/storage"
)

func GetInfo(opt GetInfoOptions) (models.Host, error) {
	cpuInfo, err := cpu.GetInfo()
	if err != nil {
		return models.Host{}, err
	}

	memoryInfo, err := memory.GetInfo()
	if err != nil {
		return models.Host{}, err
	}

	poolConfigPaths, err := storage.GetPoolFiles(opt.StorageDir)
	if err != nil {
		return models.Host{}, err
	}

	var storagePools []models.StoragePool
	for _, file := range poolConfigPaths {
		storagePoolInfo, err := storage.GetPoolInfo(file, opt.StorageDir)
		if err != nil {
			return models.Host{}, err
		}

		isExists := storage.IsPoolExists(storagePoolInfo.Path)

		if isExists {
			storagePoolInfo.Status = "Active"
			storagePoolInfo.TotalSize, storagePoolInfo.UsedSize, err = storage.GetSize(storagePoolInfo.Path)
			if err != nil {
				storagePoolInfo.Status = "Error"
			}
		} else {
			storagePoolInfo.Status = "Error"
		}

		storagePools = append(storagePools, storagePoolInfo)
	}
	if storagePools == nil {
		storagePools = []models.StoragePool{}
	}

	return models.Host{
		Cpu:          cpuInfo,
		Memory:       memoryInfo,
		StoragePools: storagePools,
	}, nil
}
