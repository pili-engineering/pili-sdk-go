package main

import (
    "fmt"
    "time"
    "github.com/pili-io/pili-sdk-go/pili"
)

const (
    ACCESS_KEY = ""
    SECRET_KEY = ""
)

func main() {

    // Instantiate an client
    var creds = pili.Creds(ACCESS_KEY, SECRET_KEY)
    var app = pili.NewClient(creds)

    // Create a new stream
    postdata := map[string]interface{}{
        "hub":             "gouhuo",    // requried, must be exists
        "title":           "gouhuo", // optional, default is auto-generated
        "publishKey":      "8e7a69c1",   // optional, a secret key for signing the <publishToken>
        "publishSecurity": "dynamic",    // optional, can be "dynamic" or "static", default is "dynamic"
    }
    stream, err := app.CreateStream(postdata)
    fmt.Printf("Result:%+v\n", stream)


    // Get an exist stream
    stream, err = app.GetStream(stream.Id)
    fmt.Printf("Result:%+v\n", stream)


    // Signing a publish url
    publish := pili.PublishPolicy{
        Nonce:                 time.Now().UnixNano(),  // optional, for "dynamic" only, default is: time.Now().UnixNano()
        StreamId:              stream.Id,              // required
        StreamPublishKey:      stream.PublishKey,      // required, a secret key for signing the <publishToken>
        StreamPublishSecurity: stream.PublishSecurity, // required, can be "dynamic" or "static"
    }
    fmt.Println("Publish URL is:", publish.Url())


    // Play url
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


    // Update a stream
    newdata := map[string]interface{}{
        "publishKey":      "8e7a69c2",
        "publishSecurity": "static",
    }
    stream, err = app.SetStream(stream.Id, newdata)
    fmt.Printf("Result:%+v\n", stream)


    // List exist streams
    options := map[string]interface{}{
       "marker": "nextMarker", // string, optional
       "limit" : limitCount,   // int64,  optional
    }
    streams, err := app.ListStreams(hubName, options)
    fmt.Printf("Result:%+v\n", streams)


    // Get recording segments from a stream
    options := map[string]int64{
       "start": startUnixTimeStamp, // int64, optional
       "end"  : endUnixTimeStamp,   // int64, optional
    }
    segments, err := app.GetStreamSegments(stream.Id, options)
    fmt.Printf("Result:%+v\n", segments)


    // Delete a stream
    result, err := app.DelStream(stream.Id)
    fmt.Printf("Result:%+v\n", result)

}
