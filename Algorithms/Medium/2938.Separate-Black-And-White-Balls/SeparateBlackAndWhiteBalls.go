package medium

// Time: O(n)
// Space: O(1)
func minimumSteps(s string) int64 {
	var count int64
	var zeroes int64
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == 48 {
			zeroes++
		}

		if s[i] == 49 {
			count += zeroes
		}
	}

	return count
}
