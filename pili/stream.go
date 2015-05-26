package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

func Sign(secret, data []byte) (token string) {
	h := hmac.New(sha1.New, secret)
	h.Write(data)
	token = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}

// Publish URL
// -------------------------------------------------------------------------------
func (s Stream) RtmpPublishUrl(rtmpPublishHost string, nonce int64) (url string) {
	switch s.PublishSecurity {
	case "dynamic":
		if nonce == 0 {
			nonce = time.Now().UnixNano()
		}
		url = s.rtmpPublishDynamicUrl(rtmpPublishHost, nonce)
	case "static":
		url = s.rtmpPublishStaticUrl(rtmpPublishHost)
	}
	return
}

func (s Stream) rtmpPublishDynamicUrl(rtmpPublishHost string, nonce int64) (url string) {
	url = fmt.Sprintf("%s?nonce=%d&token=%s", s.rtmpPublishBaseUrl(rtmpPublishHost), nonce, s.PublishToken(rtmpPublishHost, nonce))
	return
}

func (s Stream) rtmpPublishStaticUrl(rtmpPublishHost string) (url string) {
	url = fmt.Sprintf("%s?key=%s", s.rtmpPublishBaseUrl(rtmpPublishHost), s.PublishKey)
	return
}

func (s Stream) rtmpPublishBaseUrl(rtmpPublishHost string) (url string) {
	url = fmt.Sprintf("rtmp://%s/%s/%s", rtmpPublishHost, s.Hub, s.Title)
	return
}

func (s Stream) PublishToken(rtmpPublishHost string, nonce int64) (token string) {
	u, _ := url.Parse(s.rtmpPublishBaseUrl(rtmpPublishHost))
	uriStr := u.Path
	if u.RawQuery != "" {
		uriStr += "?" + u.RawQuery
	}
	uriStr = fmt.Sprintf("%s?nonce=%d", uriStr, nonce)
	token = Sign([]byte(s.PublishKey), []byte(uriStr))
	return
}

// Play URL
// --------------------------------------------------------------------------------
func (s Stream) RtmpLiveUrl(rtmpPlayHost, profile string) (url string) {
	url = fmt.Sprintf("rtmp://%s/%s/%s", rtmpPlayHost, s.Hub, s.Title)
	if profile != "" {
		url = fmt.Sprintf("%s@%s", url, profile)
	}
	return
}

func (s Stream) HlsLiveUrl(hlsPlayHost, profile string) (url string) {
	url = fmt.Sprintf("http://%s/%s/%s.m3u8", hlsPlayHost, s.Hub, s.Title)
	if profile != "" {
		url = fmt.Sprintf("http://%s/%s/%s@%s.m3u8", hlsPlayHost, s.Hub, s.Title, profile)
	}
	return
}

func (s Stream) HlsPlaybackUrl(hlsPlayHost, profile string, start, end int64) (url string) {
	url = fmt.Sprintf("http://%s/%s/%s.m3u8?start=%d&end=%d", hlsPlayHost, s.Hub, s.Title, start, end)
	if profile != "" {
		url = fmt.Sprintf("http://%s/%s/%s@%s.m3u8?start=%d&end=%d", hlsPlayHost, s.Hub, s.Title, profile, start, end)
	}
	return
}
