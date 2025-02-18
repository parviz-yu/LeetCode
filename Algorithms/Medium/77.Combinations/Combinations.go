package medium

// Time: O(k ∗ C(n,k))
// Space: O(k)
func combine(n int, k int) [][]int {
	result := make([][]int, 0, n)

	var backtrack func(acc []int, next int)
	backtrack = func(acc []int, next int) {
		if len(acc) == k {
			result = append(result, append([]int{}, acc...))
			return
		}

		for i := next; i < n+1; i++ {
			acc = append(acc, i)
			backtrack(acc, i+1)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]int{}, 1)
	return result
}
