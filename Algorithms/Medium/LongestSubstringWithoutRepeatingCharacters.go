package medium

// Time: O(n)
// Space: O(k), k = longest substring
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]struct{})
	var begin, longestSubStr int

	for end := range s {
		for _, ok := window[s[end]]; ok; _, ok = window[s[end]] {
			delete(window, s[begin])
			begin++
		}

		window[s[end]] = struct{}{}
		longestSubStr = max(longestSubStr, len(window))
	}

	return longestSubStr
}
