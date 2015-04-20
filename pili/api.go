package pili

import (
	"fmt"
	"net/http"
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

var API_BASE_URL = "http://pili.qiniuapi.com/v1"

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

func URI_ListStreams(hub string, marker, limit int64) string {
	return fmt.Sprintf("%s/streams?hub=%s&marker=%d&limit=%d", API_BASE_URL, hub, marker, limit)
}

func URI_GetStreamSegments(id string, start, end int64) string {
	return fmt.Sprintf("%s/streams/%s/segments?start=%d&end=%d", API_BASE_URL, id, start, end)
}

// -----------------------------------------------------------------------------------------------------------

type Stream struct {
	Id              string `json:"id"`
	Hub             string `json:"hub"`
	Title           string `json:"title"`
	PublishKey      string `json:"publishKey"`
	PublishSecurity string `json:"publishSecurity"`
}

type StreamList struct {
	Marker uint32    `json:"marker"`
	Items  []*Stream `json:"items"`
}

type StreamSegment struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type StreamSegmentList struct {
	segments []*StreamSegment `json:"segments"`
}

// -----------------------------------------------------------------------------------------------------------

func (app API_Client) CreateStream(data interface{}) (ret Stream, err error) {
	err = app.Conn.PostCall(&ret, URI_NewStream(), data)
	return
}

func (app API_Client) GetStream(id string) (ret Stream, err error) {
	err = app.Conn.GetCall(&ret, URI_GetStream(id))
	return
}

func (app API_Client) SetStream(id string, data interface{}) (ret Stream, err error) {
	err = app.Conn.PostCall(&ret, URI_SetStream(id), data)
	return
}

func (app API_Client) DelStream(id string) (ret interface{}, err error) {
	err = app.Conn.DelCall(&ret, URI_DelStream(id))
	return
}

func (app API_Client) ListStreams(hub string, marker, limit int64) (ret StreamList, err error) {
	err = app.Conn.GetCall(&ret, URI_ListStreams(hub, marker, limit))
	return
}

func (app API_Client) GetStreamSegments(id string, start, end int64) (ret StreamSegmentList, err error) {
	err = app.Conn.GetCall(&ret, URI_GetStreamSegments(id, start, end))
	return
}
