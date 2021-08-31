package lmqtt

import (
	"net"

	"github.com/lab5e/lmqtt/config"
)

type Options func(srv *server)

// WithConfig set the config of the server
func WithConfig(config config.Config) Options {
	return func(srv *server) {
		srv.config = config
	}
}

// WithTCPListener set  tcp listener(s) of the server. Default listen on  :1883.
func WithTCPListener(lns ...net.Listener) Options {
	return func(srv *server) {
		srv.tcpListener = append(srv.tcpListener, lns...)
	}
}

// WithHook set hooks of the server. Notice: WithPlugin() will overwrite hooks.
func WithHook(hooks Hooks) Options {
	return func(srv *server) {
		srv.hooks = hooks
	}
}
