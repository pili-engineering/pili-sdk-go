# Pili Streaming Cloud server-side library for Golang

## Features

- Stream Create,Get,List
    - [x] hub.CreateStream(options={Title,PublishKey,PublishSecurity})
    - [x] hub.GetStream(stream.Id)
    - [x] hub.ListStreams(options={Marker,Limit, Title})
- Stream operations else
    - [x] stream.ToJSONString()
    - [x] stream.Status()
    - [x] stream.Update(options={PublishKey,PublishSecurity})
    - [x] stream.Refresh()
    - [x] stream.RtmpPublishUrl()
    - [x] stream.RtmpLiveUrls()
    - [x] stream.HlsLiveUrls()
    - [x] stream.HttpFlvLiveUrls()
    - [x] stream.HlsPlaybackUrls(start, end int64)
    - [x] stream.Segments(options={Start,End, Limit})
    - [x] stream.SaveAs(name,format string, start,end int64, options={notifyUrl})
    - [x] stream.Snapshot(name,format string, options={time, notifyUrl})
    - [x] stream.Delete()

## Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Configuration](#configuration)
    - [Hub](#hub)
        - [Instantiate a Pili Hub object](#instantiate-a-pili-hub-object)
        - [Create a new Stream](#create-a-new-stream)
        - [Get a Stream](#get-a-stream)
        - [List Streams](#List-streams)
    - [Stream](#stream)
        - [To JSON string](#to-json-string)
        - [Update a Stream](#update-a-stream)
        - [Disable a Stream](#disable-a-stream)
        - [Enable a Stream](#enable-a-stream)
        - [Generate RTMP publish URL](#generate-rtmp-publish-url)
        - [Generate RTMP live play URLs](#generate-rtmp-live-play-urls)
        - [Generate HLS live play URLs](generate-hls-live-play-urls)
        - [Generate Http-Flv live play URLs](generate-http-flv-live-play-urls)
        - [Get Stream status](#get-stream-status)
        - [Get Stream segments](#get-stream-segments)
        - [Generate HLS playback URLs](generate-hls-playback-urls)
        - [Save Stream as a file](#save-stream-as-a-file)
        - [Snapshot Stream](#snapshot-stream)
        - [Delete a Stream](#delete-a-stream)
- [History](#history)


## Installaion

```
// install latest version
$ go get github.com/pili-engineering/pili-sdk-go/pili
```

## Usage

### Configuration

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
    // pili.qiniuapi.com as deafult
    // pili-lte.qiniuapi.com is the latest RC version
    //
    pili.API_HOST = "pili-lte.qiniuapi.com"

}
```

### Hub

#### Instantiate an Pili Hub object

```go

func main() {

	credentials := pili.NewCredentials(ACCESS_KEY, SECRET_KEY)
	hub := pili.NewHub(credentials, HUB_NAME)

    // ...
}
```

#### Create a new Stream

```go
options := pili.OptionalArguments{        // optional
    Title:           "stream_name",       // optional, auto-generated as default
    PublishKey:      "some_secret_words", // optional, auto-generated as default
    PublishSecurity: "dynamic",           // optional, can be "dynamic" or "static", "dynamic" as default
}
stream, err := hub.CreateStream(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("CreateStream:\n", stream)
/*
{0xc208036018 z1.coding.55d826b0d409d209d93f6e37 2015-08-22 15:37:20.397 +0800 CST 2015-08-24 09:41:55.32 +0800 CST 55d826b0d409d209d93f6e37 coding false new_secret_words static [] {map[rtmp:as2h9t.publish.z1.pili.qiniup.com] map[http:as2h9t.live1-http.z1.pili.qiniucdn.com rtmp:as2h9t.live1-rtmp.z1.pili.qiniucdn.com] map[http:as2h9t.playback1.z1.pili.qiniucdn.com]}}
*/
```

#### Get a Stream

```go
stream, err = hub.GetStream(stream.Id)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("GetStream:\n", stream)
/*
 {0xc208036018 z1.coding.55d826b0d409d209d93f6e37 2015-08-22 15:37:20.397 +0800 CST 2015-08-24 09:41:55.32 +0800 CST 55d826b0d409d209d93f6e37 coding false new_secret_words static [] {map[rtmp:as2h9t.publish.z1.pili.qiniup.com] map[http:as2h9t.live1-http.z1.pili.qiniucdn.com rtmp:as2h9t.live1-rtmp.z1.pili.qiniucdn.com] map[http:as2h9t.playback1.z1.pili.qiniucdn.com]}}
*/
```

#### List streams

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
&{0xc208036018 z1.coding.55d826b0d409d209d93f6e37 2015-08-22 15:37:20.397 +0800 CST 2015-08-24 09:41:55.32 +0800 CST 55d826b0d409d209d93f6e37 coding false new_secret_words static [] {map[rtmp:as2h9t.publish.z1.pili.qiniup.com] map[http:as2h9t.live1-http.z1.pili.qiniucdn.com rtmp:as2h9t.live1-rtmp.z1.pili.qiniucdn.com] map[http:as2h9t.playback1.z1.pili.qiniucdn.com]}}
*/
```

### Stream

#### To JSON String
```go
streamJson, err := stream.ToJSONString()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream ToJSONString:\n", streamJson)
/*
{"id":"z1.coding.55d826b0d409d209d93f6e37","createdAt":"2015-08-22T15:37:20.397+08:00","updatedAt":"2015-08-24T09:41:55.32+08:00","title":"55d826b0d409d209d93f6e37","hub":"coding","disabled":false,"publishKey":"new_secret_words","publishSecurity":"static","hosts":{"publish":{"rtmp":"as2h9t.publish.z1.pili.qiniup.com"},"live":{"http":"as2h9t.live1-http.z1.pili.qiniucdn.com","rtmp":"as2h9t.live1-rtmp.z1.pili.qiniucdn.com"},"playback":{"http":"as2h9t.playback1.z1.pili.qiniucdn.com"}}}
*/
```

#### Update a Stream

```go
stream.PublishKey = "new_secret_words" // optional
stream.PublishSecurity = "static"      // optional
stream, err = stream.Update()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Updated:\n", stream)
/*
{0xc208036018 z1.coding.55d826b0d409d209d93f6e37 2015-08-22 15:37:20.397 +0800 CST 2015-08-24 09:44:52.318147076 +0800 CST 55d826b0d409d209d93f6e37 coding false new_secret_words static [] {map[rtmp:as2h9t.publish.z1.pili.qiniup.com] map[http:as2h9t.live1-http.z1.pili.qiniucdn.com rtmp:as2h9t.live1-rtmp.z1.pili.qiniucdn.com] map[http:as2h9t.playback1.z1.pili.qiniucdn.com]}}
*/
```

#### Disable a stream

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

#### Enable a Stream

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

#### Generate RTMP publish URL

```go
url := stream.RtmpPublishUrl()
fmt.Println("Stream RtmpPublishUrl:\n", url)
/*
rtmp://as2h9t.publish.z1.pili.qiniup.com/coding/55d826b0d409d209d93f6e37?key=new_secret_words
*/
```

#### Generate RTMP live play URLs

```go
urls, err := stream.RtmpLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("RtmpLiveUrls:", urls)
/*
map[ORIGIN:rtmp://as2h9t.live1-rtmp.z1.pili.qiniucdn.com/coding/55d826b0d409d209d93f6e37]
*/
```

#### Generate HLS play live URLs

```go
urls, err = stream.HlsLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HlsLiveUrls:", urls)
/*
map[ORIGIN:http://as2h9t.live1-http.z1.pili.qiniucdn.com/coding/55d826b0d409d209d93f6e37.m3u8]
*/
```

#### Generate Http-Flv live play URLs

```go
urls, err = stream.HttpFlvLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HttpFlvLiveUrls:", urls)
/*
[ORIGIN:http://as2h9t.live1-http.z1.pili.qiniucdn.com/coding/55d826b0d409d209d93f6e37.flv]
*/
```

#### Get Stream status

```go
streamStatus, err := stream.Status()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Status:\n", streamStatus)
/*
 {114.81.254.172:36317 connected 16870.200000000001 {42.200000000000003 14.733333333333333 0.066666666666666666}}
*/
```

#### Get Stream segments

```go
options = pili.OptionalArguments{ // optional
    Start: 1440379800, // optional, in second, unix timestamp
    End:   1440479880, // optional, in second, unix timestamp
    Limit: 20,         // optional, uint
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

#### Generate HLS playback URLs

```go
start := 1440379847
end := 1440379857
urls, err = stream.HlsPlaybackUrls(int64(start), int64(end))
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HlsPlaybackUrls:", urls)
/*
map[ORIGIN:http://as2h9t.playback1.z1.pili.qiniucdn.com/coding/55d826b0d409d209d93f6e37.m3u8?start=1440379847&end=1440379857]
*/
```

#### Save Stream as a file

```go
name := "fileName.mp4" // required, string
format := "mp4"        // required, string
start = 1440379847     // required, int64, in second, unix timestamp
end = 1440379857       // required, int64, in second, unix timestamp
options = pili.OptionalArguments{
    NotifyUrl: "http://remote_callback_url",
} // optional
saveAsRes, err := stream.SaveAs(name, format, int64(start), int64(end), options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream save as:\n", saveAsRes)
/*
{http://as2h9t.vod1.z1.pili.qiniucdn.com/recordings/z1.coding.55d826b0d409d209d93f6e37/fileName.m3u8 http://as2h9t.vod1.z1.pili.qiniucdn.com/recordings/z1.coding.55d826b0d409d209d93f6e37/fileName.mp4 z1.55da7715f51b82403b01e985}
*/
```

While invoking `saveAs()` and `snapshot()`, you can get processing state via Qiniu FOP Service using `persistentId`.  
API: `curl -D GET http://api.qiniu.com/status/get/prefop?id={PersistentId}`  
Doc reference: <http://developer.qiniu.com/docs/v6/api/overview/fop/persistent-fop.html#pfop-status>  

#### Snapshot Stream

```go
name = "fileName.jpg" // required, string
format = "jpg"        // required, string
options = pili.OptionalArguments{
    Time:      1440379847, // optional, int64, in second, unit timestamp
    NotifyUrl: "http://remote_callback_url",
} // optional
snapshotRes, err := stream.Snapshot(name, format, options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Snapshot:\n", snapshotRes)
/*
{http://as2h9t.static1.z1.pili.qiniucdn.com/snapshots/z1.coding.55d826b0d409d209d93f6e37/fileName.jpg z1.55da7716f51b82403b01e986}
*/
```

#### Delete a Stream

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
