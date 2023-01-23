package entity

// Host defines model for Host.
type Host struct {
	// Cpu is host's CPU
	Cpu Cpu `json:"cpu"`

	// Memory is host's Memory
	Memory Memory `json:"memory"`

	// StoragePools is host's storage pools
	StoragePools []StoragePool `json:"storage_pools"`
}

// Cpu provides CPU info
//
// This struct depends on gopsutil
// detail: https://godocs.io/github.com/shirou/gopsutil/v3/cpu
type Cpu struct {
	// Counts returns the number of logical cores in the system
	Counts int `json:"counts"`

	// Percent calculates the percentage of used either per CPU or combined
	Percent float64 `json:"percent"`
}

// Memory provides Memory info
//
// This struct depends on gopsutil
// detail: https://godocs.io/github.com/shirou/gopsutil/v3/mem
type Memory struct {
	// This is the kernel's notion of free
	Free uint64 `json:"free"`

	// Percentage of RAM used by programs
	Percent float64 `json:"percent"`

	// Total amount of RAM on this system
	Total uint64 `json:"total"`

	// RAM used by programs
	Used uint64 `json:"used"`
}

// StoragePool provides storage pools info
//
// This struct depends on gopsutil
// detail: https://godocs.io/github.com/shirou/gopsutil/v3/disk
type StoragePool struct {
	// Name of storage pool
	Name string `json:"name"`

	// Path to storage pool
	Path string `json:"path"`

	// Status of storage pool
	Status StoragePoolStatus `json:"status"`

	// Total amount of storage pool
	TotalSize uint64 `json:"total_size"`

	// Used amount of storage pool
	UsedSize uint64 `json:"used_size"`
}

// StoragePoolStatus defines model for StoragePool.Status.
type StoragePoolStatus string

// Defines values for StoragePoolStatus.
const (
	ACTIVE StoragePoolStatus = "ACTIVE"
	ERROR  StoragePoolStatus = "ERROR"
)
