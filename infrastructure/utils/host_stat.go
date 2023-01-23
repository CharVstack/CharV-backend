package utils

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
)

type hostStatAccess struct {
	storageAccess models.StorageAccess
}

func NewHostStatAccess(s *models.StorageAccess) models.HostStatAccess {
	return &hostStatAccess{
		storageAccess: *s,
	}
}

func (h hostStatAccess) GetCpu() (entity.Cpu, error) {
	c, err := cpu.Counts(true)
	if err != nil {
		return entity.Cpu{}, err
	}

	p, err := cpu.Percent(0, false)
	if err != nil {
		return entity.Cpu{}, err
	}

	return entity.Cpu{
		Counts:  c,
		Percent: p[0],
	}, nil
}

func (h hostStatAccess) GetMem() (entity.Memory, error) {
	m, err := mem.VirtualMemory()
	if err != nil {
		return entity.Memory{}, err
	}

	return entity.Memory{
		Total:   m.Total,
		Used:    m.Used,
		Free:    m.Free,
		Percent: m.UsedPercent,
	}, nil
}

func (h hostStatAccess) GetStoragePools() ([]entity.StoragePool, error) {
	storages, err := h.storageAccess.Browse()
	if err != nil {
		return []entity.StoragePool{}, err
	}

	pools := make([]entity.StoragePool, len(storages))

	for i, s := range storages {
		// status
		f, err := isExistPool(s.Path)
		if err != nil || f == false {
			pools[i] = entity.StoragePool{
				Name:   s.Name,
				Path:   s.Path,
				Status: entity.ERROR,
			}
			continue
		}

		// size
		d, err := disk.Usage(s.Path)
		if err != nil {
			pools[i] = entity.StoragePool{
				Name:   s.Name,
				Path:   s.Path,
				Status: entity.ERROR,
			}
			continue
		}

		pools[i] = entity.StoragePool{
			Name:      s.Name,
			Path:      s.Path,
			Status:    entity.ACTIVE,
			TotalSize: d.Total,
			UsedSize:  d.Used,
		}
	}

	return pools, nil
}

func isExistPool(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return f.IsDir(), err
}
