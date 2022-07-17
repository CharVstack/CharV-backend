/*
 * CharVstack-API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/CharVstack/CharV-lib/host"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getStorageInfo(getInfo host.Host) []StoragePool {
	var storageInfoPointer = make([]*StoragePool, len(getInfo.StoragePools))

	for getInfoIndex, pool := range getInfo.StoragePools {
		getStoragePool := StoragePool{
			Name:      pool.Name,
			TotalSize: pool.TotalSize,
			UsedSize:  pool.UsedSize,
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
		Cpu: Cpu{
			Cpu:    getInfo.Cpu.Counts,
			CpuMhz: getInfo.Cpu.Percent,
		},
		Mem: Memory{
			Total:       getInfo.Memory.Total,
			Used:        getInfo.Memory.Used,
			Free:        getInfo.Memory.Free,
			UsedPercent: getInfo.Memory.UsedPercent,
		},
		StoragePools: getStorageInfo(getInfo),
	}

	return hostInfo
}

// GetApiV1Host - Get a host
func GetApiV1Host(c *gin.Context) {
	getInfo := host.GetInfo()
	hostInfo := GetHostInfo(getInfo)
	c.JSON(http.StatusOK, gin.H{
		"cpu":           hostInfo.Cpu,
		"mem":           hostInfo.Mem,
		"storage_pools": hostInfo.StoragePools,
	})
}