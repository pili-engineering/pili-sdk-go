package pili

import (
	"fmt"
	"runtime"
)

const (
	SDK_VERSION      = "1.0.2"
	API_VERSION      = "v1"
	DEFAULT_API_HOST = "pili.qiniuapi.com"
)

var (
	API_HOST  string
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
