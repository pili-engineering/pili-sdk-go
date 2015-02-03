# PILI SDK for Go

The PILI SDK for Go is a set of clients for PILI Services APIs, and is currently under development to implement full service coverage and other standard PILI SDK features.

## Installing

```bash
$ go get github.com/pili-io/pili-sdk-go/pili
```

## Using

```go

import (
    "github.com/pili-io/pili-sdk-go/pili"
    // ...
)


// Instantiate an client
var creds = pili.Creds(AccessKey, SecretKey)
var app = pili.NewClient(creds)


// Create a new stream
stream, err := app.CreateStream(nil)

/*
 *  or create a new stream with your custom arguments
 *
 *  key: default is auto generated
 *  is_private: default is false
 *  comment: default is blank

    postdata := map[string]interface{}{
        "key":        "stream_secret_key", // used for protected streaming
        "is_private": false,
        "comment":    "test_streaming_001",
    }
    stream, err := app.CreateStream(postdata)
*/

if err != nil {
    panic(err)
}


// Stream ID is useful, maybe we should storage it use later.
sid := stream.Id

// Get an exist stream
stream, err = app.GetStream(sid)

fmt.Printf("Result:%+v\n", stream)
fmt.Println("Stream Id:", stream.Id)
fmt.Println("Stream Key:", stream.Key)
fmt.Println("Stream is privately:", stream.IsPrivate)
fmt.Println("Stream push URL:", stream.PushUrl[0].RTMP)
fmt.Println("Stream RTMP live play URL:", stream.LiveUrl.RTMP)
fmt.Println("Stream HLS live play URL:", stream.LiveUrl.HLS)


// Signing a pushing url, then send it to the pusher client.
push := pili.PushPolicy{
    BaseUrl: stream.PushUrl[0].RTMP,
    Key:     stream.Key,
    Nonce:   time.Now().UnixNano(),
}
fmt.Println("Push Token is:", push.Token())
fmt.Println("Push URL is:", push.Url())


// If true === stream.IsPrivate, we need signing for play.
playrtmp := pili.PlayPolicy{
    BaseUrl: stream.LiveUrl.RTMP,
    Key:     stream.Key,
    Expiry:  time.Now().Unix() + 3600,
}
fmt.Println("RTMP play token is:", playrtmp.Token())
fmt.Println("RTMP play url is:", playrtmp.Url())

playhls := pili.PlayPolicy{
    BaseUrl: stream.LiveUrl.HLS,
    Key:     stream.Key,
    Expiry:  time.Now().Unix() + 3600,
}
fmt.Println("HLS play token is:", playhls.Token())
fmt.Println("HLS play url is:", playhls.Url())



// Update a stream
result, err := app.SetStream(sid, postdata)
fmt.Printf("Result:%+v\n", result)

// Get Status on a stream
result, err := app.GetStreamStatus(sid)
fmt.Printf("Result:%+v\n", result)

// List exist streams
result, err := app.ListStreams()
fmt.Printf("Result:%+v\n", result)

// Delete a stream
result, err := app.DelStream(sid)
fmt.Printf("Result:%+v\n", result)

// Get recording segments from a stream
result, err := app.GetStreamSegments(sid, starttime, endtime)
fmt.Printf("Result:%+v\n", result)

// Delete recording segments on a stream
result, err := app.DelStreamSegments(sid, starttime, endtime)
fmt.Printf("Result:%+v\n", result)

// Get the play url of those stream recording segments
result, err := app.PlayStreamSegments(sid, starttime, endtime)
fmt.Printf("Result:%+v\n", result)


```
