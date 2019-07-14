package rdns

import (
	"crypto/tls"
	"fmt"

	"github.com/miekg/dns"
)

// DoTClient is a DNS-over-TLS resolver.
type DoTClient struct {
	endpoint string
	pipeline *Pipeline
}

// DoTClientOptions contains options used by the DNS-over-TLS resolver.
type DoTClientOptions struct {
	TLSConfig *tls.Config
}

var _ Resolver = &DoTClient{}

// NewDoTClient instantiates a new DNS-over-TLS resolver.
func NewDoTClient(endpoint string, opt DoTClientOptions) *DoTClient {
	client := &dns.Client{
		Net:       "tcp-tls",
		TLSConfig: opt.TLSConfig,
	}
	return &DoTClient{
		endpoint: endpoint,
		pipeline: NewPipeline(endpoint, client),
	}
}

// Resolve a DNS query.
func (d *DoTClient) Resolve(q *dns.Msg, ci ClientInfo) (*dns.Msg, error) {
	return d.pipeline.Resolve(q)
}

func (d *DoTClient) String() string {
	return fmt.Sprintf("DoT(%s)", d.endpoint)
}