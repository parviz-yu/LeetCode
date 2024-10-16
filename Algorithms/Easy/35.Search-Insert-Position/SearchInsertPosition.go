package easy

// Time complexity: O(logn)
// Space complexity: O(1)
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if target <= nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}

	if target > nums[l] {
		l += 1
	}

	return l
}
