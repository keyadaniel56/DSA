package arrays

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	digits := "0123456789ABCDEF"

	// Handle 0 directly
	if value == 0 {
		return "0"
	}

	result := ""
	negative := false

	// Handle negative values
	if value < 0 {
		negative = true
		value = -value
	}

	// Convert number
	for value > 0 {
		rem := value % base
		result = string(digits[rem]) + result
		value = value / base
	}

	if negative {
		result = "-" + result
	}

	return result
}
