package pili

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStream(t *testing.T) {
	if skipTest() {
		log.Println("WARN: skip TestStream")
		return
	}
	client := New(testAccessKey, testSecretKey, nil)
	hub := NewHub(testHub, client)

	// Create.
	key := testStreamPrefix + "TestStream"
	_, err := hub.Create(key)
	require.NoError(t, err)
	stream, err := hub.Get(key)
	require.NoError(t, err)
	require.True(t, checkStream(stream, testHub, key, false))

	// Disable.
	err = stream.Disable()
	require.NoError(t, err)
	stream, err = hub.Get(key)
	require.NoError(t, err)
	require.True(t, checkStream(stream, testHub, key, true))

	// Enable.
	err = stream.Enable()
	require.NoError(t, err)
	stream, err = hub.Get(key)
	require.NoError(t, err)
	require.True(t, checkStream(stream, testHub, key, false))

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

	// HistoryRecord, empty.
	records, err := stream.HistoryRecord(0, 0)
	require.NoError(t, err)
	require.True(t, len(records) == 0)
}
