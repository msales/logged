package logged

import (
	"strconv"
	"sync"
	"time"
	"unsafe"
)

// pool is a pool of Buffers.
type pool struct {
	p *sync.Pool
}

// newPool creates a new instance of pool.
func newPool(size int) pool {
	return pool{p: &sync.Pool{
		New: func() interface{} {
			return &buffer{b: make([]byte, 0, size)}
		},
	}}
}

// Get retrieves a buffer from the pool, creating one if necessary.
func (p pool) Get() *buffer {
	buf := p.p.Get().(*buffer)
	buf.Reset()
	return buf
}

// Put adds a buffer to the pool.
func (p pool) Put(buf *buffer) {
	p.p.Put(buf)
}

// buffer wraps a byte slice, providing continence functions.
type buffer struct {
	b []byte
}

// AppendInt appends an integer to the underlying buffer.
func (b *buffer) AppendInt(i int64) {
	b.b = strconv.AppendInt(b.b, i, 10)
}

// AppendUint appends an unsigned integer to the underlying buffer.
func (b *buffer) AppendUint(i uint64) {
	b.b = strconv.AppendUint(b.b, i, 10)
}

// AppendFloat appends a float to the underlying buffer.
func (b *buffer) AppendFloat(f float64, fmt byte, prec, bitSize int) {
	b.b = strconv.AppendFloat(b.b, f, fmt, prec, bitSize)
}

// AppendBool appends a bool to the underlying buffer.
func (b *buffer) AppendBool(v bool) {
	b.b = strconv.AppendBool(b.b, v)
}

// AppendTime appends a time to the underlying buffer, in the given layout.
func (b *buffer) AppendTime(t time.Time, layout string) {
	b.b = t.AppendFormat(b.b, layout)
}

// WriteByte writes a single byte to the buffer.
func (b *buffer) WriteByte(v byte) error {
	b.b = append(b.b, v)
	return nil
}

// WriteString writes a string to the buffer.
func (b *buffer) WriteString(s string) {
	b.b = append(b.b, s...)
}

// Write implements io.Writer.
func (b *buffer) Write(bs []byte) (int, error) {
	b.b = append(b.b, bs...)

	return len(bs), nil
}

// Len returns the length of the underlying byte slice.
func (b *buffer) Len() int {
	return len(b.b)
}

// Cap returns the capacity of the underlying byte slice.
func (b *buffer) Cap() int {
	return cap(b.b)
}

// Bytes returns a mutable reference to the underlying byte slice.
func (b *buffer) Bytes() []byte {
	return b.b
}

// String returns a string of the underlying byte slice.
func (b *buffer) String() string {
	return *(*string)(unsafe.Pointer(&b.b))
}

// Reset resets the underlying byte slice. Subsequent writes re-use the slice's
// backing array.
func (b *buffer) Reset() {
	b.b = b.b[:0]
}
