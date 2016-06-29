package pili

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

// StreamInfo 流信息.
type StreamInfo struct {
	Hub string
	Key string

	// 禁用结束的时间, 0 表示不禁用, -1 表示永久禁用.
	disabledTill int64
}

// Disabled 判断一个流是否被禁用.
func (s *StreamInfo) Disabled() bool {
	return s.disabledTill == -1 || s.disabledTill > time.Now().Unix()
}

// String 返回格式化后的流信息
func (s *StreamInfo) String() string {
	return fmt.Sprintf("{hub:%s,key:%s,disabled:%v}", s.Hub, s.Key, s.Disabled())
}

// Stream 表示一个流对象.
type Stream struct {
	hub     string
	key     string
	baseURL string
	client  *Client
}

func newStream(hub, key string, client *Client) *Stream {
	ekey := base64.URLEncoding.EncodeToString([]byte(key))
	baseURL := fmt.Sprintf("%s%s/v2/hubs/%v/streams/%v", APIHTTPScheme, APIHost, hub, ekey)
	return &Stream{hub, key, baseURL, client}
}

// Info 获得流信息.
func (s *Stream) Info() (info *StreamInfo, err error) {
	path := s.baseURL
	var ret struct {
		DisabledTill int64 `json:"disabledTill"`
	}
	err = s.client.Call(&ret, "GET", path)
	if err != nil {
		return
	}
	info = &StreamInfo{s.hub, s.key, ret.DisabledTill}
	return
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
// status.StartAt 记录了直播开始的时间, 0 表示当前没在直播.
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

// ActivityRecord 表示一次直播记录
type ActivityRecord struct {
	Start int64 `json:"start"` // 直播开始时间
	End   int64 `json:"end"`   // 直播结束时间
}

// History 查询直播历史.
// start, end 是 Unix 时间戳, 限定了查询的时间范围, 0 值表示不限定, 系统会返回所有时间的直播历史.
func (s *Stream) HistoryActivity(start, end int64) (records []ActivityRecord, err error) {
	path := appendQuery(s.baseURL+"/historyactivity", start, end)
	var ret struct {
		Items []ActivityRecord `json:"items"`
	}
	err = s.client.Call(&ret, "GET", path)
	if err != nil {
		return
	}
	records = ret.Items
	return
}
