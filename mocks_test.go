package logged_test

import "github.com/msales/logged"

type CloseableHandler struct {
	CloseCalled bool
}

func (h *CloseableHandler) Log(msg string, lvl logged.Level, ctx []interface{}) {}

func (h *CloseableHandler) Close() error {
	h.CloseCalled = true
	return nil
}
