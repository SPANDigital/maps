# Copilot Instructions

## Codebase Structure

- Each utility function (e.g., `Keys`, `Values`, `Inverse`) is defined in its own `.go` file.
- Each function has a corresponding `_test.go` file with table-driven tests.

## Conventions

- Use Go generics for map utilities, with type parameters for key and value types.
- Functions should not assume any order for map keys or values.
- Return slices for keys/values, and maps for inverse operations.
- Handle `nil` maps gracefully (e.g., return `nil` for `Inverse(nil)`).
- Use Go's standard library for sorting in tests when comparing unordered slices.

## Patterns

- Table-driven tests: define test cases as structs, iterate with `t.Run`.
- Use `reflect.DeepEqual` for comparing slices and maps in tests.
- Use `slices.Sort` to compare unordered slices in tests.
- Keep function signatures generic and reusable.
- Document any assumptions or limitations in the README.

