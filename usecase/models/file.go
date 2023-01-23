package models

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/google/uuid"
)

type StorageAccess interface {
	Browse() ([]Storage, error)
	Read(name string) (Storage, error)
}

type Storage struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type VmDataAccess interface {
	Browse() ([]entity.Vm, error)
	Read(id uuid.UUID) (entity.Vm, error)
	Edit(id uuid.UUID, vm entity.Vm) (entity.Vm, error)
	Add(vm entity.Vm) error
	Delete(id uuid.UUID) error
}
