// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	. "github.com/goexts/generic/slices"
)

type BinOpTest struct {
	a string
	b string
	i int
}

func TestEqual(t *testing.T) {
	// Run the tests and check for allocation at the same time.
	allocs := testing.AllocsPerRun(10, func() {
		for _, tt := range compareTests {
			eql := Equal(tt.a, tt.b)
			if eql != (tt.i == 0) {
				t.Errorf(`Equal(%q, %q) = %v`, tt.a, tt.b, eql)
			}
		}
	})
	if allocs > 0 {
		t.Errorf("Equal allocated %v times", allocs)
	}
}

func TestEqualExhaustive(t *testing.T) {
	var size = 128
	if testing.Short() {
		size = 32
	}
	a := make([]rune, size)
	b := make([]rune, size)
	b_init := make([]rune, size)
	// randomish but deterministic data
	for i := 0; i < size; i++ {
		a[i] = rune(17 * i)
		b_init[i] = rune(23*i + 100)
	}

	for len := 0; len <= size; len++ {
		for x := 0; x <= size-len; x++ {
			for y := 0; y <= size-len; y++ {
				copy(b, b_init)
				copy(b[y:y+len], a[x:x+len])
				if !Equal(a[x:x+len], b[y:y+len]) || !Equal(b[y:y+len], a[x:x+len]) {
					t.Errorf("Equal(%d, %d, %d) = false", len, x, y)
				}
			}
		}
	}
}

// make sure Equal returns false for minimally different strings. The data
// is all zeros except for a single one in one location.
func TestNotEqual(t *testing.T) {
	var size = 128
	if testing.Short() {
		size = 32
	}
	a := make([]rune, size)
	b := make([]rune, size)

	for len := 0; len <= size; len++ {
		for x := 0; x <= size-len; x++ {
			for y := 0; y <= size-len; y++ {
				for diffpos := x; diffpos < x+len; diffpos++ {
					a[diffpos] = 1
					if Equal(a[x:x+len], b[y:y+len]) || Equal(b[y:y+len], a[x:x+len]) {
						t.Errorf("NotEqual(%d, %d, %d, %d) = true", len, x, y, diffpos)
					}
					a[diffpos] = 0
				}
			}
		}
	}
}

