// Package word provides word counting functionalities.
package word

import "strings"

// UseCount returns how many times a word occurs in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string separated by white spaces.
func Count(s string) int {
	return len(strings.Fields(s))
}
