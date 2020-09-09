package logged_test

import (
	"testing"

	"github.com/msales/logged"
	"github.com/stretchr/testify/assert"
)

func TestLevelFromString(t *testing.T) {
	tests := []struct {
		lvl       string
		want      logged.Level
		wantError bool
	}{
		{
			lvl:       "dbug",
			want:      logged.Debug,
			wantError: false,
		},
		{
			lvl:       "debug",
			want:      logged.Debug,
			wantError: false,
		},
		{
			lvl:       "info",
			want:      logged.Info,
			wantError: false,
		},
		{
			lvl:       "warn",
			want:      logged.Warn,
			wantError: false,
		},
		{
			lvl:       "eror",
			want:      logged.Error,
			wantError: false,
		},
		{
			lvl:       "error",
			want:      logged.Error,
			wantError: false,
		},
		{
			lvl:       "crit",
			want:      logged.Crit,
			wantError: false,
		},
		{
			lvl:       "unkn",
			want:      logged.Level(123),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.lvl, func(t *testing.T) {
			lvl, err := logged.LevelFromString(tt.lvl)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, lvl)
		})
	}
}

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

func TestFormatFromString(t *testing.T) {
	tests := []struct {
		format    string
		wantError bool
	}{
		{
			format:    "json",
			wantError: false,
		},
		{
			format:    "logfmt",
			wantError: false,
		},
		{
			format:    "unkn",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			format, err := logged.FormatFromString(tt.format)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.IsType(t, logged.FormatterFunc(nil), format)
		})
	}
}

func TestFormat_String(t *testing.T) {
	tests := []struct {
		format logged.Format
		want   string
	}{
		{
			format: logged.Json,
			want:   "json",
		},
		{
			format: logged.Logfmt,
			want:   "logfmt",
		},
		{
			format: logged.Format(123),
			want:   "unkn",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.format.String())
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
