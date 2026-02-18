package arrays

func Wdmatch(s1, s2 string) string {
	// If s1 is empty, it is always matchable
	if len(s1) == 0 {
		return s1
	}

	i := 0

	for _, char := range s2 {
		if i < len(s1) && char == rune(s1[i]) {
			i++
		}
	}

	if i == len(s1) {
		return s1
	}

	return ""
}
