package arrays

func Fprime(n int) string {
	if n <= 1 {
		return ""
	}

	result := ""
	first := true

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if !first {
				result += "*"
			}
			result += Itoa(i)
			n /= i
			first = false
		}
	}

	if n > 1 {
		if !first {
			result += "*"
		}
		result += Itoa(n)
	}

	return result
}
