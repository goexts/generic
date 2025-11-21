//go:build !race

package cond_test

import (
	"reflect"
	"testing"

	"github.com/goexts/generic/cond"
)

func TestIf(t *testing.T) {
	type args[T any] struct {
		condition bool
		trueVal   T
		falseVal  T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "true",
			args: args[string]{
				condition: true,
				trueVal:   "true",
				falseVal:  "false",
			},
			want: "true",
		},
		{
			name: "false",
			args: args[string]{
				condition: false,
				trueVal:   "true",
				falseVal:  "false",
			},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cond.If(tt.args.condition, tt.args.trueVal, tt.args.falseVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfFunc(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		trueFn    func() int
		falseFn   func() int
		expected  int
	}{
		{
			name:      "condition true - returns trueFn result",
			condition: true,
			trueFn:    func() int { return 42 },
			falseFn:   func() int { return 0 },
			expected:  42,
		},
		{
			name:      "condition false - returns falseFn result",
			condition: false,
			trueFn:    func() int { return 42 },
			falseFn:   func() int { return 0 },
			expected:  0,
		},
		{
			name:      "both functions return same value",
			condition: true,
			trueFn:    func() int { return 100 },
			falseFn:   func() int { return 100 },
			expected:  100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cond.IfFunc(tt.condition, tt.trueFn, tt.falseFn)
			if got != tt.expected {
				t.Errorf("IfFunc() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIfFunc_StringType(t *testing.T) {
	t.Run("string type handling", func(t *testing.T) {
		got := cond.IfFunc(true,
			func() string { return "success" },
			func() string { return "failure" },
		)
		if got != "success" {
			t.Errorf("Expected 'success', got %q", got)
		}
	})
}

func TestIfFunc_PointerType(t *testing.T) {
	type myStruct struct{ value int }
	t.Run("pointer type handling", func(t *testing.T) {
		trueVal := &myStruct{42}
		falseVal := &myStruct{0}

		got := cond.IfFunc(false,
			func() *myStruct { return trueVal },
			func() *myStruct { return falseVal },
		)

		if got != falseVal {
			t.Errorf("Expected falseVal pointer, got %v", got)
		}
	})
}
