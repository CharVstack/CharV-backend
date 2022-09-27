package openapi

type GetApiV1Host200Response struct {
	Cpu          Cpu           `json:"cpu"`
	Mem          Memory        `json:"mem"`
	StoragePools []StoragePool `json:"storage_pools"`
}
