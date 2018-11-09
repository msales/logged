package logged_test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/msales/logged"
	"github.com/stretchr/testify/assert"
)

func TestBufferedStreamHandler(t *testing.T) {
	buf := &bytes.Buffer{}
	h := logged.BufferedStreamHandler(buf, 2000, time.Second, logged.LogfmtFormat())

	h.Log("some message", logged.Error, []interface{}{})
	h.(io.Closer).Close()

	assert.Equal(t, "lvl=eror msg=\"some message\"\n", buf.String())
}

func TestBufferedStreamHandler_SendsMessagesAfterFlushInterval(t *testing.T) {
	buf := &bytes.Buffer{}
	h := logged.BufferedStreamHandler(buf, 2000, time.Millisecond, logged.LogfmtFormat())
	defer h.(io.Closer).Close()

	h.Log("some message", logged.Error, []interface{}{})

	time.Sleep(2 * time.Millisecond)

	assert.Equal(t, "lvl=eror msg=\"some message\"\n", buf.String())
}

func TestBufferedStreamHandler_SendsMessagesAfterFlushBytes(t *testing.T) {
	buf := &bytes.Buffer{}
	h := logged.BufferedStreamHandler(buf, 40, time.Second, logged.LogfmtFormat())
	defer h.(io.Closer).Close()

	h.Log("some message", logged.Error, []interface{}{})
	h.Log("some message", logged.Error, []interface{}{})

	time.Sleep(time.Millisecond)

	assert.Equal(t, "lvl=eror msg=\"some message\"\nlvl=eror msg=\"some message\"\n", buf.String())
}

func TestBufferedStreamHandler_DoesntWriteAfterClose(t *testing.T) {
	buf := &bytes.Buffer{}
	h := logged.BufferedStreamHandler(buf, 40, time.Second, logged.LogfmtFormat())
	h.(io.Closer).Close()

	h.Log("some message", logged.Error, []interface{}{})

	assert.Equal(t, "", buf.String())
}

func TestStreamHandler(t *testing.T) {
	buf := &bytes.Buffer{}
	h := logged.StreamHandler(buf, logged.LogfmtFormat())

	h.Log("some message", logged.Error, []interface{}{})

	assert.Equal(t, "lvl=eror msg=\"some message\"\n", buf.String())
}

func TestLevelFilterHandler(t *testing.T) {
	count := 0
	testHandler := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {
		count++
	})
	h := logged.LevelFilterHandler(logged.Info, testHandler)

	h.Log("test", logged.Debug, []interface{}{})
	h.Log("test", logged.Info, []interface{}{})

	assert.Equal(t, 1, count)
}

func TestLevelFilterHandler_TriesToCallUnderlyingClose(t *testing.T) {
	testHandler := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {})
	h := logged.LevelFilterHandler(logged.Info, testHandler)
	ch := h.(io.Closer)

	ch.Close()
}

func TestLevelFilterHandler_CallsUnderlyingClose(t *testing.T) {
	testHandler := &CloseableHandler{}
	h := logged.LevelFilterHandler(logged.Info, testHandler)
	ch := h.(io.Closer)

	ch.Close()

	assert.True(t, testHandler.CloseCalled)
}

func TestDiscardHandler(t *testing.T) {
	h := logged.DiscardHandler()

	h.Log("test", logged.Crit, []interface{}{})
}
