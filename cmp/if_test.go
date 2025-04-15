// Package cmp implements the functions, types, and interfaces for the module.
package cmp_test

import (
	"reflect"
	"testing"

	"github.com/goexts/generic/cmp"
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
			if got := cmp.If(tt.args.condition, tt.args.trueVal, tt.args.falseVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
