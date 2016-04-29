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
	url := fmt.Sprintf("%s/streams/%s", getApiBaseUrl(), s.Id)
	err = s.rpc.GetCall(&stream, url)
	if err != nil {
		return
	}
	stream.rpc = s.rpc
	return
}

func (s Stream) ToJSONString() (jsonBlob string, err error) {
	jsonBytes, err := json.Marshal(s)
	jsonBlob = string(jsonBytes)
	return
}

func (s Stream) Enable() (stream Stream, err error) {
	data := map[string]bool{"disabled": false}
	url := fmt.Sprintf("%s/streams/%s", getApiBaseUrl(), s.Id)
	err = s.rpc.PostCall(&stream, url, data)
	stream.rpc = s.rpc
	return
}

func (s Stream) Disable() (stream Stream, err error) {
	data := map[string]bool{"disabled": true}
	url := fmt.Sprintf("%s/streams/%s", getApiBaseUrl(), s.Id)
	err = s.rpc.PostCall(&stream, url, data)
	stream.rpc = s.rpc
	return
}

func (s Stream) Update() (stream Stream, err error) {
	data := map[string]interface{}{}
	if s.PublishKey != "" {
		data["publishKey"] = s.PublishKey
	}
	if s.PublishSecurity != "" {
		data["publishSecurity"] = s.PublishSecurity
	}
	url := fmt.Sprintf("%s/streams/%s", getApiBaseUrl(), s.Id)
	err = s.rpc.PostCall(&stream, url, data)
	stream.rpc = s.rpc
	return
}

func (s Stream) Delete() (ret interface{}, err error) {
	url := fmt.Sprintf("%s/streams/%s", getApiBaseUrl(), s.Id)
	err = s.rpc.DelCall(&ret, url)
	return
}

func (s Stream) Status() (ret StreamStatus, err error) {
	url := fmt.Sprintf("%s/streams/%s/status", getApiBaseUrl(), s.Id)
	err = s.rpc.GetCall(&ret, url)
	return
}

func (s Stream) Segments(args OptionalArguments) (ret StreamSegmentList, err error) {
	url := fmt.Sprintf("%s/streams/%s/segments", getApiBaseUrl(), s.Id)
	if args.Start > 0 {
		url = fmt.Sprintf("%s?start=%d", url, args.Start)
	}
	if args.End > 0 {
		url = fmt.Sprintf("%s&end=%d", url, args.End)
	}
	if args.Limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, args.Limit)
	}
	err = s.rpc.GetCall(&ret, url)
	return
}

func (s Stream) SaveAs(name, format string, start, end int64, args OptionalArguments) (ret StreamSaveAsResponse, err error) {
	data := map[string]interface{}{"name": name, "start": start, "end": end}
	if args.NotifyUrl != "" {
		data["notifyUrl"] = args.NotifyUrl
	}
	if args.UserPipeline != "" {
		data["pipeline"] = args.UserPipeline
	}
	if format != "" {
		data["format"] = format
	}
	url := fmt.Sprintf("%s/streams/%s/saveas", getApiBaseUrl(), s.Id)
	fmt.Println("saveas url:", url, "data:", data)
	err = s.rpc.PostCall(&ret, url, data)
	return
}

func (s Stream) Snapshot(name, format string, args OptionalArguments) (ret StreamSnapshotResponse, err error) {
	data := map[string]interface{}{"name": name, "format": format}
	if args.Time > 0 {
		data["time"] = args.Time
	}
	if args.NotifyUrl != "" {
		data["notifyUrl"] = args.NotifyUrl
	}
	url := fmt.Sprintf("%s/streams/%s/snapshot", getApiBaseUrl(), s.Id)
	err = s.rpc.PostCall(&ret, url, data)
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

// RTMP Live Play URLs
// --------------------------------------------------------------------------------

func (s Stream) RtmpLiveUrls() (urls map[string]string, err error) {
	urls = make(map[string]string)
	url := fmt.Sprintf("rtmp://%s/%s/%s", s.Hosts.Live["rtmp"], s.Hub, s.Title)
	urls[ORIGIN] = url
	return
}

// HLS Live Play URLs
// --------------------------------------------------------------------------------

func (s Stream) HlsLiveUrls() (urls map[string]string, err error) {
	urls = make(map[string]string)
	urls[ORIGIN] = fmt.Sprintf("http://%s/%s/%s.m3u8", s.Hosts.Live["hls"], s.Hub, s.Title)
	return
}

// Http-Flv Live Play URLs
// --------------------------------------------------------------------------------

func (s Stream) HttpFlvLiveUrls() (urls map[string]string, err error) {
	urls = make(map[string]string)
	urls[ORIGIN] = fmt.Sprintf("http://%s/%s/%s.flv", s.Hosts.Live["hdl"], s.Hub, s.Title)
	return
}

// HLS Playback URLs
// --------------------------------------------------------------------------------

func (s Stream) HlsPlaybackUrls(start, end int64) (urls map[string]string, err error) {
	name := fmt.Sprintf("%d", time.Now().Unix())
	ret, err := s.SaveAs(name, "", start, end, OptionalArguments{})
	if err != nil {
		return nil, err
	}

	urls = make(map[string]string)
	urls[ORIGIN] = ret.Url
	return
}
