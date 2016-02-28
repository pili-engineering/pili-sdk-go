package pili

import (
	"fmt"
	"runtime"
)

const (
	SDK_VERSION         = "1.5.2"
	SDK_USER_AGENT      = "pili-sdk-go"
	DEFAULT_API_VERSION = "v1"
	DEFAULT_API_HOST    = "pili.qiniuapi.com"
	ORIGIN              = "ORIGIN"
)

var (
	API_HOST  string
	USE_HTTPS bool
)

func UserAgent() string {
	return fmt.Sprintf("%s/%s %s %s/%s", SDK_USER_AGENT, SDK_VERSION, runtime.Version(), runtime.GOOS, runtime.GOARCH)
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

func getApiBaseUrl() (url string) {
	return fmt.Sprintf("%s://%s/%s", getHttpScheme(), getApiHost(), DEFAULT_API_VERSION)
}
