package arrays

import "strings"

func WordFlip(str string) string {
	// Trim leading/trailing spaces
	str = strings.TrimSpace(str)

	// Check for empty string
	if str == "" {
		return "Invalid Output\n"
	}

	// Split by any number of spaces
	words := strings.Fields(str)

	// Reverse words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join back into string
	return strings.Join(words, " ") + "\n"
}
