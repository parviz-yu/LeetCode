package easy

// Time complexity: O(logn)
// Space complexity: O(1)
func arrangeCoins(n int) int {
	l, r := 1, n
	for l < r {
		m := (l + r + 1) / 2
		if (m+1)*m/2 <= n {
			l = m
		} else {
			r = m - 1
		}
	}

	return l
}
