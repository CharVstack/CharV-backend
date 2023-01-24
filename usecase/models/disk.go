package models

import "github.com/google/uuid"

// Disk of VM
type Disk interface {
	Create(pool string, id uuid.UUID) error
	Delete(pool string, id uuid.UUID) error
}
