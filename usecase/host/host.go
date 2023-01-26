package host

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/usecase/models"
)

type hostUseCase struct {
	hostStatAccess models.HostStatAccess
}

func NewHostUseCase(h *models.HostStatAccess) models.HostUseCase {
	return &hostUseCase{
		hostStatAccess: *h,
	}
}

func (h hostUseCase) Get() (entity.Host, error) {
	c, err := h.hostStatAccess.GetCpu()
	if err != nil {
		return entity.Host{}, err
	}

	m, err := h.hostStatAccess.GetMem()
	if err != nil {
		return entity.Host{}, err
	}

	p, err := h.hostStatAccess.GetStoragePools()
	if err != nil {
		return entity.Host{}, err
	}

	return entity.Host{
		Cpu:          c,
		Memory:       m,
		StoragePools: p,
	}, nil
}
