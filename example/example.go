package main

import (
	"../pili" // or "github.com/pili-io/pili-sdk-go/pili"
	"fmt"
	"time"
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

func main() {

	// Instantiate an Pili client
	creds := pili.Creds(ACCESS_KEY, SECRET_KEY)
	client := pili.NewClient(creds)

	// Create a new stream
	hub := HUB            // required, hub must be an exists one
	title := ""           // optional, default is auto-generated
	publishKey := ""      // optional, a secret key for signing the <publishToken>, default is auto-generated
	publishSecurity := "" // optional, can be "dynamic" or "static", default is "dynamic"

	stream, err := client.CreateStream(hub, title, publishKey, publishSecurity)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", stream)

	// Get an exist stream
	stream, err = client.GetStream(stream.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", stream)

	// Get RTMP publish URL
	var nonce int64
	rtmpPublishUrl := stream.RtmpPublishUrl(RTMP_PUBLISH_HOST, nonce)
	fmt.Printf("RTMP Publish URL is:\n%+v\n\n", rtmpPublishUrl)

	// Get RTMP and HLS play URL
	profile := "" // optional, like '720p', '480p', '360p', '240p'. All profiles should be defined first.
	rtmpLiveUrl := stream.RtmpLiveUrl(RTMP_PLAY_HOST, profile)
	fmt.Printf("RTMP Play URL:\n%+v\n\n", rtmpLiveUrl)

	hlsLiveUrl := stream.HlsLiveUrl(HLS_PLAY_HOST, profile)
	fmt.Printf("HLS Play URL:\n%+v\n\n", hlsLiveUrl)

	startTime := time.Now().Unix() - 3600 // required
	endTime := time.Now().Unix()          // required
	hlsPlaybackUrl := stream.HlsPlaybackUrl(HLS_PLAY_HOST, profile, startTime, endTime)
	fmt.Printf("HLS Playback URL:\n%+v\n\n", hlsPlaybackUrl)

	// Get status
	streamStatus, err := client.GetStreamStatus(stream.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", streamStatus)

	// List streams
	hub = HUB       // required
	marker := ""    // optional
	var limit int64 // optional
	result1, err := client.ListStreams(hub, marker, limit)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result1)

	// Get recording segments from an exist stream
	// var startTime int64 // optional
	// var endTime int64   // optional
	result2, err := client.GetStreamSegments(stream.Id, startTime, endTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result2)

	// Update an exist stream
	newPublishKey := "new_secret_words"
	newPublishSecurity := "dynamic"
	disabled := true
	stream, err = client.SetStream(stream.Id, newPublishKey, newPublishSecurity, disabled)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", stream)

	// Delete
	result3, err := client.DelStream(stream.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result3)

}
