package pili

import (
	"fmt"
	"runtime"
)

const (
	SDK_VERSION               = "1.0.1"
	API_VERSION               = "v1"
	DEFAULT_API_HOST          = "pili.qiniuapi.com"
	DEFAULT_RTMP_PUBLISH_HOST = "pub.z1.glb.pili.qiniup.com"
	DEFAULT_RTMP_PLAY_HOST    = "live.z1.glb.pili.qiniucdn.com"
	DEFAULT_HLS_PLAY_HOST     = "hls.z1.glb.pili.qiniuapi.com"
)

var (
	API_HOST          string
	RTMP_PUBLISH_HOST string
	RTMP_PLAY_HOST    string
	HLS_PLAY_HOST     string

	USE_HTTPS bool
)

func UserAgent() string {
	return fmt.Sprintf("pili-sdk-go/%s %s %s/%s", SDK_VERSION, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func getHttpScheme() (scheme string) {
	scheme = "http"
	if USE_HTTPS {
		scheme = "https"
	}
	return
}

func getApiHost() (host string) {
	host = DEFAULT_API_HOST
	if API_HOST != "" {
		host = API_HOST
	}
	return
}

func getRtmpPublishHost() (host string) {
	host = DEFAULT_RTMP_PUBLISH_HOST
	if RTMP_PUBLISH_HOST != "" {
		host = RTMP_PUBLISH_HOST
	}
	return
}

func getRtmpPlayHost() (host string) {
	host = DEFAULT_RTMP_PLAY_HOST
	if RTMP_PLAY_HOST != "" {
		host = RTMP_PLAY_HOST
	}
	return
}

func getHlsPlayHost() (host string) {
	host = DEFAULT_HLS_PLAY_HOST
	if HLS_PLAY_HOST != "" {
		host = HLS_PLAY_HOST
	}
	return
}
