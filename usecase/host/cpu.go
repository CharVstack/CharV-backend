package host

import (
	backendModels "github.com/CharVstack/CharV-backend/domain/models"
	libModels "github.com/CharVstack/CharV-lib/domain/models"
)

func GetCpuInfo(getInfo libModels.Host) backendModels.Cpu {
	cpuInfo := backendModels.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: backendModels.TypeFloat64(getInfo.Cpu.Percent),
	}

	return cpuInfo
}
