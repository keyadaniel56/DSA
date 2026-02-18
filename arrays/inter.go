package arrays

import (
	"fmt"
	"os"
)

// inter returns a string of characters that appear in both s1 and s2,
// without duplicates, in the order they appear in s1.
func inter(s1, s2 string) string {
	seen := make(map[rune]bool)
	result := ""

	for _, c := range s1 {
		if !seen[c] && contains(s2, c) {
			result += string(c)
			seen[c] = true
		}
	}

	return result
}

// Helper: check if rune exists in string
func contains(s string, c rune) bool {
	for _, x := range s {
		if x == c {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	fmt.Println(inter(s1, s2))
}
