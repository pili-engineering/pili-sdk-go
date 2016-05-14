# Pili Streaming Cloud server-side library for Golang

## Features

- URL
	- [x] RTMP推流地址: client.RTMPPublishURL(domain, hub, streamKey, expireAfterDays)
	- [x] RTMP直播地址: RTMPPlayURL(domain, hub, streamKey)
	- [x] HLS直播地址: HLSPlayURL(domain, hub, streamKey)
	- [x] HDL直播地址: HDLPlayURL(domain, hub, streamKey)
	- [x] 截图直播地址: SnapshotPlayURL(domain, hub, streamKey)
- Hub
	- [x] 创建流: hub.Create(streamKey)
	- [x] 查询流: hub.Get(streamKey)
	- [x] 列出流: hub.List(prefix, limit, marker)
	- [x] 列出正在直播的流: hub.ListLive(prefix, limit, marker)
- Stream
	- [x] 禁用流: stream.Disable()
	- [x] 启用流: stream.Enable()
 	- [x] 查询直播状态: stream.LiveStatus()
	- [x] 保存直播回放: stream.Save(start, end)
	- [x] 查询直播历史: stream.HistoryRecord(start, end)

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
		- [Disable a Stream](#disable-a-stream)
		- [Enable a Stream](#enable-a-stream)
		- [Get Stream live status](#get-stream-live-status)
		- [Get Stream history record](#get-stream-history-record)
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
	client := pili.New(AccessKey, SecretKey, nil)
	// ...
}
```

### URL

#### Generate RTMP publish URL

```go
url := client.RTMPPublishURL("publish-rtmp.test.com", "PiliSDKTest", "streamkey", 60)
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
	client := pili.New(AccessKey, SecretKey, nil)
	hub := pili.NewHub("PiliSDKTest", client)
	// ...
}
```

#### Create a new Stream

```go
stream, err := hub.Create("streamkey")
if err != nil {
	return
}
fmt.Println(toJSON(stream))
/*
{"Hub":"PiliSDKTest","Key":"streamkey","DisabledTill":0}
*/
```

#### Get a Stream

```go
stream, err = hub.Get("streamkey")
if err != nil {
	return
}
fmt.Println(toJSON(stream))
/*
{"Hub":"PiliSDKTest","Key":"streamkey","DisabledTill":0}
*/
```

#### List Streams

```go
keys, marker, err := hub.List("str", 10, "")
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
keys, marker, err := hub.ListLive("str", 10, "")
if err != nil {
	return
}
fmt.Printf("keys=%v marker=%v\n", keys, marker)
/*
keys=[] marker=
*/
```

### Stream

#### Disable a Stream

```go
stream, err := hub.Get("streamkey")
if err != nil {
	return
}
fmt.Println("before disable:", toJSON(stream))

err = stream.Disable()
if err != nil {
	return
}

stream, err = hub.Get("streamkey")
if err != nil {
	return
}
fmt.Println("after disable:", toJSON(stream))
/*
before disable: {"Hub":"PiliSDKTest","Key":"streamkey","DisabledTill":0}
after disable: {"Hub":"PiliSDKTest","Key":"streamkey","DisabledTill":-1}
*/
```

#### Enable a Stream

```go
stream, err := hub.Get("streamkey")
if err != nil {
	return
}
fmt.Println("before enable:", toJSON(stream))

err = stream.Enable()
if err != nil {
	return
}

stream, err = hub.Get("streamkey")
if err != nil {
	return
}
fmt.Println("after enable:", toJSON(stream))
/*
before disable: {"Hub":"PiliSDKTest","Key":"streamkey","DisabledTill":-1}
after disable: {"Hub":"PiliSDKTest","Key":"streamkey","DisabledTill":0}
*/
```

#### Get Stream live status

```go
status, err := stream.LiveStatus()
if err != nil {
	return
}
fmt.Println(toJSON(status))
/*
{"startAt":1463022236,"clientIP":"222.73.202.226","bps":248,"fps":{"audio":45,"vedio":28,"data":0}}
*/
```

#### Get Stream history record

```go
records, err := stream.HistoryRecord(0, 0)
if err != nil {
	return
}
fmt.Println(records)
/*
[{1463022236,1463022518}]
*/
```

#### Save Stream live playback

```go
fname, err := stream.Save(0, 0)
if err != nil {
	return
}
fmt.Println(fname)
/*
recordings/z1.PiliSDKTest.streamkey/1463156847_1463157463.m3u8
*/
```
