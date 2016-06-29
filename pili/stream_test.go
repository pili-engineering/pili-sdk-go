package pili

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStreamInfo(t *testing.T) {

	info := StreamInfo{"a", "b", 0}
	require.False(t, info.Disabled())
	require.Equal(t, info.String(), "{hub:a,key:b,disabled:false}")

	info = StreamInfo{"a", "b", -1}
	require.True(t, info.Disabled())
	require.Equal(t, info.String(), "{hub:a,key:b,disabled:true}")

	info = StreamInfo{"a", "b", 12345}
	require.False(t, info.Disabled())
	require.Equal(t, info.String(), "{hub:a,key:b,disabled:false}")

	info = StreamInfo{"a", "b", time.Now().Unix() + 1000}
	require.True(t, info.Disabled())
	require.Equal(t, info.String(), "{hub:a,key:b,disabled:true}")
}

func TestStream(t *testing.T) {
	if skipTest() {
		t.SkipNow()
	}
	mac := &MAC{testAccessKey, []byte(testSecretKey)}
	client := New(mac, nil)
	hub := client.Hub(testHub)

	// Create.
	key := testStreamPrefix + "TestStream"
	_, err := hub.Create(key)
	require.NoError(t, err)
	stream := hub.Stream(key)
	require.NoError(t, err)
	info, err := stream.Info()
	require.NoError(t, err)
	require.True(t, checkStream(info, testHub, key, false))

	// Disable.
	stream = hub.Stream(key)
	err = stream.Disable()
	require.NoError(t, err)
	info, err = stream.Info()
	require.NoError(t, err)
	require.True(t, checkStream(info, testHub, key, true))

	// Enable.
	err = stream.Enable()
	require.NoError(t, err)
	info, err = stream.Info()
	require.NoError(t, err)
	require.True(t, checkStream(info, testHub, key, false))

	// LiveStatus, no live.
	_, err = stream.LiveStatus()
	require.Error(t, err)
	e, ok := err.(*Error)
	require.True(t, ok && e.Code == 619)

	// Save, not found.
	_, err = stream.Save(0, 0)
	require.Error(t, err)
	e, ok = err.(*Error)
	require.True(t, ok && e.Code == 619)

	// HistoryActivity, empty.
	records, err := stream.HistoryActivity(0, 0)
	require.NoError(t, err)
	require.True(t, len(records) == 0)
}
