package maps

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	type args[K comparable, V any, M interface{ ~map[K]V }] struct {
		m M
	}
	type testCase[K comparable, V any, M interface{ ~map[K]V }] struct {
		name string
		args args[K, V, M]
		want []K
	}
	tests := []testCase[int, string, map[int]string]{
		{
			name: "a",
			args: args[int, string, map[int]string]{
				m: map[int]string{1: "one", 2: "two"},
			},
			want: []int{1, 2},
		},
		{
			name: "b",
			args: args[int, string, map[int]string]{
				m: map[int]string{1: "one", 2: "two", 3: "three"},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
