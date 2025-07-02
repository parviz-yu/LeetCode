package medium

// Time: O(log n)
// Space: O(1)
func searchRange(nums []int, target int) []int {
	firstIdx := leftBinarySearch(nums, target)
	lastIdx := rightBinarySearch(nums, target)
	return []int{firstIdx, lastIdx}
}

func leftBinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2

		if target <= nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}

	if nums[l] != target {
		return -1
	}

	return l
}

func rightBinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r + 1) / 2

		if target >= nums[m] {
			l = m
		} else {
			r = m - 1
		}
	}

	if nums[l] != target {
		return -1
	}

	return l
}