var indexTests = []BinOpTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "baz", -1},
	{"foo", "foo", 0},
	{"oofofoofooo", "f", 2},
	{"oofofoofooo", "foo", 4},
	{"barfoobarfoo", "foo", 3},
	{"foo", "", 0},
	{"foo", "o", 1},
	{"abcABCabc", "A", 3},
	// cases with one byte strings - test IndexArray and special case in Index()
	{"", "a", -1},
	{"x", "a", -1},
	{"x", "x", 0},
	{"abc", "a", 0},
	{"abc", "b", 1},
	{"abc", "c", 2},
	{"abc", "x", -1},
	{"barfoobarfooyyyzzzyyyzzzyyyzzzyyyxxxzzzyyy", "x", 33},
	{"fofofofooofoboo", "oo", 7},
	{"fofofofofofoboo", "ob", 11},
	{"fofofofofofoboo", "boo", 12},
	{"fofofofofofoboo", "oboo", 11},
	{"fofofofofoooboo", "fooo", 8},
	{"fofofofofofoboo", "foboo", 10},
	{"fofofofofofoboo", "fofob", 8},
	{"fofofofofofofoffofoobarfoo", "foffof", 12},
	{"fofofofofoofofoffofoobarfoo", "foffof", 13},
	{"fofofofofofofoffofoobarfoo", "foffofo", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofo", 13},
	{"fofofofofoofofoffofoobarfoo", "foffofoo", 13},
	{"fofofofofofofoffofoobarfoo", "foffofoo", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofoob", 13},
	{"fofofofofofofoffofoobarfoo", "foffofoob", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofooba", 13},
	{"fofofofofofofoffofoobarfoo", "foffofooba", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofoobar", 13},
	{"fofofofofofofoffofoobarfoo", "foffofoobar", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofoobarf", 13},
	{"fofofofofofofoffofoobarfoo", "foffofoobarf", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofoobarfo", 13},
	{"fofofofofofofoffofoobarfoo", "foffofoobarfo", 12},
	{"fofofofofoofofoffofoobarfoo", "foffofoobarfoo", 13},
	{"fofofofofofofoffofoobarfoo", "foffofoobarfoo", 12},
	{"fofofofofoofofoffofoobarfoo", "ofoffofoobarfoo", 12},
	{"fofofofofofofoffofoobarfoo", "ofoffofoobarfoo", 11},
	{"fofofofofoofofoffofoobarfoo", "fofoffofoobarfoo", 11},
	{"fofofofofofofoffofoobarfoo", "fofoffofoobarfoo", 10},
	{"fofofofofoofofoffofoobarfoo", "foobars", -1},
	{"foofyfoobarfoobar", "y", 4},
	{"oooooooooooooooooooooo", "r", -1},
	{"oxoxoxoxoxoxoxoxoxoxoxoy", "oy", 22},
	{"oxoxoxoxoxoxoxoxoxoxoxox", "oy", -1},
	// test fallback to Rabin-Karp.
	{"000000000000000000000000000000000000000000000000000000000000000000000001", "0000000000000000000000000000000000000000000000000000000000000000001", 5},
}

var lastIndexTests = []BinOpTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"foo", "f", 0},
	{"oofofoofooo", "f", 7},
	{"oofofoofooo", "foo", 7},
	{"barfoobarfoo", "foo", 9},
	{"foo", "", 3},
	{"foo", "o", 2},
	{"abcABCabc", "A", 3},
	{"abcABCabc", "a", 6},
}

// Execute f on each test case.  funcName should be the name of f; it's used
// in failure reports.
func runIndexTests(t *testing.T, f func(s, sep []rune) int, funcName string, testCases []BinOpTest) {
	for _, test := range testCases {
		a := []rune(test.a)
		b := []rune(test.b)
		actual := f(a, b)
		if actual != test.i {
			t.Errorf("%s(%q,%q) = %v; want %v", funcName, a, b, actual, test.i)
		}
	}
	var allocTests = []struct {
		a []rune
		b []rune
		i int
	}{
		// case for function Index.
		{[]rune("000000000000000000000000000000000000000000000000000000000000000000000001"), []rune("0000000000000000000000000000000000000000000000000000000000000000001"), 5},
		// case for function LastIndex.
		{[]rune("000000000000000000000000000000000000000000000000000000000000000010000"), []rune("00000000000000000000000000000000000000000000000000000000000001"), 3},
	}
	allocs := testing.AllocsPerRun(100, func() {
		if i := Index(allocTests[1].a, allocTests[1].b); i != allocTests[1].i {
			t.Errorf("Index([]rune(%q), []rune(%q)) = %v; want %v", allocTests[1].a, allocTests[1].b, i, allocTests[1].i)
		}
		if i := LastIndex(allocTests[0].a, allocTests[0].b); i != allocTests[0].i {
			t.Errorf("LastIndex([]rune(%q), []rune(%q)) = %v; want %v", allocTests[0].a, allocTests[0].b, i, allocTests[0].i)
		}
	})
	if allocs != 0 {
		t.Errorf("expected no allocations, got %f", allocs)
	}
}

func TestIndex(t *testing.T)     { runIndexTests(t, Index[[]rune], "Index", indexTests) }
func TestLastIndex(t *testing.T) { runIndexTests(t, LastIndex[[]rune], "LastIndex", lastIndexTests) }

func TestIndexArray(t *testing.T) {
	for _, tt := range indexTests {
		if len(tt.b) != 1 {
			continue
		}
		a := []rune(tt.a)
		b := []rune(tt.b)[0]
		pos := IndexArray(a, b)
		if pos != tt.i {
			t.Errorf(`IndexArray(%q, '%c') = %v`, tt.a, b, pos)
		}
		// posp := IndexArrayPortable(a, b)
		// if posp != tt.i {
		// 	t.Errorf(`indexBytePortable(%q, '%c') = %v`, tt.a, b, posp)
		// }
	}
}

func TestLastIndexArray(t *testing.T) {
	testCases := []BinOpTest{
		{"", "q", -1},
		{"abcdef", "q", -1},
		{"abcdefabcdef", "a", len("abcdef")},      // something in the middle
		{"abcdefabcdef", "f", len("abcdefabcde")}, // last byte
		{"zabcdefabcdef", "z", 0},                 // first byte
		{"a☺b☻c☹d", "b", len("a☺")},               // non-ascii
	}
	for _, test := range testCases {
		actual := LastIndexArray([]byte(test.a), []byte(test.b)[0])
		if actual != test.i {
			t.Errorf("LastIndexArray(%q,%c) = %v; want %v", test.a, test.b[0], actual, test.i)
		}
	}
}

// test a larger buffer with different sizes and alignments
func TestIndexArrayBig(t *testing.T) {
	var n = 1024
	if testing.Short() {
		n = 128
	}
	b := make([]rune, n)
	for i := 0; i < n; i++ {
		// different start alignments
		b1 := b[i:]
		for j := 0; j < len(b1); j++ {
			b1[j] = 'x'
			pos := IndexArray(b1, 'x')
			if pos != j {
				t.Errorf("IndexArray(%q, 'x') = %v", b1, pos)
			}
			b1[j] = 0
			pos = IndexArray(b1, 'x')
			if pos != -1 {
				t.Errorf("IndexArray(%q, 'x') = %v", b1, pos)
			}
		}
		// different end alignments
		b1 = b[:i]
		for j := 0; j < len(b1); j++ {
			b1[j] = 'x'
			pos := IndexArray(b1, 'x')
			if pos != j {
				t.Errorf("IndexArray(%q, 'x') = %v", b1, pos)
			}
			b1[j] = 0
			pos = IndexArray(b1, 'x')
			if pos != -1 {
				t.Errorf("IndexArray(%q, 'x') = %v", b1, pos)
			}
		}
		// different start and end alignments
		b1 = b[i/2 : n-(i+1)/2]
		for j := 0; j < len(b1); j++ {
			b1[j] = 'x'
			pos := IndexArray(b1, 'x')
			if pos != j {
				t.Errorf("IndexArray(%q, 'x') = %v", b1, pos)
			}
			b1[j] = 0
			pos = IndexArray(b1, 'x')
			if pos != -1 {
				t.Errorf("IndexArray(%q, 'x') = %v", b1, pos)
			}
		}
	}
}

// test a small index across all page offsets
func TestIndexArraySmall(t *testing.T) {
	b := make([]rune, 5015) // bigger than a page
	// Make sure we find the correct byte even when straddling a page.
	for i := 0; i <= len(b)-15; i++ {
		for j := 0; j < 15; j++ {
			b[i+j] = rune(100 + j)
		}
		for j := 0; j < 15; j++ {
			p := IndexArray(b[i:i+15], rune(100+j))
			if p != j {
				t.Errorf("IndexArray(%q, %d) = %d", b[i:i+15], 100+j, p)
			}
		}
		for j := 0; j < 15; j++ {
			b[i+j] = 0
		}
	}
	// Make sure matches outside the slice never trigger.
	for i := 0; i <= len(b)-15; i++ {
		for j := 0; j < 15; j++ {
			b[i+j] = 1
		}
		for j := 0; j < 15; j++ {
			p := IndexArray(b[i:i+15], rune(0))
			if p != -1 {
				t.Errorf("IndexArray(%q, %d) = %d", b[i:i+15], 0, p)
			}
		}
		for j := 0; j < 15; j++ {
			b[i+j] = 0
		}
	}
}

// test count of a single byte across page offsets
func TestCountByte(t *testing.T) {
	b := make([]rune, 5015) // bigger than a page
	windows := []int{1, 2, 3, 4, 15, 16, 17, 31, 32, 33, 63, 64, 65, 128}
	testCountWindow := func(i, window int) {
		for j := 0; j < window; j++ {
			b[i+j] = rune(100)
			p := Count(b[i:i+window], []rune{100})
			if p != j+1 {
				t.Errorf("TestCountByte.Count(%q, 100) = %d", b[i:i+window], p)
			}
		}
	}

	maxWnd := windows[len(windows)-1]

	for i := 0; i <= 2*maxWnd; i++ {
		for _, window := range windows {
			if window > len(b[i:]) {
				window = len(b[i:])
			}
			testCountWindow(i, window)
			for j := 0; j < window; j++ {
				b[i+j] = rune(0)
			}
		}
	}
	for i := 4096 - (maxWnd + 1); i < len(b); i++ {
		for _, window := range windows {
			if window > len(b[i:]) {
				window = len(b[i:])
			}
			testCountWindow(i, window)
			for j := 0; j < window; j++ {
				b[i+j] = rune(0)
			}
		}
	}
}

// Make sure we don't count bytes outside our window
func TestCountByteNoMatch(t *testing.T) {
	b := make([]rune, 5015)
	windows := []int{1, 2, 3, 4, 15, 16, 17, 31, 32, 33, 63, 64, 65, 128}
	for i := 0; i <= len(b); i++ {
		for _, window := range windows {
			if window > len(b[i:]) {
				window = len(b[i:])
			}
			// Fill the window with non-match
			for j := 0; j < window; j++ {
				b[i+j] = rune(100)
			}
			// Try to find something that doesn't exist
			p := Count(b[i:i+window], []rune{0})
			if p != 0 {
				t.Errorf("TestCountByteNoMatch(%q, 0) = %d", b[i:i+window], p)
			}
			for j := 0; j < window; j++ {
				b[i+j] = rune(0)
			}
		}
	}
}

var bmbuf []rune

func valName(x int) string {
	if s := x >> 20; s<<20 == x {
		return fmt.Sprintf("%dM", s)
	}
	if s := x >> 10; s<<10 == x {
		return fmt.Sprintf("%dK", s)
	}
	return fmt.Sprint(x)
}

func benchBytes(b *testing.B, sizes []int, f func(b *testing.B, n int)) {
	for _, n := range sizes {
		if isRaceBuilder && n > 4<<10 {
			continue
		}
		b.Run(valName(n), func(b *testing.B) {
			if len(bmbuf) < n {
				bmbuf = make([]rune, n)
			}
			b.SetBytes(int64(n))
			f(b, n)
		})
	}
}

var indexSizes = []int{10, 32, 4 << 10, 4 << 20, 64 << 20}

var isRaceBuilder = false

func BenchmarkIndexArray(b *testing.B) {
	benchBytes(b, indexSizes, bmIndexArray(IndexArray[[]rune]))
}

// func BenchmarkIndexArrayPortable(b *testing.B) {
// 	benchBytes(b, indexSizes, bmIndexArray(IndexArrayPortable))
// }

func bmIndexArray(index func([]rune, rune) int) func(b *testing.B, n int) {
	return func(b *testing.B, n int) {
		buf := bmbuf[0:n]
		buf[n-1] = 'x'
		for i := 0; i < b.N; i++ {
			j := index(buf, 'x')
			if j != n-1 {
				b.Fatal("bad index", j)
			}
		}
		buf[n-1] = '\x00'
	}
}

func BenchmarkEqual(b *testing.B) {
	b.Run("0", func(b *testing.B) {
		var buf [4]byte
		buf1 := buf[0:0]
		buf2 := buf[1:1]
		for i := 0; i < b.N; i++ {
			eq := Equal(buf1, buf2)
			if !eq {
				b.Fatal("bad equal")
			}
		}
	})

	sizes := []int{1, 6, 9, 15, 16, 20, 32, 4 << 10, 4 << 20, 64 << 20}
	benchBytes(b, sizes, bmEqual(Equal[[]rune]))
}

func bmEqual(equal func([]rune, []rune) bool) func(b *testing.B, n int) {
	return func(b *testing.B, n int) {
		if len(bmbuf) < 2*n {
			bmbuf = make([]rune, 2*n)
		}
		buf1 := bmbuf[0:n]
		buf2 := bmbuf[n : 2*n]
		buf1[n-1] = 'x'
		buf2[n-1] = 'x'
		for i := 0; i < b.N; i++ {
			eq := equal(buf1, buf2)
			if !eq {
				b.Fatal("bad equal")
			}
		}
		buf1[n-1] = '\x00'
		buf2[n-1] = '\x00'
	}
}

func BenchmarkEqualBothUnaligned(b *testing.B) {
	sizes := []int{64, 4 << 10}
	if !isRaceBuilder {
		sizes = append(sizes, []int{4 << 20, 64 << 20}...)
	}
	maxSize := 2 * (sizes[len(sizes)-1] + 8)
	if len(bmbuf) < maxSize {
		bmbuf = make([]rune, maxSize)
	}

	for _, n := range sizes {
		for _, off := range []int{0, 1, 4, 7} {
			buf1 := bmbuf[off : off+n]
			buf2Start := (len(bmbuf) / 2) + off
			buf2 := bmbuf[buf2Start : buf2Start+n]
			buf1[n-1] = 'x'
			buf2[n-1] = 'x'
			b.Run(fmt.Sprint(n, off), func(b *testing.B) {
				b.SetBytes(int64(n))
				for i := 0; i < b.N; i++ {
					eq := Equal(buf1, buf2)
					if !eq {
						b.Fatal("bad equal")
					}
				}
			})
			buf1[n-1] = '\x00'
			buf2[n-1] = '\x00'
		}
	}
}

func BenchmarkIndex(b *testing.B) {
	benchBytes(b, indexSizes, func(b *testing.B, n int) {
		buf := bmbuf[0:n]
		buf[n-1] = 'x'
		for i := 0; i < b.N; i++ {
			j := Index(buf, buf[n-7:])
			if j != n-7 {
				b.Fatal("bad index", j)
			}
		}
		buf[n-1] = '\x00'
	})
}

func BenchmarkIndexEasy(b *testing.B) {
	benchBytes(b, indexSizes, func(b *testing.B, n int) {
		buf := bmbuf[0:n]
		buf[n-1] = 'x'
		buf[n-7] = 'x'
		for i := 0; i < b.N; i++ {
			j := Index(buf, buf[n-7:])
			if j != n-7 {
				b.Fatal("bad index", j)
			}
		}
		buf[n-1] = '\x00'
		buf[n-7] = '\x00'
	})
}

func BenchmarkCount(b *testing.B) {
	benchBytes(b, indexSizes, func(b *testing.B, n int) {
		buf := bmbuf[0:n]
		buf[n-1] = 'x'
		for i := 0; i < b.N; i++ {
			j := Count(buf, buf[n-7:])
			if j != 1 {
				b.Fatal("bad count", j)
			}
		}
		buf[n-1] = '\x00'
	})
}

func BenchmarkCountEasy(b *testing.B) {
	benchBytes(b, indexSizes, func(b *testing.B, n int) {
		buf := bmbuf[0:n]
		buf[n-1] = 'x'
		buf[n-7] = 'x'
		for i := 0; i < b.N; i++ {
			j := Count(buf, buf[n-7:])
			if j != 1 {
				b.Fatal("bad count", j)
			}
		}
		buf[n-1] = '\x00'
		buf[n-7] = '\x00'
	})
}

func BenchmarkCountSingle(b *testing.B) {
	benchBytes(b, indexSizes, func(b *testing.B, n int) {
		buf := bmbuf[0:n]
		step := 8
		for i := 0; i < len(buf); i += step {
			buf[i] = 1
		}
		expect := (len(buf) + (step - 1)) / step
		for i := 0; i < b.N; i++ {
			j := Count(buf, []rune{1})
			if j != expect {
				b.Fatal("bad count", j, expect)
			}
		}
		for i := 0; i < len(buf); i++ {
			buf[i] = 0
		}
	})
}

type RepeatTest struct {
	in, out string
	count   int
}

var longString = "a" + string(make([]rune, 1<<16)) + "z"

var RepeatTests = []RepeatTest{
	{"", "", 0},
	{"", "", 1},
	{"", "", 2},
	{"-", "", 0},
	{"-", "-", 1},
	{"-", "----------", 10},
	{"abc ", "abc abc abc ", 3},
	// Tests for results over the chunkLimit
	{string(rune(0)), string(make([]rune, 1<<16)), 1 << 16},
	{longString, longString + longString, 2},
}

func TestRepeat(t *testing.T) {
	for _, tt := range RepeatTests {
		tin := []rune(tt.in)
		tout := []rune(tt.out)
		a := Repeat(tin, tt.count)
		if !Equal(a, tout) {
			t.Errorf("Repeat(%q, %d) = %q; want %q", tin, tt.count, a, tout)
			continue
		}
	}
}

func repeat(b []rune, count int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%s", v)
			}
		}
	}()

	Repeat(b, count)

	return
}

