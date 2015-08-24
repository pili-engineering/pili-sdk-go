package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/qiniu/bytes/seekable"
	"io"
	"net/http"
)

func NewCredentials(accessKey, secretKey string) *Credentials {
	return &Credentials{accessKey, secretKey}
}

type Credentials struct {
	AccessKey string
	SecretKey string
}

func (c *Credentials) MACToken(req *http.Request) (token string, err error) {

	h := hmac.New(sha1.New, []byte(c.SecretKey))

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

	if req.Body != nil && ctType != "" && ctType != "application/octet-stream" {
		s2, err2 := seekable.New(req)
		if err2 != nil {
			return "", err2
		}
		h.Write(s2.Bytes())
	}

	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	token = c.AccessKey + ":" + sign
	return
}
