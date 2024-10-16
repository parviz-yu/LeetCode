package easy

// Time complexity: O(n)
// Space complexity: O(1)
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i == len(s)
}
