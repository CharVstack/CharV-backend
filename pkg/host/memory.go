package host

import (
	"github.com/CharVstack/CharV-backend/domain/models"
)

func GetMemoryInfo(getInfo models.Host) models.Memory {
	memoryInfo := models.Memory{
		Total:       getInfo.Mem.Total,
		Used:        getInfo.Mem.Used,
		Free:        getInfo.Mem.Free,
		UsedPercent: getInfo.Mem.UsedPercent,
	}

	return memoryInfo
}