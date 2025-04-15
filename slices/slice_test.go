//go:build !race

// Package slices_test implements the functions, types, and interfaces for the module.
package slices_test

import (
	"strconv"
	"testing"

	"github.com/goexts/generic/slices"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
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
		if out != nil {
			t.Error("Expected empty output for empty input")
		}
	})
}
