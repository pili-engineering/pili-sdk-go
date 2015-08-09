package pili

import (
	"time"
)

type Stream struct {
	conn            RPC_Client
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
		Publish map[string]string `json:"publish,omitempty"`
		Play    map[string]string `json:"play,omitempty"`
	} `json:"hosts,omitempty"`
}

type StreamList struct {
	Marker string    `json:"marker"`
	Items  []*Stream `json:"items"`
}

type StreamSegment struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type StreamSegmentList struct {
	Segments []*StreamSegment `json:"segments"`
}

type StreamStatus struct {
	Addr   string `json:"addr"`
	Status string `json:"status"`
}

type StreamSaveAsResponse struct {
	Url          string `json:"url"`
	TargetUrl    string `json:"targetUrl"`
	PersistentId string `json:"persistentId"`
}

type OptionalArguments struct {
	Title           string
	PublishKey      string
	PublishSecurity string
	Disabled        bool
	Marker          string
	Limit           uint
	Start           int64
	End             int64
	NotifyUrl       string
}
