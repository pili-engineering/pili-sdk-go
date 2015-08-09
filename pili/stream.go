package pili

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

func (s Stream) Refresh() (stream Stream, err error) {
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, s.Id)
	err = s.conn.GetCall(&stream, url)
	if err != nil {
		return
	}
	stream.conn = s.conn
	return
}

func (s Stream) ToJsonString() (jsonBlob string, err error) {
	jsonBytes, err := json.Marshal(s)
	jsonBlob = string(jsonBytes)
	return
}

func (s Stream) Enable() (stream Stream, err error) {
	data := map[string]bool{"disabled": false}
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, s.Id)
	err = s.conn.PostCall(&stream, url, data)
	stream.conn = s.conn
	return
}

func (s Stream) Disable() (stream Stream, err error) {
	data := map[string]bool{"disabled": true}
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, s.Id)
	err = s.conn.PostCall(&stream, url, data)
	stream.conn = s.conn
	return
}

func (s Stream) Update(args OptionalArguments) (stream Stream, err error) {
	data := map[string]interface{}{}
	if args.PublishKey != "" {
		data["publishKey"] = args.PublishKey
	}
	if args.PublishSecurity != "" {
		data["publishSecurity"] = args.PublishSecurity
	}
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, s.Id)
	err = s.conn.PostCall(&stream, url, data)
	stream.conn = s.conn
	return
}

func (s Stream) Delete() (ret interface{}, err error) {
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, s.Id)
	err = s.conn.DelCall(&ret, url)
	return
}

func (s Stream) Status() (ret StreamStatus, err error) {
	url := fmt.Sprintf("%s/streams/%s/status", API_BASE_URL, s.Id)
	err = s.conn.GetCall(&ret, url)
	return
}

func (s Stream) Segments(args OptionalArguments) (ret StreamSegmentList, err error) {
	url := fmt.Sprintf("%s/streams/%s/segments", API_BASE_URL, s.Id)
	if args.Start > 0 {
		url = fmt.Sprintf("%s?start=%d", url, args.Start)
	}
	if args.End > 0 {
		url = fmt.Sprintf("%s&end=%d", url, args.End)
	}
	err = s.conn.GetCall(&ret, url)
	return
}

// Publish URL
// -------------------------------------------------------------------------------
func (s Stream) RtmpPublishUrl() (url string) {
	switch s.PublishSecurity {
	case "dynamic":
		url = s.rtmpPublishDynamicUrl()
	case "static":
		url = s.rtmpPublishStaticUrl()
	}
	return
}

func (s Stream) rtmpPublishDynamicUrl() (url string) {
	nonce := time.Now().UnixNano()
	url = fmt.Sprintf("%s?nonce=%d&token=%s", s.rtmpPublishBaseUrl(), nonce, s.publishToken(nonce))
	return
}

func (s Stream) rtmpPublishStaticUrl() (url string) {
	url = fmt.Sprintf("%s?key=%s", s.rtmpPublishBaseUrl(), s.PublishKey)
	return
}

func (s Stream) rtmpPublishBaseUrl() (url string) {
	url = fmt.Sprintf("rtmp://%s/%s/%s", s.Hosts.Publish["rtmp"], s.Hub, s.Title)
	return
}

func (s Stream) publishToken(nonce int64) (token string) {
	u, _ := url.Parse(s.rtmpPublishBaseUrl())
	uriStr := u.Path
	if u.RawQuery != "" {
		uriStr += "?" + u.RawQuery
	}
	uriStr = fmt.Sprintf("%s?nonce=%d", uriStr, nonce)
	token = s.sign([]byte(s.PublishKey), []byte(uriStr))
	return
}

func (s Stream) sign(secret, data []byte) (token string) {
	h := hmac.New(sha1.New, secret)
	h.Write(data)
	token = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}

// RTMP Play URLs
// --------------------------------------------------------------------------------

func (s Stream) RtmpLiveUrls() (urls map[string]string, err error) {
	urls = make(map[string]string)
	url := fmt.Sprintf("rtmp://%s/%s/%s", s.Hosts.Play["rtmp"], s.Hub, s.Title)
	urls[ORIGIN] = url
	for _, profile := range s.Profiles {
		urls[profile] = fmt.Sprintf("%s@%s", url, profile)
	}
	return
}

// HLS Play URLs
// --------------------------------------------------------------------------------

func (s Stream) HlsLiveUrls() (urls map[string]string, err error) {
	urls = make(map[string]string)
	urls[ORIGIN] = fmt.Sprintf("http://%s/%s/%s.m3u8", s.Hosts.Play["hls"], s.Hub, s.Title)
	for _, profile := range s.Profiles {
		urls[profile] = fmt.Sprintf("http://%s/%s/%s@%s.m3u8", s.Hosts.Play["hls"], s.Hub, s.Title, profile)
	}
	return
}

// HLS Playback URLs
// --------------------------------------------------------------------------------

func (s Stream) HlsPlaybackUrls(start, end int64) (urls map[string]string, err error) {
	urls = make(map[string]string)
	urls[ORIGIN] = fmt.Sprintf("http://%s/%s/%s.m3u8?start=%d&end=%d", s.Hosts.Play["hls"], s.Hub, s.Title, start, end)
	for _, profile := range s.Profiles {
		urls[profile] = fmt.Sprintf("http://%s/%s/%s@%s.m3u8?start=%d&end=%d", s.Hosts.Play["hls"], s.Hub, s.Title, profile, start, end)
	}
	return
}
