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
func (p Mlbridge) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	qname := state.Name()
	ip := state.IP()
	jsonData := map[string]string{"Domain Name": qname, "IP": ip}
    jsonValue, _ := json.Marshal(jsonData)
    response, err := http.Post("http://127.0.0.1:5000/", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
        return plugin.NextOrFailure(p.Name(), p.Next, ctx, w, r)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
            prob, _ := strconv.ParseFloat(string(strings.Split(strings.Split(string(data), ":")[1], "\"")[1]), 8)
	    if prob < 0.5 {
	            fmt.Printf("Benign Domain: %s | Probability: %f\n", qname, (1 - prob))
	    } else {
	            fmt.Printf("Malicious Domain: %s | Probability: %f\n", qname, prob)
	    }
    }

    answers := []dns.RR{}

	if state.QType() != dns.TypeA {
		return dns.RcodeNameError, nil
	}

	rr := new(dns.A)
	rr.Hdr = dns.RR_Header{Name: qname, Rrtype: dns.TypeA, Class: dns.ClassI