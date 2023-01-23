package models

import "github.com/google/uuid"

type Socket interface {
	Create(id uuid.UUID) (any, error)
	List() ([]string, error)
	SearchFor(id uuid.UUID) bool
	Connect() error
	Send(data []byte) error
	Delete(id uuid.UUID) error
}
