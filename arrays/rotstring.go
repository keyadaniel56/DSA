package arrays

import "strings"

func RoString(s string) string {
	// Trim leading/trailing spaces
	s = strings.TrimSpace(s)
	if s == "" {
		return "\n"
	}

	// Split words by any number of spaces
	words := strings.Fields(s)
	if len(words) <= 1 {
		return words[0] + "\n"
	}

	// Rotate first word to the end
	rotated := strings.Join(append(words[1:], words[0]), " ")

	return rotated + "\n"
}
