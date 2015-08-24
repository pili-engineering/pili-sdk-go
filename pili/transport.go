package pili

import (
	"net/http"
)

type Transport struct {
	creds     Credentials
	transport http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	token, err := t.creds.MACToken(req)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Qiniu "+token)
	return t.transport.RoundTrip(req)
}

func NewTransport(creds *Credentials, transport http.RoundTripper) *Transport {
	if transport == nil {
		transport = http.DefaultTransport
	}
	t := &Transport{transport: transport}
	t.creds = *creds
	return t
}
