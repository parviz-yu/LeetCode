package easy

import "math"

// Time: O(n)
// Space: O(1)
func findMaxAverage(nums []int, k int) float64 {
	var begin int
	var currSum int
	maxSum := math.MinInt

	for end := range nums {
		currSum += nums[end]
		if end-begin+1 == k {
			maxSum = max(maxSum, currSum)
			currSum -= nums[begin]
			begin++
		}
	}

	return float64(maxSum) / float64(k)
}
