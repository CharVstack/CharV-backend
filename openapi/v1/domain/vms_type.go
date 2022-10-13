package domain

import "github.com/CharVstack/CharV-backend/openapi/v1"

// GetApiV1Vms200Response VMsの返却型
type GetApiV1Vms200Response struct {
	Vms []openapi.Vm `json:"vms"`
}
