# PILI直播 Go服务端SDK 使用指南

## 功能列表

- 直播流的创建、获取和列举
    - [x] hub.createStream()  // 创建流
    - [x] hub.getStream()  // 获取流
    - [x] hub.listStreams()  // 列举流
- 直播流的其他功能
    - [x] stream.toJsonString()  // 流信息转为json
    - [x] stream.update()      // 更新流
    - [x] stream.disable()      // 禁用流
    - [x] stream.enable()    // 启用流
    - [x] stream.rtmpPublishUrl()   // 生成推流地址
    - [x] stream.rtmpLiveUrls()    // 生成rtmp播放地址
    - [x] stream.hlsLiveUrls()   // 生成hls播放地址
    - [x] stream.httpFlvLiveUrls()   // 生成flv播放地址
    - [x] stream.status()     // 获取流状态
    - [x] stream.segments()      // 获取流片段
    - [x] stream.hlsPlaybackUrls()  // 生成hls回看地址
    - [x] stream.saveAs()        // 流另存为文件
    - [x] stream.snapshot()      // 获取快照
    - [x] stream.delete()    // 删除流

## 目录

- [安装](#installation)
- [用法](#usage)
    - [配置](#configuration)
    - [Hub](#hub)
        - [实例化hub对象](#instantiate-a-pili-hub-object)
        - [创建流](#create-a-new-stream)
        - [获取流](#get-an-exist-stream)
        - [列举流](#list-streams)
    - [直播流](#stream)
        - [流信息转为json](#to-json-string)
        - [更新流](#update-a-stream)
        - [禁用流](#disable-a-stream)
        - [启用流](#enable-a-stream)
        - [生成推流地址](#generate-rtmp-publish-url)
        - [生成rtmp播放地址](#generate-rtmp-live-play-urls)
        - [生成hls播放地址](#generate-hls-play-urls)
        - [生成flv播放地址](#generate-http-flv-live-play-urls)
        - [获取流状态](#get-stream-status)
        - [获取流片段](#get-stream-segments)
        - [生成hls回看地址](#generate-hls-playback-urls)
        - [流另存为文件](#save-stream-as-a-file)
        - [获取快照](#snapshot-stream)
        - [删除流](#delete-a-stream)
- [History](#history)


<a id="installation"></a>
## 安装

```
// 安装最新版本
$ go get github.com/pili-engineering/pili-sdk-go/pili
```

<a id="usage"></a>
## 用法:

<a id="configuration"></a>
### 配置

```go
import (
    "github.com/pili-engineering/pili-sdk-go/pili"
    // ...
)

const (
  ACCESS_KEY = "Qiniu_AccessKey"
  SECRET_KEY = "Qiniu_SecretKey"
  HUB        = "Pili_HubName"   // The Hub must be exists before use
)

func main() {

    // Change API host as necessary
    //
    // pili.qiniuapi.com as default
    // pili-lte.qiniuapi.com is the latest RC version
    //
    // pili.API_HOST = "pili.qiniuapi.com" // default

}
```

### Hub

<a id="instantiate-a-pili-hub-object"></a>
#### 实例化hub对象

```go
func main() {

  credentials := pili.NewCredentials(ACCESS_KEY, SECRET_KEY)
  hub := pili.NewHub(credentials, HUB_NAME)

    // ...
}
```

<a id="create-a-new-stream"></a>
#### 创建流

```go
options := pili.OptionalArguments{               // 选填
    Title:           "stream_title", // 选填，默认自动生成
    PublishKey:      "some_secret_words",        // 选填，默认自动生成
    PublishSecurity: "dynamic",                  // 选填, 可以为 "dynamic" 或 "static", 默认为 "dynamic"
}
stream, err := hub.CreateStream(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("CreateStream:\n", stream)
/*
{
    0xc208036018
    Id: z1.hub1.stream_title
    CreatedAt: 2015-08-22 15:37:20.397 +0800 CST
    UpdatedAt: 2015-08-24 09:41:55.32 +0800 CST
    Title: stream_title
    Hub: hub1
    Disabled: false
    PublishKey: some_secret_words
    PublishSecurity: dynamic
    Profiles: []
    Hosts: {
        Publish: map[
            rtmp:ec2s3f5.publish.z1.pili.qiniup.com
        ]
        Live: map[
            http:ec2s3f5.live1-http.z1.pili.qiniucdn.com
            rtmp:ec2s3f5.live1-rtmp.z1.pili.qiniucdn.com
        ]
        Playback: map[
            http:ec2s3f5.playback1.z1.pili.qiniucdn.com
        ]
    }
}
*/
```

<a id="get-an-exist-stream"></a>
#### 获取流

```go
stream, err = hub.GetStream(stream.Id)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("GetStream:\n", stream)
/*
{
    0xc208036018
    Id: z1.hub1.stream_title
    CreatedAt: 2015-08-22 15:37:20.397 +0800 CST
    UpdatedAt: 2015-08-24 09:41:55.32 +0800 CST
    Title: stream_title
    Hub: hub1
    Disabled: false
    PublishKey: some_secret_words
    PublishSecurity: dynamic
    Profiles: []
    Hosts: {
        Publish: map[
            rtmp:ec2s3f5.publish.z1.pili.qiniup.com
        ]
        Live: map[
            http:ec2s3f5.live1-http.z1.pili.qiniucdn.com
            rtmp:ec2s3f5.live1-rtmp.z1.pili.qiniucdn.com
        ]
        Playback: map[
            http:ec2s3f5.playback1.z1.pili.qiniucdn.com
        ]
    }
}
*/
```

<a id="list-streams"></a>
#### 列举流

```go
options = pili.OptionalArguments{ // optional
    Marker: "",        // optional, returned by server response
    Limit:  50,        // optional
    Title:  "prefix_", // optional, title prefix
}
listResult, err := hub.ListStreams(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("ListStreams:\n", listResult)
for _, stream := range listResult.Items {
    fmt.Println("Stream:\n", stream)
}
/*
{1 [0xc208036018]}
&{
    0xc208036018
    Id: z1.hub1.stream_title
    CreatedAt: 2015-08-22 15:37:20.397 +0800 CST
    UpdatedAt: 2015-08-24 09:41:55.32 +0800 CST
    Title: stream_title
    Hub: hub1
    Disabled: false
    PublishKey: some_secret_words
    PublishSecurity: dynamic
    Profiles: []
    Hosts: {
        Publish: map[
            rtmp:ec2s3f5.publish.z1.pili.qiniup.com
        ]
        Live: map[
            http:ec2s3f5.live1-http.z1.pili.qiniucdn.com
            rtmp:ec2s3f5.live1-rtmp.z1.pili.qiniucdn.com
        ]
        Playback: map[
            http:ec2s3f5.playback1.z1.pili.qiniucdn.com
        ]
    }
}
*/
```

<a id="stream"></a>
### 直播流

<a id="to-json-string"></a>
#### 流信息转为json
```go
streamJson, err := stream.ToJSONString()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream ToJSONString:\n", streamJson)
/*
{
    "id":"z1.hub1.stream_title",
    "createdAt":"2015-08-22T15:37:20.397+08:00",
    "updatedAt":"2015-08-24T09:41:55.32+08:00",
    "title":"stream_title",
    "hub":"hub1",
    "disabled":false,
    "publishKey":"some_secret_words",
    "publishSecurity":"dynamic",
    "hosts":{
        "publish":{
            "rtmp":"ec2s3f5.publish.z1.pili.qiniup.com"
        },
        "live":{
            "http":"ec2s3f5.live1-http.z1.pili.qiniucdn.com",
            "rtmp":"ec2s3f5.live1-rtmp.z1.pili.qiniucdn.com"
        },
        "playback":{
            "http":"ec2s3f5.playback1.z1.pili.qiniucdn.com"
        }
    }
}
*/
```

<a id="update-a-stream"></a>
#### 更新流

```go
stream.PublishKey = "new_secret_words" // optional
stream.PublishSecurity = "static"      // optional
stream, err = stream.Update()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Updated:\n", stream)
/*
{
    0xc208036018
    Id: z1.hub1.stream_title
    CreatedAt: 2015-08-22 15:37:20.397 +0800 CST
    UpdatedAt: 2015-08-24 09:41:55.32 +0800 CST
    Title: stream_title
    Hub: hub1
    Disabled: false
    PublishKey: new_secret_words
    PublishSecurity: static
    Profiles: []
    Hosts: {
        Publish: map[
            rtmp:ec2s3f5.publish.z1.pili.qiniup.com
        ]
        Live: map[
            http:ec2s3f5.live1-http.z1.pili.qiniucdn.com
            rtmp:ec2s3f5.live1-rtmp.z1.pili.qiniucdn.com
        ]
        Playback: map[
            http:ec2s3f5.playback1.z1.pili.qiniucdn.com
        ]
    }
}
*/
```

<a id="disable-a-stream"></a>
#### 禁用流

```go
stream, err = stream.Disable()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Disabled:\n", stream.Disabled)
/*
true
*/
```

<a id="enable-a-stream"></a>
#### 启用流

```go
stream, err = stream.Enable()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Enabled:\n", stream.Disabled)
/*
false
*/
```

<a id="generate-rtmp-publish-url"></a>
#### 生成推流地址

```go
url := stream.RtmpPublishUrl()
fmt.Println("Stream RtmpPublishUrl:\n", url)
/*
rtmp://ec2s3f5.publish.z1.pili.qiniup.com/hub1/stream_title?key=new_secret_words
*/
```

<a id="generate-rtmp-live-play-urls"></a>
#### 生成rtmp播放地址

```go
urls, err := stream.RtmpLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("RtmpLiveUrls:", urls)
/*
map[ORIGIN:rtmp://ec2s3f5.live1-rtmp.z1.pili.qiniucdn.com/hub1/stream_title]
*/
```

<a id="generate-hls-play-urls"></a>
#### 生成hls播放地址

```go
urls, err = stream.HlsLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HlsLiveUrls:", urls)
/*
map[ORIGIN:http://ec2s3f5.live1-http.z1.pili.qiniucdn.com/hub1/stream_title.m3u8]
*/
```

<a id="generate-http-flv-live-play-urls"></a>
#### 生成flv播放地址

```go
urls, err = stream.HttpFlvLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HttpFlvLiveUrls:", urls)
/*
map[ORIGIN:http://ec2s3f5.live1-http.z1.pili.qiniucdn.com/hub1/stream_title.flv]
*/
```

<a id="get-stream-status"></a>
#### 获取流状态

```go
streamStatus, err := stream.Status()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Status:\n", streamStatus)
/*
 {
    Addr: 114.81.254.172:36317
    Status: connected
    BytesPerSecond: 16870.200000000001
    FramesPerSecond: {
        Audio: 42.200000000000003
        Video: 14.733333333333333
        Data: 0.066666666666666666
    }
}
*/
```

<a id="get-stream-segments"></a>
#### 获取流片段 

```go
options = pili.OptionalArguments{ // 可选
    Start: 1440379800, // 可选, 单位为秒, 为UNIX时间戳
    End:   1440479880, // 选填, 单位为秒, 为UNIX时间戳
    Limit: 20,         // 选填, uint
}
segments, err := stream.Segments(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Segments:\n", segments)
/*
{[0xc20800b2c0 0xc20800b320 0xc20800b350 0xc20800b380 0xc20800b3b0 0xc20800b3e0]}
*/
```

<a id="generate-hls-playback-urls"></a>
#### 生成hls回看地址

```go
start := 1440379847
end := 1440379857
urls, err = stream.HlsPlaybackUrls(int64(start), int64(end))
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HlsPlaybackUrls:", urls)
/*
map[ORIGIN:http://ec2s3f5.playback1.z1.pili.qiniucdn.com/hub1/stream_title.m3u8?start=1440379847&end=1440379857]
*/
```

<a id="save-stream-as-a-file"></a>
#### 流另存为文件

```go
name := "fileName.mp4" // 必填, string 类型
format := "mp4"        // 必填, string 类型
start = 1440379847     // 必填, int64, 单位为秒, 为UNIX时间戳
end = 1440379857       // 必填, int64, 单位为秒, 为UNIX时间戳
options = pili.OptionalArguments{
    NotifyUrl: "http://remote_callback_url",
} // 选填
saveAsRes, err := stream.SaveAs(name, format, int64(start), int64(end), options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream save as:\n", saveAsRes)
/*
{
    Url: http://ec2s3f5.vod1.z1.pili.qiniucdn.com/recordings/z1.hub1.stream_title/fileName.m3u8
    TargetUrl: http://ec2s3f5.vod1.z1.pili.qiniucdn.com/recordings/z1.hub1.stream_title/fileName.mp4
    PersistentId: z1.55da7715f51b82403b01e985
}
*/
```

当使用 `saveAs()` 和 `snapshot()` 的时候, 由于是异步处理， 你可以在七牛的FOP接口上使用 `persistentId`来获取处理进度.参考如下：   
API: `curl -D GET http://api.qiniu.com/status/get/prefop?id={persistentId}`  
文档说明: <http://developer.qiniu.com/docs/v6/api/overview/fop/persistent-fop.html#pfop-status> 

<a id="snapshot-stream"></a>
#### 获取快照

```go
name = "fileName.jpg" // 必填, string 类型
format = "jpg"        // 必填, string 类型
options = pili.OptionalArguments{
    Time:      1440379847, // 选填, int64, 单位为秒, 为UNIX时间戳
    NotifyUrl: "http://remote_callback_url",
} // 选填
snapshotRes, err := stream.Snapshot(name, format, options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Snapshot:\n", snapshotRes)
/*
{
    TargetUrl: http://ec2s3f5.static1.z1.pili.qiniucdn.com/snapshots/z1.hub1.stream_title/fileName.jpg
    PersistentId: z1.55da7716f51b82403b01e986
}
*/
```

<a id="delete-a-stream"></a>
#### 删除流

```go
deleteResult, err := stream.Delete()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Deleted:\n", deleteResult)
/*
<nil>
*/
```

## History

- 1.5.0
    - Add stream.HttpFlvLiveUrls()
    - Add stream.Snapshot(name,format string, options={time, notifyUrl})

- 1.3.1
    - Updated stream.Update() logic

- 1.3.0
    - Add stream.SaveAs(name,format string, start,end int64, options={notifyUrl})

- 1.2.0
    - Add Stream operations
      - stream.ToJSONString()
        - stream.Update(options={PublishKey,PublishSecurity})
      - stream.Disable()
      - stream.Enable()
      - stream.RtmpPublishUrl()
      - stream.RtmpLiveUrls()
      - stream.HlsLiveUrls()
      - stream.Status()
        - stream.Segments(options={Start,End, Limit})
        - stream.HlsPlaybackUrls(start, end int64)
      - stream.Delete()
    - Update hub functions
        - hub.CreateStream(options={Title,PublishKey,PublishSecurity})
      - hub.GetStream(stream.Id)
        - hub.ListStreams(options={Marker,Limit, Title})
