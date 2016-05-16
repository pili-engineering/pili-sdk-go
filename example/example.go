package main

import (
	"encoding/json"
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

func toJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
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

	// 初始化 client.
	mac := &pili.MAC{AccessKey: AccessKey, SecretKey: []byte(SecretKey)}
	client := pili.New(mac, nil)

	// 初始化 Hub.
	hub := client.Hub(HubName)

	// 获得不存在的流.
	keyA := streamKeyPrefix + "A"
	_, err := hub.Get(keyA)
	if !pili.IsNotExists(err) {
		log.Printf("WARN: keyA=%s should be not exist but got err=(%v)\n", keyA, err)
		return
	}
	fmt.Printf("keyA=%s 不存在\n", keyA)

	// 创建流.
	_, err = hub.Create(keyA)
	if err != nil {
		log.Printf("ERROR: keyA=%s create failed err=(%v)\n", keyA, err)
		return
	}
	fmt.Printf("keyA=%s 创建\n", keyA)

	// 获得流.
	streamA, err := hub.Get(keyA)
	if err != nil {
		log.Printf("ERROR: keyA=%s get failed err=(%v)\n", keyA, err)
		return
	}
	fmt.Printf("keyA=%s 查询: %v\n", keyA, toJSON(streamA))

	// 创建重复的流.
	_, err = hub.Create(keyA)
	if !pili.IsExists(err) {
		log.Printf("ERROR: keyA=%s should be exists but got err=(%v)\n", keyA, err)
		return
	}
	fmt.Printf("keyA=%s 已存在\n", keyA)

	// 创建另一路流.
	keyB := streamKeyPrefix + "B"
	streamB, err := hub.Create(keyB)
	if err != nil {
		log.Printf("ERROR: keyB=%s create failed %v\n", keyB, err)
		return
	}
	fmt.Printf("keyB=%s 创建: %v\n", keyB, toJSON(streamB))

	// 列出所有流.
	keys, marker, err := hub.List(streamKeyPrefix, 0, "")
	if err != nil {
		log.Printf("ERROR: list failed %v\n", err)
		return
	}
	fmt.Printf("hub=%s 列出流: keys=%v marker=%v\n", HubName, keys, marker)

	// 列出正在直播的流.
	keys, marker, err = hub.ListLive(streamKeyPrefix, 0, "")
	if err != nil {
		log.Printf("ERROR: list live failed %v\n", err)
		return
	}
	fmt.Printf("hub=%s 列出正在直播的流: keys=%v marker=%v\n", HubName, keys, marker)

	// 禁用流.
	err = streamA.Disable()
	if err != nil {
		log.Printf("ERROR: keyA=%s disable failed err=(%v)\n", keyA, err)
		return
	}
	streamA, err = hub.Get(keyA)
	if err != nil {
		log.Printf("ERROR: keyA=%s get failed err=(%v)\n", keyA, err)
		return
	}
	fmt.Printf("keyA=%s 禁用: %v\n", keyA, toJSON(streamA))

	// 启用流.
	err = streamA.Enable()
	if err != nil {
		log.Printf("ERROR: keyA=%s enable failed %v\n", keyA, err)
		return
	}
	streamA, err = hub.Get(keyA)
	if err != nil {
		log.Printf("ERROR: keyA=%s get failed %v\n", keyA, err)
		return
	}
	fmt.Printf("keyA=%s 启用: %v\n", keyA, toJSON(streamA))

	// 查询直播状态.
	status, err := streamA.LiveStatus()
	fmt.Printf("keyA=%s 直播状态: status=%v err=(%v)\n", keyA, toJSON(status), err)

	// 查询推流历史.
	records, err := streamA.HistoryRecord(0, 0)
	fmt.Printf("keyA=%s 推流历史: records=%v err=(%v)\n", keyA, records, err)

	// 保存直播数据.
	fname, err := streamA.Save(0, 0)
	fmt.Printf("keyA=%s 保存直播数据: fname=%s err=(%v)\n", keyA, fname, err)

	// RTMP 推流地址.
	url := pili.RTMPPublishURL("publish-rtmp.test.com", HubName, keyA, mac, 3600)
	fmt.Printf("keyA=%s RTMP推流地址=%s\n", keyA, url)

	// RTMP 直播放址.
	url = pili.RTMPPlayURL("live-rtmp.test.com", HubName, keyA)
	fmt.Printf("keyA=%s RTMP直播地址=%s\n", keyA, url)

	// HLS 直播地址.
	url = pili.HLSPlayURL("live-hls.test.com", HubName, keyA)
	fmt.Printf("keyA=%s HLS直播地址=%s\n", keyA, url)

	// HDL 直播地址.
	url = pili.HDLPlayURL("live-hdl.test.com", HubName, keyA)
	fmt.Printf("keyA=%s HDL直播地址=%s\n", keyA, url)

	// 截图直播地址
	url = pili.SnapshotPlayURL("live-snapshot.test.com", HubName, keyA)
	fmt.Printf("keyA=%s 截图直播地址=%s\n", keyA, url)
}
