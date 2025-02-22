package medium

// Time: O(n * 2^n)
// Space: O(n)
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func(combs []int, sum, idx int)
	backtrack = func(combs []int, sum, idx int) {
		if sum > target {
			return
		}

		if sum == target {
			result = append(result, append([]int{}, combs...))
			return
		}

		for i := idx; i < len(candidates); i++ {
			combs = append(combs, candidates[i])
			backtrack(combs, sum+candidates[i], i)
			combs = combs[:len(combs)-1]
		}
	}

	backtrack([]int{}, 0, 0)
	return result
}