// See Issue golang.org/issue/16237
func TestRepeatCatchesOverflow(t *testing.T) {
	tests := [...]struct {
		s      string
		count  int
		errStr string
	}{
		0: {"--", -2147483647, "negative"},
		1: {"", int(^uint(0) >> 1), ""},
		2: {"-", 10, ""},
		3: {"gopher", 0, ""},
		4: {"-", -1, "negative"},
		5: {"--", -102, "negative"},
		6: {string(make([]rune, 255)), int((^uint(0))/255 + 1), "overflow"},
	}

	for i, tt := range tests {
		err := repeat([]rune(tt.s), tt.count)
		if tt.errStr == "" {
			if err != nil {
				t.Errorf("#%d panicked %v", i, err)
			}
			continue
		}

		if err == nil || !strings.Contains(err.Error(), tt.errStr) {
			t.Errorf("#%d expected %q got %q", i, tt.errStr, err)
		}
	}
}

func runesEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, r := range a {
		if r != b[i] {
			return false
		}
	}
	return true
}

type RunesTest struct {
	in    string
	out   []rune
	lossy bool
}

var RunesTests = []RunesTest{
	{"", []rune{}, false},
	{" ", []rune{32}, false},
	{"ABC", []rune{65, 66, 67}, false},
	{"abc", []rune{97, 98, 99}, false},
	{"\u65e5\u672c\u8a9e", []rune{26085, 26412, 35486}, false},
	{"ab\x80c", []rune{97, 98, 0xFFFD, 99}, true},
	{"ab\xc0c", []rune{97, 98, 0xFFFD, 99}, true},
}

