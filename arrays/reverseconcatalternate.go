package arrays

func RevConcatAlternate(slice1, slice2 []int) []int {
	// Helper to reverse a slice
	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	// Make copies to avoid modifying originals
	s1 := append([]int(nil), slice1...)
	s2 := append([]int(nil), slice2...)
	reverse(s1)
	reverse(s2)

	result := []int{}

	i, j := 0, 0

	// Determine which slice is "first" based on length
	firstLonger := len(s1) >= len(s2)

	for i < len(s1) || j < len(s2) {
		// If both slices still have elements
		if i < len(s1) && j < len(s2) {
			if firstLonger {
				result = append(result, s1[i], s2[j])
			} else {
				result = append(result, s2[j], s1[i])
			}
			i++
			j++
		} else if i < len(s1) {
			result = append(result, s1[i])
			i++
		} else {
			result = append(result, s2[j])
			j++
		}
		// Once the remaining parts are equal length, start with first slice
		if (len(s1) - i) == (len(s2) - j) {
			firstLonger = true
		}
	}

	return result
}
