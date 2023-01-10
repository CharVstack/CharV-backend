package qemu

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gamoutatsumi/go-vncproxy"
)

func NewVNCProxy(vmId string, logger vncproxy.Logger, socketsDir string) *vncproxy.Proxy {
	sockPath := append(strings.Split(socketsDir, "/"), "vnc-"+vmId+".sock")
	return vncproxy.New(&vncproxy.Config{
		Logger: logger,
		TokenHandler: func(r *http.Request) (addr, mode string, err error) {
			return filepath.Join(sockPath...), "unix", nil
		},
	})
}
