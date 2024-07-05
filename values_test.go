package maps

import (
	"reflect"
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
				m: map[int]string{1: "one", 2: "two", 3: "three"},
			},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(Values(tt.args.m)); !reflect.DeepEqual(got, sort(tt.want)) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
