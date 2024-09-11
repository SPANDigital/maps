package maps

import (
	"reflect"
	"slices"
	"testing"
)

func TestValues(t *testing.T) {
	type args[K comparable, V any, M interface{ ~map[K]V }] struct {
		m M
	}
	type testCase[K comparable, V any, M interface{ ~map[K]V }] struct {
		name string
		args args[K, V, M]
		want []V
	}
	tests := []testCase[int, string, map[int]string]{
		{
			name: "a",
			args: args[int, string, map[int]string]{
				m: map[int]string{1: "one", 2: "two"},
			},
			want: []string{"one", "two"},
		},
		{
			name: "b",
			args: args[int, string, map[int]string]{
				m: map[int]string{1: "alpha", 2: "beta", 3: "delta"},
			},
			want: []string{"alpha", "beta", "delta"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Values(tt.args.m)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
