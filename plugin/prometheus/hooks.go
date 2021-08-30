package prometheus

import (
	"github.com/lab5e/gmqtt/server"
)

func (p *Prometheus) HookWrapper() server.HookWrapper {
	return server.HookWrapper{}
}
