package main

import (
	"./pili" // or "github.com/pili-io/pili-sdk-go/pili"
	"fmt"
	"time"
)

const (
	ACCESS_KEY = "YOUR_QINIU_ACCESS_KEY"
	SECRET_KEY = "YOUR_QINIU_SECRET_KEY"
	HUB        = "HUB_NAME"
)

func main() {

	creds := pili.Creds(ACCESS_KEY, SECRET_KEY)
	app := pili.NewClient(creds)

	// Create
	postdata := map[string]interface{}{
		"hub":             HUB,            // required
		"title":           "stream_name",  // optional, default is auto-generated
		"publishKey":      "secret_words", // optional, a secret key for signing the <publishToken>
		"publishSecurity": "dynamic",      // optional, can be "dynamic" or "static", default is "dynamic"
	}
	stream, err := app.CreateStream(postdata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", stream)

	// Query
	stream, err = app.GetStream(stream.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", stream)

	// Signing a publish url
	publish := pili.PublishPolicy{
		Nonce:                 time.Now().UnixNano(),  // optional, for "dynamic" only, default is: time.Now().UnixNano()
		StreamId:              stream.Id,              // required
		StreamPublishKey:      stream.PublishKey,      // required, a secret key for signing the <publishToken>
		StreamPublishSecurity: stream.PublishSecurity, // required, can be "dynamic" or "static"
	}
	fmt.Println("Publish URL is:", publish.Url())

	// List streams
	/*
	   options := map[string]interface{}{
	      "marker": <nextMarker:string>,
	      "limit" : <limitCount:int64>,
	   }
	*/
	result1, err := app.ListStreams(HUB, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result1)

	// List stream segments
	/*
	   options := map[string]int64{}{
	       "start": <startTime:int64>,
	       "end"  : <endTime:int64>,
	   }
	*/
	result2, err := app.GetStreamSegments(stream.Id, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result2)

	// Update
	newdata := map[string]interface{}{
		"publishKey":      "secret_words",
		"publishSecurity": "dynamic",
	}
	stream, err = app.SetStream(stream.Id, newdata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", stream)

	// Delete
	result3, err := app.DelStream(stream.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:%+v\n", result3)

	// Play
	pili.RTMP_PLAY_HOST = "live.z1.glb.pili.qiniucdn.com"
	pili.HLS_PLAY_HOST = "hls1.z1.glb.pili.qiniuapi.com"

	play := pili.PlayPolicy{
		StreamId: stream.Id, // required
	}
	fmt.Printf("RTMP Play URL:%+v\n", play.RtmpLiveUrl(""))
	fmt.Printf("HLS Play URL:%+v\n", play.HlsLiveUrl(""))
	fmt.Printf("HLS Playback URL:%+v\n", play.HlsPlaybackUrl(1429678551, 1429689551, ""))

	fmt.Printf("RTMP 720P Play URL:%+v\n", play.RtmpLiveUrl("720p"))
	fmt.Printf("HLS 480P Play URL:%+v\n", play.HlsLiveUrl("480p"))
	fmt.Printf("HLS 360P Playback URL:%+v\n", play.HlsPlaybackUrl(1429678551, 1429689551, "360p"))
}
