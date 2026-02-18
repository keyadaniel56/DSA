package arrays

func Hidenp(s1, s2 string) string {
	if len(s1) == 0 {
		return "1"
	}

	i := 0
	for _, char := range s2 {
		if char == rune(s1[i]) {
			i++
			if i == len(s1) {
				return "1"
			}
		}
	}
	return "0"
}
