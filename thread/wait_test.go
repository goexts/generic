//go:build !race

// Package thread implements the functions, types, and interfaces for the module.
package thread

import (
	"reflect"
	"testing"
)

func TestWait(t *testing.T) {
	type args[T any] struct {
		v <-chan T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "test",
			args: args[int]{
				v: Async(func() int {
					return 1
				}),
			},
			want: 1,
		},
		{
			name: "test",
			args: args[int]{
				v: Async(func() int {
					return 2
				}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wait(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wait() = %v, want %v", got, tt.want)
			}
		})
	}
}
