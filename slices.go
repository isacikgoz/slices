package slices

import (
	"math/rand"
)

// Delete removes the i. element from s.
func Delete[T any](s []T, i int) []T {
	if i > len(s)-1 {
		return s
	}

	var e T
	copy(s[i:], s[i+1:])
	s[len(s)-1] = e
	s = s[:len(s)-1]

	return s
}

// Insert inserts a T or a slice of T into s from i. index.
func Insert[T any](s []T, i int, e ...T) []T {
	if i > len(s)-1 && len(s) != 0 {
		i = len(s) - 1
	} else if len(s) == 0 {
		i = 0
	} else if i < 0 {
		s = append(e, s...)
	}

	s = append(s[:i], append(e, s[i:]...)...)
	return s
}

// Duplicate copies slice s into another slice.
func Duplicate[T any](s []T) []T {
	dup := make([]T, len(s))
	copy(dup, s)

	return dup
}

// Push adds an element at the end of s.
func Push[T any](s []T, e T) []T {
	return append(s, e)
}

// Pop returns the last element from s. Returns zero value of T of s is empty.
func Pop[T any](s []T) (T, []T) {
	if len(s) == 0 {
		var e T
		return e, s
	}

	return s[len(s)-1], s[:len(s)-1]
}

// Shuffle pseudo-randomizes the order of elements of the s.
func Shuffle[T any](s []T) []T {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

// Reverse reverses the order of s.
func Reverse[T any](s []T) []T {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}

	return s
}

// Unique removes duplicate values in a slice
func Unique[T comparable](src []T) []T {
	var result []T
	var elemMap = make(map[T]struct{})
	for i := range src {
		if _, ok := elemMap[src[i]]; !ok {
			result = append(result, src[i])
			elemMap[src[i]] = struct{}{}
		}
	}

	return result
}

// Filter filters the elements of s according to the boolean value of the predicate.
func Filter[T any](s []T, f func(T) bool) []T {
	var n []T
	for _, e := range s {
		if f(e) {
			n = append(n, e)
		}
	}

	return n
}

// Diff returns the values from src that are not present in other slices
func Diff[T comparable](src []T, arrays ...[]T) []T {
	var elemMap = make(map[T]struct{})
	for _, array := range arrays {
		for i := range array {
			elemMap[array[i]] = struct{}{}
		}
	}
	var result = make([]T, 0, len(src))
	for i := range src {
		if _, ok := elemMap[src[i]]; !ok {
			result = append(result, src[i])
		}
	}
	return result
}
