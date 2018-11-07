package logged_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/msales/logged"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		lvl  logged.Level
		want string
	}{
		{
			lvl:  logged.Debug,
			want: "dbug",
		},
		{
			lvl:  logged.Info,
			want: "info",
		},
		{
			lvl:  logged.Warn,
			want: "warn",
		},
		{
			lvl:  logged.Error,
			want: "eror",
		},
		{
			lvl:  logged.Crit,
			want: "crit",
		},
		{
			lvl:  logged.Level(123),
			want: "unkn",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.lvl.String())
	}
}

func TestNew(t *testing.T) {
	h := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {})

	l := logged.New(h)

	assert.Implements(t, (*logged.Logger)(nil), l)
}

func TestLogger(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(l logged.Logger)
		wantMsg string
		wantLvl logged.Level
		wantCtx []interface{}
	}{
		{
			name:    "Debug",
			fn:      func(l logged.Logger) { l.Debug("debug", "level", "debug") },
			wantMsg: "debug",
			wantLvl: logged.Debug,
			wantCtx: []interface{}{"level", "debug"},
		},
		{
			name:    "Info",
			fn:      func(l logged.Logger) { l.Info("info", "level", "info") },
			wantMsg: "info",
			wantLvl: logged.Info,
			wantCtx: []interface{}{"level", "info"},
		},
		{
			name:    "Warn",
			fn:      func(l logged.Logger) { l.Warn("warn", "level", "warn") },
			wantMsg: "warn",
			wantLvl: logged.Warn,
			wantCtx: []interface{}{"level", "warn"},
		},
		{
			name:    "Error",
			fn:      func(l logged.Logger) { l.Error("error", "level", "error") },
			wantMsg: "error",
			wantLvl: logged.Error,
			wantCtx: []interface{}{"level", "error"},
		},
		{
			name:    "Crit",
			fn:      func(l logged.Logger) { l.Crit("critical", "level", "critical") },
			wantMsg: "critical",
			wantLvl: logged.Crit,
			wantCtx: []interface{}{"level", "critical"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var outMsg string
			var outLvl logged.Level
			var outCtx []interface{}

			h := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {
				outMsg = msg
				outLvl = lvl
				outCtx = ctx
			})
			l := logged.New(h)

			tt.fn(l)

			assert.Equal(t, tt.wantMsg, outMsg)
			assert.Equal(t, tt.wantLvl, outLvl)
			assert.Equal(t, tt.wantCtx, outCtx)
		})
	}
}

func TestLogger_MergesCtx(t *testing.T) {
	var out []interface{}
	h := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {
		out = ctx
	})
	l := logged.New(h, "a", "b")

	l.Debug("test", "c", "d")

	assert.Equal(t, []interface{}{"a", "b", "c", "d"}, out)
}

func TestLogger_NormalizesCtx(t *testing.T) {
	var out []interface{}
	h := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {
		out = ctx
	})
	l := logged.New(h)

	l.Debug("test", "a")

	assert.Len(t, out, 4)
	assert.Equal(t, nil, out[1])
}

func TestLogger_TriesToCallUnderlyingClose(t *testing.T) {
	h := logged.HandlerFunc(func(msg string, lvl logged.Level, ctx []interface{}) {})
	l := logged.New(h)

	l.Close()
}

func TestLogger_CallsUnderlyingClose(t *testing.T) {
	h := &CloseableHandler{}
	l := logged.New(h)

	l.Close()

	assert.True(t, h.CloseCalled)
}
