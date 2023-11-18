package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, helpers.Sum(nums)
	for i, val := range nums {
		rightSum -= val
		if leftSum == rightSum {
			return i
		} else {
			leftSum += val
		}
	}

	return -1
}
