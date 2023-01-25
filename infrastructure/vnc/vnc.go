package vnc

import (
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"net/http"
	"path/filepath"

	"github.com/gamoutatsumi/go-vncproxy"
)

func NewVNCProxy(vmId string, logger vncproxy.Logger, path system.Paths, logLevel uint32) *vncproxy.Proxy {
	return vncproxy.New(&vncproxy.Config{
		Logger: logger,
		TokenHandler: func(r *http.Request) (addr, mode string, err error) {
			return filepath.Join(path.VNC, vmId+".sock"), "unix", nil
		},
		LogLevel: logLevel,
	})
}
