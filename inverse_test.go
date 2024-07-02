package maps

import (
	"reflect"
	"testing"
)

func TestInverse(t *testing.T) {
	type args[K comparable, V comparable] struct {
		in map[K]V
	}
	type testCase[K comparable, V comparable] struct {
		name string
		args args[K, V]
		want map[V]K
	}
	tests := []testCase[string, string]{
		{
			name: "empty",
			args: args[string, string]{
				in: nil,
			},
			want: nil,
		},
		{
			name: "empty",
			args: args[string, string]{
				in: map[string]string{},
			},
			want: map[string]string{},
		},
		{
			name: "foobar",
			args: args[string, string]{
				in: map[string]string{
					"foo": "bar",
				},
			},
			want: map[string]string{
				"bar": "foo",
			},
		},
		{
			name: "multiple",
			args: args[string, string]{
				in: map[string]string{
					"foo": "bar",
					"bar": "baz",
				},
			},
			want: map[string]string{
				"bar": "foo",
				"baz": "bar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Inverse(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
