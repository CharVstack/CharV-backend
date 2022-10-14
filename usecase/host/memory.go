package host

import (
	"github.com/CharVstack/CharV-backend/adapters"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetMemoryInfo(getInfo host.Host) adapters.Memory {
	memoryInfo := adapters.Memory{
		Total:       adapters.TypeUint64(getInfo.Memory.Total),
		Used:        adapters.TypeUint64(getInfo.Memory.Used),
		Free:        adapters.TypeUint64(getInfo.Memory.Free),
		UsedPercent: adapters.TypeFloat64(getInfo.Memory.UsedPercent),
	}

	return memoryInfo
}
