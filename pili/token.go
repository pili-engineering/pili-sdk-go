package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strconv"
	//"time"
)

func Sign(secret, data []byte) (token string) {
	h := hmac.New(sha1.New, secret)
	h.Write(data)
	token = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}

// ----------------------------------------------------------

type PublishPolicy struct {
	BaseUrl    string
	PublishKey string
	Nonce      int64
}

func (p PublishPolicy) Sign() (publishToken, publishUrl string, err error) {
	u, err := url.Parse(p.BaseUrl)
	if err != nil {
		return
	}
	uriStr := u.Path
	if u.RawQuery != "" {
		uriStr += "?" + u.RawQuery
	}
	nonce := strconv.FormatInt(p.Nonce, 10)
	uriStr = uriStr + "?nonce=" + nonce
	publishToken = Sign([]byte(p.PublishKey), []byte(uriStr))
	publishUrl = p.BaseUrl + "?nonce=" + nonce + "&token=" + publishToken
	return
}

func (p PublishPolicy) Token() (publishToken string) {
	publishToken, _, _ = p.Sign()
	return
}

func (p PublishPolicy) Url() (publishUrl string) {
	_, publishUrl, _ = p.Sign()
	return
}
