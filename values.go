package maps

import (
	stdmaps "maps"
	"slices"
)

// Values returns all values from the map as a slice.
//
// Deprecated: Use slices.Collect(maps.Values(m)) from the standard library instead.
// This function will be removed in a future version.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return slices.Collect(stdmaps.Values(m))
}