var cutTests = []struct {
	s, sep        string
	before, after string
	found         bool
}{
	{"abc", "b", "a", "c", true},
	{"abc", "a", "", "bc", true},
	{"abc", "c", "ab", "", true},
	{"abc", "abc", "", "", true},
	{"abc", "", "", "abc", true},
	{"abc", "d", "abc", "", false},
	{"", "d", "", "", false},
	{"", "", "", "", true},
}

func TestCut(t *testing.T) {
	for _, tt := range cutTests {
		if before, after, found := Cut([]rune(tt.s), []rune(tt.sep)); string(before) != tt.before || string(after) != tt.after || found != tt.found {
			t.Errorf("Cut(%q, %q) = %q, %q, %v, want %q, %q, %v", tt.s, tt.sep, before, after, found, tt.before, tt.after, tt.found)
		}
	}
}

//
// var cutPrefixTests = []struct {
// 	s, sep string
// 	after  string
// 	found  bool
// }{
// 	{"abc", "a", "bc", true},
// 	{"abc", "abc", "", true},
// 	{"abc", "", "abc", true},
// 	{"abc", "d", "abc", false},
// 	{"", "d", "", false},
// 	{"", "", "", true},
// }
//
// func TestCutPrefix(t *testing.T) {
// 	for _, tt := range cutPrefixTests {
// 		if after, found := CutPrefix([]rune(tt.s), []rune(tt.sep)); string(after) != tt.after || found != tt.found {
// 			t.Errorf("CutPrefix(%q, %q) = %q, %v, want %q, %v", tt.s, tt.sep, after, found, tt.after, tt.found)
// 		}
// 	}
// }

