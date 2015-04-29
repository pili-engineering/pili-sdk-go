package pili

import (
	"fmt"
	"net/http"
	"time"
)

type API_Client struct {
	Conn RPC_Client
}

func NewClient(mac *Mac) API_Client {
	t := NewTransport(mac, nil)
	client := &http.Client{Transport: t}
	return API_Client{RPC_Client{client}}
}

// -----------------------------------------------------------------------------------------------------------

var API_BASE_URL = fmt.Sprintf("%s://%s/%s", getHttpScheme(), getApiHost(), API_VERSION)

func URI_NewStream() string {
	return fmt.Sprintf("%s/streams", API_BASE_URL)
}

func URI_GetStream(id string) string {
	return fmt.Sprintf("%s/streams/%s", API_BASE_URL, id)
}

func URI_SetStream(id string) string {
	return fmt.Sprintf("%s/streams/%s", API_BASE_URL, id)
}

func URI_DelStream(id string) string {
	return fmt.Sprintf("%s/streams/%s", API_BASE_URL, id)
}

func URI_ListStreams(hub string, options map[string]interface{}) (httpurl string) {
	httpurl = fmt.Sprintf("%s/streams?hub=%s", API_BASE_URL, hub)
	if marker, ok := options["marker"]; ok && marker != "" {
		httpurl = fmt.Sprintf("%s&marker=%s", httpurl, marker)
	}
	if limit, ok := options["limit"]; ok && limit != 0 {
		httpurl = fmt.Sprintf("%s&limit=%d", httpurl, limit)
	}
	return
}

func URI_GetStreamSegments(id string, options map[string]interface{}) (httpurl string) {
	httpurl = fmt.Sprintf("%s/streams/%s/segments", API_BASE_URL, id)
	if start, ok := options["start"]; ok && start != 0 {
		httpurl = fmt.Sprintf("%s?start=%d", httpurl, start)
	}
	if end, ok := options["end"]; ok && end != 0 {
		httpurl = fmt.Sprintf("%s&end=%d", httpurl, end)
	}
	return
}

// -----------------------------------------------------------------------------------------------------------

type Stream struct {
	Id              string    `json:"id"`
	Hub             string    `json:"hub"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Title           string    `json:"title"`
	PublishKey      string    `json:"publishKey"`
	PublishSecurity string    `json:"publishSecurity"`
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

// -----------------------------------------------------------------------------------------------------------

func (app API_Client) CreateStream(hub, title, publishKey, publishSecurity string) (ret Stream, err error) {
	data := map[string]interface{}{"hub": hub}
	if title != "" {
		data["title"] = title
	}
	if publishKey != "" {
		data["publishKey"] = publishKey
	}
	if publishSecurity != "" {
		data["publishSecurity"] = publishSecurity
	}
	err = app.Conn.PostCall(&ret, URI_NewStream(), data)
	return
}

func (app API_Client) GetStream(id string) (ret Stream, err error) {
	err = app.Conn.GetCall(&ret, URI_GetStream(id))
	return
}

func (app API_Client) SetStream(id, publishKey, publishSecurity string) (ret Stream, err error) {
	data := map[string]interface{}{
		"publishKey":      publishKey,
		"publishSecurity": publishSecurity,
	}
	err = app.Conn.PostCall(&ret, URI_SetStream(id), data)
	return
}

func (app API_Client) DelStream(id string) (ret interface{}, err error) {
	err = app.Conn.DelCall(&ret, URI_DelStream(id))
	return
}

func (app API_Client) ListStreams(hub, marker string, limit int64) (ret StreamList, err error) {
	options := map[string]interface{}{
		"marker": marker,
		"limit":  limit,
	}
	err = app.Conn.GetCall(&ret, URI_ListStreams(hub, options))
	return
}

func (app API_Client) GetStreamSegments(id string, startTime, endTime int64) (ret StreamSegmentList, err error) {
	options := map[string]interface{}{}
	if startTime != 0 {
		options["start"] = startTime
	}
	if endTime != 0 {
		options["end"] = endTime
	}
	err = app.Conn.GetCall(&ret, URI_GetStreamSegments(id, options))
	return
}
