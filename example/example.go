package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pili-engineering/pili-sdk-go/pili"
)

var (
	AccessKey = "<QINIU ACCESS KEY>" // 替换成自己 Qiniu 账号的 AccessKey.
	SecretKey = "<QINIU SECRET KEY>" // 替换成自己 Qiniu 账号的 SecretKey.
	HubName   = "<PILI HUB NAME>"    // Hub 必须事先存在.
)

func init() {
	AccessKey = os.Getenv("PILI_ACCESS_KEY")
	SecretKey = os.Getenv("PILI_SECRET_KEY")
}

func createStream(hub *pili.Hub, key string) {
	stream, err := hub.Create(key)
	if err != nil {
		return
	}
	info, err := stream.Info()
	if err != nil {
		return
	}
	fmt.Println(info)
}

func getStream(hub *pili.Hub, key string) {
	stream := hub.Stream(key)
	info, err := stream.Info()
	if err != nil {
		return
	}
	fmt.Println(info)
}

func listStreams(hub *pili.Hub, prefix string) {
	keys, marker, err := hub.List(prefix, 10, "")
	if err != nil {
		return
	}
	fmt.Printf("keys=%v marker=%v\n", keys, marker)
}

func listLiveStreams(hub *pili.Hub, prefix string) {
	keys, marker, err := hub.ListLive(prefix, 10, "")
	if err != nil {
		return
	}
	fmt.Printf("keys=%v marker=%v\n", keys, marker)
}

func disableStream(hub *pili.Hub, key string) {
	stream := hub.Stream(key)
	info, err := stream.Info()
	if err != nil {
		return
	}
	fmt.Println("before disable:", info)

	err = stream.Disable()
	if err != nil {
		return
	}

	info, err = stream.Info()
	if err != nil {
		return
	}
	fmt.Println("after disable:", info)
}

func enableStream(hub *pili.Hub, key string) {
	stream := hub.Stream(key)
	info, err := stream.Info()
	if err != nil {
		return
	}
	fmt.Println("before enable:", info)

	err = stream.Enable()
	if err != nil {
		return
	}

	info, err = stream.Info()
	if err != nil {
		return
	}
	fmt.Println("after enable:", info)
}

func liveStatus(hub *pili.Hub, key string) {
	stream := hub.Stream(key)
	status, err := stream.LiveStatus()
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", status)
}

func historyActivity(hub *pili.Hub, key string) {
	stream := hub.Stream(key)
	records, err := stream.HistoryActivity(0, 0)
	if err != nil {
		return
	}
	fmt.Println(records)
}

func savePlayback(hub *pili.Hub, key string) {
	stream := hub.Stream(key)
	fname, err := stream.Save(0, 0)
	if err != nil {
		return
	}
	fmt.Println(fname)
}

func main() {
	// 检查 AccessKey、SecretKey 的配置情况.
	if AccessKey == "" || SecretKey == "" {
		log.Printf("WARN: AccessKey=%s SecretKey=%s\n", AccessKey, SecretKey)
		return
	}
	streamKeyPrefix := "sdkexample" + strconv.FormatInt(time.Now().UnixNano(), 10)

	//HubName = "PiliSDKTest"
	//pili.APIHost = "10.200.20.28:7778"

	// 初始化 client & hub.
	mac := &pili.MAC{AccessKey: AccessKey, SecretKey: []byte(SecretKey)}
	client := pili.New(mac, nil)
	hub := client.Hub(HubName)

	keyA := streamKeyPrefix + "A"
	fmt.Println("获得不存在的流A:")
	streamA := hub.Stream(keyA)
	_, err := streamA.Info()
	fmt.Println(err, "IsNotExists", pili.IsNotExists(err))

	fmt.Println("创建流:")
	createStream(hub, keyA)

	fmt.Println("获得流:")
	getStream(hub, keyA)

	fmt.Println("创建重复流:")
	_, err = hub.Create(keyA)
	fmt.Println(err, "IsExists", pili.IsExists(err))

	keyB := streamKeyPrefix + "B"
	fmt.Println("创建另一路流:")
	createStream(hub, keyB)

	fmt.Println("列出流:")
	listStreams(hub, "carter")

	fmt.Println("列出正在直播的流:")
	listLiveStreams(hub, "carter")

	fmt.Println("禁用流:")
	disableStream(hub, keyA)

	fmt.Println("启用流:")
	enableStream(hub, keyA)

	fmt.Println("查询直播状态:")
	liveStatus(hub, keyA)

	fmt.Println("查询推流历史:")
	historyActivity(hub, keyA)

	fmt.Println("保存直播数据:")
	savePlayback(hub, keyA)

	fmt.Println("RTMP 推流地址:")
	url := pili.RTMPPublishURL("publish-rtmp.test.com", HubName, keyA, mac, 3600)
	fmt.Println(url)

	fmt.Println("RTMP 直播放址:")
	url = pili.RTMPPlayURL("live-rtmp.test.com", HubName, keyA)
	fmt.Println(url)

	fmt.Println("HLS 直播地址:")
	url = pili.HLSPlayURL("live-hls.test.com", HubName, keyA)
	fmt.Println(url)

	fmt.Println("HDL 直播地址:")
	url = pili.HDLPlayURL("live-hdl.test.com", HubName, keyA)
	fmt.Println(url)

	fmt.Println("截图直播地址:")
	url = pili.SnapshotPlayURL("live-snapshot.test.com", HubName, keyA)
	fmt.Println(url)
}
