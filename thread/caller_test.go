//go:build !race

// Package thread implements the functions, types, and interfaces for the module.
package thread

import (
	"errors"
	"fmt"
	"testing"
)

func TestTry(t *testing.T) {
	type args[T any] struct {
		f       func() (T, error)
		catch   func(err error)
		finally func()
		then    func(T)
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want struct {
			catch   error
			then    T
			finally bool
		}
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				f: func() (int, error) {
					return 1, nil
				},
				catch:   nil,
				finally: nil,
				then: func(i int) {

				},
			},
			want: struct {
				catch   error
				then    int
				finally bool
			}{
				catch:   nil,
				then:    1,
				finally: false,
			},
		},
		{
			name: "test2",
			args: args[int]{
				f: func() (int, error) {
					return 0, errors.New("error")
				},
				catch: func(err error) {

				},
				finally: func() {

				},
			},
			want: struct {
				catch   error
				then    int
				finally bool
			}{
				catch:   errors.New("error"),
				then:    0,
				finally: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, act := Try(tt.args.f)

			if tt.args.catch != nil {
				got.Catch(func(err error) {
					fmt.Println("catch")
					if err != nil && err.Error() != tt.want.catch.Error() {
						t.Errorf("Try() error = %v, wantErr %v", err, tt.want.catch)
					}
					if err == nil && tt.want.catch != nil {
						t.Errorf("Try() error = %v, want %v", err, tt.want.then)
					}
				})
			}
			if tt.args.then != nil {
				got.Then(func(v int) {
					fmt.Println("then")
					if v != tt.want.then {
						t.Errorf("Try() then = %v, want %v", v, tt.want.then)
					}
				})
			}
			if tt.args.finally != nil {
				got.Finally(func() {
					fmt.Println("finally")
					if tt.args.finally != nil && !tt.want.finally {
						t.Errorf("Try() finally = %v, want %v", tt.args.finally != nil, tt.want.finally)
					}
					if tt.args.finally == nil && tt.want.finally {
						t.Errorf("Try() finally = %v, want %v", tt.args.finally == nil, tt.want.finally)
					}
				})
			}
			act()
			// time.Sleep(time.Millisecond)
		})
	}
}
