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

type VmUseCase interface {
	Create(req entity.VmCore) (entity.Vm, error)
	ReadAll() ([]entity.Vm, error)
	ReadById(id uuid.UUID) (entity.Vm, error)
	Update(id uuid.UUID, vm entity.Vm) (entity.Vm, error)
	Delete(id uuid.UUID) error

	// Power
	GetPower(id uuid.UUID) State
	Start(id uuid.UUID) error
	Restart(id uuid.UUID) error
	Shutdown(id uuid.UUID) error
}
