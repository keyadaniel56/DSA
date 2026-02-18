package arrays

// Union returns the characters that appear in either s1 or s2, without duplicates, in order.
func Union(s1, s2 string) string {
	seen := make(map[rune]bool)
	result := ""

	// Add unique characters from the first string
	for _, char := range s1 {
		if !seen[char] {
			result += string(char)
			seen[char] = true
		}
	}

	// Add unique characters from the second string
	for _, char := range s2 {
		if !seen[char] {
			result += string(char)
			seen[char] = true
		}
	}

	return result
}
