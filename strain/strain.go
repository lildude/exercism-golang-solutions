// Package strain implements `keep` and `discard` operation on collections.
package strain

// Ints is a collection of integers.
type Ints []int

// Lists is a collection of integer slices.
type Lists [][]int

// Strings is a collection of strings.
type Strings []string

// Keep creates a new collection of ints that passes some test.
func (ints Ints) Keep(f func(int) bool) Ints {
	var kept Ints
	for _, i := range ints {
		if f(i) {
			kept = append(kept, i)
		}
	}
	return kept
}

// Discard creates a new collection of ints that don't pass some test.
func (ints Ints) Discard(f func(int) bool) Ints {
	return ints.Keep(func(i int) bool { return !f(i) })
}

// Keep creates a new collection of int slices that passes some test.
func (lists Lists) Keep(f func([]int) bool) Lists {
	var kept Lists
	for _, l := range lists {
		if f(l) {
			kept = append(kept, l)
		}
	}
	return kept
}

// Keep creates a new collection of int slices that passes some test.
func (strings Strings) Keep(f func(string) bool) Strings {
	var kept Strings
	for _, s := range strings {
		if f(s) {
			kept = append(kept, s)
		}
	}
	return kept
}
