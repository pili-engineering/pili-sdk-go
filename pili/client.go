package pili

import (
	"fmt"
	"net/http"
)

type Client struct {
	conn RPC_Client
	hub  string
}

func NewClient(mac *Mac, hub string) Client {
	t := NewTransport(mac, nil)
	tc := &http.Client{Transport: t}
	return Client{conn: RPC_Client{tc}, hub: hub}
}

func (c Client) CreateStream(args OptionalArguments) (stream Stream, err error) {
	data := map[string]interface{}{"hub": c.hub}

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
	stream.conn = c.conn

	return
}

func (c Client) GetStream(id string) (stream Stream, err error) {
	url := fmt.Sprintf("%s/streams/%s", API_BASE_URL, id)
	err = c.conn.GetCall(&stream, url)
	stream.conn = c.conn
	return
}

func (c Client) ListStreams(args OptionalArguments) (ret StreamList, err error) {

	url := fmt.Sprintf("%s/streams?hub=%s", API_BASE_URL, c.hub)

	if args.Marker != "" {
		url = fmt.Sprintf("%s&marker=%s", url, args.Marker)
	}

	if args.Limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, args.Limit)
	}

	err = c.conn.GetCall(&ret, url)

	streams := make([]*Stream, len(ret.Items))

	for _, stream := range ret.Items {
		stream.conn = c.conn
		streams = append(streams, stream)
	}

	ret.Items = streams

	return
}
