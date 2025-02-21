package medium

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var backtrack func(combs []int, sum int, start int)
	backtrack = func(combs []int, sum int, start int) {
		if len(combs) > k {
			return
		}

		if sum > n {
			return
		}

		if len(combs) == k && sum == n {
			result = append(result, append([]int{}, combs...))
			return
		}

		for i := start; i < 10; i++ {
			combs = append(combs, i)
			backtrack(combs, sum+i, i+1)
			combs = combs[:len(combs)-1]
		}
	}

	backtrack([]int{}, 0, 1)
	return result
}
