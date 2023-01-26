package utils

import (
	"github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/google/uuid"
)

type id struct{}

func NewID() models.ID {
	return &id{}
}

func (i id) GenID() (uuid.UUID, error) {
	return uuid.NewRandom() // UUID Version 4
}
