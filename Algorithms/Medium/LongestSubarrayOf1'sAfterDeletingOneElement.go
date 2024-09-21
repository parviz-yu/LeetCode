package medium

// Time: O(n)
// Space: O(1)
func longestSubarray(nums []int) int {
	var begin, zeros, maxSubLen int

	for end := range nums {
		if nums[end] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[begin] == 0 {
				zeros--
			}
			begin++
		}

		maxSubLen = max(maxSubLen, end-begin)
	}

	return maxSubLen
}
