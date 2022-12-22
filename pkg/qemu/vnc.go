package qemu

import (
	"net/http"
	"path/filepath"

	"github.com/gamoutatsumi/go-vncproxy"
)

func NewVNCProxy(vmId string) *vncproxy.Proxy {
	return vncproxy.New(&vncproxy.Config{
		LogLevel: vncproxy.InfoLevel,
		TokenHandler: func(r *http.Request) (addr, mode string, err error) {
			return filepath.Join("var", "run", "charvstack", "vnc-"+vmId+".sock"), "unix", nil
		},
	})
}
