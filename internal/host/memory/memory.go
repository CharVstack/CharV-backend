package memory

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetInfo() (models.Memory, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return models.Memory{}, err
	}

	memTotal := memInfo.Total
	memUsed := memInfo.Used
	memFree := memInfo.Free
	memUsedPercent := memInfo.UsedPercent

	return models.Memory{
		Total:       memTotal,
		Used:        memUsed,
		Free:        memFree,
		UsedPercent: memUsedPercent,
	}, nil
}
