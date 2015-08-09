# Pili server-side library for Golang

## Features

- [x] Stream operations (Create, Get, Update, Disable, Enable, Refresh, Delete)
- [x] Get Streams list
- [x] Get Stream status
- [x] Get Stream segments
- [x] Generate RTMP publish URL
- [x] Generate RTMP and HLS live play URL
- [x] Generate HLS playback URL

## Content

- [Installation](#Installation)
- [Usage](#Usage)
	- [Configuration](#Configuration)
	- [Client](#Client)
		- [Create a Pili client](#Create-a-Pili-client)
		- [Create a stream](#Create-a-stream)
		- [Get a stream](#Get-a-stream)
		- [List streams](#List-streams)
	- [Stream](#Stream)
		- [To JSON String](#To-JSON-String)
    	- [Update a stream](#Update-a-stream)
		- [Refresh a stream](#Refresh-a-stream)
		- [Disable a stream](#Disable-a-stream)
		- [Enable a stream](#Enable-a-stream)
		- [Get stream segments](#Get-stream-segments)
		- [Get stream status](#Get-stream-status)
		- [Generate RTMP publish URL](#Generate-RTMP-publish-URL)
		- [Generate RTMP live play URL](#Generate-RTMP-live-play-URL)
		- [Generate HLS live play URL](#Generate-HLS-live-play-URL)
		- [Generate HLS playback URL](#Generate-HLS-playback-URL)
		- [Delete a stream](#Delete-a-stream)
- [History](#History)

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
	HUB        = "Pili_HubName"
)
```

### Client

#### Create a Pili client

```go

func main() {

    credentials := pili.Creds(ACCESS_KEY, SECRET_KEY)
    client := pili.NewClient(credentials, HUB)

    // ...
}
```

#### Create a stream

```go
options := pili.OptionalArguments{        // optional
    Title:           "stream_name",       // optional, auto-generated as default
    PublishKey:      "some_secret_words", // optional, auto-generated as default
    PublishSecurity: "static",            // optional, can be "dynamic" or "static", "dynamic" as default
}
stream, err := client.CreateStream(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("CreateStream:\n", stream)
/*
{
	ID:z1.live.5544ee03fb16df2e330001c5
	CreatedAt:2015-05-02 23:32:19.608 +0800 +0800
	UpdatedAt:2015-05-02 23:32:19.608 +0800 +0800
	Title:5544ee03fb16df2e330001c5
	Hub:live
	Disabled:false
	PublishKey:2769c4753656d244
	PublishSecurity:dynamic
	Profiles:[720p 480p 360p 240p]
	Hosts:{
		Publish:map[
			rtmp:5icsm3.pub.z1.pili.qiniup.com
		]
		Play:map[
			hls:5icsm3.hls1.z1.pili.qiniucdn.com
			rtmp:5icsm3.live1.z1.pili.qiniucdn.com
		]
	}
}
*/
```

#### Get a stream

```go
stream, err = client.GetStream(stream.Id)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("GetStream:\n", stream)
```

#### List streams

```go
options = pili.OptionalArguments{ // optional
    Marker: "",                   // optional, returned by server response
    Limit:  10,                   // optional
}
listResult, err := client.ListStreams(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("ListStreams:\n", listResult)
for _, stream := range listResult.Items {
    fmt.Println("Stream:\n", stream)
}
/*
&{
	ID:z1.live.5544ee03fb16df2e330001c5
	CreatedAt:2015-05-02 23:32:19.608 +0800 +0800
	UpdatedAt:2015-05-02 23:32:19.608 +0800 +0800
	Title:5544ee03fb16df2e330001c5
	Hub:live
	Disabled:false
	PublishKey:2769c4753656d244
	PublishSecurity:dynamic
	Profiles:[720p 480p 360p 240p]
	Hosts:{
		Publish:map[
			rtmp:5icsm3.pub.z1.pili.qiniup.com
		]
		Play:map[
			hls:5icsm3.hls1.z1.pili.qiniucdn.com
			rtmp:5icsm3.live1.z1.pili.qiniucdn.com
		]
	}
}
...
*/
```

### Stream

#### To JSON String
```go
streamJson, err := stream.ToJsonString()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream ToJsonString:\n", streamJson)
/*
{
    "id":"z1.live.5544ee03fb16df2e330001c5",
    "createdAt":"2015-05-02T23:32:19.608+08:00",
    "updatedAt":"2015-05-02T23:32:19.608+08:00",
    "title":"5544ee03fb16df2e330001c5",
    "hub":"live",
    "disabled":false,
    "publishKey":"2769c4753656d244",
    "publishSecurity":"dynamic",
    "profiles": ["720p", "480p", "360p", "240p"],
    "hosts":{
        "publish":{
            "rtmp":"5icsm3.pub.z1.pili.qiniup.com"
        },
        "play":{
            "hls":"5icsm3.hls1.z1.pili.qiniucdn.com",
            "rtmp":"5icsm3.live1.z1.pili.qiniucdn.com"
        }
    }
}
*/
```

#### Update a stream

```go
options = pili.OptionalArguments{  // optional
    PublishKey:      "publishKey", // optional
    PublishSecurity: "dynamic",    // optional
}
stream, err = stream.Update(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Updated:\n", stream)
```

#### Refresh a stream

```go
stream, err = stream.Refresh()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Refreshed:\n", stream)
```

#### Disable a stream

```go
stream, err = stream.Disable()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Disabled:\n", stream)
```

#### Enable a stream

```go
stream, err = stream.Enable()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Enabled:\n", stream)
```

#### Get stream segments

```go
options = pili.OptionalArguments{ // optional
    Start: 1439121809,            // optional, in second, unix timestamp
    End:   1439125409,            // optional, in second, unix timestamp
}
segments, err := stream.Segments(options)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Segments:\n", segments)
// [{Start:1437283946 End:1437283999} {Start:1437284946 End:1437285946}]
```

#### Get stream status

```go
streamStatus, err := stream.Status()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Status:\n", streamStatus)
// {Addr:106.187.43.211:51393 Status:disconnected}
```

#### Generate RTMP publish URL

```go
url := stream.RtmpPublishUrl()
fmt.Println("Stream RtmpPublishUrl:\n", url)
// "rtmp://customized.example.com/hub/title?key=publishKey"
```

#### Generate RTMP live play URLs

```go
urls, err := stream.RtmpLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("RtmpLiveUrls:")
for k, v := range urls {
    fmt.Printf("%s:%s\n", k, v)
}
fmt.Println("Original RtmpLiveUrl:\n", urls[pili.ORIGIN])
// "rtmp://customized.example.com/hub/title"
```

#### Generate HLS live play URLs

```go
urls, err = stream.HlsLiveUrls()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HlsLiveUrls:")
for k, v := range urls {
    fmt.Printf("%s:%s\n", k, v)
}
fmt.Println("Original HlsLiveUrl:\n", urls[pili.ORIGIN])
// "http://customized.example.com/hub/title.m3u8"
```

#### Generate HLS playback URLs

```go
start := 1439121809 // required, in second, unix timestamp
end   := 1439125409 // required, in second, unix timestamp
urls, err = stream.HlsPlaybackUrls(int64(start), int64(end))
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("HlsPlaybackUrls:")
for k, v := range urls {
    fmt.Printf("%s:%s\n", k, v)
}
fmt.Println("Original HlsPlaybackUrl:\n", urls[pili.ORIGIN])
// "http://customized.example.com/hub/title.m3u8?start=1439121809&end=1439125409"
```

#### Delete a stream

```go
deleteResult, err := stream.Delete()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println("Stream Deleted:\n", deleteResult)
// <nil>
```

## History

- 1.2.0
    - Update Stream object
    - Add new Stream functions
    - Update Client functions
