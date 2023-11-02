package maps

import (
	"net"
	"net/http"
	"time"
)

// DefaultClient creates and returns an HTTP client with custom settings.
func DefaultClient(timeoutSeconds time.Duration) *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: timeoutSeconds * time.Second,
		}).Dial,
		DisableCompression:  false,
		TLSHandshakeTimeout: timeoutSeconds * time.Second,
	}
	return &http.Client{
		Timeout:   time.Second * timeoutSeconds,
		Transport: netTransport,
	}
}
