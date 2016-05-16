package pili

import "fmt"

// Hub 表示一个 Hub 对象.
type Hub struct {
	hub     string
	baseURL string
	client  *Client
}

func newHub(hub string, client *Client) *Hub {
	baseURL := APIHTTPScheme + APIHost + "/v2/hubs/" + hub
	return &Hub{hub, baseURL, client}
}

type createArgs struct {
	Key string `json:"key"`
}

// Create 创建一个流对象.
// 使用一个合法 RTMPPublishURL 发起推流就会自动创建流对象.
// 一般情况下不需要调用这个 API, 除非是想提前对这一个流做一些特殊配置.
func (h *Hub) Create(streamKey string) (stream *Stream, err error) {
	args := &createArgs{streamKey}
	path := h.baseURL + "/streams"
	err = h.client.CallWithJSON(nil, "POST", path, args)
	if err != nil {
		return
	}
	stream = newStream(h.hub, streamKey, h.client)
	return
}

// Stream 初始化一个流对象
func (h *Hub) Stream(key string) *Stream {
	return newStream(h.hub, key, h.client)
}

// ---------------------------------------------------------------------------------------

type listItem struct {
	Key string `json:"key"`
}

func (h *Hub) list(live bool, prefix string, limit int, marker string) (keys []string, omarker string, err error) {

	path := fmt.Sprintf("%s/streams?liveonly=%v&prefix=%s&limit=%d&marker=%s", h.baseURL, live, prefix, limit, marker)
	var ret struct {
		Items  []listItem `json:"items"`
		Marker string     `json:"marker"`
	}
	err = h.client.Call(&ret, "GET", path)
	if err != nil {
		return
	}
	keys = make([]string, len(ret.Items))
	for i, item := range ret.Items {
		keys[i] = item.Key
	}
	omarker = ret.Marker
	return
}

// List 根据 prefix 遍历 Hub 的流列表.
// limit 限定了一次最多可以返回的流个数, 实际返回的流个数可能小于这个 limit 值.
// marker 是上一次遍历得到的流标.
// omarker 记录了此次遍历到的游标, 在下次请求时应该带上, 如果 omarker 为 "" 表示已经遍历完所有流.
//
// For example:
//
//     var keys []string
//     var marker string
//     var err error
//     for {
//         keys, marker, err = hub.List("", 0, marker)
//         if err != nil {
//             break
//         }
//
//         // Do something with keys.
//         ...
//
//         if marker == "" {
//             break
//         }
//     }
func (h *Hub) List(prefix string, limit int, marker string) (keys []string, omarker string, err error) {
	return h.list(false, prefix, limit, marker)
}

// ListLive 根据 prefix 遍历 Hub 的流列表.
// limit 限定了一次最多可以返回的流个数, 实际返回的流个数可能小于这个 limit 值.
// marker 是上一次遍历得到的流标.
// omarker 记录了此次遍历到的游标, 在下次请求时应该带上, 如果 omarker 为 "" 表示已经遍历完所有流.
//
// For example:
//
//     var keys []string
//     var marker string
//     var err error
//     for {
//         keys, marker, err = hub.ListLive("", 0, marker)
//         if err != nil {
//             break
//         }
//
//         // Do something with keys.
//         ...
//
//         if marker == "" {
//             break
//         }
//     }
func (h *Hub) ListLive(prefix string, limit int, marker string) (keys []string, omarker string, err error) {
	return h.list(true, prefix, limit, marker)
}

// -----------------------------------------------------------------------------

// IsExists 判断一个 error 是否表示资源存在.
func IsExists(err error) bool {
	e, ok := err.(*Error)
	return ok && e.Code == 614
}

// IsNotExists 判断一个 error 是否表示资源不存在.
func IsNotExists(err error) bool {
	e, ok := err.(*Error)
	return ok && e.Code == 612
}
