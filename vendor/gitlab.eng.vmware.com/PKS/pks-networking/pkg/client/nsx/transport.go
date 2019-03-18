package nsx

import (
	rc "github.com/go-openapi/runtime/client"
	"net/http"
)

// nsxKeepAliveTransport is a wrapper on KeepAliveTransport
type nsxKeepAliveTransport struct {
	header             http.Header
	keepAliveTransport http.RoundTripper
	Transport          http.RoundTripper
}

func NsxKeepAliveTransport(rt http.RoundTripper) http.RoundTripper {
	return &nsxKeepAliveTransport{
		Transport:          rt,
		keepAliveTransport: rc.KeepAliveTransport(rt),
	}
}

func (k *nsxKeepAliveTransport) WithHeader(header http.Header) http.RoundTripper {
	k.header = header
	return k
}

func (k *nsxKeepAliveTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, v := range k.header {
		r.Header[k] = v
	}

	resp, err := k.keepAliveTransport.RoundTrip(r)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
