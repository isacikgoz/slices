package slices

import (
	"math/rand"
)

// Delete removes the i. element from s.
func Delete[Element any](s []Element, i int) []Element {
	if i > len(s) -1 {
		return s
	}

	copy(s[i:], s[i+1:])
	s[len(s)-1] = s[0]
	s = s[:len(s)-1]

	return s
}

// Insert inserts an element or a slice of element into s from i. index.
func Insert[Element any](s []Element, i int, e ...Element) []Element {
	if i > len(s) -1 && len(s) != 0{
		i = len(s) -1
	} else if len(s) == 0 {
		i = 0
	} else if i < 0 {
		s = append(e, s...)
	}

	s = append(s[:i], append(e, s[i:]...)...)
	return s
}

// Duplicate copies slice s into another slice.
func Duplicate[Element any](s []Element) []Element {
	dup := make([]Element, len(s))
	copy(dup, s)

	return dup
}

// Push adds an element at the end of s.
func Push[Element any](s []Element, e Element) []Element {
	return append(s, e)
}

// Pop returns the last element from s. Returns zero value of Element of s is empty.
func Pop[Element any](s []Element) (Element, []Element) {
	if len(s) == 0 {
		var e Element
		return e, s
	}

	return s[len(s)-1], s[:len(s)-1]
}

// Shuffle pseudo-randomizes the order of elements of the s.
func Shuffle[Element any](s []Element) []Element {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

// Reverse reverses the order of s.
func Reverse[Element any](s []Element) []Element {
	for i := len(s)/2-1; i >= 0; i-- {
		opp := len(s)-1-i
		s[i], s[opp] = s[opp], s[i]
	}

	return s
}
