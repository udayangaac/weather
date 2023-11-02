package maps

import (
	"net"
	"net/http"
	"time"
)

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
