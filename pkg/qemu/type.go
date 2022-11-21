package qemu

import "github.com/google/uuid"

type InstallOpts struct {
	Name       string
	Memory     uint64
	VCpu       int
	Image      string
	Disk       string
	Id         uuid.UUID
	SocketPath string
}

type StartOpts struct {
	Disk string
}
