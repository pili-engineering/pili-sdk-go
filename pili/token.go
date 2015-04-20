package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
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

type PushPolicy struct {
	BaseUrl string
	Key     string
	Nonce   int64
}

func (p PushPolicy) Sign() (token, url string) {
	nonce := p.Nonce
	url = p.BaseUrl + "?nonce=" + strconv.FormatInt(nonce, 10)
	token = Sign([]byte(p.Key), []byte(url))
	url = url + "&token=" + token
	return
}

func (p PushPolicy) Token() (token string) {
	token, _ = p.Sign()
	return
}

func (p PushPolicy) Url() (url string) {
	_, url = p.Sign()
	return
}

// ----------------------------------------------------------
/*
type PlayPolicy struct {
	BaseUrl string
	Key     string
	Expiry  int64
}

func (p PlayPolicy) Sign() (token, url string) {
	expiry := p.Expiry
	if expiry == 0 {
		expiry = time.Now().Unix() + 3600
	}
	url = p.BaseUrl + "?expiry=" + strconv.FormatInt(expiry, 10)
	token = Sign([]byte(p.Key), []byte(url))
	url = url + "&token=" + token
	return
}

func (p PlayPolicy) Token() (token string) {
	token, _ = p.Sign()
	return
}

func (p PlayPolicy) Url() (url string) {
	_, url = p.Sign()
	return

}
*/
