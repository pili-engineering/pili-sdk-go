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
    "fmt"
)


// Instantiate an client
var mac = pili.Mac{AccessKey, SecresultKey}
var app = pili.NewClient(&mac)


// Create a new stream
stream, err := app.CreateStream(nil)

/*
 *  or create a new stream with your custom arguments
 *
 *  key: default is auto generated
 *  is_private: default is false
 *  comment: default is blank

    postdata := map[string]interface{}{
        "key":        "8a7e79a8-dbb6-492a-a159-c291138cb461", // secret key like password for protected streaming
        "is_private": false,
        "comment":    "test_streaming_001",
    }
    stream, err := app.CreateStream(postdata)
*/

if err != nil {
    panic(err)
}

fmt.Printf("Result:%+v\n", stream)
fmt.Println("Stream Id:", stream.Id) // This is the only thing should write to the database
fmt.Println("Stream Key:", stream.Key)
fmt.Println("Stream is privately:", stream.IsPrivate)
fmt.Println("Stream push URL:", stream.PushUrl[0].RTMP)
fmt.Println("Stream RTMP live play URL:", stream.LiveUrl.RTMP)
fmt.Println("Stream HLS live play URL:", stream.LiveUrl.HLS)

sid := stream.Id

// Get an exist stream
stream, err := app.GetStream(sid)
fmt.Printf("Result:%+v\n", stream)

// Update a stream
stream, err := app.SetStream(sid, postdata)
fmt.Printf("Result:%+v\n", stream)

// Get Status on a stream
result, err := app.GetStreamStatus(sid)
fmt.Printf("Result:%+v\n", result)

// List exist streams
result, err := app.ListStreams()
fmt.Printf("Result:%+v\n", result)

// Delete a stream
stream, err := app.DelStream(sid)
fmt.Printf("Result:%+v\n", stream)

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
