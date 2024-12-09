package medium

// Time: O(m + n)
// Space: O(m)
func isArraySpecial(nums []int, queries [][]int) []bool {
	noSpecial := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		curr, prev := nums[i], nums[i-1]
		if curr%2 == prev%2 {
			noSpecial[i] = noSpecial[i-1] + 1
		} else {
			noSpecial[i] = noSpecial[i-1]
		}
	}

	result := make([]bool, len(queries))
	for i := range queries {
		start, end := queries[i][0], queries[i][1]
		if noSpecial[end]-noSpecial[start] == 0 {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
