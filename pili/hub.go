package pili

import (
	"fmt"
)

type Hub struct {
	rpc     *RPC
	hubName string
}

func NewHub(creds *Credentials, hubName string) Hub {
	return Hub{rpc: NewRPC(creds), hubName: hubName}
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
	url := fmt.Sprintf("%s/streams", getApiBaseUrl())
	err = c.rpc.PostCall(&stream, url, data)
	if err != nil {
		return
	}
	stream.rpc = c.rpc
	return
}

func (c Hub) GetStream(id string) (stream Stream, err error) {
	url := fmt.Sprintf("%s/streams/%s", getApiBaseUrl(), id)
	err = c.rpc.GetCall(&stream, url)
	if err != nil {
		return
	}
	stream.rpc = c.rpc
	return
}

func (c Hub) ListStreams(args OptionalArguments) (ret StreamList, err error) {
	url := fmt.Sprintf("%s/streams?hub=%s", getApiBaseUrl(), c.hubName)
	if args.Status != "" {
		url = fmt.Sprintf("%s&status=%s", url, args.Status)
	}
	if args.Marker != "" {
		url = fmt.Sprintf("%s&marker=%s", url, args.Marker)
	}
	if args.Limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, args.Limit)
	}
	if args.Title != "" {
		url = fmt.Sprintf("%s&title=%s", url, args.Title)
	}
	resultWrapper := StreamList{}
	err = c.rpc.GetCall(&resultWrapper, url)
	if err != nil {
		return
	}
	count := len(resultWrapper.Items)
	streams := make([]*Stream, count)
	for i := 0; i < count; i++ {
		streams[i] = resultWrapper.Items[i]
		streams[i].rpc = c.rpc
	}
	ret.Items = streams
	ret.Marker = resultWrapper.Marker
	ret.End = resultWrapper.End
	return
}
