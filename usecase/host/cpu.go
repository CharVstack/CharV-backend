package host

import (
	"github.com/CharVstack/CharV-lib/domain/models"
)

func GetCpuInfo(getInfo models.Host) models.Cpu {
	cpuInfo := models.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: getInfo.Cpu.Percent,
	}

	return cpuInfo
}
