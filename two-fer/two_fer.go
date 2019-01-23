// Package twofer implements a simple function for two-for-one sharing.
package twofer

// ShareWith returns a string showing who we're sharing with.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
