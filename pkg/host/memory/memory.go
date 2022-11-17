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

	return models.Memory{
		Total:   memInfo.Total,
		Used:    memInfo.Used,
		Free:    memInfo.Free,
		Percent: memInfo.UsedPercent,
	}, nil
}
