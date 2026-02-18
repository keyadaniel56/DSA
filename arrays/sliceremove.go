package arrays

func SliceRemove(slice []int, num int) []int {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == num {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
