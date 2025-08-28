package runes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/slices/runes"
)

func TestRunes(t *testing.T) {
	r := runes.Runes{'a', 'b', 'c'}

	t.Run("Read", func(t *testing.T) {
		assert.Equal(t, []rune{'b', 'c'}, r.Read(1, 2))
		assert.Nil(t, r.Read(1, 3))
		assert.Nil(t, r.Read(-1, 2))
	})

	t.Run("ReadString", func(t *testing.T) {
		assert.Equal(t, "bc", r.ReadString(1, 2))
	})

	t.Run("Index", func(t *testing.T) {
		assert.Equal(t, 1, r.Index([]rune{'b', 'c'}))
		assert.Equal(t, -1, r.Index([]rune{'b', 'd'}))
	})

	t.Run("FindString", func(t *testing.T) {
		assert.Equal(t, 1, r.FindString("bc"))
	})

	t.Run("StringArray", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, r.StringArray())
	})

	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "abc", r.String())
	})

	t.Run("ToBytes", func(t *testing.T) {
		assert.Equal(t, []byte("abc"), r.ToBytes())
	})

	t.Run("StringToRunes", func(t *testing.T) {
		assert.Equal(t, runes.Runes{'a', 'b', 'c'}, runes.StringToRunes("abc"))
	})
}

func TestRunesTrim(t *testing.T) {
	testCases := []struct {
		name   string
		r      runes.Runes
		cutset string
		want   []rune
	}{
		{"trim spaces", runes.Runes("  hello world  "), " ", []rune("hello world")},
		{"trim chars", runes.Runes("__hello__"), "_", []rune("hello")},
		{"no trim needed", runes.Runes("hello"), " ", []rune("hello")},
		{"trim all", runes.Runes("abab"), "ab", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.r.Trim(tc.cutset)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRunesTrimSpace(t *testing.T) {
	r := runes.Runes("  \t\n hello world \n\t  ")
	want := []rune("hello world")
	got := r.TrimSpace()
	assert.Equal(t, want, got)
}

func TestRunesTrimPrefix(t *testing.T) {
	r := runes.Runes("__hello")
	prefix := []rune("__")
	want := []rune("hello")
	got := r.TrimPrefix(prefix)
	assert.Equal(t, want, got)
}

func TestRunesTrimSuffix(t *testing.T) {
	r := runes.Runes("hello__")
	suffix := []rune("__")
	want := []rune("hello")
	got := r.TrimSuffix(suffix)
	assert.Equal(t, want, got)
}

func TestRunesReplace(t *testing.T) {
	testCases := []struct {
		name string
		r    runes.Runes
		old  []rune
		new  []rune
		n    int
		want []rune
	}{
		{"replace once", runes.Runes("hello world"), []rune("l"), []rune("L"), 1, []rune("heLlo world")},
		{"replace all", runes.Runes("hello world"), []rune("l"), []rune("L"), -1, []rune("heLLo worLd")},
		{"no replacement", runes.Runes("hello"), []rune("x"), []rune("X"), -1, []rune("hello")},
		{"empty old", runes.Runes("hello"), []rune(""), []rune("X"), -1, []rune("hello")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.r.Replace(tc.old, tc.new, tc.n)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRunesContains(t *testing.T) {
	r := runes.Runes("hello world")
	assert.True(t, r.Contains([]rune("world")))
	assert.False(t, r.Contains([]rune("galaxy")))
}

func TestRunesHasPrefix(t *testing.T) {
	r := runes.Runes("hello world")
	assert.True(t, r.HasPrefix([]rune("hello")))
	assert.False(t, r.HasPrefix([]rune("world")))
}

func TestRunesHasSuffix(t *testing.T) {
	r := runes.Runes("hello world")
	assert.True(t, r.HasSuffix([]rune("world")))
	assert.False(t, r.HasSuffix([]rune("hello")))
}
