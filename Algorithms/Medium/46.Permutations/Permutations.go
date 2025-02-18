package medium

// Time: O(n*n!)
// Space: O(n)
func permute(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	var backtrack func(accum []int, seen map[int]bool)
	backtrack = func(accum []int, seen map[int]bool) {
		if len(accum) == len(nums) {
			result = append(result, append([]int{}, accum...))
			return
		}

		for _, num := range nums {
			if !seen[num] {
				seen[num] = true
				accum = append(accum, num)
				backtrack(accum, seen)
				accum = accum[:len(accum)-1]
				seen[num] = false
			}
		}
	}

	backtrack([]int{}, map[int]bool{})

	return result
}