// var cutSuffixTests = []struct {
// 	s, sep string
// 	before string
// 	found  bool
// }{
// 	{"abc", "bc", "a", true},
// 	{"abc", "abc", "", true},
// 	{"abc", "", "abc", true},
// 	{"abc", "d", "abc", false},
// 	{"", "d", "", false},
// 	{"", "", "", true},
// }
//
// func TestCutSuffix(t *testing.T) {
// 	for _, tt := range cutSuffixTests {
// 		if before, found := CutSuffix([]rune(tt.s), []rune(tt.sep)); string(before) != tt.before || found != tt.found {
// 			t.Errorf("CutSuffix(%q, %q) = %q, %v, want %q, %v", tt.s, tt.sep, before, found, tt.before, tt.found)
// 		}
// 	}
// }
//
// func TestBufferGrowNegative(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err == nil {
// 			t.Fatal("Grow(-1) should have panicked")
// 		}
// 	}()
// 	var b Buffer
// 	b.Grow(-1)
// }
//
// func TestBufferTruncateNegative(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err == nil {
// 			t.Fatal("Truncate(-1) should have panicked")
// 		}
// 	}()
// 	var b Buffer
// 	b.Truncate(-1)
// }
//
// func TestBufferTruncateOutOfRange(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err == nil {
// 			t.Fatal("Truncate(20) should have panicked")
// 		}
// 	}()
// 	var b Buffer
// 	b.Write(make([]rune, 10))
// 	b.Truncate(20)
// }

