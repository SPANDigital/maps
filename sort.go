package maps

import (
	"cmp"
	"slices"
)

func sort[T cmp.Ordered](in []T) []T {
	ret := in
	slices.Sort(ret)
	return ret
}
