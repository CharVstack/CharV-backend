package qemu

import (
	"net/http"
	"path/filepath"

	"github.com/gamoutatsumi/go-vncproxy"
)

func NewVNCProxy(vmId string, logger vncproxy.Logger, socketsDir string, logLevel uint32) *vncproxy.Proxy {
	return vncproxy.New(&vncproxy.Config{
		Logger: logger,
		TokenHandler: func(r *http.Request) (addr, mode string, err error) {
			return filepath.Join(socketsDir, "vnc-"+vmId+".sock"), "unix", nil
		},
		LogLevel: logLevel,
	})
}
