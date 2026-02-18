package arrays

import (
	"strconv"
	"strings"
)

func FindPairs(arrStr string, targetStr string) string {

	// Validate array format
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		return "Invalid input."
	}

	// Validate target
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		return "Invalid target sum."
	}

	// Remove brackets
	content := arrStr[1 : len(arrStr)-1]
	if content == "" {
		return "No pairs found."
	}

	parts := strings.Split(content, ",")
	arr := []int{}

	for _, p := range parts {
		p = strings.TrimSpace(p)

		num, err := strconv.Atoi(p)
		if err != nil {
			return "Invalid number: " + p
		}

		arr = append(arr, num)
	}

	// Find pairs
	pairs := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		return "No pairs found."
	}

	return "Pairs with sum " + targetStr + ": " + formatPairs(pairs)
}

func formatPairs(pairs [][]int) string {
	result := "["

	for i, pair := range pairs {
		result += "[" + strconv.Itoa(pair[0]) + " " + strconv.Itoa(pair[1]) + "]"
		if i != len(pairs)-1 {
			result += " "
		}
	}

	result += "]"
	return result
}
