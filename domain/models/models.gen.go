package models

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for DiskDevice.
const (
	DiskDeviceCdrom   DiskDevice = "cdrom"
	DiskDeviceDisk    DiskDevice = "disk"
	DiskDeviceUnknown DiskDevice = "unknown"
)

// Defines values for DiskType.
const (
	DiskTypeQcow2   DiskType = "qcow2"
	DiskTypeUnknown DiskType = "unknown"
)

// Defines values for StoragePoolStatus.
const (
	StoragePoolStatusActive StoragePoolStatus = "Active"
	StoragePoolStatusError  StoragePoolStatus = "Error"
)

// Defines values for VmStatus.
const (
	VmStatusActive  VmStatus = "active"
	VmStatusError   VmStatus = "error"
	VmStatusPending VmStatus = "pending"
	VmStatusUnknown VmStatus = "unknown"
)

// Defines values for VmPowerInfoState.
const (
	POWEREDOFF VmPowerInfoState = "POWERED_OFF"
	POWEREDON  VmPowerInfoState = "POWERED_ON"
	SUSPENDED  VmPowerInfoState = "SUSPENDED"
)

// Defines values for PostApiV1VmsVmIdPowerActionParamsAction.
const (
	Reset   PostApiV1VmsVmIdPowerActionParamsAction = "reset"
	Start   PostApiV1VmsVmIdPowerActionParamsAction = "start"
	Stop    PostApiV1VmsVmIdPowerActionParamsAction = "stop"
	Suspend PostApiV1VmsVmIdPowerActionParamsAction = "suspend"
)

// Cpu ホストのCPU情報
type Cpu struct {
	Counts  int     `json:"counts"`
	Percent float64 `json:"percent"`
}

// Devices defines model for Devices.
type Devices struct {
	Disk []Disk `json:"disk"`
}

// Disk defines model for Disk.
type Disk struct {
	Device DiskDevice `json:"device"`
	Path   string     `json:"path"`
	Type   DiskType   `json:"type"`
}

// DiskDevice defines model for Disk.Device.
type DiskDevice string

// DiskType defines model for Disk.Type.
type DiskType string

// Host defines model for Host.
type Host struct {
	// Cpu ホストのCPU情報
	Cpu Cpu `json:"cpu"`

	// Mem ホストのメモリ情報
	Mem          Memory        `json:"mem"`
	StoragePools []StoragePool `json:"storage_pools"`
}

// Memory ホストのメモリ情報
type Memory struct {
	Free        uint64  `json:"free"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

// Metadata defines model for Metadata.
type Metadata struct {
	ApiVersion string             `json:"api_version"`
	Id         openapi_types.UUID `json:"id"`
}

// StoragePool ホストが持つストレージプールの情報
type StoragePool struct {
	Name      string            `json:"name"`
	Path      string            `json:"path"`
	Status    StoragePoolStatus `json:"status"`
	TotalSize uint64            `json:"total_size"`
	UsedSize  uint64            `json:"used_size"`
}

// StoragePoolStatus defines model for StoragePool.Status.
type StoragePoolStatus string

// Vm 仮想マシンを表すモデル
type Vm struct {
	Devices  Devices  `json:"devices"`
	Memory   int      `json:"memory"`
	Metadata Metadata `json:"metadata"`
	Name     string   `json:"name"`
	Status   VmStatus `json:"status"`
	Vcpu     int      `json:"vcpu"`
}

// VmStatus defines model for Vm.Status.
type VmStatus string

// VmPowerInfo defines model for VmPowerInfo.
type VmPowerInfo struct {
	CleanPowerOff bool             `json:"clean_power_off"`
	State         VmPowerInfoState `json:"state"`
}

// VmPowerInfoState defines model for VmPowerInfo.State.
type VmPowerInfoState string

// GetAllVMsList200Response defines model for GetAllVMsList200Response.
type GetAllVMsList200Response struct {
	Data    []Vm    `json:"data"`
	Message *string `json:"message,omitempty"`
}

// GetHost200Response defines model for GetHost200Response.
type GetHost200Response struct {
	Data    Host    `json:"data"`
	Message *string `json:"message,omitempty"`
}

// GetVMByVMId200Response defines model for GetVMByVMId200Response.
type GetVMByVMId200Response struct {
	// Data 仮想マシンを表すモデル
	Data    Vm      `json:"data"`
	Message *string `json:"message,omitempty"`
}

// GetVMPowerByVMId200Response defines model for GetVMPowerByVMId200Response.
type GetVMPowerByVMId200Response struct {
	Data    VmPowerInfo `json:"data"`
	Message *string     `json:"message,omitempty"`
}

// PatchUpdateVMByVMId200Response defines model for PatchUpdateVMByVMId200Response.
type PatchUpdateVMByVMId200Response struct {
	// Data 仮想マシンを表すモデル
	Data    Vm      `json:"data"`
	Message *string `json:"message,omitempty"`
}

// PostCreateNewVM200Response defines model for PostCreateNewVM200Response.
type PostCreateNewVM200Response struct {
	// Data 仮想マシンを表すモデル
	Data    Vm      `json:"data"`
	Message *string `json:"message,omitempty"`
}

// PostApiV1VmsJSONBody defines parameters for PostApiV1Vms.
type PostApiV1VmsJSONBody struct {
	Memory int    `json:"memory"`
	Name   string `json:"name"`
	Vcpu   int    `json:"vcpu"`
}

// PatchApiV1VmsVmIdJSONBody defines parameters for PatchApiV1VmsVmId.
type PatchApiV1VmsVmIdJSONBody struct {
	Memory *int    `json:"memory,omitempty"`
	Name   *string `json:"name,omitempty"`
	Vcpu   *int    `json:"vcpu,omitempty"`
}

// PostApiV1VmsVmIdPowerActionParams defines parameters for PostApiV1VmsVmIdPowerAction.
type PostApiV1VmsVmIdPowerActionParams struct {
	Action *PostApiV1VmsVmIdPowerActionParamsAction `form:"action,omitempty" json:"action,omitempty"`
}

// PostApiV1VmsVmIdPowerActionParamsAction defines parameters for PostApiV1VmsVmIdPowerAction.
type PostApiV1VmsVmIdPowerActionParamsAction string

// PostApiV1VmsJSONRequestBody defines body for PostApiV1Vms for application/json ContentType.
type PostApiV1VmsJSONRequestBody PostApiV1VmsJSONBody

// PatchApiV1VmsVmIdJSONRequestBody defines body for PatchApiV1VmsVmId for application/json ContentType.
type PatchApiV1VmsVmIdJSONRequestBody PatchApiV1VmsVmIdJSONBody
