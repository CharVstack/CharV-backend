package domain

import "github.com/CharVstack/CharV-backend/domain/models"

type GetApiV1Host200Response struct {
	Cpu          models.Cpu           `json:"cpu"`
	Mem          models.Memory        `json:"mem"`
	StoragePools []models.StoragePool `json:"storage_pools"`
}
