package main

import (
	"fmt"
	"pili-sdk-go.v1/pili" // or "github.com/pili-engineering/pili-sdk-go/pili"
	"time"
)

const (

	// Replace with your keys here
	ACCESS_KEY = "Qiniu_AccessKey"
	SECRET_KEY = "Qiniu_SecretKey"

	// The Hub must be exists before use
	HUB_NAME = "Pili_Hub_Name"
)

func createStream(hub pili.Hub, streamId string) {
	// Create a new stream
	options := pili.OptionalArguments{ // optional
		Title:           streamId,            // optional, auto-generated as default
		PublishKey:      "some_secret_words", // optional, auto-generated as default
		PublishSecurity: "dynamic",           // optional, can be "dynamic" or "static", "dynamic" as default
	}
	stream, err := hub.CreateStream(options)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("CreateStream:%s\n", stream)
}

func getStream(hub pili.Hub, streamId string) {
	// Get a stream
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("GetStream:%s\n", stream)
}

func listStream(hub pili.Hub) {
	// List streams
	options := pili.OptionalArguments{ // optional
	//Status: "connected", // optional
	//Marker: "",          // optional, returned by server response
	//Limit:  50,          // optional
	//Title:  "",          // optional, title prefix
	}
	listResult, err := hub.ListStreams(options)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("ListStreams:%s\n", listResult)
	for _, stream := range listResult.Items {
		fmt.Printf("Stream:%s\n", stream)
	}
}

func updateStream(hub pili.Hub, streamId string) {
	// Update a stream
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stream.PublishKey = "new_secret_words" // optional
	stream.PublishSecurity = "static"      // optional
	stream, err = stream.Update()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Stream Updated:%s\n", stream)
}

func disableStream(hub pili.Hub, streamId string) {
	// Disable a stream
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Println(err)
		return
	}
	stream, err = stream.Disable()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Stream Disabled:%s\n", stream)

	// Disable a stream with disableTill
	err = stream.DisableTill(time.Now().Add(time.Hour))
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
}

func enableStream(hub pili.Hub, streamId string) {
	// Enable a stream
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Println(err)
		return
	}

	stream, err = stream.Enable()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Stream Enabled:%s\n", stream)
}

func getPlayUrl(hub pili.Hub, streamId string) {
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Generate RTMP publish URL
	url := stream.RtmpPublishUrl()
	fmt.Printf("Stream RtmpPublishUrl:%s\n", url)

	// Generate RTMP live play URLs
	urls, err := stream.RtmpLiveUrls()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("RtmpLiveUrls:%s\n", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
		return
	}

	// Generate HLS live play URLs
	urls, err = stream.HlsLiveUrls()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("HlsLiveUrls:%s\n", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
		return
	}

	// Generate Http-Flv live play URLs
	urls, err = stream.HttpFlvLiveUrls()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("HttpFlvLiveUrls:%s\n", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}

}

func getStreamStatus(hub pili.Hub, streamId string) {
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get stream status
	streamStatus, err := stream.Status()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Stream Status:%s\n", streamStatus)
}

func getSegments(hub pili.Hub, streamId string) {
	stream, err := hub.GetStream(streamId)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get stream segments
	options := pili.OptionalArguments{ // optional
		Start: 1440379800, // optional, in second, unix timestamp
		End:   1440479880, // optional, in second, unix timestamp
		Limit: 20,         // optional, uint
	}
	segments, err := stream.Segments(options)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Segments:%s\n", segments)

}

func getPlaybackUrl(stream pili.Stream) {
	// Generate HLS playback URLs
	start := 1440379847
	end := 1440379857
	urls, err := stream.HlsPlaybackUrls(int64(start), int64(end))
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("HlsPlaybackUrls:%s\n", urls)
	for k, v := range urls {
		fmt.Printf("%s:%s\n", k, v)
	}
}

//保存回放
func saveAs(stream pili.Stream) {
	// Save Stream as a file
	name := "fileName.mp4" // required, string
	start := 1440379847    // required, int64, in second, unix timestamp
	end := 1440379857      // required, int64, in second, unix timestamp
	format := "mp4"        // optional, string
	options := pili.OptionalArguments{
		NotifyUrl: "http://remote_callback_url",
	} // optional
	saveAsRes, err := stream.SaveAs(name, format, int64(start), int64(end), options)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Stream save as:%s\n", saveAsRes)
}

