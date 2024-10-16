package easy

// Time complexity: O(logn)
// Space complexity: O(1)
func guessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		m := (l + r) / 2
		ans := guess(m)
		if ans == 0 {
			return m
		}
		if ans == 1 {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func guess(m int) int {
	panic("unimplemented")
}
