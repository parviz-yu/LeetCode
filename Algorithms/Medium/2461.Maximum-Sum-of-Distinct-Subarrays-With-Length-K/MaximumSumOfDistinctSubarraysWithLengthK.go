package medium

// Time: O(n)
// Space: O(n)
func maximumSubarraySum(nums []int, k int) int64 {
	var maxSum, currSum, begin int
	count := make(map[int]int)

	for end := range nums {
		currSum += nums[end]
		count[nums[end]]++

		if end-begin+1 == k {
			// all items in window distinct
			if len(count) == k {
				maxSum = max(maxSum, currSum)
			}
			if count[nums[begin]] == 1 {
				delete(count, nums[begin])
			} else {
				count[nums[begin]]--
			}

			currSum -= nums[begin]
			begin++
		}
	}

	return int64(maxSum)
}