//获取截图
func snapshot(stream pili.Stream) {
	// Snapshot Stream
	name := "fileName.jpg" // required, string
	format := "jpg"        // required, string
	options := pili.OptionalArguments{
		Time:      1440379847, // optional, int64, in second, unit timestamp
		NotifyUrl: "http://remote_callback_url",
	} // optional
	snapshotRes, err := stream.Snapshot(name, format, options)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	fmt.Printf("Stream Snapshot:%s\n", snapshotRes)

}

//删除流
func deleteStream(stream pili.Stream) {
	// Delete a stream
	deleteResult, err := stream.Delete()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Stream Deleted:%s\n", deleteResult)
}

func toJsonString(stream pili.Stream) {
	// To JSON String
	streamJson, err := stream.ToJSONString()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
	}
	fmt.Printf("Stream ToJSONString:%s\n", streamJson)
}

func createRoom(meeting pili.Meeting, ownerid, room string) {
	option := pili.RoomOptionArguments{
		OwnerId: ownerid,
		Name:    room,
		UserMax: 10,
	}
	roomInfo, err := meeting.CreateRoom(option)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(roomInfo.RoomName)
}

func getRoom(meeting pili.Meeting, roomName string) {
	room, err := meeting.GetRoom(roomName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", room)
}

func getActiveUser(meeting pili.Meeting, roomName string) (userid string) {
	allUsers, err := meeting.ActiveUsers(roomName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", allUsers)
	if len(allUsers.Users) > 1 {
		return allUsers.Users[0].UserId
	}
	return ""
}

func rejectUser(meeting pili.Meeting, roomName, userId string) {
	if userId == "" {
		return
	}
	err := meeting.RejectUser(roomName, userId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("reject user %s\n", userId)
}

func deleteRoom(meeting pili.Meeting, roomName string) {
	ret, err := meeting.DeleteRoom(roomName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", ret)
}

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
	meeting := pili.NewMeeting(credentials)

	streamKey := fmt.Sprintf("sdk_test_%d", time.Now().UnixNano())
	room := "sdk_room_test"

	fmt.Printf("streamKey = %s\n", streamKey)
	fmt.Printf("room = %s\n", room)

	//创建流
	fmt.Println("创建流：")
	createStream(hub, streamKey)

	//获取流
	fmt.Println("获取流：")
	streamKey = fmt.Sprintf("z1.%s.%s", HUB_NAME, streamKey)
	getStream(hub, streamKey)

	//获取流列表
	fmt.Println("获取流列表：")
	listStream(hub)

	//更新流
	fmt.Println("更新流：")
	updateStream(hub, streamKey)

	//禁播流
	fmt.Println("禁播流：")
	disableStream(hub, streamKey)

	//解禁流
	fmt.Println("解禁流：")
	enableStream(hub, streamKey)

	//获取直播播放链接
	fmt.Println("获取播放链接：")
	getPlayUrl(hub, streamKey)

	//查看流状态
	fmt.Println("查看流状态：")
	getStreamStatus(hub, streamKey)

	//获取流的 segments
	fmt.Println("获取 segments:")
	getSegments(hub, streamKey)

	//---------------------------- rtc -------------------------------------//

	//创建连麦房间
	fmt.Println("创建连麦房间：")
	ownerId := "ownerid"
	createRoom(meeting, ownerId, room)

	//获取连麦房间
	fmt.Println("获取连麦房间：")
	getRoom(meeting, room)

	//获取连麦房间活跃人数
	fmt.Println("获取连麦房间内的活跃人数：")
	userid := getActiveUser(meeting, room)

	//剔除连麦用户
	fmt.Println("剔除连麦用户：")
	rejectUser(meeting, room, userid)

	//签算roomtoken
	fmt.Println("签算 room token：")
	option := pili.RoomAccessPolicy{
		Room:     "roomName",
		User:     "123",
		Perm:     "admin",
		Version:  "2.0",
		ExpireAt: time.Now().Add(time.Hour).UnixNano(),
	}
	token, err := meeting.RoomToken(credentials, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

	//删除连麦房间
	fmt.Println("删除连麦房间：")
	deleteRoom(meeting, room)
}
