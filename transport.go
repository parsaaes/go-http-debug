package go_http_debug

import (
	"net/http"
	"net/http/httputil"
	"moul.io/http2curl"
)

// TransportWithDump wraps a transport to capture request/response dumps
type TransportWithDump struct {
	// The actual transport that gets wrapped. If it's nil, it will use http.DefaultTransport.
	RootTransport http.RoundTripper
	// The request and response dumps are returned using this function.
	Handler func(req string, resp string)
}

func (d *TransportWithDump) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := getTransport(d.RootTransport)

	if d.Handler == nil {
		return transport.RoundTrip(req)
	}

	var reqDump, respDump string

	dmp, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		reqDump = err.Error()
	} else {
		reqDump = string(dmp)
	}

	resp, rtErr := transport.RoundTrip(req)
	if rtErr == nil {
		dmp, err := httputil.DumpResponse(resp, true)
		if err != nil {
			respDump = err.Error()
		} else {
			respDump = string(dmp)
		}
	}

	d.Handler(reqDump, respDump)

	return resp, rtErr
}

// TransportWithCurl wraps a transport to translate its request to a curl command.
type TransportWithCurl struct {
	// The actual transport that gets wrapped. If it's nil, it will use http.DefaultTransport.
	RootTransport http.RoundTripper
	// The curl command is returned using this function.
	Handler func(cmd string)
}

func (d *TransportWithCurl) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := getTransport(d.RootTransport)

	if d.Handler == nil {
		return transport.RoundTrip(req)
	}

	cmd, err := http2curl.GetCurlCommand(req)
	if err != nil {
		d.Handler(err.Error())
	} else {
		d.Handler(cmd.String())
	}

	return transport.RoundTrip(req)
}

func getTransport(rt http.RoundTripper) http.RoundTripper {
	if rt == nil {
		return http.DefaultTransport
	}

	return rt
}