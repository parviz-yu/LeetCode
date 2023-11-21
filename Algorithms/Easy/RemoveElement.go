package easy

// Time complexity: O(n)
// Space complexity: O(1)
func removeElement(nums []int, val int) int {
	n := 0
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			if nums[r] != val {
				nums[r], nums[l] = nums[l], nums[r]
				l++
			}
			n++
			r--
		} else {
			l++
		}
	}

	return len(nums) - n
}
