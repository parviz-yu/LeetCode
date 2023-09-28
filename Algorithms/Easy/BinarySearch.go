package easy

// Time complexity: O(n)
// Space complexity: O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
