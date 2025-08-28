package bytes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/slices/bytes"
)

func TestBytes(t *testing.T) {
	b := bytes.Bytes("abc")

	t.Run("Read", func(t *testing.T) {
		assert.Equal(t, []byte("bc"), b.Read(1, 2))
		assert.Nil(t, b.Read(1, 3))
		assert.Nil(t, b.Read(-1, 2))
	})

	t.Run("ReadString", func(t *testing.T) {
		assert.Equal(t, "bc", b.ReadString(1, 2))
	})

	t.Run("Index", func(t *testing.T) {
		assert.Equal(t, 1, b.Index([]byte("bc")))
		assert.Equal(t, -1, b.Index([]byte("bd")))
	})

	t.Run("FindString", func(t *testing.T) {
		assert.Equal(t, 1, b.FindString("bc"))
	})

	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "abc", b.String())
	})
}

func TestBytesTrim(t *testing.T) {
	testCases := []struct {
		name   string
		b      bytes.Bytes
		cutset string
		want   []byte
	}{
		{"trim spaces", bytes.Bytes("  hello world  "), " ", []byte("hello world")},
		{"trim chars", bytes.Bytes("__hello__"), "_", []byte("hello")},
		{"no trim needed", bytes.Bytes("hello"), " ", []byte("hello")},
		{"trim all", bytes.Bytes("abab"), "ab", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.b.Trim(tc.cutset)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBytesTrimSpace(t *testing.T) {
	b := bytes.Bytes("  \t\n hello world \n\t  ")
	want := []byte("hello world")
	got := b.TrimSpace()
	assert.Equal(t, want, got)
}

func TestBytesTrimPrefix(t *testing.T) {
	b := bytes.Bytes("__hello")
	prefix := []byte("__")
	want := []byte("hello")
	got := b.TrimPrefix(prefix)
	assert.Equal(t, want, got)
}

func TestBytesTrimSuffix(t *testing.T) {
	b := bytes.Bytes("hello__")
	suffix := []byte("__")
	want := []byte("hello")
	got := b.TrimSuffix(suffix)
	assert.Equal(t, want, got)
}

func TestBytesReplace(t *testing.T) {
	testCases := []struct {
		name string
		b    bytes.Bytes
		old  []byte
		new  []byte
		n    int
		want []byte
	}{
		{"replace once", bytes.Bytes("hello world"), []byte("l"), []byte("L"), 1, []byte("heLlo world")},
		{"replace all", bytes.Bytes("hello world"), []byte("l"), []byte("L"), -1, []byte("heLLo worLd")},
		{"no replacement", bytes.Bytes("hello"), []byte("x"), []byte("X"), -1, []byte("hello")},
		{"empty old", bytes.Bytes("hello"), []byte(""), []byte("X"), -1, []byte("XhXeXlXlXoX")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.b.Replace(tc.old, tc.new, tc.n)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBytesContains(t *testing.T) {
	b := bytes.Bytes("hello world")
	assert.True(t, b.Contains([]byte("world")))
	assert.False(t, b.Contains([]byte("galaxy")))
}

func TestBytesHasPrefix(t *testing.T) {
	b := bytes.Bytes("hello world")
	assert.True(t, b.HasPrefix([]byte("hello")))
	assert.False(t, b.HasPrefix([]byte("world")))
}

func TestBytesHasSuffix(t *testing.T) {
	b := bytes.Bytes("hello world")
	assert.True(t, b.HasSuffix([]byte("world")))
	assert.False(t, b.HasSuffix([]byte("hello")))
}
