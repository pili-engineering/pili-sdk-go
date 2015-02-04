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

var API_BASE_URL = "http://api.pili.qiniu.com/v1"

func URI_NewStream() string {
	return fmt.Sprintf("%s/streams", API_BASE_URL)
}

func URI_ListStreams() string {
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

func URI_GetStreamStatus(id string) string {
	return fmt.Sprintf("%s/streams/%s/status", API_BASE_URL, id)
}

func URI_GetStreamSegments(id string, starttime, endtime int64) string {
	return fmt.Sprintf("%s/streams/%s/segments?starttime=%d&endtime=%d", API_BASE_URL, id, starttime, endtime)
}

func URI_DelStreamSegments(id string, starttime, endtime int64) string {
	return fmt.Sprintf("%s/streams/%s/segments?starttime=%d&endtime=%d", API_BASE_URL, id, starttime, endtime)
}

func URI_PlayStreamSegments(id string, starttime, endtime int64) string {
	return fmt.Sprintf("%s/streams/%s/segments/play?starttime=%d&endtime=%d", API_BASE_URL, id, starttime, endtime)
}

// -----------------------------------------------------------------------------------------------------------

type StreamUrl struct {
	RTMP string `json:"RTMP"`
	HLS  string `json:"HLS"`
}

type LivePlayUrl struct {
	*StreamUrl `json:"[original]"`
}

type Stream struct {
	Id          string       `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Application string       `json:"application"`
	IsPrivate   bool         `json:"is_private"`
	Key         string       `json:"key"`
	Comment     string       `json:"comment"`
	PushUrl     []*StreamUrl `json:"push_url"`
	LiveUrl     *LivePlayUrl `json:"live_url"`
}

type StreamStatus struct {
	Status string `json:"status"`
}

type StreamList struct {
	Total   uint32    `json:"total"`
	Streams []*Stream `json:"streams"`
}

type StreamSegment struct {
	StartTime int64 `json:"starttime"`
	EndTime   int64 `json:"endtime"`
}

type StreamSegmentList struct {
	Total uint32           `json:"total"`
	List  []*StreamSegment `json:"list"`
}

type StreamSegmentPlay struct {
	Url string `json:"url"`
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

func (app API_Client) DelStream(id string) (ret Stream, err error) {
	err = app.Conn.DelCall(&ret, URI_DelStream(id))
	return
}

func (app API_Client) GetStreamStatus(id string) (ret StreamStatus, err error) {
	err = app.Conn.GetCall(&ret, URI_GetStreamStatus(id))
	return
}

func (app API_Client) ListStreams() (ret StreamList, err error) {
	err = app.Conn.GetCall(&ret, URI_ListStreams())
	return
}

func (app API_Client) GetStreamSegments(id string, starttime, endtime int64) (ret StreamSegmentList, err error) {
	err = app.Conn.GetCall(&ret, URI_GetStreamSegments(id, starttime, endtime))
	return
}

func (app API_Client) PlayStreamSegments(id string, starttime, endtime int64) (ret StreamSegmentPlay, err error) {
	err = app.Conn.GetCall(&ret, URI_PlayStreamSegments(id, starttime, endtime))
	return
}

func (app API_Client) DelStreamSegments(id string, starttime, endtime int64) (ret interface{}, err error) {
	err = app.Conn.DelCall(&ret, URI_DelStreamSegments(id, starttime, endtime))
	return
}
