package main

import (
	"../pili" // or "github.com/pili-engineering/pili-sdk-go/pili"
	"fmt"
)

const (

	// Replace with your keys here
	ACCESS_KEY = "Qiniu_AccessKey"
	SECRET_KEY = "Qiniu_SecretKey"

	// The Hub must be exists before use
	HUB_NAME = "Pili_Hub_Name"
)

func main() {

	// Change API host as necessary
	//
	// pili.qiniuapi.com as default
	// pili-lte.qiniuapi.com is the latest RC version
	//
	// pili.API_HOST = "pili.qiniuapi.com" // default

	// Instantiate a Pili Hub object
	credentials := pili.NewCredentials(ACCESS_KEY, SECRET_KEY)
	hub := pili.NewHub(credentials, HUB_NAME)

	// Create a new stream
	options := pili.OptionalArguments{ // optional
		Title:           "stream_name",       // optional, auto-generated as default
		PublishKey:      "some_secret_words", // optional, auto-generated as default
		PublishSecurity: "dynamic",           // optional, can be "dynamic" or "static", "dynamic" as default
	}
	stream, err := hub.CreateStream(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("CreateStream:\n", stream)

	// Get a stream
	stream, err = hub.GetStream(stream.Id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("GetStream:\n", stream)

	// List streams
	options = pili.OptionalArguments{ // optional
	//Status: "connected", // optional
	//Marker: "",          // optional, returned by server response
	//Limit:  50,          // optional
	//Title:  "",          // optional, title prefix
	}
	listResult, err := hub.ListStreams(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("ListStreams:\n", listResult)
	for _, stream := range listResult.Items {
		fmt.Println("Stream:\n", stream)
	}

	// To JSON String
	streamJson, err := stream.ToJSONString()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream ToJSONString:\n", streamJson)

	// Update a stream
	stream.PublishKey = "new_secret_words" // optional
	stream.PublishSecurity = "static"      // optional
	stream, err = stream.Update()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Updated:\n", stream)

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

	// Generate RTMP publish URL
	url := stream.RtmpPublishUrl()
	fmt.Println("Stream RtmpPublishUrl:\n", url)

	// Generate RTMP live play URLs
	urls, err := stream.RtmpLiveUrls()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("RtmpLiveUrls:", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}

	// Generate HLS live play URLs
	urls, err = stream.HlsLiveUrls()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("HlsLiveUrls:", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}

	// Generate Http-Flv live play URLs
	urls, err = stream.HttpFlvLiveUrls()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("HttpFlvLiveUrls:", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}

	// Get stream status
	streamStatus, err := stream.Status()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Status:\n", streamStatus)

	// Get stream segments
	options = pili.OptionalArguments{ // optional
		Start: 1440379800, // optional, in second, unix timestamp
		End:   1440479880, // optional, in second, unix timestamp
		Limit: 20,         // optional, uint
	}
	segments, err := stream.Segments(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Segments:\n", segments)

	// Generate HLS playback URLs
	start := 1440379847
	end := 1440379857
	urls, err = stream.HlsPlaybackUrls(int64(start), int64(end))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("HlsPlaybackUrls:", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}

	// Save Stream as a file
	name := "fileName.mp4" // required, string
	start = 1440379847     // required, int64, in second, unix timestamp
	end = 1440379857       // required, int64, in second, unix timestamp
	format := "mp4"        // optional, string
	options = pili.OptionalArguments{
		NotifyUrl: "http://remote_callback_url",
	} // optional
	saveAsRes, err := stream.SaveAs(name, format, int64(start), int64(end), options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream save as:\n", saveAsRes)

	// Snapshot Stream
	name = "fileName.jpg" // required, string
	format = "jpg"        // required, string
	options = pili.OptionalArguments{
		Time:      1440379847, // optional, int64, in second, unit timestamp
		NotifyUrl: "http://remote_callback_url",
	} // optional
	snapshotRes, err := stream.Snapshot(name, format, options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Snapshot:\n", snapshotRes)

	// Delete a stream
	deleteResult, err := stream.Delete()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stream Deleted:\n", deleteResult)
}
