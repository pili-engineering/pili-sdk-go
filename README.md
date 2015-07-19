# Pili server-side library for Golang

## Features

- [x] Stream operations (Create, Delete, Update, Get)
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
    	- [Update a stream](#Update-a-stream)
		- [Delete a stream](#Delete-a-stream)
		- [Get stream segments](#Get-stream-segments)
		- [Get stream status](#Get-stream-status)
		- [Generate RTMP publish URL](#Generate-RTMP-publish-URL)
		- [Generate RTMP live play URL](#Generate-RTMP-live-play-URL)
		- [Generate HLS live play URL](#Generate-HLS-live-play-URL)
		- [Generate HLS playback URL](#Generate-HLS-playback-URL)
		- [To JSON String](#To-JSON-String)
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

    var credentials = pili.Creds(ACCESS_KEY, SECRET_KEY)
    var client = pili.NewClient(credentials, HUB)

    // ...
}
```

#### Create a stream

```go
options := pili.OptionalArguments {      // optional
	Title           : "test1",           // optional, default is auto-generated
	PublishKey      : "SomeSecretWords", // optional, a secret key for signing the `publishToken`, default is auto-generated
	PublishSecurity : "static",          // optional, can be "dynamic" or "static", "dynamic" as default
}

stream, err := client.CreateStream(&options)
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(stream)
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
stream, err = client.GetStream(streamId)
if err != nil {
	fmt.Println("error:", err)
}
fmt.Println(stream)
```

#### List streams

```go
options := pili.OptionalArguments { // optional
	Marker: "NextMarker",           // optional
	Limit: 100,                     // optional
}

result, err := client.listStreams(&options)
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(result)
// {Marker:NextMarker Items:[0x1042c180 0x1042c060]}

for _, stream := range s.Items {
	fmt.Println(stream)
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

#### Update a stream

```go
options := pili.OptionalArguments { // optional
  PublishKey     : "publishKey",    // optional
  PublishSecrity : "dynamic",       // optional
  Disabled       : false            // optional
}

stream, err := stream.Update(&options)
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(stream)
```

#### Delete a stream

```go
_, err := stream.Delete()
if err != nil {
	fmt.Println("error:", err)
}
```

#### Get stream segments

```go
options := pili.OptionalArguments { // optional
  Start : startTime,                // optional, in second, unix timestamp
  End   : endTime,                  // optional, in second, unix timestamp
}

segments, err := stream.Segments(&options)
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(segments)
// [{Start:1437283946 End:1437283999} {Start:1437284946 End:1437285946}]
```

#### Get stream status

```go
result, err := stream.Status()
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(result)
// {Addr:106.187.43.211:51393 Status:disconnected}
```

#### Generate RTMP publish URL

```go
url, _ := stream.RtmpPublishUrl()
fmt.Println(url)
// "rtmp://customized.example.com/hub/title?key=publishKey"
```

#### Generate RTMP live play URLs

```go
urls, err: = stream.RtmpLiveUrls()
if err != nil {
	fmt.Println("error:", err)
}

for k, v := range urls {
	fmt.Printf("%s:%s\n", k, v)
}

fmt.Println(urls["ORIGIN"])
// "rtmp://customized.example.com/hub/title"
```

#### Generate HLS live play URLs

```go
urls, err: = stream.HlsLiveUrls()
if err != nil {
	fmt.Println("error:", err)
}

for k, v := range urls {
	fmt.Printf("%s:%s\n", k, v)
}

fmt.Println(urls["ORIGIN"])
// "http://customized.example.com/hub/title.m3u8"
```

#### Generate HLS playback URLs

```go
// startTime: required, in second, unix timestamp
// endTime  : required, in second, unix timestamp

urls, err: = stream.HlsPlaybackUrls(startTime, endTime)
if err != nil {
	fmt.Println("error:", err)
}

for k, v := range urls {
	fmt.Printf("%s:%s\n", k, v)
}

fmt.Println(urls["ORIGIN"])
// "http://customized.example.com/hub/title.m3u8?start=1437283946&end=1437285946"
```

#### To JSON String
```go
streamJson, err := stream.ToJSONString()
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(streamJson)
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

## History

- 1.2.0
    - Update Stream object
    - Add new Stream functions
    - Update Client functions
