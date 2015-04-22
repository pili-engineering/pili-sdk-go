package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const STREAM_ID_SEPARATOR = "."

func resolveStreamId(sid string) (hub, title string) {
	a := strings.Split(sid, STREAM_ID_SEPARATOR)
	hub = a[1]
	title = a[2]
	return
}

func Sign(secret, data []byte) (token string) {
	h := hmac.New(sha1.New, secret)
	h.Write(data)
	token = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}

// ----------------------------------------------------------

type PublishPolicy struct {
	StreamId              string // required, format: <zone>.<hub>.<title>
	StreamPublishKey      string // required, a secret key for signing the <publishToken>
	StreamPublishSecurity string // required, can be "dynamic" or "static"
	Nonce                 int64  // optional, for "dynamic" only, default is time.Now().UnixNano()
}

func (p PublishPolicy) baseUrl() (url string) {
	hub, title := resolveStreamId(p.StreamId)
	url = fmt.Sprintf("rtmp://%s/%s/%s", RTMP_PUBLISH_HOST, hub, title)
	return
}

func (p PublishPolicy) staticUrl() (url string) {
	hub, title := resolveStreamId(p.StreamId)
	url = fmt.Sprintf("rtmp://%s/%s/%s?key=%s", RTMP_PUBLISH_HOST, hub, title, p.StreamPublishKey)
	return
}

func (p PublishPolicy) dynamicUrl() (url string) {
	url = fmt.Sprintf("%s?nonce=%d&token=%s", p.baseUrl(), p.NonceVal(), p.Token())
	return
}

func (p PublishPolicy) NonceVal() (nonce int64) {
	nonce = p.Nonce
	if nonce == 0 {
		nonce = time.Now().UnixNano()
	}
	return
}

func (p PublishPolicy) Token() (token string) {
	u, _ := url.Parse(p.baseUrl())
	uriStr := u.Path
	if u.RawQuery != "" {
		uriStr += "?" + u.RawQuery
	}
	uriStr = fmt.Sprintf("%s?nonce=%d", uriStr, p.NonceVal())
	token = Sign([]byte(p.StreamPublishKey), []byte(uriStr))
	return
}

func (p PublishPolicy) Url() (url string) {
	switch p.StreamPublishSecurity {
	case "dynamic":
		url = p.dynamicUrl()
	case "static":
		url = p.staticUrl()
	}
	return
}

// ----------------------------------------------------------

type PlayPolicy struct {
	StreamId string // required, format: <zone>.<hub>.<title>
}

func (p PlayPolicy) RtmpLiveUrl(preset string) (url string) {
	hub, title := resolveStreamId(p.StreamId)
	url = fmt.Sprintf("rtmp://%s/%s/%s", RTMP_PUBLISH_HOST, hub, title)
	if preset != "" {
		url = fmt.Sprintf("rtmp://%s/%s/%s@%s", RTMP_PUBLISH_HOST, hub, title, preset)
	}
	return
}

func (p PlayPolicy) HlsLiveUrl(preset string) (url string) {
	hub, title := resolveStreamId(p.StreamId)
	url = fmt.Sprintf("http://%s/%s/%s.m3u8", HLS_PLAY_HOST, hub, title)
	if preset != "" {
		url = fmt.Sprintf("http://%s/%s/%s@%s.m3u8", HLS_PLAY_HOST, hub, title, preset)
	}
	return
}

func (p PlayPolicy) HlsPlaybackUrl(start, end int64, preset string) (url string) {
	hub, title := resolveStreamId(p.StreamId)
	url = fmt.Sprintf("http://%s/%s/%s.m3u8?start=%d&end=%d", HLS_PLAY_HOST, hub, title, start, end)
	if preset != "" {
		url = fmt.Sprintf("http://%s/%s/%s@%s.m3u8?start=%d&end=%d", HLS_PLAY_HOST, hub, title, preset, start, end)
	}
	return
}
