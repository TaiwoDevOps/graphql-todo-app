package utils

import "strings"

// Normalize normalizes a string by converting it to lowercase and trimming whitespace
func Normalize(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}
