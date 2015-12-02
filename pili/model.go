package pili

import (
	"time"
)

type Stream struct {
	rpc             *RPC
	Id              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Title           string    `json:"title"`
	Hub             string    `json:"hub"`
	Disabled        bool      `json:"disabled"`
	PublishKey      string    `json:"publishKey"`
	PublishSecurity string    `json:"publishSecurity"`
	Profiles        []string  `json:"profiles,omitempty"`
	Hosts           struct {
		Publish  map[string]string `json:"publish,omitempty"`
		Live     map[string]string `json:"live,omitempty"`
		Playback map[string]string `json:"playback,omitempty"`
	} `json:"hosts,omitempty"`
}

type StreamList struct {
	Marker string    `json:"marker"`
	Items  []*Stream `json:"items"`
	End    bool      `json:"end"`
}

type StreamId struct {
	Id string `json:"id"`
}

type StreamIdList struct {
	Marker string      `json:"marker"`
	Items  []*StreamId `json:"items"`
}

type StreamSegment struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type StreamSegmentList struct {
	Start    int64            `json:"start"`
	End      int64            `json:"end"`
	Duration int64            `json:"duration"`
	Segments []*StreamSegment `json:"segments"`
}

type StreamStatus struct {
	Addr            string  `json:"addr"`
	StartFrom       string  `json:"startFrom"`
	Status          string  `json:"status"`
	BytesPerSecond  float64 `json:"bytesPerSecond"`
	FramesPerSecond struct {
		Audio float64 `json:"audio"`
		Video float64 `json:"video"`
		Data  float64 `json:"data"`
	} `json:"framesPerSecond"`
}

type StreamSaveAsResponse struct {
	Url          string `json:"url"`
	TargetUrl    string `json:"targetUrl"`
	PersistentId string `json:"persistentId"`
}

type StreamSnapshotResponse struct {
	TargetUrl    string `json:"targetUrl"`
	PersistentId string `json:"persistentId"`
}

type OptionalArguments struct {
	Title           string
	PublishKey      string
	PublishSecurity string
	Disabled        bool
	Hub             string
	Idonly          string
	Status          string
	Marker          string
	Limit           uint
	Start           int64
	End             int64
	Time            int64
	NotifyUrl       string
	UserPipeline    string
}
