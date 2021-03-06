package logged_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/msales/logged"
	"github.com/stretchr/testify/assert"
)

func TestJsonFormat(t *testing.T) {
	f := logged.JSONFormat()

	b := f.Format("some message", logged.Error, []interface{}{"x", 1, "y", 3.2, "bool", true,
		"carriage_return", "bang" + string('\r') + "foo", "tab", "bar	baz", "newline", "foo\nbar", "escape", string('\\')})

	expect := []byte(`{"lvl":"eror","msg":"some message","x":1,"y":3.2,"bool":true,"carriage_return":"bang\rfoo","tab":"bar\tbaz","newline":"foo\nbar","escape":"\\"}` + "\n")
	assert.Equal(t, expect, b)

	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	assert.NoError(t, err)
}

func TestJsonFormat_KeyError(t *testing.T) {
	f := logged.JSONFormat()

	b := f.Format("some message", logged.Error, []interface{}{1, "y"})

	expect := []byte(`{"lvl":"eror","msg":"some message","LOGGED_ERROR":1}` + "\n")
	assert.Equal(t, expect, b)

	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	assert.NoError(t, err)
}

func TestJsonFormat_Ints(t *testing.T) {
	f := logged.JSONFormat()

	b := f.Format("", logged.Error, []interface{}{"int", 1, "int8", int8(2), "int16", int16(3), "int32", int32(4), "int64", int64(5)})

	expect := []byte(`{"lvl":"eror","msg":"","int":1,"int8":2,"int16":3,"int32":4,"int64":5}` + "\n")
	assert.Equal(t, expect, b)

	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	assert.NoError(t, err)
}

func TestJsonFormat_Uints(t *testing.T) {
	f := logged.JSONFormat()

	b := f.Format("", logged.Error, []interface{}{"uint", uint(1), "uint8", uint8(2), "uint16", uint16(3), "uint32", uint32(4), "uint64", uint64(5)})

	expect := []byte(`{"lvl":"eror","msg":"","uint":1,"uint8":2,"uint16":3,"uint32":4,"uint64":5}` + "\n")
	assert.Equal(t, expect, b)

	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	assert.NoError(t, err)
}

func TestJsonFormat_Floats(t *testing.T) {
	f := logged.JSONFormat()

	b := f.Format("", logged.Error, []interface{}{"float32", float32(1), "float64", float64(4.56)})

	expect := []byte(`{"lvl":"eror","msg":"","float32":1,"float64":4.56}` + "\n")
	assert.Equal(t, expect, b)

	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	assert.NoError(t, err)
}

func TestJsonFormat_Time(t *testing.T) {
	f := logged.JSONFormat()

	b := f.Format("", logged.Error, []interface{}{"time", time.Unix(1541573670, 0).UTC()})

	expect := []byte(`{"lvl":"eror","msg":"","time":"2018-11-07T06:54:30+0000"}` + "\n")
	assert.Equal(t, expect, b)

	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	assert.NoError(t, err)
}

func TestJsonFormat_Unknown(t *testing.T) {
	obj := struct {
		Name string
	}{Name: "test"}

	f := logged.JSONFormat()

	b := f.Format("", logged.Error, []interface{}{"what", obj, "nil", nil})

	expect := []byte(`{"lvl":"eror","msg":"","what":"{Name:test}","nil":null}` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat(t *testing.T) {
	f := logged.LogfmtFormat()

	b := f.Format("some message", logged.Error, []interface{}{"x", 1, "y", 3.2, "bool", true, "equals", "=", "quote", "\"",
		"carriage_return", "bang" + string('\r') + "foo", "tab", "bar	baz", "newline", "foo\nbar", "escape", string('\\')})

	expect := []byte(`lvl=eror msg="some message" x=1 y=3.200 bool=true equals="=" quote="\"" carriage_return="bang\rfoo" tab="bar\tbaz" newline="foo\nbar" escape=\\` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat_KeyError(t *testing.T) {
	f := logged.LogfmtFormat()

	b := f.Format("some message", logged.Error, []interface{}{1, "y"})

	expect := []byte(`lvl=eror msg="some message" LOGGED_ERROR=1` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat_Ints(t *testing.T) {
	f := logged.LogfmtFormat()

	b := f.Format("", logged.Error, []interface{}{"int", 1, "int8", int8(2), "int16", int16(3), "int32", int32(4), "int64", int64(5)})

	expect := []byte(`lvl=eror msg= int=1 int8=2 int16=3 int32=4 int64=5` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat_Uints(t *testing.T) {
	f := logged.LogfmtFormat()

	b := f.Format("", logged.Error, []interface{}{"uint", uint(1), "uint8", uint8(2), "uint16", uint16(3), "uint32", uint32(4), "uint64", uint64(5)})

	expect := []byte(`lvl=eror msg= uint=1 uint8=2 uint16=3 uint32=4 uint64=5` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat_Floats(t *testing.T) {
	f := logged.LogfmtFormat()

	b := f.Format("", logged.Error, []interface{}{"float32", float32(1.23), "float64", float64(4.56)})

	expect := []byte(`lvl=eror msg= float32=1.230 float64=4.560` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat_Time(t *testing.T) {
	f := logged.LogfmtFormat()

	b := f.Format("", logged.Error, []interface{}{"time", time.Unix(1541573670, 0).UTC()})

	expect := []byte(`lvl=eror msg= time=2018-11-07T06:54:30+0000` + "\n")
	assert.Equal(t, expect, b)
}

func TestLogfmtFormat_Unknown(t *testing.T) {
	obj := struct {
		Name string
	}{Name: "test"}

	f := logged.LogfmtFormat()

	b := f.Format("", logged.Error, []interface{}{"what", obj, "nil", nil})

	expect := []byte(`lvl=eror msg= what={Name:test} nil=` + "\n")
	assert.Equal(t, expect, b)
}
