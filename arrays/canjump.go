package arrays

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}

	pos := 0
	for pos < len(nums)-1 {
		if nums[pos] == 0 {
			return false
		}
		next := int(pos) + int(nums[pos])
		if next >= len(nums)-1 {
			return true
		}
		pos = next
	}
	return true
}
