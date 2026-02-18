package arrays

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(str); i++ {

		// Ignore spaces completely
		if str[i] == ' ' {
			continue
		}

		count++

		// If it's the 6th character, skip it
		if count == 6 {
			count = 0
			continue
		}

		result += string(str[i])

		// After 5 valid characters, add space
		if count == 5 {
			result += " "
		}
	}

	// Remove trailing space if exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
