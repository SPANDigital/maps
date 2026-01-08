package maps

import (
	stdmaps "maps"
	"slices"
)

// Keys returns all keys from the map as a slice.
//
// Deprecated: Use slices.Collect(maps.Keys(m)) from the standard library instead.
// This function will be removed in a future version.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return slices.Collect(stdmaps.Keys(m))
}
