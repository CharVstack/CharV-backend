package host

import (
	"github.com/CharVstack/CharV-backend/adapters"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetCpuInfo(getInfo host.Host) adapters.Cpu {
	cpuInfo := adapters.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: adapters.TypeFloat64(getInfo.Cpu.Percent),
	}

	return cpuInfo
}