var containsTests = []struct {
	b, subslice []rune
	want        bool
}{
	{[]rune("hello"), []rune("hel"), true},
	{[]rune("日本語"), []rune("日本"), true},
	{[]rune("hello"), []rune("Hello, world"), false},
	{[]rune("東京"), []rune("京東"), false},
}

func TestContains(t *testing.T) {
	for _, tt := range containsTests {
		if got := Contains(tt.b, tt.subslice); got != tt.want {
			t.Errorf("Contains(%q, %q) = %v, want %v", tt.b, tt.subslice, got, tt.want)
		}
	}
}

var makeFieldsInput = func() []rune {
	x := make([]rune, 1<<20)
	// Input is ~10% space, ~10% 2-byte UTF-8, rest ASCII non-space.
	for i := range x {
		switch rand.Intn(10) {
		case 0:
			x[i] = ' '
		case 1:
			if i > 0 && x[i-1] == 'x' {
				copy(x[i-1:], []rune("χ"))
				break
			}
			fallthrough
		default:
			x[i] = 'x'
		}
	}
	return x
}

var makeFieldsInputASCII = func() []rune {
	x := make([]rune, 1<<20)
	// Input is ~10% space, rest ASCII non-space.
	for i := range x {
		if rand.Intn(10) == 0 {
			x[i] = ' '
		} else {
			x[i] = 'x'
		}
	}
	return x
}

