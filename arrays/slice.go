package arrays

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	length := len(a)
	start := nbrs[0]

	// Handle negative start
	if start < 0 {
		start = length + start
	}

	// If only one number → slice to end
	end := length
	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end = length + end
		}
	}

	// Bounds check
	if start < 0 || start > length || end < 0 || end > length || start >= end {
		return nil
	}

	result := []string{}
	for i := start; i < end; i++ {
		result = append(result, a[i])
	}

	return result
}
