package models

import "github.com/CharVstack/CharV-backend/entity"

type HostUseCase interface {
	Get() (entity.Host, error)
}
