package medium

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			combs = append(combs, candidates[i])
			backtrack(combs, sum+candidates[i], i+1)
			combs = combs[:len(combs)-1]
		}
	}

	sort.Ints(candidates)
	backtrack([]int{}, 0, 0)
	return result
}
