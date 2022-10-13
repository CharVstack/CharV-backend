package domain

import "github.com/CharVstack/CharV-backend/openapi/v1"

type GetApiV1Host200Response struct {
	Cpu          openapi.Cpu           `json:"cpu"`
	Mem          openapi.Memory        `json:"mem"`
	StoragePools []openapi.StoragePool `json:"storage_pools"`
}
