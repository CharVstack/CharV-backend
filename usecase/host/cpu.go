package host

import (
	"github.com/CharVstack/CharV-backend/openapi/v1"
	"github.com/CharVstack/CharV-lib/pkg/host"
)

func GetCpuInfo(getInfo host.Host) openapi.Cpu {
	cpuInfo := openapi.Cpu{
		Counts:  getInfo.Cpu.Counts,
		Percent: openapi.TypeFloat64(getInfo.Cpu.Percent),
	}

	return cpuInfo
}
