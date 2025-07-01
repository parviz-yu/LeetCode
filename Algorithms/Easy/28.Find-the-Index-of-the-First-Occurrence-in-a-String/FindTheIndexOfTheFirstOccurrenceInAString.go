package easy

// Time: O(m*n)
// Space: O(1)
func strStr(haystack string, needle string) int {
	start := 0
	end := start + len(needle)

	for end <= len(haystack) {
		if haystack[start:end] == needle {
			return start
		}
		start++
		end++
	}

	return -1
}
