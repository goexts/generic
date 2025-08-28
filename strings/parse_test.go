package strings

import (
	"reflect"
	"testing"
)

func TestParseOr(t *testing.T) {
	type args[T any] struct {
		s   string
		def []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "int",
			run: func(t *testing.T) {
				tests := []testCase[int]{
					{
						name: "valid",
						args: args[int]{s: "123"},
						want: 123,
					},
					{
						name: "invalid with default",
						args: args[int]{s: "abc", def: []int{456}},
						want: 456,
					},
				}
				for _, tt := range tests {
					t.Run(tt.name, func(t *testing.T) {
						if got := ParseOr(tt.args.s, tt.args.def...); !reflect.DeepEqual(got, tt.want) {
							t.Errorf("ParseOr() = %v, want %v", got, tt.want)
						}
					})
				}
			},
		},
		{
			name: "string",
			run: func(t *testing.T) {
				tests := []testCase[string]{
					{
						name: "valid",
						args: args[string]{s: "abc"},
						want: "abc",
					},
				}
				for _, tt := range tests {
					t.Run(tt.name, func(t *testing.T) {
						if got := ParseOr(tt.args.s, tt.args.def...); !reflect.DeepEqual(got, tt.want) {
							t.Errorf("ParseOr() = %v, want %v", got, tt.want)
						}
					})
				}
			},
		},
		{
			name: "struct",
			run: func(t *testing.T) {
				type myStruct struct {
					A int
					B string
				}
				tests := []testCase[myStruct]{
					{
						name: "valid",
						args: args[myStruct]{s: `{"A":1,"B":"b"}`},
						want: myStruct{A: 1, B: "b"},
					},
					{
						name: "invalid with default",
						args: args[myStruct]{s: `{"A":1,"B":"b"`, def: []myStruct{{A: 2, B: "c"}}},
						want: myStruct{A: 2, B: "c"},
					},
				}
				for _, tt := range tests {
					t.Run(tt.name, func(t *testing.T) {
						if got := ParseOr(tt.args.s, tt.args.def...); !reflect.DeepEqual(got, tt.want) {
							t.Errorf("ParseOr() = %v, want %v", got, tt.want)
						}
					})
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}
