package pili

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkStream(stream *Stream, hub, key string, disabled bool) bool {
	return stream.Hub == hub && stream.Key == key && (stream.DisabledTill != 0) == disabled
}

func TestHub(t *testing.T) {

	if skipTest() {
		t.SkipNow()
	}
	client := New(&MAC{testAccessKey, []byte(testSecretKey)}, nil)
	hub := client.Hub(testHub)
	prefix := testStreamPrefix + "TestHub"

	// Create keyA, success.
	keyA := prefix + "A"
	streamA, err := hub.Create(keyA)
	require.NoError(t, err)
	require.True(t, checkStream(streamA, testHub, keyA, false))

	// Get keyA, success.
	streamA, err = hub.Get(keyA)
	require.NoError(t, err)
	require.True(t, checkStream(streamA, testHub, keyA, false))

	// Create keyA, exists.
	streamA, err = hub.Create(keyA)
	require.True(t, IsExists(err))

	// Get keyB, not exists.
	keyB := prefix + "B"
	_, err = hub.Get(keyB)
	require.True(t, IsNotExists(err))

	// Create keyB, success.
	streamB, err := hub.Create(keyB)
	require.NoError(t, err)
	require.True(t, checkStream(streamB, testHub, keyB, false))

	// List all.
	keys, marker, err := hub.List(prefix, 0, "")
	require.NoError(t, err)
	sort.Strings(keys)
	require.Equal(t, keys, []string{keyA, keyB})
	if marker != "" {
		keys, marker, err = hub.List(prefix, 0, marker)
		require.NoError(t, err)
		require.True(t, len(keys) == 0 && marker == "", "keys=%v marker=%v", keys, marker)
	}

	// List one by one.
	keys0, marker, err := hub.List(prefix, 1, "")
	require.NoError(t, err)
	require.True(t, len(keys0) == 1 && (keys0[0] == keyA || keys0[0] == keyB))
	require.NotEqual(t, marker, "")
	keys1, marker, err := hub.List(prefix, 1, marker)
	require.NoError(t, err)
	require.True(t, len(keys1) == 1 && keys1[0] != keys0[0] && (keys1[0] == keyA || keys1[0] == keyB), "got keys1=%v keys0=%v", keys1, keys0)
	require.NotEqual(t, marker, "")
	keys, marker, err = hub.List(prefix, 1, marker)
	require.NoError(t, err)
	require.True(t, len(keys) == 0 && marker == "")

	// ListLive all.
	keys, marker, err = hub.ListLive(prefix, 0, "")
	require.NoError(t, err)
	require.Equal(t, keys, []string{})
	require.Equal(t, marker, "")
}
