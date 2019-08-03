// Package twofer contains a function that returns string
package twofer

import "fmt"

// ShareWith returns the string "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
