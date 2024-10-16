package medium

// Time: O(n)
// Space: O(1)
func longestOnes(nums []int, k int) int {
	var begin, windowState int
	var maxOnes int

	for end := range nums {
		if nums[end] == 0 {
			windowState++
		}

		for windowState > k {
			if nums[begin] == 0 {
				windowState--
			}
			begin++
		}

		maxOnes = max(maxOnes, end-begin+1)
	}

	return maxOnes
}
