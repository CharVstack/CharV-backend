package storage

import (
	"github.com/shirou/gopsutil/v3/disk"
)

func GetSize(path string) (size uint64, usedSize uint64, err error) {
	diskUsage, err := disk.Usage(path)
	if err != nil {
		return 0, 0, err
	}
	return diskUsage.Total, diskUsage.Used, err
}
