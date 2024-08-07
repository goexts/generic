package maps

import (
	"testing"
)

func TestMapToKVs(t *testing.T) {
	type args struct {
		m map[int]string
	}
	type testCase struct {
		name string
		args args
		want []struct {
			Key int
			Val string
		}
	}
	tests := []testCase{
		{
			name: "test",
			args: args{
				m: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
			want: []struct {
				Key int
				Val string
			}{
				{
					Key: 1,
					Val: "a",
				},
				{
					Key: 2,
					Val: "b",
				},
				{
					Key: 3,
					Val: "c",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapToKVs(tt.args.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapToKVs() = %v, want %v", got, tt.want)
			}
			for i, v := range got {
				if v.Key != tt.want[i].Key || v.Val != tt.want[i].Val {
					t.Errorf("MapToKVs() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
