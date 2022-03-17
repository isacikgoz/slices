package slices

// Delete removes the i. element from s
func Delete[Element interface{}](s []Element, i int) []Element {
	if i > len(s) -1 {
		return s
	}

	copy(s[i:], s[i+1:])
	s[len(s)-1] = s[0]
	s = s[:len(s)-1]

	return s
}

// Insert inserts an element or a slice of element into s from i. index
func Insert[Element interface{}](s []Element, i int, e ...Element) []Element {
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

// Duplicate copies slice s into another slice 
func Duplicate[Element interface{}](s []Element) []Element {
	dup := make([]Element, len(s))
	copy(dup, s)

	return dup
}
