package domain

import "github.com/CharVstack/CharV-backend/adapters"

type GetApiV1Host200Response struct {
	Cpu          adapters.Cpu           `json:"cpu"`
	Mem          adapters.Memory        `json:"mem"`
	StoragePools []adapters.StoragePool `json:"storage_pools"`
}
