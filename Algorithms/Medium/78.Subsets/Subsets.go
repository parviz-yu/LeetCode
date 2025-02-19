package medium

// Time: O(2^n)
// Space: O(n)
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrack func(acc []int, start, end int)
	backtrack = func(acc []int, start, end int) {
		result = append(result, append([]int{}, acc...))

		if start == end {
			return
		}

		for i := start; i < end; i++ {
			acc = append(acc, nums[i])
			backtrack(acc, i+1, end)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]int{}, 0, len(nums))
	return result
}
