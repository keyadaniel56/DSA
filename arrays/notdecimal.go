package arrays

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	sign := ""
	start := 0

	// Handle negative sign
	if dec[0] == '-' {
		sign = "-"
		start = 1
		if len(dec) == 1 {
			return dec + "\n"
		}
	}

	dotIndex := -1

	// Validate characters
	for i := start; i < len(dec); i++ {
		if dec[i] == '.' {
			if dotIndex != -1 {
				return dec + "\n" // multiple dots
			}
			dotIndex = i
		} else if dec[i] < '0' || dec[i] > '9' {
			return dec + "\n"
		}
	}

	// No decimal point
	if dotIndex == -1 {
		return dec + "\n"
	}

	// Check if decimal part is only zeros
	allZero := true
	for i := dotIndex + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZero = false
			break
		}
	}

	if allZero {
		return dec + "\n"
	}

	// Remove decimal point
	result := sign
	for i := start; i < len(dec); i++ {
		if dec[i] != '.' {
			result += string(dec[i])
		}
	}

	return result + "\n"
}
