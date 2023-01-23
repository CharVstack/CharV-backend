package models

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/google/uuid"
)

type HostStatAccess interface {
	GetCpu() (entity.Cpu, error)
	GetMem() (entity.Memory, error)
	GetStoragePools() ([]entity.StoragePool, error)
}

type ID interface {
	GenID() (uuid.UUID, error)
}

type Command interface {
	Run(name string, args []string) error
}
