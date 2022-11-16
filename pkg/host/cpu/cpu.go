package cpu

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetInfo() (models.Cpu, error) {
	cpuCounts, err := cpu.Counts(true)
	if err != nil {
		return models.Cpu{}, err
	}

	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return models.Cpu{}, err
	}

	return models.Cpu{
		Counts:  cpuCounts,
		Percent: cpuPercent[0],
	}, nil
}
