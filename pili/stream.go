package pili

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

type streamInfo struct {
	Hub string
	Key string

	// 禁用结束的时间, 0 表示不禁用, -1 表示永久禁用.
	DisabledTill int64
}

// Stream 表示一个流对象.
type Stream struct {
	streamInfo
	baseURL string
	client  *Client
}

func newStream(info *streamInfo, client *Client) *Stream {
	ekey := base64.URLEncoding.EncodeToString([]byte(info.Key))
	baseURL := fmt.Sprintf("%s%s/v2/hubs/%v/streams/%v", APIHTTPScheme, APIHost, info.Hub, ekey)
	return &Stream{*info, baseURL, client}
}

type disabledArgs struct {
	DisabledTill int64 `json:"disabledTill"`
}

// Disable 禁用一个流.
func (s *Stream) Disable() error {
	args := &disabledArgs{-1}
	path := s.baseURL + "/disabled"
	return s.client.CallWithJSON(nil, "POST", path, args)
}

// Enable 启用一个流.
func (s *Stream) Enable() error {
	args := &disabledArgs{0}
	path := s.baseURL + "/disabled"
	return s.client.CallWithJSON(nil, "POST", path, args)
}

// FPSStatus 帧率状态
type FPSStatus struct {
	Audio int `json:"audio"`
	Video int `json:"video"`
	Data  int `json:"data"`
}

// LiveStatus 直播状态
type LiveStatus struct {
	// 直播开始的 Unix 时间戳, 0 表示当前没在直播.
	StartAt int64 `json:"startAt"`

	// 直播的客户端 IP.
	ClientIP string `json:"clientIP"`

	// 直播的码率、帧率信息.
	BPS int       `json:"bps"`
	FPS FPSStatus `json:"fps"`
}

// LiveStatus 查询直播状态.
func (s *Stream) LiveStatus() (status *LiveStatus, err error) {
	path := s.baseURL + "/live"
	err = s.client.Call(&status, "GET", path)
	return
}

func appendQuery(path string, start, end int64) string {
	flag := "?"
	if start > 0 {
		path += flag + "start=" + strconv.FormatInt(start, 10)
		flag = "&"
	}
	if end > 0 {
		path += flag + "end=" + strconv.FormatInt(end, 10)
	}
	return path
}

type saveArgs struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// Save 保存直播回放.
// start, end 是 Unix 时间戳, 限定了保存的直播的时间范围, 0 值表示不限定, 系统会默认保存最近一次直播的内容.
// fname 保存的文件名, 由系统生成.
func (s *Stream) Save(start, end int64) (fname string, err error) {
	path := appendQuery(s.baseURL+"/saveas", start, end)
	args := &saveArgs{start, end}
	var ret struct {
		Fname string `json:"fname"`
	}
	err = s.client.CallWithJSON(&ret, "POST", path, args)
	if err != nil {
		return
	}
	fname = ret.Fname
	return
}

// Record 表示一次直播记录
type Record struct {
	Start int64 `json:"start"` // 直播开始时间
	End   int64 `json:"end"`   // 直播结束时间
}

// History 查询直播历史.
// start, end 是 Unix 时间戳, 限定了查询的时间范围, 0 值表示不限定, 系统会返回所有时间的直播历史.
func (s *Stream) HistoryRecord(start, end int64) (records []Record, err error) {
	path := appendQuery(s.baseURL+"/historyrecord", start, end)
	var ret struct {
		Items []Record `json:"items"`
	}
	err = s.client.Call(&ret, "GET", path)
	if err != nil {
		return
	}
	records = ret.Items
	return
}
