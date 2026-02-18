// package arrays

// func ConcatAlternate(slice1, slice2 []int) []int {
// 	first, second := slice1, slice2
// 	if len(slice2) > len(slice1) {
// 		first, second = slice1, slice2
// 	}
// 	result := []int{}
// 	minLen := len(first)
// 	if len(second) < minLen {
// 		minLen = len(second)
// 	}

// 	for i := 0; i < minLen; i++ {
// 		result = append(result, first[i], second[i])
// 	}

// 	if len(first) > minLen {
// 		result = append(result, first[minLen:]...)
// 	}
// 	return result
// }

package arrays

func ConcatAlternate(slice1, slice2 []int) []int {
	first, second := slice1, slice2
	if len(slice2) > len(slice1) {
		first, second = slice1, slice2
	}
	result := []int{}
	minLen := len(first)
	if len(second) < minLen {
		minLen = len(second)
	}
	for i := 0; i < minLen; i++ {
		result = append(result, first[i], second[i])
	}
	if len(first) > minLen {
		result = append(result, first[minLen:]...)
	}
	return result
}
