//go:build !race

// Package slices_test implements the functions, types, and interfaces for the module.
package slices_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/slices"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

func TestAppend(t *testing.T) {
	arr := []int{1, 2, 3}
	newArr, oldSize := slices.Append(arr, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, newArr)
	assert.Equal(t, 3, oldSize)

	s := make([]int, 2, 4)
	s[0] = 1
	s[1] = 2
	s2, _ := slices.Append(s, 3)
	assert.Equal(t, []int{1, 2}, s)
	assert.Equal(t, []int{1, 2, 3}, s2)
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		value int
		want  bool
	}{
		{"contains value", []int{1, 2, 3}, 2, true},
		{"does not contain value", []int{1, 2, 3}, 4, false},
		{"empty slice", []int{}, 1, false},
		{"nil slice", nil, 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Contains(tc.slice, tc.value)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCopyAt(t *testing.T) {
	testCases := []struct {
		name        string
		source      []int
		target      []int
		index       int
		expected    []int
		shouldPanic bool
	}{
		{
			name:     "copy elements correctly",
			source:   []int{1, 2, 3, 4, 5},
			target:   []int{6, 7, 8},
			index:    2,
			expected: []int{1, 2, 6, 7, 8},
		},
		{
			name:     "target fits within source capacity",
			source:   make([]int, 5, 10),
			target:   []int{6, 7, 8},
			index:    2,
			expected: []int{0, 0, 6, 7, 8},
		},
		{
			name:     "return modified target array",
			source:   []int{1, 2, 3},
			target:   []int{4, 5},
			index:    1,
			expected: []int{1, 4, 5},
		},
		{
			name:     "source insufficient capacity",
			source:   make([]int, 3),
			target:   []int{4, 5},
			index:    2,
			expected: []int{0, 0, 4, 5},
		},
		{
			name:        "negative index values",
			source:      []int{1, 2, 3},
			target:      []int{4, 5},
			index:       -1,
			shouldPanic: true,
		},
		{
			name:     "index greater than source length",
			source:   []int{1, 2},
			target:   []int{3},
			index:    3,
			expected: []int{1, 2, 0, 3},
		},
		{
			name:     "empty source and target arrays",
			source:   []int{},
			target:   []int{},
			index:    0,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				assert.Panics(t, func() {
					slices.CopyAt(tc.source, tc.target, tc.index)
				})
			} else {
				result := slices.CopyAt(tc.source, tc.target, tc.index)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestCount(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		sub  []int
		want int
	}{
		{"count occurrences", []int{1, 2, 1, 2, 1}, []int{1, 2}, 2},
		{"no occurrences", []int{1, 2, 3}, []int{4}, 0},
		{"empty sub", []int{1, 2, 3}, []int{}, 4},
		{"overlapping", []int{1, 1, 1}, []int{1, 1}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Count(tc.s, tc.sub)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCountArray(t *testing.T) {
	testCases := []struct {
		name string
		ss   []int
		s    int
		want int
	}{
		{"count occurrences", []int{1, 2, 1, 2, 1}, 1, 3},
		{"no occurrences", []int{1, 2, 3}, 4, 0},
		{"empty slice", []int{}, 1, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.CountArray(tc.ss, tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCut(t *testing.T) {
	testCases := []struct {
		name   string
		s      []int
		sep    []int
		before []int
		after  []int
		found  bool
	}{
		{"sep found", []int{1, 2, 3, 4}, []int{2, 3}, []int{1}, []int{4}, true},
		{"sep not found", []int{1, 2, 3}, []int{4}, []int{1, 2, 3}, nil, false},
		{"empty sep", []int{1, 2, 3}, []int{}, []int{}, []int{1, 2, 3}, true},
		{"sep at start", []int{1, 2, 3}, []int{1, 2}, []int{}, []int{3}, true},
		{"sep at end", []int{1, 2, 3}, []int{3}, []int{1, 2}, []int{}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			before, after, found := slices.Cut(tc.s, tc.sep)
			assert.Equal(t, tc.before, before)
			assert.Equal(t, tc.after, after)
			assert.Equal(t, tc.found, found)
		})
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		f     func(int) bool
		want  []int
	}{
		{"filter even", []int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 0 }, []int{2, 4}},
		{"filter none", []int{1, 3, 5}, func(i int) bool { return i%2 == 0 }, []int{}},
		{"filter all", []int{2, 4, 6}, func(i int) bool { return i%2 == 0 }, []int{2, 4, 6}},
		{"empty slice", []int{}, func(_ int) bool { return true }, []int{}},
		{"nil slice", nil, func(_ int) bool { return true }, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Filter(tc.slice, tc.f)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFilterIncluded(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		include []int
		want    []int
	}{
		{"basic inclusion", []int{1, 2, 3, 4, 5}, []int{2, 4, 6}, []int{2, 4}},
		{"no matches", []int{1, 2, 3}, []int{4, 5}, []int{}},
		{"all match", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty include", []int{1, 2, 3}, []int{}, []int{}},
		{"empty slice", []int{}, []int{1, 2, 3}, []int{}},
		{"duplicates in source", []int{1, 2, 2, 3, 3, 3}, []int{2, 4}, []int{2, 2}},
		{"large include list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]int{2, 4, 6, 8, 10, 12},
			[]int{2, 4, 6, 8, 10}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slices.FilterIncluded(tc.slice, tc.include)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestFilterExcluded(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		exclude []int
		want    []int
	}{
		{"basic exclusion", []int{1, 2, 3, 4, 5}, []int{2, 4, 6}, []int{1, 3, 5}},
		{"no matches", []int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3}},
		{"all match", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"empty exclude", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"empty slice", []int{}, []int{1, 2, 3}, []int{}},
		{"duplicates in source", []int{1, 2, 2, 3, 3, 3}, []int{2}, []int{1, 3, 3, 3}},
		{"large exclude list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]int{2, 4, 6, 8, 10, 12},
			[]int{1, 3, 5, 7, 9, 11}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slices.FilterExcluded(tc.slice, tc.exclude)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestIndexSlice(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		sub  []int
		want int
	}{
		{"simple found", []int{1, 2, 3, 1, 2}, []int{2, 3}, 1},
		{"not found", []int{1, 2, 3}, []int{4}, -1},
		{"empty sub", []int{1, 2, 3}, []int{}, 0},
		{"sub equals s", []int{1, 2, 3}, []int{1, 2, 3}, 0},
		{"sub longer than s", []int{1, 2}, []int{1, 2, 3}, -1},
		{"multiple matches", []int{1, 2, 1, 2, 3}, []int{1, 2}, 0},
		{"single element sub", []int{1, 2, 3, 2, 1}, []int{2}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.IndexSlice(tc.s, tc.sub)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestInsertWith(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		value    int
		fn       func(a, b int) bool
		expected []int
	}{
		{
			name:     "insert in the middle",
			slice:    []int{1, 2, 4, 5},
			value:    3,
			fn:       func(a, b int) bool { return a > b },
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "insert at the beginning",
			slice:    []int{2, 3, 4, 5},
			value:    1,
			fn:       func(a, b int) bool { return a > b },
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "insert at the end",
			slice:    []int{1, 2, 3, 4},
			value:    5,
			fn:       func(a, b int) bool { return a > b },
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "insert into a nil slice",
			slice:    nil,
			value:    1,
			fn:       func(a, b int) bool { return a > b },
			expected: []int{1},
		},
		{
			name:     "insert into an empty slice",
			slice:    []int{},
			value:    1,
			fn:       func(a, b int) bool { return a > b },
			expected: []int{1},
		},
		{
			name:     "comparison is always true",
			slice:    []int{1, 2, 3},
			value:    0,
			fn:       func(_, _ int) bool { return true },
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "comparison is always false",
			slice:    []int{1, 2, 3},
			value:    4,
			fn:       func(_, _ int) bool { return false },
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "slice with duplicates",
			slice:    []int{1, 2, 2, 3},
			value:    2,
			fn:       func(a, b int) bool { return a > b },
			expected: []int{1, 2, 2, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slices.InsertWith(tc.slice, tc.value, tc.fn)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestJoin(t *testing.T) {
	testCases := []struct {
		name string
		s    [][]int
		sep  []int
		want []int
	}{
		{"multiple slices", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{0}, []int{1, 2, 0, 3, 4, 0, 5, 6}},
		{"empty separator", [][]int{{1}, {2}, {3}}, []int{}, []int{1, 2, 3}},
		{"single slice", [][]int{{1, 2, 3}}, []int{0}, []int{1, 2, 3}},
		{"empty slices", [][]int{{}, {}, {}}, []int{0}, []int{0, 0}},
		{"empty input", [][]int{}, []int{}, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Join(tc.s, tc.sep)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLastIndexSlice(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		sep  []int
		want int
	}{
		{"simple found", []int{1, 2, 3, 1, 2}, []int{1, 2}, 3},
		{"not found", []int{1, 2, 3}, []int{4}, -1},
		{"empty sep", []int{1, 2, 3}, []int{}, 3},
		{"sep equals s", []int{1, 2, 3}, []int{1, 2, 3}, 0},
		{"sep longer than s", []int{1, 2}, []int{1, 2, 3}, -1},
		{"multiple matches", []int{1, 2, 1, 2, 3}, []int{1, 2}, 2},
		{"single element sep", []int{1, 2, 3, 2, 1}, []int{2}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.LastIndexSlice(tc.s, tc.sep)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		s := []int{1, 2, 3}
		f := func(i int) string { return "v" + strconv.Itoa(i) }
		result := slices.Map(s, f)
		assert.Equal(t, []string{"v1", "v2", "v3"}, result)
	})

	t.Run("nil slice", func(t *testing.T) {
		var s []int
		f := func(i int) int { return i }
		result := slices.Map(s, f)
		assert.Equal(t, []int{}, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		s := []int{}
		f := func(i int) int { return i }
		result := slices.Map(s, f)
		assert.Equal(t, []int{}, result)
	})
}

func TestOverWithError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		s := []int{1, 2, 3}
		var result []int
		iterator := slices.OverWithError(s, nil)
		iterator(func(_ int, v int) bool {
			result = append(result, v*2)
			return true
		})
		assert.Equal(t, []int{2, 4, 6}, result)
	})

	t.Run("with error", func(t *testing.T) {
		s := []int{1, 2, 3}
		var result []int
		iterator := slices.OverWithError(s, errors.New("test error"))
		iterator(func(_ int, v int) bool {
			result = append(result, v)
			return true
		})
		assert.Empty(t, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		s := []int{}
		var result []int
		iterator := slices.OverWithError(s, nil)
		iterator(func(_ int, v int) bool {
			result = append(result, v)
			return true
		})
		assert.Empty(t, result)
	})

	t.Run("stop iteration", func(t *testing.T) {
		s := []int{1, 2, 3}
		var result []int
		iterator := slices.OverWithError(s, nil)
		iterator(func(_ int, v int) bool {
			result = append(result, v)
			return v != 2
		})
		assert.Equal(t, []int{1, 2}, result)
	})
}

func TestRead(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		offset int
		limit  int
		want   []int
	}{
		{"read full", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
		{"read partial", []int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3, 4}},
		{"offset out of bounds", []int{1, 2, 3}, 4, 1, nil},
		{"limit out of bounds", []int{1, 2, 3}, 1, 3, nil},
		{"negative offset", []int{1, 2, 3}, -1, 2, nil},
		{"negative limit", []int{1, 2, 3}, 0, -1, nil},
		{"empty slice", []int{}, 0, 0, nil},
		{"read from nil", nil, 0, 0, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Read(tc.arr, tc.offset, tc.limit)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	t.Run("sum integers", func(t *testing.T) {
		s := []int{1, 2, 3, 4}
		f := func(acc int, v int) int { return acc + v }
		result := slices.Reduce(s, 0, f)
		assert.Equal(t, 10, result)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		f := func(acc string, v string) string { return acc + v }
		result := slices.Reduce(s, "", f)
		assert.Equal(t, "abc", result)
	})

	t.Run("empty slice", func(t *testing.T) {
		s := []int{}
		f := func(acc int, v int) int { return acc + v }
		result := slices.Reduce(s, 10, f)
		assert.Equal(t, 10, result)
	})
}

func TestRemoveWith(t *testing.T) {
	t.Run("int slices", func(t *testing.T) {
		testCases := []struct {
			name     string
			input    []int
			fn       func(a int) bool
			expected []int
		}{
			{
				name:     "remove even numbers",
				input:    []int{1, 2, 3, 4, 5},
				fn:       func(a int) bool { return a%2 == 0 },
				expected: []int{1, 3, 5},
			},
			{
				name:     "handles empty slices",
				input:    []int{},
				fn:       func(a int) bool { return a > 0 },
				expected: []int{},
			},
			{
				name:     "all elements satisfy comparison",
				input:    []int{2, 4, 6},
				fn:       func(a int) bool { return a%2 == 0 },
				expected: []int{},
			},
			{
				name:     "no elements satisfy comparison",
				input:    []int{1, 3, 5},
				fn:       func(a int) bool { return a%2 == 0 },
				expected: []int{1, 3, 5},
			},
			{
				name:     "handles duplicate elements",
				input:    []int{1, 2, 2, 3},
				fn:       func(a int) bool { return a == 2 },
				expected: []int{1, 3},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := slices.RemoveWith(tc.input, tc.fn)
				assert.Equal(t, tc.expected, result)
			})
		}
	})

	t.Run("string slices", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		fn := func(a string) bool { return a == "banana" }
		result := slices.RemoveWith(input, fn)
		expected := []string{"apple", "cherry"}
		assert.Equal(t, expected, result)
	})

	t.Run("float64 slices", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		fn := func(a float64) bool { return a > 2.0 }
		result := slices.RemoveWith(input, fn)
		expected := []float64{1.1}
		assert.Equal(t, expected, result)
	})
}

func TestRepeat(t *testing.T) {
	testCases := []struct {
		name  string
		b     []int
		count int
		want  []int
		panic bool
	}{
		{"repeat 3 times", []int{1, 2}, 3, []int{1, 2, 1, 2, 1, 2}, false},
		{"repeat 0 times", []int{1, 2}, 0, []int{}, false},
		{"empty slice", []int{}, 5, []int{}, false},
		{"negative count", []int{1, 2}, -1, nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				assert.Panics(t, func() { slices.Repeat(tc.b, tc.count) })
			} else {
				got := slices.Repeat(tc.b, tc.count)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"odd length", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"empty slice", []int{}, []int{}},
		{"nil slice", nil, nil},
		{"single element", []int{1}, []int{1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slices.Reverse(tc.slice)
			assert.Equal(t, tc.want, tc.slice)
		})
	}
}

func TestTransform(t *testing.T) {
	t.Run("basic string conversion", func(t *testing.T) {
		in := []int{1, 2, 3}
		out := slices.Transform(in, func(n int) (string, bool) {
			return string(rune(n + 64)), true // 1→A, 2→B, 3→C
		})
		expected := []string{"A", "B", "C"}
		if len(out) != len(expected) || out[0] != "A" || out[2] != "C" {
			t.Errorf("Expected %v, got %v", expected, out)
		}
	})

	t.Run("partial filtering", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		out := slices.Transform(in, func(n int) (int, bool) {
			return n * 2, n%2 == 0
		})
		if len(out) != 2 || out[0] != 4 || out[1] != 8 {
			t.Errorf("Expected [4 8], got %v", out)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		var in []string
		out := slices.Transform(in, func(s string) (int, bool) {
			return len(s), true
		})
		if len(out) != 0 {
			t.Error("Expected empty slice")
		}
	})

	t.Run("struct conversion", func(t *testing.T) {
		in := []Point{{1, 2}, {3, 4}}
		out := slices.Transform(in, func(p Point) (string, bool) {
			return p.String(), true
		})
		if out[0] != "(1,2)" || out[1] != "(3,4)" {
			t.Errorf("Expected [(1,2) (3,4)], got %v", out)
		}
	})

	t.Run("nil slice handling", func(t *testing.T) {
		var in []float64
		out := slices.Transform(in, func(f float64) (int, bool) {
			return int(f), true
		})
		assert.Equal(t, []int{}, out)
	})
}

func TestUnique(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"unique elements", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"some duplicates", []int{1, 2, 2, 3, 1, 4}, []int{1, 2, 3, 4}},
		{"all duplicates", []int{1, 1, 1, 1}, []int{1}},
		{"empty slice", []int{}, []int{}},
		{"nil slice", nil, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Unique(tc.slice)
			assert.Equal(t, tc.want, got)
		})
	}
}
