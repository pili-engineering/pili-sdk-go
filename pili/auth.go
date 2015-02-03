package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/qiniu/bytes/seekable"
	"io"
	"net/http"
)

// -----------------------------------------------------------------------------------------------------------

type Mac struct {
	AccessKey string
	SecretKey string
}

func (mac *Mac) SignRequest(req *http.Request, incbody bool) (token string, err error) {

	h := hmac.New(sha1.New, []byte(mac.SecretKey))

	u := req.URL
	data := u.Path
	if u.RawQuery != "" {
		data += "?" + u.RawQuery
	}
	io.WriteString(h, data+"\n")

	if incbody {
		s2, err2 := seekable.New(req)
		if err2 != nil {
			return "", err2
		}
		h.Write(s2.Bytes())
	}

	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	token = mac.AccessKey + ":" + sign
	return
}

// -----------------------------------------------------------------------------------------------------------

type Transport struct {
	mac       Mac
	transport http.RoundTripper
}

func incBody(req *http.Request) bool {
	if req.Body == nil {
		return false
	}
	if ct, ok := req.Header["Content-Type"]; ok {
		switch ct[0] {
		case "application/json":
			return true
		}
	}
	return false
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	token, err := t.mac.SignRequest(req, incBody(req))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "pili "+token)
	return t.transport.RoundTrip(req)
}

func NewTransport(mac *Mac, transport http.RoundTripper) *Transport {
	if transport == nil {
		transport = http.DefaultTransport
	}
	t := &Transport{transport: transport}
	t.mac = *mac
	return t
}