var bytesdata = []struct {
	name string
	data []rune
}{
	{"ASCII", makeFieldsInputASCII()},
	{"Mixed", makeFieldsInput()},
}

func makeBenchInputHard() []rune {
	tokens := [...]string{
		"<a>", "<p>", "<b>", "<strong>",
		"</a>", "</p>", "</b>", "</strong>",
		"hello", "world",
	}
	x := make([]rune, 0, 1<<20)
	for {
		i := rand.Intn(len(tokens))
		if len(x)+len(tokens[i]) >= 1<<20 {
			break
		}
		x = append(x, []rune(tokens[i])...)
	}
	return x
}

var benchInputHard = makeBenchInputHard()

func benchmarkIndexHard(b *testing.B, sep []rune) {
	for i := 0; i < b.N; i++ {
		Index(benchInputHard, sep)
	}
}

func benchmarkLastIndexHard(b *testing.B, sep []rune) {
	for i := 0; i < b.N; i++ {
		LastIndex(benchInputHard, sep)
	}
}

func benchmarkCountHard(b *testing.B, sep []rune) {
	for i := 0; i < b.N; i++ {
		Count(benchInputHard, sep)
	}
}

func BenchmarkIndexHard1(b *testing.B) { benchmarkIndexHard(b, []rune("<>")) }
func BenchmarkIndexHard2(b *testing.B) { benchmarkIndexHard(b, []rune("</pre>")) }
func BenchmarkIndexHard3(b *testing.B) { benchmarkIndexHard(b, []rune("<b>hello world</b>")) }
func BenchmarkIndexHard4(b *testing.B) {
	benchmarkIndexHard(b, []rune("<pre><b>hello</b><strong>world</strong></pre>"))
}

