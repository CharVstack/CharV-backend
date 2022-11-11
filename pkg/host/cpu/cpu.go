package cpu

import (
	"github.com/CharVstack/CharV-backend/domain/models"
)

func GetInfo(getInfo models.Host) models.Cpu {
	cpuInfo := models.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: getInfo.Cpu.Percent,
	}

	return cpuInfo
}
