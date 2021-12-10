// Package myip looks up the host's IPv4 address.
package myip

import (
	"context"
	"errors"
	"net"
	"time"
)

var ErrNotFound = errors.New("Not found")

var defaultResolver = &net.Resolver{
	PreferGo: true,
	Dial:     dialGoogleNS,
}

func dialGoogleNS(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{Timeout: time.Second}
	return d.DialContext(ctx, network, "8.8.8.8:53")
}

// LookupV4 looks up the current WAN IPv4 using the Google resolver.
//
// On error, this might return a DNS lookup error as from net.LookupIP
// or ErrNotFound if the DNS server returned an empty reply.
func LookupV4(ctx context.Context) (net.IP, error) {
	ips, err := defaultResolver.LookupIP(ctx, "ip4", "o-o.myaddr.l.google.com")
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, ErrNotFound
	}
	return ips[0], nil
}
