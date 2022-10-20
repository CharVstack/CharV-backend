package host

import (
	backendDomain "github.com/CharVstack/CharV-backend/domain"
	"github.com/CharVstack/CharV-lib/domain"
)

func GetHostInfo(getInfo domain.Host) backendDomain.GetApiV1Host200Response {

	hostInfo := backendDomain.GetApiV1Host200Response{
		Cpu:          GetCpuInfo(getInfo),
		Mem:          GetMemoryInfo(getInfo),
		StoragePools: GetStorageInfo(getInfo),
	}

	return hostInfo
}
