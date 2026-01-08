# maps

Go utilities for working with maps.

**Requires Go 1.25+**

## Functions

### Inverse
Swaps keys and values in a map. Both keys and values must be comparable.

```go
m := map[string]int{"a": 1, "b": 2}
inverted := maps.Inverse(m) // map[int]string{1: "a", 2: "b"}
```

### Keys (Deprecated)
**Deprecated:** Use `slices.Collect(maps.Keys(m))` from the standard library instead.

Retrieves keys from a map as a slice. The order of keys is not guaranteed.

```go
// Old way (deprecated)
keys := maps.Keys(m)

// New way (recommended)
keys := slices.Collect(maps.Keys(m))
```

### Values (Deprecated)
**Deprecated:** Use `slices.Collect(maps.Values(m))` from the standard library instead.

Retrieves values from a map as a slice. The order of values is not guaranteed.

```go
// Old way (deprecated)
values := maps.Values(m)

// New way (recommended)
values := slices.Collect(maps.Values(m))
```

## Migration Guide

Go 1.25+ includes a standard library `maps` package with iterator-based `Keys()` and `Values()` functions. To migrate:

1. Import both `maps` and `slices` from the standard library
2. Replace `maps.Keys(m)` with `slices.Collect(maps.Keys(m))`
3. Replace `maps.Values(m)` with `slices.Collect(maps.Values(m))`

The `Keys` and `Values` functions in this package will be removed in a future version.