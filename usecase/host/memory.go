package host

import (
	"github.com/CharVstack/CharV-backend/openapi/v1"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetMemoryInfo(getInfo host.Host) openapi.Memory {
	memoryInfo := openapi.Memory{
		Total:       openapi.TypeUint64(getInfo.Memory.Total),
		Used:        openapi.TypeUint64(getInfo.Memory.Used),
		Free:        openapi.TypeUint64(getInfo.Memory.Free),
		UsedPercent: openapi.TypeFloat64(getInfo.Memory.UsedPercent),
	}

	return memoryInfo
}
