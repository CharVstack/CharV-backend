package util

import (
	"errors"

	"github.com/CharVstack/CharV-backend/pkg/host/memory"
)

func ExistsSufficientMemory(guestMemory uint64) error {
	hostMemory, err := memory.GetInfo()
	if err != nil {
		return err
	}

	if hostMemory.Free/(1024*1024) <= guestMemory {
		return errors.New("err: vm cannot be created because the memory specified for the guest is larger than the free memory of the host.")
	}
	return nil
}
