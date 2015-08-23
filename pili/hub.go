package pili

import (
	"fmt"
	"net/http"
)

type Hub struct {
	conn    RPC_Client
	hubName string
}

func NewHub(creds *Credentials, hubName string) Hub {
	t := NewTransport(creds, nil)
	tc := &http.Client{Transport: t}
	return Hub{conn: RPC_Client{tc}, hubName: hubName}
}

func (c Hub) CreateStream(args OptionalArguments) (stream Stream, err error) {
	data := map[string]interface{}{"hub": c.hubName}
	if args.Title != "" {
		data["title"] = args.Title
	}
	if args.PublishKey != "" {
		data["publishKey"] = args.PublishKey
	}
	if args.PublishSecurity != "" {
		data["publishSecurity"] = args.PublishSecurity
	}
	url := fmt.Sprintf("%s/streams", API_BASE_URL)
	err = c.conn.PostCall(&stream, url, data)
	if err != nil {
		return
	}
	stream.conn = c.conn
	return
}

func (c Hub) GetStream(id string) (stream Stream, err error) {
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, id)
	err = c.conn.GetCall(&stream, url)
	if err != nil {
		return
	}
	stream.conn = c.conn
	return
}

func (c Hub) ListStreams(args OptionalArguments) (ret StreamList, err error) {
	url := fmt.Sprintf("%s/streams?hub=%s", API_BASE_URL, c.hubName)
	if args.Marker != "" {
		url = fmt.Sprintf("%s&marker=%s", url, args.Marker)
	}
	if args.Limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, args.Limit)
	}
	resultWrapper := StreamList{}
	err = c.conn.GetCall(&resultWrapper, url)
	if err != nil {
		return
	}
	count := len(resultWrapper.Items)
	streams := make([]*Stream, count)
	for i := 0; i < count; i++ {
		streams[i] = resultWrapper.Items[i]
		streams[i].conn = c.conn
	}
	ret.Items = streams
	ret.Marker = resultWrapper.Marker
	return
}
