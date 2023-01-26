package models

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/google/uuid"
)

type State string

// State of VM power
const (
	RUNNING  State = "RUNNING"
	SHUTDOWN State = "SHUTDOWN"
	UNKNOWN  State = "UNKNOWN"
)

// VmUseCase provides VM control
type VmUseCase interface {
	Create(req entity.VmCore) (entity.Vm, error)
	ReadAll() ([]entity.Vm, error)
	ReadById(id uuid.UUID) (entity.Vm, error)
	Update(id uuid.UUID, vm entity.VmCore) (entity.Vm, error)
	Delete(id uuid.UUID) error

	// Power
	GetPower(id uuid.UUID) State
	Start(id uuid.UUID) error
	Restart(id uuid.UUID) error
	Shutdown(id uuid.UUID) error
}

// VmDataAccess provides access to VM's structure files
type VmDataAccess interface {
	Browse() ([]entity.Vm, error)
	Read(id uuid.UUID) (entity.Vm, error)
	Edit(id uuid.UUID, vm entity.VmCore) (entity.Vm, error)
	Add(vm entity.Vm) error
	Delete(id uuid.UUID) error
}
