package medium

// Time complexity: O(n)
// Space complexity: O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] >= nums[0] {
			l = m + 1
		} else {
			r = m
		}
	}

	r = len(nums) - 1
	if target >= nums[0] {
		l, r = 0, l
	}

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
