package main

import (
	"../pili" // or "github.com/pili-io/pili-sdk-go/pili"
	"fmt"
	"time"
)

const (
	ACCESS_KEY = "YOUR_ACCESS_KEY"
	SECRET_KEY = "YOUR_SECRET_KEY"
	HUB        = "YOUR_HUB_NAME"
)

func main() {

	// Instantiate an Pili client
	creds := pili.Creds(ACCESS_KEY, SECRET_KEY)
	client := pili.NewClient(creds)

	// Create a new stream
	hub := HUB            // required, <Hub> must be an exists one
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

	// Signing a RTMP publish URL
	publish := pili.PublishPolicy{
		Nonce:                 time.Now().UnixNano(),  // optional, for "dynamic" only, default is: time.Now().UnixNano()
		StreamId:              stream.Id,              // required
		StreamPublishKey:      stream.PublishKey,      // required, a secret key for signing the <publishToken>
		StreamPublishSecurity: stream.PublishSecurity, // required, can be "dynamic" or "static"
	}
	fmt.Printf("Publish URL is:\n%+v\n\n", publish.Url())

	// Generate Play URLs
	pili.RTMP_PLAY_HOST = "live.z1.glb.pili.qiniucdn.com"
	pili.HLS_PLAY_HOST = "hls1.z1.glb.pili.qiniuapi.com"

	play := pili.PlayPolicy{
		StreamId: stream.Id, // required
	}

	preset := "" // optional, just like '720p', '480p', '360p', '240p'. All presets should be defined first.
	fmt.Printf("RTMP Play URL:\n%+v\n\n", play.RtmpLiveUrl(preset))
	fmt.Printf("HLS Play URL:\n%+v\n\n", play.HlsLiveUrl(preset))
	fmt.Printf("HLS Playback URL:\n%+v\n\n", play.HlsPlaybackUrl(1429678551, 1429689551, preset))

	fmt.Printf("RTMP 720P Play URL:\n%+v\n\n", play.RtmpLiveUrl("720p"))
	fmt.Printf("HLS 480P Play URL:\n%+v\n\n", play.HlsLiveUrl("480p"))
	fmt.Printf("HLS 360P Playback URL:\n%+v\n\n", play.HlsPlaybackUrl(1429678551, 1429689551, "360p"))

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
	var startTime int64 // optional
	var endTime int64   // optional
	result2, err := client.GetStreamSegments(stream.Id, startTime, endTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result2)

	// Update an exist stream
	newPublishKey := "new_secret_words"
	newPublishSecurity := "dynamic"
	stream, err = client.SetStream(stream.Id, newPublishKey, newPublishSecurity)
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
