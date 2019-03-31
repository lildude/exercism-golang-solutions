// Package accumulate provides a method that performs a provided operation on all elements of a collection.
package accumulate

// Accumulate takes a collection and an operation to perform on each element of the collection and returns a new
// collection containing the result of applying that operation to each element of the input collection.
func Accumulate(c []string, f func(string) string) []string {
	nc := []string{}
	for _, e := range c {
		nc = append(nc, f(e))
	}
	return nc
}
