# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library providing generic utility functions for working with maps. The library requires Go 1.25+ and leverages Go generics to provide type-safe map operations.

## Core Utilities

- **Inverse**: Swaps keys and values (both must be comparable) - **Active**
- **Keys**: (DEPRECATED) Extracts keys from a map as a slice - use `slices.Collect(maps.Keys(m))` from stdlib instead
- **Values**: (DEPRECATED) Extracts values from a map as a slice - use `slices.Collect(maps.Values(m))` from stdlib instead

### Standard Library Migration

Go 1.25+ includes a `maps` package in the standard library with iterator-based `Keys()` and `Values()` functions. The functions in this library that overlap with the standard library are deprecated and will be removed in a future version.

## Development Commands

### Running Tests

```bash
# Run all tests
go test ./...

# Run a specific test
go test -run TestInverse
go test -run TestKeys
go test -run TestValues

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

### Building

```bash
# Build the package
go build ./...

# Verify all code compiles
go build -v ./...
```

### Linting

```bash
# Run go vet
go vet ./...

# Format code
go fmt ./...
```

## Code Architecture

### File Organization

Each utility function follows a strict one-function-per-file pattern:
- `<function>.go` - Implementation
- `<function>_test.go` - Table-driven tests

All code lives in the root `maps` package (no subdirectories).

### Generic Type Parameters

Functions use Go generics with constraints:
- Map keys must be `comparable`
- The `Keys` and `Values` functions use `~map[K]V` constraint to accept any map type
- `Inverse` requires both keys and values to be `comparable`

Example signature pattern:
```go
func Keys[M ~map[K]V, K comparable, V any](m M) []K
```

### Deprecated Functions

The `Keys` and `Values` functions are deprecated as of Go 1.25+. The standard library now provides these via iterators:

```go
// Old way (deprecated)
keys := maps.Keys(m)

// New way (recommended) - note the aliased import
import (
    stdmaps "maps"
    "slices"
)
keys := slices.Collect(stdmaps.Keys(m))
```

When working with this codebase, be aware that these functions use `stdmaps` as an alias for the standard library `maps` package to avoid naming conflicts with this package.

### Nil Handling

All functions handle `nil` maps gracefully:
- `Inverse(nil)` returns `nil`
- `Keys(nil)` and `Values(nil)` delegate to standard library (which returns empty iterators)

## Testing Conventions

### Table-Driven Tests

All tests use the table-driven pattern with:
- Generic `args` and `testCase` structs
- `t.Run()` for each test case
- `reflect.DeepEqual` for comparing results

### Unordered Comparisons

For `Keys` and `Values` tests, use `slices.Sort` before comparison since map iteration order is not guaranteed:

```go
slices.Sort(got)
slices.Sort(tt.want)
if !reflect.DeepEqual(got, tt.want) {
    t.Errorf("...")
}
```

### Test Case Naming

Use descriptive test case names like "empty", "foobar", "multiple" that indicate what scenario is being tested.
