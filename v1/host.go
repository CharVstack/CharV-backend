package openapi

import (
	"github.com/CharVstack/CharV-lib/host"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getCpuInfo(getInfo host.Host) Cpu {
	cpuInfo := Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: TypeFloat64(getInfo.Cpu.Percent),
	}

	return cpuInfo
}

func getMemoryInfo(getInfo host.Host) Memory {
	memoryInfo := Memory{
		Total:       TypeUint64(getInfo.Memory.Total),
		Used:        TypeUint64(getInfo.Memory.Used),
		Free:        TypeUint64(getInfo.Memory.Free),
		UsedPercent: TypeFloat64(getInfo.Memory.UsedPercent),
	}

	return memoryInfo
}

func getStorageInfo(getInfo host.Host) []StoragePool {
	var storageInfoPointer = make([]*StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := StoragePool{
			Name:      pool.Name,
			TotalSize: TypeUint64(pool.TotalSize),
			UsedSize:  TypeUint64(pool.UsedSize),
			Path:      pool.Path,
			Status:    string(pool.Status),
		}
		storageInfoPointer[getInfoIndex] = &getStoragePool
	}

	var storageInfo = make([]StoragePool, len(storageInfoPointer))

	for getInfoIndex, pool := range storageInfoPointer {
		storageInfo[getInfoIndex] = *pool
	}

	return storageInfo
}

func GetHostInfo(getInfo host.Host) GetApiV1Host200Response {

	hostInfo := GetApiV1Host200Response{
		Cpu:          getCpuInfo(getInfo),
		Mem:          getMemoryInfo(getInfo),
		StoragePools: getStorageInfo(getInfo),
	}

	return hostInfo
}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	getInfo := host.GetInfo()
	hostInfo := GetHostInfo(getInfo)
	c.JSON(http.StatusOK, gin.H{
		"cpu":           hostInfo.Cpu,
		"mem":           hostInfo.Mem,
		"storage_pools": hostInfo.StoragePools,
	})
}
