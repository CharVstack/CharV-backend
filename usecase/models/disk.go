package models

import "github.com/google/uuid"

// Disk of VM
type Disk interface {
	Create(id uuid.UUID) error
	Delete(id uuid.UUID) error
}
