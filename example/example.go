package main

import (
	"../pili" // or "github.com/pili-io/pili-sdk-go/pili"
	"fmt"
)

const (

	// Replace with your keys here
	ACCESS_KEY = "Qiniu_AccessKey"
	SECRET_KEY = "Qiniu_SecretKey"

	// Replace with your hub name
	HUB = "Pili_Hub"
)

func main() {

	// Instantiate an Pili client
	credentials := pili.Creds(ACCESS_KEY, SECRET_KEY)
	client := pili.NewClient(credentials, HUB)

	// Create a new stream
	options := pili.OptionalArguments{ // optional
		Title:           "stream_name",       // optional, auto-generated as default
		PublishKey:      "some_secret_words", // optional, auto-generated as default
		PublishSecurity: "static",            // optional, can be "dynamic" or "static", "dynamic" as default
	}
	stream, err := client.CreateStream(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("CreateStream:\n", stream)

	// Get a stream
	stream, err = client.GetStream(stream.Id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("GetStream:\n", stream)

	// List streams
	options = pili.OptionalArguments{ // optional
		Marker: "", // optional, returned by server response
		Limit:  10, // optional
	}
	listResult, err := client.ListStreams(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("ListStreams:\n", listResult)
	for _, stream := range listResult.Items {
		fmt.Println("Stream:\n", stream)
	}

	// To JSON String
	streamJson, err := stream.ToJsonString()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream ToJsonString:\n", streamJson)

	// Update a stream
	options = pili.OptionalArguments{ // optional
		PublishKey:      "publishKey", // optional
		PublishSecurity: "dynamic",    // optional
	}
	stream, err = stream.Update(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Updated:\n", stream)

	// Refresh a stream
	stream, err = stream.Refresh()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Refreshed:\n", stream)

	// Disable a stream
	stream, err = stream.Disable()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Disabled:\n", stream)

	// Enable a stream
	stream, err = stream.Enable()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Enabled:\n", stream)

	// Get stream segments
	options = pili.OptionalArguments{ // optional
		Start: 1439121809, // optional, in second, unix timestamp
		End:   1439125409, // optional, in second, unix timestamp
	}
	segments, err := stream.Segments(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Segments:\n", segments)

	// Get stream status
	streamStatus, err := stream.Status()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Status:\n", streamStatus)

	// Generate RTMP publish URL
	url := stream.RtmpPublishUrl()
	fmt.Println("Stream RtmpPublishUrl:\n", url)

	// Generate RTMP live play URLs
	urls, err := stream.RtmpLiveUrls()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("RtmpLiveUrls:")
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}
	fmt.Println("Original RtmpLiveUrl:\n", urls[pili.ORIGIN])

	// Generate HLS live play URLs
	urls, err = stream.HlsLiveUrls()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("HlsLiveUrls:")
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}
	fmt.Println("Original HlsLiveUrl:\n", urls[pili.ORIGIN])

	// Generate HLS playback URLs
	start := 1439121809
	end := 1439125409
	urls, err = stream.HlsPlaybackUrls(int64(start), int64(end))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("HlsPlaybackUrls:")
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}
	fmt.Println("Original HlsPlaybackUrl:\n", urls[pili.ORIGIN])

	// Delete a stream
	deleteResult, err := stream.Delete()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Deleted:\n", deleteResult)
}
