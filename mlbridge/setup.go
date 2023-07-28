
package mlbridge

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/caddyserver/caddy"
)

func init() {
	caddy.RegisterPlugin("mlbridge", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})