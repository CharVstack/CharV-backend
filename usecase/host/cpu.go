package host

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-lib/domain"
)

func GetCpuInfo(getInfo domain.Host) models.Cpu {
	cpuInfo := models.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: models.TypeFloat64(getInfo.Cpu.Percent),
	}

	return cpuInfo
}
