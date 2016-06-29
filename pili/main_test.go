package pili

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	testAccessKey    = ""
	testSecretKey    = ""
	testHub          = ""
	testStreamPrefix = "sdktest" + strconv.FormatInt(time.Now().UnixNano(), 10)
)

// Local test environment
// func init() {
// 	testAccessKey = "7O7hf7Ld1RrC_fpZdFvU8aCgOPuhw2K4eapYOdII"
// 	testSecretKey = "6Rq7rMSUHHqOgo0DJjh15tHsGUBEH9QhWqqyj4ka"
// 	testHub = "PiliSDKTest"
// 	APIHost = "10.200.20.28:7778"
// }

func skipTest() bool {
	return testAccessKey == "" || testSecretKey == "" || testHub == ""
}

func TestURL(t *testing.T) {
	if skipTest() {
		t.SkipNow()
	}

	mac := &MAC{testAccessKey, []byte(testSecretKey)}
	expect := "rtmp://publish-rtmp.test.com/" + testHub + "/key?e="
	result := RTMPPublishURL("publish-rtmp.test.com", testHub, "key", mac, 10)
	require.True(t, strings.HasPrefix(result, expect))

	expect = "rtmp://live-rtmp.test.com/" + testHub + "/key"
	result = RTMPPlayURL("live-rtmp.test.com", testHub, "key")
	require.Equal(t, result, expect)

	expect = "http://live-hls.test.com/" + testHub + "/key.m3u8"
	result = HLSPlayURL("live-hls.test.com", testHub, "key")
	require.Equal(t, result, expect)

	expect = "http://live-hdl.test.com/" + testHub + "/key.flv"
	result = HDLPlayURL("live-hdl.test.com", testHub, "key")
	require.Equal(t, result, expect)

	expect = "http://live-snapshot.test.com/" + testHub + "/key.jpg"
	result = SnapshotPlayURL("live-snapshot.test.com", testHub, "key")
	require.Equal(t, result, expect)
}
