package host

import (
	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetCpuInfo(getInfo host.Host) models.Cpu {
	cpuInfo := models.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: models.TypeFloat64(getInfo.Cpu.Percent),
	}

	return cpuInfo
}
