package medium

// Time: O(n)
// Space: O(n)
func maxCount(banned []int, n int, maxSum int) int {
	seen := make(map[int]bool, len(banned))
	for _, b := range banned {
		seen[b] = true
	}

	var count int
	var acc int
	for i := 1; i <= n; i++ {
		if !seen[i] && acc+i <= maxSum {
			acc += i
			count++
		}
	}

	return count
}
