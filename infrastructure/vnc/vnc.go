package vnc

import (
	"net/http"
	"path/filepath"

	"github.com/gamoutatsumi/go-vncproxy"
)

func NewVNCProxy(vmId string, logger vncproxy.Logger, sockDir string, logLevel uint32) *vncproxy.Proxy {
	return vncproxy.New(&vncproxy.Config{
		Logger: logger,
		TokenHandler: func(r *http.Request) (addr, mode string, err error) {
			return filepath.Join(sockDir, vmId+".sock"), "unix", nil
		},
		LogLevel: logLevel,
	})
}