func BenchmarkLastIndexHard1(b *testing.B) { benchmarkLastIndexHard(b, []rune("<>")) }
func BenchmarkLastIndexHard2(b *testing.B) { benchmarkLastIndexHard(b, []rune("</pre>")) }
func BenchmarkLastIndexHard3(b *testing.B) { benchmarkLastIndexHard(b, []rune("<b>hello world</b>")) }

func BenchmarkCountHard1(b *testing.B) { benchmarkCountHard(b, []rune("<>")) }
func BenchmarkCountHard2(b *testing.B) { benchmarkCountHard(b, []rune("</pre>")) }
func BenchmarkCountHard3(b *testing.B) { benchmarkCountHard(b, []rune("<b>hello world</b>")) }
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat([]rune("-"), 80)
	}
}

func BenchmarkRepeatLarge(b *testing.B) {
	s := Repeat([]rune("@"), 8*1024)
	for j := 8; j <= 30; j++ {
		for _, k := range []int{1, 16, 4097} {
			s := s[:k]
			n := (1 << j) / k
			if n == 0 {
				continue
			}
			b.Run(fmt.Sprintf("%d/%d", 1<<j, k), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Repeat(s, n)
				}
				b.SetBytes(int64(n * len(s)))
			})
		}
	}
}

func BenchmarkIndexPeriodic(b *testing.B) {
	key := []rune{1, 1}
	for _, skip := range [...]int{2, 4, 8, 16, 32, 64} {
		b.Run(fmt.Sprintf("IndexPeriodic%d", skip), func(b *testing.B) {
			buf := make([]rune, 1<<16)
			for i := 0; i < len(buf); i += skip {
				buf[i] = 1
			}
			for i := 0; i < b.N; i++ {
				Index(buf, key)
			}
		})
	}
}

//
//func TestClone(t *testing.T) {
//	var cloneTests = [][]rune{
//		[]rune(nil),
//		[]rune{},
//		Clone([]rune{}),
//		[]rune(strings.Repeat("a", 42))[:0],
//		[]rune(strings.Repeat("a", 42))[:0:0],
//		[]rune("short"),
//		[]rune(strings.Repeat("a", 42)),
//	}
//	for _, input := range cloneTests {
//		clone := Clone(input)
//		if !Equal(clone, input) {
//			t.Errorf("Clone(%q) = %q; want %q", input, clone, input)
//		}
//
//		if input == nil && clone != nil {
//			t.Errorf("Clone(%#v) return value should be equal to nil slice.", input)
//		}
//
//		if input != nil && clone == nil {
//			t.Errorf("Clone(%#v) return value should not be equal to nil slice.", input)
//		}
//
//		if cap(input) != 0 && unsafe.SliceData(input) == unsafe.SliceData(clone) {
//			t.Errorf("Clone(%q) return value should not reference inputs backing memory.", input)
//		}
//	}
//}
