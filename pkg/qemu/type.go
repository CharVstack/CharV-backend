package qemu

type InstallOpts struct {
	Name   string
	Memory uint64
	VCpu   uint64
	Image  string
	Disk   string
}

type StartOpts struct {
	Disk string
}
