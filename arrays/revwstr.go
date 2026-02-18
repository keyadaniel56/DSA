package arrays

import "strings"

func ReverseWords(s string) string {
	if s == "" {
		return "\n"
	}

	// Split words by space
	words := strings.Split(s, " ")
	reversed := ""

	// Reverse the slice
	for i := len(words) - 1; i >= 0; i-- {
		reversed += words[i]
		if i != 0 {
			reversed += " "
		}
	}

	return reversed + "\n"
}
