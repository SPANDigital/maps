package maps

func Inverse[K comparable, V comparable](in map[K]V) map[V]K {
	if in == nil {
		return nil
	}
	out := make(map[V]K)
	for k, v := range in {
		out[v] = k
	}
	return out
}
