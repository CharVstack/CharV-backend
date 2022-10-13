package host

import (
	"github.com/CharVstack/CharV-backend/domain"
	"github.com/CharVstack/CharV-lib/host"
)

func GetHostInfo(getInfo host.Host) domain.GetApiV1Host200Response {

	hostInfo := domain.GetApiV1Host200Response{
		Cpu:          GetCpuInfo(getInfo),
		Mem:          GetMemoryInfo(getInfo),
		StoragePools: GetStorageInfo(getInfo),
	}

	return hostInfo
}
