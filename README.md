![build](https://img.shields.io/github/workflow/status/isacikgoz/slices/Test) ![goversion](https://img.shields.io/github/go-mod/go-version/isacikgoz/slices) [![Go Reference](https://pkg.go.dev/badge/github.com/isacikgoz/slices.svg)](https://pkg.go.dev/github.com/isacikgoz/slices) ![coverage](https://img.shields.io/codecov/c/github/isacikgoz/slices)

# Slices

`slices` package provides some utility functions on Go slice types. It leverages the generic types hence requires Go 1.18 and above.

## Download

You can simply run `go get github.com/isacikgoz/slices` to start using in your own code.

## Examples

Delete with preserving the order of the slice

```Go
s := []int{0, 1, 2, 3}
s = slices.Delete(s, 0) // removes the first element from the slice
```

Insert another slice from an index

```Go
s1 := []int{0, 3, 4, 5}
s2 := []int{1, 2}
s1 = slices.Insert(s1, 1, s2...) // will have (0, 1, 2, 3, 4, 5)
```

## License

[BSD-3-Clause](/LICENSE)
