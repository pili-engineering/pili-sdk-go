# Pili Streaming Cloud server-side library for Golang

## Features

- URL
	- [x] RTMP推流地址: RTMPPublishURL(domain, hub, streamKey, mac, expireAfterDays)
	- [x] RTMP直播地址: RTMPPlayURL(domain, hub, streamKey)
	- [x] HLS直播地址: HLSPlayURL(domain, hub, streamKey)
	- [x] HDL直播地址: HDLPlayURL(domain, hub, streamKey)
	- [x] 截图直播地址: SnapshotPlayURL(domain, hub, streamKey)
- Hub
	- [x] 创建流: hub.Create(streamKey)
	- [x] 获得流: hub.Stream(streamKey)
	- [x] 列出流: hub.List(prefix, limit, marker)
	- [x] 列出正在直播的流: hub.ListLive(prefix, limit, marker)
- Stream
	- [x] 流信息: stream.Info()
	- [x] 禁用流: stream.Disable()
	- [x] 启用流: stream.Enable()
 	- [x] 查询直播状态: stream.LiveStatus()
	- [x] 保存直播回放: stream.Save(start, end)
	- [x] 查询直播历史: stream.HistoryActivity(start, end)

## Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Configuration](#configuration)
	- [URL](#url)
		- [Generate RTMP publish URL](#generate-rtmp-publish-url)
		- [Generate RTMP play URL](#generate-rtmp-play-url)
		- [Generate HLS play URL](#generate-hls-play-url)
		- [Generate HDL play URL](#generate-hdl-play-url)
		- [Generate Snapshot play URL](#generate-snapshot-play-url)
	- [Hub](#hub)
		- [Instantiate a Pili Hub object](#instantiate-a-pili-hub-object)
		- [Create a new Stream](#create-a-new-stream)
		- [Get a Stream](#get-a-stream)
		- [List Streams](#list-streams)
		- [List live Streams](#list-live-streams)
	- [Stream](#stream)
		- [Get Stream info](#get-stream-info)
		- [Disable a Stream](#disable-a-stream)
		- [Enable a Stream](#enable-a-stream)
		- [Get Stream live status](#get-stream-live-status)
		- [Get Stream history activity](#get-stream-history-activity)
		- [Save Stream live playback](#save-stream-live-playback)

## Installation

before next step, install git.

```
// install latest version
$ go get github.com/pili-engineering/pili-sdk-go/pili
```

## Usage

### Configuration

```go
package main

import (
	// ...
	"github.com/pili-engineering/pili-sdk-go/pili"
)

var (
	AccessKey = "<QINIU ACCESS KEY>" // 替换成自己 Qiniu 账号的 AccessKey.
	SecretKey = "<QINIU SECRET KEY>" // 替换成自己 Qiniu 账号的 SecretKey.
	HubName   = "<PILI HUB NAME>"    // Hub 必须事先存在.
)

func main() {
	// ...
	mac := &pili.MAC{AccessKey, []byte(SecretKey)}
	client := pili.New(mac, nil)
	// ...
}
```

### URL

#### Generate RTMP publish URL

```go
url := pili.RTMPPublishURL("publish-rtmp.test.com", "PiliSDKTest", "streamkey", mac, 60)
fmt.Println(url)
/*
rtmp://publish-rtmp.test.com/PiliSDKTest/streamkey?e=1463023142&token=7O7hf7Ld1RrC_fpZdFvU8aCgOPuhw2K4eapYOdII:-5IVlpFNNGJHwv-2qKwVIakC0ME=
*/
```

#### Generate RTMP play URL

```go
url := pili.RTMPPlayURL("live-rtmp.test.com", "PiliSDKTest", "streamkey")
fmt.Println(url)
/*
rtmp://live-rtmp.test.com/PiliSDKTest/streamkey
*/
```

#### Generate HLS play URL

```go
url := pili.HLSPlayURL("live-hls.test.com", "PiliSDKTest", "streamkey")
fmt.Println(url)
/*
http://live-hls.test.com/PiliSDKTest/streamkey.m3u8
*/
```

#### Generate HDL play URL

```go
url := pili.HDLPlayURL("live-hdl.test.com", "PiliSDKTest", "streamkey")
fmt.Println(url)
/*
http://live-hdl.test.com/PiliSDKTest/streamkey.flv
*/
```

#### Generate Snapshot play URL

```go
url := pili.SnapshotPlayURL("live-snapshot.test.com", "PiliSDKTest", "streamkey")
fmt.Println(url)
/*
http://live-snapshot.test.com/PiliSDKTest/streamkey.jpg
*/
```

### Hub

#### Instantiate a Pili Hub object

```go
func main() {
	mac := &pili.MAC{AccessKey, []byte(SecretKey)}
	client := pili.New(mac, nil)
	hub := client.Hub("PiliSDKTest")
	// ...
}
```

#### Create a new Stream

```go
stream, err := hub.Create(key)
if err != nil {
	return
}
info, err := stream.Info()
if err != nil {
	return
}
fmt.Println(info)
/*
{hub:PiliSDKTest,key:streamkey,disabled:false}
*/
```

#### Get a Stream

```go
stream := hub.Stream(key)
info, err := stream.Info()
if err != nil {
	return
}
fmt.Println(info)
/*
{hub:PiliSDKTest,key:streamkey,disabled:false}
*/
```

#### List Streams

```go
keys, marker, err := hub.List(prefix, 10, "")
if err != nil {
	return
}
fmt.Printf("keys=%v marker=%v\n", keys, marker)
/*
keys=[streamkey] marker=
*/
```

#### List live Streams

```go
keys, marker, err := hub.ListLive(prefix, 10, "")
if err != nil {
	return
}
fmt.Printf("keys=%v marker=%v\n", keys, marker)
/*
keys=[streamkey] marker=
*/
```

### Stream

#### Get Stream info

```go
stream := hub.Stream(key)
info, err := stream.Info()
if err != nil {
	return
}
fmt.Println(info)
/*
{hub:PiliSDKTest,key:streamkey,disabled:false}
*/
```

#### Disable a Stream

```go
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
/*
before disable: {hub:PiliSDKTest,key:streamkey,disabled:false}
after disable: {hub:PiliSDKTest,key:streamkey,disabled:true}
*/
```

#### Enable a Stream

```go
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
/*
before enable: {hub:PiliSDKTest,key:streamkey,disabled:true}
after enable: {hub:PiliSDKTest,key:streamkey,disabled:false}
*/
```

#### Get Stream live status

```go
stream := hub.Stream(key)
status, err := stream.LiveStatus()
if err != nil {
	return
}
fmt.Printf("%+v\n", status)
/*
&{StartAt:1463382400 ClientIP:172.21.1.214:52897 BPS:128854 FPS:{Audio:38 Video:23 Data:0}}
*/
```

#### Get Stream history activity

```go
stream := hub.Stream(key)
records, err := stream.HistoryActivity(0, 0)
if err != nil {
	return
}
fmt.Println(records)
/*
[{1463382401 1463382441}]
*/
```

#### Save Stream live playback

```go
stream := hub.Stream(key)
fname, err := stream.Save(0, 0)
if err != nil {
	return
}
fmt.Println(fname)
/*
recordings/z1.PiliSDKTest.streamkey/1463156847_1463157463.m3u8
*/
```
