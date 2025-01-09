package easy

import "strings"

// Time: O(n)
// Space: O(1)
func prefixCount(words []string, pref string) int {
	var count int
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}

	return count
}
