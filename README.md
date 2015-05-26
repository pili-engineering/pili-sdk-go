# Pili server-side library for Golang

## Installation

```bash
$ go get github.com/pili-io/pili-sdk-go/pili
```

## Usage

### Configuration

```go

import (
    "github.com/pili-io/pili-sdk-go/pili"
    // ...
)

const (
	// Replace with your customized domains
	RTMP_PUBLISH_HOST = "xxx.pub.z1.pili.qiniup.com"
	RTMP_PLAY_HOST    = "xxx.live1.z1.pili.qiniucdn.com"
	HLS_PLAY_HOST     = "xxx.hls1.z1.pili.qiniucdn.com"

	// Replace with your keys here
	ACCESS_KEY = "QiniuAccessKey"
	SECRET_KEY = "QiniuSecretKey"

	// Replace with your hub name
	HUB = "hubName"
)
```


### Instantiate an Pili client

```go

func main() {

    var creds = pili.Creds(ACCESS_KEY, SECRET_KEY)
    var client = pili.NewClient(creds)

    // ...
}
```


### Create a new stream

```go
hub             := HUB // required, hub must be an exists one
title           := ""  // optional, default is auto-generated
publishKey      := ""  // optional, a secret key for signing the <publishToken>, default is auto-generated
publishSecurity := ""  // optional, can be "dynamic" or "static", default is "dynamic"

stream, err := client.CreateStream(hub, title, publishKey, publishSecurity)
if err != nil {
    panic(err)
}

fmt.Printf("Result:%+v\n", stream)
```


### Get stream

```go
stream, err = client.GetStream(stream.Id)
if err != nil {
    panic(err)
}
fmt.Printf("Result:%+v\n", stream)
```

### Get RTMP publish URL

```go
var nonce int64
publishUrl := stream.RtmpPublishUrl(RTMP_PUBLISH_HOST, nonce)
fmt.Printf("RTMP Publish URL is:\n%+v\n\n", publishUrl)
```


### Get RTMP and HLS play URL

```go
// optional, like '720p', '480p', '360p', '240p'. All profiles should be defined first.
profile := "480p"

rtmpLiveUrl := stream.RtmpLiveUrl(RTMP_PLAY_HOST, profile)
fmt.Printf("RTMP Play URL:\n%+v\n\n", rtmpLiveUrl)

hlsLiveUrl := stream.HlsLiveUrl(HLS_PLAY_HOST, profile)
fmt.Printf("HLS Play URL:\n%+v\n\n", hlsLiveUrl)

startTime := time.Now().Unix() - 3600 // required
endTime := time.Now().Unix()          // required
hlsPlaybackUrl := stream.HlsPlaybackUrl(HLS_PLAY_HOST, profile, startTime, endTime)
fmt.Printf("HLS Playback URL:\n%+v\n\n", hlsPlaybackUrl)
```


### Get status

```go
streamStatus, err := client.GetStreamStatus(stream.Id)
if err != nil {
    panic(err)
}
fmt.Printf("Result:%+v\n", streamStatus)
```


### List streams

```go
hub    := HUB // required
marker := ""  // optional
limit  := 0   // optional

result, err := client.ListStreams(hub, marker, int64(limit))
if err != nil {
    panic(err)
}

fmt.Printf("Result:%+v\n", result)
```


### Get recording segments from an exist stream

```go
var startTime int64 // optional
var endTime int64   // optional

segments, err := client.GetStreamSegments(stream.Id, startTime, endTime)
if err != nil {
    panic(err)
}

fmt.Printf("Result:%+v\n", segments)
```


### Update an exist stream

```go
newPublishKey      := "new_secret_words"
newPublishSecurity := "dynamic"
disabled           := true

stream, err = client.SetStream(stream.Id, newPublishKey, newPublishSecurity, disabled)
if err != nil {
    panic(err)
}

fmt.Printf("Result:%+v\n", stream)
```


### Delete stream

```go
del, err := client.DelStream(stream.Id)
if err != nil {
    panic(err)
}
fmt.Printf("Result:%+v\n", del)
```
