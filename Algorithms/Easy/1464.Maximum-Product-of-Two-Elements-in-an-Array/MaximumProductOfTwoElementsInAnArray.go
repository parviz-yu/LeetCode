package easy

import "slices"

// Time: O(nlogn)
// Space: O(n)
func maxProduct(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	return (nums[n-1] - 1) * (nums[n-2] - 1)
}

// Time: O(n)
// Space: O(1)
func maxProduct_2(nums []int) int {
	var max1, max2 int
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max2 = num
		}
	}

	return (max1 - 1) * (max2 - 1)
}
