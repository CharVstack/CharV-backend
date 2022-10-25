package host

import (
	backendDomain "github.com/CharVstack/CharV-backend/domain"
	libModels "github.com/CharVstack/CharV-lib/domain/models"
)

func GetHostInfo(getInfo libModels.Host) backendDomain.GetApiV1Host200Response {

	hostInfo := backendDomain.GetApiV1Host200Response{
		Cpu:          GetCpuInfo(getInfo),
		Mem:          GetMemoryInfo(getInfo),
		StoragePools: GetStorageInfo(getInfo),
	}

	return hostInfo
}
