package medium

import "sort"

// Time: O(n*log(n))
// Space: O(n)
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	var number int
	for l < r {
		if nums[l]+nums[r] == k {
			number++
			l++
			r--
		} else if nums[l]+nums[r] > k {
			r--
		} else {
			l++
		}
	}

	return number
}
