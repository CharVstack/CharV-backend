package entity

import "github.com/google/uuid"

// VmCore contains the most important part of Vm struct type
type VmCore struct {
	// Cpu is amount of VM's logical core
	Cpu int `json:"cpu"`
	// Memory
	Memory uint64 `json:"memory"`
	// Name
	Name string `json:"name"`
}

// Vm defines model for VM
type Vm struct {
	VmCore
	// ID specifies unique VM and is represented by UUID version 4.
	ID             uuid.UUID          `json:"id"`
	Devices        Devices            `json:"devices"`
	Boot           BootDevice         `json:"boot"`
	Virtualization VirtualizationType `json:"virtualization"`
	Daemonize      bool               `json:"daemonize"`
}

// Devices defines model for Devices.
type Devices struct {
	OS   Disk   `json:"os"`
	Disk []Disk `json:"disk"`
}

// Disk defines model for Disk.
type Disk struct {
	Device DiskDevice `json:"device"`
	Name   string     `json:"name"`
	Pool   string     `json:"pool"`
	Type   DiskType   `json:"type"`
}

// DiskDevice defines model for Disk.Device.
type DiskDevice string

// DiskType defines model for Disk.Type.
type DiskType string

// BootDevice defines model for Boot.Device
type BootDevice string

// VirtualizationType defines model for Virtualization.Type
type VirtualizationType string

// Defines values for DiskDevice.
const (
	DiskDeviceCdrom   DiskDevice = "cdrom"
	DiskDeviceDisk    DiskDevice = "disk"
	DiskDeviceUnknown DiskDevice = "unknown"
)

// Defines values for DiskType.
const (
	DiskTypeQcow2   DiskType = "qcow2"
	DiskTypeRaw     DiskType = "raw"
	DiskTypeIso     DiskType = "iso"
	DiskTypeUnknown DiskType = "unknown"
)

// Defines values for BootDevice.
const (
	BootDeviceCdrom   BootDevice = "cdrom"
	BootDeviceDisk    BootDevice = "disk"
	BootDeviceUnknown BootDevice = "unknown"
)

// Defines values for VirtualizationType
const (
	VirtualizationTypeKvm     VirtualizationType = "kvm"
	VirtualizationTypeQemu    VirtualizationType = "qemu"
	VirtualizationTypeUnknown VirtualizationType = "unknown"
)
