package host

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetMemoryInfo(getInfo host.Host) models.Memory {
	memoryInfo := models.Memory{
		Total:       models.TypeUint64(getInfo.Memory.Total),
		Used:        models.TypeUint64(getInfo.Memory.Used),
		Free:        models.TypeUint64(getInfo.Memory.Free),
		UsedPercent: models.TypeFloat64(getInfo.Memory.UsedPercent),
	}

	return memoryInfo
}
