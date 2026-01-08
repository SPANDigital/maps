package maps

import (
	stdmaps "maps"
	"reflect"
	"slices"
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
			got := Keys(tt.args.m)
			slices.Sort(got)
			want := tt.want
			slices.Sort(want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Keys() = %v, want %v", got, want)
			}
		})
	}
}

// Example showing how to use the standard library maps.Keys with slices.Collect.
// This is the recommended approach going forward.
func ExampleKeys_standardLibrary() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}

	// New way: use standard library
	keys := slices.Collect(stdmaps.Keys(m))
	slices.Sort(keys)

	// Output can't be guaranteed due to map iteration order,
	// but this demonstrates the pattern
	_ = keys
}
