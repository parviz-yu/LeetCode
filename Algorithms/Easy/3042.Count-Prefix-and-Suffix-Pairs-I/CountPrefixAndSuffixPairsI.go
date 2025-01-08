package easy

import "strings"

// Time: O(n^2)
// Space: O(1)
func countPrefixSuffixPairs(words []string) int {
	var count int

	for i := range words {
		sub := words[i]
		for j := i + 1; j < len(words); j++ {
			s := words[j]
			if strings.HasPrefix(s, sub) && strings.HasSuffix(s, sub) {
				count++
			}
		}
	}

	return count
}
