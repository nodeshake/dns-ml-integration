// Package mlbridge implements a plugin
package mlbridge

import (
	"context"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"strings"
	"net/http"
	"strconv"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

// Mlbridge is a plugin in CoreDNS
type Mlbridge struct{
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (p Mlbridge) ServeDNS(ctx context.Context, w dns.