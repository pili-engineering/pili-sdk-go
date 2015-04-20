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
postdata := map[string]interface{}{
    "title":           "streamName", // optional, default is auto-generated
    "hub":             "hubName",    // requried, must be exists
    "publishKey":      "8e7a69c1",   // required, a secret key for signing pubishToken
    "publishSecurity": "dynamic",    // required, "dynamic" or "static"
}
stream, err := app.CreateStream(postdata)

if err != nil {
    panic(err)
}


// Get an exist stream
stream, err = app.GetStream(stream.Id)

fmt.Printf("Result:%+v\n", stream)


// Signing a publish url, then send it to the publisher client.
publish := pili.PublishPolicy{
    BaseUrl: "rtmp://<rtmpPublishHost>/<hubName>/<streamName>",
    Key:     stream.PublishKey,
    Nonce:   time.Now().UnixNano(),
}
fmt.Println("Publish Token is:", publish.Token())
fmt.Println("Publish URL is:", publish.Url())


// Update a stream
newdata := map[string]interface{}{
	"publishKey":      "8e7a69c2",
	"publishSecurity": "static",
}
stream, err = app.SetStream(stream.Id, newdata)
if err != nil {
	panic(err)
}
fmt.Printf("Result:%+v\n", result)


// List exist streams
result, err := app.ListStreams(hubName, nextMaker, limitCount)
fmt.Printf("Result:%+v\n", result)


// Delete a stream
result, err := app.DelStream(stream.Id)
fmt.Printf("Result:%+v\n", result)


// Get recording segments from a stream
result, err := app.GetStreamSegments(stream.Id, startUnixTimeStamp, endUnixTimeStamp)
fmt.Printf("Result:%+v\n", result)

```
