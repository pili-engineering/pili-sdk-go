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

func Creds(accessKey, secretKey string) *Mac {
	return &Mac{accessKey, secretKey}
}

func incBody(req *http.Request, ctType string) bool {

	return req.Body != nil && ctType != "" && ctType != "application/octet-stream"
}

func (mac *Mac) SignRequest(req *http.Request) (token string, err error) {

	h := hmac.New(sha1.New, []byte(mac.SecretKey))

	u := req.URL
	data := req.Method + " " + u.Path
	if u.RawQuery != "" {
		data += "?" + u.RawQuery
	}
	io.WriteString(h, data+"\nHost: "+req.Host)

	ctType := req.Header.Get("Content-Type")
	if ctType != "" {
		io.WriteString(h, "\nContent-Type: "+ctType)
	}
	io.WriteString(h, "\n\n")

	if incBody(req, ctType) {
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

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	token, err := t.mac.SignRequest(req)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Qiniu "+token)
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
