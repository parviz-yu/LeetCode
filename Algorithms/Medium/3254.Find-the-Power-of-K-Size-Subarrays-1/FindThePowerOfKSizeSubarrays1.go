package medium

// Time: O(n)
// Space: O(1)
func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	result := make([]int, 0)
	var windowState, begin int
	prev := nums[0] - 1

	for end := range nums {
		if nums[end]-prev == 1 {
			windowState++
		}

		prev = nums[end]

		if end-begin+1 == k {
			if windowState == k {
				result = append(result, prev)
			} else {
				result = append(result, -1)
			}

			if nums[begin]+1 == nums[begin+1] {
				windowState--
			}

			begin++
		}
	}

	return result
}
